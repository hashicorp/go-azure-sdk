package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdListOperationId{}

func TestNewUserIdDriveIdListOperationID(t *testing.T) {
	id := NewUserIdDriveIdListOperationID("userIdValue", "driveIdValue", "richLongRunningOperationIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.DriveId != "driveIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveIdValue")
	}

	if id.RichLongRunningOperationId != "richLongRunningOperationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RichLongRunningOperationId'", id.RichLongRunningOperationId, "richLongRunningOperationIdValue")
	}
}

func TestFormatUserIdDriveIdListOperationID(t *testing.T) {
	actual := NewUserIdDriveIdListOperationID("userIdValue", "driveIdValue", "richLongRunningOperationIdValue").ID()
	expected := "/users/userIdValue/drives/driveIdValue/list/operations/richLongRunningOperationIdValue"
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/drives",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/drives/driveIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/drives/driveIdValue/list",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/drives/driveIdValue/list/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/drives/driveIdValue/list/operations/richLongRunningOperationIdValue",
			Expected: &UserIdDriveIdListOperationId{
				UserId:                     "userIdValue",
				DriveId:                    "driveIdValue",
				RichLongRunningOperationId: "richLongRunningOperationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/drives/driveIdValue/list/operations/richLongRunningOperationIdValue/extra",
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/drives",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dRiVeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/drives/driveIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dRiVeS/dRiVeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/drives/driveIdValue/list",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dRiVeS/dRiVeIdVaLuE/lIsT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/drives/driveIdValue/list/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dRiVeS/dRiVeIdVaLuE/lIsT/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/drives/driveIdValue/list/operations/richLongRunningOperationIdValue",
			Expected: &UserIdDriveIdListOperationId{
				UserId:                     "userIdValue",
				DriveId:                    "driveIdValue",
				RichLongRunningOperationId: "richLongRunningOperationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/drives/driveIdValue/list/operations/richLongRunningOperationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dRiVeS/dRiVeIdVaLuE/lIsT/oPeRaTiOnS/rIcHlOnGrUnNiNgOpErAtIoNiDvAlUe",
			Expected: &UserIdDriveIdListOperationId{
				UserId:                     "uSeRiDvAlUe",
				DriveId:                    "dRiVeIdVaLuE",
				RichLongRunningOperationId: "rIcHlOnGrUnNiNgOpErAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dRiVeS/dRiVeIdVaLuE/lIsT/oPeRaTiOnS/rIcHlOnGrUnNiNgOpErAtIoNiDvAlUe/extra",
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
