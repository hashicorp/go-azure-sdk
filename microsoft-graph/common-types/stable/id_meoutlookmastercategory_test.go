package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOutlookMasterCategoryId{}

func TestNewMeOutlookMasterCategoryID(t *testing.T) {
	id := NewMeOutlookMasterCategoryID("outlookCategoryId")

	if id.OutlookCategoryId != "outlookCategoryId" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookCategoryId'", id.OutlookCategoryId, "outlookCategoryId")
	}
}

func TestFormatMeOutlookMasterCategoryID(t *testing.T) {
	actual := NewMeOutlookMasterCategoryID("outlookCategoryId").ID()
	expected := "/me/outlook/masterCategories/outlookCategoryId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOutlookMasterCategoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOutlookMasterCategoryId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/masterCategories",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/outlook/masterCategories/outlookCategoryId",
			Expected: &MeOutlookMasterCategoryId{
				OutlookCategoryId: "outlookCategoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/outlook/masterCategories/outlookCategoryId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOutlookMasterCategoryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OutlookCategoryId != v.Expected.OutlookCategoryId {
			t.Fatalf("Expected %q but got %q for OutlookCategoryId", v.Expected.OutlookCategoryId, actual.OutlookCategoryId)
		}

	}
}

func TestParseMeOutlookMasterCategoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOutlookMasterCategoryId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/masterCategories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/mAsTeRcAtEgOrIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/outlook/masterCategories/outlookCategoryId",
			Expected: &MeOutlookMasterCategoryId{
				OutlookCategoryId: "outlookCategoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/outlook/masterCategories/outlookCategoryId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/mAsTeRcAtEgOrIeS/oUtLoOkCaTeGoRyId",
			Expected: &MeOutlookMasterCategoryId{
				OutlookCategoryId: "oUtLoOkCaTeGoRyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/mAsTeRcAtEgOrIeS/oUtLoOkCaTeGoRyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOutlookMasterCategoryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OutlookCategoryId != v.Expected.OutlookCategoryId {
			t.Fatalf("Expected %q but got %q for OutlookCategoryId", v.Expected.OutlookCategoryId, actual.OutlookCategoryId)
		}

	}
}

func TestSegmentsForMeOutlookMasterCategoryId(t *testing.T) {
	segments := MeOutlookMasterCategoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOutlookMasterCategoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
