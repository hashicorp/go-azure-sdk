package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdRootAnalyticsItemActivityStatId{}

func TestNewMeDriveIdRootAnalyticsItemActivityStatID(t *testing.T) {
	id := NewMeDriveIdRootAnalyticsItemActivityStatID("driveId", "itemActivityStatId")

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.ItemActivityStatId != "itemActivityStatId" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemActivityStatId'", id.ItemActivityStatId, "itemActivityStatId")
	}
}

func TestFormatMeDriveIdRootAnalyticsItemActivityStatID(t *testing.T) {
	actual := NewMeDriveIdRootAnalyticsItemActivityStatID("driveId", "itemActivityStatId").ID()
	expected := "/me/drives/driveId/root/analytics/itemActivityStats/itemActivityStatId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDriveIdRootAnalyticsItemActivityStatID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdRootAnalyticsItemActivityStatId
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
			Input: "/me/drives/driveId/root/analytics",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/root/analytics/itemActivityStats",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/root/analytics/itemActivityStats/itemActivityStatId",
			Expected: &MeDriveIdRootAnalyticsItemActivityStatId{
				DriveId:            "driveId",
				ItemActivityStatId: "itemActivityStatId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/root/analytics/itemActivityStats/itemActivityStatId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdRootAnalyticsItemActivityStatID(v.Input)
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

		if actual.ItemActivityStatId != v.Expected.ItemActivityStatId {
			t.Fatalf("Expected %q but got %q for ItemActivityStatId", v.Expected.ItemActivityStatId, actual.ItemActivityStatId)
		}

	}
}

func TestParseMeDriveIdRootAnalyticsItemActivityStatIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdRootAnalyticsItemActivityStatId
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
			Input: "/me/drives/driveId/root/analytics",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/rOoT/aNaLyTiCs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/root/analytics/itemActivityStats",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/rOoT/aNaLyTiCs/iTeMaCtIvItYsTaTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/root/analytics/itemActivityStats/itemActivityStatId",
			Expected: &MeDriveIdRootAnalyticsItemActivityStatId{
				DriveId:            "driveId",
				ItemActivityStatId: "itemActivityStatId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/root/analytics/itemActivityStats/itemActivityStatId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/rOoT/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiD",
			Expected: &MeDriveIdRootAnalyticsItemActivityStatId{
				DriveId:            "dRiVeId",
				ItemActivityStatId: "iTeMaCtIvItYsTaTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/rOoT/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdRootAnalyticsItemActivityStatIDInsensitively(v.Input)
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

		if actual.ItemActivityStatId != v.Expected.ItemActivityStatId {
			t.Fatalf("Expected %q but got %q for ItemActivityStatId", v.Expected.ItemActivityStatId, actual.ItemActivityStatId)
		}

	}
}

func TestSegmentsForMeDriveIdRootAnalyticsItemActivityStatId(t *testing.T) {
	segments := MeDriveIdRootAnalyticsItemActivityStatId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDriveIdRootAnalyticsItemActivityStatId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
