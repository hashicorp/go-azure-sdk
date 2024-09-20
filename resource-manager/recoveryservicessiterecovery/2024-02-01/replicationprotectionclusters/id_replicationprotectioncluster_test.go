package replicationprotectionclusters

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReplicationProtectionClusterId{}

func TestNewReplicationProtectionClusterID(t *testing.T) {
	id := NewReplicationProtectionClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "fabricName", "protectionContainerName", "replicationProtectionClusterName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.VaultName != "resourceName" {
		t.Fatalf("Expected %q but got %q for Segment 'VaultName'", id.VaultName, "resourceName")
	}

	if id.ReplicationFabricName != "fabricName" {
		t.Fatalf("Expected %q but got %q for Segment 'ReplicationFabricName'", id.ReplicationFabricName, "fabricName")
	}

	if id.ReplicationProtectionContainerName != "protectionContainerName" {
		t.Fatalf("Expected %q but got %q for Segment 'ReplicationProtectionContainerName'", id.ReplicationProtectionContainerName, "protectionContainerName")
	}

	if id.ReplicationProtectionClusterName != "replicationProtectionClusterName" {
		t.Fatalf("Expected %q but got %q for Segment 'ReplicationProtectionClusterName'", id.ReplicationProtectionClusterName, "replicationProtectionClusterName")
	}
}

func TestFormatReplicationProtectionClusterID(t *testing.T) {
	actual := NewReplicationProtectionClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "fabricName", "protectionContainerName", "replicationProtectionClusterName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceName/replicationFabrics/fabricName/replicationProtectionContainers/protectionContainerName/replicationProtectionClusters/replicationProtectionClusterName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReplicationProtectionClusterID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReplicationProtectionClusterId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceName/replicationFabrics",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceName/replicationFabrics/fabricName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceName/replicationFabrics/fabricName/replicationProtectionContainers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceName/replicationFabrics/fabricName/replicationProtectionContainers/protectionContainerName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceName/replicationFabrics/fabricName/replicationProtectionContainers/protectionContainerName/replicationProtectionClusters",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceName/replicationFabrics/fabricName/replicationProtectionContainers/protectionContainerName/replicationProtectionClusters/replicationProtectionClusterName",
			Expected: &ReplicationProtectionClusterId{
				SubscriptionId:                     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                  "example-resource-group",
				VaultName:                          "resourceName",
				ReplicationFabricName:              "fabricName",
				ReplicationProtectionContainerName: "protectionContainerName",
				ReplicationProtectionClusterName:   "replicationProtectionClusterName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceName/replicationFabrics/fabricName/replicationProtectionContainers/protectionContainerName/replicationProtectionClusters/replicationProtectionClusterName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReplicationProtectionClusterID(v.Input)
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

		if actual.VaultName != v.Expected.VaultName {
			t.Fatalf("Expected %q but got %q for VaultName", v.Expected.VaultName, actual.VaultName)
		}

		if actual.ReplicationFabricName != v.Expected.ReplicationFabricName {
			t.Fatalf("Expected %q but got %q for ReplicationFabricName", v.Expected.ReplicationFabricName, actual.ReplicationFabricName)
		}

		if actual.ReplicationProtectionContainerName != v.Expected.ReplicationProtectionContainerName {
			t.Fatalf("Expected %q but got %q for ReplicationProtectionContainerName", v.Expected.ReplicationProtectionContainerName, actual.ReplicationProtectionContainerName)
		}

		if actual.ReplicationProtectionClusterName != v.Expected.ReplicationProtectionClusterName {
			t.Fatalf("Expected %q but got %q for ReplicationProtectionClusterName", v.Expected.ReplicationProtectionClusterName, actual.ReplicationProtectionClusterName)
		}

	}
}

func TestParseReplicationProtectionClusterIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReplicationProtectionClusterId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/rEsOuRcEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceName/replicationFabrics",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/rEsOuRcEnAmE/rEpLiCaTiOnFaBrIcS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceName/replicationFabrics/fabricName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/rEsOuRcEnAmE/rEpLiCaTiOnFaBrIcS/fAbRiCnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceName/replicationFabrics/fabricName/replicationProtectionContainers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/rEsOuRcEnAmE/rEpLiCaTiOnFaBrIcS/fAbRiCnAmE/rEpLiCaTiOnPrOtEcTiOnCoNtAiNeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceName/replicationFabrics/fabricName/replicationProtectionContainers/protectionContainerName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/rEsOuRcEnAmE/rEpLiCaTiOnFaBrIcS/fAbRiCnAmE/rEpLiCaTiOnPrOtEcTiOnCoNtAiNeRs/pRoTeCtIoNcOnTaInErNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceName/replicationFabrics/fabricName/replicationProtectionContainers/protectionContainerName/replicationProtectionClusters",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/rEsOuRcEnAmE/rEpLiCaTiOnFaBrIcS/fAbRiCnAmE/rEpLiCaTiOnPrOtEcTiOnCoNtAiNeRs/pRoTeCtIoNcOnTaInErNaMe/rEpLiCaTiOnPrOtEcTiOnClUsTeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceName/replicationFabrics/fabricName/replicationProtectionContainers/protectionContainerName/replicationProtectionClusters/replicationProtectionClusterName",
			Expected: &ReplicationProtectionClusterId{
				SubscriptionId:                     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                  "example-resource-group",
				VaultName:                          "resourceName",
				ReplicationFabricName:              "fabricName",
				ReplicationProtectionContainerName: "protectionContainerName",
				ReplicationProtectionClusterName:   "replicationProtectionClusterName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceName/replicationFabrics/fabricName/replicationProtectionContainers/protectionContainerName/replicationProtectionClusters/replicationProtectionClusterName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/rEsOuRcEnAmE/rEpLiCaTiOnFaBrIcS/fAbRiCnAmE/rEpLiCaTiOnPrOtEcTiOnCoNtAiNeRs/pRoTeCtIoNcOnTaInErNaMe/rEpLiCaTiOnPrOtEcTiOnClUsTeRs/rEpLiCaTiOnPrOtEcTiOnClUsTeRnAmE",
			Expected: &ReplicationProtectionClusterId{
				SubscriptionId:                     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                  "eXaMpLe-rEsOuRcE-GrOuP",
				VaultName:                          "rEsOuRcEnAmE",
				ReplicationFabricName:              "fAbRiCnAmE",
				ReplicationProtectionContainerName: "pRoTeCtIoNcOnTaInErNaMe",
				ReplicationProtectionClusterName:   "rEpLiCaTiOnPrOtEcTiOnClUsTeRnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/rEsOuRcEnAmE/rEpLiCaTiOnFaBrIcS/fAbRiCnAmE/rEpLiCaTiOnPrOtEcTiOnCoNtAiNeRs/pRoTeCtIoNcOnTaInErNaMe/rEpLiCaTiOnPrOtEcTiOnClUsTeRs/rEpLiCaTiOnPrOtEcTiOnClUsTeRnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReplicationProtectionClusterIDInsensitively(v.Input)
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

		if actual.VaultName != v.Expected.VaultName {
			t.Fatalf("Expected %q but got %q for VaultName", v.Expected.VaultName, actual.VaultName)
		}

		if actual.ReplicationFabricName != v.Expected.ReplicationFabricName {
			t.Fatalf("Expected %q but got %q for ReplicationFabricName", v.Expected.ReplicationFabricName, actual.ReplicationFabricName)
		}

		if actual.ReplicationProtectionContainerName != v.Expected.ReplicationProtectionContainerName {
			t.Fatalf("Expected %q but got %q for ReplicationProtectionContainerName", v.Expected.ReplicationProtectionContainerName, actual.ReplicationProtectionContainerName)
		}

		if actual.ReplicationProtectionClusterName != v.Expected.ReplicationProtectionClusterName {
			t.Fatalf("Expected %q but got %q for ReplicationProtectionClusterName", v.Expected.ReplicationProtectionClusterName, actual.ReplicationProtectionClusterName)
		}

	}
}

func TestSegmentsForReplicationProtectionClusterId(t *testing.T) {
	segments := ReplicationProtectionClusterId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReplicationProtectionClusterId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
