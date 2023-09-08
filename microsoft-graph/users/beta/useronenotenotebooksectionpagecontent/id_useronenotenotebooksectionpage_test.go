package useronenotenotebooksectionpagecontent

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteNotebookSectionPageId{}

func TestNewUserOnenoteNotebookSectionPageID(t *testing.T) {
	id := NewUserOnenoteNotebookSectionPageID("userIdValue", "notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.NotebookId != "notebookIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'NotebookId'", id.NotebookId, "notebookIdValue")
	}

	if id.OnenoteSectionId != "onenoteSectionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenoteSectionId'", id.OnenoteSectionId, "onenoteSectionIdValue")
	}

	if id.OnenotePageId != "onenotePageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenotePageId'", id.OnenotePageId, "onenotePageIdValue")
	}
}

func TestFormatUserOnenoteNotebookSectionPageID(t *testing.T) {
	actual := NewUserOnenoteNotebookSectionPageID("userIdValue", "notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue").ID()
	expected := "/users/userIdValue/onenote/notebooks/notebookIdValue/sections/onenoteSectionIdValue/pages/onenotePageIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserOnenoteNotebookSectionPageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOnenoteNotebookSectionPageId
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
			Input: "/users/userIdValue/onenote/notebooks/notebookIdValue/sections",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote/notebooks/notebookIdValue/sections/onenoteSectionIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote/notebooks/notebookIdValue/sections/onenoteSectionIdValue/pages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onenote/notebooks/notebookIdValue/sections/onenoteSectionIdValue/pages/onenotePageIdValue",
			Expected: &UserOnenoteNotebookSectionPageId{
				UserId:           "userIdValue",
				NotebookId:       "notebookIdValue",
				OnenoteSectionId: "onenoteSectionIdValue",
				OnenotePageId:    "onenotePageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onenote/notebooks/notebookIdValue/sections/onenoteSectionIdValue/pages/onenotePageIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOnenoteNotebookSectionPageID(v.Input)
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

		if actual.OnenotePageId != v.Expected.OnenotePageId {
			t.Fatalf("Expected %q but got %q for OnenotePageId", v.Expected.OnenotePageId, actual.OnenotePageId)
		}

	}
}

func TestParseUserOnenoteNotebookSectionPageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOnenoteNotebookSectionPageId
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
			Input: "/users/userIdValue/onenote/notebooks/notebookIdValue/sections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/nOtEbOoKs/nOtEbOoKiDvAlUe/sEcTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote/notebooks/notebookIdValue/sections/onenoteSectionIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/nOtEbOoKs/nOtEbOoKiDvAlUe/sEcTiOnS/oNeNoTeSeCtIoNiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote/notebooks/notebookIdValue/sections/onenoteSectionIdValue/pages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/nOtEbOoKs/nOtEbOoKiDvAlUe/sEcTiOnS/oNeNoTeSeCtIoNiDvAlUe/pAgEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onenote/notebooks/notebookIdValue/sections/onenoteSectionIdValue/pages/onenotePageIdValue",
			Expected: &UserOnenoteNotebookSectionPageId{
				UserId:           "userIdValue",
				NotebookId:       "notebookIdValue",
				OnenoteSectionId: "onenoteSectionIdValue",
				OnenotePageId:    "onenotePageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onenote/notebooks/notebookIdValue/sections/onenoteSectionIdValue/pages/onenotePageIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/nOtEbOoKs/nOtEbOoKiDvAlUe/sEcTiOnS/oNeNoTeSeCtIoNiDvAlUe/pAgEs/oNeNoTePaGeIdVaLuE",
			Expected: &UserOnenoteNotebookSectionPageId{
				UserId:           "uSeRiDvAlUe",
				NotebookId:       "nOtEbOoKiDvAlUe",
				OnenoteSectionId: "oNeNoTeSeCtIoNiDvAlUe",
				OnenotePageId:    "oNeNoTePaGeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/nOtEbOoKs/nOtEbOoKiDvAlUe/sEcTiOnS/oNeNoTeSeCtIoNiDvAlUe/pAgEs/oNeNoTePaGeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOnenoteNotebookSectionPageIDInsensitively(v.Input)
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

		if actual.OnenotePageId != v.Expected.OnenotePageId {
			t.Fatalf("Expected %q but got %q for OnenotePageId", v.Expected.OnenotePageId, actual.OnenotePageId)
		}

	}
}

func TestSegmentsForUserOnenoteNotebookSectionPageId(t *testing.T) {
	segments := UserOnenoteNotebookSectionPageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserOnenoteNotebookSectionPageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
