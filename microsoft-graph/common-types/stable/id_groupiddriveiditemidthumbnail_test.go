package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdItemIdThumbnailId{}

func TestNewGroupIdDriveIdItemIdThumbnailID(t *testing.T) {
	id := NewGroupIdDriveIdItemIdThumbnailID("groupIdValue", "driveIdValue", "driveItemIdValue", "thumbnailSetIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.DriveId != "driveIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveIdValue")
	}

	if id.DriveItemId != "driveItemIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveItemId'", id.DriveItemId, "driveItemIdValue")
	}

	if id.ThumbnailSetId != "thumbnailSetIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ThumbnailSetId'", id.ThumbnailSetId, "thumbnailSetIdValue")
	}
}

func TestFormatGroupIdDriveIdItemIdThumbnailID(t *testing.T) {
	actual := NewGroupIdDriveIdItemIdThumbnailID("groupIdValue", "driveIdValue", "driveItemIdValue", "thumbnailSetIdValue").ID()
	expected := "/groups/groupIdValue/drives/driveIdValue/items/driveItemIdValue/thumbnails/thumbnailSetIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdDriveIdItemIdThumbnailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdDriveIdItemIdThumbnailId
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
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/items",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/items/driveItemIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/items/driveItemIdValue/thumbnails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/drives/driveIdValue/items/driveItemIdValue/thumbnails/thumbnailSetIdValue",
			Expected: &GroupIdDriveIdItemIdThumbnailId{
				GroupId:        "groupIdValue",
				DriveId:        "driveIdValue",
				DriveItemId:    "driveItemIdValue",
				ThumbnailSetId: "thumbnailSetIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/drives/driveIdValue/items/driveItemIdValue/thumbnails/thumbnailSetIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdDriveIdItemIdThumbnailID(v.Input)
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

		if actual.DriveItemId != v.Expected.DriveItemId {
			t.Fatalf("Expected %q but got %q for DriveItemId", v.Expected.DriveItemId, actual.DriveItemId)
		}

		if actual.ThumbnailSetId != v.Expected.ThumbnailSetId {
			t.Fatalf("Expected %q but got %q for ThumbnailSetId", v.Expected.ThumbnailSetId, actual.ThumbnailSetId)
		}

	}
}

func TestParseGroupIdDriveIdItemIdThumbnailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdDriveIdItemIdThumbnailId
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
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/items",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE/iTeMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/items/driveItemIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE/iTeMs/dRiVeItEmIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/items/driveItemIdValue/thumbnails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE/iTeMs/dRiVeItEmIdVaLuE/tHuMbNaIlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/drives/driveIdValue/items/driveItemIdValue/thumbnails/thumbnailSetIdValue",
			Expected: &GroupIdDriveIdItemIdThumbnailId{
				GroupId:        "groupIdValue",
				DriveId:        "driveIdValue",
				DriveItemId:    "driveItemIdValue",
				ThumbnailSetId: "thumbnailSetIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/drives/driveIdValue/items/driveItemIdValue/thumbnails/thumbnailSetIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE/iTeMs/dRiVeItEmIdVaLuE/tHuMbNaIlS/tHuMbNaIlSeTiDvAlUe",
			Expected: &GroupIdDriveIdItemIdThumbnailId{
				GroupId:        "gRoUpIdVaLuE",
				DriveId:        "dRiVeIdVaLuE",
				DriveItemId:    "dRiVeItEmIdVaLuE",
				ThumbnailSetId: "tHuMbNaIlSeTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE/iTeMs/dRiVeItEmIdVaLuE/tHuMbNaIlS/tHuMbNaIlSeTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdDriveIdItemIdThumbnailIDInsensitively(v.Input)
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

		if actual.DriveItemId != v.Expected.DriveItemId {
			t.Fatalf("Expected %q but got %q for DriveItemId", v.Expected.DriveItemId, actual.DriveItemId)
		}

		if actual.ThumbnailSetId != v.Expected.ThumbnailSetId {
			t.Fatalf("Expected %q but got %q for ThumbnailSetId", v.Expected.ThumbnailSetId, actual.ThumbnailSetId)
		}

	}
}

func TestSegmentsForGroupIdDriveIdItemIdThumbnailId(t *testing.T) {
	segments := GroupIdDriveIdItemIdThumbnailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdDriveIdItemIdThumbnailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
