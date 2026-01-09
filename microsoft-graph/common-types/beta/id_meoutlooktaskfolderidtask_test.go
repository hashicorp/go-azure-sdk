package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOutlookTaskFolderIdTaskId{}

func TestNewMeOutlookTaskFolderIdTaskID(t *testing.T) {
	id := NewMeOutlookTaskFolderIdTaskID("outlookTaskFolderId", "outlookTaskId")

	if id.OutlookTaskFolderId != "outlookTaskFolderId" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskFolderId'", id.OutlookTaskFolderId, "outlookTaskFolderId")
	}

	if id.OutlookTaskId != "outlookTaskId" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskId'", id.OutlookTaskId, "outlookTaskId")
	}
}

func TestFormatMeOutlookTaskFolderIdTaskID(t *testing.T) {
	actual := NewMeOutlookTaskFolderIdTaskID("outlookTaskFolderId", "outlookTaskId").ID()
	expected := "/me/outlook/taskFolders/outlookTaskFolderId/tasks/outlookTaskId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOutlookTaskFolderIdTaskID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOutlookTaskFolderIdTaskId
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
			// Valid URI
			Input: "/me/outlook/taskFolders/outlookTaskFolderId/tasks/outlookTaskId",
			Expected: &MeOutlookTaskFolderIdTaskId{
				OutlookTaskFolderId: "outlookTaskFolderId",
				OutlookTaskId:       "outlookTaskId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/outlook/taskFolders/outlookTaskFolderId/tasks/outlookTaskId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOutlookTaskFolderIdTaskID(v.Input)
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

	}
}

func TestParseMeOutlookTaskFolderIdTaskIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOutlookTaskFolderIdTaskId
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
			// Valid URI
			Input: "/me/outlook/taskFolders/outlookTaskFolderId/tasks/outlookTaskId",
			Expected: &MeOutlookTaskFolderIdTaskId{
				OutlookTaskFolderId: "outlookTaskFolderId",
				OutlookTaskId:       "outlookTaskId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/outlook/taskFolders/outlookTaskFolderId/tasks/outlookTaskId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKfOlDeRs/oUtLoOkTaSkFoLdErId/tAsKs/oUtLoOkTaSkId",
			Expected: &MeOutlookTaskFolderIdTaskId{
				OutlookTaskFolderId: "oUtLoOkTaSkFoLdErId",
				OutlookTaskId:       "oUtLoOkTaSkId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKfOlDeRs/oUtLoOkTaSkFoLdErId/tAsKs/oUtLoOkTaSkId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOutlookTaskFolderIdTaskIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForMeOutlookTaskFolderIdTaskId(t *testing.T) {
	segments := MeOutlookTaskFolderIdTaskId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOutlookTaskFolderIdTaskId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
