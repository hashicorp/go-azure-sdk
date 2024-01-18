package managementgroups

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SubscriptionId{}

func TestNewSubscriptionID(t *testing.T) {
	id := NewSubscriptionID("groupIdValue", "12345678-1234-9876-4563-123456789012")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}
}

func TestFormatSubscriptionID(t *testing.T) {
	actual := NewSubscriptionID("groupIdValue", "12345678-1234-9876-4563-123456789012").ID()
	expected := "/providers/Microsoft.Management/managementGroups/groupIdValue/subscriptions/12345678-1234-9876-4563-123456789012"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSubscriptionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SubscriptionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/groupIdValue/subscriptions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/groupIdValue/subscriptions/12345678-1234-9876-4563-123456789012",
			Expected: &SubscriptionId{
				GroupId:        "groupIdValue",
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/groupIdValue/subscriptions/12345678-1234-9876-4563-123456789012/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSubscriptionID(v.Input)
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

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

	}
}

func TestParseSubscriptionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SubscriptionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/gRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/groupIdValue/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/gRoUpIdVaLuE/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/groupIdValue/subscriptions/12345678-1234-9876-4563-123456789012",
			Expected: &SubscriptionId{
				GroupId:        "groupIdValue",
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/groupIdValue/subscriptions/12345678-1234-9876-4563-123456789012/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/gRoUpIdVaLuE/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Expected: &SubscriptionId{
				GroupId:        "gRoUpIdVaLuE",
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/gRoUpIdVaLuE/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSubscriptionIDInsensitively(v.Input)
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

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

	}
}

func TestSegmentsForSubscriptionId(t *testing.T) {
	segments := SubscriptionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SubscriptionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
