package serverendpointresource

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServerEndpointId{}

func TestNewServerEndpointID(t *testing.T) {
	id := NewServerEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName", "serverEndpointName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.StorageSyncServiceName != "storageSyncServiceName" {
		t.Fatalf("Expected %q but got %q for Segment 'StorageSyncServiceName'", id.StorageSyncServiceName, "storageSyncServiceName")
	}

	if id.SyncGroupName != "syncGroupName" {
		t.Fatalf("Expected %q but got %q for Segment 'SyncGroupName'", id.SyncGroupName, "syncGroupName")
	}

	if id.ServerEndpointName != "serverEndpointName" {
		t.Fatalf("Expected %q but got %q for Segment 'ServerEndpointName'", id.ServerEndpointName, "serverEndpointName")
	}
}

func TestFormatServerEndpointID(t *testing.T) {
	actual := NewServerEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName", "serverEndpointName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageSync/storageSyncServices/storageSyncServiceName/syncGroups/syncGroupName/serverEndpoints/serverEndpointName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServerEndpointID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServerEndpointId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageSync",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageSync/storageSyncServices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageSync/storageSyncServices/storageSyncServiceName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageSync/storageSyncServices/storageSyncServiceName/syncGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageSync/storageSyncServices/storageSyncServiceName/syncGroups/syncGroupName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageSync/storageSyncServices/storageSyncServiceName/syncGroups/syncGroupName/serverEndpoints",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageSync/storageSyncServices/storageSyncServiceName/syncGroups/syncGroupName/serverEndpoints/serverEndpointName",
			Expected: &ServerEndpointId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:      "example-resource-group",
				StorageSyncServiceName: "storageSyncServiceName",
				SyncGroupName:          "syncGroupName",
				ServerEndpointName:     "serverEndpointName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageSync/storageSyncServices/storageSyncServiceName/syncGroups/syncGroupName/serverEndpoints/serverEndpointName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServerEndpointID(v.Input)
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

		if actual.StorageSyncServiceName != v.Expected.StorageSyncServiceName {
			t.Fatalf("Expected %q but got %q for StorageSyncServiceName", v.Expected.StorageSyncServiceName, actual.StorageSyncServiceName)
		}

		if actual.SyncGroupName != v.Expected.SyncGroupName {
			t.Fatalf("Expected %q but got %q for SyncGroupName", v.Expected.SyncGroupName, actual.SyncGroupName)
		}

		if actual.ServerEndpointName != v.Expected.ServerEndpointName {
			t.Fatalf("Expected %q but got %q for ServerEndpointName", v.Expected.ServerEndpointName, actual.ServerEndpointName)
		}

	}
}

func TestParseServerEndpointIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServerEndpointId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageSync",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sToRaGeSyNc",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageSync/storageSyncServices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sToRaGeSyNc/sToRaGeSyNcSeRvIcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageSync/storageSyncServices/storageSyncServiceName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sToRaGeSyNc/sToRaGeSyNcSeRvIcEs/sToRaGeSyNcSeRvIcEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageSync/storageSyncServices/storageSyncServiceName/syncGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sToRaGeSyNc/sToRaGeSyNcSeRvIcEs/sToRaGeSyNcSeRvIcEnAmE/sYnCgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageSync/storageSyncServices/storageSyncServiceName/syncGroups/syncGroupName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sToRaGeSyNc/sToRaGeSyNcSeRvIcEs/sToRaGeSyNcSeRvIcEnAmE/sYnCgRoUpS/sYnCgRoUpNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageSync/storageSyncServices/storageSyncServiceName/syncGroups/syncGroupName/serverEndpoints",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sToRaGeSyNc/sToRaGeSyNcSeRvIcEs/sToRaGeSyNcSeRvIcEnAmE/sYnCgRoUpS/sYnCgRoUpNaMe/sErVeReNdPoInTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageSync/storageSyncServices/storageSyncServiceName/syncGroups/syncGroupName/serverEndpoints/serverEndpointName",
			Expected: &ServerEndpointId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:      "example-resource-group",
				StorageSyncServiceName: "storageSyncServiceName",
				SyncGroupName:          "syncGroupName",
				ServerEndpointName:     "serverEndpointName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageSync/storageSyncServices/storageSyncServiceName/syncGroups/syncGroupName/serverEndpoints/serverEndpointName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sToRaGeSyNc/sToRaGeSyNcSeRvIcEs/sToRaGeSyNcSeRvIcEnAmE/sYnCgRoUpS/sYnCgRoUpNaMe/sErVeReNdPoInTs/sErVeReNdPoInTnAmE",
			Expected: &ServerEndpointId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:      "eXaMpLe-rEsOuRcE-GrOuP",
				StorageSyncServiceName: "sToRaGeSyNcSeRvIcEnAmE",
				SyncGroupName:          "sYnCgRoUpNaMe",
				ServerEndpointName:     "sErVeReNdPoInTnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sToRaGeSyNc/sToRaGeSyNcSeRvIcEs/sToRaGeSyNcSeRvIcEnAmE/sYnCgRoUpS/sYnCgRoUpNaMe/sErVeReNdPoInTs/sErVeReNdPoInTnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServerEndpointIDInsensitively(v.Input)
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

		if actual.StorageSyncServiceName != v.Expected.StorageSyncServiceName {
			t.Fatalf("Expected %q but got %q for StorageSyncServiceName", v.Expected.StorageSyncServiceName, actual.StorageSyncServiceName)
		}

		if actual.SyncGroupName != v.Expected.SyncGroupName {
			t.Fatalf("Expected %q but got %q for SyncGroupName", v.Expected.SyncGroupName, actual.SyncGroupName)
		}

		if actual.ServerEndpointName != v.Expected.ServerEndpointName {
			t.Fatalf("Expected %q but got %q for ServerEndpointName", v.Expected.ServerEndpointName, actual.ServerEndpointName)
		}

	}
}

func TestSegmentsForServerEndpointId(t *testing.T) {
	segments := ServerEndpointId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServerEndpointId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
