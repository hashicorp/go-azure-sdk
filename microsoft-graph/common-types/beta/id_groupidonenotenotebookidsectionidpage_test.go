package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdOnenoteNotebookIdSectionIdPageId{}

func TestNewGroupIdOnenoteNotebookIdSectionIdPageID(t *testing.T) {
	id := NewGroupIdOnenoteNotebookIdSectionIdPageID("groupId", "notebookId", "onenoteSectionId", "onenotePageId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.NotebookId != "notebookId" {
		t.Fatalf("Expected %q but got %q for Segment 'NotebookId'", id.NotebookId, "notebookId")
	}

	if id.OnenoteSectionId != "onenoteSectionId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenoteSectionId'", id.OnenoteSectionId, "onenoteSectionId")
	}

	if id.OnenotePageId != "onenotePageId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenotePageId'", id.OnenotePageId, "onenotePageId")
	}
}

func TestFormatGroupIdOnenoteNotebookIdSectionIdPageID(t *testing.T) {
	actual := NewGroupIdOnenoteNotebookIdSectionIdPageID("groupId", "notebookId", "onenoteSectionId", "onenotePageId").ID()
	expected := "/groups/groupId/onenote/notebooks/notebookId/sections/onenoteSectionId/pages/onenotePageId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdOnenoteNotebookIdSectionIdPageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdOnenoteNotebookIdSectionIdPageId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/onenote",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/onenote/notebooks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/onenote/notebooks/notebookId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/onenote/notebooks/notebookId/sections",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/onenote/notebooks/notebookId/sections/onenoteSectionId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/onenote/notebooks/notebookId/sections/onenoteSectionId/pages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/onenote/notebooks/notebookId/sections/onenoteSectionId/pages/onenotePageId",
			Expected: &GroupIdOnenoteNotebookIdSectionIdPageId{
				GroupId:          "groupId",
				NotebookId:       "notebookId",
				OnenoteSectionId: "onenoteSectionId",
				OnenotePageId:    "onenotePageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/onenote/notebooks/notebookId/sections/onenoteSectionId/pages/onenotePageId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdOnenoteNotebookIdSectionIdPageID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.NotebookId != v.Expected.NotebookId {
			t.Fatalf("Expected %q but got %q for NotebookId", v.Expected.NotebookId, actual.NotebookId)
		}

		if actual.OnenoteSectionId != v.Expected.OnenoteSectionId {
			t.Fatalf("Expected %q but got %q for OnenoteSectionId", v.Expected.OnenoteSectionId, actual.OnenoteSectionId)
		}

		if actual.OnenotePageId != v.Expected.OnenotePageId {
			t.Fatalf("Expected %q but got %q for OnenotePageId", v.Expected.OnenotePageId, actual.OnenotePageId)
		}

	}
}

func TestParseGroupIdOnenoteNotebookIdSectionIdPageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdOnenoteNotebookIdSectionIdPageId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/onenote",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/oNeNoTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/onenote/notebooks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/oNeNoTe/nOtEbOoKs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/onenote/notebooks/notebookId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/oNeNoTe/nOtEbOoKs/nOtEbOoKiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/onenote/notebooks/notebookId/sections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/oNeNoTe/nOtEbOoKs/nOtEbOoKiD/sEcTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/onenote/notebooks/notebookId/sections/onenoteSectionId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/oNeNoTe/nOtEbOoKs/nOtEbOoKiD/sEcTiOnS/oNeNoTeSeCtIoNiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/onenote/notebooks/notebookId/sections/onenoteSectionId/pages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/oNeNoTe/nOtEbOoKs/nOtEbOoKiD/sEcTiOnS/oNeNoTeSeCtIoNiD/pAgEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/onenote/notebooks/notebookId/sections/onenoteSectionId/pages/onenotePageId",
			Expected: &GroupIdOnenoteNotebookIdSectionIdPageId{
				GroupId:          "groupId",
				NotebookId:       "notebookId",
				OnenoteSectionId: "onenoteSectionId",
				OnenotePageId:    "onenotePageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/onenote/notebooks/notebookId/sections/onenoteSectionId/pages/onenotePageId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/oNeNoTe/nOtEbOoKs/nOtEbOoKiD/sEcTiOnS/oNeNoTeSeCtIoNiD/pAgEs/oNeNoTePaGeId",
			Expected: &GroupIdOnenoteNotebookIdSectionIdPageId{
				GroupId:          "gRoUpId",
				NotebookId:       "nOtEbOoKiD",
				OnenoteSectionId: "oNeNoTeSeCtIoNiD",
				OnenotePageId:    "oNeNoTePaGeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/oNeNoTe/nOtEbOoKs/nOtEbOoKiD/sEcTiOnS/oNeNoTeSeCtIoNiD/pAgEs/oNeNoTePaGeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdOnenoteNotebookIdSectionIdPageIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.NotebookId != v.Expected.NotebookId {
			t.Fatalf("Expected %q but got %q for NotebookId", v.Expected.NotebookId, actual.NotebookId)
		}

		if actual.OnenoteSectionId != v.Expected.OnenoteSectionId {
			t.Fatalf("Expected %q but got %q for OnenoteSectionId", v.Expected.OnenoteSectionId, actual.OnenoteSectionId)
		}

		if actual.OnenotePageId != v.Expected.OnenotePageId {
			t.Fatalf("Expected %q but got %q for OnenotePageId", v.Expected.OnenotePageId, actual.OnenotePageId)
		}

	}
}

func TestSegmentsForGroupIdOnenoteNotebookIdSectionIdPageId(t *testing.T) {
	segments := GroupIdOnenoteNotebookIdSectionIdPageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdOnenoteNotebookIdSectionIdPageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
