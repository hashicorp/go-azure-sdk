package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdListOperationId{}

func TestNewUserIdDriveIdListOperationID(t *testing.T) {
	id := NewUserIdDriveIdListOperationID("userId", "driveId", "richLongRunningOperationId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.RichLongRunningOperationId != "richLongRunningOperationId" {
		t.Fatalf("Expected %q but got %q for Segment 'RichLongRunningOperationId'", id.RichLongRunningOperationId, "richLongRunningOperationId")
	}
}

func TestFormatUserIdDriveIdListOperationID(t *testing.T) {
	actual := NewUserIdDriveIdListOperationID("userId", "driveId", "richLongRunningOperationId").ID()
	expected := "/users/userId/drives/driveId/list/operations/richLongRunningOperationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdDriveIdListOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdListOperationId
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
			Input: "/users/userId/drives/driveId/list",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/list/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/list/operations/richLongRunningOperationId",
			Expected: &UserIdDriveIdListOperationId{
				UserId:                     "userId",
				DriveId:                    "driveId",
				RichLongRunningOperationId: "richLongRunningOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/list/operations/richLongRunningOperationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdListOperationID(v.Input)
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

		if actual.RichLongRunningOperationId != v.Expected.RichLongRunningOperationId {
			t.Fatalf("Expected %q but got %q for RichLongRunningOperationId", v.Expected.RichLongRunningOperationId, actual.RichLongRunningOperationId)
		}

	}
}

func TestParseUserIdDriveIdListOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdListOperationId
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
			Input: "/users/userId/drives/driveId/list",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/lIsT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/list/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/lIsT/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/list/operations/richLongRunningOperationId",
			Expected: &UserIdDriveIdListOperationId{
				UserId:                     "userId",
				DriveId:                    "driveId",
				RichLongRunningOperationId: "richLongRunningOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/list/operations/richLongRunningOperationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/lIsT/oPeRaTiOnS/rIcHlOnGrUnNiNgOpErAtIoNiD",
			Expected: &UserIdDriveIdListOperationId{
				UserId:                     "uSeRiD",
				DriveId:                    "dRiVeId",
				RichLongRunningOperationId: "rIcHlOnGrUnNiNgOpErAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/lIsT/oPeRaTiOnS/rIcHlOnGrUnNiNgOpErAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdListOperationIDInsensitively(v.Input)
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

		if actual.RichLongRunningOperationId != v.Expected.RichLongRunningOperationId {
			t.Fatalf("Expected %q but got %q for RichLongRunningOperationId", v.Expected.RichLongRunningOperationId, actual.RichLongRunningOperationId)
		}

	}
}

func TestSegmentsForUserIdDriveIdListOperationId(t *testing.T) {
	segments := UserIdDriveIdListOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdDriveIdListOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
