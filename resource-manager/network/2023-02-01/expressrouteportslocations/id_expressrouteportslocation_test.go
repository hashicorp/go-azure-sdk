package expressrouteportslocations

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ExpressRoutePortsLocationId{}

func TestNewExpressRoutePortsLocationID(t *testing.T) {
	id := NewExpressRoutePortsLocationID("12345678-1234-9876-4563-123456789012", "expressRoutePortsLocationValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ExpressRoutePortsLocationName != "expressRoutePortsLocationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ExpressRoutePortsLocationName'", id.ExpressRoutePortsLocationName, "expressRoutePortsLocationValue")
	}
}

func TestFormatExpressRoutePortsLocationID(t *testing.T) {
	actual := NewExpressRoutePortsLocationID("12345678-1234-9876-4563-123456789012", "expressRoutePortsLocationValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/expressRoutePortsLocations/expressRoutePortsLocationValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseExpressRoutePortsLocationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ExpressRoutePortsLocationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/expressRoutePortsLocations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/expressRoutePortsLocations/expressRoutePortsLocationValue",
			Expected: &ExpressRoutePortsLocationId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ExpressRoutePortsLocationName: "expressRoutePortsLocationValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/expressRoutePortsLocations/expressRoutePortsLocationValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseExpressRoutePortsLocationID(v.Input)
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

		if actual.ExpressRoutePortsLocationName != v.Expected.ExpressRoutePortsLocationName {
			t.Fatalf("Expected %q but got %q for ExpressRoutePortsLocationName", v.Expected.ExpressRoutePortsLocationName, actual.ExpressRoutePortsLocationName)
		}

	}
}

func TestParseExpressRoutePortsLocationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ExpressRoutePortsLocationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.nEtWoRk",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/expressRoutePortsLocations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.nEtWoRk/eXpReSsRoUtEpOrTsLoCaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/expressRoutePortsLocations/expressRoutePortsLocationValue",
			Expected: &ExpressRoutePortsLocationId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ExpressRoutePortsLocationName: "expressRoutePortsLocationValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/expressRoutePortsLocations/expressRoutePortsLocationValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.nEtWoRk/eXpReSsRoUtEpOrTsLoCaTiOnS/eXpReSsRoUtEpOrTsLoCaTiOnVaLuE",
			Expected: &ExpressRoutePortsLocationId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ExpressRoutePortsLocationName: "eXpReSsRoUtEpOrTsLoCaTiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.nEtWoRk/eXpReSsRoUtEpOrTsLoCaTiOnS/eXpReSsRoUtEpOrTsLoCaTiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseExpressRoutePortsLocationIDInsensitively(v.Input)
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

		if actual.ExpressRoutePortsLocationName != v.Expected.ExpressRoutePortsLocationName {
			t.Fatalf("Expected %q but got %q for ExpressRoutePortsLocationName", v.Expected.ExpressRoutePortsLocationName, actual.ExpressRoutePortsLocationName)
		}

	}
}

func TestSegmentsForExpressRoutePortsLocationId(t *testing.T) {
	segments := ExpressRoutePortsLocationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ExpressRoutePortsLocationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
