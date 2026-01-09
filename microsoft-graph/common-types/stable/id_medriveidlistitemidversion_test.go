package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListItemIdVersionId{}

func TestNewMeDriveIdListItemIdVersionID(t *testing.T) {
	id := NewMeDriveIdListItemIdVersionID("driveId", "listItemId", "listItemVersionId")

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.ListItemId != "listItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'ListItemId'", id.ListItemId, "listItemId")
	}

	if id.ListItemVersionId != "listItemVersionId" {
		t.Fatalf("Expected %q but got %q for Segment 'ListItemVersionId'", id.ListItemVersionId, "listItemVersionId")
	}
}

func TestFormatMeDriveIdListItemIdVersionID(t *testing.T) {
	actual := NewMeDriveIdListItemIdVersionID("driveId", "listItemId", "listItemVersionId").ID()
	expected := "/me/drives/driveId/list/items/listItemId/versions/listItemVersionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDriveIdListItemIdVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdListItemIdVersionId
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
			Input: "/me/drives/driveId/list/items/listItemId/versions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/list/items/listItemId/versions/listItemVersionId",
			Expected: &MeDriveIdListItemIdVersionId{
				DriveId:           "driveId",
				ListItemId:        "listItemId",
				ListItemVersionId: "listItemVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/list/items/listItemId/versions/listItemVersionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdListItemIdVersionID(v.Input)
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

		if actual.ListItemVersionId != v.Expected.ListItemVersionId {
			t.Fatalf("Expected %q but got %q for ListItemVersionId", v.Expected.ListItemVersionId, actual.ListItemVersionId)
		}

	}
}

func TestParseMeDriveIdListItemIdVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdListItemIdVersionId
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
			Input: "/me/drives/driveId/list/items/listItemId/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT/iTeMs/lIsTiTeMiD/vErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/list/items/listItemId/versions/listItemVersionId",
			Expected: &MeDriveIdListItemIdVersionId{
				DriveId:           "driveId",
				ListItemId:        "listItemId",
				ListItemVersionId: "listItemVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/list/items/listItemId/versions/listItemVersionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT/iTeMs/lIsTiTeMiD/vErSiOnS/lIsTiTeMvErSiOnId",
			Expected: &MeDriveIdListItemIdVersionId{
				DriveId:           "dRiVeId",
				ListItemId:        "lIsTiTeMiD",
				ListItemVersionId: "lIsTiTeMvErSiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT/iTeMs/lIsTiTeMiD/vErSiOnS/lIsTiTeMvErSiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdListItemIdVersionIDInsensitively(v.Input)
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

		if actual.ListItemVersionId != v.Expected.ListItemVersionId {
			t.Fatalf("Expected %q but got %q for ListItemVersionId", v.Expected.ListItemVersionId, actual.ListItemVersionId)
		}

	}
}

func TestSegmentsForMeDriveIdListItemIdVersionId(t *testing.T) {
	segments := MeDriveIdListItemIdVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDriveIdListItemIdVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
