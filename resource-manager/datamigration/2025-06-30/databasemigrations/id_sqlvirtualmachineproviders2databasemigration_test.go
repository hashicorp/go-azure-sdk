package databasemigrations

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SqlVirtualMachineProviders2DatabaseMigrationId{}

func TestNewSqlVirtualMachineProviders2DatabaseMigrationID(t *testing.T) {
	id := NewSqlVirtualMachineProviders2DatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlVirtualMachineName", "databaseMigrationName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.SqlVirtualMachineName != "sqlVirtualMachineName" {
		t.Fatalf("Expected %q but got %q for Segment 'SqlVirtualMachineName'", id.SqlVirtualMachineName, "sqlVirtualMachineName")
	}

	if id.DatabaseMigrationName != "databaseMigrationName" {
		t.Fatalf("Expected %q but got %q for Segment 'DatabaseMigrationName'", id.DatabaseMigrationName, "databaseMigrationName")
	}
}

func TestFormatSqlVirtualMachineProviders2DatabaseMigrationID(t *testing.T) {
	actual := NewSqlVirtualMachineProviders2DatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlVirtualMachineName", "databaseMigrationName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/sqlVirtualMachineName/providers/Microsoft.DataMigration/databaseMigrations/databaseMigrationName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSqlVirtualMachineProviders2DatabaseMigrationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SqlVirtualMachineProviders2DatabaseMigrationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/sqlVirtualMachineName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/sqlVirtualMachineName/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/sqlVirtualMachineName/providers/Microsoft.DataMigration",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/sqlVirtualMachineName/providers/Microsoft.DataMigration/databaseMigrations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/sqlVirtualMachineName/providers/Microsoft.DataMigration/databaseMigrations/databaseMigrationName",
			Expected: &SqlVirtualMachineProviders2DatabaseMigrationId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "example-resource-group",
				SqlVirtualMachineName: "sqlVirtualMachineName",
				DatabaseMigrationName: "databaseMigrationName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/sqlVirtualMachineName/providers/Microsoft.DataMigration/databaseMigrations/databaseMigrationName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSqlVirtualMachineProviders2DatabaseMigrationID(v.Input)
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

		if actual.SqlVirtualMachineName != v.Expected.SqlVirtualMachineName {
			t.Fatalf("Expected %q but got %q for SqlVirtualMachineName", v.Expected.SqlVirtualMachineName, actual.SqlVirtualMachineName)
		}

		if actual.DatabaseMigrationName != v.Expected.DatabaseMigrationName {
			t.Fatalf("Expected %q but got %q for DatabaseMigrationName", v.Expected.DatabaseMigrationName, actual.DatabaseMigrationName)
		}

	}
}

func TestParseSqlVirtualMachineProviders2DatabaseMigrationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SqlVirtualMachineProviders2DatabaseMigrationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQlViRtUaLmAcHiNe/sQlViRtUaLmAcHiNeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/sqlVirtualMachineName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQlViRtUaLmAcHiNe/sQlViRtUaLmAcHiNeS/sQlViRtUaLmAcHiNeNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/sqlVirtualMachineName/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQlViRtUaLmAcHiNe/sQlViRtUaLmAcHiNeS/sQlViRtUaLmAcHiNeNaMe/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/sqlVirtualMachineName/providers/Microsoft.DataMigration",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQlViRtUaLmAcHiNe/sQlViRtUaLmAcHiNeS/sQlViRtUaLmAcHiNeNaMe/pRoViDeRs/mIcRoSoFt.dAtAmIgRaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/sqlVirtualMachineName/providers/Microsoft.DataMigration/databaseMigrations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQlViRtUaLmAcHiNe/sQlViRtUaLmAcHiNeS/sQlViRtUaLmAcHiNeNaMe/pRoViDeRs/mIcRoSoFt.dAtAmIgRaTiOn/dAtAbAsEmIgRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/sqlVirtualMachineName/providers/Microsoft.DataMigration/databaseMigrations/databaseMigrationName",
			Expected: &SqlVirtualMachineProviders2DatabaseMigrationId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "example-resource-group",
				SqlVirtualMachineName: "sqlVirtualMachineName",
				DatabaseMigrationName: "databaseMigrationName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/sqlVirtualMachineName/providers/Microsoft.DataMigration/databaseMigrations/databaseMigrationName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQlViRtUaLmAcHiNe/sQlViRtUaLmAcHiNeS/sQlViRtUaLmAcHiNeNaMe/pRoViDeRs/mIcRoSoFt.dAtAmIgRaTiOn/dAtAbAsEmIgRaTiOnS/dAtAbAsEmIgRaTiOnNaMe",
			Expected: &SqlVirtualMachineProviders2DatabaseMigrationId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "eXaMpLe-rEsOuRcE-GrOuP",
				SqlVirtualMachineName: "sQlViRtUaLmAcHiNeNaMe",
				DatabaseMigrationName: "dAtAbAsEmIgRaTiOnNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQlViRtUaLmAcHiNe/sQlViRtUaLmAcHiNeS/sQlViRtUaLmAcHiNeNaMe/pRoViDeRs/mIcRoSoFt.dAtAmIgRaTiOn/dAtAbAsEmIgRaTiOnS/dAtAbAsEmIgRaTiOnNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSqlVirtualMachineProviders2DatabaseMigrationIDInsensitively(v.Input)
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

		if actual.SqlVirtualMachineName != v.Expected.SqlVirtualMachineName {
			t.Fatalf("Expected %q but got %q for SqlVirtualMachineName", v.Expected.SqlVirtualMachineName, actual.SqlVirtualMachineName)
		}

		if actual.DatabaseMigrationName != v.Expected.DatabaseMigrationName {
			t.Fatalf("Expected %q but got %q for DatabaseMigrationName", v.Expected.DatabaseMigrationName, actual.DatabaseMigrationName)
		}

	}
}

func TestSegmentsForSqlVirtualMachineProviders2DatabaseMigrationId(t *testing.T) {
	segments := SqlVirtualMachineProviders2DatabaseMigrationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SqlVirtualMachineProviders2DatabaseMigrationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
