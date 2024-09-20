package templatespecversions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &VersionId{}

func TestNewVersionID(t *testing.T) {
	id := NewVersionID("templateSpecName", "templateSpecVersion")

	if id.BuiltInTemplateSpecName != "templateSpecName" {
		t.Fatalf("Expected %q but got %q for Segment 'BuiltInTemplateSpecName'", id.BuiltInTemplateSpecName, "templateSpecName")
	}

	if id.VersionName != "templateSpecVersion" {
		t.Fatalf("Expected %q but got %q for Segment 'VersionName'", id.VersionName, "templateSpecVersion")
	}
}

func TestFormatVersionID(t *testing.T) {
	actual := NewVersionID("templateSpecName", "templateSpecVersion").ID()
	expected := "/providers/Microsoft.Resources/builtInTemplateSpecs/templateSpecName/versions/templateSpecVersion"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VersionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Resources",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Resources/builtInTemplateSpecs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Resources/builtInTemplateSpecs/templateSpecName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Resources/builtInTemplateSpecs/templateSpecName/versions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Resources/builtInTemplateSpecs/templateSpecName/versions/templateSpecVersion",
			Expected: &VersionId{
				BuiltInTemplateSpecName: "templateSpecName",
				VersionName:             "templateSpecVersion",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Resources/builtInTemplateSpecs/templateSpecName/versions/templateSpecVersion/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVersionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BuiltInTemplateSpecName != v.Expected.BuiltInTemplateSpecName {
			t.Fatalf("Expected %q but got %q for BuiltInTemplateSpecName", v.Expected.BuiltInTemplateSpecName, actual.BuiltInTemplateSpecName)
		}

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
		}

	}
}

func TestParseVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VersionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Resources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.rEsOuRcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Resources/builtInTemplateSpecs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/bUiLtInTeMpLaTeSpEcS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Resources/builtInTemplateSpecs/templateSpecName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/bUiLtInTeMpLaTeSpEcS/tEmPlAtEsPeCnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Resources/builtInTemplateSpecs/templateSpecName/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/bUiLtInTeMpLaTeSpEcS/tEmPlAtEsPeCnAmE/vErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Resources/builtInTemplateSpecs/templateSpecName/versions/templateSpecVersion",
			Expected: &VersionId{
				BuiltInTemplateSpecName: "templateSpecName",
				VersionName:             "templateSpecVersion",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Resources/builtInTemplateSpecs/templateSpecName/versions/templateSpecVersion/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/bUiLtInTeMpLaTeSpEcS/tEmPlAtEsPeCnAmE/vErSiOnS/tEmPlAtEsPeCvErSiOn",
			Expected: &VersionId{
				BuiltInTemplateSpecName: "tEmPlAtEsPeCnAmE",
				VersionName:             "tEmPlAtEsPeCvErSiOn",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/bUiLtInTeMpLaTeSpEcS/tEmPlAtEsPeCnAmE/vErSiOnS/tEmPlAtEsPeCvErSiOn/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVersionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BuiltInTemplateSpecName != v.Expected.BuiltInTemplateSpecName {
			t.Fatalf("Expected %q but got %q for BuiltInTemplateSpecName", v.Expected.BuiltInTemplateSpecName, actual.BuiltInTemplateSpecName)
		}

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
		}

	}
}

func TestSegmentsForVersionId(t *testing.T) {
	segments := VersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("VersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
