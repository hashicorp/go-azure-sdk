package shares

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ShareId{}

func TestNewShareID(t *testing.T) {
	id := NewShareID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName", "shareName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.DataBoxEdgeDeviceName != "dataBoxEdgeDeviceName" {
		t.Fatalf("Expected %q but got %q for Segment 'DataBoxEdgeDeviceName'", id.DataBoxEdgeDeviceName, "dataBoxEdgeDeviceName")
	}

	if id.ShareName != "shareName" {
		t.Fatalf("Expected %q but got %q for Segment 'ShareName'", id.ShareName, "shareName")
	}
}

func TestFormatShareID(t *testing.T) {
	actual := NewShareID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName", "shareName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceName/shares/shareName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseShareID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ShareId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceName/shares",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceName/shares/shareName",
			Expected: &ShareId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "example-resource-group",
				DataBoxEdgeDeviceName: "dataBoxEdgeDeviceName",
				ShareName:             "shareName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceName/shares/shareName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseShareID(v.Input)
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

		if actual.ShareName != v.Expected.ShareName {
			t.Fatalf("Expected %q but got %q for ShareName", v.Expected.ShareName, actual.ShareName)
		}

	}
}

func TestParseShareIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ShareId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtAbOxEdGe/dAtAbOxEdGeDeViCeS/dAtAbOxEdGeDeViCeNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceName/shares",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtAbOxEdGe/dAtAbOxEdGeDeViCeS/dAtAbOxEdGeDeViCeNaMe/sHaReS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceName/shares/shareName",
			Expected: &ShareId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "example-resource-group",
				DataBoxEdgeDeviceName: "dataBoxEdgeDeviceName",
				ShareName:             "shareName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/dataBoxEdgeDeviceName/shares/shareName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtAbOxEdGe/dAtAbOxEdGeDeViCeS/dAtAbOxEdGeDeViCeNaMe/sHaReS/sHaReNaMe",
			Expected: &ShareId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "eXaMpLe-rEsOuRcE-GrOuP",
				DataBoxEdgeDeviceName: "dAtAbOxEdGeDeViCeNaMe",
				ShareName:             "sHaReNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtAbOxEdGe/dAtAbOxEdGeDeViCeS/dAtAbOxEdGeDeViCeNaMe/sHaReS/sHaReNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseShareIDInsensitively(v.Input)
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

		if actual.ShareName != v.Expected.ShareName {
			t.Fatalf("Expected %q but got %q for ShareName", v.Expected.ShareName, actual.ShareName)
		}

	}
}

func TestSegmentsForShareId(t *testing.T) {
	segments := ShareId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ShareId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
