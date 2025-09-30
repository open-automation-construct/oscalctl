package cciparsing

import (
	"embed"
	"encoding/xml"
	"io"
	"os"
	"regexp"
	"strings"
)

//go:embed assets/cci_list.xml
var embeddedFS embed.FS

// CCIDocument represents the structure of the CCI XML document
type CCIDocument struct {
	XMLName  xml.Name  `xml:"cci_list"`
	CCIItems []CCIItem `xml:"cci_items>cci_item"`
}

// CCIItem represents a single CCI item in the document
type CCIItem struct {
	ID         string     `xml:"id,attr"`
	Status     string     `xml:"status"`
	PublishDate string    `xml:"publishdate"`
	Contributor string    `xml:"contributor"`
	Definition string     `xml:"definition"`
	Type       string     `xml:"type"`
	References References `xml:"references"`
}

// References contains a list of reference elements
type References struct {
	References []Reference `xml:"reference"`
}

// Reference represents a single reference in a CCI item
type Reference struct {
	Creator   string `xml:"creator,attr"`
	Title     string `xml:"title,attr"`
	Version   string `xml:"version,attr"`
	Location  string `xml:"location,attr"`
	Index     string `xml:"index,attr"`
}

// normalizeControlID converts control IDs like 'SA-4 (7) (a)' to OSCAL format like 'sa-4.7'
func normalizeControlID(index string) string {
	// Match base control and enhancement
	re := regexp.MustCompile(`([A-Z]+-\d+)\s*(?:\((\d+)\))?`)
	matches := re.FindStringSubmatch(index)
	
	if len(matches) < 2 {
		return ""
	}
	
	base := strings.ToLower(matches[1])
	
	if len(matches) > 2 && matches[2] != "" {
		return base + "." + matches[2]
	}
	
	return base
}

// getPreferredControlID finds the most relevant control ID based on preference order
func getPreferredControlID(references []Reference) string {
	// Define preference order for references
	preferenceOrder := []string{
		"NIST SP 800-53 Revision 5",
		"NIST SP 800-53 Revision 4",
		"NIST SP 800-53",
		"NIST SP 800-53A",
	}
	
	var selectedRef *Reference
	
	// Find the highest preference reference
	for _, pref := range preferenceOrder {
		for i := range references {
			if references[i].Title == pref {
				selectedRef = &references[i]
				break
			}
		}
		if selectedRef != nil {
			break
		}
	}
	
	// If a reference was found, normalize the index to OSCAL format
	if selectedRef != nil {
		return normalizeControlID(selectedRef.Index)
	}
	
	return ""
}

// IsValidOSCALToken checks if a string conforms to the OSCAL token pattern
// pattern: "^(\\p{L}|_)(\\p{L}|\\p{N}|[.\\-_])*$"
func IsValidOSCALToken(s string) bool {
	if s == "" {
		return false
	}
	
	// First character must be a letter or underscore
	firstChar := []rune(s)[0]
	if !((firstChar >= 'a' && firstChar <= 'z') || 
	     (firstChar >= 'A' && firstChar <= 'Z') || 
	     firstChar == '_') {
		return false
	}
	
	// Rest of characters must be letters, numbers, dots, hyphens, or underscores
	for _, c := range s[1:] {
		if !((c >= 'a' && c <= 'z') || 
		     (c >= 'A' && c <= 'Z') || 
		     (c >= '0' && c <= '9') || 
		     c == '.' || c == '-' || c == '_') {
			return false
		}
	}
	
	return true
}

// GetEmbeddedCCIControlMap parses the embedded CCI XML document
func GetEmbeddedCCIControlMap() (map[string]string, error) {
	file, err := embeddedFS.Open("assets/cci_list.xml")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	return ParseCCIDocumentReader(file)
}

// ParseCCIDocument parses a CCI XML document from a file path and returns a map of CCI IDs to control IDs
// If filePath is empty, uses the embedded CCI document
func ParseCCIDocument(filePath string) (map[string]string, error) {
	// If no file path is provided, use the embedded CCI list
	if filePath == "" {
		return GetEmbeddedCCIControlMap()
	}
	
	// Otherwise, read from the specified file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	return ParseCCIDocumentReader(file)
}

// ParseCCIDocumentReader parses a CCI XML document from a reader
func ParseCCIDocumentReader(reader io.Reader) (map[string]string, error) {
	var document CCIDocument
	decoder := xml.NewDecoder(reader)
	
	if err := decoder.Decode(&document); err != nil {
		return nil, err
	}
	
	result := make(map[string]string)
	
	for _, item := range document.CCIItems {
		controlID := getPreferredControlID(item.References.References)
		
		// Only include control IDs that match the OSCAL token pattern
		if controlID != "" && IsValidOSCALToken(controlID) {
			result[item.ID] = controlID
		}
	}
	
	return result, nil
}