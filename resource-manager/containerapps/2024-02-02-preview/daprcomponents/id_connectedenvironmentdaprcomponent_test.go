package daprcomponents

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ConnectedEnvironmentDaprComponentId{}

func TestNewConnectedEnvironmentDaprComponentID(t *testing.T) {
	id := NewConnectedEnvironmentDaprComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentValue", "daprComponentValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.ConnectedEnvironmentName != "connectedEnvironmentValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ConnectedEnvironmentName'", id.ConnectedEnvironmentName, "connectedEnvironmentValue")
	}

	if id.DaprComponentName != "daprComponentValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DaprComponentName'", id.DaprComponentName, "daprComponentValue")
	}
}

func TestFormatConnectedEnvironmentDaprComponentID(t *testing.T) {
	actual := NewConnectedEnvironmentDaprComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentValue", "daprComponentValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.App/connectedEnvironments/connectedEnvironmentValue/daprComponents/daprComponentValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseConnectedEnvironmentDaprComponentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ConnectedEnvironmentDaprComponentId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.App",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.App/connectedEnvironments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.App/connectedEnvironments/connectedEnvironmentValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.App/connectedEnvironments/connectedEnvironmentValue/daprComponents",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.App/connectedEnvironments/connectedEnvironmentValue/daprComponents/daprComponentValue",
			Expected: &ConnectedEnvironmentDaprComponentId{
				SubscriptionId:           "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:        "example-resource-group",
				ConnectedEnvironmentName: "connectedEnvironmentValue",
				DaprComponentName:        "daprComponentValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.App/connectedEnvironments/connectedEnvironmentValue/daprComponents/daprComponentValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseConnectedEnvironmentDaprComponentID(v.Input)
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

		if actual.ConnectedEnvironmentName != v.Expected.ConnectedEnvironmentName {
			t.Fatalf("Expected %q but got %q for ConnectedEnvironmentName", v.Expected.ConnectedEnvironmentName, actual.ConnectedEnvironmentName)
		}

		if actual.DaprComponentName != v.Expected.DaprComponentName {
			t.Fatalf("Expected %q but got %q for DaprComponentName", v.Expected.DaprComponentName, actual.DaprComponentName)
		}

	}
}

func TestParseConnectedEnvironmentDaprComponentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ConnectedEnvironmentDaprComponentId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.App",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aPp",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.App/connectedEnvironments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aPp/cOnNeCtEdEnViRoNmEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.App/connectedEnvironments/connectedEnvironmentValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aPp/cOnNeCtEdEnViRoNmEnTs/cOnNeCtEdEnViRoNmEnTvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.App/connectedEnvironments/connectedEnvironmentValue/daprComponents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aPp/cOnNeCtEdEnViRoNmEnTs/cOnNeCtEdEnViRoNmEnTvAlUe/dApRcOmPoNeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.App/connectedEnvironments/connectedEnvironmentValue/daprComponents/daprComponentValue",
			Expected: &ConnectedEnvironmentDaprComponentId{
				SubscriptionId:           "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:        "example-resource-group",
				ConnectedEnvironmentName: "connectedEnvironmentValue",
				DaprComponentName:        "daprComponentValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.App/connectedEnvironments/connectedEnvironmentValue/daprComponents/daprComponentValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aPp/cOnNeCtEdEnViRoNmEnTs/cOnNeCtEdEnViRoNmEnTvAlUe/dApRcOmPoNeNtS/dApRcOmPoNeNtVaLuE",
			Expected: &ConnectedEnvironmentDaprComponentId{
				SubscriptionId:           "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:        "eXaMpLe-rEsOuRcE-GrOuP",
				ConnectedEnvironmentName: "cOnNeCtEdEnViRoNmEnTvAlUe",
				DaprComponentName:        "dApRcOmPoNeNtVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aPp/cOnNeCtEdEnViRoNmEnTs/cOnNeCtEdEnViRoNmEnTvAlUe/dApRcOmPoNeNtS/dApRcOmPoNeNtVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseConnectedEnvironmentDaprComponentIDInsensitively(v.Input)
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

		if actual.ConnectedEnvironmentName != v.Expected.ConnectedEnvironmentName {
			t.Fatalf("Expected %q but got %q for ConnectedEnvironmentName", v.Expected.ConnectedEnvironmentName, actual.ConnectedEnvironmentName)
		}

		if actual.DaprComponentName != v.Expected.DaprComponentName {
			t.Fatalf("Expected %q but got %q for DaprComponentName", v.Expected.DaprComponentName, actual.DaprComponentName)
		}

	}
}

func TestSegmentsForConnectedEnvironmentDaprComponentId(t *testing.T) {
	segments := ConnectedEnvironmentDaprComponentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ConnectedEnvironmentDaprComponentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
