package resourceguards

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeleteResourceGuardProxyRequestId{}

func TestNewDeleteResourceGuardProxyRequestID(t *testing.T) {
	id := NewDeleteResourceGuardProxyRequestID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName", "deleteResourceGuardProxyRequestName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.ResourceGuardName != "resourceGuardName" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGuardName'", id.ResourceGuardName, "resourceGuardName")
	}

	if id.DeleteResourceGuardProxyRequestName != "deleteResourceGuardProxyRequestName" {
		t.Fatalf("Expected %q but got %q for Segment 'DeleteResourceGuardProxyRequestName'", id.DeleteResourceGuardProxyRequestName, "deleteResourceGuardProxyRequestName")
	}
}

func TestFormatDeleteResourceGuardProxyRequestID(t *testing.T) {
	actual := NewDeleteResourceGuardProxyRequestID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName", "deleteResourceGuardProxyRequestName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/resourceGuards/resourceGuardName/deleteResourceGuardProxyRequests/deleteResourceGuardProxyRequestName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeleteResourceGuardProxyRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeleteResourceGuardProxyRequestId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/resourceGuards",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/resourceGuards/resourceGuardName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/resourceGuards/resourceGuardName/deleteResourceGuardProxyRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/resourceGuards/resourceGuardName/deleteResourceGuardProxyRequests/deleteResourceGuardProxyRequestName",
			Expected: &DeleteResourceGuardProxyRequestId{
				SubscriptionId:                      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                   "example-resource-group",
				ResourceGuardName:                   "resourceGuardName",
				DeleteResourceGuardProxyRequestName: "deleteResourceGuardProxyRequestName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/resourceGuards/resourceGuardName/deleteResourceGuardProxyRequests/deleteResourceGuardProxyRequestName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeleteResourceGuardProxyRequestID(v.Input)
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

		if actual.ResourceGuardName != v.Expected.ResourceGuardName {
			t.Fatalf("Expected %q but got %q for ResourceGuardName", v.Expected.ResourceGuardName, actual.ResourceGuardName)
		}

		if actual.DeleteResourceGuardProxyRequestName != v.Expected.DeleteResourceGuardProxyRequestName {
			t.Fatalf("Expected %q but got %q for DeleteResourceGuardProxyRequestName", v.Expected.DeleteResourceGuardProxyRequestName, actual.DeleteResourceGuardProxyRequestName)
		}

	}
}

func TestParseDeleteResourceGuardProxyRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeleteResourceGuardProxyRequestId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/resourceGuards",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtApRoTeCtIoN/rEsOuRcEgUaRdS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/resourceGuards/resourceGuardName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtApRoTeCtIoN/rEsOuRcEgUaRdS/rEsOuRcEgUaRdNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/resourceGuards/resourceGuardName/deleteResourceGuardProxyRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtApRoTeCtIoN/rEsOuRcEgUaRdS/rEsOuRcEgUaRdNaMe/dElEtErEsOuRcEgUaRdPrOxYrEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/resourceGuards/resourceGuardName/deleteResourceGuardProxyRequests/deleteResourceGuardProxyRequestName",
			Expected: &DeleteResourceGuardProxyRequestId{
				SubscriptionId:                      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                   "example-resource-group",
				ResourceGuardName:                   "resourceGuardName",
				DeleteResourceGuardProxyRequestName: "deleteResourceGuardProxyRequestName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DataProtection/resourceGuards/resourceGuardName/deleteResourceGuardProxyRequests/deleteResourceGuardProxyRequestName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtApRoTeCtIoN/rEsOuRcEgUaRdS/rEsOuRcEgUaRdNaMe/dElEtErEsOuRcEgUaRdPrOxYrEqUeStS/dElEtErEsOuRcEgUaRdPrOxYrEqUeStNaMe",
			Expected: &DeleteResourceGuardProxyRequestId{
				SubscriptionId:                      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                   "eXaMpLe-rEsOuRcE-GrOuP",
				ResourceGuardName:                   "rEsOuRcEgUaRdNaMe",
				DeleteResourceGuardProxyRequestName: "dElEtErEsOuRcEgUaRdPrOxYrEqUeStNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dAtApRoTeCtIoN/rEsOuRcEgUaRdS/rEsOuRcEgUaRdNaMe/dElEtErEsOuRcEgUaRdPrOxYrEqUeStS/dElEtErEsOuRcEgUaRdPrOxYrEqUeStNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeleteResourceGuardProxyRequestIDInsensitively(v.Input)
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

		if actual.ResourceGuardName != v.Expected.ResourceGuardName {
			t.Fatalf("Expected %q but got %q for ResourceGuardName", v.Expected.ResourceGuardName, actual.ResourceGuardName)
		}

		if actual.DeleteResourceGuardProxyRequestName != v.Expected.DeleteResourceGuardProxyRequestName {
			t.Fatalf("Expected %q but got %q for DeleteResourceGuardProxyRequestName", v.Expected.DeleteResourceGuardProxyRequestName, actual.DeleteResourceGuardProxyRequestName)
		}

	}
}

func TestSegmentsForDeleteResourceGuardProxyRequestId(t *testing.T) {
	segments := DeleteResourceGuardProxyRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeleteResourceGuardProxyRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
