package reachabilityanalysisintent

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReachabilityAnalysisIntentId{}

func TestNewReachabilityAnalysisIntentID(t *testing.T) {
	id := NewReachabilityAnalysisIntentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "networkManagerName", "verifierWorkspaceName", "reachabilityAnalysisIntentName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.NetworkManagerName != "networkManagerName" {
		t.Fatalf("Expected %q but got %q for Segment 'NetworkManagerName'", id.NetworkManagerName, "networkManagerName")
	}

	if id.VerifierWorkspaceName != "verifierWorkspaceName" {
		t.Fatalf("Expected %q but got %q for Segment 'VerifierWorkspaceName'", id.VerifierWorkspaceName, "verifierWorkspaceName")
	}

	if id.ReachabilityAnalysisIntentName != "reachabilityAnalysisIntentName" {
		t.Fatalf("Expected %q but got %q for Segment 'ReachabilityAnalysisIntentName'", id.ReachabilityAnalysisIntentName, "reachabilityAnalysisIntentName")
	}
}

func TestFormatReachabilityAnalysisIntentID(t *testing.T) {
	actual := NewReachabilityAnalysisIntentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "networkManagerName", "verifierWorkspaceName", "reachabilityAnalysisIntentName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/verifierWorkspaces/verifierWorkspaceName/reachabilityAnalysisIntents/reachabilityAnalysisIntentName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReachabilityAnalysisIntentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReachabilityAnalysisIntentId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/verifierWorkspaces",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/verifierWorkspaces/verifierWorkspaceName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/verifierWorkspaces/verifierWorkspaceName/reachabilityAnalysisIntents",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/verifierWorkspaces/verifierWorkspaceName/reachabilityAnalysisIntents/reachabilityAnalysisIntentName",
			Expected: &ReachabilityAnalysisIntentId{
				SubscriptionId:                 "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:              "example-resource-group",
				NetworkManagerName:             "networkManagerName",
				VerifierWorkspaceName:          "verifierWorkspaceName",
				ReachabilityAnalysisIntentName: "reachabilityAnalysisIntentName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/verifierWorkspaces/verifierWorkspaceName/reachabilityAnalysisIntents/reachabilityAnalysisIntentName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReachabilityAnalysisIntentID(v.Input)
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

		if actual.VerifierWorkspaceName != v.Expected.VerifierWorkspaceName {
			t.Fatalf("Expected %q but got %q for VerifierWorkspaceName", v.Expected.VerifierWorkspaceName, actual.VerifierWorkspaceName)
		}

		if actual.ReachabilityAnalysisIntentName != v.Expected.ReachabilityAnalysisIntentName {
			t.Fatalf("Expected %q but got %q for ReachabilityAnalysisIntentName", v.Expected.ReachabilityAnalysisIntentName, actual.ReachabilityAnalysisIntentName)
		}

	}
}

func TestParseReachabilityAnalysisIntentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReachabilityAnalysisIntentId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/verifierWorkspaces",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs/nEtWoRkMaNaGeRnAmE/vErIfIeRwOrKsPaCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/verifierWorkspaces/verifierWorkspaceName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs/nEtWoRkMaNaGeRnAmE/vErIfIeRwOrKsPaCeS/vErIfIeRwOrKsPaCeNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/verifierWorkspaces/verifierWorkspaceName/reachabilityAnalysisIntents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs/nEtWoRkMaNaGeRnAmE/vErIfIeRwOrKsPaCeS/vErIfIeRwOrKsPaCeNaMe/rEaChAbIlItYaNaLySiSiNtEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/verifierWorkspaces/verifierWorkspaceName/reachabilityAnalysisIntents/reachabilityAnalysisIntentName",
			Expected: &ReachabilityAnalysisIntentId{
				SubscriptionId:                 "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:              "example-resource-group",
				NetworkManagerName:             "networkManagerName",
				VerifierWorkspaceName:          "verifierWorkspaceName",
				ReachabilityAnalysisIntentName: "reachabilityAnalysisIntentName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/verifierWorkspaces/verifierWorkspaceName/reachabilityAnalysisIntents/reachabilityAnalysisIntentName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs/nEtWoRkMaNaGeRnAmE/vErIfIeRwOrKsPaCeS/vErIfIeRwOrKsPaCeNaMe/rEaChAbIlItYaNaLySiSiNtEnTs/rEaChAbIlItYaNaLySiSiNtEnTnAmE",
			Expected: &ReachabilityAnalysisIntentId{
				SubscriptionId:                 "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:              "eXaMpLe-rEsOuRcE-GrOuP",
				NetworkManagerName:             "nEtWoRkMaNaGeRnAmE",
				VerifierWorkspaceName:          "vErIfIeRwOrKsPaCeNaMe",
				ReachabilityAnalysisIntentName: "rEaChAbIlItYaNaLySiSiNtEnTnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs/nEtWoRkMaNaGeRnAmE/vErIfIeRwOrKsPaCeS/vErIfIeRwOrKsPaCeNaMe/rEaChAbIlItYaNaLySiSiNtEnTs/rEaChAbIlItYaNaLySiSiNtEnTnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReachabilityAnalysisIntentIDInsensitively(v.Input)
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

		if actual.VerifierWorkspaceName != v.Expected.VerifierWorkspaceName {
			t.Fatalf("Expected %q but got %q for VerifierWorkspaceName", v.Expected.VerifierWorkspaceName, actual.VerifierWorkspaceName)
		}

		if actual.ReachabilityAnalysisIntentName != v.Expected.ReachabilityAnalysisIntentName {
			t.Fatalf("Expected %q but got %q for ReachabilityAnalysisIntentName", v.Expected.ReachabilityAnalysisIntentName, actual.ReachabilityAnalysisIntentName)
		}

	}
}

func TestSegmentsForReachabilityAnalysisIntentId(t *testing.T) {
	segments := ReachabilityAnalysisIntentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReachabilityAnalysisIntentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
