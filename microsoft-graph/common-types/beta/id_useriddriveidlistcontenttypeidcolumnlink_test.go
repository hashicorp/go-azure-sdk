package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdListContentTypeIdColumnLinkId{}

func TestNewUserIdDriveIdListContentTypeIdColumnLinkID(t *testing.T) {
	id := NewUserIdDriveIdListContentTypeIdColumnLinkID("userId", "driveId", "contentTypeId", "columnLinkId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.ContentTypeId != "contentTypeId" {
		t.Fatalf("Expected %q but got %q for Segment 'ContentTypeId'", id.ContentTypeId, "contentTypeId")
	}

	if id.ColumnLinkId != "columnLinkId" {
		t.Fatalf("Expected %q but got %q for Segment 'ColumnLinkId'", id.ColumnLinkId, "columnLinkId")
	}
}

func TestFormatUserIdDriveIdListContentTypeIdColumnLinkID(t *testing.T) {
	actual := NewUserIdDriveIdListContentTypeIdColumnLinkID("userId", "driveId", "contentTypeId", "columnLinkId").ID()
	expected := "/users/userId/drives/driveId/list/contentTypes/contentTypeId/columnLinks/columnLinkId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdDriveIdListContentTypeIdColumnLinkID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdListContentTypeIdColumnLinkId
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
			Input: "/users/userId/drives/driveId/list/contentTypes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/list/contentTypes/contentTypeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/list/contentTypes/contentTypeId/columnLinks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/list/contentTypes/contentTypeId/columnLinks/columnLinkId",
			Expected: &UserIdDriveIdListContentTypeIdColumnLinkId{
				UserId:        "userId",
				DriveId:       "driveId",
				ContentTypeId: "contentTypeId",
				ColumnLinkId:  "columnLinkId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/list/contentTypes/contentTypeId/columnLinks/columnLinkId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdListContentTypeIdColumnLinkID(v.Input)
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

		if actual.ContentTypeId != v.Expected.ContentTypeId {
			t.Fatalf("Expected %q but got %q for ContentTypeId", v.Expected.ContentTypeId, actual.ContentTypeId)
		}

		if actual.ColumnLinkId != v.Expected.ColumnLinkId {
			t.Fatalf("Expected %q but got %q for ColumnLinkId", v.Expected.ColumnLinkId, actual.ColumnLinkId)
		}

	}
}

func TestParseUserIdDriveIdListContentTypeIdColumnLinkIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdListContentTypeIdColumnLinkId
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
			Input: "/users/userId/drives/driveId/list/contentTypes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/lIsT/cOnTeNtTyPeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/list/contentTypes/contentTypeId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/lIsT/cOnTeNtTyPeS/cOnTeNtTyPeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/list/contentTypes/contentTypeId/columnLinks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/lIsT/cOnTeNtTyPeS/cOnTeNtTyPeId/cOlUmNlInKs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/list/contentTypes/contentTypeId/columnLinks/columnLinkId",
			Expected: &UserIdDriveIdListContentTypeIdColumnLinkId{
				UserId:        "userId",
				DriveId:       "driveId",
				ContentTypeId: "contentTypeId",
				ColumnLinkId:  "columnLinkId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/list/contentTypes/contentTypeId/columnLinks/columnLinkId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/lIsT/cOnTeNtTyPeS/cOnTeNtTyPeId/cOlUmNlInKs/cOlUmNlInKiD",
			Expected: &UserIdDriveIdListContentTypeIdColumnLinkId{
				UserId:        "uSeRiD",
				DriveId:       "dRiVeId",
				ContentTypeId: "cOnTeNtTyPeId",
				ColumnLinkId:  "cOlUmNlInKiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/lIsT/cOnTeNtTyPeS/cOnTeNtTyPeId/cOlUmNlInKs/cOlUmNlInKiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdListContentTypeIdColumnLinkIDInsensitively(v.Input)
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

		if actual.ContentTypeId != v.Expected.ContentTypeId {
			t.Fatalf("Expected %q but got %q for ContentTypeId", v.Expected.ContentTypeId, actual.ContentTypeId)
		}

		if actual.ColumnLinkId != v.Expected.ColumnLinkId {
			t.Fatalf("Expected %q but got %q for ColumnLinkId", v.Expected.ColumnLinkId, actual.ColumnLinkId)
		}

	}
}

func TestSegmentsForUserIdDriveIdListContentTypeIdColumnLinkId(t *testing.T) {
	segments := UserIdDriveIdListContentTypeIdColumnLinkId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdDriveIdListContentTypeIdColumnLinkId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
