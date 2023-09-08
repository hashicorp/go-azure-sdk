package useroutlooktaskfolder

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOutlookTaskFolderId{}

func TestNewUserOutlookTaskFolderID(t *testing.T) {
	id := NewUserOutlookTaskFolderID("userIdValue", "outlookTaskFolderIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.OutlookTaskFolderId != "outlookTaskFolderIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskFolderId'", id.OutlookTaskFolderId, "outlookTaskFolderIdValue")
	}
}

func TestFormatUserOutlookTaskFolderID(t *testing.T) {
	actual := NewUserOutlookTaskFolderID("userIdValue", "outlookTaskFolderIdValue").ID()
	expected := "/users/userIdValue/outlook/taskFolders/outlookTaskFolderIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserOutlookTaskFolderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOutlookTaskFolderId
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
			Input: "/users/userIdValue/outlook",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/outlook/taskFolders",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/outlook/taskFolders/outlookTaskFolderIdValue",
			Expected: &UserOutlookTaskFolderId{
				UserId:              "userIdValue",
				OutlookTaskFolderId: "outlookTaskFolderIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/outlook/taskFolders/outlookTaskFolderIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOutlookTaskFolderID(v.Input)
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

		if actual.OutlookTaskFolderId != v.Expected.OutlookTaskFolderId {
			t.Fatalf("Expected %q but got %q for OutlookTaskFolderId", v.Expected.OutlookTaskFolderId, actual.OutlookTaskFolderId)
		}

	}
}

func TestParseUserOutlookTaskFolderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOutlookTaskFolderId
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
			Input: "/users/userIdValue/outlook",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oUtLoOk",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/outlook/taskFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oUtLoOk/tAsKfOlDeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/outlook/taskFolders/outlookTaskFolderIdValue",
			Expected: &UserOutlookTaskFolderId{
				UserId:              "userIdValue",
				OutlookTaskFolderId: "outlookTaskFolderIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/outlook/taskFolders/outlookTaskFolderIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oUtLoOk/tAsKfOlDeRs/oUtLoOkTaSkFoLdErIdVaLuE",
			Expected: &UserOutlookTaskFolderId{
				UserId:              "uSeRiDvAlUe",
				OutlookTaskFolderId: "oUtLoOkTaSkFoLdErIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oUtLoOk/tAsKfOlDeRs/oUtLoOkTaSkFoLdErIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOutlookTaskFolderIDInsensitively(v.Input)
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

		if actual.OutlookTaskFolderId != v.Expected.OutlookTaskFolderId {
			t.Fatalf("Expected %q but got %q for OutlookTaskFolderId", v.Expected.OutlookTaskFolderId, actual.OutlookTaskFolderId)
		}

	}
}

func TestSegmentsForUserOutlookTaskFolderId(t *testing.T) {
	segments := UserOutlookTaskFolderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserOutlookTaskFolderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
