package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdRootListItemDocumentSetVersionId{}

func TestNewGroupIdDriveIdRootListItemDocumentSetVersionID(t *testing.T) {
	id := NewGroupIdDriveIdRootListItemDocumentSetVersionID("groupId", "driveId", "documentSetVersionId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.DocumentSetVersionId != "documentSetVersionId" {
		t.Fatalf("Expected %q but got %q for Segment 'DocumentSetVersionId'", id.DocumentSetVersionId, "documentSetVersionId")
	}
}

func TestFormatGroupIdDriveIdRootListItemDocumentSetVersionID(t *testing.T) {
	actual := NewGroupIdDriveIdRootListItemDocumentSetVersionID("groupId", "driveId", "documentSetVersionId").ID()
	expected := "/groups/groupId/drives/driveId/root/listItem/documentSetVersions/documentSetVersionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdDriveIdRootListItemDocumentSetVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdDriveIdRootListItemDocumentSetVersionId
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
			Input: "/groups/groupId/drives/driveId/root/listItem",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/root/listItem/documentSetVersions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/drives/driveId/root/listItem/documentSetVersions/documentSetVersionId",
			Expected: &GroupIdDriveIdRootListItemDocumentSetVersionId{
				GroupId:              "groupId",
				DriveId:              "driveId",
				DocumentSetVersionId: "documentSetVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/drives/driveId/root/listItem/documentSetVersions/documentSetVersionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdDriveIdRootListItemDocumentSetVersionID(v.Input)
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

		if actual.DocumentSetVersionId != v.Expected.DocumentSetVersionId {
			t.Fatalf("Expected %q but got %q for DocumentSetVersionId", v.Expected.DocumentSetVersionId, actual.DocumentSetVersionId)
		}

	}
}

func TestParseGroupIdDriveIdRootListItemDocumentSetVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdDriveIdRootListItemDocumentSetVersionId
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
			Input: "/groups/groupId/drives/driveId/root/listItem",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/rOoT/lIsTiTeM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/drives/driveId/root/listItem/documentSetVersions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/rOoT/lIsTiTeM/dOcUmEnTsEtVeRsIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/drives/driveId/root/listItem/documentSetVersions/documentSetVersionId",
			Expected: &GroupIdDriveIdRootListItemDocumentSetVersionId{
				GroupId:              "groupId",
				DriveId:              "driveId",
				DocumentSetVersionId: "documentSetVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/drives/driveId/root/listItem/documentSetVersions/documentSetVersionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/rOoT/lIsTiTeM/dOcUmEnTsEtVeRsIoNs/dOcUmEnTsEtVeRsIoNiD",
			Expected: &GroupIdDriveIdRootListItemDocumentSetVersionId{
				GroupId:              "gRoUpId",
				DriveId:              "dRiVeId",
				DocumentSetVersionId: "dOcUmEnTsEtVeRsIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/dRiVeS/dRiVeId/rOoT/lIsTiTeM/dOcUmEnTsEtVeRsIoNs/dOcUmEnTsEtVeRsIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdDriveIdRootListItemDocumentSetVersionIDInsensitively(v.Input)
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

		if actual.DocumentSetVersionId != v.Expected.DocumentSetVersionId {
			t.Fatalf("Expected %q but got %q for DocumentSetVersionId", v.Expected.DocumentSetVersionId, actual.DocumentSetVersionId)
		}

	}
}

func TestSegmentsForGroupIdDriveIdRootListItemDocumentSetVersionId(t *testing.T) {
	segments := GroupIdDriveIdRootListItemDocumentSetVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdDriveIdRootListItemDocumentSetVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
