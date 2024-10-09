package appserviceplans

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &VirtualNetworkConnectionGatewayId{}

func TestNewVirtualNetworkConnectionGatewayID(t *testing.T) {
	id := NewVirtualNetworkConnectionGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName", "virtualNetworkConnectionName", "gatewayName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.ServerFarmName != "serverFarmName" {
		t.Fatalf("Expected %q but got %q for Segment 'ServerFarmName'", id.ServerFarmName, "serverFarmName")
	}

	if id.VirtualNetworkConnectionName != "virtualNetworkConnectionName" {
		t.Fatalf("Expected %q but got %q for Segment 'VirtualNetworkConnectionName'", id.VirtualNetworkConnectionName, "virtualNetworkConnectionName")
	}

	if id.GatewayName != "gatewayName" {
		t.Fatalf("Expected %q but got %q for Segment 'GatewayName'", id.GatewayName, "gatewayName")
	}
}

func TestFormatVirtualNetworkConnectionGatewayID(t *testing.T) {
	actual := NewVirtualNetworkConnectionGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName", "virtualNetworkConnectionName", "gatewayName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/serverFarms/serverFarmName/virtualNetworkConnections/virtualNetworkConnectionName/gateways/gatewayName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseVirtualNetworkConnectionGatewayID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VirtualNetworkConnectionGatewayId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/serverFarms",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/serverFarms/serverFarmName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/serverFarms/serverFarmName/virtualNetworkConnections",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/serverFarms/serverFarmName/virtualNetworkConnections/virtualNetworkConnectionName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/serverFarms/serverFarmName/virtualNetworkConnections/virtualNetworkConnectionName/gateways",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/serverFarms/serverFarmName/virtualNetworkConnections/virtualNetworkConnectionName/gateways/gatewayName",
			Expected: &VirtualNetworkConnectionGatewayId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:            "example-resource-group",
				ServerFarmName:               "serverFarmName",
				VirtualNetworkConnectionName: "virtualNetworkConnectionName",
				GatewayName:                  "gatewayName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/serverFarms/serverFarmName/virtualNetworkConnections/virtualNetworkConnectionName/gateways/gatewayName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVirtualNetworkConnectionGatewayID(v.Input)
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

		if actual.ServerFarmName != v.Expected.ServerFarmName {
			t.Fatalf("Expected %q but got %q for ServerFarmName", v.Expected.ServerFarmName, actual.ServerFarmName)
		}

		if actual.VirtualNetworkConnectionName != v.Expected.VirtualNetworkConnectionName {
			t.Fatalf("Expected %q but got %q for VirtualNetworkConnectionName", v.Expected.VirtualNetworkConnectionName, actual.VirtualNetworkConnectionName)
		}

		if actual.GatewayName != v.Expected.GatewayName {
			t.Fatalf("Expected %q but got %q for GatewayName", v.Expected.GatewayName, actual.GatewayName)
		}

	}
}

func TestParseVirtualNetworkConnectionGatewayIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VirtualNetworkConnectionGatewayId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/serverFarms",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/sErVeRfArMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/serverFarms/serverFarmName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/sErVeRfArMs/sErVeRfArMnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/serverFarms/serverFarmName/virtualNetworkConnections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/sErVeRfArMs/sErVeRfArMnAmE/vIrTuAlNeTwOrKcOnNeCtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/serverFarms/serverFarmName/virtualNetworkConnections/virtualNetworkConnectionName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/sErVeRfArMs/sErVeRfArMnAmE/vIrTuAlNeTwOrKcOnNeCtIoNs/vIrTuAlNeTwOrKcOnNeCtIoNnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/serverFarms/serverFarmName/virtualNetworkConnections/virtualNetworkConnectionName/gateways",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/sErVeRfArMs/sErVeRfArMnAmE/vIrTuAlNeTwOrKcOnNeCtIoNs/vIrTuAlNeTwOrKcOnNeCtIoNnAmE/gAtEwAyS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/serverFarms/serverFarmName/virtualNetworkConnections/virtualNetworkConnectionName/gateways/gatewayName",
			Expected: &VirtualNetworkConnectionGatewayId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:            "example-resource-group",
				ServerFarmName:               "serverFarmName",
				VirtualNetworkConnectionName: "virtualNetworkConnectionName",
				GatewayName:                  "gatewayName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/serverFarms/serverFarmName/virtualNetworkConnections/virtualNetworkConnectionName/gateways/gatewayName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/sErVeRfArMs/sErVeRfArMnAmE/vIrTuAlNeTwOrKcOnNeCtIoNs/vIrTuAlNeTwOrKcOnNeCtIoNnAmE/gAtEwAyS/gAtEwAyNaMe",
			Expected: &VirtualNetworkConnectionGatewayId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:            "eXaMpLe-rEsOuRcE-GrOuP",
				ServerFarmName:               "sErVeRfArMnAmE",
				VirtualNetworkConnectionName: "vIrTuAlNeTwOrKcOnNeCtIoNnAmE",
				GatewayName:                  "gAtEwAyNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/sErVeRfArMs/sErVeRfArMnAmE/vIrTuAlNeTwOrKcOnNeCtIoNs/vIrTuAlNeTwOrKcOnNeCtIoNnAmE/gAtEwAyS/gAtEwAyNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVirtualNetworkConnectionGatewayIDInsensitively(v.Input)
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

		if actual.ServerFarmName != v.Expected.ServerFarmName {
			t.Fatalf("Expected %q but got %q for ServerFarmName", v.Expected.ServerFarmName, actual.ServerFarmName)
		}

		if actual.VirtualNetworkConnectionName != v.Expected.VirtualNetworkConnectionName {
			t.Fatalf("Expected %q but got %q for VirtualNetworkConnectionName", v.Expected.VirtualNetworkConnectionName, actual.VirtualNetworkConnectionName)
		}

		if actual.GatewayName != v.Expected.GatewayName {
			t.Fatalf("Expected %q but got %q for GatewayName", v.Expected.GatewayName, actual.GatewayName)
		}

	}
}

func TestSegmentsForVirtualNetworkConnectionGatewayId(t *testing.T) {
	segments := VirtualNetworkConnectionGatewayId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("VirtualNetworkConnectionGatewayId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
