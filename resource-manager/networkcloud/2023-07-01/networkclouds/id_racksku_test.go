package networkclouds

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RackSkuId{}

func TestNewRackSkuID(t *testing.T) {
	id := NewRackSkuID("12345678-1234-9876-4563-123456789012", "rackSkuValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.RackSkuName != "rackSkuValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RackSkuName'", id.RackSkuName, "rackSkuValue")
	}
}

func TestFormatRackSkuID(t *testing.T) {
	actual := NewRackSkuID("12345678-1234-9876-4563-123456789012", "rackSkuValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.NetworkCloud/rackSkus/rackSkuValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRackSkuID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RackSkuId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.NetworkCloud",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.NetworkCloud/rackSkus",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.NetworkCloud/rackSkus/rackSkuValue",
			Expected: &RackSkuId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				RackSkuName:    "rackSkuValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.NetworkCloud/rackSkus/rackSkuValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRackSkuID(v.Input)
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

		if actual.RackSkuName != v.Expected.RackSkuName {
			t.Fatalf("Expected %q but got %q for RackSkuName", v.Expected.RackSkuName, actual.RackSkuName)
		}

	}
}

func TestParseRackSkuIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RackSkuId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.NetworkCloud",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.nEtWoRkClOuD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.NetworkCloud/rackSkus",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.nEtWoRkClOuD/rAcKsKuS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.NetworkCloud/rackSkus/rackSkuValue",
			Expected: &RackSkuId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				RackSkuName:    "rackSkuValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.NetworkCloud/rackSkus/rackSkuValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.nEtWoRkClOuD/rAcKsKuS/rAcKsKuVaLuE",
			Expected: &RackSkuId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				RackSkuName:    "rAcKsKuVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.nEtWoRkClOuD/rAcKsKuS/rAcKsKuVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRackSkuIDInsensitively(v.Input)
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

		if actual.RackSkuName != v.Expected.RackSkuName {
			t.Fatalf("Expected %q but got %q for RackSkuName", v.Expected.RackSkuName, actual.RackSkuName)
		}

	}
}

func TestSegmentsForRackSkuId(t *testing.T) {
	segments := RackSkuId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RackSkuId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
