package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListItemIdActivityId{}

func TestNewMeDriveIdListItemIdActivityID(t *testing.T) {
	id := NewMeDriveIdListItemIdActivityID("driveId", "listItemId", "itemActivityOLDId")

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.ListItemId != "listItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'ListItemId'", id.ListItemId, "listItemId")
	}

	if id.ItemActivityOLDId != "itemActivityOLDId" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemActivityOLDId'", id.ItemActivityOLDId, "itemActivityOLDId")
	}
}

func TestFormatMeDriveIdListItemIdActivityID(t *testing.T) {
	actual := NewMeDriveIdListItemIdActivityID("driveId", "listItemId", "itemActivityOLDId").ID()
	expected := "/me/drives/driveId/list/items/listItemId/activities/itemActivityOLDId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDriveIdListItemIdActivityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdListItemIdActivityId
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
			Input: "/me/drives/driveId/list",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/list/items",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/list/items/listItemId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/list/items/listItemId/activities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/list/items/listItemId/activities/itemActivityOLDId",
			Expected: &MeDriveIdListItemIdActivityId{
				DriveId:           "driveId",
				ListItemId:        "listItemId",
				ItemActivityOLDId: "itemActivityOLDId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/list/items/listItemId/activities/itemActivityOLDId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdListItemIdActivityID(v.Input)
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

		if actual.ListItemId != v.Expected.ListItemId {
			t.Fatalf("Expected %q but got %q for ListItemId", v.Expected.ListItemId, actual.ListItemId)
		}

		if actual.ItemActivityOLDId != v.Expected.ItemActivityOLDId {
			t.Fatalf("Expected %q but got %q for ItemActivityOLDId", v.Expected.ItemActivityOLDId, actual.ItemActivityOLDId)
		}

	}
}

func TestParseMeDriveIdListItemIdActivityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdListItemIdActivityId
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
			Input: "/me/drives/driveId/list",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/list/items",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT/iTeMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/list/items/listItemId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT/iTeMs/lIsTiTeMiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/list/items/listItemId/activities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT/iTeMs/lIsTiTeMiD/aCtIvItIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/list/items/listItemId/activities/itemActivityOLDId",
			Expected: &MeDriveIdListItemIdActivityId{
				DriveId:           "driveId",
				ListItemId:        "listItemId",
				ItemActivityOLDId: "itemActivityOLDId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/list/items/listItemId/activities/itemActivityOLDId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT/iTeMs/lIsTiTeMiD/aCtIvItIeS/iTeMaCtIvItYoLdId",
			Expected: &MeDriveIdListItemIdActivityId{
				DriveId:           "dRiVeId",
				ListItemId:        "lIsTiTeMiD",
				ItemActivityOLDId: "iTeMaCtIvItYoLdId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT/iTeMs/lIsTiTeMiD/aCtIvItIeS/iTeMaCtIvItYoLdId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdListItemIdActivityIDInsensitively(v.Input)
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

		if actual.ListItemId != v.Expected.ListItemId {
			t.Fatalf("Expected %q but got %q for ListItemId", v.Expected.ListItemId, actual.ListItemId)
		}

		if actual.ItemActivityOLDId != v.Expected.ItemActivityOLDId {
			t.Fatalf("Expected %q but got %q for ItemActivityOLDId", v.Expected.ItemActivityOLDId, actual.ItemActivityOLDId)
		}

	}
}

func TestSegmentsForMeDriveIdListItemIdActivityId(t *testing.T) {
	segments := MeDriveIdListItemIdActivityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDriveIdListItemIdActivityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
