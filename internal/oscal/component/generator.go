package component

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"

	"github.com/open-automation-construct/oscalctl/internal/cciparsing"
	"github.com/open-automation-construct/oscalctl/internal/cklb"
)

func GenerateComponent(inputPath, outputPath, cciPath string) error {
	// Read and parse the input STIG checklist
	checklist, err := readSTIGChecklist(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read STIG checklist: %w", err)
	}

	// Parse CCI document - this will use the embedded one if cciPath is empty
	cciControlMap, err := cciparsing.ParseCCIDocument(cciPath)
	if err != nil {
		return fmt.Errorf("failed to parse CCI document: %w", err)
	}

	// Generate OSCAL component
	component, err := createComponent(checklist, cciControlMap)
	if err != nil {
		return fmt.Errorf("failed to create OSCAL component: %w", err)
	}

	// Output the component
	if err := writeComponent(component, outputPath); err != nil {
		return fmt.Errorf("failed to write OSCAL component: %w", err)
	}

	fmt.Printf("Generated OSCAL component at %s\n", outputPath)
	return nil
}

// readSTIGChecklist reads and parses a STIG checklist file
func readSTIGChecklist(path string) (*cklb.Checklist, error) {
	checklist := &cklb.Checklist{}
	
	// Try to load the checklist file
	err := checklist.LoadFromFile(path)
	if err != nil {
		return nil, fmt.Errorf("error loading checklist: %v", err)
	}
	
	// Validate the checklist
	isValid, errors := checklist.Validate()
	if !isValid {
		return nil, fmt.Errorf("checklist validation failed: %v", errors)
	}

	return checklist, nil
}

// extractCCINumbers extracts CCI identifiers from a rule
func extractCCINumbers(rule cklb.STIGRule) []string {
	var cciNumbers []string
	
	for _, cci := range rule.CCIs {
		// Ensure CCI is in the format we expect (e.g., "CCI-000001")
		if len(cci) > 0 {
			// You could add additional validation or formatting here if needed
			cciNumbers = append(cciNumbers, cci)
		}
	}
	
	return cciNumbers
}

func createComponent(checklist *cklb.Checklist, cciControlMap map[string]string) (*oscalTypes.ComponentDefinition, error) {
    // Generate UUIDs
    componentDefUUID := uuid.New().String()
    componentUUID := uuid.New().String()
    
    // Build metadata
    lastModified := time.Now()
    metadata := oscalTypes.Metadata{
        Title: checklist.Data.Title,
        LastModified: lastModified,
        Version: "1.0.0",
        OscalVersion: "1.1.3",
    }
    
    // Build the component
    definedComponent := oscalTypes.DefinedComponent{
        UUID: componentUUID,
        Type: "software",
        Title: checklist.Data.Title,
        Description: fmt.Sprintf("Component definition generated from STIG: %s", checklist.Data.Title),
    }
    
    // Set control implementation sets
    controlImplementationSets := buildControlImplementationSets(checklist, cciControlMap)
    definedComponent.ControlImplementations = &controlImplementationSets
    
    // Create the component definition
    component := &oscalTypes.ComponentDefinition{
        UUID: componentDefUUID,
        Metadata: metadata,
        Components: &[]oscalTypes.DefinedComponent{definedComponent},
    }

    return component, nil
}

// buildControlImplementationSets builds control implementation sets from STIG rules
func buildControlImplementationSets(checklist *cklb.Checklist, cciControlMap map[string]string) []oscalTypes.ControlImplementationSet {
    implementationUUID := uuid.New().String()
    
    implementationSet := oscalTypes.ControlImplementationSet{
        UUID: implementationUUID,
        Source: "https://raw.githubusercontent.com/usnistgov/oscal-content/main/nist.gov/SP800-53/rev5/json/NIST_SP-800-53_rev5_catalog.json", 
        Description: fmt.Sprintf("Control implementation for %s", checklist.Data.Title),
        ImplementedRequirements: []oscalTypes.ImplementedRequirementControlImplementation{},
    }

    // Process each STIG rule
    for _, stig := range checklist.Data.STIGs {
        for _, rule := range stig.Rules {
            reqUUID := uuid.New().String()
            
            // Extract CCIs from the rule
            cciNumbers := extractCCINumbers(rule)
            
            // Determine control ID from CCIs if available
            controlId := "unknown"
            if len(cciNumbers) > 0 && cciControlMap != nil {
                // Try to find a control ID for any of the CCIs
                for _, cci := range cciNumbers {
                    if control, exists := cciControlMap[cci]; exists && control != "" {
                        controlId = control
                        break
                    }
                }
            }
            
            // Add all CCIs to remarks
            remarks := rule.RuleID
            if len(cciNumbers) > 0 {
                remarks += " - CCIs: " + fmt.Sprintf("%v", cciNumbers)
            }
            
            // This matches the ImplementedRequirementControlImplementation struct definition
            requirement := oscalTypes.ImplementedRequirementControlImplementation{
                UUID:        reqUUID,
                ControlId:   controlId,
                Description: rule.RuleTitle,
                Remarks:     remarks,
            }
            
            implementationSet.ImplementedRequirements = append(
                implementationSet.ImplementedRequirements, 
                requirement,
            )
        }
    }

    return []oscalTypes.ControlImplementationSet{implementationSet}
}

// writeComponent writes the OSCAL component to a JSON file
func writeComponent(component *oscalTypes.ComponentDefinition, path string) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Marshal component to JSON
	jsonData, err := json.MarshalIndent(component, "", "  ")
	if err != nil {
		return err
	}

	// Write to file
	return os.WriteFile(path, jsonData, 0644)
}