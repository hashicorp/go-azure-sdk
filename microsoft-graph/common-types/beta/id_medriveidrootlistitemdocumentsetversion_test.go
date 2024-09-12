package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdRootListItemDocumentSetVersionId{}

func TestNewMeDriveIdRootListItemDocumentSetVersionID(t *testing.T) {
	id := NewMeDriveIdRootListItemDocumentSetVersionID("driveIdValue", "documentSetVersionIdValue")

	if id.DriveId != "driveIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveIdValue")
	}

	if id.DocumentSetVersionId != "documentSetVersionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DocumentSetVersionId'", id.DocumentSetVersionId, "documentSetVersionIdValue")
	}
}

func TestFormatMeDriveIdRootListItemDocumentSetVersionID(t *testing.T) {
	actual := NewMeDriveIdRootListItemDocumentSetVersionID("driveIdValue", "documentSetVersionIdValue").ID()
	expected := "/me/drives/driveIdValue/root/listItem/documentSetVersions/documentSetVersionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDriveIdRootListItemDocumentSetVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdRootListItemDocumentSetVersionId
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
			Input: "/me/drives/driveIdValue/root",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/root/listItem",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/root/listItem/documentSetVersions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveIdValue/root/listItem/documentSetVersions/documentSetVersionIdValue",
			Expected: &MeDriveIdRootListItemDocumentSetVersionId{
				DriveId:              "driveIdValue",
				DocumentSetVersionId: "documentSetVersionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveIdValue/root/listItem/documentSetVersions/documentSetVersionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdRootListItemDocumentSetVersionID(v.Input)
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

		if actual.DocumentSetVersionId != v.Expected.DocumentSetVersionId {
			t.Fatalf("Expected %q but got %q for DocumentSetVersionId", v.Expected.DocumentSetVersionId, actual.DocumentSetVersionId)
		}

	}
}

func TestParseMeDriveIdRootListItemDocumentSetVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdRootListItemDocumentSetVersionId
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
			Input: "/me/drives/driveIdValue/root",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/rOoT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/root/listItem",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/rOoT/lIsTiTeM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/root/listItem/documentSetVersions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/rOoT/lIsTiTeM/dOcUmEnTsEtVeRsIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveIdValue/root/listItem/documentSetVersions/documentSetVersionIdValue",
			Expected: &MeDriveIdRootListItemDocumentSetVersionId{
				DriveId:              "driveIdValue",
				DocumentSetVersionId: "documentSetVersionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveIdValue/root/listItem/documentSetVersions/documentSetVersionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/rOoT/lIsTiTeM/dOcUmEnTsEtVeRsIoNs/dOcUmEnTsEtVeRsIoNiDvAlUe",
			Expected: &MeDriveIdRootListItemDocumentSetVersionId{
				DriveId:              "dRiVeIdVaLuE",
				DocumentSetVersionId: "dOcUmEnTsEtVeRsIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/rOoT/lIsTiTeM/dOcUmEnTsEtVeRsIoNs/dOcUmEnTsEtVeRsIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdRootListItemDocumentSetVersionIDInsensitively(v.Input)
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

		if actual.DocumentSetVersionId != v.Expected.DocumentSetVersionId {
			t.Fatalf("Expected %q but got %q for DocumentSetVersionId", v.Expected.DocumentSetVersionId, actual.DocumentSetVersionId)
		}

	}
}

func TestSegmentsForMeDriveIdRootListItemDocumentSetVersionId(t *testing.T) {
	segments := MeDriveIdRootListItemDocumentSetVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDriveIdRootListItemDocumentSetVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
