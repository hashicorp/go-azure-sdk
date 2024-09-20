package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdItemIdListItemDocumentSetVersionId{}

func TestNewUserIdDriveIdItemIdListItemDocumentSetVersionID(t *testing.T) {
	id := NewUserIdDriveIdItemIdListItemDocumentSetVersionID("userId", "driveId", "driveItemId", "documentSetVersionId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.DriveItemId != "driveItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveItemId'", id.DriveItemId, "driveItemId")
	}

	if id.DocumentSetVersionId != "documentSetVersionId" {
		t.Fatalf("Expected %q but got %q for Segment 'DocumentSetVersionId'", id.DocumentSetVersionId, "documentSetVersionId")
	}
}

func TestFormatUserIdDriveIdItemIdListItemDocumentSetVersionID(t *testing.T) {
	actual := NewUserIdDriveIdItemIdListItemDocumentSetVersionID("userId", "driveId", "driveItemId", "documentSetVersionId").ID()
	expected := "/users/userId/drives/driveId/items/driveItemId/listItem/documentSetVersions/documentSetVersionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdDriveIdItemIdListItemDocumentSetVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdItemIdListItemDocumentSetVersionId
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
			Input: "/users/userId/drives/driveId/items",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/items/driveItemId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem/documentSetVersions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem/documentSetVersions/documentSetVersionId",
			Expected: &UserIdDriveIdItemIdListItemDocumentSetVersionId{
				UserId:               "userId",
				DriveId:              "driveId",
				DriveItemId:          "driveItemId",
				DocumentSetVersionId: "documentSetVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem/documentSetVersions/documentSetVersionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdItemIdListItemDocumentSetVersionID(v.Input)
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

		if actual.DriveItemId != v.Expected.DriveItemId {
			t.Fatalf("Expected %q but got %q for DriveItemId", v.Expected.DriveItemId, actual.DriveItemId)
		}

		if actual.DocumentSetVersionId != v.Expected.DocumentSetVersionId {
			t.Fatalf("Expected %q but got %q for DocumentSetVersionId", v.Expected.DocumentSetVersionId, actual.DocumentSetVersionId)
		}

	}
}

func TestParseUserIdDriveIdItemIdListItemDocumentSetVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdItemIdListItemDocumentSetVersionId
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
			Input: "/users/userId/drives/driveId/items",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/iTeMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/items/driveItemId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/lIsTiTeM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem/documentSetVersions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/lIsTiTeM/dOcUmEnTsEtVeRsIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem/documentSetVersions/documentSetVersionId",
			Expected: &UserIdDriveIdItemIdListItemDocumentSetVersionId{
				UserId:               "userId",
				DriveId:              "driveId",
				DriveItemId:          "driveItemId",
				DocumentSetVersionId: "documentSetVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/items/driveItemId/listItem/documentSetVersions/documentSetVersionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/lIsTiTeM/dOcUmEnTsEtVeRsIoNs/dOcUmEnTsEtVeRsIoNiD",
			Expected: &UserIdDriveIdItemIdListItemDocumentSetVersionId{
				UserId:               "uSeRiD",
				DriveId:              "dRiVeId",
				DriveItemId:          "dRiVeItEmId",
				DocumentSetVersionId: "dOcUmEnTsEtVeRsIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/lIsTiTeM/dOcUmEnTsEtVeRsIoNs/dOcUmEnTsEtVeRsIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdItemIdListItemDocumentSetVersionIDInsensitively(v.Input)
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

		if actual.DriveItemId != v.Expected.DriveItemId {
			t.Fatalf("Expected %q but got %q for DriveItemId", v.Expected.DriveItemId, actual.DriveItemId)
		}

		if actual.DocumentSetVersionId != v.Expected.DocumentSetVersionId {
			t.Fatalf("Expected %q but got %q for DocumentSetVersionId", v.Expected.DocumentSetVersionId, actual.DocumentSetVersionId)
		}

	}
}

func TestSegmentsForUserIdDriveIdItemIdListItemDocumentSetVersionId(t *testing.T) {
	segments := UserIdDriveIdItemIdListItemDocumentSetVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdDriveIdItemIdListItemDocumentSetVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
