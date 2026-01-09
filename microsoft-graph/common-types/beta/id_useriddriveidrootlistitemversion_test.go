package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdRootListItemVersionId{}

func TestNewUserIdDriveIdRootListItemVersionID(t *testing.T) {
	id := NewUserIdDriveIdRootListItemVersionID("userId", "driveId", "listItemVersionId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.ListItemVersionId != "listItemVersionId" {
		t.Fatalf("Expected %q but got %q for Segment 'ListItemVersionId'", id.ListItemVersionId, "listItemVersionId")
	}
}

func TestFormatUserIdDriveIdRootListItemVersionID(t *testing.T) {
	actual := NewUserIdDriveIdRootListItemVersionID("userId", "driveId", "listItemVersionId").ID()
	expected := "/users/userId/drives/driveId/root/listItem/versions/listItemVersionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdDriveIdRootListItemVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdRootListItemVersionId
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
			Input: "/users/userId/drives/driveId/root/listItem",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/root/listItem/versions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/root/listItem/versions/listItemVersionId",
			Expected: &UserIdDriveIdRootListItemVersionId{
				UserId:            "userId",
				DriveId:           "driveId",
				ListItemVersionId: "listItemVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/root/listItem/versions/listItemVersionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdRootListItemVersionID(v.Input)
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

		if actual.ListItemVersionId != v.Expected.ListItemVersionId {
			t.Fatalf("Expected %q but got %q for ListItemVersionId", v.Expected.ListItemVersionId, actual.ListItemVersionId)
		}

	}
}

func TestParseUserIdDriveIdRootListItemVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdRootListItemVersionId
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
			Input: "/users/userId/drives/driveId/root/listItem",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/rOoT/lIsTiTeM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/root/listItem/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/rOoT/lIsTiTeM/vErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/root/listItem/versions/listItemVersionId",
			Expected: &UserIdDriveIdRootListItemVersionId{
				UserId:            "userId",
				DriveId:           "driveId",
				ListItemVersionId: "listItemVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/root/listItem/versions/listItemVersionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/rOoT/lIsTiTeM/vErSiOnS/lIsTiTeMvErSiOnId",
			Expected: &UserIdDriveIdRootListItemVersionId{
				UserId:            "uSeRiD",
				DriveId:           "dRiVeId",
				ListItemVersionId: "lIsTiTeMvErSiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/rOoT/lIsTiTeM/vErSiOnS/lIsTiTeMvErSiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdRootListItemVersionIDInsensitively(v.Input)
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

		if actual.ListItemVersionId != v.Expected.ListItemVersionId {
			t.Fatalf("Expected %q but got %q for ListItemVersionId", v.Expected.ListItemVersionId, actual.ListItemVersionId)
		}

	}
}

func TestSegmentsForUserIdDriveIdRootListItemVersionId(t *testing.T) {
	segments := UserIdDriveIdRootListItemVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdDriveIdRootListItemVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
