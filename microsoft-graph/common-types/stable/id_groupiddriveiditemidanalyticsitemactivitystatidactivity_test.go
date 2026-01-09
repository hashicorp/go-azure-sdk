package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId{}

func TestNewGroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityID(t *testing.T) {
	id := NewGroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityID("groupId", "driveId", "driveItemId", "itemActivityStatId", "itemActivityId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.DriveItemId != "driveItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveItemId'", id.DriveItemId, "driveItemId")
	}

	if id.ItemActivityStatId != "itemActivityStatId" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemActivityStatId'", id.ItemActivityStatId, "itemActivityStatId")
	}

	if id.ItemActivityId != "itemActivityId" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemActivityId'", id.ItemActivityId, "itemActivityId")
	}
}

func TestFormatGroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityID(t *testing.T) {
	actual := NewGroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityID("groupId", "driveId", "driveItemId", "itemActivityStatId", "itemActivityId").ID()
	expected := "/groups/groupId/drives/driveId/items/driveItemId/analytics/itemActivityStats/itemActivityStatId/activities/itemActivityId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/items",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/items/driveItemId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/items/driveItemId/analytics",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/items/driveItemId/analytics/itemActivityStats",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/items/driveItemId/analytics/itemActivityStats/itemActivityStatId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/items/driveItemId/analytics/itemActivityStats/itemActivityStatId/activities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/drives/driveId/items/driveItemId/analytics/itemActivityStats/itemActivityStatId/activities/itemActivityId",
			Expected: &GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId{
				GroupId:            "groupId",
				DriveId:            "driveId",
				DriveItemId:        "driveItemId",
				ItemActivityStatId: "itemActivityStatId",
				ItemActivityId:     "itemActivityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/drives/driveId/items/driveItemId/analytics/itemActivityStats/itemActivityStatId/activities/itemActivityId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.DriveId != v.Expected.DriveId {
			t.Fatalf("Expected %q but got %q for DriveId", v.Expected.DriveId, actual.DriveId)
		}

		if actual.DriveItemId != v.Expected.DriveItemId {
			t.Fatalf("Expected %q but got %q for DriveItemId", v.Expected.DriveItemId, actual.DriveItemId)
		}

		if actual.ItemActivityStatId != v.Expected.ItemActivityStatId {
			t.Fatalf("Expected %q but got %q for ItemActivityStatId", v.Expected.ItemActivityStatId, actual.ItemActivityStatId)
		}

		if actual.ItemActivityId != v.Expected.ItemActivityId {
			t.Fatalf("Expected %q but got %q for ItemActivityId", v.Expected.ItemActivityId, actual.ItemActivityId)
		}

	}
}

func TestParseGroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/items",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/iTeMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/items/driveItemId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/items/driveItemId/analytics",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/aNaLyTiCs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/items/driveItemId/analytics/itemActivityStats",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/aNaLyTiCs/iTeMaCtIvItYsTaTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/items/driveItemId/analytics/itemActivityStats/itemActivityStatId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/items/driveItemId/analytics/itemActivityStats/itemActivityStatId/activities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiD/aCtIvItIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/drives/driveId/items/driveItemId/analytics/itemActivityStats/itemActivityStatId/activities/itemActivityId",
			Expected: &GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId{
				GroupId:            "groupId",
				DriveId:            "driveId",
				DriveItemId:        "driveItemId",
				ItemActivityStatId: "itemActivityStatId",
				ItemActivityId:     "itemActivityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/drives/driveId/items/driveItemId/analytics/itemActivityStats/itemActivityStatId/activities/itemActivityId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiD/aCtIvItIeS/iTeMaCtIvItYiD",
			Expected: &GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId{
				GroupId:            "gRoUpId",
				DriveId:            "dRiVeId",
				DriveItemId:        "dRiVeItEmId",
				ItemActivityStatId: "iTeMaCtIvItYsTaTiD",
				ItemActivityId:     "iTeMaCtIvItYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiD/aCtIvItIeS/iTeMaCtIvItYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.DriveId != v.Expected.DriveId {
			t.Fatalf("Expected %q but got %q for DriveId", v.Expected.DriveId, actual.DriveId)
		}

		if actual.DriveItemId != v.Expected.DriveItemId {
			t.Fatalf("Expected %q but got %q for DriveItemId", v.Expected.DriveItemId, actual.DriveItemId)
		}

		if actual.ItemActivityStatId != v.Expected.ItemActivityStatId {
			t.Fatalf("Expected %q but got %q for ItemActivityStatId", v.Expected.ItemActivityStatId, actual.ItemActivityStatId)
		}

		if actual.ItemActivityId != v.Expected.ItemActivityId {
			t.Fatalf("Expected %q but got %q for ItemActivityId", v.Expected.ItemActivityId, actual.ItemActivityId)
		}

	}
}

func TestSegmentsForGroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId(t *testing.T) {
	segments := GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
