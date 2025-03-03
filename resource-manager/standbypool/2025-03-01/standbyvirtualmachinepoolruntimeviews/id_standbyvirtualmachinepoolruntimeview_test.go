package standbyvirtualmachinepoolruntimeviews

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &StandbyVirtualMachinePoolRuntimeViewId{}

func TestNewStandbyVirtualMachinePoolRuntimeViewID(t *testing.T) {
	id := NewStandbyVirtualMachinePoolRuntimeViewID("12345678-1234-9876-4563-123456789012", "example-resource-group", "standbyVirtualMachinePoolName", "runtimeViewName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.StandbyVirtualMachinePoolName != "standbyVirtualMachinePoolName" {
		t.Fatalf("Expected %q but got %q for Segment 'StandbyVirtualMachinePoolName'", id.StandbyVirtualMachinePoolName, "standbyVirtualMachinePoolName")
	}

	if id.RuntimeViewName != "runtimeViewName" {
		t.Fatalf("Expected %q but got %q for Segment 'RuntimeViewName'", id.RuntimeViewName, "runtimeViewName")
	}
}

func TestFormatStandbyVirtualMachinePoolRuntimeViewID(t *testing.T) {
	actual := NewStandbyVirtualMachinePoolRuntimeViewID("12345678-1234-9876-4563-123456789012", "example-resource-group", "standbyVirtualMachinePoolName", "runtimeViewName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/standbyVirtualMachinePoolName/runtimeViews/runtimeViewName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseStandbyVirtualMachinePoolRuntimeViewID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *StandbyVirtualMachinePoolRuntimeViewId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/standbyVirtualMachinePoolName/runtimeViews",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/standbyVirtualMachinePoolName/runtimeViews/runtimeViewName",
			Expected: &StandbyVirtualMachinePoolRuntimeViewId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "example-resource-group",
				StandbyVirtualMachinePoolName: "standbyVirtualMachinePoolName",
				RuntimeViewName:               "runtimeViewName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/standbyVirtualMachinePoolName/runtimeViews/runtimeViewName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseStandbyVirtualMachinePoolRuntimeViewID(v.Input)
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

		if actual.RuntimeViewName != v.Expected.RuntimeViewName {
			t.Fatalf("Expected %q but got %q for RuntimeViewName", v.Expected.RuntimeViewName, actual.RuntimeViewName)
		}

	}
}

func TestParseStandbyVirtualMachinePoolRuntimeViewIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *StandbyVirtualMachinePoolRuntimeViewId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/standbyVirtualMachinePoolName/runtimeViews",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sTaNdByPoOl/sTaNdByViRtUaLmAcHiNePoOlS/sTaNdByViRtUaLmAcHiNePoOlNaMe/rUnTiMeViEwS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/standbyVirtualMachinePoolName/runtimeViews/runtimeViewName",
			Expected: &StandbyVirtualMachinePoolRuntimeViewId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "example-resource-group",
				StandbyVirtualMachinePoolName: "standbyVirtualMachinePoolName",
				RuntimeViewName:               "runtimeViewName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/standbyVirtualMachinePoolName/runtimeViews/runtimeViewName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sTaNdByPoOl/sTaNdByViRtUaLmAcHiNePoOlS/sTaNdByViRtUaLmAcHiNePoOlNaMe/rUnTiMeViEwS/rUnTiMeViEwNaMe",
			Expected: &StandbyVirtualMachinePoolRuntimeViewId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "eXaMpLe-rEsOuRcE-GrOuP",
				StandbyVirtualMachinePoolName: "sTaNdByViRtUaLmAcHiNePoOlNaMe",
				RuntimeViewName:               "rUnTiMeViEwNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sTaNdByPoOl/sTaNdByViRtUaLmAcHiNePoOlS/sTaNdByViRtUaLmAcHiNePoOlNaMe/rUnTiMeViEwS/rUnTiMeViEwNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseStandbyVirtualMachinePoolRuntimeViewIDInsensitively(v.Input)
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

		if actual.RuntimeViewName != v.Expected.RuntimeViewName {
			t.Fatalf("Expected %q but got %q for RuntimeViewName", v.Expected.RuntimeViewName, actual.RuntimeViewName)
		}

	}
}

func TestSegmentsForStandbyVirtualMachinePoolRuntimeViewId(t *testing.T) {
	segments := StandbyVirtualMachinePoolRuntimeViewId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("StandbyVirtualMachinePoolRuntimeViewId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
