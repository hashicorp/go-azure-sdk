package virtualnetworklinks

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &VirtualNetworkLinkId{}

func TestNewVirtualNetworkLinkID(t *testing.T) {
	id := NewVirtualNetworkLinkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsForwardingRulesetName", "virtualNetworkLinkName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.DnsForwardingRulesetName != "dnsForwardingRulesetName" {
		t.Fatalf("Expected %q but got %q for Segment 'DnsForwardingRulesetName'", id.DnsForwardingRulesetName, "dnsForwardingRulesetName")
	}

	if id.VirtualNetworkLinkName != "virtualNetworkLinkName" {
		t.Fatalf("Expected %q but got %q for Segment 'VirtualNetworkLinkName'", id.VirtualNetworkLinkName, "virtualNetworkLinkName")
	}
}

func TestFormatVirtualNetworkLinkID(t *testing.T) {
	actual := NewVirtualNetworkLinkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsForwardingRulesetName", "virtualNetworkLinkName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/dnsForwardingRulesets/dnsForwardingRulesetName/virtualNetworkLinks/virtualNetworkLinkName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseVirtualNetworkLinkID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VirtualNetworkLinkId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/dnsForwardingRulesets",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/dnsForwardingRulesets/dnsForwardingRulesetName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/dnsForwardingRulesets/dnsForwardingRulesetName/virtualNetworkLinks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/dnsForwardingRulesets/dnsForwardingRulesetName/virtualNetworkLinks/virtualNetworkLinkName",
			Expected: &VirtualNetworkLinkId{
				SubscriptionId:           "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:        "example-resource-group",
				DnsForwardingRulesetName: "dnsForwardingRulesetName",
				VirtualNetworkLinkName:   "virtualNetworkLinkName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/dnsForwardingRulesets/dnsForwardingRulesetName/virtualNetworkLinks/virtualNetworkLinkName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVirtualNetworkLinkID(v.Input)
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

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.DnsForwardingRulesetName != v.Expected.DnsForwardingRulesetName {
			t.Fatalf("Expected %q but got %q for DnsForwardingRulesetName", v.Expected.DnsForwardingRulesetName, actual.DnsForwardingRulesetName)
		}

		if actual.VirtualNetworkLinkName != v.Expected.VirtualNetworkLinkName {
			t.Fatalf("Expected %q but got %q for VirtualNetworkLinkName", v.Expected.VirtualNetworkLinkName, actual.VirtualNetworkLinkName)
		}

	}
}

func TestParseVirtualNetworkLinkIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VirtualNetworkLinkId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/dnsForwardingRulesets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/dNsFoRwArDiNgRuLeSeTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/dnsForwardingRulesets/dnsForwardingRulesetName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/dNsFoRwArDiNgRuLeSeTs/dNsFoRwArDiNgRuLeSeTnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/dnsForwardingRulesets/dnsForwardingRulesetName/virtualNetworkLinks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/dNsFoRwArDiNgRuLeSeTs/dNsFoRwArDiNgRuLeSeTnAmE/vIrTuAlNeTwOrKlInKs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/dnsForwardingRulesets/dnsForwardingRulesetName/virtualNetworkLinks/virtualNetworkLinkName",
			Expected: &VirtualNetworkLinkId{
				SubscriptionId:           "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:        "example-resource-group",
				DnsForwardingRulesetName: "dnsForwardingRulesetName",
				VirtualNetworkLinkName:   "virtualNetworkLinkName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/dnsForwardingRulesets/dnsForwardingRulesetName/virtualNetworkLinks/virtualNetworkLinkName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/dNsFoRwArDiNgRuLeSeTs/dNsFoRwArDiNgRuLeSeTnAmE/vIrTuAlNeTwOrKlInKs/vIrTuAlNeTwOrKlInKnAmE",
			Expected: &VirtualNetworkLinkId{
				SubscriptionId:           "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:        "eXaMpLe-rEsOuRcE-GrOuP",
				DnsForwardingRulesetName: "dNsFoRwArDiNgRuLeSeTnAmE",
				VirtualNetworkLinkName:   "vIrTuAlNeTwOrKlInKnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/dNsFoRwArDiNgRuLeSeTs/dNsFoRwArDiNgRuLeSeTnAmE/vIrTuAlNeTwOrKlInKs/vIrTuAlNeTwOrKlInKnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVirtualNetworkLinkIDInsensitively(v.Input)
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

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.DnsForwardingRulesetName != v.Expected.DnsForwardingRulesetName {
			t.Fatalf("Expected %q but got %q for DnsForwardingRulesetName", v.Expected.DnsForwardingRulesetName, actual.DnsForwardingRulesetName)
		}

		if actual.VirtualNetworkLinkName != v.Expected.VirtualNetworkLinkName {
			t.Fatalf("Expected %q but got %q for VirtualNetworkLinkName", v.Expected.VirtualNetworkLinkName, actual.VirtualNetworkLinkName)
		}

	}
}

func TestSegmentsForVirtualNetworkLinkId(t *testing.T) {
	segments := VirtualNetworkLinkId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("VirtualNetworkLinkId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
