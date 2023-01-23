package recoverypoints

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = RecoveryPointId{}

func TestNewRecoveryPointID(t *testing.T) {
	id := NewRecoveryPointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationProtectionContainerValue", "replicationProtectedItemValue", "recoveryPointValue")

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

	if id.ReplicationProtectionContainerName != "replicationProtectionContainerValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ReplicationProtectionContainerName'", id.ReplicationProtectionContainerName, "replicationProtectionContainerValue")
	}

	if id.ReplicationProtectedItemName != "replicationProtectedItemValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ReplicationProtectedItemName'", id.ReplicationProtectedItemName, "replicationProtectedItemValue")
	}

	if id.RecoveryPointName != "recoveryPointValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RecoveryPointName'", id.RecoveryPointName, "recoveryPointValue")
	}
}

func TestFormatRecoveryPointID(t *testing.T) {
	actual := NewRecoveryPointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationProtectionContainerValue", "replicationProtectedItemValue", "recoveryPointValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationProtectionContainers/replicationProtectionContainerValue/replicationProtectedItems/replicationProtectedItemValue/recoveryPoints/recoveryPointValue"
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationProtectionContainers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationProtectionContainers/replicationProtectionContainerValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationProtectionContainers/replicationProtectionContainerValue/replicationProtectedItems",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationProtectionContainers/replicationProtectionContainerValue/replicationProtectedItems/replicationProtectedItemValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationProtectionContainers/replicationProtectionContainerValue/replicationProtectedItems/replicationProtectedItemValue/recoveryPoints",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationProtectionContainers/replicationProtectionContainerValue/replicationProtectedItems/replicationProtectedItemValue/recoveryPoints/recoveryPointValue",
			Expected: &RecoveryPointId{
				SubscriptionId:                     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                  "example-resource-group",
				VaultName:                          "vaultValue",
				ReplicationFabricName:              "replicationFabricValue",
				ReplicationProtectionContainerName: "replicationProtectionContainerValue",
				ReplicationProtectedItemName:       "replicationProtectedItemValue",
				RecoveryPointName:                  "recoveryPointValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationProtectionContainers/replicationProtectionContainerValue/replicationProtectedItems/replicationProtectedItemValue/recoveryPoints/recoveryPointValue/extra",
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

		if actual.ReplicationFabricName != v.Expected.ReplicationFabricName {
			t.Fatalf("Expected %q but got %q for ReplicationFabricName", v.Expected.ReplicationFabricName, actual.ReplicationFabricName)
		}

		if actual.ReplicationProtectionContainerName != v.Expected.ReplicationProtectionContainerName {
			t.Fatalf("Expected %q but got %q for ReplicationProtectionContainerName", v.Expected.ReplicationProtectionContainerName, actual.ReplicationProtectionContainerName)
		}

		if actual.ReplicationProtectedItemName != v.Expected.ReplicationProtectedItemName {
			t.Fatalf("Expected %q but got %q for ReplicationProtectedItemName", v.Expected.ReplicationProtectedItemName, actual.ReplicationProtectedItemName)
		}

		if actual.RecoveryPointName != v.Expected.RecoveryPointName {
			t.Fatalf("Expected %q but got %q for RecoveryPointName", v.Expected.RecoveryPointName, actual.RecoveryPointName)
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationProtectionContainers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtVaLuE/rEpLiCaTiOnFaBrIcS/rEpLiCaTiOnFaBrIcVaLuE/rEpLiCaTiOnPrOtEcTiOnCoNtAiNeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationProtectionContainers/replicationProtectionContainerValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtVaLuE/rEpLiCaTiOnFaBrIcS/rEpLiCaTiOnFaBrIcVaLuE/rEpLiCaTiOnPrOtEcTiOnCoNtAiNeRs/rEpLiCaTiOnPrOtEcTiOnCoNtAiNeRvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationProtectionContainers/replicationProtectionContainerValue/replicationProtectedItems",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtVaLuE/rEpLiCaTiOnFaBrIcS/rEpLiCaTiOnFaBrIcVaLuE/rEpLiCaTiOnPrOtEcTiOnCoNtAiNeRs/rEpLiCaTiOnPrOtEcTiOnCoNtAiNeRvAlUe/rEpLiCaTiOnPrOtEcTeDiTeMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationProtectionContainers/replicationProtectionContainerValue/replicationProtectedItems/replicationProtectedItemValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtVaLuE/rEpLiCaTiOnFaBrIcS/rEpLiCaTiOnFaBrIcVaLuE/rEpLiCaTiOnPrOtEcTiOnCoNtAiNeRs/rEpLiCaTiOnPrOtEcTiOnCoNtAiNeRvAlUe/rEpLiCaTiOnPrOtEcTeDiTeMs/rEpLiCaTiOnPrOtEcTeDiTeMvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationProtectionContainers/replicationProtectionContainerValue/replicationProtectedItems/replicationProtectedItemValue/recoveryPoints",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtVaLuE/rEpLiCaTiOnFaBrIcS/rEpLiCaTiOnFaBrIcVaLuE/rEpLiCaTiOnPrOtEcTiOnCoNtAiNeRs/rEpLiCaTiOnPrOtEcTiOnCoNtAiNeRvAlUe/rEpLiCaTiOnPrOtEcTeDiTeMs/rEpLiCaTiOnPrOtEcTeDiTeMvAlUe/rEcOvErYpOiNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationProtectionContainers/replicationProtectionContainerValue/replicationProtectedItems/replicationProtectedItemValue/recoveryPoints/recoveryPointValue",
			Expected: &RecoveryPointId{
				SubscriptionId:                     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                  "example-resource-group",
				VaultName:                          "vaultValue",
				ReplicationFabricName:              "replicationFabricValue",
				ReplicationProtectionContainerName: "replicationProtectionContainerValue",
				ReplicationProtectedItemName:       "replicationProtectedItemValue",
				RecoveryPointName:                  "recoveryPointValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/vaultValue/replicationFabrics/replicationFabricValue/replicationProtectionContainers/replicationProtectionContainerValue/replicationProtectedItems/replicationProtectedItemValue/recoveryPoints/recoveryPointValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtVaLuE/rEpLiCaTiOnFaBrIcS/rEpLiCaTiOnFaBrIcVaLuE/rEpLiCaTiOnPrOtEcTiOnCoNtAiNeRs/rEpLiCaTiOnPrOtEcTiOnCoNtAiNeRvAlUe/rEpLiCaTiOnPrOtEcTeDiTeMs/rEpLiCaTiOnPrOtEcTeDiTeMvAlUe/rEcOvErYpOiNtS/rEcOvErYpOiNtVaLuE",
			Expected: &RecoveryPointId{
				SubscriptionId:                     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                  "eXaMpLe-rEsOuRcE-GrOuP",
				VaultName:                          "vAuLtVaLuE",
				ReplicationFabricName:              "rEpLiCaTiOnFaBrIcVaLuE",
				ReplicationProtectionContainerName: "rEpLiCaTiOnPrOtEcTiOnCoNtAiNeRvAlUe",
				ReplicationProtectedItemName:       "rEpLiCaTiOnPrOtEcTeDiTeMvAlUe",
				RecoveryPointName:                  "rEcOvErYpOiNtVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/vAuLtVaLuE/rEpLiCaTiOnFaBrIcS/rEpLiCaTiOnFaBrIcVaLuE/rEpLiCaTiOnPrOtEcTiOnCoNtAiNeRs/rEpLiCaTiOnPrOtEcTiOnCoNtAiNeRvAlUe/rEpLiCaTiOnPrOtEcTeDiTeMs/rEpLiCaTiOnPrOtEcTeDiTeMvAlUe/rEcOvErYpOiNtS/rEcOvErYpOiNtVaLuE/extra",
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

		if actual.ReplicationFabricName != v.Expected.ReplicationFabricName {
			t.Fatalf("Expected %q but got %q for ReplicationFabricName", v.Expected.ReplicationFabricName, actual.ReplicationFabricName)
		}

		if actual.ReplicationProtectionContainerName != v.Expected.ReplicationProtectionContainerName {
			t.Fatalf("Expected %q but got %q for ReplicationProtectionContainerName", v.Expected.ReplicationProtectionContainerName, actual.ReplicationProtectionContainerName)
		}

		if actual.ReplicationProtectedItemName != v.Expected.ReplicationProtectedItemName {
			t.Fatalf("Expected %q but got %q for ReplicationProtectedItemName", v.Expected.ReplicationProtectedItemName, actual.ReplicationProtectedItemName)
		}

		if actual.RecoveryPointName != v.Expected.RecoveryPointName {
			t.Fatalf("Expected %q but got %q for RecoveryPointName", v.Expected.RecoveryPointName, actual.RecoveryPointName)
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
