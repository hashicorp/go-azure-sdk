package availabilitygrouplisteners

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AvailabilityGroupListenerId{}

func TestNewAvailabilityGroupListenerID(t *testing.T) {
	id := NewAvailabilityGroupListenerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlVirtualMachineGroupName", "availabilityGroupListenerName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.SqlVirtualMachineGroupName != "sqlVirtualMachineGroupName" {
		t.Fatalf("Expected %q but got %q for Segment 'SqlVirtualMachineGroupName'", id.SqlVirtualMachineGroupName, "sqlVirtualMachineGroupName")
	}

	if id.AvailabilityGroupListenerName != "availabilityGroupListenerName" {
		t.Fatalf("Expected %q but got %q for Segment 'AvailabilityGroupListenerName'", id.AvailabilityGroupListenerName, "availabilityGroupListenerName")
	}
}

func TestFormatAvailabilityGroupListenerID(t *testing.T) {
	actual := NewAvailabilityGroupListenerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlVirtualMachineGroupName", "availabilityGroupListenerName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/sqlVirtualMachineGroupName/availabilityGroupListeners/availabilityGroupListenerName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAvailabilityGroupListenerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AvailabilityGroupListenerId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/sqlVirtualMachineGroupName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/sqlVirtualMachineGroupName/availabilityGroupListeners",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/sqlVirtualMachineGroupName/availabilityGroupListeners/availabilityGroupListenerName",
			Expected: &AvailabilityGroupListenerId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "example-resource-group",
				SqlVirtualMachineGroupName:    "sqlVirtualMachineGroupName",
				AvailabilityGroupListenerName: "availabilityGroupListenerName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/sqlVirtualMachineGroupName/availabilityGroupListeners/availabilityGroupListenerName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAvailabilityGroupListenerID(v.Input)
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

		if actual.SqlVirtualMachineGroupName != v.Expected.SqlVirtualMachineGroupName {
			t.Fatalf("Expected %q but got %q for SqlVirtualMachineGroupName", v.Expected.SqlVirtualMachineGroupName, actual.SqlVirtualMachineGroupName)
		}

		if actual.AvailabilityGroupListenerName != v.Expected.AvailabilityGroupListenerName {
			t.Fatalf("Expected %q but got %q for AvailabilityGroupListenerName", v.Expected.AvailabilityGroupListenerName, actual.AvailabilityGroupListenerName)
		}

	}
}

func TestParseAvailabilityGroupListenerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AvailabilityGroupListenerId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQlViRtUaLmAcHiNe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQlViRtUaLmAcHiNe/sQlViRtUaLmAcHiNeGrOuPs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/sqlVirtualMachineGroupName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQlViRtUaLmAcHiNe/sQlViRtUaLmAcHiNeGrOuPs/sQlViRtUaLmAcHiNeGrOuPnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/sqlVirtualMachineGroupName/availabilityGroupListeners",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQlViRtUaLmAcHiNe/sQlViRtUaLmAcHiNeGrOuPs/sQlViRtUaLmAcHiNeGrOuPnAmE/aVaIlAbIlItYgRoUpLiStEnErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/sqlVirtualMachineGroupName/availabilityGroupListeners/availabilityGroupListenerName",
			Expected: &AvailabilityGroupListenerId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "example-resource-group",
				SqlVirtualMachineGroupName:    "sqlVirtualMachineGroupName",
				AvailabilityGroupListenerName: "availabilityGroupListenerName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/sqlVirtualMachineGroupName/availabilityGroupListeners/availabilityGroupListenerName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQlViRtUaLmAcHiNe/sQlViRtUaLmAcHiNeGrOuPs/sQlViRtUaLmAcHiNeGrOuPnAmE/aVaIlAbIlItYgRoUpLiStEnErS/aVaIlAbIlItYgRoUpLiStEnErNaMe",
			Expected: &AvailabilityGroupListenerId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "eXaMpLe-rEsOuRcE-GrOuP",
				SqlVirtualMachineGroupName:    "sQlViRtUaLmAcHiNeGrOuPnAmE",
				AvailabilityGroupListenerName: "aVaIlAbIlItYgRoUpLiStEnErNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQlViRtUaLmAcHiNe/sQlViRtUaLmAcHiNeGrOuPs/sQlViRtUaLmAcHiNeGrOuPnAmE/aVaIlAbIlItYgRoUpLiStEnErS/aVaIlAbIlItYgRoUpLiStEnErNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAvailabilityGroupListenerIDInsensitively(v.Input)
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

		if actual.SqlVirtualMachineGroupName != v.Expected.SqlVirtualMachineGroupName {
			t.Fatalf("Expected %q but got %q for SqlVirtualMachineGroupName", v.Expected.SqlVirtualMachineGroupName, actual.SqlVirtualMachineGroupName)
		}

		if actual.AvailabilityGroupListenerName != v.Expected.AvailabilityGroupListenerName {
			t.Fatalf("Expected %q but got %q for AvailabilityGroupListenerName", v.Expected.AvailabilityGroupListenerName, actual.AvailabilityGroupListenerName)
		}

	}
}

func TestSegmentsForAvailabilityGroupListenerId(t *testing.T) {
	segments := AvailabilityGroupListenerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AvailabilityGroupListenerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
