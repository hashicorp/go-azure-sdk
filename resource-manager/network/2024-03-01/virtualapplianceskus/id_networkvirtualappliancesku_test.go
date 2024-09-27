package virtualapplianceskus

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &NetworkVirtualApplianceSkuId{}

func TestNewNetworkVirtualApplianceSkuID(t *testing.T) {
	id := NewNetworkVirtualApplianceSkuID("12345678-1234-9876-4563-123456789012", "skuName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.NetworkVirtualApplianceSkuName != "skuName" {
		t.Fatalf("Expected %q but got %q for Segment 'NetworkVirtualApplianceSkuName'", id.NetworkVirtualApplianceSkuName, "skuName")
	}
}

func TestFormatNetworkVirtualApplianceSkuID(t *testing.T) {
	actual := NewNetworkVirtualApplianceSkuID("12345678-1234-9876-4563-123456789012", "skuName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/networkVirtualApplianceSkus/skuName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseNetworkVirtualApplianceSkuID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *NetworkVirtualApplianceSkuId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/networkVirtualApplianceSkus",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/networkVirtualApplianceSkus/skuName",
			Expected: &NetworkVirtualApplianceSkuId{
				SubscriptionId:                 "12345678-1234-9876-4563-123456789012",
				NetworkVirtualApplianceSkuName: "skuName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/networkVirtualApplianceSkus/skuName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseNetworkVirtualApplianceSkuID(v.Input)
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

		if actual.NetworkVirtualApplianceSkuName != v.Expected.NetworkVirtualApplianceSkuName {
			t.Fatalf("Expected %q but got %q for NetworkVirtualApplianceSkuName", v.Expected.NetworkVirtualApplianceSkuName, actual.NetworkVirtualApplianceSkuName)
		}

	}
}

func TestParseNetworkVirtualApplianceSkuIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *NetworkVirtualApplianceSkuId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/networkVirtualApplianceSkus",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkViRtUaLaPpLiAnCeSkUs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/networkVirtualApplianceSkus/skuName",
			Expected: &NetworkVirtualApplianceSkuId{
				SubscriptionId:                 "12345678-1234-9876-4563-123456789012",
				NetworkVirtualApplianceSkuName: "skuName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/networkVirtualApplianceSkus/skuName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkViRtUaLaPpLiAnCeSkUs/sKuNaMe",
			Expected: &NetworkVirtualApplianceSkuId{
				SubscriptionId:                 "12345678-1234-9876-4563-123456789012",
				NetworkVirtualApplianceSkuName: "sKuNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkViRtUaLaPpLiAnCeSkUs/sKuNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseNetworkVirtualApplianceSkuIDInsensitively(v.Input)
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

		if actual.NetworkVirtualApplianceSkuName != v.Expected.NetworkVirtualApplianceSkuName {
			t.Fatalf("Expected %q but got %q for NetworkVirtualApplianceSkuName", v.Expected.NetworkVirtualApplianceSkuName, actual.NetworkVirtualApplianceSkuName)
		}

	}
}

func TestSegmentsForNetworkVirtualApplianceSkuId(t *testing.T) {
	segments := NetworkVirtualApplianceSkuId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("NetworkVirtualApplianceSkuId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
