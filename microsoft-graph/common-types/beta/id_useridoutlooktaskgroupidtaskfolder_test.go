package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOutlookTaskGroupIdTaskFolderId{}

func TestNewUserIdOutlookTaskGroupIdTaskFolderID(t *testing.T) {
	id := NewUserIdOutlookTaskGroupIdTaskFolderID("userId", "outlookTaskGroupId", "outlookTaskFolderId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.OutlookTaskGroupId != "outlookTaskGroupId" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskGroupId'", id.OutlookTaskGroupId, "outlookTaskGroupId")
	}

	if id.OutlookTaskFolderId != "outlookTaskFolderId" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskFolderId'", id.OutlookTaskFolderId, "outlookTaskFolderId")
	}
}

func TestFormatUserIdOutlookTaskGroupIdTaskFolderID(t *testing.T) {
	actual := NewUserIdOutlookTaskGroupIdTaskFolderID("userId", "outlookTaskGroupId", "outlookTaskFolderId").ID()
	expected := "/users/userId/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdOutlookTaskGroupIdTaskFolderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOutlookTaskGroupIdTaskFolderId
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
			Input: "/users/userId/outlook",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook/taskGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook/taskGroups/outlookTaskGroupId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook/taskGroups/outlookTaskGroupId/taskFolders",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId",
			Expected: &UserIdOutlookTaskGroupIdTaskFolderId{
				UserId:              "userId",
				OutlookTaskGroupId:  "outlookTaskGroupId",
				OutlookTaskFolderId: "outlookTaskFolderId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOutlookTaskGroupIdTaskFolderID(v.Input)
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

		if actual.OutlookTaskGroupId != v.Expected.OutlookTaskGroupId {
			t.Fatalf("Expected %q but got %q for OutlookTaskGroupId", v.Expected.OutlookTaskGroupId, actual.OutlookTaskGroupId)
		}

		if actual.OutlookTaskFolderId != v.Expected.OutlookTaskFolderId {
			t.Fatalf("Expected %q but got %q for OutlookTaskFolderId", v.Expected.OutlookTaskFolderId, actual.OutlookTaskFolderId)
		}

	}
}

func TestParseUserIdOutlookTaskGroupIdTaskFolderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOutlookTaskGroupIdTaskFolderId
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
			Input: "/users/userId/outlook",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook/taskGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk/tAsKgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook/taskGroups/outlookTaskGroupId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook/taskGroups/outlookTaskGroupId/taskFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiD/tAsKfOlDeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId",
			Expected: &UserIdOutlookTaskGroupIdTaskFolderId{
				UserId:              "userId",
				OutlookTaskGroupId:  "outlookTaskGroupId",
				OutlookTaskFolderId: "outlookTaskFolderId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiD/tAsKfOlDeRs/oUtLoOkTaSkFoLdErId",
			Expected: &UserIdOutlookTaskGroupIdTaskFolderId{
				UserId:              "uSeRiD",
				OutlookTaskGroupId:  "oUtLoOkTaSkGrOuPiD",
				OutlookTaskFolderId: "oUtLoOkTaSkFoLdErId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiD/tAsKfOlDeRs/oUtLoOkTaSkFoLdErId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOutlookTaskGroupIdTaskFolderIDInsensitively(v.Input)
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

		if actual.OutlookTaskGroupId != v.Expected.OutlookTaskGroupId {
			t.Fatalf("Expected %q but got %q for OutlookTaskGroupId", v.Expected.OutlookTaskGroupId, actual.OutlookTaskGroupId)
		}

		if actual.OutlookTaskFolderId != v.Expected.OutlookTaskFolderId {
			t.Fatalf("Expected %q but got %q for OutlookTaskFolderId", v.Expected.OutlookTaskFolderId, actual.OutlookTaskFolderId)
		}

	}
}

func TestSegmentsForUserIdOutlookTaskGroupIdTaskFolderId(t *testing.T) {
	segments := UserIdOutlookTaskGroupIdTaskFolderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdOutlookTaskGroupIdTaskFolderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
