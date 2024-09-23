package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdListColumnId{}

func TestNewUserIdDriveIdListColumnID(t *testing.T) {
	id := NewUserIdDriveIdListColumnID("userId", "driveId", "columnDefinitionId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.ColumnDefinitionId != "columnDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'ColumnDefinitionId'", id.ColumnDefinitionId, "columnDefinitionId")
	}
}

func TestFormatUserIdDriveIdListColumnID(t *testing.T) {
	actual := NewUserIdDriveIdListColumnID("userId", "driveId", "columnDefinitionId").ID()
	expected := "/users/userId/drives/driveId/list/columns/columnDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdDriveIdListColumnID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdListColumnId
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
			Input: "/users/userId/drives/driveId/list/columns",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/list/columns/columnDefinitionId",
			Expected: &UserIdDriveIdListColumnId{
				UserId:             "userId",
				DriveId:            "driveId",
				ColumnDefinitionId: "columnDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/list/columns/columnDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdListColumnID(v.Input)
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

		if actual.ColumnDefinitionId != v.Expected.ColumnDefinitionId {
			t.Fatalf("Expected %q but got %q for ColumnDefinitionId", v.Expected.ColumnDefinitionId, actual.ColumnDefinitionId)
		}

	}
}

func TestParseUserIdDriveIdListColumnIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdListColumnId
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
			Input: "/users/userId/drives/driveId/list/columns",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/lIsT/cOlUmNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/list/columns/columnDefinitionId",
			Expected: &UserIdDriveIdListColumnId{
				UserId:             "userId",
				DriveId:            "driveId",
				ColumnDefinitionId: "columnDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/list/columns/columnDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/lIsT/cOlUmNs/cOlUmNdEfInItIoNiD",
			Expected: &UserIdDriveIdListColumnId{
				UserId:             "uSeRiD",
				DriveId:            "dRiVeId",
				ColumnDefinitionId: "cOlUmNdEfInItIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/lIsT/cOlUmNs/cOlUmNdEfInItIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdListColumnIDInsensitively(v.Input)
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

		if actual.ColumnDefinitionId != v.Expected.ColumnDefinitionId {
			t.Fatalf("Expected %q but got %q for ColumnDefinitionId", v.Expected.ColumnDefinitionId, actual.ColumnDefinitionId)
		}

	}
}

func TestSegmentsForUserIdDriveIdListColumnId(t *testing.T) {
	segments := UserIdDriveIdListColumnId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdDriveIdListColumnId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
