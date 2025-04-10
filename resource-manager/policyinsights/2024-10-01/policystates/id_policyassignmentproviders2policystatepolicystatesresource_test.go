package policystates

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId{}

func TestNewPolicyAssignmentProviders2PolicyStatePolicyStatesResourceID(t *testing.T) {
	id := NewPolicyAssignmentProviders2PolicyStatePolicyStatesResourceID("12345678-1234-9876-4563-123456789012", "policyAssignmentName", "default")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.PolicyAssignmentName != "policyAssignmentName" {
		t.Fatalf("Expected %q but got %q for Segment 'PolicyAssignmentName'", id.PolicyAssignmentName, "policyAssignmentName")
	}

	if id.PolicyStatesResource != "default" {
		t.Fatalf("Expected %q but got %q for Segment 'PolicyStatesResource'", id.PolicyStatesResource, "default")
	}
}

func TestFormatPolicyAssignmentProviders2PolicyStatePolicyStatesResourceID(t *testing.T) {
	actual := NewPolicyAssignmentProviders2PolicyStatePolicyStatesResourceID("12345678-1234-9876-4563-123456789012", "policyAssignmentName", "default").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyAssignments/policyAssignmentName/providers/Microsoft.PolicyInsights/policyStates/default"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyAssignmentProviders2PolicyStatePolicyStatesResourceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyAssignments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyAssignments/policyAssignmentName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyAssignments/policyAssignmentName/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyAssignments/policyAssignmentName/providers/Microsoft.PolicyInsights",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyAssignments/policyAssignmentName/providers/Microsoft.PolicyInsights/policyStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyAssignments/policyAssignmentName/providers/Microsoft.PolicyInsights/policyStates/default",
			Expected: &PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				PolicyAssignmentName: "policyAssignmentName",
				PolicyStatesResource: "default",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyAssignments/policyAssignmentName/providers/Microsoft.PolicyInsights/policyStates/default/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyAssignmentProviders2PolicyStatePolicyStatesResourceID(v.Input)
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

		if actual.PolicyAssignmentName != v.Expected.PolicyAssignmentName {
			t.Fatalf("Expected %q but got %q for PolicyAssignmentName", v.Expected.PolicyAssignmentName, actual.PolicyAssignmentName)
		}

		if actual.PolicyStatesResource != v.Expected.PolicyStatesResource {
			t.Fatalf("Expected %q but got %q for PolicyStatesResource", v.Expected.PolicyStatesResource, actual.PolicyStatesResource)
		}

	}
}

func TestParsePolicyAssignmentProviders2PolicyStatePolicyStatesResourceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYaSsIgNmEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyAssignments/policyAssignmentName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYaSsIgNmEnTs/pOlIcYaSsIgNmEnTnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyAssignments/policyAssignmentName/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYaSsIgNmEnTs/pOlIcYaSsIgNmEnTnAmE/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyAssignments/policyAssignmentName/providers/Microsoft.PolicyInsights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYaSsIgNmEnTs/pOlIcYaSsIgNmEnTnAmE/pRoViDeRs/mIcRoSoFt.pOlIcYiNsIgHtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyAssignments/policyAssignmentName/providers/Microsoft.PolicyInsights/policyStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYaSsIgNmEnTs/pOlIcYaSsIgNmEnTnAmE/pRoViDeRs/mIcRoSoFt.pOlIcYiNsIgHtS/pOlIcYsTaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyAssignments/policyAssignmentName/providers/Microsoft.PolicyInsights/policyStates/default",
			Expected: &PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				PolicyAssignmentName: "policyAssignmentName",
				PolicyStatesResource: "default",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyAssignments/policyAssignmentName/providers/Microsoft.PolicyInsights/policyStates/default/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYaSsIgNmEnTs/pOlIcYaSsIgNmEnTnAmE/pRoViDeRs/mIcRoSoFt.pOlIcYiNsIgHtS/pOlIcYsTaTeS/dEfAuLt",
			Expected: &PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				PolicyAssignmentName: "pOlIcYaSsIgNmEnTnAmE",
				PolicyStatesResource: "default",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYaSsIgNmEnTs/pOlIcYaSsIgNmEnTnAmE/pRoViDeRs/mIcRoSoFt.pOlIcYiNsIgHtS/pOlIcYsTaTeS/dEfAuLt/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyAssignmentProviders2PolicyStatePolicyStatesResourceIDInsensitively(v.Input)
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

		if actual.PolicyAssignmentName != v.Expected.PolicyAssignmentName {
			t.Fatalf("Expected %q but got %q for PolicyAssignmentName", v.Expected.PolicyAssignmentName, actual.PolicyAssignmentName)
		}

		if actual.PolicyStatesResource != v.Expected.PolicyStatesResource {
			t.Fatalf("Expected %q but got %q for PolicyStatesResource", v.Expected.PolicyStatesResource, actual.PolicyStatesResource)
		}

	}
}

func TestSegmentsForPolicyAssignmentProviders2PolicyStatePolicyStatesResourceId(t *testing.T) {
	segments := PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
