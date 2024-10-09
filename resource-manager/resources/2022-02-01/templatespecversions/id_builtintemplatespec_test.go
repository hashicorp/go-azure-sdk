package templatespecversions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &BuiltInTemplateSpecId{}

func TestNewBuiltInTemplateSpecID(t *testing.T) {
	id := NewBuiltInTemplateSpecID("builtInTemplateSpecName")

	if id.BuiltInTemplateSpecName != "builtInTemplateSpecName" {
		t.Fatalf("Expected %q but got %q for Segment 'BuiltInTemplateSpecName'", id.BuiltInTemplateSpecName, "builtInTemplateSpecName")
	}
}

func TestFormatBuiltInTemplateSpecID(t *testing.T) {
	actual := NewBuiltInTemplateSpecID("builtInTemplateSpecName").ID()
	expected := "/providers/Microsoft.Resources/builtInTemplateSpecs/builtInTemplateSpecName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseBuiltInTemplateSpecID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BuiltInTemplateSpecId
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
			// Valid URI
			Input: "/providers/Microsoft.Resources/builtInTemplateSpecs/builtInTemplateSpecName",
			Expected: &BuiltInTemplateSpecId{
				BuiltInTemplateSpecName: "builtInTemplateSpecName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Resources/builtInTemplateSpecs/builtInTemplateSpecName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBuiltInTemplateSpecID(v.Input)
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

	}
}

func TestParseBuiltInTemplateSpecIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BuiltInTemplateSpecId
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
			// Valid URI
			Input: "/providers/Microsoft.Resources/builtInTemplateSpecs/builtInTemplateSpecName",
			Expected: &BuiltInTemplateSpecId{
				BuiltInTemplateSpecName: "builtInTemplateSpecName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Resources/builtInTemplateSpecs/builtInTemplateSpecName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/bUiLtInTeMpLaTeSpEcS/bUiLtInTeMpLaTeSpEcNaMe",
			Expected: &BuiltInTemplateSpecId{
				BuiltInTemplateSpecName: "bUiLtInTeMpLaTeSpEcNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/bUiLtInTeMpLaTeSpEcS/bUiLtInTeMpLaTeSpEcNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBuiltInTemplateSpecIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForBuiltInTemplateSpecId(t *testing.T) {
	segments := BuiltInTemplateSpecId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("BuiltInTemplateSpecId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
