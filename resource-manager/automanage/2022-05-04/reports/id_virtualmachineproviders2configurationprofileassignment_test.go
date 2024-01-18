package reports

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &VirtualMachineProviders2ConfigurationProfileAssignmentId{}

func TestNewVirtualMachineProviders2ConfigurationProfileAssignmentID(t *testing.T) {
	id := NewVirtualMachineProviders2ConfigurationProfileAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue", "configurationProfileAssignmentValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.VirtualMachineName != "virtualMachineValue" {
		t.Fatalf("Expected %q but got %q for Segment 'VirtualMachineName'", id.VirtualMachineName, "virtualMachineValue")
	}

	if id.ConfigurationProfileAssignmentName != "configurationProfileAssignmentValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ConfigurationProfileAssignmentName'", id.ConfigurationProfileAssignmentName, "configurationProfileAssignmentValue")
	}
}

func TestFormatVirtualMachineProviders2ConfigurationProfileAssignmentID(t *testing.T) {
	actual := NewVirtualMachineProviders2ConfigurationProfileAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue", "configurationProfileAssignmentValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachines/virtualMachineValue/providers/Microsoft.AutoManage/configurationProfileAssignments/configurationProfileAssignmentValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseVirtualMachineProviders2ConfigurationProfileAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VirtualMachineProviders2ConfigurationProfileAssignmentId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachines",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachines/virtualMachineValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachines/virtualMachineValue/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachines/virtualMachineValue/providers/Microsoft.AutoManage",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachines/virtualMachineValue/providers/Microsoft.AutoManage/configurationProfileAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachines/virtualMachineValue/providers/Microsoft.AutoManage/configurationProfileAssignments/configurationProfileAssignmentValue",
			Expected: &VirtualMachineProviders2ConfigurationProfileAssignmentId{
				SubscriptionId:                     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                  "example-resource-group",
				VirtualMachineName:                 "virtualMachineValue",
				ConfigurationProfileAssignmentName: "configurationProfileAssignmentValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachines/virtualMachineValue/providers/Microsoft.AutoManage/configurationProfileAssignments/configurationProfileAssignmentValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVirtualMachineProviders2ConfigurationProfileAssignmentID(v.Input)
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

		if actual.ConfigurationProfileAssignmentName != v.Expected.ConfigurationProfileAssignmentName {
			t.Fatalf("Expected %q but got %q for ConfigurationProfileAssignmentName", v.Expected.ConfigurationProfileAssignmentName, actual.ConfigurationProfileAssignmentName)
		}

	}
}

func TestParseVirtualMachineProviders2ConfigurationProfileAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VirtualMachineProviders2ConfigurationProfileAssignmentId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachines",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachines/virtualMachineValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEs/vIrTuAlMaChInEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachines/virtualMachineValue/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEs/vIrTuAlMaChInEvAlUe/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachines/virtualMachineValue/providers/Microsoft.AutoManage",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEs/vIrTuAlMaChInEvAlUe/pRoViDeRs/mIcRoSoFt.aUtOmAnAgE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachines/virtualMachineValue/providers/Microsoft.AutoManage/configurationProfileAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEs/vIrTuAlMaChInEvAlUe/pRoViDeRs/mIcRoSoFt.aUtOmAnAgE/cOnFiGuRaTiOnPrOfIlEaSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachines/virtualMachineValue/providers/Microsoft.AutoManage/configurationProfileAssignments/configurationProfileAssignmentValue",
			Expected: &VirtualMachineProviders2ConfigurationProfileAssignmentId{
				SubscriptionId:                     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                  "example-resource-group",
				VirtualMachineName:                 "virtualMachineValue",
				ConfigurationProfileAssignmentName: "configurationProfileAssignmentValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachines/virtualMachineValue/providers/Microsoft.AutoManage/configurationProfileAssignments/configurationProfileAssignmentValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEs/vIrTuAlMaChInEvAlUe/pRoViDeRs/mIcRoSoFt.aUtOmAnAgE/cOnFiGuRaTiOnPrOfIlEaSsIgNmEnTs/cOnFiGuRaTiOnPrOfIlEaSsIgNmEnTvAlUe",
			Expected: &VirtualMachineProviders2ConfigurationProfileAssignmentId{
				SubscriptionId:                     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                  "eXaMpLe-rEsOuRcE-GrOuP",
				VirtualMachineName:                 "vIrTuAlMaChInEvAlUe",
				ConfigurationProfileAssignmentName: "cOnFiGuRaTiOnPrOfIlEaSsIgNmEnTvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEs/vIrTuAlMaChInEvAlUe/pRoViDeRs/mIcRoSoFt.aUtOmAnAgE/cOnFiGuRaTiOnPrOfIlEaSsIgNmEnTs/cOnFiGuRaTiOnPrOfIlEaSsIgNmEnTvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVirtualMachineProviders2ConfigurationProfileAssignmentIDInsensitively(v.Input)
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

		if actual.ConfigurationProfileAssignmentName != v.Expected.ConfigurationProfileAssignmentName {
			t.Fatalf("Expected %q but got %q for ConfigurationProfileAssignmentName", v.Expected.ConfigurationProfileAssignmentName, actual.ConfigurationProfileAssignmentName)
		}

	}
}

func TestSegmentsForVirtualMachineProviders2ConfigurationProfileAssignmentId(t *testing.T) {
	segments := VirtualMachineProviders2ConfigurationProfileAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("VirtualMachineProviders2ConfigurationProfileAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
