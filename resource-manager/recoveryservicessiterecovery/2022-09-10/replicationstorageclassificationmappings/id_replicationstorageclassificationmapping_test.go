package replicationstorageclassificationmappings

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ReplicationStorageClassificationMappingId{}

func TestNewReplicationStorageClassificationMappingID(t *testing.T) {
	id := NewReplicationStorageClassificationMappingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue", "fabricValue", "storageClassificationValue", "storageClassificationMappingValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.ResourceName != "resourceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceName'", id.ResourceName, "resourceValue")
	}

	if id.FabricName != "fabricValue" {
		t.Fatalf("Expected %q but got %q for Segment 'FabricName'", id.FabricName, "fabricValue")
	}

	if id.StorageClassificationName != "storageClassificationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'StorageClassificationName'", id.StorageClassificationName, "storageClassificationValue")
	}

	if id.StorageClassificationMappingName != "storageClassificationMappingValue" {
		t.Fatalf("Expected %q but got %q for Segment 'StorageClassificationMappingName'", id.StorageClassificationMappingName, "storageClassificationMappingValue")
	}
}

func TestFormatReplicationStorageClassificationMappingID(t *testing.T) {
	actual := NewReplicationStorageClassificationMappingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue", "fabricValue", "storageClassificationValue", "storageClassificationMappingValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceValue/replicationFabrics/fabricValue/replicationStorageClassifications/storageClassificationValue/replicationStorageClassificationMappings/storageClassificationMappingValue"
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceValue/replicationFabrics",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceValue/replicationFabrics/fabricValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceValue/replicationFabrics/fabricValue/replicationStorageClassifications",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceValue/replicationFabrics/fabricValue/replicationStorageClassifications/storageClassificationValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceValue/replicationFabrics/fabricValue/replicationStorageClassifications/storageClassificationValue/replicationStorageClassificationMappings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceValue/replicationFabrics/fabricValue/replicationStorageClassifications/storageClassificationValue/replicationStorageClassificationMappings/storageClassificationMappingValue",
			Expected: &ReplicationStorageClassificationMappingId{
				SubscriptionId:                   "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                "example-resource-group",
				ResourceName:                     "resourceValue",
				FabricName:                       "fabricValue",
				StorageClassificationName:        "storageClassificationValue",
				StorageClassificationMappingName: "storageClassificationMappingValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceValue/replicationFabrics/fabricValue/replicationStorageClassifications/storageClassificationValue/replicationStorageClassificationMappings/storageClassificationMappingValue/extra",
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

		if actual.ResourceName != v.Expected.ResourceName {
			t.Fatalf("Expected %q but got %q for ResourceName", v.Expected.ResourceName, actual.ResourceName)
		}

		if actual.FabricName != v.Expected.FabricName {
			t.Fatalf("Expected %q but got %q for FabricName", v.Expected.FabricName, actual.FabricName)
		}

		if actual.StorageClassificationName != v.Expected.StorageClassificationName {
			t.Fatalf("Expected %q but got %q for StorageClassificationName", v.Expected.StorageClassificationName, actual.StorageClassificationName)
		}

		if actual.StorageClassificationMappingName != v.Expected.StorageClassificationMappingName {
			t.Fatalf("Expected %q but got %q for StorageClassificationMappingName", v.Expected.StorageClassificationMappingName, actual.StorageClassificationMappingName)
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/rEsOuRcEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceValue/replicationFabrics",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/rEsOuRcEvAlUe/rEpLiCaTiOnFaBrIcS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceValue/replicationFabrics/fabricValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/rEsOuRcEvAlUe/rEpLiCaTiOnFaBrIcS/fAbRiCvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceValue/replicationFabrics/fabricValue/replicationStorageClassifications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/rEsOuRcEvAlUe/rEpLiCaTiOnFaBrIcS/fAbRiCvAlUe/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceValue/replicationFabrics/fabricValue/replicationStorageClassifications/storageClassificationValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/rEsOuRcEvAlUe/rEpLiCaTiOnFaBrIcS/fAbRiCvAlUe/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNs/sToRaGeClAsSiFiCaTiOnVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceValue/replicationFabrics/fabricValue/replicationStorageClassifications/storageClassificationValue/replicationStorageClassificationMappings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/rEsOuRcEvAlUe/rEpLiCaTiOnFaBrIcS/fAbRiCvAlUe/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNs/sToRaGeClAsSiFiCaTiOnVaLuE/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNmApPiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceValue/replicationFabrics/fabricValue/replicationStorageClassifications/storageClassificationValue/replicationStorageClassificationMappings/storageClassificationMappingValue",
			Expected: &ReplicationStorageClassificationMappingId{
				SubscriptionId:                   "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                "example-resource-group",
				ResourceName:                     "resourceValue",
				FabricName:                       "fabricValue",
				StorageClassificationName:        "storageClassificationValue",
				StorageClassificationMappingName: "storageClassificationMappingValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RecoveryServices/vaults/resourceValue/replicationFabrics/fabricValue/replicationStorageClassifications/storageClassificationValue/replicationStorageClassificationMappings/storageClassificationMappingValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/rEsOuRcEvAlUe/rEpLiCaTiOnFaBrIcS/fAbRiCvAlUe/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNs/sToRaGeClAsSiFiCaTiOnVaLuE/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNmApPiNgS/sToRaGeClAsSiFiCaTiOnMaPpInGvAlUe",
			Expected: &ReplicationStorageClassificationMappingId{
				SubscriptionId:                   "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                "eXaMpLe-rEsOuRcE-GrOuP",
				ResourceName:                     "rEsOuRcEvAlUe",
				FabricName:                       "fAbRiCvAlUe",
				StorageClassificationName:        "sToRaGeClAsSiFiCaTiOnVaLuE",
				StorageClassificationMappingName: "sToRaGeClAsSiFiCaTiOnMaPpInGvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEcOvErYsErViCeS/vAuLtS/rEsOuRcEvAlUe/rEpLiCaTiOnFaBrIcS/fAbRiCvAlUe/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNs/sToRaGeClAsSiFiCaTiOnVaLuE/rEpLiCaTiOnStOrAgEcLaSsIfIcAtIoNmApPiNgS/sToRaGeClAsSiFiCaTiOnMaPpInGvAlUe/extra",
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

		if actual.ResourceName != v.Expected.ResourceName {
			t.Fatalf("Expected %q but got %q for ResourceName", v.Expected.ResourceName, actual.ResourceName)
		}

		if actual.FabricName != v.Expected.FabricName {
			t.Fatalf("Expected %q but got %q for FabricName", v.Expected.FabricName, actual.FabricName)
		}

		if actual.StorageClassificationName != v.Expected.StorageClassificationName {
			t.Fatalf("Expected %q but got %q for StorageClassificationName", v.Expected.StorageClassificationName, actual.StorageClassificationName)
		}

		if actual.StorageClassificationMappingName != v.Expected.StorageClassificationMappingName {
			t.Fatalf("Expected %q but got %q for StorageClassificationMappingName", v.Expected.StorageClassificationMappingName, actual.StorageClassificationMappingName)
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
