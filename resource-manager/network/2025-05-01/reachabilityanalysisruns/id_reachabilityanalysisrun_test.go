package reachabilityanalysisruns

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReachabilityAnalysisRunId{}

func TestNewReachabilityAnalysisRunID(t *testing.T) {
	id := NewReachabilityAnalysisRunID("12345678-1234-9876-4563-123456789012", "example-resource-group", "networkManagerName", "verifierWorkspaceName", "reachabilityAnalysisRunName")

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

	if id.ReachabilityAnalysisRunName != "reachabilityAnalysisRunName" {
		t.Fatalf("Expected %q but got %q for Segment 'ReachabilityAnalysisRunName'", id.ReachabilityAnalysisRunName, "reachabilityAnalysisRunName")
	}
}

func TestFormatReachabilityAnalysisRunID(t *testing.T) {
	actual := NewReachabilityAnalysisRunID("12345678-1234-9876-4563-123456789012", "example-resource-group", "networkManagerName", "verifierWorkspaceName", "reachabilityAnalysisRunName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/verifierWorkspaces/verifierWorkspaceName/reachabilityAnalysisRuns/reachabilityAnalysisRunName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReachabilityAnalysisRunID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReachabilityAnalysisRunId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/verifierWorkspaces/verifierWorkspaceName/reachabilityAnalysisRuns",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/verifierWorkspaces/verifierWorkspaceName/reachabilityAnalysisRuns/reachabilityAnalysisRunName",
			Expected: &ReachabilityAnalysisRunId{
				SubscriptionId:              "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:           "example-resource-group",
				NetworkManagerName:          "networkManagerName",
				VerifierWorkspaceName:       "verifierWorkspaceName",
				ReachabilityAnalysisRunName: "reachabilityAnalysisRunName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/verifierWorkspaces/verifierWorkspaceName/reachabilityAnalysisRuns/reachabilityAnalysisRunName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReachabilityAnalysisRunID(v.Input)
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

		if actual.ReachabilityAnalysisRunName != v.Expected.ReachabilityAnalysisRunName {
			t.Fatalf("Expected %q but got %q for ReachabilityAnalysisRunName", v.Expected.ReachabilityAnalysisRunName, actual.ReachabilityAnalysisRunName)
		}

	}
}

func TestParseReachabilityAnalysisRunIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReachabilityAnalysisRunId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/verifierWorkspaces/verifierWorkspaceName/reachabilityAnalysisRuns",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs/nEtWoRkMaNaGeRnAmE/vErIfIeRwOrKsPaCeS/vErIfIeRwOrKsPaCeNaMe/rEaChAbIlItYaNaLySiSrUnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/verifierWorkspaces/verifierWorkspaceName/reachabilityAnalysisRuns/reachabilityAnalysisRunName",
			Expected: &ReachabilityAnalysisRunId{
				SubscriptionId:              "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:           "example-resource-group",
				NetworkManagerName:          "networkManagerName",
				VerifierWorkspaceName:       "verifierWorkspaceName",
				ReachabilityAnalysisRunName: "reachabilityAnalysisRunName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkManagers/networkManagerName/verifierWorkspaces/verifierWorkspaceName/reachabilityAnalysisRuns/reachabilityAnalysisRunName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs/nEtWoRkMaNaGeRnAmE/vErIfIeRwOrKsPaCeS/vErIfIeRwOrKsPaCeNaMe/rEaChAbIlItYaNaLySiSrUnS/rEaChAbIlItYaNaLySiSrUnNaMe",
			Expected: &ReachabilityAnalysisRunId{
				SubscriptionId:              "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:           "eXaMpLe-rEsOuRcE-GrOuP",
				NetworkManagerName:          "nEtWoRkMaNaGeRnAmE",
				VerifierWorkspaceName:       "vErIfIeRwOrKsPaCeNaMe",
				ReachabilityAnalysisRunName: "rEaChAbIlItYaNaLySiSrUnNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRs/nEtWoRkMaNaGeRnAmE/vErIfIeRwOrKsPaCeS/vErIfIeRwOrKsPaCeNaMe/rEaChAbIlItYaNaLySiSrUnS/rEaChAbIlItYaNaLySiSrUnNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReachabilityAnalysisRunIDInsensitively(v.Input)
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

		if actual.ReachabilityAnalysisRunName != v.Expected.ReachabilityAnalysisRunName {
			t.Fatalf("Expected %q but got %q for ReachabilityAnalysisRunName", v.Expected.ReachabilityAnalysisRunName, actual.ReachabilityAnalysisRunName)
		}

	}
}

func TestSegmentsForReachabilityAnalysisRunId(t *testing.T) {
	segments := ReachabilityAnalysisRunId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReachabilityAnalysisRunId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
