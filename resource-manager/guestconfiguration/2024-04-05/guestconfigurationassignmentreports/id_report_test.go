package guestconfigurationassignmentreports

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportId{}

func TestNewReportID(t *testing.T) {
	id := NewReportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetName", "guestConfigurationAssignmentName", "reportName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.VirtualMachineScaleSetName != "virtualMachineScaleSetName" {
		t.Fatalf("Expected %q but got %q for Segment 'VirtualMachineScaleSetName'", id.VirtualMachineScaleSetName, "virtualMachineScaleSetName")
	}

	if id.GuestConfigurationAssignmentName != "guestConfigurationAssignmentName" {
		t.Fatalf("Expected %q but got %q for Segment 'GuestConfigurationAssignmentName'", id.GuestConfigurationAssignmentName, "guestConfigurationAssignmentName")
	}

	if id.ReportName != "reportName" {
		t.Fatalf("Expected %q but got %q for Segment 'ReportName'", id.ReportName, "reportName")
	}
}

func TestFormatReportID(t *testing.T) {
	actual := NewReportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetName", "guestConfigurationAssignmentName", "reportName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetName/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/guestConfigurationAssignmentName/reports/reportName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetName/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetName/providers/Microsoft.GuestConfiguration",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetName/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetName/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/guestConfigurationAssignmentName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetName/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/guestConfigurationAssignmentName/reports",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetName/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/guestConfigurationAssignmentName/reports/reportName",
			Expected: &ReportId{
				SubscriptionId:                   "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                "example-resource-group",
				VirtualMachineScaleSetName:       "virtualMachineScaleSetName",
				GuestConfigurationAssignmentName: "guestConfigurationAssignmentName",
				ReportName:                       "reportName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetName/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/guestConfigurationAssignmentName/reports/reportName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportID(v.Input)
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

		if actual.VirtualMachineScaleSetName != v.Expected.VirtualMachineScaleSetName {
			t.Fatalf("Expected %q but got %q for VirtualMachineScaleSetName", v.Expected.VirtualMachineScaleSetName, actual.VirtualMachineScaleSetName)
		}

		if actual.GuestConfigurationAssignmentName != v.Expected.GuestConfigurationAssignmentName {
			t.Fatalf("Expected %q but got %q for GuestConfigurationAssignmentName", v.Expected.GuestConfigurationAssignmentName, actual.GuestConfigurationAssignmentName)
		}

		if actual.ReportName != v.Expected.ReportName {
			t.Fatalf("Expected %q but got %q for ReportName", v.Expected.ReportName, actual.ReportName)
		}

	}
}

func TestParseReportIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetName/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTnAmE/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetName/providers/Microsoft.GuestConfiguration",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTnAmE/pRoViDeRs/mIcRoSoFt.gUeStCoNfIgUrAtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetName/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTnAmE/pRoViDeRs/mIcRoSoFt.gUeStCoNfIgUrAtIoN/gUeStCoNfIgUrAtIoNaSsIgNmEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetName/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/guestConfigurationAssignmentName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTnAmE/pRoViDeRs/mIcRoSoFt.gUeStCoNfIgUrAtIoN/gUeStCoNfIgUrAtIoNaSsIgNmEnTs/gUeStCoNfIgUrAtIoNaSsIgNmEnTnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetName/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/guestConfigurationAssignmentName/reports",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTnAmE/pRoViDeRs/mIcRoSoFt.gUeStCoNfIgUrAtIoN/gUeStCoNfIgUrAtIoNaSsIgNmEnTs/gUeStCoNfIgUrAtIoNaSsIgNmEnTnAmE/rEpOrTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetName/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/guestConfigurationAssignmentName/reports/reportName",
			Expected: &ReportId{
				SubscriptionId:                   "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                "example-resource-group",
				VirtualMachineScaleSetName:       "virtualMachineScaleSetName",
				GuestConfigurationAssignmentName: "guestConfigurationAssignmentName",
				ReportName:                       "reportName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetName/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/guestConfigurationAssignmentName/reports/reportName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTnAmE/pRoViDeRs/mIcRoSoFt.gUeStCoNfIgUrAtIoN/gUeStCoNfIgUrAtIoNaSsIgNmEnTs/gUeStCoNfIgUrAtIoNaSsIgNmEnTnAmE/rEpOrTs/rEpOrTnAmE",
			Expected: &ReportId{
				SubscriptionId:                   "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                "eXaMpLe-rEsOuRcE-GrOuP",
				VirtualMachineScaleSetName:       "vIrTuAlMaChInEsCaLeSeTnAmE",
				GuestConfigurationAssignmentName: "gUeStCoNfIgUrAtIoNaSsIgNmEnTnAmE",
				ReportName:                       "rEpOrTnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTnAmE/pRoViDeRs/mIcRoSoFt.gUeStCoNfIgUrAtIoN/gUeStCoNfIgUrAtIoNaSsIgNmEnTs/gUeStCoNfIgUrAtIoNaSsIgNmEnTnAmE/rEpOrTs/rEpOrTnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportIDInsensitively(v.Input)
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

		if actual.VirtualMachineScaleSetName != v.Expected.VirtualMachineScaleSetName {
			t.Fatalf("Expected %q but got %q for VirtualMachineScaleSetName", v.Expected.VirtualMachineScaleSetName, actual.VirtualMachineScaleSetName)
		}

		if actual.GuestConfigurationAssignmentName != v.Expected.GuestConfigurationAssignmentName {
			t.Fatalf("Expected %q but got %q for GuestConfigurationAssignmentName", v.Expected.GuestConfigurationAssignmentName, actual.GuestConfigurationAssignmentName)
		}

		if actual.ReportName != v.Expected.ReportName {
			t.Fatalf("Expected %q but got %q for ReportName", v.Expected.ReportName, actual.ReportName)
		}

	}
}

func TestSegmentsForReportId(t *testing.T) {
	segments := ReportId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
