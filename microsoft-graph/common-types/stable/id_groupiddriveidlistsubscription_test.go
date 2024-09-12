package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdListSubscriptionId{}

func TestNewGroupIdDriveIdListSubscriptionID(t *testing.T) {
	id := NewGroupIdDriveIdListSubscriptionID("groupIdValue", "driveIdValue", "subscriptionIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.DriveId != "driveIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveIdValue")
	}

	if id.SubscriptionId != "subscriptionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "subscriptionIdValue")
	}
}

func TestFormatGroupIdDriveIdListSubscriptionID(t *testing.T) {
	actual := NewGroupIdDriveIdListSubscriptionID("groupIdValue", "driveIdValue", "subscriptionIdValue").ID()
	expected := "/groups/groupIdValue/drives/driveIdValue/list/subscriptions/subscriptionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdDriveIdListSubscriptionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdDriveIdListSubscriptionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/list",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/list/subscriptions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/drives/driveIdValue/list/subscriptions/subscriptionIdValue",
			Expected: &GroupIdDriveIdListSubscriptionId{
				GroupId:        "groupIdValue",
				DriveId:        "driveIdValue",
				SubscriptionId: "subscriptionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/drives/driveIdValue/list/subscriptions/subscriptionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdDriveIdListSubscriptionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.DriveId != v.Expected.DriveId {
			t.Fatalf("Expected %q but got %q for DriveId", v.Expected.DriveId, actual.DriveId)
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

	}
}

func TestParseGroupIdDriveIdListSubscriptionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdDriveIdListSubscriptionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/list",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE/lIsT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/drives/driveIdValue/list/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE/lIsT/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/drives/driveIdValue/list/subscriptions/subscriptionIdValue",
			Expected: &GroupIdDriveIdListSubscriptionId{
				GroupId:        "groupIdValue",
				DriveId:        "driveIdValue",
				SubscriptionId: "subscriptionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/drives/driveIdValue/list/subscriptions/subscriptionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE/lIsT/sUbScRiPtIoNs/sUbScRiPtIoNiDvAlUe",
			Expected: &GroupIdDriveIdListSubscriptionId{
				GroupId:        "gRoUpIdVaLuE",
				DriveId:        "dRiVeIdVaLuE",
				SubscriptionId: "sUbScRiPtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/dRiVeS/dRiVeIdVaLuE/lIsT/sUbScRiPtIoNs/sUbScRiPtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdDriveIdListSubscriptionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.DriveId != v.Expected.DriveId {
			t.Fatalf("Expected %q but got %q for DriveId", v.Expected.DriveId, actual.DriveId)
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

	}
}

func TestSegmentsForGroupIdDriveIdListSubscriptionId(t *testing.T) {
	segments := GroupIdDriveIdListSubscriptionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdDriveIdListSubscriptionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
