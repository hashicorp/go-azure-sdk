package adminrules

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SecurityAdminConfigurationRuleCollectionRuleId{}

func TestNewSecurityAdminConfigurationRuleCollectionRuleID(t *testing.T) {
	id := NewSecurityAdminConfigurationRuleCollectionRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "networkManagerName", "configurationName", "ruleCollectionName", "ruleName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.NetworkManagerName != "networkManagerName" {
		t.Fatalf("Expected %q but got %q for Segment 'NetworkManagerName'", id.NetworkManagerName, "networkManagerName")
	}

	if id.SecurityAdminConfigurationName != "configurationName" {
		t.Fatalf("Expected %q but got %q for Segment 'SecurityAdminConfigurationName'", id.SecurityAdminConfigurationName, "configurationName")
	}

	if id.RuleCollectionName != "ruleCollectionName" {
		t.Fatalf("Expected %q but got %q for Segment 'RuleCollectionName'", id.RuleCollectionName, "ruleCollectionName")
	}

	if id.RuleName != "ruleName" {
		t.Fatalf("Expected %q but got %q for Segment 'RuleName'", id.RuleName, "ruleName")
	}
}

func TestFormatSecurityAdminConfigurationRuleCollectionRuleID(t *testing.T) {
	actual := NewSecurityAdminConfigurationRuleCollectionRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "networkManagerName", "configurationName", "ruleCollectionName", "ruleName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/configurationName/ruleCollections/ruleCollectionName/rules/ruleName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSecurityAdminConfigurationRuleCollectionRuleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SecurityAdminConfigurationRuleCollectionRuleId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/configurationName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/configurationName/ruleCollections",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/configurationName/ruleCollections/ruleCollectionName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/configurationName/ruleCollections/ruleCollectionName/rules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/configurationName/ruleCollections/ruleCollectionName/rules/ruleName",
			Expected: &SecurityAdminConfigurationRuleCollectionRuleId{
				SubscriptionId:                 "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:              "example-resource-group",
				NetworkManagerName:             "networkManagerName",
				SecurityAdminConfigurationName: "configurationName",
				RuleCollectionName:             "ruleCollectionName",
				RuleName:                       "ruleName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/configurationName/ruleCollections/ruleCollectionName/rules/ruleName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSecurityAdminConfigurationRuleCollectionRuleID(v.Input)
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

		if actual.RuleName != v.Expected.RuleName {
			t.Fatalf("Expected %q but got %q for RuleName", v.Expected.RuleName, actual.RuleName)
		}

	}
}

func TestParseSecurityAdminConfigurationRuleCollectionRuleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SecurityAdminConfigurationRuleCollectionRuleId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/configurationName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs/nEtWoRkMaNaGeRnAmE/sEcUrItYaDmInCoNfIgUrAtIoNs/cOnFiGuRaTiOnNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/configurationName/ruleCollections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs/nEtWoRkMaNaGeRnAmE/sEcUrItYaDmInCoNfIgUrAtIoNs/cOnFiGuRaTiOnNaMe/rUlEcOlLeCtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/configurationName/ruleCollections/ruleCollectionName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs/nEtWoRkMaNaGeRnAmE/sEcUrItYaDmInCoNfIgUrAtIoNs/cOnFiGuRaTiOnNaMe/rUlEcOlLeCtIoNs/rUlEcOlLeCtIoNnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/configurationName/ruleCollections/ruleCollectionName/rules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs/nEtWoRkMaNaGeRnAmE/sEcUrItYaDmInCoNfIgUrAtIoNs/cOnFiGuRaTiOnNaMe/rUlEcOlLeCtIoNs/rUlEcOlLeCtIoNnAmE/rUlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/configurationName/ruleCollections/ruleCollectionName/rules/ruleName",
			Expected: &SecurityAdminConfigurationRuleCollectionRuleId{
				SubscriptionId:                 "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:              "example-resource-group",
				NetworkManagerName:             "networkManagerName",
				SecurityAdminConfigurationName: "configurationName",
				RuleCollectionName:             "ruleCollectionName",
				RuleName:                       "ruleName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/securityAdminConfigurations/configurationName/ruleCollections/ruleCollectionName/rules/ruleName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs/nEtWoRkMaNaGeRnAmE/sEcUrItYaDmInCoNfIgUrAtIoNs/cOnFiGuRaTiOnNaMe/rUlEcOlLeCtIoNs/rUlEcOlLeCtIoNnAmE/rUlEs/rUlEnAmE",
			Expected: &SecurityAdminConfigurationRuleCollectionRuleId{
				SubscriptionId:                 "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:              "eXaMpLe-rEsOuRcE-GrOuP",
				NetworkManagerName:             "nEtWoRkMaNaGeRnAmE",
				SecurityAdminConfigurationName: "cOnFiGuRaTiOnNaMe",
				RuleCollectionName:             "rUlEcOlLeCtIoNnAmE",
				RuleName:                       "rUlEnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs/nEtWoRkMaNaGeRnAmE/sEcUrItYaDmInCoNfIgUrAtIoNs/cOnFiGuRaTiOnNaMe/rUlEcOlLeCtIoNs/rUlEcOlLeCtIoNnAmE/rUlEs/rUlEnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSecurityAdminConfigurationRuleCollectionRuleIDInsensitively(v.Input)
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

		if actual.RuleName != v.Expected.RuleName {
			t.Fatalf("Expected %q but got %q for RuleName", v.Expected.RuleName, actual.RuleName)
		}

	}
}

func TestSegmentsForSecurityAdminConfigurationRuleCollectionRuleId(t *testing.T) {
	segments := SecurityAdminConfigurationRuleCollectionRuleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SecurityAdminConfigurationRuleCollectionRuleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
