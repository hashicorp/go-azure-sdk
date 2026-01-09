package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdRootAnalyticsItemActivityStatId{}

func TestNewGroupIdDriveIdRootAnalyticsItemActivityStatID(t *testing.T) {
	id := NewGroupIdDriveIdRootAnalyticsItemActivityStatID("groupId", "driveId", "itemActivityStatId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.ItemActivityStatId != "itemActivityStatId" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemActivityStatId'", id.ItemActivityStatId, "itemActivityStatId")
	}
}

func TestFormatGroupIdDriveIdRootAnalyticsItemActivityStatID(t *testing.T) {
	actual := NewGroupIdDriveIdRootAnalyticsItemActivityStatID("groupId", "driveId", "itemActivityStatId").ID()
	expected := "/groups/groupId/drives/driveId/root/analytics/itemActivityStats/itemActivityStatId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdDriveIdRootAnalyticsItemActivityStatID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdDriveIdRootAnalyticsItemActivityStatId
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
			Input: "/groups/groupId/drives/driveId/root",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/root/analytics",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/root/analytics/itemActivityStats",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/drives/driveId/root/analytics/itemActivityStats/itemActivityStatId",
			Expected: &GroupIdDriveIdRootAnalyticsItemActivityStatId{
				GroupId:            "groupId",
				DriveId:            "driveId",
				ItemActivityStatId: "itemActivityStatId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/drives/driveId/root/analytics/itemActivityStats/itemActivityStatId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdDriveIdRootAnalyticsItemActivityStatID(v.Input)
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

		if actual.ItemActivityStatId != v.Expected.ItemActivityStatId {
			t.Fatalf("Expected %q but got %q for ItemActivityStatId", v.Expected.ItemActivityStatId, actual.ItemActivityStatId)
		}

	}
}

func TestParseGroupIdDriveIdRootAnalyticsItemActivityStatIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdDriveIdRootAnalyticsItemActivityStatId
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
			Input: "/groups/groupId/drives/driveId/root",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/rOoT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/root/analytics",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/rOoT/aNaLyTiCs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/root/analytics/itemActivityStats",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/rOoT/aNaLyTiCs/iTeMaCtIvItYsTaTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/drives/driveId/root/analytics/itemActivityStats/itemActivityStatId",
			Expected: &GroupIdDriveIdRootAnalyticsItemActivityStatId{
				GroupId:            "groupId",
				DriveId:            "driveId",
				ItemActivityStatId: "itemActivityStatId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/drives/driveId/root/analytics/itemActivityStats/itemActivityStatId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/rOoT/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiD",
			Expected: &GroupIdDriveIdRootAnalyticsItemActivityStatId{
				GroupId:            "gRoUpId",
				DriveId:            "dRiVeId",
				ItemActivityStatId: "iTeMaCtIvItYsTaTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/rOoT/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdDriveIdRootAnalyticsItemActivityStatIDInsensitively(v.Input)
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

		if actual.ItemActivityStatId != v.Expected.ItemActivityStatId {
			t.Fatalf("Expected %q but got %q for ItemActivityStatId", v.Expected.ItemActivityStatId, actual.ItemActivityStatId)
		}

	}
}

func TestSegmentsForGroupIdDriveIdRootAnalyticsItemActivityStatId(t *testing.T) {
	segments := GroupIdDriveIdRootAnalyticsItemActivityStatId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdDriveIdRootAnalyticsItemActivityStatId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
