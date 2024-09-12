package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdListItemIdPermissionId{}

func TestNewGroupIdDriveIdListItemIdPermissionID(t *testing.T) {
	id := NewGroupIdDriveIdListItemIdPermissionID("groupIdValue", "driveIdValue", "listItemIdValue", "permissionIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.DriveId != "driveIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveIdValue")
	}

	if id.ListItemId != "listItemIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ListItemId'", id.ListItemId, "listItemIdValue")
	}

	if id.PermissionId != "permissionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PermissionId'", id.PermissionId, "permissionIdValue")
	}
}

func TestFormatGroupIdDriveIdListItemIdPermissionID(t *testing.T) {
	actual := NewGroupIdDriveIdListItemIdPermissionID("groupIdValue", "driveIdValue", "listItemIdValue", "permissionIdValue").ID()
	expected := "/groups/groupIdValue/drives/driveIdValue/list/items/listItemIdValue/permissions/permissionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdDriveIdListItemIdPermissionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdDriveIdListItemIdPermissionId
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
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/list",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/list/items",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/list/items/listItemIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/list/items/listItemIdValue/permissions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/drives/driveIdValue/list/items/listItemIdValue/permissions/permissionIdValue",
			Expected: &GroupIdDriveIdListItemIdPermissionId{
				GroupId:      "groupIdValue",
				DriveId:      "driveIdValue",
				ListItemId:   "listItemIdValue",
				PermissionId: "permissionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/drives/driveIdValue/list/items/listItemIdValue/permissions/permissionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdDriveIdListItemIdPermissionID(v.Input)
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

		if actual.ListItemId != v.Expected.ListItemId {
			t.Fatalf("Expected %q but got %q for ListItemId", v.Expected.ListItemId, actual.ListItemId)
		}

		if actual.PermissionId != v.Expected.PermissionId {
			t.Fatalf("Expected %q but got %q for PermissionId", v.Expected.PermissionId, actual.PermissionId)
		}

	}
}

func TestParseGroupIdDriveIdListItemIdPermissionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdDriveIdListItemIdPermissionId
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
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/list",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE/lIsT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/list/items",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE/lIsT/iTeMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/list/items/listItemIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE/lIsT/iTeMs/lIsTiTeMiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/list/items/listItemIdValue/permissions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE/lIsT/iTeMs/lIsTiTeMiDvAlUe/pErMiSsIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/drives/driveIdValue/list/items/listItemIdValue/permissions/permissionIdValue",
			Expected: &GroupIdDriveIdListItemIdPermissionId{
				GroupId:      "groupIdValue",
				DriveId:      "driveIdValue",
				ListItemId:   "listItemIdValue",
				PermissionId: "permissionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/drives/driveIdValue/list/items/listItemIdValue/permissions/permissionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE/lIsT/iTeMs/lIsTiTeMiDvAlUe/pErMiSsIoNs/pErMiSsIoNiDvAlUe",
			Expected: &GroupIdDriveIdListItemIdPermissionId{
				GroupId:      "gRoUpIdVaLuE",
				DriveId:      "dRiVeIdVaLuE",
				ListItemId:   "lIsTiTeMiDvAlUe",
				PermissionId: "pErMiSsIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE/lIsT/iTeMs/lIsTiTeMiDvAlUe/pErMiSsIoNs/pErMiSsIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdDriveIdListItemIdPermissionIDInsensitively(v.Input)
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

		if actual.ListItemId != v.Expected.ListItemId {
			t.Fatalf("Expected %q but got %q for ListItemId", v.Expected.ListItemId, actual.ListItemId)
		}

		if actual.PermissionId != v.Expected.PermissionId {
			t.Fatalf("Expected %q but got %q for PermissionId", v.Expected.PermissionId, actual.PermissionId)
		}

	}
}

func TestSegmentsForGroupIdDriveIdListItemIdPermissionId(t *testing.T) {
	segments := GroupIdDriveIdListItemIdPermissionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdDriveIdListItemIdPermissionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
