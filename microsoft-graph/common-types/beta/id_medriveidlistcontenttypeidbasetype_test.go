package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListContentTypeIdBaseTypeId{}

func TestNewMeDriveIdListContentTypeIdBaseTypeID(t *testing.T) {
	id := NewMeDriveIdListContentTypeIdBaseTypeID("driveId", "contentTypeId", "contentTypeId1")

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.ContentTypeId != "contentTypeId" {
		t.Fatalf("Expected %q but got %q for Segment 'ContentTypeId'", id.ContentTypeId, "contentTypeId")
	}

	if id.ContentTypeId1 != "contentTypeId1" {
		t.Fatalf("Expected %q but got %q for Segment 'ContentTypeId1'", id.ContentTypeId1, "contentTypeId1")
	}
}

func TestFormatMeDriveIdListContentTypeIdBaseTypeID(t *testing.T) {
	actual := NewMeDriveIdListContentTypeIdBaseTypeID("driveId", "contentTypeId", "contentTypeId1").ID()
	expected := "/me/drives/driveId/list/contentTypes/contentTypeId/baseTypes/contentTypeId1"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDriveIdListContentTypeIdBaseTypeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdListContentTypeIdBaseTypeId
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
			Input: "/me/drives/driveId/list/contentTypes/contentTypeId/baseTypes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/list/contentTypes/contentTypeId/baseTypes/contentTypeId1",
			Expected: &MeDriveIdListContentTypeIdBaseTypeId{
				DriveId:        "driveId",
				ContentTypeId:  "contentTypeId",
				ContentTypeId1: "contentTypeId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/list/contentTypes/contentTypeId/baseTypes/contentTypeId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdListContentTypeIdBaseTypeID(v.Input)
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

		if actual.ContentTypeId1 != v.Expected.ContentTypeId1 {
			t.Fatalf("Expected %q but got %q for ContentTypeId1", v.Expected.ContentTypeId1, actual.ContentTypeId1)
		}

	}
}

func TestParseMeDriveIdListContentTypeIdBaseTypeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdListContentTypeIdBaseTypeId
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
			Input: "/me/drives/driveId/list/contentTypes/contentTypeId/baseTypes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT/cOnTeNtTyPeS/cOnTeNtTyPeId/bAsEtYpEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/list/contentTypes/contentTypeId/baseTypes/contentTypeId1",
			Expected: &MeDriveIdListContentTypeIdBaseTypeId{
				DriveId:        "driveId",
				ContentTypeId:  "contentTypeId",
				ContentTypeId1: "contentTypeId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/list/contentTypes/contentTypeId/baseTypes/contentTypeId1/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT/cOnTeNtTyPeS/cOnTeNtTyPeId/bAsEtYpEs/cOnTeNtTyPeId1",
			Expected: &MeDriveIdListContentTypeIdBaseTypeId{
				DriveId:        "dRiVeId",
				ContentTypeId:  "cOnTeNtTyPeId",
				ContentTypeId1: "cOnTeNtTyPeId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/lIsT/cOnTeNtTyPeS/cOnTeNtTyPeId/bAsEtYpEs/cOnTeNtTyPeId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdListContentTypeIdBaseTypeIDInsensitively(v.Input)
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

		if actual.ContentTypeId1 != v.Expected.ContentTypeId1 {
			t.Fatalf("Expected %q but got %q for ContentTypeId1", v.Expected.ContentTypeId1, actual.ContentTypeId1)
		}

	}
}

func TestSegmentsForMeDriveIdListContentTypeIdBaseTypeId(t *testing.T) {
	segments := MeDriveIdListContentTypeIdBaseTypeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDriveIdListContentTypeIdBaseTypeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
