package policystates

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId{}

func TestNewPolicySetDefinitionProviders2PolicyStatePolicyStatesResourceID(t *testing.T) {
	id := NewPolicySetDefinitionProviders2PolicyStatePolicyStatesResourceID("12345678-1234-9876-4563-123456789012", "policySetDefinitionName", "default")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.PolicySetDefinitionName != "policySetDefinitionName" {
		t.Fatalf("Expected %q but got %q for Segment 'PolicySetDefinitionName'", id.PolicySetDefinitionName, "policySetDefinitionName")
	}

	if id.PolicyStatesResource != "default" {
		t.Fatalf("Expected %q but got %q for Segment 'PolicyStatesResource'", id.PolicyStatesResource, "default")
	}
}

func TestFormatPolicySetDefinitionProviders2PolicyStatePolicyStatesResourceID(t *testing.T) {
	actual := NewPolicySetDefinitionProviders2PolicyStatePolicyStatesResourceID("12345678-1234-9876-4563-123456789012", "policySetDefinitionName", "default").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName/providers/Microsoft.PolicyInsights/policyStates/default"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicySetDefinitionProviders2PolicyStatePolicyStatesResourceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policySetDefinitions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName/providers/Microsoft.PolicyInsights",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName/providers/Microsoft.PolicyInsights/policyStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName/providers/Microsoft.PolicyInsights/policyStates/default",
			Expected: &PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				PolicySetDefinitionName: "policySetDefinitionName",
				PolicyStatesResource:    "default",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName/providers/Microsoft.PolicyInsights/policyStates/default/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicySetDefinitionProviders2PolicyStatePolicyStatesResourceID(v.Input)
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

		if actual.PolicySetDefinitionName != v.Expected.PolicySetDefinitionName {
			t.Fatalf("Expected %q but got %q for PolicySetDefinitionName", v.Expected.PolicySetDefinitionName, actual.PolicySetDefinitionName)
		}

		if actual.PolicyStatesResource != v.Expected.PolicyStatesResource {
			t.Fatalf("Expected %q but got %q for PolicyStatesResource", v.Expected.PolicyStatesResource, actual.PolicyStatesResource)
		}

	}
}

func TestParsePolicySetDefinitionProviders2PolicyStatePolicyStatesResourceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policySetDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS/pOlIcYsEtDeFiNiTiOnNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS/pOlIcYsEtDeFiNiTiOnNaMe/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName/providers/Microsoft.PolicyInsights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS/pOlIcYsEtDeFiNiTiOnNaMe/pRoViDeRs/mIcRoSoFt.pOlIcYiNsIgHtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName/providers/Microsoft.PolicyInsights/policyStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS/pOlIcYsEtDeFiNiTiOnNaMe/pRoViDeRs/mIcRoSoFt.pOlIcYiNsIgHtS/pOlIcYsTaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName/providers/Microsoft.PolicyInsights/policyStates/default",
			Expected: &PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				PolicySetDefinitionName: "policySetDefinitionName",
				PolicyStatesResource:    "default",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName/providers/Microsoft.PolicyInsights/policyStates/default/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS/pOlIcYsEtDeFiNiTiOnNaMe/pRoViDeRs/mIcRoSoFt.pOlIcYiNsIgHtS/pOlIcYsTaTeS/dEfAuLt",
			Expected: &PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				PolicySetDefinitionName: "pOlIcYsEtDeFiNiTiOnNaMe",
				PolicyStatesResource:    "default",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS/pOlIcYsEtDeFiNiTiOnNaMe/pRoViDeRs/mIcRoSoFt.pOlIcYiNsIgHtS/pOlIcYsTaTeS/dEfAuLt/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicySetDefinitionProviders2PolicyStatePolicyStatesResourceIDInsensitively(v.Input)
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

		if actual.PolicySetDefinitionName != v.Expected.PolicySetDefinitionName {
			t.Fatalf("Expected %q but got %q for PolicySetDefinitionName", v.Expected.PolicySetDefinitionName, actual.PolicySetDefinitionName)
		}

		if actual.PolicyStatesResource != v.Expected.PolicyStatesResource {
			t.Fatalf("Expected %q but got %q for PolicyStatesResource", v.Expected.PolicyStatesResource, actual.PolicyStatesResource)
		}

	}
}

func TestSegmentsForPolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId(t *testing.T) {
	segments := PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
