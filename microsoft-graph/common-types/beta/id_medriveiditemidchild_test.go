package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdItemIdChildId{}

func TestNewMeDriveIdItemIdChildID(t *testing.T) {
	id := NewMeDriveIdItemIdChildID("driveIdValue", "driveItemIdValue", "driveItemId1Value")

	if id.DriveId != "driveIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveIdValue")
	}

	if id.DriveItemId != "driveItemIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveItemId'", id.DriveItemId, "driveItemIdValue")
	}

	if id.DriveItemId1 != "driveItemId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveItemId1'", id.DriveItemId1, "driveItemId1Value")
	}
}

func TestFormatMeDriveIdItemIdChildID(t *testing.T) {
	actual := NewMeDriveIdItemIdChildID("driveIdValue", "driveItemIdValue", "driveItemId1Value").ID()
	expected := "/me/drives/driveIdValue/items/driveItemIdValue/children/driveItemId1Value"
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
			Input: "/me/drives/driveIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/items",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/items/driveItemIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/items/driveItemIdValue/children",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveIdValue/items/driveItemIdValue/children/driveItemId1Value",
			Expected: &MeDriveIdItemIdChildId{
				DriveId:      "driveIdValue",
				DriveItemId:  "driveItemIdValue",
				DriveItemId1: "driveItemId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveIdValue/items/driveItemIdValue/children/driveItemId1Value/extra",
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
			Input: "/me/drives/driveIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/items",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/iTeMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/items/driveItemIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/iTeMs/dRiVeItEmIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/items/driveItemIdValue/children",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/iTeMs/dRiVeItEmIdVaLuE/cHiLdReN",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveIdValue/items/driveItemIdValue/children/driveItemId1Value",
			Expected: &MeDriveIdItemIdChildId{
				DriveId:      "driveIdValue",
				DriveItemId:  "driveItemIdValue",
				DriveItemId1: "driveItemId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveIdValue/items/driveItemIdValue/children/driveItemId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/iTeMs/dRiVeItEmIdVaLuE/cHiLdReN/dRiVeItEmId1vAlUe",
			Expected: &MeDriveIdItemIdChildId{
				DriveId:      "dRiVeIdVaLuE",
				DriveItemId:  "dRiVeItEmIdVaLuE",
				DriveItemId1: "dRiVeItEmId1vAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/iTeMs/dRiVeItEmIdVaLuE/cHiLdReN/dRiVeItEmId1vAlUe/extra",
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
