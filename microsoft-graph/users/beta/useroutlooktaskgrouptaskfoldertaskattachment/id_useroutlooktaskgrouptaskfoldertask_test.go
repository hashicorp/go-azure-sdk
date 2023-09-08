package useroutlooktaskgrouptaskfoldertaskattachment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOutlookTaskGroupTaskFolderTaskId{}

func TestNewUserOutlookTaskGroupTaskFolderTaskID(t *testing.T) {
	id := NewUserOutlookTaskGroupTaskFolderTaskID("userIdValue", "outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.OutlookTaskGroupId != "outlookTaskGroupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskGroupId'", id.OutlookTaskGroupId, "outlookTaskGroupIdValue")
	}

	if id.OutlookTaskFolderId != "outlookTaskFolderIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskFolderId'", id.OutlookTaskFolderId, "outlookTaskFolderIdValue")
	}

	if id.OutlookTaskId != "outlookTaskIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskId'", id.OutlookTaskId, "outlookTaskIdValue")
	}
}

func TestFormatUserOutlookTaskGroupTaskFolderTaskID(t *testing.T) {
	actual := NewUserOutlookTaskGroupTaskFolderTaskID("userIdValue", "outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue").ID()
	expected := "/users/userIdValue/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue/tasks/outlookTaskIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserOutlookTaskGroupTaskFolderTaskID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOutlookTaskGroupTaskFolderTaskId
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
			Input: "/users/userIdValue/outlook/taskGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/outlook/taskGroups/outlookTaskGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue/tasks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue/tasks/outlookTaskIdValue",
			Expected: &UserOutlookTaskGroupTaskFolderTaskId{
				UserId:              "userIdValue",
				OutlookTaskGroupId:  "outlookTaskGroupIdValue",
				OutlookTaskFolderId: "outlookTaskFolderIdValue",
				OutlookTaskId:       "outlookTaskIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue/tasks/outlookTaskIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOutlookTaskGroupTaskFolderTaskID(v.Input)
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

		if actual.OutlookTaskId != v.Expected.OutlookTaskId {
			t.Fatalf("Expected %q but got %q for OutlookTaskId", v.Expected.OutlookTaskId, actual.OutlookTaskId)
		}

	}
}

func TestParseUserOutlookTaskGroupTaskFolderTaskIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOutlookTaskGroupTaskFolderTaskId
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
			Input: "/users/userIdValue/outlook/taskGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oUtLoOk/tAsKgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/outlook/taskGroups/outlookTaskGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiDvAlUe/tAsKfOlDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiDvAlUe/tAsKfOlDeRs/oUtLoOkTaSkFoLdErIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiDvAlUe/tAsKfOlDeRs/oUtLoOkTaSkFoLdErIdVaLuE/tAsKs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue/tasks/outlookTaskIdValue",
			Expected: &UserOutlookTaskGroupTaskFolderTaskId{
				UserId:              "userIdValue",
				OutlookTaskGroupId:  "outlookTaskGroupIdValue",
				OutlookTaskFolderId: "outlookTaskFolderIdValue",
				OutlookTaskId:       "outlookTaskIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue/tasks/outlookTaskIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiDvAlUe/tAsKfOlDeRs/oUtLoOkTaSkFoLdErIdVaLuE/tAsKs/oUtLoOkTaSkIdVaLuE",
			Expected: &UserOutlookTaskGroupTaskFolderTaskId{
				UserId:              "uSeRiDvAlUe",
				OutlookTaskGroupId:  "oUtLoOkTaSkGrOuPiDvAlUe",
				OutlookTaskFolderId: "oUtLoOkTaSkFoLdErIdVaLuE",
				OutlookTaskId:       "oUtLoOkTaSkIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiDvAlUe/tAsKfOlDeRs/oUtLoOkTaSkFoLdErIdVaLuE/tAsKs/oUtLoOkTaSkIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOutlookTaskGroupTaskFolderTaskIDInsensitively(v.Input)
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

		if actual.OutlookTaskId != v.Expected.OutlookTaskId {
			t.Fatalf("Expected %q but got %q for OutlookTaskId", v.Expected.OutlookTaskId, actual.OutlookTaskId)
		}

	}
}

func TestSegmentsForUserOutlookTaskGroupTaskFolderTaskId(t *testing.T) {
	segments := UserOutlookTaskGroupTaskFolderTaskId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserOutlookTaskGroupTaskFolderTaskId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
