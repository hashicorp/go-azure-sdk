package containers

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ContainerId{}

func TestNewContainerID(t *testing.T) {
	id := NewContainerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "storageAccountValue", "containerValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.DataBoxEdgeDeviceName != "dataBoxEdgeDeviceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DataBoxEdgeDeviceName'", id.DataBoxEdgeDeviceName, "dataBoxEdgeDeviceValue")
	}

	if id.StorageAccountName != "storageAccountValue" {
		t.Fatalf("Expected %q but got %q for Segment 'StorageAccountName'", id.StorageAccountName, "storageAccountValue")
	}

	if id.ContainerName != "containerValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ContainerName'", id.ContainerName, "containerValue")
	}
}

func TestFormatContainerID(t *testing.T) {
	actual := NewContainerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "storageAccountValue", "containerValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceValue/storageAccounts/storageAccountValue/containers/containerValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseContainerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ContainerId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceValue/storageAccounts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceValue/storageAccounts/storageAccountValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceValue/storageAccounts/storageAccountValue/containers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceValue/storageAccounts/storageAccountValue/containers/containerValue",
			Expected: &ContainerId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "example-resource-group",
				DataBoxEdgeDeviceName: "dataBoxEdgeDeviceValue",
				StorageAccountName:    "storageAccountValue",
				ContainerName:         "containerValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceValue/storageAccounts/storageAccountValue/containers/containerValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseContainerID(v.Input)
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

		if actual.DataBoxEdgeDeviceName != v.Expected.DataBoxEdgeDeviceName {
			t.Fatalf("Expected %q but got %q for DataBoxEdgeDeviceName", v.Expected.DataBoxEdgeDeviceName, actual.DataBoxEdgeDeviceName)
		}

		if actual.StorageAccountName != v.Expected.StorageAccountName {
			t.Fatalf("Expected %q but got %q for StorageAccountName", v.Expected.StorageAccountName, actual.StorageAccountName)
		}

		if actual.ContainerName != v.Expected.ContainerName {
			t.Fatalf("Expected %q but got %q for ContainerName", v.Expected.ContainerName, actual.ContainerName)
		}

	}
}

func TestParseContainerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ContainerId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtAbOxEdGe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtAbOxEdGe/dAtAbOxEdGeDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtAbOxEdGe/dAtAbOxEdGeDeViCeS/dAtAbOxEdGeDeViCeVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceValue/storageAccounts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtAbOxEdGe/dAtAbOxEdGeDeViCeS/dAtAbOxEdGeDeViCeVaLuE/sToRaGeAcCoUnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceValue/storageAccounts/storageAccountValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtAbOxEdGe/dAtAbOxEdGeDeViCeS/dAtAbOxEdGeDeViCeVaLuE/sToRaGeAcCoUnTs/sToRaGeAcCoUnTvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceValue/storageAccounts/storageAccountValue/containers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtAbOxEdGe/dAtAbOxEdGeDeViCeS/dAtAbOxEdGeDeViCeVaLuE/sToRaGeAcCoUnTs/sToRaGeAcCoUnTvAlUe/cOnTaInErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceValue/storageAccounts/storageAccountValue/containers/containerValue",
			Expected: &ContainerId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "example-resource-group",
				DataBoxEdgeDeviceName: "dataBoxEdgeDeviceValue",
				StorageAccountName:    "storageAccountValue",
				ContainerName:         "containerValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceValue/storageAccounts/storageAccountValue/containers/containerValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtAbOxEdGe/dAtAbOxEdGeDeViCeS/dAtAbOxEdGeDeViCeVaLuE/sToRaGeAcCoUnTs/sToRaGeAcCoUnTvAlUe/cOnTaInErS/cOnTaInErVaLuE",
			Expected: &ContainerId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "eXaMpLe-rEsOuRcE-GrOuP",
				DataBoxEdgeDeviceName: "dAtAbOxEdGeDeViCeVaLuE",
				StorageAccountName:    "sToRaGeAcCoUnTvAlUe",
				ContainerName:         "cOnTaInErVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtAbOxEdGe/dAtAbOxEdGeDeViCeS/dAtAbOxEdGeDeViCeVaLuE/sToRaGeAcCoUnTs/sToRaGeAcCoUnTvAlUe/cOnTaInErS/cOnTaInErVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseContainerIDInsensitively(v.Input)
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

		if actual.DataBoxEdgeDeviceName != v.Expected.DataBoxEdgeDeviceName {
			t.Fatalf("Expected %q but got %q for DataBoxEdgeDeviceName", v.Expected.DataBoxEdgeDeviceName, actual.DataBoxEdgeDeviceName)
		}

		if actual.StorageAccountName != v.Expected.StorageAccountName {
			t.Fatalf("Expected %q but got %q for StorageAccountName", v.Expected.StorageAccountName, actual.StorageAccountName)
		}

		if actual.ContainerName != v.Expected.ContainerName {
			t.Fatalf("Expected %q but got %q for ContainerName", v.Expected.ContainerName, actual.ContainerName)
		}

	}
}

func TestSegmentsForContainerId(t *testing.T) {
	segments := ContainerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ContainerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
