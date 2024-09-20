package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOutlookTaskGroupIdTaskFolderId{}

func TestNewMeOutlookTaskGroupIdTaskFolderID(t *testing.T) {
	id := NewMeOutlookTaskGroupIdTaskFolderID("outlookTaskGroupId", "outlookTaskFolderId")

	if id.OutlookTaskGroupId != "outlookTaskGroupId" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskGroupId'", id.OutlookTaskGroupId, "outlookTaskGroupId")
	}

	if id.OutlookTaskFolderId != "outlookTaskFolderId" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskFolderId'", id.OutlookTaskFolderId, "outlookTaskFolderId")
	}
}

func TestFormatMeOutlookTaskGroupIdTaskFolderID(t *testing.T) {
	actual := NewMeOutlookTaskGroupIdTaskFolderID("outlookTaskGroupId", "outlookTaskFolderId").ID()
	expected := "/me/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOutlookTaskGroupIdTaskFolderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOutlookTaskGroupIdTaskFolderId
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
			Input: "/me/outlook/taskGroups/outlookTaskGroupId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskGroups/outlookTaskGroupId/taskFolders",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId",
			Expected: &MeOutlookTaskGroupIdTaskFolderId{
				OutlookTaskGroupId:  "outlookTaskGroupId",
				OutlookTaskFolderId: "outlookTaskFolderId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOutlookTaskGroupIdTaskFolderID(v.Input)
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

	}
}

func TestParseMeOutlookTaskGroupIdTaskFolderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOutlookTaskGroupIdTaskFolderId
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
			Input: "/me/outlook/taskGroups/outlookTaskGroupId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/taskGroups/outlookTaskGroupId/taskFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiD/tAsKfOlDeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId",
			Expected: &MeOutlookTaskGroupIdTaskFolderId{
				OutlookTaskGroupId:  "outlookTaskGroupId",
				OutlookTaskFolderId: "outlookTaskFolderId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/outlook/taskGroups/outlookTaskGroupId/taskFolders/outlookTaskFolderId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiD/tAsKfOlDeRs/oUtLoOkTaSkFoLdErId",
			Expected: &MeOutlookTaskGroupIdTaskFolderId{
				OutlookTaskGroupId:  "oUtLoOkTaSkGrOuPiD",
				OutlookTaskFolderId: "oUtLoOkTaSkFoLdErId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKgRoUpS/oUtLoOkTaSkGrOuPiD/tAsKfOlDeRs/oUtLoOkTaSkFoLdErId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOutlookTaskGroupIdTaskFolderIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForMeOutlookTaskGroupIdTaskFolderId(t *testing.T) {
	segments := MeOutlookTaskGroupIdTaskFolderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOutlookTaskGroupIdTaskFolderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
