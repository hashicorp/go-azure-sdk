package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdListItemId{}

func TestNewUserIdDriveIdListItemID(t *testing.T) {
	id := NewUserIdDriveIdListItemID("userId", "driveId", "listItemId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.ListItemId != "listItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'ListItemId'", id.ListItemId, "listItemId")
	}
}

func TestFormatUserIdDriveIdListItemID(t *testing.T) {
	actual := NewUserIdDriveIdListItemID("userId", "driveId", "listItemId").ID()
	expected := "/users/userId/drives/driveId/list/items/listItemId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdDriveIdListItemID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdListItemId
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
			Input: "/users/userId/drives/driveId/list",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/list/items",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/list/items/listItemId",
			Expected: &UserIdDriveIdListItemId{
				UserId:     "userId",
				DriveId:    "driveId",
				ListItemId: "listItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/list/items/listItemId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdListItemID(v.Input)
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

		if actual.ListItemId != v.Expected.ListItemId {
			t.Fatalf("Expected %q but got %q for ListItemId", v.Expected.ListItemId, actual.ListItemId)
		}

	}
}

func TestParseUserIdDriveIdListItemIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdListItemId
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
			Input: "/users/userId/drives/driveId/list",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/lIsT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/list/items",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/lIsT/iTeMs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/list/items/listItemId",
			Expected: &UserIdDriveIdListItemId{
				UserId:     "userId",
				DriveId:    "driveId",
				ListItemId: "listItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/list/items/listItemId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/lIsT/iTeMs/lIsTiTeMiD",
			Expected: &UserIdDriveIdListItemId{
				UserId:     "uSeRiD",
				DriveId:    "dRiVeId",
				ListItemId: "lIsTiTeMiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/lIsT/iTeMs/lIsTiTeMiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdListItemIDInsensitively(v.Input)
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

		if actual.ListItemId != v.Expected.ListItemId {
			t.Fatalf("Expected %q but got %q for ListItemId", v.Expected.ListItemId, actual.ListItemId)
		}

	}
}

func TestSegmentsForUserIdDriveIdListItemId(t *testing.T) {
	segments := UserIdDriveIdListItemId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdDriveIdListItemId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
