package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdRootThumbnailId{}

func TestNewGroupIdDriveIdRootThumbnailID(t *testing.T) {
	id := NewGroupIdDriveIdRootThumbnailID("groupId", "driveId", "thumbnailSetId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.ThumbnailSetId != "thumbnailSetId" {
		t.Fatalf("Expected %q but got %q for Segment 'ThumbnailSetId'", id.ThumbnailSetId, "thumbnailSetId")
	}
}

func TestFormatGroupIdDriveIdRootThumbnailID(t *testing.T) {
	actual := NewGroupIdDriveIdRootThumbnailID("groupId", "driveId", "thumbnailSetId").ID()
	expected := "/groups/groupId/drives/driveId/root/thumbnails/thumbnailSetId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdDriveIdRootThumbnailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdDriveIdRootThumbnailId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/root",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/root/thumbnails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/drives/driveId/root/thumbnails/thumbnailSetId",
			Expected: &GroupIdDriveIdRootThumbnailId{
				GroupId:        "groupId",
				DriveId:        "driveId",
				ThumbnailSetId: "thumbnailSetId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/drives/driveId/root/thumbnails/thumbnailSetId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdDriveIdRootThumbnailID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.DriveId != v.Expected.DriveId {
			t.Fatalf("Expected %q but got %q for DriveId", v.Expected.DriveId, actual.DriveId)
		}

		if actual.ThumbnailSetId != v.Expected.ThumbnailSetId {
			t.Fatalf("Expected %q but got %q for ThumbnailSetId", v.Expected.ThumbnailSetId, actual.ThumbnailSetId)
		}

	}
}

func TestParseGroupIdDriveIdRootThumbnailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdDriveIdRootThumbnailId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/root",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/rOoT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/root/thumbnails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/rOoT/tHuMbNaIlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/drives/driveId/root/thumbnails/thumbnailSetId",
			Expected: &GroupIdDriveIdRootThumbnailId{
				GroupId:        "groupId",
				DriveId:        "driveId",
				ThumbnailSetId: "thumbnailSetId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/drives/driveId/root/thumbnails/thumbnailSetId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/rOoT/tHuMbNaIlS/tHuMbNaIlSeTiD",
			Expected: &GroupIdDriveIdRootThumbnailId{
				GroupId:        "gRoUpId",
				DriveId:        "dRiVeId",
				ThumbnailSetId: "tHuMbNaIlSeTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/rOoT/tHuMbNaIlS/tHuMbNaIlSeTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdDriveIdRootThumbnailIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.DriveId != v.Expected.DriveId {
			t.Fatalf("Expected %q but got %q for DriveId", v.Expected.DriveId, actual.DriveId)
		}

		if actual.ThumbnailSetId != v.Expected.ThumbnailSetId {
			t.Fatalf("Expected %q but got %q for ThumbnailSetId", v.Expected.ThumbnailSetId, actual.ThumbnailSetId)
		}

	}
}

func TestSegmentsForGroupIdDriveIdRootThumbnailId(t *testing.T) {
	segments := GroupIdDriveIdRootThumbnailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdDriveIdRootThumbnailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
