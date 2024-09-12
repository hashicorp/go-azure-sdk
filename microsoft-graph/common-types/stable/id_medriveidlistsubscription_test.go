package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListSubscriptionId{}

func TestNewMeDriveIdListSubscriptionID(t *testing.T) {
	id := NewMeDriveIdListSubscriptionID("driveIdValue", "subscriptionIdValue")

	if id.DriveId != "driveIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveIdValue")
	}

	if id.SubscriptionId != "subscriptionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "subscriptionIdValue")
	}
}

func TestFormatMeDriveIdListSubscriptionID(t *testing.T) {
	actual := NewMeDriveIdListSubscriptionID("driveIdValue", "subscriptionIdValue").ID()
	expected := "/me/drives/driveIdValue/list/subscriptions/subscriptionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDriveIdListSubscriptionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdListSubscriptionId
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
			Input: "/me/drives/driveIdValue/list/subscriptions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveIdValue/list/subscriptions/subscriptionIdValue",
			Expected: &MeDriveIdListSubscriptionId{
				DriveId:        "driveIdValue",
				SubscriptionId: "subscriptionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveIdValue/list/subscriptions/subscriptionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdListSubscriptionID(v.Input)
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

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

	}
}

func TestParseMeDriveIdListSubscriptionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdListSubscriptionId
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
			Input: "/me/drives/driveIdValue/list/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/lIsT/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveIdValue/list/subscriptions/subscriptionIdValue",
			Expected: &MeDriveIdListSubscriptionId{
				DriveId:        "driveIdValue",
				SubscriptionId: "subscriptionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveIdValue/list/subscriptions/subscriptionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/lIsT/sUbScRiPtIoNs/sUbScRiPtIoNiDvAlUe",
			Expected: &MeDriveIdListSubscriptionId{
				DriveId:        "dRiVeIdVaLuE",
				SubscriptionId: "sUbScRiPtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/lIsT/sUbScRiPtIoNs/sUbScRiPtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdListSubscriptionIDInsensitively(v.Input)
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

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

	}
}

func TestSegmentsForMeDriveIdListSubscriptionId(t *testing.T) {
	segments := MeDriveIdListSubscriptionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDriveIdListSubscriptionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
