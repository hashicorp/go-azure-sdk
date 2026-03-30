package adminrulecollections

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SecurityAdminConfigurationRuleCollectionId{}

func TestNewSecurityAdminConfigurationRuleCollectionID(t *testing.T) {
	id := NewSecurityAdminConfigurationRuleCollectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "networkManagerName", "securityAdminConfigurationName", "ruleCollectionName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.NetworkManagerName != "networkManagerName" {
		t.Fatalf("Expected %q but got %q for Segment 'NetworkManagerName'", id.NetworkManagerName, "networkManagerName")
	}

	if id.SecurityAdminConfigurationName != "securityAdminConfigurationName" {
		t.Fatalf("Expected %q but got %q for Segment 'SecurityAdminConfigurationName'", id.SecurityAdminConfigurationName, "securityAdminConfigurationName")
	}

	if id.RuleCollectionName != "ruleCollectionName" {
		t.Fatalf("Expected %q but got %q for Segment 'RuleCollectionName'", id.RuleCollectionName, "ruleCollectionName")
	}
}

func TestFormatSecurityAdminConfigurationRuleCollectionID(t *testing.T) {
	actual := NewSecurityAdminConfigurationRuleCollectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "networkManagerName", "securityAdminConfigurationName", "ruleCollectionName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/securityAdminConfigurationName/ruleCollections/ruleCollectionName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSecurityAdminConfigurationRuleCollectionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SecurityAdminConfigurationRuleCollectionId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/securityAdminConfigurationName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/securityAdminConfigurationName/ruleCollections",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/securityAdminConfigurationName/ruleCollections/ruleCollectionName",
			Expected: &SecurityAdminConfigurationRuleCollectionId{
				SubscriptionId:                 "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:              "example-resource-group",
				NetworkManagerName:             "networkManagerName",
				SecurityAdminConfigurationName: "securityAdminConfigurationName",
				RuleCollectionName:             "ruleCollectionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/securityAdminConfigurationName/ruleCollections/ruleCollectionName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSecurityAdminConfigurationRuleCollectionID(v.Input)
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

		if actual.NetworkManagerName != v.Expected.NetworkManagerName {
			t.Fatalf("Expected %q but got %q for NetworkManagerName", v.Expected.NetworkManagerName, actual.NetworkManagerName)
		}

		if actual.SecurityAdminConfigurationName != v.Expected.SecurityAdminConfigurationName {
			t.Fatalf("Expected %q but got %q for SecurityAdminConfigurationName", v.Expected.SecurityAdminConfigurationName, actual.SecurityAdminConfigurationName)
		}

		if actual.RuleCollectionName != v.Expected.RuleCollectionName {
			t.Fatalf("Expected %q but got %q for RuleCollectionName", v.Expected.RuleCollectionName, actual.RuleCollectionName)
		}

	}
}

func TestParseSecurityAdminConfigurationRuleCollectionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SecurityAdminConfigurationRuleCollectionId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs/nEtWoRkMaNaGeRnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs/nEtWoRkMaNaGeRnAmE/sEcUrItYaDmInCoNfIgUrAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/securityAdminConfigurationName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs/nEtWoRkMaNaGeRnAmE/sEcUrItYaDmInCoNfIgUrAtIoNs/sEcUrItYaDmInCoNfIgUrAtIoNnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/securityAdminConfigurationName/ruleCollections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs/nEtWoRkMaNaGeRnAmE/sEcUrItYaDmInCoNfIgUrAtIoNs/sEcUrItYaDmInCoNfIgUrAtIoNnAmE/rUlEcOlLeCtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/securityAdminConfigurationName/ruleCollections/ruleCollectionName",
			Expected: &SecurityAdminConfigurationRuleCollectionId{
				SubscriptionId:                 "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:              "example-resource-group",
				NetworkManagerName:             "networkManagerName",
				SecurityAdminConfigurationName: "securityAdminConfigurationName",
				RuleCollectionName:             "ruleCollectionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/securityAdminConfigurationName/ruleCollections/ruleCollectionName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs/nEtWoRkMaNaGeRnAmE/sEcUrItYaDmInCoNfIgUrAtIoNs/sEcUrItYaDmInCoNfIgUrAtIoNnAmE/rUlEcOlLeCtIoNs/rUlEcOlLeCtIoNnAmE",
			Expected: &SecurityAdminConfigurationRuleCollectionId{
				SubscriptionId:                 "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:              "eXaMpLe-rEsOuRcE-GrOuP",
				NetworkManagerName:             "nEtWoRkMaNaGeRnAmE",
				SecurityAdminConfigurationName: "sEcUrItYaDmInCoNfIgUrAtIoNnAmE",
				RuleCollectionName:             "rUlEcOlLeCtIoNnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs/nEtWoRkMaNaGeRnAmE/sEcUrItYaDmInCoNfIgUrAtIoNs/sEcUrItYaDmInCoNfIgUrAtIoNnAmE/rUlEcOlLeCtIoNs/rUlEcOlLeCtIoNnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSecurityAdminConfigurationRuleCollectionIDInsensitively(v.Input)
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

		if actual.NetworkManagerName != v.Expected.NetworkManagerName {
			t.Fatalf("Expected %q but got %q for NetworkManagerName", v.Expected.NetworkManagerName, actual.NetworkManagerName)
		}

		if actual.SecurityAdminConfigurationName != v.Expected.SecurityAdminConfigurationName {
			t.Fatalf("Expected %q but got %q for SecurityAdminConfigurationName", v.Expected.SecurityAdminConfigurationName, actual.SecurityAdminConfigurationName)
		}

		if actual.RuleCollectionName != v.Expected.RuleCollectionName {
			t.Fatalf("Expected %q but got %q for RuleCollectionName", v.Expected.RuleCollectionName, actual.RuleCollectionName)
		}

	}
}

func TestSegmentsForSecurityAdminConfigurationRuleCollectionId(t *testing.T) {
	segments := SecurityAdminConfigurationRuleCollectionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SecurityAdminConfigurationRuleCollectionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
