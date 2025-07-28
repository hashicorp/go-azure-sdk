package deletedbackupinstances

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeletedBackupInstanceId{}

func TestNewDeletedBackupInstanceID(t *testing.T) {
	id := NewDeletedBackupInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "backupVaultName", "deletedBackupInstanceName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.BackupVaultName != "backupVaultName" {
		t.Fatalf("Expected %q but got %q for Segment 'BackupVaultName'", id.BackupVaultName, "backupVaultName")
	}

	if id.DeletedBackupInstanceName != "deletedBackupInstanceName" {
		t.Fatalf("Expected %q but got %q for Segment 'DeletedBackupInstanceName'", id.DeletedBackupInstanceName, "deletedBackupInstanceName")
	}
}

func TestFormatDeletedBackupInstanceID(t *testing.T) {
	actual := NewDeletedBackupInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "backupVaultName", "deletedBackupInstanceName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/backupVaults/backupVaultName/deletedBackupInstances/deletedBackupInstanceName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeletedBackupInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeletedBackupInstanceId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/backupVaults",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/backupVaults/backupVaultName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/backupVaults/backupVaultName/deletedBackupInstances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/backupVaults/backupVaultName/deletedBackupInstances/deletedBackupInstanceName",
			Expected: &DeletedBackupInstanceId{
				SubscriptionId:            "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:         "example-resource-group",
				BackupVaultName:           "backupVaultName",
				DeletedBackupInstanceName: "deletedBackupInstanceName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/backupVaults/backupVaultName/deletedBackupInstances/deletedBackupInstanceName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeletedBackupInstanceID(v.Input)
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

		if actual.BackupVaultName != v.Expected.BackupVaultName {
			t.Fatalf("Expected %q but got %q for BackupVaultName", v.Expected.BackupVaultName, actual.BackupVaultName)
		}

		if actual.DeletedBackupInstanceName != v.Expected.DeletedBackupInstanceName {
			t.Fatalf("Expected %q but got %q for DeletedBackupInstanceName", v.Expected.DeletedBackupInstanceName, actual.DeletedBackupInstanceName)
		}

	}
}

func TestParseDeletedBackupInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeletedBackupInstanceId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtApRoTeCtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/backupVaults",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtApRoTeCtIoN/bAcKuPvAuLtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/backupVaults/backupVaultName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtApRoTeCtIoN/bAcKuPvAuLtS/bAcKuPvAuLtNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/backupVaults/backupVaultName/deletedBackupInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtApRoTeCtIoN/bAcKuPvAuLtS/bAcKuPvAuLtNaMe/dElEtEdBaCkUpInStAnCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/backupVaults/backupVaultName/deletedBackupInstances/deletedBackupInstanceName",
			Expected: &DeletedBackupInstanceId{
				SubscriptionId:            "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:         "example-resource-group",
				BackupVaultName:           "backupVaultName",
				DeletedBackupInstanceName: "deletedBackupInstanceName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/backupVaults/backupVaultName/deletedBackupInstances/deletedBackupInstanceName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtApRoTeCtIoN/bAcKuPvAuLtS/bAcKuPvAuLtNaMe/dElEtEdBaCkUpInStAnCeS/dElEtEdBaCkUpInStAnCeNaMe",
			Expected: &DeletedBackupInstanceId{
				SubscriptionId:            "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:         "eXaMpLe-rEsOuRcE-GrOuP",
				BackupVaultName:           "bAcKuPvAuLtNaMe",
				DeletedBackupInstanceName: "dElEtEdBaCkUpInStAnCeNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtApRoTeCtIoN/bAcKuPvAuLtS/bAcKuPvAuLtNaMe/dElEtEdBaCkUpInStAnCeS/dElEtEdBaCkUpInStAnCeNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeletedBackupInstanceIDInsensitively(v.Input)
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

		if actual.BackupVaultName != v.Expected.BackupVaultName {
			t.Fatalf("Expected %q but got %q for BackupVaultName", v.Expected.BackupVaultName, actual.BackupVaultName)
		}

		if actual.DeletedBackupInstanceName != v.Expected.DeletedBackupInstanceName {
			t.Fatalf("Expected %q but got %q for DeletedBackupInstanceName", v.Expected.DeletedBackupInstanceName, actual.DeletedBackupInstanceName)
		}

	}
}

func TestSegmentsForDeletedBackupInstanceId(t *testing.T) {
	segments := DeletedBackupInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeletedBackupInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
