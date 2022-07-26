package virtualmachinescalesetvmextensions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = VirtualMachineScaleSetVirtualMachineId{}

func TestNewVirtualMachineScaleSetVirtualMachineID(t *testing.T) {
	id := NewVirtualMachineScaleSetVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetValue", "instanceIdValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.VirtualMachineScaleSetName != "virtualMachineScaleSetValue" {
		t.Fatalf("Expected %q but got %q for Segment 'VirtualMachineScaleSetName'", id.VirtualMachineScaleSetName, "virtualMachineScaleSetValue")
	}

	if id.InstanceId != "instanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'InstanceId'", id.InstanceId, "instanceIdValue")
	}
}

func TestFormatVirtualMachineScaleSetVirtualMachineID(t *testing.T) {
	actual := NewVirtualMachineScaleSetVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetValue", "instanceIdValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/instanceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseVirtualMachineScaleSetVirtualMachineID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VirtualMachineScaleSetVirtualMachineId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/instanceIdValue",
			Expected: &VirtualMachineScaleSetVirtualMachineId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "example-resource-group",
				VirtualMachineScaleSetName: "virtualMachineScaleSetValue",
				InstanceId:                 "instanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/instanceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVirtualMachineScaleSetVirtualMachineID(v.Input)
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

		if actual.InstanceId != v.Expected.InstanceId {
			t.Fatalf("Expected %q but got %q for InstanceId", v.Expected.InstanceId, actual.InstanceId)
		}

	}
}

func TestParseVirtualMachineScaleSetVirtualMachineIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VirtualMachineScaleSetVirtualMachineId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTvAlUe/vIrTuAlMaChInEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/instanceIdValue",
			Expected: &VirtualMachineScaleSetVirtualMachineId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "example-resource-group",
				VirtualMachineScaleSetName: "virtualMachineScaleSetValue",
				InstanceId:                 "instanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/virtualMachineScaleSets/virtualMachineScaleSetValue/virtualMachines/instanceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTvAlUe/vIrTuAlMaChInEs/iNsTaNcEiDvAlUe",
			Expected: &VirtualMachineScaleSetVirtualMachineId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "eXaMpLe-rEsOuRcE-GrOuP",
				VirtualMachineScaleSetName: "vIrTuAlMaChInEsCaLeSeTvAlUe",
				InstanceId:                 "iNsTaNcEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/vIrTuAlMaChInEsCaLeSeTs/vIrTuAlMaChInEsCaLeSeTvAlUe/vIrTuAlMaChInEs/iNsTaNcEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVirtualMachineScaleSetVirtualMachineIDInsensitively(v.Input)
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

		if actual.InstanceId != v.Expected.InstanceId {
			t.Fatalf("Expected %q but got %q for InstanceId", v.Expected.InstanceId, actual.InstanceId)
		}

	}
}

func TestSegmentsForVirtualMachineScaleSetVirtualMachineId(t *testing.T) {
	segments := VirtualMachineScaleSetVirtualMachineId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("VirtualMachineScaleSetVirtualMachineId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
