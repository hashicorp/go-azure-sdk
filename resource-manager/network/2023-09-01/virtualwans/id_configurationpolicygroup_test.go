package virtualwans

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ConfigurationPolicyGroupId{}

func TestNewConfigurationPolicyGroupID(t *testing.T) {
	id := NewConfigurationPolicyGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vpnServerConfigurationName", "configurationPolicyGroupName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.VpnServerConfigurationName != "vpnServerConfigurationName" {
		t.Fatalf("Expected %q but got %q for Segment 'VpnServerConfigurationName'", id.VpnServerConfigurationName, "vpnServerConfigurationName")
	}

	if id.ConfigurationPolicyGroupName != "configurationPolicyGroupName" {
		t.Fatalf("Expected %q but got %q for Segment 'ConfigurationPolicyGroupName'", id.ConfigurationPolicyGroupName, "configurationPolicyGroupName")
	}
}

func TestFormatConfigurationPolicyGroupID(t *testing.T) {
	actual := NewConfigurationPolicyGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vpnServerConfigurationName", "configurationPolicyGroupName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfigurationName/configurationPolicyGroups/configurationPolicyGroupName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseConfigurationPolicyGroupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ConfigurationPolicyGroupId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnServerConfigurations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfigurationName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfigurationName/configurationPolicyGroups",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfigurationName/configurationPolicyGroups/configurationPolicyGroupName",
			Expected: &ConfigurationPolicyGroupId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:            "example-resource-group",
				VpnServerConfigurationName:   "vpnServerConfigurationName",
				ConfigurationPolicyGroupName: "configurationPolicyGroupName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfigurationName/configurationPolicyGroups/configurationPolicyGroupName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseConfigurationPolicyGroupID(v.Input)
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

		if actual.VpnServerConfigurationName != v.Expected.VpnServerConfigurationName {
			t.Fatalf("Expected %q but got %q for VpnServerConfigurationName", v.Expected.VpnServerConfigurationName, actual.VpnServerConfigurationName)
		}

		if actual.ConfigurationPolicyGroupName != v.Expected.ConfigurationPolicyGroupName {
			t.Fatalf("Expected %q but got %q for ConfigurationPolicyGroupName", v.Expected.ConfigurationPolicyGroupName, actual.ConfigurationPolicyGroupName)
		}

	}
}

func TestParseConfigurationPolicyGroupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ConfigurationPolicyGroupId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnServerConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/vPnSeRvErCoNfIgUrAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfigurationName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/vPnSeRvErCoNfIgUrAtIoNs/vPnSeRvErCoNfIgUrAtIoNnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfigurationName/configurationPolicyGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/vPnSeRvErCoNfIgUrAtIoNs/vPnSeRvErCoNfIgUrAtIoNnAmE/cOnFiGuRaTiOnPoLiCyGrOuPs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfigurationName/configurationPolicyGroups/configurationPolicyGroupName",
			Expected: &ConfigurationPolicyGroupId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:            "example-resource-group",
				VpnServerConfigurationName:   "vpnServerConfigurationName",
				ConfigurationPolicyGroupName: "configurationPolicyGroupName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfigurationName/configurationPolicyGroups/configurationPolicyGroupName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/vPnSeRvErCoNfIgUrAtIoNs/vPnSeRvErCoNfIgUrAtIoNnAmE/cOnFiGuRaTiOnPoLiCyGrOuPs/cOnFiGuRaTiOnPoLiCyGrOuPnAmE",
			Expected: &ConfigurationPolicyGroupId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:            "eXaMpLe-rEsOuRcE-GrOuP",
				VpnServerConfigurationName:   "vPnSeRvErCoNfIgUrAtIoNnAmE",
				ConfigurationPolicyGroupName: "cOnFiGuRaTiOnPoLiCyGrOuPnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/vPnSeRvErCoNfIgUrAtIoNs/vPnSeRvErCoNfIgUrAtIoNnAmE/cOnFiGuRaTiOnPoLiCyGrOuPs/cOnFiGuRaTiOnPoLiCyGrOuPnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseConfigurationPolicyGroupIDInsensitively(v.Input)
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

		if actual.VpnServerConfigurationName != v.Expected.VpnServerConfigurationName {
			t.Fatalf("Expected %q but got %q for VpnServerConfigurationName", v.Expected.VpnServerConfigurationName, actual.VpnServerConfigurationName)
		}

		if actual.ConfigurationPolicyGroupName != v.Expected.ConfigurationPolicyGroupName {
			t.Fatalf("Expected %q but got %q for ConfigurationPolicyGroupName", v.Expected.ConfigurationPolicyGroupName, actual.ConfigurationPolicyGroupName)
		}

	}
}

func TestSegmentsForConfigurationPolicyGroupId(t *testing.T) {
	segments := ConfigurationPolicyGroupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ConfigurationPolicyGroupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
