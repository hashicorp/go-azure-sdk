package views

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ViewId{}

func TestNewViewID(t *testing.T) {
	id := NewViewID("viewName")

	if id.ViewName != "viewName" {
		t.Fatalf("Expected %q but got %q for Segment 'ViewName'", id.ViewName, "viewName")
	}
}

func TestFormatViewID(t *testing.T) {
	actual := NewViewID("viewName").ID()
	expected := "/providers/Microsoft.CostManagement/views/viewName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseViewID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ViewId
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
			Input: "/providers/Microsoft.CostManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.CostManagement/views",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.CostManagement/views/viewName",
			Expected: &ViewId{
				ViewName: "viewName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.CostManagement/views/viewName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseViewID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ViewName != v.Expected.ViewName {
			t.Fatalf("Expected %q but got %q for ViewName", v.Expected.ViewName, actual.ViewName)
		}

	}
}

func TestParseViewIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ViewId
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
			Input: "/providers/Microsoft.CostManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cOsTmAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.CostManagement/views",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cOsTmAnAgEmEnT/vIeWs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.CostManagement/views/viewName",
			Expected: &ViewId{
				ViewName: "viewName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.CostManagement/views/viewName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cOsTmAnAgEmEnT/vIeWs/vIeWnAmE",
			Expected: &ViewId{
				ViewName: "vIeWnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cOsTmAnAgEmEnT/vIeWs/vIeWnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseViewIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ViewName != v.Expected.ViewName {
			t.Fatalf("Expected %q but got %q for ViewName", v.Expected.ViewName, actual.ViewName)
		}

	}
}

func TestSegmentsForViewId(t *testing.T) {
	segments := ViewId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ViewId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
