package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListOperationId{}

func TestNewMeDriveIdListOperationID(t *testing.T) {
	id := NewMeDriveIdListOperationID("driveId", "richLongRunningOperationId")

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.RichLongRunningOperationId != "richLongRunningOperationId" {
		t.Fatalf("Expected %q but got %q for Segment 'RichLongRunningOperationId'", id.RichLongRunningOperationId, "richLongRunningOperationId")
	}
}

func TestFormatMeDriveIdListOperationID(t *testing.T) {
	actual := NewMeDriveIdListOperationID("driveId", "richLongRunningOperationId").ID()
	expected := "/me/drives/driveId/list/operations/richLongRunningOperationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDriveIdListOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdListOperationId
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
			Input: "/me/drives/driveId/list/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/list/operations/richLongRunningOperationId",
			Expected: &MeDriveIdListOperationId{
				DriveId:                    "driveId",
				RichLongRunningOperationId: "richLongRunningOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/list/operations/richLongRunningOperationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdListOperationID(v.Input)
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

		if actual.RichLongRunningOperationId != v.Expected.RichLongRunningOperationId {
			t.Fatalf("Expected %q but got %q for RichLongRunningOperationId", v.Expected.RichLongRunningOperationId, actual.RichLongRunningOperationId)
		}

	}
}

func TestParseMeDriveIdListOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdListOperationId
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
			Input: "/me/drives/driveId/list/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/list/operations/richLongRunningOperationId",
			Expected: &MeDriveIdListOperationId{
				DriveId:                    "driveId",
				RichLongRunningOperationId: "richLongRunningOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/list/operations/richLongRunningOperationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT/oPeRaTiOnS/rIcHlOnGrUnNiNgOpErAtIoNiD",
			Expected: &MeDriveIdListOperationId{
				DriveId:                    "dRiVeId",
				RichLongRunningOperationId: "rIcHlOnGrUnNiNgOpErAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT/oPeRaTiOnS/rIcHlOnGrUnNiNgOpErAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdListOperationIDInsensitively(v.Input)
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

		if actual.RichLongRunningOperationId != v.Expected.RichLongRunningOperationId {
			t.Fatalf("Expected %q but got %q for RichLongRunningOperationId", v.Expected.RichLongRunningOperationId, actual.RichLongRunningOperationId)
		}

	}
}

func TestSegmentsForMeDriveIdListOperationId(t *testing.T) {
	segments := MeDriveIdListOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDriveIdListOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
