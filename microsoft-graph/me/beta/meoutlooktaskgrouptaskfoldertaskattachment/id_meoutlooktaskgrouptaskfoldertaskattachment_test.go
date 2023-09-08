package meoutlooktaskgrouptaskfoldertaskattachment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOutlookTaskGroupTaskFolderTaskAttachmentId{}

func TestNewMeOutlookTaskGroupTaskFolderTaskAttachmentID(t *testing.T) {
	id := NewMeOutlookTaskGroupTaskFolderTaskAttachmentID("outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue", "attachmentIdValue")

	if id.OutlookTaskGroupId != "outlookTaskGroupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskGroupId'", id.OutlookTaskGroupId, "outlookTaskGroupIdValue")
	}

	if id.OutlookTaskFolderId != "outlookTaskFolderIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskFolderId'", id.OutlookTaskFolderId, "outlookTaskFolderIdValue")
	}

	if id.OutlookTaskId != "outlookTaskIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskId'", id.OutlookTaskId, "outlookTaskIdValue")
	}

	if id.AttachmentId != "attachmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AttachmentId'", id.AttachmentId, "attachmentIdValue")
	}
}

func TestFormatMeOutlookTaskGroupTaskFolderTaskAttachmentID(t *testing.T) {
	actual := NewMeOutlookTaskGroupTaskFolderTaskAttachmentID("outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue", "attachmentIdValue").ID()
	expected := "/me/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue/tasks/outlookTaskIdValue/attachments/attachmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOutlookTaskGroupTaskFolderTaskAttachmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOutlookTaskGroupTaskFolderTaskAttachmentId
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
			Input: "/me/outlook",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskGroups/outlookTaskGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue/tasks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue/tasks/outlookTaskIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue/tasks/outlookTaskIdValue/attachments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue/tasks/outlookTaskIdValue/attachments/attachmentIdValue",
			Expected: &MeOutlookTaskGroupTaskFolderTaskAttachmentId{
				OutlookTaskGroupId:  "outlookTaskGroupIdValue",
				OutlookTaskFolderId: "outlookTaskFolderIdValue",
				OutlookTaskId:       "outlookTaskIdValue",
				AttachmentId:        "attachmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue/tasks/outlookTaskIdValue/attachments/attachmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOutlookTaskGroupTaskFolderTaskAttachmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
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

func TestParseMeOutlookTaskGroupTaskFolderTaskAttachmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOutlookTaskGroupTaskFolderTaskAttachmentId
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
			Input: "/me/outlook",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskGroups/outlookTaskGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiDvAlUe/tAsKfOlDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiDvAlUe/tAsKfOlDeRs/oUtLoOkTaSkFoLdErIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiDvAlUe/tAsKfOlDeRs/oUtLoOkTaSkFoLdErIdVaLuE/tAsKs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue/tasks/outlookTaskIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiDvAlUe/tAsKfOlDeRs/oUtLoOkTaSkFoLdErIdVaLuE/tAsKs/oUtLoOkTaSkIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue/tasks/outlookTaskIdValue/attachments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiDvAlUe/tAsKfOlDeRs/oUtLoOkTaSkFoLdErIdVaLuE/tAsKs/oUtLoOkTaSkIdVaLuE/aTtAcHmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue/tasks/outlookTaskIdValue/attachments/attachmentIdValue",
			Expected: &MeOutlookTaskGroupTaskFolderTaskAttachmentId{
				OutlookTaskGroupId:  "outlookTaskGroupIdValue",
				OutlookTaskFolderId: "outlookTaskFolderIdValue",
				OutlookTaskId:       "outlookTaskIdValue",
				AttachmentId:        "attachmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/outlook/taskGroups/outlookTaskGroupIdValue/taskFolders/outlookTaskFolderIdValue/tasks/outlookTaskIdValue/attachments/attachmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiDvAlUe/tAsKfOlDeRs/oUtLoOkTaSkFoLdErIdVaLuE/tAsKs/oUtLoOkTaSkIdVaLuE/aTtAcHmEnTs/aTtAcHmEnTiDvAlUe",
			Expected: &MeOutlookTaskGroupTaskFolderTaskAttachmentId{
				OutlookTaskGroupId:  "oUtLoOkTaSkGrOuPiDvAlUe",
				OutlookTaskFolderId: "oUtLoOkTaSkFoLdErIdVaLuE",
				OutlookTaskId:       "oUtLoOkTaSkIdVaLuE",
				AttachmentId:        "aTtAcHmEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiDvAlUe/tAsKfOlDeRs/oUtLoOkTaSkFoLdErIdVaLuE/tAsKs/oUtLoOkTaSkIdVaLuE/aTtAcHmEnTs/aTtAcHmEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOutlookTaskGroupTaskFolderTaskAttachmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
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

func TestSegmentsForMeOutlookTaskGroupTaskFolderTaskAttachmentId(t *testing.T) {
	segments := MeOutlookTaskGroupTaskFolderTaskAttachmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOutlookTaskGroupTaskFolderTaskAttachmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
