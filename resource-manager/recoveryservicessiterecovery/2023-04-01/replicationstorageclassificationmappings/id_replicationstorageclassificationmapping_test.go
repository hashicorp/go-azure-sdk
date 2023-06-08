package replicationstorageclassificationmappings

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ReplicationStorageClassificationMappingId{}

func TestNewReplicationStorageClassificationMappingID(t *testing.T) {
	id := NewReplicationStorageClassificationMappingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationStorageClassificationValue", "replicationStorageClassificationMappingValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.VaultName != "vaultValue" {
		t.Fatalf("Expected %q but got %q for Segment 'VaultName'", id.VaultName, "vaultValue")
	}

	if id.ReplicationFabricName != "replicationFabricValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ReplicationFabricName'", id.ReplicationFabricName, "replicationFabricValue")
	}

	if id.ReplicationStorageClassificationName != "replicationStorageClassificationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ReplicationStorageClassificationName'", id.ReplicationStorageClassificationName, "replicationStorageClassificationValue")
	}

	if id.ReplicationStorageClassificationMappingName != "replicationStorageClassificationMappingValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ReplicationStorageClassificationMappingName'", id.ReplicationStorageClassificationMappingName, "replicationStorageClassificationMappingValue")
	}
}

func TestFormatReplicationStorageClassificationMappingID(t *testing.T) {
	actual := NewReplicationStorageClassificationMappingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationStorageClassificationValue", "replicationStorageClassificationMappingValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationStorageClassifications/replicationStorageClassificationValue/replicationStorageClassificationMappings/replicationStorageClassificationMappingValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReplicationStorageClassificationMappingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReplicationStorageClassificationMappingId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationStorageClassifications",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationStorageClassifications/replicationStorageClassificationValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationStorageClassifications/replicationStorageClassificationValue/replicationStorageClassificationMappings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationStorageClassifications/replicationStorageClassificationValue/replicationStorageClassificationMappings/replicationStorageClassificationMappingValue",
			Expected: &ReplicationStorageClassificationMappingId{
				SubscriptionId:                              "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                           "example-resource-group",
				VaultName:                                   "vaultValue",
				ReplicationFabricName:                       "replicationFabricValue",
				ReplicationStorageClassificationName:        "replicationStorageClassificationValue",
				ReplicationStorageClassificationMappingName: "replicationStorageClassificationMappingValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationStorageClassifications/replicationStorageClassificationValue/replicationStorageClassificationMappings/replicationStorageClassificationMappingValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReplicationStorageClassificationMappingID(v.Input)
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

		if actual.ReplicationStorageClassificationName != v.Expected.ReplicationStorageClassificationName {
			t.Fatalf("Expected %q but got %q for ReplicationStorageClassificationName", v.Expected.ReplicationStorageClassificationName, actual.ReplicationStorageClassificationName)
		}

		if actual.ReplicationStorageClassificationMappingName != v.Expected.ReplicationStorageClassificationMappingName {
			t.Fatalf("Expected %q but got %q for ReplicationStorageClassificationMappingName", v.Expected.ReplicationStorageClassificationMappingName, actual.ReplicationStorageClassificationMappingName)
		}

	}
}

func TestParseReplicationStorageClassificationMappingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReplicationStorageClassificationMappingId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtVaLuE/rEpLiCaTiOnFaBrIcS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtVaLuE/rEpLiCaTiOnFaBrIcS/rEpLiCaTiOnFaBrIcVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationStorageClassifications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtVaLuE/rEpLiCaTiOnFaBrIcS/rEpLiCaTiOnFaBrIcVaLuE/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationStorageClassifications/replicationStorageClassificationValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtVaLuE/rEpLiCaTiOnFaBrIcS/rEpLiCaTiOnFaBrIcVaLuE/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNs/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationStorageClassifications/replicationStorageClassificationValue/replicationStorageClassificationMappings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtVaLuE/rEpLiCaTiOnFaBrIcS/rEpLiCaTiOnFaBrIcVaLuE/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNs/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNvAlUe/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNmApPiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationStorageClassifications/replicationStorageClassificationValue/replicationStorageClassificationMappings/replicationStorageClassificationMappingValue",
			Expected: &ReplicationStorageClassificationMappingId{
				SubscriptionId:                              "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                           "example-resource-group",
				VaultName:                                   "vaultValue",
				ReplicationFabricName:                       "replicationFabricValue",
				ReplicationStorageClassificationName:        "replicationStorageClassificationValue",
				ReplicationStorageClassificationMappingName: "replicationStorageClassificationMappingValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationStorageClassifications/replicationStorageClassificationValue/replicationStorageClassificationMappings/replicationStorageClassificationMappingValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtVaLuE/rEpLiCaTiOnFaBrIcS/rEpLiCaTiOnFaBrIcVaLuE/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNs/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNvAlUe/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNmApPiNgS/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNmApPiNgVaLuE",
			Expected: &ReplicationStorageClassificationMappingId{
				SubscriptionId:                              "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                           "eXaMpLe-rEsOuRcE-GrOuP",
				VaultName:                                   "vAuLtVaLuE",
				ReplicationFabricName:                       "rEpLiCaTiOnFaBrIcVaLuE",
				ReplicationStorageClassificationName:        "rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNvAlUe",
				ReplicationStorageClassificationMappingName: "rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNmApPiNgVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtVaLuE/rEpLiCaTiOnFaBrIcS/rEpLiCaTiOnFaBrIcVaLuE/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNs/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNvAlUe/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNmApPiNgS/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNmApPiNgVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReplicationStorageClassificationMappingIDInsensitively(v.Input)
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

		if actual.ReplicationStorageClassificationName != v.Expected.ReplicationStorageClassificationName {
			t.Fatalf("Expected %q but got %q for ReplicationStorageClassificationName", v.Expected.ReplicationStorageClassificationName, actual.ReplicationStorageClassificationName)
		}

		if actual.ReplicationStorageClassificationMappingName != v.Expected.ReplicationStorageClassificationMappingName {
			t.Fatalf("Expected %q but got %q for ReplicationStorageClassificationMappingName", v.Expected.ReplicationStorageClassificationMappingName, actual.ReplicationStorageClassificationMappingName)
		}

	}
}

func TestSegmentsForReplicationStorageClassificationMappingId(t *testing.T) {
	segments := ReplicationStorageClassificationMappingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReplicationStorageClassificationMappingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
