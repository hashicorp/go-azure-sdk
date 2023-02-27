// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedprivateendpoints

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ManagedPrivateEndpointId{}

func TestNewManagedPrivateEndpointID(t *testing.T) {
	id := NewManagedPrivateEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "managedVirtualNetworkValue", "managedPrivateEndpointValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.FactoryName != "factoryValue" {
		t.Fatalf("Expected %q but got %q for Segment 'FactoryName'", id.FactoryName, "factoryValue")
	}

	if id.ManagedVirtualNetworkName != "managedVirtualNetworkValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedVirtualNetworkName'", id.ManagedVirtualNetworkName, "managedVirtualNetworkValue")
	}

	if id.ManagedPrivateEndpointName != "managedPrivateEndpointValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedPrivateEndpointName'", id.ManagedPrivateEndpointName, "managedPrivateEndpointValue")
	}
}

func TestFormatManagedPrivateEndpointID(t *testing.T) {
	actual := NewManagedPrivateEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "managedVirtualNetworkValue", "managedPrivateEndpointValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataFactory/factories/factoryValue/managedVirtualNetworks/managedVirtualNetworkValue/managedPrivateEndpoints/managedPrivateEndpointValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseManagedPrivateEndpointID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ManagedPrivateEndpointId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataFactory",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataFactory/factories",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataFactory/factories/factoryValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataFactory/factories/factoryValue/managedVirtualNetworks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataFactory/factories/factoryValue/managedVirtualNetworks/managedVirtualNetworkValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataFactory/factories/factoryValue/managedVirtualNetworks/managedVirtualNetworkValue/managedPrivateEndpoints",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataFactory/factories/factoryValue/managedVirtualNetworks/managedVirtualNetworkValue/managedPrivateEndpoints/managedPrivateEndpointValue",
			Expected: &ManagedPrivateEndpointId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "example-resource-group",
				FactoryName:                "factoryValue",
				ManagedVirtualNetworkName:  "managedVirtualNetworkValue",
				ManagedPrivateEndpointName: "managedPrivateEndpointValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataFactory/factories/factoryValue/managedVirtualNetworks/managedVirtualNetworkValue/managedPrivateEndpoints/managedPrivateEndpointValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseManagedPrivateEndpointID(v.Input)
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

		if actual.FactoryName != v.Expected.FactoryName {
			t.Fatalf("Expected %q but got %q for FactoryName", v.Expected.FactoryName, actual.FactoryName)
		}

		if actual.ManagedVirtualNetworkName != v.Expected.ManagedVirtualNetworkName {
			t.Fatalf("Expected %q but got %q for ManagedVirtualNetworkName", v.Expected.ManagedVirtualNetworkName, actual.ManagedVirtualNetworkName)
		}

		if actual.ManagedPrivateEndpointName != v.Expected.ManagedPrivateEndpointName {
			t.Fatalf("Expected %q but got %q for ManagedPrivateEndpointName", v.Expected.ManagedPrivateEndpointName, actual.ManagedPrivateEndpointName)
		}

	}
}

func TestParseManagedPrivateEndpointIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ManagedPrivateEndpointId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataFactory",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtAfAcToRy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataFactory/factories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtAfAcToRy/fAcToRiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataFactory/factories/factoryValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtAfAcToRy/fAcToRiEs/fAcToRyVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataFactory/factories/factoryValue/managedVirtualNetworks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtAfAcToRy/fAcToRiEs/fAcToRyVaLuE/mAnAgEdViRtUaLnEtWoRkS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataFactory/factories/factoryValue/managedVirtualNetworks/managedVirtualNetworkValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtAfAcToRy/fAcToRiEs/fAcToRyVaLuE/mAnAgEdViRtUaLnEtWoRkS/mAnAgEdViRtUaLnEtWoRkVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataFactory/factories/factoryValue/managedVirtualNetworks/managedVirtualNetworkValue/managedPrivateEndpoints",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtAfAcToRy/fAcToRiEs/fAcToRyVaLuE/mAnAgEdViRtUaLnEtWoRkS/mAnAgEdViRtUaLnEtWoRkVaLuE/mAnAgEdPrIvAtEeNdPoInTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataFactory/factories/factoryValue/managedVirtualNetworks/managedVirtualNetworkValue/managedPrivateEndpoints/managedPrivateEndpointValue",
			Expected: &ManagedPrivateEndpointId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "example-resource-group",
				FactoryName:                "factoryValue",
				ManagedVirtualNetworkName:  "managedVirtualNetworkValue",
				ManagedPrivateEndpointName: "managedPrivateEndpointValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataFactory/factories/factoryValue/managedVirtualNetworks/managedVirtualNetworkValue/managedPrivateEndpoints/managedPrivateEndpointValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtAfAcToRy/fAcToRiEs/fAcToRyVaLuE/mAnAgEdViRtUaLnEtWoRkS/mAnAgEdViRtUaLnEtWoRkVaLuE/mAnAgEdPrIvAtEeNdPoInTs/mAnAgEdPrIvAtEeNdPoInTvAlUe",
			Expected: &ManagedPrivateEndpointId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "eXaMpLe-rEsOuRcE-GrOuP",
				FactoryName:                "fAcToRyVaLuE",
				ManagedVirtualNetworkName:  "mAnAgEdViRtUaLnEtWoRkVaLuE",
				ManagedPrivateEndpointName: "mAnAgEdPrIvAtEeNdPoInTvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtAfAcToRy/fAcToRiEs/fAcToRyVaLuE/mAnAgEdViRtUaLnEtWoRkS/mAnAgEdViRtUaLnEtWoRkVaLuE/mAnAgEdPrIvAtEeNdPoInTs/mAnAgEdPrIvAtEeNdPoInTvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseManagedPrivateEndpointIDInsensitively(v.Input)
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

		if actual.FactoryName != v.Expected.FactoryName {
			t.Fatalf("Expected %q but got %q for FactoryName", v.Expected.FactoryName, actual.FactoryName)
		}

		if actual.ManagedVirtualNetworkName != v.Expected.ManagedVirtualNetworkName {
			t.Fatalf("Expected %q but got %q for ManagedVirtualNetworkName", v.Expected.ManagedVirtualNetworkName, actual.ManagedVirtualNetworkName)
		}

		if actual.ManagedPrivateEndpointName != v.Expected.ManagedPrivateEndpointName {
			t.Fatalf("Expected %q but got %q for ManagedPrivateEndpointName", v.Expected.ManagedPrivateEndpointName, actual.ManagedPrivateEndpointName)
		}

	}
}

func TestSegmentsForManagedPrivateEndpointId(t *testing.T) {
	segments := ManagedPrivateEndpointId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ManagedPrivateEndpointId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
