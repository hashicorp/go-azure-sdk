package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdItemIdListItemPermissionId{}

func TestNewUserIdDriveIdItemIdListItemPermissionID(t *testing.T) {
	id := NewUserIdDriveIdItemIdListItemPermissionID("userId", "driveId", "driveItemId", "permissionId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.DriveItemId != "driveItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveItemId'", id.DriveItemId, "driveItemId")
	}

	if id.PermissionId != "permissionId" {
		t.Fatalf("Expected %q but got %q for Segment 'PermissionId'", id.PermissionId, "permissionId")
	}
}

func TestFormatUserIdDriveIdItemIdListItemPermissionID(t *testing.T) {
	actual := NewUserIdDriveIdItemIdListItemPermissionID("userId", "driveId", "driveItemId", "permissionId").ID()
	expected := "/users/userId/drives/driveId/items/driveItemId/listItem/permissions/permissionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdDriveIdItemIdListItemPermissionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdItemIdListItemPermissionId
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
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem/permissions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem/permissions/permissionId",
			Expected: &UserIdDriveIdItemIdListItemPermissionId{
				UserId:       "userId",
				DriveId:      "driveId",
				DriveItemId:  "driveItemId",
				PermissionId: "permissionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem/permissions/permissionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdItemIdListItemPermissionID(v.Input)
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

		if actual.PermissionId != v.Expected.PermissionId {
			t.Fatalf("Expected %q but got %q for PermissionId", v.Expected.PermissionId, actual.PermissionId)
		}

	}
}

func TestParseUserIdDriveIdItemIdListItemPermissionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdItemIdListItemPermissionId
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
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem/permissions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/lIsTiTeM/pErMiSsIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem/permissions/permissionId",
			Expected: &UserIdDriveIdItemIdListItemPermissionId{
				UserId:       "userId",
				DriveId:      "driveId",
				DriveItemId:  "driveItemId",
				PermissionId: "permissionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem/permissions/permissionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/lIsTiTeM/pErMiSsIoNs/pErMiSsIoNiD",
			Expected: &UserIdDriveIdItemIdListItemPermissionId{
				UserId:       "uSeRiD",
				DriveId:      "dRiVeId",
				DriveItemId:  "dRiVeItEmId",
				PermissionId: "pErMiSsIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/lIsTiTeM/pErMiSsIoNs/pErMiSsIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdItemIdListItemPermissionIDInsensitively(v.Input)
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

		if actual.PermissionId != v.Expected.PermissionId {
			t.Fatalf("Expected %q but got %q for PermissionId", v.Expected.PermissionId, actual.PermissionId)
		}

	}
}

func TestSegmentsForUserIdDriveIdItemIdListItemPermissionId(t *testing.T) {
	segments := UserIdDriveIdItemIdListItemPermissionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdDriveIdItemIdListItemPermissionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
