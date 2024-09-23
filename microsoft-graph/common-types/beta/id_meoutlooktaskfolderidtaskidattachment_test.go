package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOutlookTaskFolderIdTaskIdAttachmentId{}

func TestNewMeOutlookTaskFolderIdTaskIdAttachmentID(t *testing.T) {
	id := NewMeOutlookTaskFolderIdTaskIdAttachmentID("outlookTaskFolderId", "outlookTaskId", "attachmentId")

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

func TestFormatMeOutlookTaskFolderIdTaskIdAttachmentID(t *testing.T) {
	actual := NewMeOutlookTaskFolderIdTaskIdAttachmentID("outlookTaskFolderId", "outlookTaskId", "attachmentId").ID()
	expected := "/me/outlook/taskFolders/outlookTaskFolderId/tasks/outlookTaskId/attachments/attachmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOutlookTaskFolderIdTaskIdAttachmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOutlookTaskFolderIdTaskIdAttachmentId
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
			Input: "/me/outlook/taskFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskFolders/outlookTaskFolderId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskFolders/outlookTaskFolderId/tasks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskFolders/outlookTaskFolderId/tasks/outlookTaskId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskFolders/outlookTaskFolderId/tasks/outlookTaskId/attachments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/outlook/taskFolders/outlookTaskFolderId/tasks/outlookTaskId/attachments/attachmentId",
			Expected: &MeOutlookTaskFolderIdTaskIdAttachmentId{
				OutlookTaskFolderId: "outlookTaskFolderId",
				OutlookTaskId:       "outlookTaskId",
				AttachmentId:        "attachmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/outlook/taskFolders/outlookTaskFolderId/tasks/outlookTaskId/attachments/attachmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOutlookTaskFolderIdTaskIdAttachmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
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

func TestParseMeOutlookTaskFolderIdTaskIdAttachmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOutlookTaskFolderIdTaskIdAttachmentId
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
			Input: "/me/outlook/taskFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKfOlDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskFolders/outlookTaskFolderId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKfOlDeRs/oUtLoOkTaSkFoLdErId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskFolders/outlookTaskFolderId/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKfOlDeRs/oUtLoOkTaSkFoLdErId/tAsKs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskFolders/outlookTaskFolderId/tasks/outlookTaskId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKfOlDeRs/oUtLoOkTaSkFoLdErId/tAsKs/oUtLoOkTaSkId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskFolders/outlookTaskFolderId/tasks/outlookTaskId/attachments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKfOlDeRs/oUtLoOkTaSkFoLdErId/tAsKs/oUtLoOkTaSkId/aTtAcHmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/outlook/taskFolders/outlookTaskFolderId/tasks/outlookTaskId/attachments/attachmentId",
			Expected: &MeOutlookTaskFolderIdTaskIdAttachmentId{
				OutlookTaskFolderId: "outlookTaskFolderId",
				OutlookTaskId:       "outlookTaskId",
				AttachmentId:        "attachmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/outlook/taskFolders/outlookTaskFolderId/tasks/outlookTaskId/attachments/attachmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKfOlDeRs/oUtLoOkTaSkFoLdErId/tAsKs/oUtLoOkTaSkId/aTtAcHmEnTs/aTtAcHmEnTiD",
			Expected: &MeOutlookTaskFolderIdTaskIdAttachmentId{
				OutlookTaskFolderId: "oUtLoOkTaSkFoLdErId",
				OutlookTaskId:       "oUtLoOkTaSkId",
				AttachmentId:        "aTtAcHmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKfOlDeRs/oUtLoOkTaSkFoLdErId/tAsKs/oUtLoOkTaSkId/aTtAcHmEnTs/aTtAcHmEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOutlookTaskFolderIdTaskIdAttachmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
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

func TestSegmentsForMeOutlookTaskFolderIdTaskIdAttachmentId(t *testing.T) {
	segments := MeOutlookTaskFolderIdTaskIdAttachmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOutlookTaskFolderIdTaskIdAttachmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
