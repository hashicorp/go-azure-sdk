package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenoteNotebookIdSectionGroupId{}

func TestNewMeOnenoteNotebookIdSectionGroupID(t *testing.T) {
	id := NewMeOnenoteNotebookIdSectionGroupID("notebookId", "sectionGroupId")

	if id.NotebookId != "notebookId" {
		t.Fatalf("Expected %q but got %q for Segment 'NotebookId'", id.NotebookId, "notebookId")
	}

	if id.SectionGroupId != "sectionGroupId" {
		t.Fatalf("Expected %q but got %q for Segment 'SectionGroupId'", id.SectionGroupId, "sectionGroupId")
	}
}

func TestFormatMeOnenoteNotebookIdSectionGroupID(t *testing.T) {
	actual := NewMeOnenoteNotebookIdSectionGroupID("notebookId", "sectionGroupId").ID()
	expected := "/me/onenote/notebooks/notebookId/sectionGroups/sectionGroupId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnenoteNotebookIdSectionGroupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnenoteNotebookIdSectionGroupId
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
			// Valid URI
			Input: "/me/onenote/notebooks/notebookId/sectionGroups/sectionGroupId",
			Expected: &MeOnenoteNotebookIdSectionGroupId{
				NotebookId:     "notebookId",
				SectionGroupId: "sectionGroupId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onenote/notebooks/notebookId/sectionGroups/sectionGroupId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnenoteNotebookIdSectionGroupID(v.Input)
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

	}
}

func TestParseMeOnenoteNotebookIdSectionGroupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnenoteNotebookIdSectionGroupId
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
			// Valid URI
			Input: "/me/onenote/notebooks/notebookId/sectionGroups/sectionGroupId",
			Expected: &MeOnenoteNotebookIdSectionGroupId{
				NotebookId:     "notebookId",
				SectionGroupId: "sectionGroupId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onenote/notebooks/notebookId/sectionGroups/sectionGroupId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/nOtEbOoKs/nOtEbOoKiD/sEcTiOnGrOuPs/sEcTiOnGrOuPiD",
			Expected: &MeOnenoteNotebookIdSectionGroupId{
				NotebookId:     "nOtEbOoKiD",
				SectionGroupId: "sEcTiOnGrOuPiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/nOtEbOoKs/nOtEbOoKiD/sEcTiOnGrOuPs/sEcTiOnGrOuPiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnenoteNotebookIdSectionGroupIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForMeOnenoteNotebookIdSectionGroupId(t *testing.T) {
	segments := MeOnenoteNotebookIdSectionGroupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnenoteNotebookIdSectionGroupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
