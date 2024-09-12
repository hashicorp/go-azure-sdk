package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdRootSubscriptionId{}

func TestNewMeDriveIdRootSubscriptionID(t *testing.T) {
	id := NewMeDriveIdRootSubscriptionID("driveIdValue", "subscriptionIdValue")

	if id.DriveId != "driveIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveIdValue")
	}

	if id.SubscriptionId != "subscriptionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "subscriptionIdValue")
	}
}

func TestFormatMeDriveIdRootSubscriptionID(t *testing.T) {
	actual := NewMeDriveIdRootSubscriptionID("driveIdValue", "subscriptionIdValue").ID()
	expected := "/me/drives/driveIdValue/root/subscriptions/subscriptionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDriveIdRootSubscriptionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdRootSubscriptionId
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
			Input: "/me/drives/driveIdValue/root/subscriptions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveIdValue/root/subscriptions/subscriptionIdValue",
			Expected: &MeDriveIdRootSubscriptionId{
				DriveId:        "driveIdValue",
				SubscriptionId: "subscriptionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveIdValue/root/subscriptions/subscriptionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdRootSubscriptionID(v.Input)
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

func TestParseMeDriveIdRootSubscriptionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDriveIdRootSubscriptionId
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
			Input: "/me/drives/driveIdValue/root/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/rOoT/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/drives/driveIdValue/root/subscriptions/subscriptionIdValue",
			Expected: &MeDriveIdRootSubscriptionId{
				DriveId:        "driveIdValue",
				SubscriptionId: "subscriptionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/drives/driveIdValue/root/subscriptions/subscriptionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/rOoT/sUbScRiPtIoNs/sUbScRiPtIoNiDvAlUe",
			Expected: &MeDriveIdRootSubscriptionId{
				DriveId:        "dRiVeIdVaLuE",
				SubscriptionId: "sUbScRiPtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dRiVeS/dRiVeIdVaLuE/rOoT/sUbScRiPtIoNs/sUbScRiPtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDriveIdRootSubscriptionIDInsensitively(v.Input)
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

func TestSegmentsForMeDriveIdRootSubscriptionId(t *testing.T) {
	segments := MeDriveIdRootSubscriptionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDriveIdRootSubscriptionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
