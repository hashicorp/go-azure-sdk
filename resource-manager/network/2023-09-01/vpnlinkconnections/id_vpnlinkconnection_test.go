package vpnlinkconnections

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = VpnLinkConnectionId{}

func TestNewVpnLinkConnectionID(t *testing.T) {
	id := NewVpnLinkConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vpnGatewayValue", "vpnConnectionValue", "vpnLinkConnectionValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.VpnGatewayName != "vpnGatewayValue" {
		t.Fatalf("Expected %q but got %q for Segment 'VpnGatewayName'", id.VpnGatewayName, "vpnGatewayValue")
	}

	if id.VpnConnectionName != "vpnConnectionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'VpnConnectionName'", id.VpnConnectionName, "vpnConnectionValue")
	}

	if id.VpnLinkConnectionName != "vpnLinkConnectionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'VpnLinkConnectionName'", id.VpnLinkConnectionName, "vpnLinkConnectionValue")
	}
}

func TestFormatVpnLinkConnectionID(t *testing.T) {
	actual := NewVpnLinkConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vpnGatewayValue", "vpnConnectionValue", "vpnLinkConnectionValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnGateways/vpnGatewayValue/vpnConnections/vpnConnectionValue/vpnLinkConnections/vpnLinkConnectionValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseVpnLinkConnectionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VpnLinkConnectionId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnGateways",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnGateways/vpnGatewayValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnGateways/vpnGatewayValue/vpnConnections",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnGateways/vpnGatewayValue/vpnConnections/vpnConnectionValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnGateways/vpnGatewayValue/vpnConnections/vpnConnectionValue/vpnLinkConnections",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnGateways/vpnGatewayValue/vpnConnections/vpnConnectionValue/vpnLinkConnections/vpnLinkConnectionValue",
			Expected: &VpnLinkConnectionId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "example-resource-group",
				VpnGatewayName:        "vpnGatewayValue",
				VpnConnectionName:     "vpnConnectionValue",
				VpnLinkConnectionName: "vpnLinkConnectionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnGateways/vpnGatewayValue/vpnConnections/vpnConnectionValue/vpnLinkConnections/vpnLinkConnectionValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVpnLinkConnectionID(v.Input)
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

		if actual.VpnGatewayName != v.Expected.VpnGatewayName {
			t.Fatalf("Expected %q but got %q for VpnGatewayName", v.Expected.VpnGatewayName, actual.VpnGatewayName)
		}

		if actual.VpnConnectionName != v.Expected.VpnConnectionName {
			t.Fatalf("Expected %q but got %q for VpnConnectionName", v.Expected.VpnConnectionName, actual.VpnConnectionName)
		}

		if actual.VpnLinkConnectionName != v.Expected.VpnLinkConnectionName {
			t.Fatalf("Expected %q but got %q for VpnLinkConnectionName", v.Expected.VpnLinkConnectionName, actual.VpnLinkConnectionName)
		}

	}
}

func TestParseVpnLinkConnectionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VpnLinkConnectionId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnGateways",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/vPnGaTeWaYs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnGateways/vpnGatewayValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/vPnGaTeWaYs/vPnGaTeWaYvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnGateways/vpnGatewayValue/vpnConnections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/vPnGaTeWaYs/vPnGaTeWaYvAlUe/vPnCoNnEcTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnGateways/vpnGatewayValue/vpnConnections/vpnConnectionValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/vPnGaTeWaYs/vPnGaTeWaYvAlUe/vPnCoNnEcTiOnS/vPnCoNnEcTiOnVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnGateways/vpnGatewayValue/vpnConnections/vpnConnectionValue/vpnLinkConnections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/vPnGaTeWaYs/vPnGaTeWaYvAlUe/vPnCoNnEcTiOnS/vPnCoNnEcTiOnVaLuE/vPnLiNkCoNnEcTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnGateways/vpnGatewayValue/vpnConnections/vpnConnectionValue/vpnLinkConnections/vpnLinkConnectionValue",
			Expected: &VpnLinkConnectionId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "example-resource-group",
				VpnGatewayName:        "vpnGatewayValue",
				VpnConnectionName:     "vpnConnectionValue",
				VpnLinkConnectionName: "vpnLinkConnectionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnGateways/vpnGatewayValue/vpnConnections/vpnConnectionValue/vpnLinkConnections/vpnLinkConnectionValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/vPnGaTeWaYs/vPnGaTeWaYvAlUe/vPnCoNnEcTiOnS/vPnCoNnEcTiOnVaLuE/vPnLiNkCoNnEcTiOnS/vPnLiNkCoNnEcTiOnVaLuE",
			Expected: &VpnLinkConnectionId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "eXaMpLe-rEsOuRcE-GrOuP",
				VpnGatewayName:        "vPnGaTeWaYvAlUe",
				VpnConnectionName:     "vPnCoNnEcTiOnVaLuE",
				VpnLinkConnectionName: "vPnLiNkCoNnEcTiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/vPnGaTeWaYs/vPnGaTeWaYvAlUe/vPnCoNnEcTiOnS/vPnCoNnEcTiOnVaLuE/vPnLiNkCoNnEcTiOnS/vPnLiNkCoNnEcTiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVpnLinkConnectionIDInsensitively(v.Input)
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

		if actual.VpnGatewayName != v.Expected.VpnGatewayName {
			t.Fatalf("Expected %q but got %q for VpnGatewayName", v.Expected.VpnGatewayName, actual.VpnGatewayName)
		}

		if actual.VpnConnectionName != v.Expected.VpnConnectionName {
			t.Fatalf("Expected %q but got %q for VpnConnectionName", v.Expected.VpnConnectionName, actual.VpnConnectionName)
		}

		if actual.VpnLinkConnectionName != v.Expected.VpnLinkConnectionName {
			t.Fatalf("Expected %q but got %q for VpnLinkConnectionName", v.Expected.VpnLinkConnectionName, actual.VpnLinkConnectionName)
		}

	}
}

func TestSegmentsForVpnLinkConnectionId(t *testing.T) {
	segments := VpnLinkConnectionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("VpnLinkConnectionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
