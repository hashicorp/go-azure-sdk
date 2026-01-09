package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdItemIdListItemActivityId{}

func TestNewUserIdDriveIdItemIdListItemActivityID(t *testing.T) {
	id := NewUserIdDriveIdItemIdListItemActivityID("userId", "driveId", "driveItemId", "itemActivityOLDId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.DriveItemId != "driveItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveItemId'", id.DriveItemId, "driveItemId")
	}

	if id.ItemActivityOLDId != "itemActivityOLDId" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemActivityOLDId'", id.ItemActivityOLDId, "itemActivityOLDId")
	}
}

func TestFormatUserIdDriveIdItemIdListItemActivityID(t *testing.T) {
	actual := NewUserIdDriveIdItemIdListItemActivityID("userId", "driveId", "driveItemId", "itemActivityOLDId").ID()
	expected := "/users/userId/drives/driveId/items/driveItemId/listItem/activities/itemActivityOLDId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdDriveIdItemIdListItemActivityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdItemIdListItemActivityId
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
			Input: "/users/userId/drives",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/items",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/items/driveItemId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem/activities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem/activities/itemActivityOLDId",
			Expected: &UserIdDriveIdItemIdListItemActivityId{
				UserId:            "userId",
				DriveId:           "driveId",
				DriveItemId:       "driveItemId",
				ItemActivityOLDId: "itemActivityOLDId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem/activities/itemActivityOLDId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdItemIdListItemActivityID(v.Input)
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

		if actual.DriveId != v.Expected.DriveId {
			t.Fatalf("Expected %q but got %q for DriveId", v.Expected.DriveId, actual.DriveId)
		}

		if actual.DriveItemId != v.Expected.DriveItemId {
			t.Fatalf("Expected %q but got %q for DriveItemId", v.Expected.DriveItemId, actual.DriveItemId)
		}

		if actual.ItemActivityOLDId != v.Expected.ItemActivityOLDId {
			t.Fatalf("Expected %q but got %q for ItemActivityOLDId", v.Expected.ItemActivityOLDId, actual.ItemActivityOLDId)
		}

	}
}

func TestParseUserIdDriveIdItemIdListItemActivityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdItemIdListItemActivityId
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
			Input: "/users/userId/drives",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/items",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/iTeMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/items/driveItemId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/lIsTiTeM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem/activities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/lIsTiTeM/aCtIvItIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem/activities/itemActivityOLDId",
			Expected: &UserIdDriveIdItemIdListItemActivityId{
				UserId:            "userId",
				DriveId:           "driveId",
				DriveItemId:       "driveItemId",
				ItemActivityOLDId: "itemActivityOLDId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem/activities/itemActivityOLDId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/lIsTiTeM/aCtIvItIeS/iTeMaCtIvItYoLdId",
			Expected: &UserIdDriveIdItemIdListItemActivityId{
				UserId:            "uSeRiD",
				DriveId:           "dRiVeId",
				DriveItemId:       "dRiVeItEmId",
				ItemActivityOLDId: "iTeMaCtIvItYoLdId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/lIsTiTeM/aCtIvItIeS/iTeMaCtIvItYoLdId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdItemIdListItemActivityIDInsensitively(v.Input)
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

		if actual.DriveId != v.Expected.DriveId {
			t.Fatalf("Expected %q but got %q for DriveId", v.Expected.DriveId, actual.DriveId)
		}

		if actual.DriveItemId != v.Expected.DriveItemId {
			t.Fatalf("Expected %q but got %q for DriveItemId", v.Expected.DriveItemId, actual.DriveItemId)
		}

		if actual.ItemActivityOLDId != v.Expected.ItemActivityOLDId {
			t.Fatalf("Expected %q but got %q for ItemActivityOLDId", v.Expected.ItemActivityOLDId, actual.ItemActivityOLDId)
		}

	}
}

func TestSegmentsForUserIdDriveIdItemIdListItemActivityId(t *testing.T) {
	segments := UserIdDriveIdItemIdListItemActivityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdDriveIdItemIdListItemActivityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
