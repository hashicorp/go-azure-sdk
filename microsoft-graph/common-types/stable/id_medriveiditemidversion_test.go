package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdItemIdVersionId{}

func TestNewMeDriveIdItemIdVersionID(t *testing.T) {
	id := NewMeDriveIdItemIdVersionID("driveIdValue", "driveItemIdValue", "driveItemVersionIdValue")

	if id.DriveId != "driveIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveIdValue")
	}

	if id.DriveItemId != "driveItemIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveItemId'", id.DriveItemId, "driveItemIdValue")
	}

	if id.DriveItemVersionId != "driveItemVersionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveItemVersionId'", id.DriveItemVersionId, "driveItemVersionIdValue")
	}
}

func TestFormatMeDriveIdItemIdVersionID(t *testing.T) {
	actual := NewMeDriveIdItemIdVersionID("driveIdValue", "driveItemIdValue", "driveItemVersionIdValue").ID()
	expected := "/me/drives/driveIdValue/items/driveItemIdValue/versions/driveItemVersionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDriveIdItemIdVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdItemIdVersionId
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
			Input: "/me/drives/driveIdValue/items/driveItemIdValue/versions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveIdValue/items/driveItemIdValue/versions/driveItemVersionIdValue",
			Expected: &MeDriveIdItemIdVersionId{
				DriveId:            "driveIdValue",
				DriveItemId:        "driveItemIdValue",
				DriveItemVersionId: "driveItemVersionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveIdValue/items/driveItemIdValue/versions/driveItemVersionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdItemIdVersionID(v.Input)
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

		if actual.DriveItemVersionId != v.Expected.DriveItemVersionId {
			t.Fatalf("Expected %q but got %q for DriveItemVersionId", v.Expected.DriveItemVersionId, actual.DriveItemVersionId)
		}

	}
}

func TestParseMeDriveIdItemIdVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdItemIdVersionId
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
			Input: "/me/drives/driveIdValue/items/driveItemIdValue/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/iTeMs/dRiVeItEmIdVaLuE/vErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveIdValue/items/driveItemIdValue/versions/driveItemVersionIdValue",
			Expected: &MeDriveIdItemIdVersionId{
				DriveId:            "driveIdValue",
				DriveItemId:        "driveItemIdValue",
				DriveItemVersionId: "driveItemVersionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveIdValue/items/driveItemIdValue/versions/driveItemVersionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/iTeMs/dRiVeItEmIdVaLuE/vErSiOnS/dRiVeItEmVeRsIoNiDvAlUe",
			Expected: &MeDriveIdItemIdVersionId{
				DriveId:            "dRiVeIdVaLuE",
				DriveItemId:        "dRiVeItEmIdVaLuE",
				DriveItemVersionId: "dRiVeItEmVeRsIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/iTeMs/dRiVeItEmIdVaLuE/vErSiOnS/dRiVeItEmVeRsIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdItemIdVersionIDInsensitively(v.Input)
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

		if actual.DriveItemVersionId != v.Expected.DriveItemVersionId {
			t.Fatalf("Expected %q but got %q for DriveItemVersionId", v.Expected.DriveItemVersionId, actual.DriveItemVersionId)
		}

	}
}

func TestSegmentsForMeDriveIdItemIdVersionId(t *testing.T) {
	segments := MeDriveIdItemIdVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDriveIdItemIdVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
