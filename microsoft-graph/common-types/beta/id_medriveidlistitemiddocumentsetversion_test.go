package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListItemIdDocumentSetVersionId{}

func TestNewMeDriveIdListItemIdDocumentSetVersionID(t *testing.T) {
	id := NewMeDriveIdListItemIdDocumentSetVersionID("driveIdValue", "listItemIdValue", "documentSetVersionIdValue")

	if id.DriveId != "driveIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveIdValue")
	}

	if id.ListItemId != "listItemIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ListItemId'", id.ListItemId, "listItemIdValue")
	}

	if id.DocumentSetVersionId != "documentSetVersionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DocumentSetVersionId'", id.DocumentSetVersionId, "documentSetVersionIdValue")
	}
}

func TestFormatMeDriveIdListItemIdDocumentSetVersionID(t *testing.T) {
	actual := NewMeDriveIdListItemIdDocumentSetVersionID("driveIdValue", "listItemIdValue", "documentSetVersionIdValue").ID()
	expected := "/me/drives/driveIdValue/list/items/listItemIdValue/documentSetVersions/documentSetVersionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDriveIdListItemIdDocumentSetVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdListItemIdDocumentSetVersionId
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
			Input: "/me/drives/driveIdValue/list/items",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/list/items/listItemIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/list/items/listItemIdValue/documentSetVersions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveIdValue/list/items/listItemIdValue/documentSetVersions/documentSetVersionIdValue",
			Expected: &MeDriveIdListItemIdDocumentSetVersionId{
				DriveId:              "driveIdValue",
				ListItemId:           "listItemIdValue",
				DocumentSetVersionId: "documentSetVersionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveIdValue/list/items/listItemIdValue/documentSetVersions/documentSetVersionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdListItemIdDocumentSetVersionID(v.Input)
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

		if actual.ListItemId != v.Expected.ListItemId {
			t.Fatalf("Expected %q but got %q for ListItemId", v.Expected.ListItemId, actual.ListItemId)
		}

		if actual.DocumentSetVersionId != v.Expected.DocumentSetVersionId {
			t.Fatalf("Expected %q but got %q for DocumentSetVersionId", v.Expected.DocumentSetVersionId, actual.DocumentSetVersionId)
		}

	}
}

func TestParseMeDriveIdListItemIdDocumentSetVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdListItemIdDocumentSetVersionId
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
			Input: "/me/drives/driveIdValue/list/items",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/lIsT/iTeMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/list/items/listItemIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/lIsT/iTeMs/lIsTiTeMiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/list/items/listItemIdValue/documentSetVersions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/lIsT/iTeMs/lIsTiTeMiDvAlUe/dOcUmEnTsEtVeRsIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveIdValue/list/items/listItemIdValue/documentSetVersions/documentSetVersionIdValue",
			Expected: &MeDriveIdListItemIdDocumentSetVersionId{
				DriveId:              "driveIdValue",
				ListItemId:           "listItemIdValue",
				DocumentSetVersionId: "documentSetVersionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveIdValue/list/items/listItemIdValue/documentSetVersions/documentSetVersionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/lIsT/iTeMs/lIsTiTeMiDvAlUe/dOcUmEnTsEtVeRsIoNs/dOcUmEnTsEtVeRsIoNiDvAlUe",
			Expected: &MeDriveIdListItemIdDocumentSetVersionId{
				DriveId:              "dRiVeIdVaLuE",
				ListItemId:           "lIsTiTeMiDvAlUe",
				DocumentSetVersionId: "dOcUmEnTsEtVeRsIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/lIsT/iTeMs/lIsTiTeMiDvAlUe/dOcUmEnTsEtVeRsIoNs/dOcUmEnTsEtVeRsIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdListItemIdDocumentSetVersionIDInsensitively(v.Input)
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

		if actual.ListItemId != v.Expected.ListItemId {
			t.Fatalf("Expected %q but got %q for ListItemId", v.Expected.ListItemId, actual.ListItemId)
		}

		if actual.DocumentSetVersionId != v.Expected.DocumentSetVersionId {
			t.Fatalf("Expected %q but got %q for DocumentSetVersionId", v.Expected.DocumentSetVersionId, actual.DocumentSetVersionId)
		}

	}
}

func TestSegmentsForMeDriveIdListItemIdDocumentSetVersionId(t *testing.T) {
	segments := MeDriveIdListItemIdDocumentSetVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDriveIdListItemIdDocumentSetVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
