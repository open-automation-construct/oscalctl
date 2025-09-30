package cklb

import (
	"encoding/json"
	"fmt"
	"os"
)

// Checklist implements the ChecklistInterface
type Checklist struct {
	Data ChecklistFile
}

// LoadFromFile loads a CKLB file into the Checklist struct
func (c *Checklist) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	
	return json.Unmarshal(data, &c.Data)
}

// SaveToFile saves the Checklist struct to a CKLB file
func (c *Checklist) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(c.Data, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(filename, data, 0644)
}

// GetSTIGs returns all STIGs in the checklist
func (c *Checklist) GetSTIGs() []STIG {
	return c.Data.STIGs
}

// GetRulesWithStatus returns all rules with the specified status
func (c *Checklist) GetRulesWithStatus(status string) []STIGRule {
	var rules []STIGRule
	
	for _, stig := range c.Data.STIGs {
		for _, rule := range stig.Rules {
			if rule.Status == status {
				rules = append(rules, rule)
			}
		}
	}
	
	return rules
}

// UpdateRuleStatus updates the status of a rule
func (c *Checklist) UpdateRuleStatus(ruleID string, status string) error {
	for i, stig := range c.Data.STIGs {
		for j, rule := range stig.Rules {
			if rule.RuleID == ruleID {
				c.Data.STIGs[i].Rules[j].Status = status
				return nil
			}
		}
	}
	
	return fmt.Errorf("rule %s not found", ruleID)
}

// AddComment adds a comment to a rule
func (c *Checklist) AddComment(ruleID string, comment string) error {
	for i, stig := range c.Data.STIGs {
		for j, rule := range stig.Rules {
			if rule.RuleID == ruleID {
				c.Data.STIGs[i].Rules[j].Comments = comment
				return nil
			}
		}
	}
	
	return fmt.Errorf("rule %s not found", ruleID)
}

// GetTargetInfo returns the target info
func (c *Checklist) GetTargetInfo() TargetData {
	return c.Data.TargetData
}

// UpdateTargetInfo updates the target info
func (c *Checklist) UpdateTargetInfo(targetData TargetData) error {
	c.Data.TargetData = targetData
	return nil
}