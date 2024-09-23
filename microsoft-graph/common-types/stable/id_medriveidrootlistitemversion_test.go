package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdRootListItemVersionId{}

func TestNewMeDriveIdRootListItemVersionID(t *testing.T) {
	id := NewMeDriveIdRootListItemVersionID("driveId", "listItemVersionId")

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.ListItemVersionId != "listItemVersionId" {
		t.Fatalf("Expected %q but got %q for Segment 'ListItemVersionId'", id.ListItemVersionId, "listItemVersionId")
	}
}

func TestFormatMeDriveIdRootListItemVersionID(t *testing.T) {
	actual := NewMeDriveIdRootListItemVersionID("driveId", "listItemVersionId").ID()
	expected := "/me/drives/driveId/root/listItem/versions/listItemVersionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDriveIdRootListItemVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdRootListItemVersionId
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
			Input: "/me/drives/driveId/root",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/root/listItem",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/root/listItem/versions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/root/listItem/versions/listItemVersionId",
			Expected: &MeDriveIdRootListItemVersionId{
				DriveId:           "driveId",
				ListItemVersionId: "listItemVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/root/listItem/versions/listItemVersionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdRootListItemVersionID(v.Input)
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

		if actual.ListItemVersionId != v.Expected.ListItemVersionId {
			t.Fatalf("Expected %q but got %q for ListItemVersionId", v.Expected.ListItemVersionId, actual.ListItemVersionId)
		}

	}
}

func TestParseMeDriveIdRootListItemVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdRootListItemVersionId
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
			Input: "/me/drives/driveId/root",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/rOoT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/root/listItem",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/rOoT/lIsTiTeM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/root/listItem/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/rOoT/lIsTiTeM/vErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/root/listItem/versions/listItemVersionId",
			Expected: &MeDriveIdRootListItemVersionId{
				DriveId:           "driveId",
				ListItemVersionId: "listItemVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/root/listItem/versions/listItemVersionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/rOoT/lIsTiTeM/vErSiOnS/lIsTiTeMvErSiOnId",
			Expected: &MeDriveIdRootListItemVersionId{
				DriveId:           "dRiVeId",
				ListItemVersionId: "lIsTiTeMvErSiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/rOoT/lIsTiTeM/vErSiOnS/lIsTiTeMvErSiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdRootListItemVersionIDInsensitively(v.Input)
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

		if actual.ListItemVersionId != v.Expected.ListItemVersionId {
			t.Fatalf("Expected %q but got %q for ListItemVersionId", v.Expected.ListItemVersionId, actual.ListItemVersionId)
		}

	}
}

func TestSegmentsForMeDriveIdRootListItemVersionId(t *testing.T) {
	segments := MeDriveIdRootListItemVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDriveIdRootListItemVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
