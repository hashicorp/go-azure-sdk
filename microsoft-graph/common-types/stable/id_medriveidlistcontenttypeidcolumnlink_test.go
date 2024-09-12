package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListContentTypeIdColumnLinkId{}

func TestNewMeDriveIdListContentTypeIdColumnLinkID(t *testing.T) {
	id := NewMeDriveIdListContentTypeIdColumnLinkID("driveIdValue", "contentTypeIdValue", "columnLinkIdValue")

	if id.DriveId != "driveIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveIdValue")
	}

	if id.ContentTypeId != "contentTypeIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ContentTypeId'", id.ContentTypeId, "contentTypeIdValue")
	}

	if id.ColumnLinkId != "columnLinkIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ColumnLinkId'", id.ColumnLinkId, "columnLinkIdValue")
	}
}

func TestFormatMeDriveIdListContentTypeIdColumnLinkID(t *testing.T) {
	actual := NewMeDriveIdListContentTypeIdColumnLinkID("driveIdValue", "contentTypeIdValue", "columnLinkIdValue").ID()
	expected := "/me/drives/driveIdValue/list/contentTypes/contentTypeIdValue/columnLinks/columnLinkIdValue"
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
			Input: "/me/drives/driveIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/list",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/list/contentTypes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/list/contentTypes/contentTypeIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/list/contentTypes/contentTypeIdValue/columnLinks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveIdValue/list/contentTypes/contentTypeIdValue/columnLinks/columnLinkIdValue",
			Expected: &MeDriveIdListContentTypeIdColumnLinkId{
				DriveId:       "driveIdValue",
				ContentTypeId: "contentTypeIdValue",
				ColumnLinkId:  "columnLinkIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveIdValue/list/contentTypes/contentTypeIdValue/columnLinks/columnLinkIdValue/extra",
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
			Input: "/me/drives/driveIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/list",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/lIsT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/list/contentTypes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/lIsT/cOnTeNtTyPeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/list/contentTypes/contentTypeIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/lIsT/cOnTeNtTyPeS/cOnTeNtTyPeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/list/contentTypes/contentTypeIdValue/columnLinks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/lIsT/cOnTeNtTyPeS/cOnTeNtTyPeIdVaLuE/cOlUmNlInKs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveIdValue/list/contentTypes/contentTypeIdValue/columnLinks/columnLinkIdValue",
			Expected: &MeDriveIdListContentTypeIdColumnLinkId{
				DriveId:       "driveIdValue",
				ContentTypeId: "contentTypeIdValue",
				ColumnLinkId:  "columnLinkIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveIdValue/list/contentTypes/contentTypeIdValue/columnLinks/columnLinkIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/lIsT/cOnTeNtTyPeS/cOnTeNtTyPeIdVaLuE/cOlUmNlInKs/cOlUmNlInKiDvAlUe",
			Expected: &MeDriveIdListContentTypeIdColumnLinkId{
				DriveId:       "dRiVeIdVaLuE",
				ContentTypeId: "cOnTeNtTyPeIdVaLuE",
				ColumnLinkId:  "cOlUmNlInKiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/lIsT/cOnTeNtTyPeS/cOnTeNtTyPeIdVaLuE/cOlUmNlInKs/cOlUmNlInKiDvAlUe/extra",
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
