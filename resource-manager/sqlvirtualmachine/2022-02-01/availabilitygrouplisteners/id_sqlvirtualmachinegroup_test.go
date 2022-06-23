package availabilitygrouplisteners

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = SqlVirtualMachineGroupId{}

func TestNewSqlVirtualMachineGroupID(t *testing.T) {
	id := NewSqlVirtualMachineGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlVirtualMachineGroupValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.SqlVirtualMachineGroupName != "sqlVirtualMachineGroupValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SqlVirtualMachineGroupName'", id.SqlVirtualMachineGroupName, "sqlVirtualMachineGroupValue")
	}
}

func TestFormatSqlVirtualMachineGroupID(t *testing.T) {
	actual := NewSqlVirtualMachineGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlVirtualMachineGroupValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/sqlVirtualMachineGroupValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSqlVirtualMachineGroupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SqlVirtualMachineGroupId
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
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/sqlVirtualMachineGroupValue",
			Expected: &SqlVirtualMachineGroupId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "example-resource-group",
				SqlVirtualMachineGroupName: "sqlVirtualMachineGroupValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/sqlVirtualMachineGroupValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSqlVirtualMachineGroupID(v.Input)
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

	}
}

func TestParseSqlVirtualMachineGroupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SqlVirtualMachineGroupId
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
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/sqlVirtualMachineGroupValue",
			Expected: &SqlVirtualMachineGroupId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "example-resource-group",
				SqlVirtualMachineGroupName: "sqlVirtualMachineGroupValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/sqlVirtualMachineGroupValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQlViRtUaLmAcHiNe/sQlViRtUaLmAcHiNeGrOuPs/sQlViRtUaLmAcHiNeGrOuPvAlUe",
			Expected: &SqlVirtualMachineGroupId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "eXaMpLe-rEsOuRcE-GrOuP",
				SqlVirtualMachineGroupName: "sQlViRtUaLmAcHiNeGrOuPvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQlViRtUaLmAcHiNe/sQlViRtUaLmAcHiNeGrOuPs/sQlViRtUaLmAcHiNeGrOuPvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSqlVirtualMachineGroupIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForSqlVirtualMachineGroupId(t *testing.T) {
	segments := SqlVirtualMachineGroupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SqlVirtualMachineGroupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
