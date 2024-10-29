package standbyvirtualmachines

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &StandbyVirtualMachineId{}

func TestNewStandbyVirtualMachineID(t *testing.T) {
	id := NewStandbyVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "standbyVirtualMachinePoolName", "standbyVirtualMachineName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.StandbyVirtualMachinePoolName != "standbyVirtualMachinePoolName" {
		t.Fatalf("Expected %q but got %q for Segment 'StandbyVirtualMachinePoolName'", id.StandbyVirtualMachinePoolName, "standbyVirtualMachinePoolName")
	}

	if id.StandbyVirtualMachineName != "standbyVirtualMachineName" {
		t.Fatalf("Expected %q but got %q for Segment 'StandbyVirtualMachineName'", id.StandbyVirtualMachineName, "standbyVirtualMachineName")
	}
}

func TestFormatStandbyVirtualMachineID(t *testing.T) {
	actual := NewStandbyVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "standbyVirtualMachinePoolName", "standbyVirtualMachineName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/standbyVirtualMachinePoolName/standbyVirtualMachines/standbyVirtualMachineName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseStandbyVirtualMachineID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *StandbyVirtualMachineId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyVirtualMachinePools",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/standbyVirtualMachinePoolName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/standbyVirtualMachinePoolName/standbyVirtualMachines",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/standbyVirtualMachinePoolName/standbyVirtualMachines/standbyVirtualMachineName",
			Expected: &StandbyVirtualMachineId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "example-resource-group",
				StandbyVirtualMachinePoolName: "standbyVirtualMachinePoolName",
				StandbyVirtualMachineName:     "standbyVirtualMachineName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/standbyVirtualMachinePoolName/standbyVirtualMachines/standbyVirtualMachineName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseStandbyVirtualMachineID(v.Input)
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

		if actual.StandbyVirtualMachinePoolName != v.Expected.StandbyVirtualMachinePoolName {
			t.Fatalf("Expected %q but got %q for StandbyVirtualMachinePoolName", v.Expected.StandbyVirtualMachinePoolName, actual.StandbyVirtualMachinePoolName)
		}

		if actual.StandbyVirtualMachineName != v.Expected.StandbyVirtualMachineName {
			t.Fatalf("Expected %q but got %q for StandbyVirtualMachineName", v.Expected.StandbyVirtualMachineName, actual.StandbyVirtualMachineName)
		}

	}
}

func TestParseStandbyVirtualMachineIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *StandbyVirtualMachineId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sTaNdByPoOl",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyVirtualMachinePools",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sTaNdByPoOl/sTaNdByViRtUaLmAcHiNePoOlS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/standbyVirtualMachinePoolName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sTaNdByPoOl/sTaNdByViRtUaLmAcHiNePoOlS/sTaNdByViRtUaLmAcHiNePoOlNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/standbyVirtualMachinePoolName/standbyVirtualMachines",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sTaNdByPoOl/sTaNdByViRtUaLmAcHiNePoOlS/sTaNdByViRtUaLmAcHiNePoOlNaMe/sTaNdByViRtUaLmAcHiNeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/standbyVirtualMachinePoolName/standbyVirtualMachines/standbyVirtualMachineName",
			Expected: &StandbyVirtualMachineId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "example-resource-group",
				StandbyVirtualMachinePoolName: "standbyVirtualMachinePoolName",
				StandbyVirtualMachineName:     "standbyVirtualMachineName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/standbyVirtualMachinePoolName/standbyVirtualMachines/standbyVirtualMachineName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sTaNdByPoOl/sTaNdByViRtUaLmAcHiNePoOlS/sTaNdByViRtUaLmAcHiNePoOlNaMe/sTaNdByViRtUaLmAcHiNeS/sTaNdByViRtUaLmAcHiNeNaMe",
			Expected: &StandbyVirtualMachineId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "eXaMpLe-rEsOuRcE-GrOuP",
				StandbyVirtualMachinePoolName: "sTaNdByViRtUaLmAcHiNePoOlNaMe",
				StandbyVirtualMachineName:     "sTaNdByViRtUaLmAcHiNeNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sTaNdByPoOl/sTaNdByViRtUaLmAcHiNePoOlS/sTaNdByViRtUaLmAcHiNePoOlNaMe/sTaNdByViRtUaLmAcHiNeS/sTaNdByViRtUaLmAcHiNeNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseStandbyVirtualMachineIDInsensitively(v.Input)
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

		if actual.StandbyVirtualMachinePoolName != v.Expected.StandbyVirtualMachinePoolName {
			t.Fatalf("Expected %q but got %q for StandbyVirtualMachinePoolName", v.Expected.StandbyVirtualMachinePoolName, actual.StandbyVirtualMachinePoolName)
		}

		if actual.StandbyVirtualMachineName != v.Expected.StandbyVirtualMachineName {
			t.Fatalf("Expected %q but got %q for StandbyVirtualMachineName", v.Expected.StandbyVirtualMachineName, actual.StandbyVirtualMachineName)
		}

	}
}

func TestSegmentsForStandbyVirtualMachineId(t *testing.T) {
	segments := StandbyVirtualMachineId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("StandbyVirtualMachineId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
