package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdRootAnalyticsItemActivityStatIdActivityId{}

func TestNewMeDriveIdRootAnalyticsItemActivityStatIdActivityID(t *testing.T) {
	id := NewMeDriveIdRootAnalyticsItemActivityStatIdActivityID("driveId", "itemActivityStatId", "itemActivityId")

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.ItemActivityStatId != "itemActivityStatId" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemActivityStatId'", id.ItemActivityStatId, "itemActivityStatId")
	}

	if id.ItemActivityId != "itemActivityId" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemActivityId'", id.ItemActivityId, "itemActivityId")
	}
}

func TestFormatMeDriveIdRootAnalyticsItemActivityStatIdActivityID(t *testing.T) {
	actual := NewMeDriveIdRootAnalyticsItemActivityStatIdActivityID("driveId", "itemActivityStatId", "itemActivityId").ID()
	expected := "/me/drives/driveId/root/analytics/itemActivityStats/itemActivityStatId/activities/itemActivityId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDriveIdRootAnalyticsItemActivityStatIdActivityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdRootAnalyticsItemActivityStatIdActivityId
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
			// Incomplete URI
			Input: "/me/drives/driveId/root/analytics/itemActivityStats/itemActivityStatId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/root/analytics/itemActivityStats/itemActivityStatId/activities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/root/analytics/itemActivityStats/itemActivityStatId/activities/itemActivityId",
			Expected: &MeDriveIdRootAnalyticsItemActivityStatIdActivityId{
				DriveId:            "driveId",
				ItemActivityStatId: "itemActivityStatId",
				ItemActivityId:     "itemActivityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/root/analytics/itemActivityStats/itemActivityStatId/activities/itemActivityId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdRootAnalyticsItemActivityStatIdActivityID(v.Input)
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

		if actual.ItemActivityId != v.Expected.ItemActivityId {
			t.Fatalf("Expected %q but got %q for ItemActivityId", v.Expected.ItemActivityId, actual.ItemActivityId)
		}

	}
}

func TestParseMeDriveIdRootAnalyticsItemActivityStatIdActivityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdRootAnalyticsItemActivityStatIdActivityId
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
			// Incomplete URI
			Input: "/me/drives/driveId/root/analytics/itemActivityStats/itemActivityStatId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/rOoT/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveId/root/analytics/itemActivityStats/itemActivityStatId/activities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/rOoT/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiD/aCtIvItIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveId/root/analytics/itemActivityStats/itemActivityStatId/activities/itemActivityId",
			Expected: &MeDriveIdRootAnalyticsItemActivityStatIdActivityId{
				DriveId:            "driveId",
				ItemActivityStatId: "itemActivityStatId",
				ItemActivityId:     "itemActivityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveId/root/analytics/itemActivityStats/itemActivityStatId/activities/itemActivityId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/rOoT/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiD/aCtIvItIeS/iTeMaCtIvItYiD",
			Expected: &MeDriveIdRootAnalyticsItemActivityStatIdActivityId{
				DriveId:            "dRiVeId",
				ItemActivityStatId: "iTeMaCtIvItYsTaTiD",
				ItemActivityId:     "iTeMaCtIvItYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeId/rOoT/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiD/aCtIvItIeS/iTeMaCtIvItYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdRootAnalyticsItemActivityStatIdActivityIDInsensitively(v.Input)
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

		if actual.ItemActivityId != v.Expected.ItemActivityId {
			t.Fatalf("Expected %q but got %q for ItemActivityId", v.Expected.ItemActivityId, actual.ItemActivityId)
		}

	}
}

func TestSegmentsForMeDriveIdRootAnalyticsItemActivityStatIdActivityId(t *testing.T) {
	segments := MeDriveIdRootAnalyticsItemActivityStatIdActivityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDriveIdRootAnalyticsItemActivityStatIdActivityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
