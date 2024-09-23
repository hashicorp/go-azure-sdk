package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnenoteNotebookIdSectionId{}

func TestNewUserIdOnenoteNotebookIdSectionID(t *testing.T) {
	id := NewUserIdOnenoteNotebookIdSectionID("userId", "notebookId", "onenoteSectionId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.NotebookId != "notebookId" {
		t.Fatalf("Expected %q but got %q for Segment 'NotebookId'", id.NotebookId, "notebookId")
	}

	if id.OnenoteSectionId != "onenoteSectionId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenoteSectionId'", id.OnenoteSectionId, "onenoteSectionId")
	}
}

func TestFormatUserIdOnenoteNotebookIdSectionID(t *testing.T) {
	actual := NewUserIdOnenoteNotebookIdSectionID("userId", "notebookId", "onenoteSectionId").ID()
	expected := "/users/userId/onenote/notebooks/notebookId/sections/onenoteSectionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdOnenoteNotebookIdSectionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOnenoteNotebookIdSectionId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onenote",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onenote/notebooks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onenote/notebooks/notebookId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onenote/notebooks/notebookId/sections",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/onenote/notebooks/notebookId/sections/onenoteSectionId",
			Expected: &UserIdOnenoteNotebookIdSectionId{
				UserId:           "userId",
				NotebookId:       "notebookId",
				OnenoteSectionId: "onenoteSectionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/onenote/notebooks/notebookId/sections/onenoteSectionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOnenoteNotebookIdSectionID(v.Input)
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

		if actual.OnenoteSectionId != v.Expected.OnenoteSectionId {
			t.Fatalf("Expected %q but got %q for OnenoteSectionId", v.Expected.OnenoteSectionId, actual.OnenoteSectionId)
		}

	}
}

func TestParseUserIdOnenoteNotebookIdSectionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOnenoteNotebookIdSectionId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onenote",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNeNoTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onenote/notebooks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNeNoTe/nOtEbOoKs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onenote/notebooks/notebookId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNeNoTe/nOtEbOoKs/nOtEbOoKiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onenote/notebooks/notebookId/sections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNeNoTe/nOtEbOoKs/nOtEbOoKiD/sEcTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/onenote/notebooks/notebookId/sections/onenoteSectionId",
			Expected: &UserIdOnenoteNotebookIdSectionId{
				UserId:           "userId",
				NotebookId:       "notebookId",
				OnenoteSectionId: "onenoteSectionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/onenote/notebooks/notebookId/sections/onenoteSectionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNeNoTe/nOtEbOoKs/nOtEbOoKiD/sEcTiOnS/oNeNoTeSeCtIoNiD",
			Expected: &UserIdOnenoteNotebookIdSectionId{
				UserId:           "uSeRiD",
				NotebookId:       "nOtEbOoKiD",
				OnenoteSectionId: "oNeNoTeSeCtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNeNoTe/nOtEbOoKs/nOtEbOoKiD/sEcTiOnS/oNeNoTeSeCtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOnenoteNotebookIdSectionIDInsensitively(v.Input)
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

		if actual.OnenoteSectionId != v.Expected.OnenoteSectionId {
			t.Fatalf("Expected %q but got %q for OnenoteSectionId", v.Expected.OnenoteSectionId, actual.OnenoteSectionId)
		}

	}
}

func TestSegmentsForUserIdOnenoteNotebookIdSectionId(t *testing.T) {
	segments := UserIdOnenoteNotebookIdSectionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdOnenoteNotebookIdSectionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
