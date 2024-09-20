package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenoteNotebookIdSectionGroupIdSectionId{}

func TestNewMeOnenoteNotebookIdSectionGroupIdSectionID(t *testing.T) {
	id := NewMeOnenoteNotebookIdSectionGroupIdSectionID("notebookId", "sectionGroupId", "onenoteSectionId")

	if id.NotebookId != "notebookId" {
		t.Fatalf("Expected %q but got %q for Segment 'NotebookId'", id.NotebookId, "notebookId")
	}

	if id.SectionGroupId != "sectionGroupId" {
		t.Fatalf("Expected %q but got %q for Segment 'SectionGroupId'", id.SectionGroupId, "sectionGroupId")
	}

	if id.OnenoteSectionId != "onenoteSectionId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenoteSectionId'", id.OnenoteSectionId, "onenoteSectionId")
	}
}

func TestFormatMeOnenoteNotebookIdSectionGroupIdSectionID(t *testing.T) {
	actual := NewMeOnenoteNotebookIdSectionGroupIdSectionID("notebookId", "sectionGroupId", "onenoteSectionId").ID()
	expected := "/me/onenote/notebooks/notebookId/sectionGroups/sectionGroupId/sections/onenoteSectionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnenoteNotebookIdSectionGroupIdSectionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnenoteNotebookIdSectionGroupIdSectionId
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
			Input: "/me/onenote/notebooks/notebookId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/notebooks/notebookId/sectionGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/notebooks/notebookId/sectionGroups/sectionGroupId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/notebooks/notebookId/sectionGroups/sectionGroupId/sections",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onenote/notebooks/notebookId/sectionGroups/sectionGroupId/sections/onenoteSectionId",
			Expected: &MeOnenoteNotebookIdSectionGroupIdSectionId{
				NotebookId:       "notebookId",
				SectionGroupId:   "sectionGroupId",
				OnenoteSectionId: "onenoteSectionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onenote/notebooks/notebookId/sectionGroups/sectionGroupId/sections/onenoteSectionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnenoteNotebookIdSectionGroupIdSectionID(v.Input)
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

		if actual.OnenoteSectionId != v.Expected.OnenoteSectionId {
			t.Fatalf("Expected %q but got %q for OnenoteSectionId", v.Expected.OnenoteSectionId, actual.OnenoteSectionId)
		}

	}
}

func TestParseMeOnenoteNotebookIdSectionGroupIdSectionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnenoteNotebookIdSectionGroupIdSectionId
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
			Input: "/me/onenote/notebooks/notebookId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/nOtEbOoKs/nOtEbOoKiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/notebooks/notebookId/sectionGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/nOtEbOoKs/nOtEbOoKiD/sEcTiOnGrOuPs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/notebooks/notebookId/sectionGroups/sectionGroupId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/nOtEbOoKs/nOtEbOoKiD/sEcTiOnGrOuPs/sEcTiOnGrOuPiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/notebooks/notebookId/sectionGroups/sectionGroupId/sections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/nOtEbOoKs/nOtEbOoKiD/sEcTiOnGrOuPs/sEcTiOnGrOuPiD/sEcTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onenote/notebooks/notebookId/sectionGroups/sectionGroupId/sections/onenoteSectionId",
			Expected: &MeOnenoteNotebookIdSectionGroupIdSectionId{
				NotebookId:       "notebookId",
				SectionGroupId:   "sectionGroupId",
				OnenoteSectionId: "onenoteSectionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onenote/notebooks/notebookId/sectionGroups/sectionGroupId/sections/onenoteSectionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/nOtEbOoKs/nOtEbOoKiD/sEcTiOnGrOuPs/sEcTiOnGrOuPiD/sEcTiOnS/oNeNoTeSeCtIoNiD",
			Expected: &MeOnenoteNotebookIdSectionGroupIdSectionId{
				NotebookId:       "nOtEbOoKiD",
				SectionGroupId:   "sEcTiOnGrOuPiD",
				OnenoteSectionId: "oNeNoTeSeCtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/nOtEbOoKs/nOtEbOoKiD/sEcTiOnGrOuPs/sEcTiOnGrOuPiD/sEcTiOnS/oNeNoTeSeCtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnenoteNotebookIdSectionGroupIdSectionIDInsensitively(v.Input)
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

		if actual.OnenoteSectionId != v.Expected.OnenoteSectionId {
			t.Fatalf("Expected %q but got %q for OnenoteSectionId", v.Expected.OnenoteSectionId, actual.OnenoteSectionId)
		}

	}
}

func TestSegmentsForMeOnenoteNotebookIdSectionGroupIdSectionId(t *testing.T) {
	segments := MeOnenoteNotebookIdSectionGroupIdSectionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnenoteNotebookIdSectionGroupIdSectionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
