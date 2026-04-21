package restorabledroppeddatabases

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RestorableDroppedDatabaseId{}

func TestNewRestorableDroppedDatabaseID(t *testing.T) {
	id := NewRestorableDroppedDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "restorableDroppedDatabaseId")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.ServerName != "serverName" {
		t.Fatalf("Expected %q but got %q for Segment 'ServerName'", id.ServerName, "serverName")
	}

	if id.RestorableDroppedDatabaseId != "restorableDroppedDatabaseId" {
		t.Fatalf("Expected %q but got %q for Segment 'RestorableDroppedDatabaseId'", id.RestorableDroppedDatabaseId, "restorableDroppedDatabaseId")
	}
}

func TestFormatRestorableDroppedDatabaseID(t *testing.T) {
	actual := NewRestorableDroppedDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "restorableDroppedDatabaseId").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/restorableDroppedDatabases/restorableDroppedDatabaseId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRestorableDroppedDatabaseID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RestorableDroppedDatabaseId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/restorableDroppedDatabases",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/restorableDroppedDatabases/restorableDroppedDatabaseId",
			Expected: &RestorableDroppedDatabaseId{
				SubscriptionId:              "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:           "example-resource-group",
				ServerName:                  "serverName",
				RestorableDroppedDatabaseId: "restorableDroppedDatabaseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/restorableDroppedDatabases/restorableDroppedDatabaseId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRestorableDroppedDatabaseID(v.Input)
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

		if actual.ServerName != v.Expected.ServerName {
			t.Fatalf("Expected %q but got %q for ServerName", v.Expected.ServerName, actual.ServerName)
		}

		if actual.RestorableDroppedDatabaseId != v.Expected.RestorableDroppedDatabaseId {
			t.Fatalf("Expected %q but got %q for RestorableDroppedDatabaseId", v.Expected.RestorableDroppedDatabaseId, actual.RestorableDroppedDatabaseId)
		}

	}
}

func TestParseRestorableDroppedDatabaseIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RestorableDroppedDatabaseId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs/sErVeRnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/restorableDroppedDatabases",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs/sErVeRnAmE/rEsToRaBlEdRoPpEdDaTaBaSeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/restorableDroppedDatabases/restorableDroppedDatabaseId",
			Expected: &RestorableDroppedDatabaseId{
				SubscriptionId:              "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:           "example-resource-group",
				ServerName:                  "serverName",
				RestorableDroppedDatabaseId: "restorableDroppedDatabaseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/restorableDroppedDatabases/restorableDroppedDatabaseId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs/sErVeRnAmE/rEsToRaBlEdRoPpEdDaTaBaSeS/rEsToRaBlEdRoPpEdDaTaBaSeId",
			Expected: &RestorableDroppedDatabaseId{
				SubscriptionId:              "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:           "eXaMpLe-rEsOuRcE-GrOuP",
				ServerName:                  "sErVeRnAmE",
				RestorableDroppedDatabaseId: "rEsToRaBlEdRoPpEdDaTaBaSeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs/sErVeRnAmE/rEsToRaBlEdRoPpEdDaTaBaSeS/rEsToRaBlEdRoPpEdDaTaBaSeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRestorableDroppedDatabaseIDInsensitively(v.Input)
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

		if actual.ServerName != v.Expected.ServerName {
			t.Fatalf("Expected %q but got %q for ServerName", v.Expected.ServerName, actual.ServerName)
		}

		if actual.RestorableDroppedDatabaseId != v.Expected.RestorableDroppedDatabaseId {
			t.Fatalf("Expected %q but got %q for RestorableDroppedDatabaseId", v.Expected.RestorableDroppedDatabaseId, actual.RestorableDroppedDatabaseId)
		}

	}
}

func TestSegmentsForRestorableDroppedDatabaseId(t *testing.T) {
	segments := RestorableDroppedDatabaseId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RestorableDroppedDatabaseId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
