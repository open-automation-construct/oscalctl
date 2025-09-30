package cklb

import (
	"fmt"
)

// Validate checks if the checklist is valid according to the schema
func (c *Checklist) Validate() (bool, []string) {
	// Implement validation logic
	var errors []string
	
	// Basic validation
	if c.Data.Title == "" {
		errors = append(errors, "Missing title")
	}
	if c.Data.ID == "" {
		errors = append(errors, "Missing ID")
	}
	
	// Validate STIGs
	if len(c.Data.STIGs) > 0 {
		for i, stig := range c.Data.STIGs {
			// Check required STIG fields
			if stig.STIGName == "" {
				errors = append(errors, fmt.Sprintf("STIG[%d]: Missing stig_name", i))
			}
			if stig.DisplayName == "" {
				errors = append(errors, fmt.Sprintf("STIG[%d]: Missing display_name", i))
			}
			if stig.STIGID == "" {
				errors = append(errors, fmt.Sprintf("STIG[%d]: Missing stig_id", i))
			}
			if stig.ReleaseInfo == "" {
				errors = append(errors, fmt.Sprintf("STIG[%d]: Missing release_info", i))
			}
			if stig.UUID == "" {
				errors = append(errors, fmt.Sprintf("STIG[%d]: Missing uuid", i))
			}
			if stig.Size == 0 {
				errors = append(errors, fmt.Sprintf("STIG[%d]: Missing or zero size", i))
			}
			
			// Validate rules
			for j, rule := range stig.Rules {
				if rule.UUID == "" {
					errors = append(errors, fmt.Sprintf("STIG[%d].Rule[%d]: Missing uuid", i, j))
				}
				if rule.RuleID == "" {
					errors = append(errors, fmt.Sprintf("STIG[%d].Rule[%d]: Missing rule_id", i, j))
				}
				
				// Validate status if present
				if rule.Status != "" {
					validStatus := []string{"not_reviewed", "not_applicable", "open", "not_a_finding"}
					isValid := false
					for _, status := range validStatus {
						if rule.Status == status {
							isValid = true
							break
						}
					}
					if !isValid {
						errors = append(errors, fmt.Sprintf("STIG[%d].Rule[%d]: Invalid status '%s'", i, j, rule.Status))
					}
				}
			}
		}
	}
	
	return len(errors) == 0, errors
}