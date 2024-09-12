package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdRootListItemActivityId{}

func TestNewGroupIdDriveIdRootListItemActivityID(t *testing.T) {
	id := NewGroupIdDriveIdRootListItemActivityID("groupIdValue", "driveIdValue", "itemActivityOLDIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.DriveId != "driveIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveIdValue")
	}

	if id.ItemActivityOLDId != "itemActivityOLDIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemActivityOLDId'", id.ItemActivityOLDId, "itemActivityOLDIdValue")
	}
}

func TestFormatGroupIdDriveIdRootListItemActivityID(t *testing.T) {
	actual := NewGroupIdDriveIdRootListItemActivityID("groupIdValue", "driveIdValue", "itemActivityOLDIdValue").ID()
	expected := "/groups/groupIdValue/drives/driveIdValue/root/listItem/activities/itemActivityOLDIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdDriveIdRootListItemActivityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdDriveIdRootListItemActivityId
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
			Input: "/groups/groupIdValue/drives/driveIdValue/root",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/root/listItem",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/root/listItem/activities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/drives/driveIdValue/root/listItem/activities/itemActivityOLDIdValue",
			Expected: &GroupIdDriveIdRootListItemActivityId{
				GroupId:           "groupIdValue",
				DriveId:           "driveIdValue",
				ItemActivityOLDId: "itemActivityOLDIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/drives/driveIdValue/root/listItem/activities/itemActivityOLDIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdDriveIdRootListItemActivityID(v.Input)
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

		if actual.ItemActivityOLDId != v.Expected.ItemActivityOLDId {
			t.Fatalf("Expected %q but got %q for ItemActivityOLDId", v.Expected.ItemActivityOLDId, actual.ItemActivityOLDId)
		}

	}
}

func TestParseGroupIdDriveIdRootListItemActivityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdDriveIdRootListItemActivityId
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
			Input: "/groups/groupIdValue/drives/driveIdValue/root",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE/rOoT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/root/listItem",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE/rOoT/lIsTiTeM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/root/listItem/activities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE/rOoT/lIsTiTeM/aCtIvItIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/drives/driveIdValue/root/listItem/activities/itemActivityOLDIdValue",
			Expected: &GroupIdDriveIdRootListItemActivityId{
				GroupId:           "groupIdValue",
				DriveId:           "driveIdValue",
				ItemActivityOLDId: "itemActivityOLDIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/drives/driveIdValue/root/listItem/activities/itemActivityOLDIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE/rOoT/lIsTiTeM/aCtIvItIeS/iTeMaCtIvItYoLdIdVaLuE",
			Expected: &GroupIdDriveIdRootListItemActivityId{
				GroupId:           "gRoUpIdVaLuE",
				DriveId:           "dRiVeIdVaLuE",
				ItemActivityOLDId: "iTeMaCtIvItYoLdIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE/rOoT/lIsTiTeM/aCtIvItIeS/iTeMaCtIvItYoLdIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdDriveIdRootListItemActivityIDInsensitively(v.Input)
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

		if actual.ItemActivityOLDId != v.Expected.ItemActivityOLDId {
			t.Fatalf("Expected %q but got %q for ItemActivityOLDId", v.Expected.ItemActivityOLDId, actual.ItemActivityOLDId)
		}

	}
}

func TestSegmentsForGroupIdDriveIdRootListItemActivityId(t *testing.T) {
	segments := GroupIdDriveIdRootListItemActivityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdDriveIdRootListItemActivityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
