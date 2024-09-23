package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdItemIdListItemPermissionId{}

func TestNewMeDriveIdItemIdListItemPermissionID(t *testing.T) {
	id := NewMeDriveIdItemIdListItemPermissionID("driveId", "driveItemId", "permissionId")

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

func TestFormatMeDriveIdItemIdListItemPermissionID(t *testing.T) {
	actual := NewMeDriveIdItemIdListItemPermissionID("driveId", "driveItemId", "permissionId").ID()
	expected := "/me/drives/driveId/items/driveItemId/listItem/permissions/permissionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDriveIdItemIdListItemPermissionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdItemIdListItemPermissionId
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
			Input: "/me/drives",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/items",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/items/driveItemId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/items/driveItemId/listItem",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/items/driveItemId/listItem/permissions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/items/driveItemId/listItem/permissions/permissionId",
			Expected: &MeDriveIdItemIdListItemPermissionId{
				DriveId:      "driveId",
				DriveItemId:  "driveItemId",
				PermissionId: "permissionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/items/driveItemId/listItem/permissions/permissionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdItemIdListItemPermissionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
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

func TestParseMeDriveIdItemIdListItemPermissionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdItemIdListItemPermissionId
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
			Input: "/me/drives",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/items",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/iTeMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/items/driveItemId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/items/driveItemId/listItem",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/lIsTiTeM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/items/driveItemId/listItem/permissions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/lIsTiTeM/pErMiSsIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/items/driveItemId/listItem/permissions/permissionId",
			Expected: &MeDriveIdItemIdListItemPermissionId{
				DriveId:      "driveId",
				DriveItemId:  "driveItemId",
				PermissionId: "permissionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/items/driveItemId/listItem/permissions/permissionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/lIsTiTeM/pErMiSsIoNs/pErMiSsIoNiD",
			Expected: &MeDriveIdItemIdListItemPermissionId{
				DriveId:      "dRiVeId",
				DriveItemId:  "dRiVeItEmId",
				PermissionId: "pErMiSsIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/lIsTiTeM/pErMiSsIoNs/pErMiSsIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdItemIdListItemPermissionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
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

func TestSegmentsForMeDriveIdItemIdListItemPermissionId(t *testing.T) {
	segments := MeDriveIdItemIdListItemPermissionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDriveIdItemIdListItemPermissionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
