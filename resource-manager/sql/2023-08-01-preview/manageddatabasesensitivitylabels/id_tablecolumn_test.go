package manageddatabasesensitivitylabels

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &TableColumnId{}

func TestNewTableColumnID(t *testing.T) {
	id := NewTableColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseName", "schemaName", "tableName", "columnName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.ManagedInstanceName != "managedInstanceName" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedInstanceName'", id.ManagedInstanceName, "managedInstanceName")
	}

	if id.DatabaseName != "databaseName" {
		t.Fatalf("Expected %q but got %q for Segment 'DatabaseName'", id.DatabaseName, "databaseName")
	}

	if id.SchemaName != "schemaName" {
		t.Fatalf("Expected %q but got %q for Segment 'SchemaName'", id.SchemaName, "schemaName")
	}

	if id.TableName != "tableName" {
		t.Fatalf("Expected %q but got %q for Segment 'TableName'", id.TableName, "tableName")
	}

	if id.ColumnName != "columnName" {
		t.Fatalf("Expected %q but got %q for Segment 'ColumnName'", id.ColumnName, "columnName")
	}
}

func TestFormatTableColumnID(t *testing.T) {
	actual := NewTableColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseName", "schemaName", "tableName", "columnName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName/databases/databaseName/schemas/schemaName/tables/tableName/columns/columnName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseTableColumnID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TableColumnId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName/databases",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName/databases/databaseName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName/databases/databaseName/schemas",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName/databases/databaseName/schemas/schemaName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName/databases/databaseName/schemas/schemaName/tables",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName/databases/databaseName/schemas/schemaName/tables/tableName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName/databases/databaseName/schemas/schemaName/tables/tableName/columns",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName/databases/databaseName/schemas/schemaName/tables/tableName/columns/columnName",
			Expected: &TableColumnId{
				SubscriptionId:      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:   "example-resource-group",
				ManagedInstanceName: "managedInstanceName",
				DatabaseName:        "databaseName",
				SchemaName:          "schemaName",
				TableName:           "tableName",
				ColumnName:          "columnName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName/databases/databaseName/schemas/schemaName/tables/tableName/columns/columnName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTableColumnID(v.Input)
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

		if actual.ManagedInstanceName != v.Expected.ManagedInstanceName {
			t.Fatalf("Expected %q but got %q for ManagedInstanceName", v.Expected.ManagedInstanceName, actual.ManagedInstanceName)
		}

		if actual.DatabaseName != v.Expected.DatabaseName {
			t.Fatalf("Expected %q but got %q for DatabaseName", v.Expected.DatabaseName, actual.DatabaseName)
		}

		if actual.SchemaName != v.Expected.SchemaName {
			t.Fatalf("Expected %q but got %q for SchemaName", v.Expected.SchemaName, actual.SchemaName)
		}

		if actual.TableName != v.Expected.TableName {
			t.Fatalf("Expected %q but got %q for TableName", v.Expected.TableName, actual.TableName)
		}

		if actual.ColumnName != v.Expected.ColumnName {
			t.Fatalf("Expected %q but got %q for ColumnName", v.Expected.ColumnName, actual.ColumnName)
		}

	}
}

func TestParseTableColumnIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TableColumnId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/mAnAgEdInStAnCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/mAnAgEdInStAnCeS/mAnAgEdInStAnCeNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName/databases",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/mAnAgEdInStAnCeS/mAnAgEdInStAnCeNaMe/dAtAbAsEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName/databases/databaseName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/mAnAgEdInStAnCeS/mAnAgEdInStAnCeNaMe/dAtAbAsEs/dAtAbAsEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName/databases/databaseName/schemas",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/mAnAgEdInStAnCeS/mAnAgEdInStAnCeNaMe/dAtAbAsEs/dAtAbAsEnAmE/sChEmAs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName/databases/databaseName/schemas/schemaName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/mAnAgEdInStAnCeS/mAnAgEdInStAnCeNaMe/dAtAbAsEs/dAtAbAsEnAmE/sChEmAs/sChEmAnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName/databases/databaseName/schemas/schemaName/tables",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/mAnAgEdInStAnCeS/mAnAgEdInStAnCeNaMe/dAtAbAsEs/dAtAbAsEnAmE/sChEmAs/sChEmAnAmE/tAbLeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName/databases/databaseName/schemas/schemaName/tables/tableName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/mAnAgEdInStAnCeS/mAnAgEdInStAnCeNaMe/dAtAbAsEs/dAtAbAsEnAmE/sChEmAs/sChEmAnAmE/tAbLeS/tAbLeNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName/databases/databaseName/schemas/schemaName/tables/tableName/columns",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/mAnAgEdInStAnCeS/mAnAgEdInStAnCeNaMe/dAtAbAsEs/dAtAbAsEnAmE/sChEmAs/sChEmAnAmE/tAbLeS/tAbLeNaMe/cOlUmNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName/databases/databaseName/schemas/schemaName/tables/tableName/columns/columnName",
			Expected: &TableColumnId{
				SubscriptionId:      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:   "example-resource-group",
				ManagedInstanceName: "managedInstanceName",
				DatabaseName:        "databaseName",
				SchemaName:          "schemaName",
				TableName:           "tableName",
				ColumnName:          "columnName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/managedInstances/managedInstanceName/databases/databaseName/schemas/schemaName/tables/tableName/columns/columnName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/mAnAgEdInStAnCeS/mAnAgEdInStAnCeNaMe/dAtAbAsEs/dAtAbAsEnAmE/sChEmAs/sChEmAnAmE/tAbLeS/tAbLeNaMe/cOlUmNs/cOlUmNnAmE",
			Expected: &TableColumnId{
				SubscriptionId:      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:   "eXaMpLe-rEsOuRcE-GrOuP",
				ManagedInstanceName: "mAnAgEdInStAnCeNaMe",
				DatabaseName:        "dAtAbAsEnAmE",
				SchemaName:          "sChEmAnAmE",
				TableName:           "tAbLeNaMe",
				ColumnName:          "cOlUmNnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/mAnAgEdInStAnCeS/mAnAgEdInStAnCeNaMe/dAtAbAsEs/dAtAbAsEnAmE/sChEmAs/sChEmAnAmE/tAbLeS/tAbLeNaMe/cOlUmNs/cOlUmNnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTableColumnIDInsensitively(v.Input)
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

		if actual.ManagedInstanceName != v.Expected.ManagedInstanceName {
			t.Fatalf("Expected %q but got %q for ManagedInstanceName", v.Expected.ManagedInstanceName, actual.ManagedInstanceName)
		}

		if actual.DatabaseName != v.Expected.DatabaseName {
			t.Fatalf("Expected %q but got %q for DatabaseName", v.Expected.DatabaseName, actual.DatabaseName)
		}

		if actual.SchemaName != v.Expected.SchemaName {
			t.Fatalf("Expected %q but got %q for SchemaName", v.Expected.SchemaName, actual.SchemaName)
		}

		if actual.TableName != v.Expected.TableName {
			t.Fatalf("Expected %q but got %q for TableName", v.Expected.TableName, actual.TableName)
		}

		if actual.ColumnName != v.Expected.ColumnName {
			t.Fatalf("Expected %q but got %q for ColumnName", v.Expected.ColumnName, actual.ColumnName)
		}

	}
}

func TestSegmentsForTableColumnId(t *testing.T) {
	segments := TableColumnId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("TableColumnId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
