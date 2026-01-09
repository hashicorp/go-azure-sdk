package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListContentTypeIdColumnLinkId{}

func TestNewMeDriveIdListContentTypeIdColumnLinkID(t *testing.T) {
	id := NewMeDriveIdListContentTypeIdColumnLinkID("driveId", "contentTypeId", "columnLinkId")

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

func TestFormatMeDriveIdListContentTypeIdColumnLinkID(t *testing.T) {
	actual := NewMeDriveIdListContentTypeIdColumnLinkID("driveId", "contentTypeId", "columnLinkId").ID()
	expected := "/me/drives/driveId/list/contentTypes/contentTypeId/columnLinks/columnLinkId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDriveIdListContentTypeIdColumnLinkID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdListContentTypeIdColumnLinkId
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
			Input: "/me/drives/driveId/list",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/list/contentTypes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/list/contentTypes/contentTypeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/list/contentTypes/contentTypeId/columnLinks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/list/contentTypes/contentTypeId/columnLinks/columnLinkId",
			Expected: &MeDriveIdListContentTypeIdColumnLinkId{
				DriveId:       "driveId",
				ContentTypeId: "contentTypeId",
				ColumnLinkId:  "columnLinkId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/list/contentTypes/contentTypeId/columnLinks/columnLinkId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdListContentTypeIdColumnLinkID(v.Input)
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

		if actual.ContentTypeId != v.Expected.ContentTypeId {
			t.Fatalf("Expected %q but got %q for ContentTypeId", v.Expected.ContentTypeId, actual.ContentTypeId)
		}

		if actual.ColumnLinkId != v.Expected.ColumnLinkId {
			t.Fatalf("Expected %q but got %q for ColumnLinkId", v.Expected.ColumnLinkId, actual.ColumnLinkId)
		}

	}
}

func TestParseMeDriveIdListContentTypeIdColumnLinkIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdListContentTypeIdColumnLinkId
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
			Input: "/me/drives/driveId/list",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/list/contentTypes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT/cOnTeNtTyPeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/list/contentTypes/contentTypeId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT/cOnTeNtTyPeS/cOnTeNtTyPeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/list/contentTypes/contentTypeId/columnLinks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT/cOnTeNtTyPeS/cOnTeNtTyPeId/cOlUmNlInKs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/list/contentTypes/contentTypeId/columnLinks/columnLinkId",
			Expected: &MeDriveIdListContentTypeIdColumnLinkId{
				DriveId:       "driveId",
				ContentTypeId: "contentTypeId",
				ColumnLinkId:  "columnLinkId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/list/contentTypes/contentTypeId/columnLinks/columnLinkId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT/cOnTeNtTyPeS/cOnTeNtTyPeId/cOlUmNlInKs/cOlUmNlInKiD",
			Expected: &MeDriveIdListContentTypeIdColumnLinkId{
				DriveId:       "dRiVeId",
				ContentTypeId: "cOnTeNtTyPeId",
				ColumnLinkId:  "cOlUmNlInKiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT/cOnTeNtTyPeS/cOnTeNtTyPeId/cOlUmNlInKs/cOlUmNlInKiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdListContentTypeIdColumnLinkIDInsensitively(v.Input)
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

		if actual.ContentTypeId != v.Expected.ContentTypeId {
			t.Fatalf("Expected %q but got %q for ContentTypeId", v.Expected.ContentTypeId, actual.ContentTypeId)
		}

		if actual.ColumnLinkId != v.Expected.ColumnLinkId {
			t.Fatalf("Expected %q but got %q for ColumnLinkId", v.Expected.ColumnLinkId, actual.ColumnLinkId)
		}

	}
}

func TestSegmentsForMeDriveIdListContentTypeIdColumnLinkId(t *testing.T) {
	segments := MeDriveIdListContentTypeIdColumnLinkId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDriveIdListContentTypeIdColumnLinkId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
