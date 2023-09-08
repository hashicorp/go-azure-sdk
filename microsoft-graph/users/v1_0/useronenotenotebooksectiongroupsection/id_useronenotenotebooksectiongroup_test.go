package useronenotenotebooksectiongroupsection

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteNotebookSectionGroupId{}

func TestNewUserOnenoteNotebookSectionGroupID(t *testing.T) {
	id := NewUserOnenoteNotebookSectionGroupID("userIdValue", "notebookIdValue", "sectionGroupIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.NotebookId != "notebookIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'NotebookId'", id.NotebookId, "notebookIdValue")
	}

	if id.SectionGroupId != "sectionGroupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SectionGroupId'", id.SectionGroupId, "sectionGroupIdValue")
	}
}

func TestFormatUserOnenoteNotebookSectionGroupID(t *testing.T) {
	actual := NewUserOnenoteNotebookSectionGroupID("userIdValue", "notebookIdValue", "sectionGroupIdValue").ID()
	expected := "/users/userIdValue/onenote/notebooks/notebookIdValue/sectionGroups/sectionGroupIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserOnenoteNotebookSectionGroupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOnenoteNotebookSectionGroupId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote/notebooks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote/notebooks/notebookIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote/notebooks/notebookIdValue/sectionGroups",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onenote/notebooks/notebookIdValue/sectionGroups/sectionGroupIdValue",
			Expected: &UserOnenoteNotebookSectionGroupId{
				UserId:         "userIdValue",
				NotebookId:     "notebookIdValue",
				SectionGroupId: "sectionGroupIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onenote/notebooks/notebookIdValue/sectionGroups/sectionGroupIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOnenoteNotebookSectionGroupID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.NotebookId != v.Expected.NotebookId {
			t.Fatalf("Expected %q but got %q for NotebookId", v.Expected.NotebookId, actual.NotebookId)
		}

		if actual.SectionGroupId != v.Expected.SectionGroupId {
			t.Fatalf("Expected %q but got %q for SectionGroupId", v.Expected.SectionGroupId, actual.SectionGroupId)
		}

	}
}

func TestParseUserOnenoteNotebookSectionGroupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOnenoteNotebookSectionGroupId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote/notebooks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/nOtEbOoKs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote/notebooks/notebookIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/nOtEbOoKs/nOtEbOoKiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote/notebooks/notebookIdValue/sectionGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/nOtEbOoKs/nOtEbOoKiDvAlUe/sEcTiOnGrOuPs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onenote/notebooks/notebookIdValue/sectionGroups/sectionGroupIdValue",
			Expected: &UserOnenoteNotebookSectionGroupId{
				UserId:         "userIdValue",
				NotebookId:     "notebookIdValue",
				SectionGroupId: "sectionGroupIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onenote/notebooks/notebookIdValue/sectionGroups/sectionGroupIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/nOtEbOoKs/nOtEbOoKiDvAlUe/sEcTiOnGrOuPs/sEcTiOnGrOuPiDvAlUe",
			Expected: &UserOnenoteNotebookSectionGroupId{
				UserId:         "uSeRiDvAlUe",
				NotebookId:     "nOtEbOoKiDvAlUe",
				SectionGroupId: "sEcTiOnGrOuPiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/nOtEbOoKs/nOtEbOoKiDvAlUe/sEcTiOnGrOuPs/sEcTiOnGrOuPiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOnenoteNotebookSectionGroupIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.NotebookId != v.Expected.NotebookId {
			t.Fatalf("Expected %q but got %q for NotebookId", v.Expected.NotebookId, actual.NotebookId)
		}

		if actual.SectionGroupId != v.Expected.SectionGroupId {
			t.Fatalf("Expected %q but got %q for SectionGroupId", v.Expected.SectionGroupId, actual.SectionGroupId)
		}

	}
}

func TestSegmentsForUserOnenoteNotebookSectionGroupId(t *testing.T) {
	segments := UserOnenoteNotebookSectionGroupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserOnenoteNotebookSectionGroupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
