package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdRootThumbnailId{}

func TestNewUserIdDriveIdRootThumbnailID(t *testing.T) {
	id := NewUserIdDriveIdRootThumbnailID("userId", "driveId", "thumbnailSetId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.ThumbnailSetId != "thumbnailSetId" {
		t.Fatalf("Expected %q but got %q for Segment 'ThumbnailSetId'", id.ThumbnailSetId, "thumbnailSetId")
	}
}

func TestFormatUserIdDriveIdRootThumbnailID(t *testing.T) {
	actual := NewUserIdDriveIdRootThumbnailID("userId", "driveId", "thumbnailSetId").ID()
	expected := "/users/userId/drives/driveId/root/thumbnails/thumbnailSetId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdDriveIdRootThumbnailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdRootThumbnailId
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
			Input: "/users/userId/drives/driveId/root/thumbnails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/root/thumbnails/thumbnailSetId",
			Expected: &UserIdDriveIdRootThumbnailId{
				UserId:         "userId",
				DriveId:        "driveId",
				ThumbnailSetId: "thumbnailSetId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/root/thumbnails/thumbnailSetId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdRootThumbnailID(v.Input)
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

		if actual.ThumbnailSetId != v.Expected.ThumbnailSetId {
			t.Fatalf("Expected %q but got %q for ThumbnailSetId", v.Expected.ThumbnailSetId, actual.ThumbnailSetId)
		}

	}
}

func TestParseUserIdDriveIdRootThumbnailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdRootThumbnailId
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
			Input: "/users/userId/drives/driveId/root/thumbnails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/rOoT/tHuMbNaIlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/root/thumbnails/thumbnailSetId",
			Expected: &UserIdDriveIdRootThumbnailId{
				UserId:         "userId",
				DriveId:        "driveId",
				ThumbnailSetId: "thumbnailSetId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/root/thumbnails/thumbnailSetId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/rOoT/tHuMbNaIlS/tHuMbNaIlSeTiD",
			Expected: &UserIdDriveIdRootThumbnailId{
				UserId:         "uSeRiD",
				DriveId:        "dRiVeId",
				ThumbnailSetId: "tHuMbNaIlSeTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/rOoT/tHuMbNaIlS/tHuMbNaIlSeTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdRootThumbnailIDInsensitively(v.Input)
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

		if actual.ThumbnailSetId != v.Expected.ThumbnailSetId {
			t.Fatalf("Expected %q but got %q for ThumbnailSetId", v.Expected.ThumbnailSetId, actual.ThumbnailSetId)
		}

	}
}

func TestSegmentsForUserIdDriveIdRootThumbnailId(t *testing.T) {
	segments := UserIdDriveIdRootThumbnailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdDriveIdRootThumbnailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
