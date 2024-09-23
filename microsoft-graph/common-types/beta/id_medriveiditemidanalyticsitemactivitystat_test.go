package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdItemIdAnalyticsItemActivityStatId{}

func TestNewMeDriveIdItemIdAnalyticsItemActivityStatID(t *testing.T) {
	id := NewMeDriveIdItemIdAnalyticsItemActivityStatID("driveId", "driveItemId", "itemActivityStatId")

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.DriveItemId != "driveItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveItemId'", id.DriveItemId, "driveItemId")
	}

	if id.ItemActivityStatId != "itemActivityStatId" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemActivityStatId'", id.ItemActivityStatId, "itemActivityStatId")
	}
}

func TestFormatMeDriveIdItemIdAnalyticsItemActivityStatID(t *testing.T) {
	actual := NewMeDriveIdItemIdAnalyticsItemActivityStatID("driveId", "driveItemId", "itemActivityStatId").ID()
	expected := "/me/drives/driveId/items/driveItemId/analytics/itemActivityStats/itemActivityStatId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDriveIdItemIdAnalyticsItemActivityStatID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdItemIdAnalyticsItemActivityStatId
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
			Input: "/me/drives/driveId/items",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/items/driveItemId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/items/driveItemId/analytics",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/items/driveItemId/analytics/itemActivityStats",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/items/driveItemId/analytics/itemActivityStats/itemActivityStatId",
			Expected: &MeDriveIdItemIdAnalyticsItemActivityStatId{
				DriveId:            "driveId",
				DriveItemId:        "driveItemId",
				ItemActivityStatId: "itemActivityStatId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/items/driveItemId/analytics/itemActivityStats/itemActivityStatId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdItemIdAnalyticsItemActivityStatID(v.Input)
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

		if actual.DriveItemId != v.Expected.DriveItemId {
			t.Fatalf("Expected %q but got %q for DriveItemId", v.Expected.DriveItemId, actual.DriveItemId)
		}

		if actual.ItemActivityStatId != v.Expected.ItemActivityStatId {
			t.Fatalf("Expected %q but got %q for ItemActivityStatId", v.Expected.ItemActivityStatId, actual.ItemActivityStatId)
		}

	}
}

func TestParseMeDriveIdItemIdAnalyticsItemActivityStatIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdItemIdAnalyticsItemActivityStatId
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
			Input: "/me/drives/driveId/items",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/iTeMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/items/driveItemId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/items/driveItemId/analytics",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/aNaLyTiCs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/items/driveItemId/analytics/itemActivityStats",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/aNaLyTiCs/iTeMaCtIvItYsTaTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/items/driveItemId/analytics/itemActivityStats/itemActivityStatId",
			Expected: &MeDriveIdItemIdAnalyticsItemActivityStatId{
				DriveId:            "driveId",
				DriveItemId:        "driveItemId",
				ItemActivityStatId: "itemActivityStatId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/items/driveItemId/analytics/itemActivityStats/itemActivityStatId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiD",
			Expected: &MeDriveIdItemIdAnalyticsItemActivityStatId{
				DriveId:            "dRiVeId",
				DriveItemId:        "dRiVeItEmId",
				ItemActivityStatId: "iTeMaCtIvItYsTaTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/iTeMs/dRiVeItEmId/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdItemIdAnalyticsItemActivityStatIDInsensitively(v.Input)
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

		if actual.DriveItemId != v.Expected.DriveItemId {
			t.Fatalf("Expected %q but got %q for DriveItemId", v.Expected.DriveItemId, actual.DriveItemId)
		}

		if actual.ItemActivityStatId != v.Expected.ItemActivityStatId {
			t.Fatalf("Expected %q but got %q for ItemActivityStatId", v.Expected.ItemActivityStatId, actual.ItemActivityStatId)
		}

	}
}

func TestSegmentsForMeDriveIdItemIdAnalyticsItemActivityStatId(t *testing.T) {
	segments := MeDriveIdItemIdAnalyticsItemActivityStatId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDriveIdItemIdAnalyticsItemActivityStatId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
