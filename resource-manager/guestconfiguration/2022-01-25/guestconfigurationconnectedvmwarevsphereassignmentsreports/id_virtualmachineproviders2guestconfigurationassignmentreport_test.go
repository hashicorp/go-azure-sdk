package guestconfigurationconnectedvmwarevsphereassignmentsreports

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = VirtualMachineProviders2GuestConfigurationAssignmentReportId{}

func TestNewVirtualMachineProviders2GuestConfigurationAssignmentReportID(t *testing.T) {
	id := NewVirtualMachineProviders2GuestConfigurationAssignmentReportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue", "guestConfigurationAssignmentValue", "reportIdValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.VirtualMachineName != "virtualMachineValue" {
		t.Fatalf("Expected %q but got %q for Segment 'VirtualMachineName'", id.VirtualMachineName, "virtualMachineValue")
	}

	if id.GuestConfigurationAssignmentName != "guestConfigurationAssignmentValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GuestConfigurationAssignmentName'", id.GuestConfigurationAssignmentName, "guestConfigurationAssignmentValue")
	}

	if id.ReportId != "reportIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ReportId'", id.ReportId, "reportIdValue")
	}
}

func TestFormatVirtualMachineProviders2GuestConfigurationAssignmentReportID(t *testing.T) {
	actual := NewVirtualMachineProviders2GuestConfigurationAssignmentReportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue", "guestConfigurationAssignmentValue", "reportIdValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/virtualMachineValue/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/guestConfigurationAssignmentValue/reports/reportIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseVirtualMachineProviders2GuestConfigurationAssignmentReportID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VirtualMachineProviders2GuestConfigurationAssignmentReportId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/virtualMachineValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/virtualMachineValue/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/virtualMachineValue/providers/Microsoft.GuestConfiguration",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/virtualMachineValue/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/virtualMachineValue/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/guestConfigurationAssignmentValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/virtualMachineValue/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/guestConfigurationAssignmentValue/reports",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/virtualMachineValue/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/guestConfigurationAssignmentValue/reports/reportIdValue",
			Expected: &VirtualMachineProviders2GuestConfigurationAssignmentReportId{
				SubscriptionId:                   "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                "example-resource-group",
				VirtualMachineName:               "virtualMachineValue",
				GuestConfigurationAssignmentName: "guestConfigurationAssignmentValue",
				ReportId:                         "reportIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/virtualMachineValue/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/guestConfigurationAssignmentValue/reports/reportIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVirtualMachineProviders2GuestConfigurationAssignmentReportID(v.Input)
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

		if actual.VirtualMachineName != v.Expected.VirtualMachineName {
			t.Fatalf("Expected %q but got %q for VirtualMachineName", v.Expected.VirtualMachineName, actual.VirtualMachineName)
		}

		if actual.GuestConfigurationAssignmentName != v.Expected.GuestConfigurationAssignmentName {
			t.Fatalf("Expected %q but got %q for GuestConfigurationAssignmentName", v.Expected.GuestConfigurationAssignmentName, actual.GuestConfigurationAssignmentName)
		}

		if actual.ReportId != v.Expected.ReportId {
			t.Fatalf("Expected %q but got %q for ReportId", v.Expected.ReportId, actual.ReportId)
		}

	}
}

