package capabilitytypes

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &CapabilityTypeId{}

func TestNewCapabilityTypeID(t *testing.T) {
	id := NewCapabilityTypeID("12345678-1234-9876-4563-123456789012", "locationName", "targetTypeName", "capabilityTypeName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.LocationName != "locationName" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "locationName")
	}

	if id.TargetTypeName != "targetTypeName" {
		t.Fatalf("Expected %q but got %q for Segment 'TargetTypeName'", id.TargetTypeName, "targetTypeName")
	}

	if id.CapabilityTypeName != "capabilityTypeName" {
		t.Fatalf("Expected %q but got %q for Segment 'CapabilityTypeName'", id.CapabilityTypeName, "capabilityTypeName")
	}
}

func TestFormatCapabilityTypeID(t *testing.T) {
	actual := NewCapabilityTypeID("12345678-1234-9876-4563-123456789012", "locationName", "targetTypeName", "capabilityTypeName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Chaos/locations/locationName/targetTypes/targetTypeName/capabilityTypes/capabilityTypeName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseCapabilityTypeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CapabilityTypeId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Chaos",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Chaos/locations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Chaos/locations/locationName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Chaos/locations/locationName/targetTypes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Chaos/locations/locationName/targetTypes/targetTypeName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Chaos/locations/locationName/targetTypes/targetTypeName/capabilityTypes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Chaos/locations/locationName/targetTypes/targetTypeName/capabilityTypes/capabilityTypeName",
			Expected: &CapabilityTypeId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				LocationName:       "locationName",
				TargetTypeName:     "targetTypeName",
				CapabilityTypeName: "capabilityTypeName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Chaos/locations/locationName/targetTypes/targetTypeName/capabilityTypes/capabilityTypeName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCapabilityTypeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.TargetTypeName != v.Expected.TargetTypeName {
			t.Fatalf("Expected %q but got %q for TargetTypeName", v.Expected.TargetTypeName, actual.TargetTypeName)
		}

		if actual.CapabilityTypeName != v.Expected.CapabilityTypeName {
			t.Fatalf("Expected %q but got %q for CapabilityTypeName", v.Expected.CapabilityTypeName, actual.CapabilityTypeName)
		}

	}
}

func TestParseCapabilityTypeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CapabilityTypeId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Chaos",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cHaOs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Chaos/locations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cHaOs/lOcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Chaos/locations/locationName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cHaOs/lOcAtIoNs/lOcAtIoNnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Chaos/locations/locationName/targetTypes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cHaOs/lOcAtIoNs/lOcAtIoNnAmE/tArGeTtYpEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Chaos/locations/locationName/targetTypes/targetTypeName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cHaOs/lOcAtIoNs/lOcAtIoNnAmE/tArGeTtYpEs/tArGeTtYpEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Chaos/locations/locationName/targetTypes/targetTypeName/capabilityTypes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cHaOs/lOcAtIoNs/lOcAtIoNnAmE/tArGeTtYpEs/tArGeTtYpEnAmE/cApAbIlItYtYpEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Chaos/locations/locationName/targetTypes/targetTypeName/capabilityTypes/capabilityTypeName",
			Expected: &CapabilityTypeId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				LocationName:       "locationName",
				TargetTypeName:     "targetTypeName",
				CapabilityTypeName: "capabilityTypeName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Chaos/locations/locationName/targetTypes/targetTypeName/capabilityTypes/capabilityTypeName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cHaOs/lOcAtIoNs/lOcAtIoNnAmE/tArGeTtYpEs/tArGeTtYpEnAmE/cApAbIlItYtYpEs/cApAbIlItYtYpEnAmE",
			Expected: &CapabilityTypeId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				LocationName:       "lOcAtIoNnAmE",
				TargetTypeName:     "tArGeTtYpEnAmE",
				CapabilityTypeName: "cApAbIlItYtYpEnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cHaOs/lOcAtIoNs/lOcAtIoNnAmE/tArGeTtYpEs/tArGeTtYpEnAmE/cApAbIlItYtYpEs/cApAbIlItYtYpEnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCapabilityTypeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.TargetTypeName != v.Expected.TargetTypeName {
			t.Fatalf("Expected %q but got %q for TargetTypeName", v.Expected.TargetTypeName, actual.TargetTypeName)
		}

		if actual.CapabilityTypeName != v.Expected.CapabilityTypeName {
			t.Fatalf("Expected %q but got %q for CapabilityTypeName", v.Expected.CapabilityTypeName, actual.CapabilityTypeName)
		}

	}
}

func TestSegmentsForCapabilityTypeId(t *testing.T) {
	segments := CapabilityTypeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("CapabilityTypeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
