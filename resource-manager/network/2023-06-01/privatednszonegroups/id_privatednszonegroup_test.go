package privatednszonegroups

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PrivateDnsZoneGroupId{}

func TestNewPrivateDnsZoneGroupID(t *testing.T) {
	id := NewPrivateDnsZoneGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateEndpointValue", "privateDnsZoneGroupValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.PrivateEndpointName != "privateEndpointValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PrivateEndpointName'", id.PrivateEndpointName, "privateEndpointValue")
	}

	if id.PrivateDnsZoneGroupName != "privateDnsZoneGroupValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PrivateDnsZoneGroupName'", id.PrivateDnsZoneGroupName, "privateDnsZoneGroupValue")
	}
}

func TestFormatPrivateDnsZoneGroupID(t *testing.T) {
	actual := NewPrivateDnsZoneGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateEndpointValue", "privateDnsZoneGroupValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/privateEndpoints/privateEndpointValue/privateDnsZoneGroups/privateDnsZoneGroupValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePrivateDnsZoneGroupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PrivateDnsZoneGroupId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/privateEndpoints",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/privateEndpoints/privateEndpointValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/privateEndpoints/privateEndpointValue/privateDnsZoneGroups",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/privateEndpoints/privateEndpointValue/privateDnsZoneGroups/privateDnsZoneGroupValue",
			Expected: &PrivateDnsZoneGroupId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:       "example-resource-group",
				PrivateEndpointName:     "privateEndpointValue",
				PrivateDnsZoneGroupName: "privateDnsZoneGroupValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/privateEndpoints/privateEndpointValue/privateDnsZoneGroups/privateDnsZoneGroupValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePrivateDnsZoneGroupID(v.Input)
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

		if actual.PrivateEndpointName != v.Expected.PrivateEndpointName {
			t.Fatalf("Expected %q but got %q for PrivateEndpointName", v.Expected.PrivateEndpointName, actual.PrivateEndpointName)
		}

		if actual.PrivateDnsZoneGroupName != v.Expected.PrivateDnsZoneGroupName {
			t.Fatalf("Expected %q but got %q for PrivateDnsZoneGroupName", v.Expected.PrivateDnsZoneGroupName, actual.PrivateDnsZoneGroupName)
		}

	}
}

func TestParsePrivateDnsZoneGroupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PrivateDnsZoneGroupId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/privateEndpoints",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/pRiVaTeEnDpOiNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/privateEndpoints/privateEndpointValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/pRiVaTeEnDpOiNtS/pRiVaTeEnDpOiNtVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/privateEndpoints/privateEndpointValue/privateDnsZoneGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/pRiVaTeEnDpOiNtS/pRiVaTeEnDpOiNtVaLuE/pRiVaTeDnSzOnEgRoUpS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/privateEndpoints/privateEndpointValue/privateDnsZoneGroups/privateDnsZoneGroupValue",
			Expected: &PrivateDnsZoneGroupId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:       "example-resource-group",
				PrivateEndpointName:     "privateEndpointValue",
				PrivateDnsZoneGroupName: "privateDnsZoneGroupValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/privateEndpoints/privateEndpointValue/privateDnsZoneGroups/privateDnsZoneGroupValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/pRiVaTeEnDpOiNtS/pRiVaTeEnDpOiNtVaLuE/pRiVaTeDnSzOnEgRoUpS/pRiVaTeDnSzOnEgRoUpVaLuE",
			Expected: &PrivateDnsZoneGroupId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:       "eXaMpLe-rEsOuRcE-GrOuP",
				PrivateEndpointName:     "pRiVaTeEnDpOiNtVaLuE",
				PrivateDnsZoneGroupName: "pRiVaTeDnSzOnEgRoUpVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/pRiVaTeEnDpOiNtS/pRiVaTeEnDpOiNtVaLuE/pRiVaTeDnSzOnEgRoUpS/pRiVaTeDnSzOnEgRoUpVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePrivateDnsZoneGroupIDInsensitively(v.Input)
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

		if actual.PrivateEndpointName != v.Expected.PrivateEndpointName {
			t.Fatalf("Expected %q but got %q for PrivateEndpointName", v.Expected.PrivateEndpointName, actual.PrivateEndpointName)
		}

		if actual.PrivateDnsZoneGroupName != v.Expected.PrivateDnsZoneGroupName {
			t.Fatalf("Expected %q but got %q for PrivateDnsZoneGroupName", v.Expected.PrivateDnsZoneGroupName, actual.PrivateDnsZoneGroupName)
		}

	}
}

func TestSegmentsForPrivateDnsZoneGroupId(t *testing.T) {
	segments := PrivateDnsZoneGroupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PrivateDnsZoneGroupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