func TestParseVirtualMachineProviders2GuestConfigurationAssignmentReportIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VirtualMachineProviders2GuestConfigurationAssignmentReportId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOnNeCtEdVmWaReVsPhErE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOnNeCtEdVmWaReVsPhErE/vIrTuAlMaChInEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/virtualMachineValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOnNeCtEdVmWaReVsPhErE/vIrTuAlMaChInEs/vIrTuAlMaChInEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/virtualMachineValue/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOnNeCtEdVmWaReVsPhErE/vIrTuAlMaChInEs/vIrTuAlMaChInEvAlUe/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/virtualMachineValue/providers/Microsoft.GuestConfiguration",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOnNeCtEdVmWaReVsPhErE/vIrTuAlMaChInEs/vIrTuAlMaChInEvAlUe/pRoViDeRs/mIcRoSoFt.gUeStCoNfIgUrAtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/virtualMachineValue/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOnNeCtEdVmWaReVsPhErE/vIrTuAlMaChInEs/vIrTuAlMaChInEvAlUe/pRoViDeRs/mIcRoSoFt.gUeStCoNfIgUrAtIoN/gUeStCoNfIgUrAtIoNaSsIgNmEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/virtualMachineValue/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/guestConfigurationAssignmentValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOnNeCtEdVmWaReVsPhErE/vIrTuAlMaChInEs/vIrTuAlMaChInEvAlUe/pRoViDeRs/mIcRoSoFt.gUeStCoNfIgUrAtIoN/gUeStCoNfIgUrAtIoNaSsIgNmEnTs/gUeStCoNfIgUrAtIoNaSsIgNmEnTvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/virtualMachineValue/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/guestConfigurationAssignmentValue/reports",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOnNeCtEdVmWaReVsPhErE/vIrTuAlMaChInEs/vIrTuAlMaChInEvAlUe/pRoViDeRs/mIcRoSoFt.gUeStCoNfIgUrAtIoN/gUeStCoNfIgUrAtIoNaSsIgNmEnTs/gUeStCoNfIgUrAtIoNaSsIgNmEnTvAlUe/rEpOrTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/virtualMachineValue/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/guestConfigurationAssignmentValue/reports/reportIdValue",
			Expected: &VirtualMachineProviders2GuestConfigurationAssignmentReportId{
				SubscriptionId:                   "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                "example-resource-group",
				VirtualMachineName:               "virtualMachineValue",
				GuestConfigurationAssignmentName: "guestConfigurationAssignmentValue",
				ReportId:                         "reportIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/virtualMachineValue/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/guestConfigurationAssignmentValue/reports/reportIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOnNeCtEdVmWaReVsPhErE/vIrTuAlMaChInEs/vIrTuAlMaChInEvAlUe/pRoViDeRs/mIcRoSoFt.gUeStCoNfIgUrAtIoN/gUeStCoNfIgUrAtIoNaSsIgNmEnTs/gUeStCoNfIgUrAtIoNaSsIgNmEnTvAlUe/rEpOrTs/rEpOrTiDvAlUe",
			Expected: &VirtualMachineProviders2GuestConfigurationAssignmentReportId{
				SubscriptionId:                   "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                "eXaMpLe-rEsOuRcE-GrOuP",
				VirtualMachineName:               "vIrTuAlMaChInEvAlUe",
				GuestConfigurationAssignmentName: "gUeStCoNfIgUrAtIoNaSsIgNmEnTvAlUe",
				ReportId:                         "rEpOrTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOnNeCtEdVmWaReVsPhErE/vIrTuAlMaChInEs/vIrTuAlMaChInEvAlUe/pRoViDeRs/mIcRoSoFt.gUeStCoNfIgUrAtIoN/gUeStCoNfIgUrAtIoNaSsIgNmEnTs/gUeStCoNfIgUrAtIoNaSsIgNmEnTvAlUe/rEpOrTs/rEpOrTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVirtualMachineProviders2GuestConfigurationAssignmentReportIDInsensitively(v.Input)
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

		if actual.VirtualMachineName != v.Expected.VirtualMachineName {
			t.Fatalf("Expected %q but got %q for VirtualMachineName", v.Expected.VirtualMachineName, actual.VirtualMachineName)
		}

		if actual.GuestConfigurationAssignmentName != v.Expected.GuestConfigurationAssignmentName {
			t.Fatalf("Expected %q but got %q for GuestConfigurationAssignmentName", v.Expected.GuestConfigurationAssignmentName, actual.GuestConfigurationAssignmentName)
		}

		if actual.ReportId != v.Expected.ReportId {
			t.Fatalf("Expected %q but got %q for ReportId", v.Expected.ReportId, actual.ReportId)
		}

	}
}

func TestSegmentsForVirtualMachineProviders2GuestConfigurationAssignmentReportId(t *testing.T) {
	segments := VirtualMachineProviders2GuestConfigurationAssignmentReportId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("VirtualMachineProviders2GuestConfigurationAssignmentReportId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
