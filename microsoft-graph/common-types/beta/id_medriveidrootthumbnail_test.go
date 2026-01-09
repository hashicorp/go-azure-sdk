package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdRootThumbnailId{}

func TestNewMeDriveIdRootThumbnailID(t *testing.T) {
	id := NewMeDriveIdRootThumbnailID("driveId", "thumbnailSetId")

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.ThumbnailSetId != "thumbnailSetId" {
		t.Fatalf("Expected %q but got %q for Segment 'ThumbnailSetId'", id.ThumbnailSetId, "thumbnailSetId")
	}
}

func TestFormatMeDriveIdRootThumbnailID(t *testing.T) {
	actual := NewMeDriveIdRootThumbnailID("driveId", "thumbnailSetId").ID()
	expected := "/me/drives/driveId/root/thumbnails/thumbnailSetId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDriveIdRootThumbnailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdRootThumbnailId
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
			Input: "/me/drives/driveId/root",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/root/thumbnails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/root/thumbnails/thumbnailSetId",
			Expected: &MeDriveIdRootThumbnailId{
				DriveId:        "driveId",
				ThumbnailSetId: "thumbnailSetId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/root/thumbnails/thumbnailSetId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdRootThumbnailID(v.Input)
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

		if actual.ThumbnailSetId != v.Expected.ThumbnailSetId {
			t.Fatalf("Expected %q but got %q for ThumbnailSetId", v.Expected.ThumbnailSetId, actual.ThumbnailSetId)
		}

	}
}

func TestParseMeDriveIdRootThumbnailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdRootThumbnailId
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
			Input: "/me/drives/driveId/root",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/rOoT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/root/thumbnails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/rOoT/tHuMbNaIlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/root/thumbnails/thumbnailSetId",
			Expected: &MeDriveIdRootThumbnailId{
				DriveId:        "driveId",
				ThumbnailSetId: "thumbnailSetId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/root/thumbnails/thumbnailSetId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/rOoT/tHuMbNaIlS/tHuMbNaIlSeTiD",
			Expected: &MeDriveIdRootThumbnailId{
				DriveId:        "dRiVeId",
				ThumbnailSetId: "tHuMbNaIlSeTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/rOoT/tHuMbNaIlS/tHuMbNaIlSeTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdRootThumbnailIDInsensitively(v.Input)
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

		if actual.ThumbnailSetId != v.Expected.ThumbnailSetId {
			t.Fatalf("Expected %q but got %q for ThumbnailSetId", v.Expected.ThumbnailSetId, actual.ThumbnailSetId)
		}

	}
}

func TestSegmentsForMeDriveIdRootThumbnailId(t *testing.T) {
	segments := MeDriveIdRootThumbnailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDriveIdRootThumbnailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
