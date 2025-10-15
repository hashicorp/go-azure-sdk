package dataflowgraph

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DataflowGraphId{}

func TestNewDataflowGraphID(t *testing.T) {
	id := NewDataflowGraphID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instanceName", "dataflowProfileName", "dataflowGraphName")

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

	if id.DataflowGraphName != "dataflowGraphName" {
		t.Fatalf("Expected %q but got %q for Segment 'DataflowGraphName'", id.DataflowGraphName, "dataflowGraphName")
	}
}

func TestFormatDataflowGraphID(t *testing.T) {
	actual := NewDataflowGraphID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instanceName", "dataflowProfileName", "dataflowGraphName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances/instanceName/dataflowProfiles/dataflowProfileName/dataflowGraphs/dataflowGraphName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDataflowGraphID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DataflowGraphId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances/instanceName/dataflowProfiles/dataflowProfileName/dataflowGraphs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances/instanceName/dataflowProfiles/dataflowProfileName/dataflowGraphs/dataflowGraphName",
			Expected: &DataflowGraphId{
				SubscriptionId:      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:   "example-resource-group",
				InstanceName:        "instanceName",
				DataflowProfileName: "dataflowProfileName",
				DataflowGraphName:   "dataflowGraphName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances/instanceName/dataflowProfiles/dataflowProfileName/dataflowGraphs/dataflowGraphName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDataflowGraphID(v.Input)
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

		if actual.DataflowGraphName != v.Expected.DataflowGraphName {
			t.Fatalf("Expected %q but got %q for DataflowGraphName", v.Expected.DataflowGraphName, actual.DataflowGraphName)
		}

	}
}

func TestParseDataflowGraphIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DataflowGraphId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances/instanceName/dataflowProfiles/dataflowProfileName/dataflowGraphs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.iOtOpErAtIoNs/iNsTaNcEs/iNsTaNcEnAmE/dAtAfLoWpRoFiLeS/dAtAfLoWpRoFiLeNaMe/dAtAfLoWgRaPhS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances/instanceName/dataflowProfiles/dataflowProfileName/dataflowGraphs/dataflowGraphName",
			Expected: &DataflowGraphId{
				SubscriptionId:      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:   "example-resource-group",
				InstanceName:        "instanceName",
				DataflowProfileName: "dataflowProfileName",
				DataflowGraphName:   "dataflowGraphName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.IoTOperations/instances/instanceName/dataflowProfiles/dataflowProfileName/dataflowGraphs/dataflowGraphName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.iOtOpErAtIoNs/iNsTaNcEs/iNsTaNcEnAmE/dAtAfLoWpRoFiLeS/dAtAfLoWpRoFiLeNaMe/dAtAfLoWgRaPhS/dAtAfLoWgRaPhNaMe",
			Expected: &DataflowGraphId{
				SubscriptionId:      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:   "eXaMpLe-rEsOuRcE-GrOuP",
				InstanceName:        "iNsTaNcEnAmE",
				DataflowProfileName: "dAtAfLoWpRoFiLeNaMe",
				DataflowGraphName:   "dAtAfLoWgRaPhNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.iOtOpErAtIoNs/iNsTaNcEs/iNsTaNcEnAmE/dAtAfLoWpRoFiLeS/dAtAfLoWpRoFiLeNaMe/dAtAfLoWgRaPhS/dAtAfLoWgRaPhNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDataflowGraphIDInsensitively(v.Input)
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

		if actual.DataflowGraphName != v.Expected.DataflowGraphName {
			t.Fatalf("Expected %q but got %q for DataflowGraphName", v.Expected.DataflowGraphName, actual.DataflowGraphName)
		}

	}
}

func TestSegmentsForDataflowGraphId(t *testing.T) {
	segments := DataflowGraphId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DataflowGraphId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
