package resourceproviders

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SourceControlId{}

func TestNewSourceControlID(t *testing.T) {
	id := NewSourceControlID("sourceControlName")

	if id.SourceControlName != "sourceControlName" {
		t.Fatalf("Expected %q but got %q for Segment 'SourceControlName'", id.SourceControlName, "sourceControlName")
	}
}

func TestFormatSourceControlID(t *testing.T) {
	actual := NewSourceControlID("sourceControlName").ID()
	expected := "/providers/Microsoft.Web/sourceControls/sourceControlName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSourceControlID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SourceControlId
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
			Input: "/providers/Microsoft.Web",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Web/sourceControls",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Web/sourceControls/sourceControlName",
			Expected: &SourceControlId{
				SourceControlName: "sourceControlName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Web/sourceControls/sourceControlName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSourceControlID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SourceControlName != v.Expected.SourceControlName {
			t.Fatalf("Expected %q but got %q for SourceControlName", v.Expected.SourceControlName, actual.SourceControlName)
		}

	}
}

func TestParseSourceControlIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SourceControlId
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
			Input: "/providers/Microsoft.Web",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.wEb",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Web/sourceControls",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.wEb/sOuRcEcOnTrOlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Web/sourceControls/sourceControlName",
			Expected: &SourceControlId{
				SourceControlName: "sourceControlName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Web/sourceControls/sourceControlName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.wEb/sOuRcEcOnTrOlS/sOuRcEcOnTrOlNaMe",
			Expected: &SourceControlId{
				SourceControlName: "sOuRcEcOnTrOlNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.wEb/sOuRcEcOnTrOlS/sOuRcEcOnTrOlNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSourceControlIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SourceControlName != v.Expected.SourceControlName {
			t.Fatalf("Expected %q but got %q for SourceControlName", v.Expected.SourceControlName, actual.SourceControlName)
		}

	}
}

func TestSegmentsForSourceControlId(t *testing.T) {
	segments := SourceControlId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SourceControlId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
