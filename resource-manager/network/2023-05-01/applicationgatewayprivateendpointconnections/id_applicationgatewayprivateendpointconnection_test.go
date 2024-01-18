package applicationgatewayprivateendpointconnections

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationGatewayPrivateEndpointConnectionId{}

func TestNewApplicationGatewayPrivateEndpointConnectionID(t *testing.T) {
	id := NewApplicationGatewayPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "applicationGatewayValue", "privateEndpointConnectionValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.ApplicationGatewayName != "applicationGatewayValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationGatewayName'", id.ApplicationGatewayName, "applicationGatewayValue")
	}

	if id.PrivateEndpointConnectionName != "privateEndpointConnectionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PrivateEndpointConnectionName'", id.PrivateEndpointConnectionName, "privateEndpointConnectionValue")
	}
}

func TestFormatApplicationGatewayPrivateEndpointConnectionID(t *testing.T) {
	actual := NewApplicationGatewayPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "applicationGatewayValue", "privateEndpointConnectionValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/applicationGateways/applicationGatewayValue/privateEndpointConnections/privateEndpointConnectionValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseApplicationGatewayPrivateEndpointConnectionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationGatewayPrivateEndpointConnectionId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/applicationGateways",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/applicationGateways/applicationGatewayValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/applicationGateways/applicationGatewayValue/privateEndpointConnections",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/applicationGateways/applicationGatewayValue/privateEndpointConnections/privateEndpointConnectionValue",
			Expected: &ApplicationGatewayPrivateEndpointConnectionId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "example-resource-group",
				ApplicationGatewayName:        "applicationGatewayValue",
				PrivateEndpointConnectionName: "privateEndpointConnectionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/applicationGateways/applicationGatewayValue/privateEndpointConnections/privateEndpointConnectionValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationGatewayPrivateEndpointConnectionID(v.Input)
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

		if actual.ApplicationGatewayName != v.Expected.ApplicationGatewayName {
			t.Fatalf("Expected %q but got %q for ApplicationGatewayName", v.Expected.ApplicationGatewayName, actual.ApplicationGatewayName)
		}

		if actual.PrivateEndpointConnectionName != v.Expected.PrivateEndpointConnectionName {
			t.Fatalf("Expected %q but got %q for PrivateEndpointConnectionName", v.Expected.PrivateEndpointConnectionName, actual.PrivateEndpointConnectionName)
		}

	}
}

func TestParseApplicationGatewayPrivateEndpointConnectionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationGatewayPrivateEndpointConnectionId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/applicationGateways",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/aPpLiCaTiOnGaTeWaYs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/applicationGateways/applicationGatewayValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/aPpLiCaTiOnGaTeWaYs/aPpLiCaTiOnGaTeWaYvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/applicationGateways/applicationGatewayValue/privateEndpointConnections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/aPpLiCaTiOnGaTeWaYs/aPpLiCaTiOnGaTeWaYvAlUe/pRiVaTeEnDpOiNtCoNnEcTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/applicationGateways/applicationGatewayValue/privateEndpointConnections/privateEndpointConnectionValue",
			Expected: &ApplicationGatewayPrivateEndpointConnectionId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "example-resource-group",
				ApplicationGatewayName:        "applicationGatewayValue",
				PrivateEndpointConnectionName: "privateEndpointConnectionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/applicationGateways/applicationGatewayValue/privateEndpointConnections/privateEndpointConnectionValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/aPpLiCaTiOnGaTeWaYs/aPpLiCaTiOnGaTeWaYvAlUe/pRiVaTeEnDpOiNtCoNnEcTiOnS/pRiVaTeEnDpOiNtCoNnEcTiOnVaLuE",
			Expected: &ApplicationGatewayPrivateEndpointConnectionId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "eXaMpLe-rEsOuRcE-GrOuP",
				ApplicationGatewayName:        "aPpLiCaTiOnGaTeWaYvAlUe",
				PrivateEndpointConnectionName: "pRiVaTeEnDpOiNtCoNnEcTiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/aPpLiCaTiOnGaTeWaYs/aPpLiCaTiOnGaTeWaYvAlUe/pRiVaTeEnDpOiNtCoNnEcTiOnS/pRiVaTeEnDpOiNtCoNnEcTiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationGatewayPrivateEndpointConnectionIDInsensitively(v.Input)
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

		if actual.ApplicationGatewayName != v.Expected.ApplicationGatewayName {
			t.Fatalf("Expected %q but got %q for ApplicationGatewayName", v.Expected.ApplicationGatewayName, actual.ApplicationGatewayName)
		}

		if actual.PrivateEndpointConnectionName != v.Expected.PrivateEndpointConnectionName {
			t.Fatalf("Expected %q but got %q for PrivateEndpointConnectionName", v.Expected.PrivateEndpointConnectionName, actual.PrivateEndpointConnectionName)
		}

	}
}

func TestSegmentsForApplicationGatewayPrivateEndpointConnectionId(t *testing.T) {
	segments := ApplicationGatewayPrivateEndpointConnectionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ApplicationGatewayPrivateEndpointConnectionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
