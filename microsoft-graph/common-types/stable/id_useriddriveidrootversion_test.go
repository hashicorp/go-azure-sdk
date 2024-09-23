package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdRootVersionId{}

func TestNewUserIdDriveIdRootVersionID(t *testing.T) {
	id := NewUserIdDriveIdRootVersionID("userId", "driveId", "driveItemVersionId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.DriveItemVersionId != "driveItemVersionId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveItemVersionId'", id.DriveItemVersionId, "driveItemVersionId")
	}
}

func TestFormatUserIdDriveIdRootVersionID(t *testing.T) {
	actual := NewUserIdDriveIdRootVersionID("userId", "driveId", "driveItemVersionId").ID()
	expected := "/users/userId/drives/driveId/root/versions/driveItemVersionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdDriveIdRootVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdRootVersionId
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
			Input: "/users/userId/drives/driveId/root",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/root/versions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/root/versions/driveItemVersionId",
			Expected: &UserIdDriveIdRootVersionId{
				UserId:             "userId",
				DriveId:            "driveId",
				DriveItemVersionId: "driveItemVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/root/versions/driveItemVersionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdRootVersionID(v.Input)
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

		if actual.DriveItemVersionId != v.Expected.DriveItemVersionId {
			t.Fatalf("Expected %q but got %q for DriveItemVersionId", v.Expected.DriveItemVersionId, actual.DriveItemVersionId)
		}

	}
}

func TestParseUserIdDriveIdRootVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdRootVersionId
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
			Input: "/users/userId/drives/driveId/root",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/rOoT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/root/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/rOoT/vErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/root/versions/driveItemVersionId",
			Expected: &UserIdDriveIdRootVersionId{
				UserId:             "userId",
				DriveId:            "driveId",
				DriveItemVersionId: "driveItemVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/root/versions/driveItemVersionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/rOoT/vErSiOnS/dRiVeItEmVeRsIoNiD",
			Expected: &UserIdDriveIdRootVersionId{
				UserId:             "uSeRiD",
				DriveId:            "dRiVeId",
				DriveItemVersionId: "dRiVeItEmVeRsIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/rOoT/vErSiOnS/dRiVeItEmVeRsIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdRootVersionIDInsensitively(v.Input)
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

		if actual.DriveItemVersionId != v.Expected.DriveItemVersionId {
			t.Fatalf("Expected %q but got %q for DriveItemVersionId", v.Expected.DriveItemVersionId, actual.DriveItemVersionId)
		}

	}
}

func TestSegmentsForUserIdDriveIdRootVersionId(t *testing.T) {
	segments := UserIdDriveIdRootVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdDriveIdRootVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
