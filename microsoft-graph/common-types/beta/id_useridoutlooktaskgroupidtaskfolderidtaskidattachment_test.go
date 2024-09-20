package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId{}

func TestNewUserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID(t *testing.T) {
	id := NewUserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID("userId", "outlookTaskGroupId", "outlookTaskFolderId", "outlookTaskId", "attachmentId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.OutlookTaskGroupId != "outlookTaskGroupId" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskGroupId'", id.OutlookTaskGroupId, "outlookTaskGroupId")
	}

	if id.OutlookTaskFolderId != "outlookTaskFolderId" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskFolderId'", id.OutlookTaskFolderId, "outlookTaskFolderId")
	}

	if id.OutlookTaskId != "outlookTaskId" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskId'", id.OutlookTaskId, "outlookTaskId")
	}

	if id.AttachmentId != "attachmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'AttachmentId'", id.AttachmentId, "attachmentId")
	}
}

func TestFormatUserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID(t *testing.T) {
	actual := NewUserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID("userId", "outlookTaskGroupId", "outlookTaskFolderId", "outlookTaskId", "attachmentId").ID()
	expected := "/users/userId/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId/tasks/outlookTaskId/attachments/attachmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId
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
			// Incomplete URI
			Input: "/users/userId/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId/tasks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId/tasks/outlookTaskId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId/tasks/outlookTaskId/attachments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId/tasks/outlookTaskId/attachments/attachmentId",
			Expected: &UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId{
				UserId:              "userId",
				OutlookTaskGroupId:  "outlookTaskGroupId",
				OutlookTaskFolderId: "outlookTaskFolderId",
				OutlookTaskId:       "outlookTaskId",
				AttachmentId:        "attachmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId/tasks/outlookTaskId/attachments/attachmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID(v.Input)
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

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestParseUserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId
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
			// Incomplete URI
			Input: "/users/userId/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiD/tAsKfOlDeRs/oUtLoOkTaSkFoLdErId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiD/tAsKfOlDeRs/oUtLoOkTaSkFoLdErId/tAsKs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId/tasks/outlookTaskId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiD/tAsKfOlDeRs/oUtLoOkTaSkFoLdErId/tAsKs/oUtLoOkTaSkId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId/tasks/outlookTaskId/attachments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiD/tAsKfOlDeRs/oUtLoOkTaSkFoLdErId/tAsKs/oUtLoOkTaSkId/aTtAcHmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId/tasks/outlookTaskId/attachments/attachmentId",
			Expected: &UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId{
				UserId:              "userId",
				OutlookTaskGroupId:  "outlookTaskGroupId",
				OutlookTaskFolderId: "outlookTaskFolderId",
				OutlookTaskId:       "outlookTaskId",
				AttachmentId:        "attachmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId/tasks/outlookTaskId/attachments/attachmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiD/tAsKfOlDeRs/oUtLoOkTaSkFoLdErId/tAsKs/oUtLoOkTaSkId/aTtAcHmEnTs/aTtAcHmEnTiD",
			Expected: &UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId{
				UserId:              "uSeRiD",
				OutlookTaskGroupId:  "oUtLoOkTaSkGrOuPiD",
				OutlookTaskFolderId: "oUtLoOkTaSkFoLdErId",
				OutlookTaskId:       "oUtLoOkTaSkId",
				AttachmentId:        "aTtAcHmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiD/tAsKfOlDeRs/oUtLoOkTaSkFoLdErId/tAsKs/oUtLoOkTaSkId/aTtAcHmEnTs/aTtAcHmEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentIDInsensitively(v.Input)
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

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestSegmentsForUserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId(t *testing.T) {
	segments := UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
