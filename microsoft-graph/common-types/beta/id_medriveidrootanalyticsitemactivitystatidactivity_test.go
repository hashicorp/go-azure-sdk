package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdRootAnalyticsItemActivityStatIdActivityId{}

func TestNewMeDriveIdRootAnalyticsItemActivityStatIdActivityID(t *testing.T) {
	id := NewMeDriveIdRootAnalyticsItemActivityStatIdActivityID("driveIdValue", "itemActivityStatIdValue", "itemActivityIdValue")

	if id.DriveId != "driveIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveIdValue")
	}

	if id.ItemActivityStatId != "itemActivityStatIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemActivityStatId'", id.ItemActivityStatId, "itemActivityStatIdValue")
	}

	if id.ItemActivityId != "itemActivityIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemActivityId'", id.ItemActivityId, "itemActivityIdValue")
	}
}

func TestFormatMeDriveIdRootAnalyticsItemActivityStatIdActivityID(t *testing.T) {
	actual := NewMeDriveIdRootAnalyticsItemActivityStatIdActivityID("driveIdValue", "itemActivityStatIdValue", "itemActivityIdValue").ID()
	expected := "/me/drives/driveIdValue/root/analytics/itemActivityStats/itemActivityStatIdValue/activities/itemActivityIdValue"
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
			Input: "/me/drives/driveIdValue/root/analytics",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/root/analytics/itemActivityStats",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/root/analytics/itemActivityStats/itemActivityStatIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/root/analytics/itemActivityStats/itemActivityStatIdValue/activities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveIdValue/root/analytics/itemActivityStats/itemActivityStatIdValue/activities/itemActivityIdValue",
			Expected: &MeDriveIdRootAnalyticsItemActivityStatIdActivityId{
				DriveId:            "driveIdValue",
				ItemActivityStatId: "itemActivityStatIdValue",
				ItemActivityId:     "itemActivityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveIdValue/root/analytics/itemActivityStats/itemActivityStatIdValue/activities/itemActivityIdValue/extra",
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
			Input: "/me/drives/driveIdValue/root/analytics",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/rOoT/aNaLyTiCs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/root/analytics/itemActivityStats",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/rOoT/aNaLyTiCs/iTeMaCtIvItYsTaTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/root/analytics/itemActivityStats/itemActivityStatIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/rOoT/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/drives/driveIdValue/root/analytics/itemActivityStats/itemActivityStatIdValue/activities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/rOoT/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiDvAlUe/aCtIvItIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveIdValue/root/analytics/itemActivityStats/itemActivityStatIdValue/activities/itemActivityIdValue",
			Expected: &MeDriveIdRootAnalyticsItemActivityStatIdActivityId{
				DriveId:            "driveIdValue",
				ItemActivityStatId: "itemActivityStatIdValue",
				ItemActivityId:     "itemActivityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveIdValue/root/analytics/itemActivityStats/itemActivityStatIdValue/activities/itemActivityIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/rOoT/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiDvAlUe/aCtIvItIeS/iTeMaCtIvItYiDvAlUe",
			Expected: &MeDriveIdRootAnalyticsItemActivityStatIdActivityId{
				DriveId:            "dRiVeIdVaLuE",
				ItemActivityStatId: "iTeMaCtIvItYsTaTiDvAlUe",
				ItemActivityId:     "iTeMaCtIvItYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/rOoT/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiDvAlUe/aCtIvItIeS/iTeMaCtIvItYiDvAlUe/extra",
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
