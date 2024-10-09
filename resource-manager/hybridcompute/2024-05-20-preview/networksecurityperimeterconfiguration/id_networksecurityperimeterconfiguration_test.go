package networksecurityperimeterconfiguration

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &NetworkSecurityPerimeterConfigurationId{}

func TestNewNetworkSecurityPerimeterConfigurationID(t *testing.T) {
	id := NewNetworkSecurityPerimeterConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateLinkScopeName", "networkSecurityPerimeterConfigurationName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.PrivateLinkScopeName != "privateLinkScopeName" {
		t.Fatalf("Expected %q but got %q for Segment 'PrivateLinkScopeName'", id.PrivateLinkScopeName, "privateLinkScopeName")
	}

	if id.NetworkSecurityPerimeterConfigurationName != "networkSecurityPerimeterConfigurationName" {
		t.Fatalf("Expected %q but got %q for Segment 'NetworkSecurityPerimeterConfigurationName'", id.NetworkSecurityPerimeterConfigurationName, "networkSecurityPerimeterConfigurationName")
	}
}

func TestFormatNetworkSecurityPerimeterConfigurationID(t *testing.T) {
	actual := NewNetworkSecurityPerimeterConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateLinkScopeName", "networkSecurityPerimeterConfigurationName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName/networkSecurityPerimeterConfigurations/networkSecurityPerimeterConfigurationName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseNetworkSecurityPerimeterConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *NetworkSecurityPerimeterConfigurationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HybridCompute",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HybridCompute/privateLinkScopes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName/networkSecurityPerimeterConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName/networkSecurityPerimeterConfigurations/networkSecurityPerimeterConfigurationName",
			Expected: &NetworkSecurityPerimeterConfigurationId{
				SubscriptionId:                            "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                         "example-resource-group",
				PrivateLinkScopeName:                      "privateLinkScopeName",
				NetworkSecurityPerimeterConfigurationName: "networkSecurityPerimeterConfigurationName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName/networkSecurityPerimeterConfigurations/networkSecurityPerimeterConfigurationName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseNetworkSecurityPerimeterConfigurationID(v.Input)
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

		if actual.PrivateLinkScopeName != v.Expected.PrivateLinkScopeName {
			t.Fatalf("Expected %q but got %q for PrivateLinkScopeName", v.Expected.PrivateLinkScopeName, actual.PrivateLinkScopeName)
		}

		if actual.NetworkSecurityPerimeterConfigurationName != v.Expected.NetworkSecurityPerimeterConfigurationName {
			t.Fatalf("Expected %q but got %q for NetworkSecurityPerimeterConfigurationName", v.Expected.NetworkSecurityPerimeterConfigurationName, actual.NetworkSecurityPerimeterConfigurationName)
		}

	}
}

func TestParseNetworkSecurityPerimeterConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *NetworkSecurityPerimeterConfigurationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HybridCompute",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HybridCompute/privateLinkScopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/pRiVaTeLiNkScOpEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/pRiVaTeLiNkScOpEs/pRiVaTeLiNkScOpEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName/networkSecurityPerimeterConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/pRiVaTeLiNkScOpEs/pRiVaTeLiNkScOpEnAmE/nEtWoRkSeCuRiTyPeRiMeTeRcOnFiGuRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName/networkSecurityPerimeterConfigurations/networkSecurityPerimeterConfigurationName",
			Expected: &NetworkSecurityPerimeterConfigurationId{
				SubscriptionId:                            "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                         "example-resource-group",
				PrivateLinkScopeName:                      "privateLinkScopeName",
				NetworkSecurityPerimeterConfigurationName: "networkSecurityPerimeterConfigurationName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName/networkSecurityPerimeterConfigurations/networkSecurityPerimeterConfigurationName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/pRiVaTeLiNkScOpEs/pRiVaTeLiNkScOpEnAmE/nEtWoRkSeCuRiTyPeRiMeTeRcOnFiGuRaTiOnS/nEtWoRkSeCuRiTyPeRiMeTeRcOnFiGuRaTiOnNaMe",
			Expected: &NetworkSecurityPerimeterConfigurationId{
				SubscriptionId:                            "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                         "eXaMpLe-rEsOuRcE-GrOuP",
				PrivateLinkScopeName:                      "pRiVaTeLiNkScOpEnAmE",
				NetworkSecurityPerimeterConfigurationName: "nEtWoRkSeCuRiTyPeRiMeTeRcOnFiGuRaTiOnNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/pRiVaTeLiNkScOpEs/pRiVaTeLiNkScOpEnAmE/nEtWoRkSeCuRiTyPeRiMeTeRcOnFiGuRaTiOnS/nEtWoRkSeCuRiTyPeRiMeTeRcOnFiGuRaTiOnNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseNetworkSecurityPerimeterConfigurationIDInsensitively(v.Input)
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

		if actual.PrivateLinkScopeName != v.Expected.PrivateLinkScopeName {
			t.Fatalf("Expected %q but got %q for PrivateLinkScopeName", v.Expected.PrivateLinkScopeName, actual.PrivateLinkScopeName)
		}

		if actual.NetworkSecurityPerimeterConfigurationName != v.Expected.NetworkSecurityPerimeterConfigurationName {
			t.Fatalf("Expected %q but got %q for NetworkSecurityPerimeterConfigurationName", v.Expected.NetworkSecurityPerimeterConfigurationName, actual.NetworkSecurityPerimeterConfigurationName)
		}

	}
}

func TestSegmentsForNetworkSecurityPerimeterConfigurationId(t *testing.T) {
	segments := NetworkSecurityPerimeterConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("NetworkSecurityPerimeterConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
