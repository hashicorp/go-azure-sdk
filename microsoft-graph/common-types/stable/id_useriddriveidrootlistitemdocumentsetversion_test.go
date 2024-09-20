package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdRootListItemDocumentSetVersionId{}

func TestNewUserIdDriveIdRootListItemDocumentSetVersionID(t *testing.T) {
	id := NewUserIdDriveIdRootListItemDocumentSetVersionID("userId", "driveId", "documentSetVersionId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.DocumentSetVersionId != "documentSetVersionId" {
		t.Fatalf("Expected %q but got %q for Segment 'DocumentSetVersionId'", id.DocumentSetVersionId, "documentSetVersionId")
	}
}

func TestFormatUserIdDriveIdRootListItemDocumentSetVersionID(t *testing.T) {
	actual := NewUserIdDriveIdRootListItemDocumentSetVersionID("userId", "driveId", "documentSetVersionId").ID()
	expected := "/users/userId/drives/driveId/root/listItem/documentSetVersions/documentSetVersionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdDriveIdRootListItemDocumentSetVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdRootListItemDocumentSetVersionId
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
			Input: "/users/userId/drives/driveId/root/listItem/documentSetVersions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/root/listItem/documentSetVersions/documentSetVersionId",
			Expected: &UserIdDriveIdRootListItemDocumentSetVersionId{
				UserId:               "userId",
				DriveId:              "driveId",
				DocumentSetVersionId: "documentSetVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/root/listItem/documentSetVersions/documentSetVersionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdRootListItemDocumentSetVersionID(v.Input)
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

		if actual.DocumentSetVersionId != v.Expected.DocumentSetVersionId {
			t.Fatalf("Expected %q but got %q for DocumentSetVersionId", v.Expected.DocumentSetVersionId, actual.DocumentSetVersionId)
		}

	}
}

func TestParseUserIdDriveIdRootListItemDocumentSetVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdRootListItemDocumentSetVersionId
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
			Input: "/users/userId/drives/driveId/root/listItem/documentSetVersions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/rOoT/lIsTiTeM/dOcUmEnTsEtVeRsIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/root/listItem/documentSetVersions/documentSetVersionId",
			Expected: &UserIdDriveIdRootListItemDocumentSetVersionId{
				UserId:               "userId",
				DriveId:              "driveId",
				DocumentSetVersionId: "documentSetVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/root/listItem/documentSetVersions/documentSetVersionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/rOoT/lIsTiTeM/dOcUmEnTsEtVeRsIoNs/dOcUmEnTsEtVeRsIoNiD",
			Expected: &UserIdDriveIdRootListItemDocumentSetVersionId{
				UserId:               "uSeRiD",
				DriveId:              "dRiVeId",
				DocumentSetVersionId: "dOcUmEnTsEtVeRsIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/rOoT/lIsTiTeM/dOcUmEnTsEtVeRsIoNs/dOcUmEnTsEtVeRsIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdRootListItemDocumentSetVersionIDInsensitively(v.Input)
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

		if actual.DocumentSetVersionId != v.Expected.DocumentSetVersionId {
			t.Fatalf("Expected %q but got %q for DocumentSetVersionId", v.Expected.DocumentSetVersionId, actual.DocumentSetVersionId)
		}

	}
}

func TestSegmentsForUserIdDriveIdRootListItemDocumentSetVersionId(t *testing.T) {
	segments := UserIdDriveIdRootListItemDocumentSetVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdDriveIdRootListItemDocumentSetVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
