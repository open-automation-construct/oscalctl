package component

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"

	"github.com/open-automation-construct/stigctl/src/internal/cklb"
)

func GenerateComponent(inputPath, outputPath string) error {
	// Read and parse the input STIG checklist
	checklist, err := readSTIGChecklist(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read STIG checklist: %w", err)
	}

	// Generate OSCAL component
	component, err := createComponent(checklist)
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

func createComponent(checklist *cklb.Checklist) (*oscalTypes.ComponentDefinition, error) {
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
    controlImplementationSets := buildControlImplementationSets(checklist)
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
func buildControlImplementationSets(checklist *cklb.Checklist) []oscalTypes.ControlImplementationSet {
    implementationUUID := uuid.New().String()
    
    implementationSet := oscalTypes.ControlImplementationSet{
        UUID: implementationUUID,
        Source: "STIG",  // Using a placeholder source - This will probably need to be a URI fragment to UUID pointing to backmatter resource  https://pages.nist.gov/OSCAL-Reference/models/v1.1.3/component-definition/json-reference/#/component-definition/components/control-implementations/source
        Description: fmt.Sprintf("Control implementation for %s", checklist.Data.Title),
        ImplementedRequirements: []oscalTypes.ImplementedRequirementControlImplementation{},
    }

    // Process each STIG rule
    for _, stig := range checklist.Data.STIGs {
        for _, rule := range stig.Rules {
            reqUUID := uuid.New().String()
            
            // This matches the ImplementedRequirementControlImplementation struct definition
            requirement := oscalTypes.ImplementedRequirementControlImplementation{
                UUID:        reqUUID,
                ControlId:   "Placeholder - need to reference CCIs",
                Description: rule.RuleTitle,
				Remarks: rule.RuleID,
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