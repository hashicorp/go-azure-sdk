package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdItemIdChildId{}

func TestNewMeDriveIdItemIdChildID(t *testing.T) {
	id := NewMeDriveIdItemIdChildID("driveId", "driveItemId", "driveItemId1")

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.DriveItemId != "driveItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveItemId'", id.DriveItemId, "driveItemId")
	}

	if id.DriveItemId1 != "driveItemId1" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveItemId1'", id.DriveItemId1, "driveItemId1")
	}
}

func TestFormatMeDriveIdItemIdChildID(t *testing.T) {
	actual := NewMeDriveIdItemIdChildID("driveId", "driveItemId", "driveItemId1").ID()
	expected := "/me/drives/driveId/items/driveItemId/children/driveItemId1"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDriveIdItemIdChildID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdItemIdChildId
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
			Input: "/me/drives/driveId/items/driveItemId/children",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/items/driveItemId/children/driveItemId1",
			Expected: &MeDriveIdItemIdChildId{
				DriveId:      "driveId",
				DriveItemId:  "driveItemId",
				DriveItemId1: "driveItemId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/items/driveItemId/children/driveItemId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdItemIdChildID(v.Input)
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

		if actual.DriveItemId1 != v.Expected.DriveItemId1 {
			t.Fatalf("Expected %q but got %q for DriveItemId1", v.Expected.DriveItemId1, actual.DriveItemId1)
		}

	}
}

func TestParseMeDriveIdItemIdChildIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdItemIdChildId
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
			Input: "/me/drives/driveId/items/driveItemId/children",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/cHiLdReN",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/items/driveItemId/children/driveItemId1",
			Expected: &MeDriveIdItemIdChildId{
				DriveId:      "driveId",
				DriveItemId:  "driveItemId",
				DriveItemId1: "driveItemId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/items/driveItemId/children/driveItemId1/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/cHiLdReN/dRiVeItEmId1",
			Expected: &MeDriveIdItemIdChildId{
				DriveId:      "dRiVeId",
				DriveItemId:  "dRiVeItEmId",
				DriveItemId1: "dRiVeItEmId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/cHiLdReN/dRiVeItEmId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdItemIdChildIDInsensitively(v.Input)
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

		if actual.DriveItemId1 != v.Expected.DriveItemId1 {
			t.Fatalf("Expected %q but got %q for DriveItemId1", v.Expected.DriveItemId1, actual.DriveItemId1)
		}

	}
}

func TestSegmentsForMeDriveIdItemIdChildId(t *testing.T) {
	segments := MeDriveIdItemIdChildId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDriveIdItemIdChildId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
