package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdListContentTypeIdColumnPositionId{}

func TestNewGroupIdDriveIdListContentTypeIdColumnPositionID(t *testing.T) {
	id := NewGroupIdDriveIdListContentTypeIdColumnPositionID("groupId", "driveId", "contentTypeId", "columnDefinitionId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.ContentTypeId != "contentTypeId" {
		t.Fatalf("Expected %q but got %q for Segment 'ContentTypeId'", id.ContentTypeId, "contentTypeId")
	}

	if id.ColumnDefinitionId != "columnDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'ColumnDefinitionId'", id.ColumnDefinitionId, "columnDefinitionId")
	}
}

func TestFormatGroupIdDriveIdListContentTypeIdColumnPositionID(t *testing.T) {
	actual := NewGroupIdDriveIdListContentTypeIdColumnPositionID("groupId", "driveId", "contentTypeId", "columnDefinitionId").ID()
	expected := "/groups/groupId/drives/driveId/list/contentTypes/contentTypeId/columnPositions/columnDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdDriveIdListContentTypeIdColumnPositionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdDriveIdListContentTypeIdColumnPositionId
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
			Input: "/groups/groupId/drives/driveId/list",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/list/contentTypes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/list/contentTypes/contentTypeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/list/contentTypes/contentTypeId/columnPositions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/drives/driveId/list/contentTypes/contentTypeId/columnPositions/columnDefinitionId",
			Expected: &GroupIdDriveIdListContentTypeIdColumnPositionId{
				GroupId:            "groupId",
				DriveId:            "driveId",
				ContentTypeId:      "contentTypeId",
				ColumnDefinitionId: "columnDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/drives/driveId/list/contentTypes/contentTypeId/columnPositions/columnDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdDriveIdListContentTypeIdColumnPositionID(v.Input)
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

		if actual.ContentTypeId != v.Expected.ContentTypeId {
			t.Fatalf("Expected %q but got %q for ContentTypeId", v.Expected.ContentTypeId, actual.ContentTypeId)
		}

		if actual.ColumnDefinitionId != v.Expected.ColumnDefinitionId {
			t.Fatalf("Expected %q but got %q for ColumnDefinitionId", v.Expected.ColumnDefinitionId, actual.ColumnDefinitionId)
		}

	}
}

func TestParseGroupIdDriveIdListContentTypeIdColumnPositionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdDriveIdListContentTypeIdColumnPositionId
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
			Input: "/groups/groupId/drives/driveId/list",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/lIsT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/list/contentTypes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/lIsT/cOnTeNtTyPeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/list/contentTypes/contentTypeId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/lIsT/cOnTeNtTyPeS/cOnTeNtTyPeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/list/contentTypes/contentTypeId/columnPositions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/lIsT/cOnTeNtTyPeS/cOnTeNtTyPeId/cOlUmNpOsItIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/drives/driveId/list/contentTypes/contentTypeId/columnPositions/columnDefinitionId",
			Expected: &GroupIdDriveIdListContentTypeIdColumnPositionId{
				GroupId:            "groupId",
				DriveId:            "driveId",
				ContentTypeId:      "contentTypeId",
				ColumnDefinitionId: "columnDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/drives/driveId/list/contentTypes/contentTypeId/columnPositions/columnDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/lIsT/cOnTeNtTyPeS/cOnTeNtTyPeId/cOlUmNpOsItIoNs/cOlUmNdEfInItIoNiD",
			Expected: &GroupIdDriveIdListContentTypeIdColumnPositionId{
				GroupId:            "gRoUpId",
				DriveId:            "dRiVeId",
				ContentTypeId:      "cOnTeNtTyPeId",
				ColumnDefinitionId: "cOlUmNdEfInItIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/lIsT/cOnTeNtTyPeS/cOnTeNtTyPeId/cOlUmNpOsItIoNs/cOlUmNdEfInItIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdDriveIdListContentTypeIdColumnPositionIDInsensitively(v.Input)
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

		if actual.ContentTypeId != v.Expected.ContentTypeId {
			t.Fatalf("Expected %q but got %q for ContentTypeId", v.Expected.ContentTypeId, actual.ContentTypeId)
		}

		if actual.ColumnDefinitionId != v.Expected.ColumnDefinitionId {
			t.Fatalf("Expected %q but got %q for ColumnDefinitionId", v.Expected.ColumnDefinitionId, actual.ColumnDefinitionId)
		}

	}
}

func TestSegmentsForGroupIdDriveIdListContentTypeIdColumnPositionId(t *testing.T) {
	segments := GroupIdDriveIdListContentTypeIdColumnPositionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdDriveIdListContentTypeIdColumnPositionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
