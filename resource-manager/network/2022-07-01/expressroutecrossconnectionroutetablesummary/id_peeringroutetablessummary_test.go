package expressroutecrossconnectionroutetablesummary

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PeeringRouteTablesSummaryId{}

func TestNewPeeringRouteTablesSummaryID(t *testing.T) {
	id := NewPeeringRouteTablesSummaryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "expressRouteCrossConnectionValue", "peeringValue", "routeTablesSummaryValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.ExpressRouteCrossConnectionName != "expressRouteCrossConnectionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ExpressRouteCrossConnectionName'", id.ExpressRouteCrossConnectionName, "expressRouteCrossConnectionValue")
	}

	if id.PeeringName != "peeringValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PeeringName'", id.PeeringName, "peeringValue")
	}

	if id.RouteTablesSummaryName != "routeTablesSummaryValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RouteTablesSummaryName'", id.RouteTablesSummaryName, "routeTablesSummaryValue")
	}
}

func TestFormatPeeringRouteTablesSummaryID(t *testing.T) {
	actual := NewPeeringRouteTablesSummaryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "expressRouteCrossConnectionValue", "peeringValue", "routeTablesSummaryValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/expressRouteCrossConnections/expressRouteCrossConnectionValue/peerings/peeringValue/routeTablesSummary/routeTablesSummaryValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePeeringRouteTablesSummaryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PeeringRouteTablesSummaryId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/expressRouteCrossConnections",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/expressRouteCrossConnections/expressRouteCrossConnectionValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/expressRouteCrossConnections/expressRouteCrossConnectionValue/peerings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/expressRouteCrossConnections/expressRouteCrossConnectionValue/peerings/peeringValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/expressRouteCrossConnections/expressRouteCrossConnectionValue/peerings/peeringValue/routeTablesSummary",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/expressRouteCrossConnections/expressRouteCrossConnectionValue/peerings/peeringValue/routeTablesSummary/routeTablesSummaryValue",
			Expected: &PeeringRouteTablesSummaryId{
				SubscriptionId:                  "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:               "example-resource-group",
				ExpressRouteCrossConnectionName: "expressRouteCrossConnectionValue",
				PeeringName:                     "peeringValue",
				RouteTablesSummaryName:          "routeTablesSummaryValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/expressRouteCrossConnections/expressRouteCrossConnectionValue/peerings/peeringValue/routeTablesSummary/routeTablesSummaryValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePeeringRouteTablesSummaryID(v.Input)
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

		if actual.ExpressRouteCrossConnectionName != v.Expected.ExpressRouteCrossConnectionName {
			t.Fatalf("Expected %q but got %q for ExpressRouteCrossConnectionName", v.Expected.ExpressRouteCrossConnectionName, actual.ExpressRouteCrossConnectionName)
		}

		if actual.PeeringName != v.Expected.PeeringName {
			t.Fatalf("Expected %q but got %q for PeeringName", v.Expected.PeeringName, actual.PeeringName)
		}

		if actual.RouteTablesSummaryName != v.Expected.RouteTablesSummaryName {
			t.Fatalf("Expected %q but got %q for RouteTablesSummaryName", v.Expected.RouteTablesSummaryName, actual.RouteTablesSummaryName)
		}

	}
}

func TestParsePeeringRouteTablesSummaryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PeeringRouteTablesSummaryId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/expressRouteCrossConnections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/eXpReSsRoUtEcRoSsCoNnEcTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/expressRouteCrossConnections/expressRouteCrossConnectionValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/eXpReSsRoUtEcRoSsCoNnEcTiOnS/eXpReSsRoUtEcRoSsCoNnEcTiOnVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/expressRouteCrossConnections/expressRouteCrossConnectionValue/peerings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/eXpReSsRoUtEcRoSsCoNnEcTiOnS/eXpReSsRoUtEcRoSsCoNnEcTiOnVaLuE/pEeRiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/expressRouteCrossConnections/expressRouteCrossConnectionValue/peerings/peeringValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/eXpReSsRoUtEcRoSsCoNnEcTiOnS/eXpReSsRoUtEcRoSsCoNnEcTiOnVaLuE/pEeRiNgS/pEeRiNgVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/expressRouteCrossConnections/expressRouteCrossConnectionValue/peerings/peeringValue/routeTablesSummary",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/eXpReSsRoUtEcRoSsCoNnEcTiOnS/eXpReSsRoUtEcRoSsCoNnEcTiOnVaLuE/pEeRiNgS/pEeRiNgVaLuE/rOuTeTaBlEsSuMmArY",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/expressRouteCrossConnections/expressRouteCrossConnectionValue/peerings/peeringValue/routeTablesSummary/routeTablesSummaryValue",
			Expected: &PeeringRouteTablesSummaryId{
				SubscriptionId:                  "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:               "example-resource-group",
				ExpressRouteCrossConnectionName: "expressRouteCrossConnectionValue",
				PeeringName:                     "peeringValue",
				RouteTablesSummaryName:          "routeTablesSummaryValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/expressRouteCrossConnections/expressRouteCrossConnectionValue/peerings/peeringValue/routeTablesSummary/routeTablesSummaryValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/eXpReSsRoUtEcRoSsCoNnEcTiOnS/eXpReSsRoUtEcRoSsCoNnEcTiOnVaLuE/pEeRiNgS/pEeRiNgVaLuE/rOuTeTaBlEsSuMmArY/rOuTeTaBlEsSuMmArYvAlUe",
			Expected: &PeeringRouteTablesSummaryId{
				SubscriptionId:                  "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:               "eXaMpLe-rEsOuRcE-GrOuP",
				ExpressRouteCrossConnectionName: "eXpReSsRoUtEcRoSsCoNnEcTiOnVaLuE",
				PeeringName:                     "pEeRiNgVaLuE",
				RouteTablesSummaryName:          "rOuTeTaBlEsSuMmArYvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/eXpReSsRoUtEcRoSsCoNnEcTiOnS/eXpReSsRoUtEcRoSsCoNnEcTiOnVaLuE/pEeRiNgS/pEeRiNgVaLuE/rOuTeTaBlEsSuMmArY/rOuTeTaBlEsSuMmArYvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePeeringRouteTablesSummaryIDInsensitively(v.Input)
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

		if actual.ExpressRouteCrossConnectionName != v.Expected.ExpressRouteCrossConnectionName {
			t.Fatalf("Expected %q but got %q for ExpressRouteCrossConnectionName", v.Expected.ExpressRouteCrossConnectionName, actual.ExpressRouteCrossConnectionName)
		}

		if actual.PeeringName != v.Expected.PeeringName {
			t.Fatalf("Expected %q but got %q for PeeringName", v.Expected.PeeringName, actual.PeeringName)
		}

		if actual.RouteTablesSummaryName != v.Expected.RouteTablesSummaryName {
			t.Fatalf("Expected %q but got %q for RouteTablesSummaryName", v.Expected.RouteTablesSummaryName, actual.RouteTablesSummaryName)
		}

	}
}

func TestSegmentsForPeeringRouteTablesSummaryId(t *testing.T) {
	segments := PeeringRouteTablesSummaryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PeeringRouteTablesSummaryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
