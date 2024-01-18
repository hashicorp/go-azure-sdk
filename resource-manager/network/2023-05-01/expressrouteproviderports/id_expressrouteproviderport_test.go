package expressrouteproviderports

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ExpressRouteProviderPortId{}

func TestNewExpressRouteProviderPortID(t *testing.T) {
	id := NewExpressRouteProviderPortID("12345678-1234-9876-4563-123456789012", "expressRouteProviderPortValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ExpressRouteProviderPortName != "expressRouteProviderPortValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ExpressRouteProviderPortName'", id.ExpressRouteProviderPortName, "expressRouteProviderPortValue")
	}
}

func TestFormatExpressRouteProviderPortID(t *testing.T) {
	actual := NewExpressRouteProviderPortID("12345678-1234-9876-4563-123456789012", "expressRouteProviderPortValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/expressRouteProviderPorts/expressRouteProviderPortValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseExpressRouteProviderPortID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ExpressRouteProviderPortId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/expressRouteProviderPorts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/expressRouteProviderPorts/expressRouteProviderPortValue",
			Expected: &ExpressRouteProviderPortId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ExpressRouteProviderPortName: "expressRouteProviderPortValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/expressRouteProviderPorts/expressRouteProviderPortValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseExpressRouteProviderPortID(v.Input)
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

		if actual.ExpressRouteProviderPortName != v.Expected.ExpressRouteProviderPortName {
			t.Fatalf("Expected %q but got %q for ExpressRouteProviderPortName", v.Expected.ExpressRouteProviderPortName, actual.ExpressRouteProviderPortName)
		}

	}
}

func TestParseExpressRouteProviderPortIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ExpressRouteProviderPortId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/expressRouteProviderPorts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.nEtWoRk/eXpReSsRoUtEpRoViDeRpOrTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/expressRouteProviderPorts/expressRouteProviderPortValue",
			Expected: &ExpressRouteProviderPortId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ExpressRouteProviderPortName: "expressRouteProviderPortValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/expressRouteProviderPorts/expressRouteProviderPortValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.nEtWoRk/eXpReSsRoUtEpRoViDeRpOrTs/eXpReSsRoUtEpRoViDeRpOrTvAlUe",
			Expected: &ExpressRouteProviderPortId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ExpressRouteProviderPortName: "eXpReSsRoUtEpRoViDeRpOrTvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.nEtWoRk/eXpReSsRoUtEpRoViDeRpOrTs/eXpReSsRoUtEpRoViDeRpOrTvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseExpressRouteProviderPortIDInsensitively(v.Input)
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

		if actual.ExpressRouteProviderPortName != v.Expected.ExpressRouteProviderPortName {
			t.Fatalf("Expected %q but got %q for ExpressRouteProviderPortName", v.Expected.ExpressRouteProviderPortName, actual.ExpressRouteProviderPortName)
		}

	}
}

func TestSegmentsForExpressRouteProviderPortId(t *testing.T) {
	segments := ExpressRouteProviderPortId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ExpressRouteProviderPortId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
