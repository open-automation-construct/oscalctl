package cciparsing

import (
	"strings"
	"testing"
)

func TestNormalizeControlID(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"SA-4 (7) (a)", "sa-4.7"},
		{"AC-1 a", "ac-1"},
		{"AC-2 (4)", "ac-2.4"},
		{"Invalid", ""},
	}
	
	for _, tc := range testCases {
		result := normalizeControlID(tc.input)
		if result != tc.expected {
			t.Errorf("normalizeControlID(%s) = %s, expected %s", tc.input, result, tc.expected)
		}
	}
}

func TestIsValidOSCALToken(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"sa-4.7", true},
		{"ac-1", true},
		{"1invalid", false},
		{"valid_token-123.45", true},
		{"-invalid", false},
		{"", false},
	}
	
	for _, tc := range testCases {
		result := IsValidOSCALToken(tc.input)
		if result != tc.expected {
			t.Errorf("IsValidOSCALToken(%s) = %v, expected %v", tc.input, result, tc.expected)
		}
	}
}

func TestGetPreferredControlID(t *testing.T) {
	references := []Reference{
		{Creator: "NIST", Title: "NIST SP 800-53", Version: "3", Index: "SA-4 a"},
		{Creator: "NIST", Title: "NIST SP 800-53A", Version: "1", Index: "SA-4.1 (i)"},
		{Creator: "NIST", Title: "NIST SP 800-53 Revision 4", Version: "4", Index: "SA-4 (7) (a)"},
	}
	
	expected := "sa-4.7"
	result := getPreferredControlID(references)
	
	if result != expected {
		t.Errorf("getPreferredControlID() = %s, expected %s", result, expected)
	}
	
	// Test with Rev 5 as highest priority
	references = append(references, Reference{
		Creator: "NIST", Title: "NIST SP 800-53 Revision 5", Version: "5", Index: "SA-4 (9) (b)",
	})
	
	expected = "sa-4.9"
	result = getPreferredControlID(references)
	
	if result != expected {
		t.Errorf("getPreferredControlID() with Rev 5 = %s, expected %s", result, expected)
	}
}

func TestParseCCIDocumentReader(t *testing.T) {
	// Simple test XML
	xmlData := `<?xml version="1.0" encoding="utf-8"?>
<cci_list xmlns="http://iase.disa.mil/cci">
  <metadata>
    <version>2025-09-19</version>
    <publishdate>2025-09-19</publishdate>
  </metadata>
  <cci_items>
    <cci_item id="CCI-000001">
      <status>draft</status>
      <publishdate>2009-05-13</publishdate>
      <contributor>DISA FSO</contributor>
      <definition>Test definition</definition>
      <type>policy</type>
      <references>
        <reference creator="NIST" title="NIST SP 800-53" version="3" location="http://csrc.nist.gov/publications/PubsSPs.html" index="AC-1 a" />
        <reference creator="NIST" title="NIST SP 800-53 Revision 4" version="4" location="http://csrc.nist.gov/publications/PubsSPs.html" index="AC-1 a 1" />
      </references>
    </cci_item>
    <cci_item id="CCI-000002">
      <status>draft</status>
      <publishdate>2009-05-13</publishdate>
      <contributor>DISA FSO</contributor>
      <definition>Test definition 2</definition>
      <type>policy</type>
      <references>
        <reference creator="NIST" title="NIST SP 800-53 Revision 5" version="5" location="http://csrc.nist.gov/publications/PubsSPs.html" index="SA-4 (7) (a)" />
      </references>
    </cci_item>
  </cci_items>
</cci_list>`
	
	expected := map[string]string{
		"CCI-000001": "ac-1",
		"CCI-000002": "sa-4.7",
	}
	
	reader := strings.NewReader(xmlData)
	result, err := ParseCCIDocumentReader(reader)
	
	if err != nil {
		t.Fatalf("ParseCCIDocumentReader() returned error: %v", err)
	}
	
	if len(result) != len(expected) {
		t.Errorf("ParseCCIDocumentReader() returned %d items, expected %d", len(result), len(expected))
	}
	
	for k, v := range expected {
		if result[k] != v {
			t.Errorf("ParseCCIDocumentReader() result[%s] = %s, expected %s", k, result[k], v)
		}
	}
}