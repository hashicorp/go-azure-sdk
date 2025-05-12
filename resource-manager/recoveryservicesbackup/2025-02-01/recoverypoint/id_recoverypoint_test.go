package recoverypoint

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RecoveryPointId{}

func TestNewRecoveryPointID(t *testing.T) {
	id := NewRecoveryPointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "backupFabricName", "protectionContainerName", "protectedItemName", "recoveryPointId")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.VaultName != "vaultName" {
		t.Fatalf("Expected %q but got %q for Segment 'VaultName'", id.VaultName, "vaultName")
	}

	if id.BackupFabricName != "backupFabricName" {
		t.Fatalf("Expected %q but got %q for Segment 'BackupFabricName'", id.BackupFabricName, "backupFabricName")
	}

	if id.ProtectionContainerName != "protectionContainerName" {
		t.Fatalf("Expected %q but got %q for Segment 'ProtectionContainerName'", id.ProtectionContainerName, "protectionContainerName")
	}

	if id.ProtectedItemName != "protectedItemName" {
		t.Fatalf("Expected %q but got %q for Segment 'ProtectedItemName'", id.ProtectedItemName, "protectedItemName")
	}

	if id.RecoveryPointId != "recoveryPointId" {
		t.Fatalf("Expected %q but got %q for Segment 'RecoveryPointId'", id.RecoveryPointId, "recoveryPointId")
	}
}

func TestFormatRecoveryPointID(t *testing.T) {
	actual := NewRecoveryPointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "backupFabricName", "protectionContainerName", "protectedItemName", "recoveryPointId").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName/backupFabrics/backupFabricName/protectionContainers/protectionContainerName/protectedItems/protectedItemName/recoveryPoints/recoveryPointId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRecoveryPointID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RecoveryPointId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName/backupFabrics",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName/backupFabrics/backupFabricName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName/backupFabrics/backupFabricName/protectionContainers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName/backupFabrics/backupFabricName/protectionContainers/protectionContainerName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName/backupFabrics/backupFabricName/protectionContainers/protectionContainerName/protectedItems",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName/backupFabrics/backupFabricName/protectionContainers/protectionContainerName/protectedItems/protectedItemName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName/backupFabrics/backupFabricName/protectionContainers/protectionContainerName/protectedItems/protectedItemName/recoveryPoints",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName/backupFabrics/backupFabricName/protectionContainers/protectionContainerName/protectedItems/protectedItemName/recoveryPoints/recoveryPointId",
			Expected: &RecoveryPointId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:       "example-resource-group",
				VaultName:               "vaultName",
				BackupFabricName:        "backupFabricName",
				ProtectionContainerName: "protectionContainerName",
				ProtectedItemName:       "protectedItemName",
				RecoveryPointId:         "recoveryPointId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName/backupFabrics/backupFabricName/protectionContainers/protectionContainerName/protectedItems/protectedItemName/recoveryPoints/recoveryPointId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRecoveryPointID(v.Input)
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

		if actual.BackupFabricName != v.Expected.BackupFabricName {
			t.Fatalf("Expected %q but got %q for BackupFabricName", v.Expected.BackupFabricName, actual.BackupFabricName)
		}

		if actual.ProtectionContainerName != v.Expected.ProtectionContainerName {
			t.Fatalf("Expected %q but got %q for ProtectionContainerName", v.Expected.ProtectionContainerName, actual.ProtectionContainerName)
		}

		if actual.ProtectedItemName != v.Expected.ProtectedItemName {
			t.Fatalf("Expected %q but got %q for ProtectedItemName", v.Expected.ProtectedItemName, actual.ProtectedItemName)
		}

		if actual.RecoveryPointId != v.Expected.RecoveryPointId {
			t.Fatalf("Expected %q but got %q for RecoveryPointId", v.Expected.RecoveryPointId, actual.RecoveryPointId)
		}

	}
}

func TestParseRecoveryPointIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RecoveryPointId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName/backupFabrics",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtNaMe/bAcKuPfAbRiCs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName/backupFabrics/backupFabricName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtNaMe/bAcKuPfAbRiCs/bAcKuPfAbRiCnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName/backupFabrics/backupFabricName/protectionContainers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtNaMe/bAcKuPfAbRiCs/bAcKuPfAbRiCnAmE/pRoTeCtIoNcOnTaInErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName/backupFabrics/backupFabricName/protectionContainers/protectionContainerName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtNaMe/bAcKuPfAbRiCs/bAcKuPfAbRiCnAmE/pRoTeCtIoNcOnTaInErS/pRoTeCtIoNcOnTaInErNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName/backupFabrics/backupFabricName/protectionContainers/protectionContainerName/protectedItems",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtNaMe/bAcKuPfAbRiCs/bAcKuPfAbRiCnAmE/pRoTeCtIoNcOnTaInErS/pRoTeCtIoNcOnTaInErNaMe/pRoTeCtEdItEmS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName/backupFabrics/backupFabricName/protectionContainers/protectionContainerName/protectedItems/protectedItemName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtNaMe/bAcKuPfAbRiCs/bAcKuPfAbRiCnAmE/pRoTeCtIoNcOnTaInErS/pRoTeCtIoNcOnTaInErNaMe/pRoTeCtEdItEmS/pRoTeCtEdItEmNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName/backupFabrics/backupFabricName/protectionContainers/protectionContainerName/protectedItems/protectedItemName/recoveryPoints",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtNaMe/bAcKuPfAbRiCs/bAcKuPfAbRiCnAmE/pRoTeCtIoNcOnTaInErS/pRoTeCtIoNcOnTaInErNaMe/pRoTeCtEdItEmS/pRoTeCtEdItEmNaMe/rEcOvErYpOiNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName/backupFabrics/backupFabricName/protectionContainers/protectionContainerName/protectedItems/protectedItemName/recoveryPoints/recoveryPointId",
			Expected: &RecoveryPointId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:       "example-resource-group",
				VaultName:               "vaultName",
				BackupFabricName:        "backupFabricName",
				ProtectionContainerName: "protectionContainerName",
				ProtectedItemName:       "protectedItemName",
				RecoveryPointId:         "recoveryPointId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultName/backupFabrics/backupFabricName/protectionContainers/protectionContainerName/protectedItems/protectedItemName/recoveryPoints/recoveryPointId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtNaMe/bAcKuPfAbRiCs/bAcKuPfAbRiCnAmE/pRoTeCtIoNcOnTaInErS/pRoTeCtIoNcOnTaInErNaMe/pRoTeCtEdItEmS/pRoTeCtEdItEmNaMe/rEcOvErYpOiNtS/rEcOvErYpOiNtId",
			Expected: &RecoveryPointId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:       "eXaMpLe-rEsOuRcE-GrOuP",
				VaultName:               "vAuLtNaMe",
				BackupFabricName:        "bAcKuPfAbRiCnAmE",
				ProtectionContainerName: "pRoTeCtIoNcOnTaInErNaMe",
				ProtectedItemName:       "pRoTeCtEdItEmNaMe",
				RecoveryPointId:         "rEcOvErYpOiNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtNaMe/bAcKuPfAbRiCs/bAcKuPfAbRiCnAmE/pRoTeCtIoNcOnTaInErS/pRoTeCtIoNcOnTaInErNaMe/pRoTeCtEdItEmS/pRoTeCtEdItEmNaMe/rEcOvErYpOiNtS/rEcOvErYpOiNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRecoveryPointIDInsensitively(v.Input)
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

		if actual.BackupFabricName != v.Expected.BackupFabricName {
			t.Fatalf("Expected %q but got %q for BackupFabricName", v.Expected.BackupFabricName, actual.BackupFabricName)
		}

		if actual.ProtectionContainerName != v.Expected.ProtectionContainerName {
			t.Fatalf("Expected %q but got %q for ProtectionContainerName", v.Expected.ProtectionContainerName, actual.ProtectionContainerName)
		}

		if actual.ProtectedItemName != v.Expected.ProtectedItemName {
			t.Fatalf("Expected %q but got %q for ProtectedItemName", v.Expected.ProtectedItemName, actual.ProtectedItemName)
		}

		if actual.RecoveryPointId != v.Expected.RecoveryPointId {
			t.Fatalf("Expected %q but got %q for RecoveryPointId", v.Expected.RecoveryPointId, actual.RecoveryPointId)
		}

	}
}

func TestSegmentsForRecoveryPointId(t *testing.T) {
	segments := RecoveryPointId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RecoveryPointId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
