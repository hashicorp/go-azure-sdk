package syncmembers

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SyncMemberId{}

func TestNewSyncMemberID(t *testing.T) {
	id := NewSyncMemberID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "syncGroupName", "syncMemberName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.ServerName != "serverName" {
		t.Fatalf("Expected %q but got %q for Segment 'ServerName'", id.ServerName, "serverName")
	}

	if id.DatabaseName != "databaseName" {
		t.Fatalf("Expected %q but got %q for Segment 'DatabaseName'", id.DatabaseName, "databaseName")
	}

	if id.SyncGroupName != "syncGroupName" {
		t.Fatalf("Expected %q but got %q for Segment 'SyncGroupName'", id.SyncGroupName, "syncGroupName")
	}

	if id.SyncMemberName != "syncMemberName" {
		t.Fatalf("Expected %q but got %q for Segment 'SyncMemberName'", id.SyncMemberName, "syncMemberName")
	}
}

func TestFormatSyncMemberID(t *testing.T) {
	actual := NewSyncMemberID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "syncGroupName", "syncMemberName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/databases/databaseName/syncGroups/syncGroupName/syncMembers/syncMemberName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSyncMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SyncMemberId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/databases",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/databases/databaseName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/databases/databaseName/syncGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/databases/databaseName/syncGroups/syncGroupName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/databases/databaseName/syncGroups/syncGroupName/syncMembers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/databases/databaseName/syncGroups/syncGroupName/syncMembers/syncMemberName",
			Expected: &SyncMemberId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "example-resource-group",
				ServerName:        "serverName",
				DatabaseName:      "databaseName",
				SyncGroupName:     "syncGroupName",
				SyncMemberName:    "syncMemberName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/databases/databaseName/syncGroups/syncGroupName/syncMembers/syncMemberName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSyncMemberID(v.Input)
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

		if actual.DatabaseName != v.Expected.DatabaseName {
			t.Fatalf("Expected %q but got %q for DatabaseName", v.Expected.DatabaseName, actual.DatabaseName)
		}

		if actual.SyncGroupName != v.Expected.SyncGroupName {
			t.Fatalf("Expected %q but got %q for SyncGroupName", v.Expected.SyncGroupName, actual.SyncGroupName)
		}

		if actual.SyncMemberName != v.Expected.SyncMemberName {
			t.Fatalf("Expected %q but got %q for SyncMemberName", v.Expected.SyncMemberName, actual.SyncMemberName)
		}

	}
}

func TestParseSyncMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SyncMemberId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/databases",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs/sErVeRnAmE/dAtAbAsEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/databases/databaseName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs/sErVeRnAmE/dAtAbAsEs/dAtAbAsEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/databases/databaseName/syncGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs/sErVeRnAmE/dAtAbAsEs/dAtAbAsEnAmE/sYnCgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/databases/databaseName/syncGroups/syncGroupName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs/sErVeRnAmE/dAtAbAsEs/dAtAbAsEnAmE/sYnCgRoUpS/sYnCgRoUpNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/databases/databaseName/syncGroups/syncGroupName/syncMembers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs/sErVeRnAmE/dAtAbAsEs/dAtAbAsEnAmE/sYnCgRoUpS/sYnCgRoUpNaMe/sYnCmEmBeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/databases/databaseName/syncGroups/syncGroupName/syncMembers/syncMemberName",
			Expected: &SyncMemberId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "example-resource-group",
				ServerName:        "serverName",
				DatabaseName:      "databaseName",
				SyncGroupName:     "syncGroupName",
				SyncMemberName:    "syncMemberName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/databases/databaseName/syncGroups/syncGroupName/syncMembers/syncMemberName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs/sErVeRnAmE/dAtAbAsEs/dAtAbAsEnAmE/sYnCgRoUpS/sYnCgRoUpNaMe/sYnCmEmBeRs/sYnCmEmBeRnAmE",
			Expected: &SyncMemberId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "eXaMpLe-rEsOuRcE-GrOuP",
				ServerName:        "sErVeRnAmE",
				DatabaseName:      "dAtAbAsEnAmE",
				SyncGroupName:     "sYnCgRoUpNaMe",
				SyncMemberName:    "sYnCmEmBeRnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs/sErVeRnAmE/dAtAbAsEs/dAtAbAsEnAmE/sYnCgRoUpS/sYnCgRoUpNaMe/sYnCmEmBeRs/sYnCmEmBeRnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSyncMemberIDInsensitively(v.Input)
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

		if actual.DatabaseName != v.Expected.DatabaseName {
			t.Fatalf("Expected %q but got %q for DatabaseName", v.Expected.DatabaseName, actual.DatabaseName)
		}

		if actual.SyncGroupName != v.Expected.SyncGroupName {
			t.Fatalf("Expected %q but got %q for SyncGroupName", v.Expected.SyncGroupName, actual.SyncGroupName)
		}

		if actual.SyncMemberName != v.Expected.SyncMemberName {
			t.Fatalf("Expected %q but got %q for SyncMemberName", v.Expected.SyncMemberName, actual.SyncMemberName)
		}

	}
}

func TestSegmentsForSyncMemberId(t *testing.T) {
	segments := SyncMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SyncMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
