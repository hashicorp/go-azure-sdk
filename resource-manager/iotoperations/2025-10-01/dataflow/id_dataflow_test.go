package dataflow

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DataflowId{}

func TestNewDataflowID(t *testing.T) {
	id := NewDataflowID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instanceName", "dataflowProfileName", "dataflowName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.InstanceName != "instanceName" {
		t.Fatalf("Expected %q but got %q for Segment 'InstanceName'", id.InstanceName, "instanceName")
	}

	if id.DataflowProfileName != "dataflowProfileName" {
		t.Fatalf("Expected %q but got %q for Segment 'DataflowProfileName'", id.DataflowProfileName, "dataflowProfileName")
	}

	if id.DataflowName != "dataflowName" {
		t.Fatalf("Expected %q but got %q for Segment 'DataflowName'", id.DataflowName, "dataflowName")
	}
}

func TestFormatDataflowID(t *testing.T) {
	actual := NewDataflowID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instanceName", "dataflowProfileName", "dataflowName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances/instanceName/dataflowProfiles/dataflowProfileName/dataflows/dataflowName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDataflowID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DataflowId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances/instanceName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances/instanceName/dataflowProfiles",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances/instanceName/dataflowProfiles/dataflowProfileName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances/instanceName/dataflowProfiles/dataflowProfileName/dataflows",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances/instanceName/dataflowProfiles/dataflowProfileName/dataflows/dataflowName",
			Expected: &DataflowId{
				SubscriptionId:      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:   "example-resource-group",
				InstanceName:        "instanceName",
				DataflowProfileName: "dataflowProfileName",
				DataflowName:        "dataflowName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances/instanceName/dataflowProfiles/dataflowProfileName/dataflows/dataflowName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDataflowID(v.Input)
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

		if actual.InstanceName != v.Expected.InstanceName {
			t.Fatalf("Expected %q but got %q for InstanceName", v.Expected.InstanceName, actual.InstanceName)
		}

		if actual.DataflowProfileName != v.Expected.DataflowProfileName {
			t.Fatalf("Expected %q but got %q for DataflowProfileName", v.Expected.DataflowProfileName, actual.DataflowProfileName)
		}

		if actual.DataflowName != v.Expected.DataflowName {
			t.Fatalf("Expected %q but got %q for DataflowName", v.Expected.DataflowName, actual.DataflowName)
		}

	}
}

func TestParseDataflowIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DataflowId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.iOtOpErAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.iOtOpErAtIoNs/iNsTaNcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances/instanceName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.iOtOpErAtIoNs/iNsTaNcEs/iNsTaNcEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances/instanceName/dataflowProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.iOtOpErAtIoNs/iNsTaNcEs/iNsTaNcEnAmE/dAtAfLoWpRoFiLeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances/instanceName/dataflowProfiles/dataflowProfileName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.iOtOpErAtIoNs/iNsTaNcEs/iNsTaNcEnAmE/dAtAfLoWpRoFiLeS/dAtAfLoWpRoFiLeNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances/instanceName/dataflowProfiles/dataflowProfileName/dataflows",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.iOtOpErAtIoNs/iNsTaNcEs/iNsTaNcEnAmE/dAtAfLoWpRoFiLeS/dAtAfLoWpRoFiLeNaMe/dAtAfLoWs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances/instanceName/dataflowProfiles/dataflowProfileName/dataflows/dataflowName",
			Expected: &DataflowId{
				SubscriptionId:      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:   "example-resource-group",
				InstanceName:        "instanceName",
				DataflowProfileName: "dataflowProfileName",
				DataflowName:        "dataflowName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances/instanceName/dataflowProfiles/dataflowProfileName/dataflows/dataflowName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.iOtOpErAtIoNs/iNsTaNcEs/iNsTaNcEnAmE/dAtAfLoWpRoFiLeS/dAtAfLoWpRoFiLeNaMe/dAtAfLoWs/dAtAfLoWnAmE",
			Expected: &DataflowId{
				SubscriptionId:      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:   "eXaMpLe-rEsOuRcE-GrOuP",
				InstanceName:        "iNsTaNcEnAmE",
				DataflowProfileName: "dAtAfLoWpRoFiLeNaMe",
				DataflowName:        "dAtAfLoWnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.iOtOpErAtIoNs/iNsTaNcEs/iNsTaNcEnAmE/dAtAfLoWpRoFiLeS/dAtAfLoWpRoFiLeNaMe/dAtAfLoWs/dAtAfLoWnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDataflowIDInsensitively(v.Input)
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

		if actual.InstanceName != v.Expected.InstanceName {
			t.Fatalf("Expected %q but got %q for InstanceName", v.Expected.InstanceName, actual.InstanceName)
		}

		if actual.DataflowProfileName != v.Expected.DataflowProfileName {
			t.Fatalf("Expected %q but got %q for DataflowProfileName", v.Expected.DataflowProfileName, actual.DataflowProfileName)
		}

		if actual.DataflowName != v.Expected.DataflowName {
			t.Fatalf("Expected %q but got %q for DataflowName", v.Expected.DataflowName, actual.DataflowName)
		}

	}
}

func TestSegmentsForDataflowId(t *testing.T) {
	segments := DataflowId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DataflowId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
