package meonenotenotebooksectiongroupsectiongroup

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnenoteNotebookSectionGroupSectionGroupId{}

func TestNewMeOnenoteNotebookSectionGroupSectionGroupID(t *testing.T) {
	id := NewMeOnenoteNotebookSectionGroupSectionGroupID("notebookIdValue", "sectionGroupIdValue", "sectionGroupId1Value")

	if id.NotebookId != "notebookIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'NotebookId'", id.NotebookId, "notebookIdValue")
	}

	if id.SectionGroupId != "sectionGroupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SectionGroupId'", id.SectionGroupId, "sectionGroupIdValue")
	}

	if id.SectionGroupId1 != "sectionGroupId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'SectionGroupId1'", id.SectionGroupId1, "sectionGroupId1Value")
	}
}

func TestFormatMeOnenoteNotebookSectionGroupSectionGroupID(t *testing.T) {
	actual := NewMeOnenoteNotebookSectionGroupSectionGroupID("notebookIdValue", "sectionGroupIdValue", "sectionGroupId1Value").ID()
	expected := "/me/onenote/notebooks/notebookIdValue/sectionGroups/sectionGroupIdValue/sectionGroups/sectionGroupId1Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnenoteNotebookSectionGroupSectionGroupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnenoteNotebookSectionGroupSectionGroupId
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
			Input: "/me/onenote",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/notebooks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/notebooks/notebookIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/notebooks/notebookIdValue/sectionGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/notebooks/notebookIdValue/sectionGroups/sectionGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/notebooks/notebookIdValue/sectionGroups/sectionGroupIdValue/sectionGroups",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onenote/notebooks/notebookIdValue/sectionGroups/sectionGroupIdValue/sectionGroups/sectionGroupId1Value",
			Expected: &MeOnenoteNotebookSectionGroupSectionGroupId{
				NotebookId:      "notebookIdValue",
				SectionGroupId:  "sectionGroupIdValue",
				SectionGroupId1: "sectionGroupId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onenote/notebooks/notebookIdValue/sectionGroups/sectionGroupIdValue/sectionGroups/sectionGroupId1Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnenoteNotebookSectionGroupSectionGroupID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.NotebookId != v.Expected.NotebookId {
			t.Fatalf("Expected %q but got %q for NotebookId", v.Expected.NotebookId, actual.NotebookId)
		}

		if actual.SectionGroupId != v.Expected.SectionGroupId {
			t.Fatalf("Expected %q but got %q for SectionGroupId", v.Expected.SectionGroupId, actual.SectionGroupId)
		}

		if actual.SectionGroupId1 != v.Expected.SectionGroupId1 {
			t.Fatalf("Expected %q but got %q for SectionGroupId1", v.Expected.SectionGroupId1, actual.SectionGroupId1)
		}

	}
}

func TestParseMeOnenoteNotebookSectionGroupSectionGroupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnenoteNotebookSectionGroupSectionGroupId
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
			Input: "/me/onenote",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/notebooks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/nOtEbOoKs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/notebooks/notebookIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/nOtEbOoKs/nOtEbOoKiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/notebooks/notebookIdValue/sectionGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/nOtEbOoKs/nOtEbOoKiDvAlUe/sEcTiOnGrOuPs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/notebooks/notebookIdValue/sectionGroups/sectionGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/nOtEbOoKs/nOtEbOoKiDvAlUe/sEcTiOnGrOuPs/sEcTiOnGrOuPiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/notebooks/notebookIdValue/sectionGroups/sectionGroupIdValue/sectionGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/nOtEbOoKs/nOtEbOoKiDvAlUe/sEcTiOnGrOuPs/sEcTiOnGrOuPiDvAlUe/sEcTiOnGrOuPs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onenote/notebooks/notebookIdValue/sectionGroups/sectionGroupIdValue/sectionGroups/sectionGroupId1Value",
			Expected: &MeOnenoteNotebookSectionGroupSectionGroupId{
				NotebookId:      "notebookIdValue",
				SectionGroupId:  "sectionGroupIdValue",
				SectionGroupId1: "sectionGroupId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onenote/notebooks/notebookIdValue/sectionGroups/sectionGroupIdValue/sectionGroups/sectionGroupId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/nOtEbOoKs/nOtEbOoKiDvAlUe/sEcTiOnGrOuPs/sEcTiOnGrOuPiDvAlUe/sEcTiOnGrOuPs/sEcTiOnGrOuPiD1VaLuE",
			Expected: &MeOnenoteNotebookSectionGroupSectionGroupId{
				NotebookId:      "nOtEbOoKiDvAlUe",
				SectionGroupId:  "sEcTiOnGrOuPiDvAlUe",
				SectionGroupId1: "sEcTiOnGrOuPiD1VaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/nOtEbOoKs/nOtEbOoKiDvAlUe/sEcTiOnGrOuPs/sEcTiOnGrOuPiDvAlUe/sEcTiOnGrOuPs/sEcTiOnGrOuPiD1VaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnenoteNotebookSectionGroupSectionGroupIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.NotebookId != v.Expected.NotebookId {
			t.Fatalf("Expected %q but got %q for NotebookId", v.Expected.NotebookId, actual.NotebookId)
		}

		if actual.SectionGroupId != v.Expected.SectionGroupId {
			t.Fatalf("Expected %q but got %q for SectionGroupId", v.Expected.SectionGroupId, actual.SectionGroupId)
		}

		if actual.SectionGroupId1 != v.Expected.SectionGroupId1 {
			t.Fatalf("Expected %q but got %q for SectionGroupId1", v.Expected.SectionGroupId1, actual.SectionGroupId1)
		}

	}
}

func TestSegmentsForMeOnenoteNotebookSectionGroupSectionGroupId(t *testing.T) {
	segments := MeOnenoteNotebookSectionGroupSectionGroupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnenoteNotebookSectionGroupSectionGroupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
