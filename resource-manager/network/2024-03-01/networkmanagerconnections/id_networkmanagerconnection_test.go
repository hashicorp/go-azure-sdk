package networkmanagerconnections

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &NetworkManagerConnectionId{}

func TestNewNetworkManagerConnectionID(t *testing.T) {
	id := NewNetworkManagerConnectionID("12345678-1234-9876-4563-123456789012", "networkManagerConnectionName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.NetworkManagerConnectionName != "networkManagerConnectionName" {
		t.Fatalf("Expected %q but got %q for Segment 'NetworkManagerConnectionName'", id.NetworkManagerConnectionName, "networkManagerConnectionName")
	}
}

func TestFormatNetworkManagerConnectionID(t *testing.T) {
	actual := NewNetworkManagerConnectionID("12345678-1234-9876-4563-123456789012", "networkManagerConnectionName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/networkManagerConnections/networkManagerConnectionName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseNetworkManagerConnectionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *NetworkManagerConnectionId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/networkManagerConnections",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/networkManagerConnections/networkManagerConnectionName",
			Expected: &NetworkManagerConnectionId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				NetworkManagerConnectionName: "networkManagerConnectionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/networkManagerConnections/networkManagerConnectionName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseNetworkManagerConnectionID(v.Input)
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

		if actual.NetworkManagerConnectionName != v.Expected.NetworkManagerConnectionName {
			t.Fatalf("Expected %q but got %q for NetworkManagerConnectionName", v.Expected.NetworkManagerConnectionName, actual.NetworkManagerConnectionName)
		}

	}
}

func TestParseNetworkManagerConnectionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *NetworkManagerConnectionId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.nEtWoRk",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/networkManagerConnections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRcOnNeCtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/networkManagerConnections/networkManagerConnectionName",
			Expected: &NetworkManagerConnectionId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				NetworkManagerConnectionName: "networkManagerConnectionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Network/networkManagerConnections/networkManagerConnectionName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRcOnNeCtIoNs/nEtWoRkMaNaGeRcOnNeCtIoNnAmE",
			Expected: &NetworkManagerConnectionId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				NetworkManagerConnectionName: "nEtWoRkMaNaGeRcOnNeCtIoNnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRcOnNeCtIoNs/nEtWoRkMaNaGeRcOnNeCtIoNnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseNetworkManagerConnectionIDInsensitively(v.Input)
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

		if actual.NetworkManagerConnectionName != v.Expected.NetworkManagerConnectionName {
			t.Fatalf("Expected %q but got %q for NetworkManagerConnectionName", v.Expected.NetworkManagerConnectionName, actual.NetworkManagerConnectionName)
		}

	}
}

func TestSegmentsForNetworkManagerConnectionId(t *testing.T) {
	segments := NetworkManagerConnectionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("NetworkManagerConnectionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
