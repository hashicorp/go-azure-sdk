package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdListItemIdDocumentSetVersionId{}

func TestNewGroupIdDriveIdListItemIdDocumentSetVersionID(t *testing.T) {
	id := NewGroupIdDriveIdListItemIdDocumentSetVersionID("groupId", "driveId", "listItemId", "documentSetVersionId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.ListItemId != "listItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'ListItemId'", id.ListItemId, "listItemId")
	}

	if id.DocumentSetVersionId != "documentSetVersionId" {
		t.Fatalf("Expected %q but got %q for Segment 'DocumentSetVersionId'", id.DocumentSetVersionId, "documentSetVersionId")
	}
}

func TestFormatGroupIdDriveIdListItemIdDocumentSetVersionID(t *testing.T) {
	actual := NewGroupIdDriveIdListItemIdDocumentSetVersionID("groupId", "driveId", "listItemId", "documentSetVersionId").ID()
	expected := "/groups/groupId/drives/driveId/list/items/listItemId/documentSetVersions/documentSetVersionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdDriveIdListItemIdDocumentSetVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdDriveIdListItemIdDocumentSetVersionId
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
			Input: "/groups/groupId/drives/driveId/list/items",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/list/items/listItemId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/list/items/listItemId/documentSetVersions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/drives/driveId/list/items/listItemId/documentSetVersions/documentSetVersionId",
			Expected: &GroupIdDriveIdListItemIdDocumentSetVersionId{
				GroupId:              "groupId",
				DriveId:              "driveId",
				ListItemId:           "listItemId",
				DocumentSetVersionId: "documentSetVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/drives/driveId/list/items/listItemId/documentSetVersions/documentSetVersionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdDriveIdListItemIdDocumentSetVersionID(v.Input)
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

		if actual.ListItemId != v.Expected.ListItemId {
			t.Fatalf("Expected %q but got %q for ListItemId", v.Expected.ListItemId, actual.ListItemId)
		}

		if actual.DocumentSetVersionId != v.Expected.DocumentSetVersionId {
			t.Fatalf("Expected %q but got %q for DocumentSetVersionId", v.Expected.DocumentSetVersionId, actual.DocumentSetVersionId)
		}

	}
}

func TestParseGroupIdDriveIdListItemIdDocumentSetVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdDriveIdListItemIdDocumentSetVersionId
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
			Input: "/groups/groupId/drives/driveId/list/items",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/lIsT/iTeMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/list/items/listItemId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/lIsT/iTeMs/lIsTiTeMiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/list/items/listItemId/documentSetVersions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/lIsT/iTeMs/lIsTiTeMiD/dOcUmEnTsEtVeRsIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/drives/driveId/list/items/listItemId/documentSetVersions/documentSetVersionId",
			Expected: &GroupIdDriveIdListItemIdDocumentSetVersionId{
				GroupId:              "groupId",
				DriveId:              "driveId",
				ListItemId:           "listItemId",
				DocumentSetVersionId: "documentSetVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/drives/driveId/list/items/listItemId/documentSetVersions/documentSetVersionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/lIsT/iTeMs/lIsTiTeMiD/dOcUmEnTsEtVeRsIoNs/dOcUmEnTsEtVeRsIoNiD",
			Expected: &GroupIdDriveIdListItemIdDocumentSetVersionId{
				GroupId:              "gRoUpId",
				DriveId:              "dRiVeId",
				ListItemId:           "lIsTiTeMiD",
				DocumentSetVersionId: "dOcUmEnTsEtVeRsIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/lIsT/iTeMs/lIsTiTeMiD/dOcUmEnTsEtVeRsIoNs/dOcUmEnTsEtVeRsIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdDriveIdListItemIdDocumentSetVersionIDInsensitively(v.Input)
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

		if actual.ListItemId != v.Expected.ListItemId {
			t.Fatalf("Expected %q but got %q for ListItemId", v.Expected.ListItemId, actual.ListItemId)
		}

		if actual.DocumentSetVersionId != v.Expected.DocumentSetVersionId {
			t.Fatalf("Expected %q but got %q for DocumentSetVersionId", v.Expected.DocumentSetVersionId, actual.DocumentSetVersionId)
		}

	}
}

func TestSegmentsForGroupIdDriveIdListItemIdDocumentSetVersionId(t *testing.T) {
	segments := GroupIdDriveIdListItemIdDocumentSetVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdDriveIdListItemIdDocumentSetVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
