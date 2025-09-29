package cklb

import (
	"encoding/json"
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