package cmd

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)

// ChecklistFile represents the root structure of a CKLB file
type ChecklistFile struct {
	Title       string      `json:"title"`
	CklbVersion string      `json:"cklb_version,omitempty"`
	ID          string      `json:"id"`
	Active      bool        `json:"active,omitempty"`
	Mode        int         `json:"mode,omitempty"`
	HasPath     bool        `json:"has_path,omitempty"`
	TargetData  TargetData  `json:"target_data,omitempty"`
	STIGs       []STIG      `json:"stigs,omitempty"`
}

// TargetData contains information about the scanned system
type TargetData struct {
	TargetType     string `json:"target_type,omitempty"`
	HostName       string `json:"host_name,omitempty"`
	IPAddress      string `json:"ip_address,omitempty"`
	MACAddress     string `json:"mac_address,omitempty"`
	FQDN           string `json:"fqdn,omitempty"`
	Comments       string `json:"comments,omitempty"`
	Role           string `json:"role,omitempty"`
	IsWebDatabase  bool   `json:"is_web_database,omitempty"`
	TechnologyArea string `json:"technology_area,omitempty"`
	WebDBSite      string `json:"web_db_site,omitempty"`
	WebDBInstance  string `json:"web_db_instance,omitempty"`
}

// STIG represents a single STIG within the checklist
type STIG struct {
	STIGName           string     `json:"stig_name"`
	DisplayName        string     `json:"display_name"`
	STIGID             string     `json:"stig_id"`
	ReleaseInfo        string     `json:"release_info"`
	UUID               string     `json:"uuid"`
	ReferenceIdentifier string     `json:"reference_identifier,omitempty"`
	Size               int        `json:"size"`
	Rules              []STIGRule `json:"rules,omitempty"`
}

// STIGRule represents an individual rule within a STIG
type STIGRule struct {
	UUID                    string        `json:"uuid"`
	STIGUUID                string        `json:"stig_uuid"`
	GroupID                 string        `json:"group_id"`
	GroupIDSrc              string        `json:"group_id_src"`
	RuleID                  string        `json:"rule_id"`
	RuleIDSrc               string        `json:"rule_id_src"`
	TargetKey               string        `json:"target_key,omitempty"`
	STIGRef                 string        `json:"stig_ref,omitempty"`
	Weight                  string        `json:"weight,omitempty"`
	Classification          string        `json:"classification,omitempty"`
	Severity                string        `json:"severity,omitempty"`
	RuleVersion             string        `json:"rule_version,omitempty"`
	RuleTitle               string        `json:"rule_title,omitempty"`
	FixText                 string        `json:"fix_text,omitempty"`
	ReferenceIdentifier     string        `json:"reference_identifier,omitempty"`
	GroupTitle              string        `json:"group_title,omitempty"`
	FalsePositives          string        `json:"false_positives,omitempty"`
	FalseNegatives          string        `json:"false_negatives,omitempty"`
	Discussion              string        `json:"discussion,omitempty"`
	CheckContent            string        `json:"check_content,omitempty"`
	Documentable            string        `json:"documentable,omitempty"`
	Mitigations             string        `json:"mitigations,omitempty"`
	PotentialImpacts        string        `json:"potential_impacts,omitempty"`
	ThirdPartyTools         string        `json:"third_party_tools,omitempty"`
	MitigationControl       string        `json:"mitigation_control,omitempty"`
	Responsibility          string        `json:"responsibility,omitempty"`
	SecurityOverrideGuidance string        `json:"security_override_guidance,omitempty"`
	IAControls              string        `json:"ia_controls,omitempty"`
	CheckContentRef         *CheckContentRef `json:"check_content_ref,omitempty"`
	LegacyIDs               []string      `json:"legacy_ids,omitempty"`
	CCIs                    []string      `json:"ccis,omitempty"`
	GroupTree               []GroupTree   `json:"group_tree,omitempty"`
	CreatedAt               string        `json:"createdAt,omitempty"`
	UpdatedAt               string        `json:"updatedAt,omitempty"`
	Status                  string        `json:"status,omitempty"`
	Overrides               json.RawMessage `json:"overrides,omitempty"`
	Comments                string        `json:"comments,omitempty"`
	FindingDetails          string        `json:"finding_details,omitempty"`
	STIGUuidDeprecated      string        `json:"STIGUuid,omitempty"`
}

// CheckContentRef represents the check content reference
type CheckContentRef struct {
	Name string `json:"name,omitempty"`
	Href string `json:"href,omitempty"`
}

// GroupTree represents one level in the group hierarchy
type GroupTree struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// ChecklistInterface defines operations that can be performed on a checklist
type ChecklistInterface interface {
	LoadFromFile(filename string) error
	SaveToFile(filename string) error
	Validate() (bool, []string)
	GetSTIGs() []STIG
	GetRulesWithStatus(status string) []STIGRule
	UpdateRuleStatus(ruleID string, status string) error
	AddComment(ruleID string, comment string) error
	GetTargetInfo() TargetData
	UpdateTargetInfo(targetData TargetData) error
}

// Checklist implements the ChecklistInterface
type Checklist struct {
	Data ChecklistFile
}

// LoadFromFile loads a CKLB file into the Checklist struct
func (c *Checklist) LoadFromFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
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
	
	return ioutil.WriteFile(filename, data, 0644)
}

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
	
	return len(errors) == 0, errors
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