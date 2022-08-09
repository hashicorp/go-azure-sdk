package incrementalrestorepoints

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = DiskRestorePointId{}

func TestNewDiskRestorePointID(t *testing.T) {
	id := NewDiskRestorePointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "restorePointCollectionValue", "vmRestorePointValue", "diskRestorePointValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.RestorePointCollectionName != "restorePointCollectionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RestorePointCollectionName'", id.RestorePointCollectionName, "restorePointCollectionValue")
	}

	if id.VmRestorePointName != "vmRestorePointValue" {
		t.Fatalf("Expected %q but got %q for Segment 'VmRestorePointName'", id.VmRestorePointName, "vmRestorePointValue")
	}

	if id.DiskRestorePointName != "diskRestorePointValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DiskRestorePointName'", id.DiskRestorePointName, "diskRestorePointValue")
	}
}

func TestFormatDiskRestorePointID(t *testing.T) {
	actual := NewDiskRestorePointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "restorePointCollectionValue", "vmRestorePointValue", "diskRestorePointValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/restorePointCollections/restorePointCollectionValue/restorePoints/vmRestorePointValue/diskRestorePoints/diskRestorePointValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDiskRestorePointID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DiskRestorePointId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/restorePointCollections",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/restorePointCollections/restorePointCollectionValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/restorePointCollections/restorePointCollectionValue/restorePoints",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/restorePointCollections/restorePointCollectionValue/restorePoints/vmRestorePointValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/restorePointCollections/restorePointCollectionValue/restorePoints/vmRestorePointValue/diskRestorePoints",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/restorePointCollections/restorePointCollectionValue/restorePoints/vmRestorePointValue/diskRestorePoints/diskRestorePointValue",
			Expected: &DiskRestorePointId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "example-resource-group",
				RestorePointCollectionName: "restorePointCollectionValue",
				VmRestorePointName:         "vmRestorePointValue",
				DiskRestorePointName:       "diskRestorePointValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/restorePointCollections/restorePointCollectionValue/restorePoints/vmRestorePointValue/diskRestorePoints/diskRestorePointValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDiskRestorePointID(v.Input)
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

		if actual.RestorePointCollectionName != v.Expected.RestorePointCollectionName {
			t.Fatalf("Expected %q but got %q for RestorePointCollectionName", v.Expected.RestorePointCollectionName, actual.RestorePointCollectionName)
		}

		if actual.VmRestorePointName != v.Expected.VmRestorePointName {
			t.Fatalf("Expected %q but got %q for VmRestorePointName", v.Expected.VmRestorePointName, actual.VmRestorePointName)
		}

		if actual.DiskRestorePointName != v.Expected.DiskRestorePointName {
			t.Fatalf("Expected %q but got %q for DiskRestorePointName", v.Expected.DiskRestorePointName, actual.DiskRestorePointName)
		}

	}
}

func TestParseDiskRestorePointIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DiskRestorePointId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/restorePointCollections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/rEsToRePoInTcOlLeCtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/restorePointCollections/restorePointCollectionValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/rEsToRePoInTcOlLeCtIoNs/rEsToRePoInTcOlLeCtIoNvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/restorePointCollections/restorePointCollectionValue/restorePoints",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/rEsToRePoInTcOlLeCtIoNs/rEsToRePoInTcOlLeCtIoNvAlUe/rEsToRePoInTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/restorePointCollections/restorePointCollectionValue/restorePoints/vmRestorePointValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/rEsToRePoInTcOlLeCtIoNs/rEsToRePoInTcOlLeCtIoNvAlUe/rEsToRePoInTs/vMrEsToRePoInTvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/restorePointCollections/restorePointCollectionValue/restorePoints/vmRestorePointValue/diskRestorePoints",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/rEsToRePoInTcOlLeCtIoNs/rEsToRePoInTcOlLeCtIoNvAlUe/rEsToRePoInTs/vMrEsToRePoInTvAlUe/dIsKrEsToRePoInTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/restorePointCollections/restorePointCollectionValue/restorePoints/vmRestorePointValue/diskRestorePoints/diskRestorePointValue",
			Expected: &DiskRestorePointId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "example-resource-group",
				RestorePointCollectionName: "restorePointCollectionValue",
				VmRestorePointName:         "vmRestorePointValue",
				DiskRestorePointName:       "diskRestorePointValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/restorePointCollections/restorePointCollectionValue/restorePoints/vmRestorePointValue/diskRestorePoints/diskRestorePointValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/rEsToRePoInTcOlLeCtIoNs/rEsToRePoInTcOlLeCtIoNvAlUe/rEsToRePoInTs/vMrEsToRePoInTvAlUe/dIsKrEsToRePoInTs/dIsKrEsToRePoInTvAlUe",
			Expected: &DiskRestorePointId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "eXaMpLe-rEsOuRcE-GrOuP",
				RestorePointCollectionName: "rEsToRePoInTcOlLeCtIoNvAlUe",
				VmRestorePointName:         "vMrEsToRePoInTvAlUe",
				DiskRestorePointName:       "dIsKrEsToRePoInTvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/rEsToRePoInTcOlLeCtIoNs/rEsToRePoInTcOlLeCtIoNvAlUe/rEsToRePoInTs/vMrEsToRePoInTvAlUe/dIsKrEsToRePoInTs/dIsKrEsToRePoInTvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDiskRestorePointIDInsensitively(v.Input)
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

		if actual.RestorePointCollectionName != v.Expected.RestorePointCollectionName {
			t.Fatalf("Expected %q but got %q for RestorePointCollectionName", v.Expected.RestorePointCollectionName, actual.RestorePointCollectionName)
		}

		if actual.VmRestorePointName != v.Expected.VmRestorePointName {
			t.Fatalf("Expected %q but got %q for VmRestorePointName", v.Expected.VmRestorePointName, actual.VmRestorePointName)
		}

		if actual.DiskRestorePointName != v.Expected.DiskRestorePointName {
			t.Fatalf("Expected %q but got %q for DiskRestorePointName", v.Expected.DiskRestorePointName, actual.DiskRestorePointName)
		}

	}
}

func TestSegmentsForDiskRestorePointId(t *testing.T) {
	segments := DiskRestorePointId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DiskRestorePointId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
