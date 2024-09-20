package autonomousdatabasebackups

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AutonomousDatabaseBackupId{}

func TestNewAutonomousDatabaseBackupID(t *testing.T) {
	id := NewAutonomousDatabaseBackupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "autonomousdatabasename", "adbbackupid")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.AutonomousDatabaseName != "autonomousdatabasename" {
		t.Fatalf("Expected %q but got %q for Segment 'AutonomousDatabaseName'", id.AutonomousDatabaseName, "autonomousdatabasename")
	}

	if id.AutonomousDatabaseBackupName != "adbbackupid" {
		t.Fatalf("Expected %q but got %q for Segment 'AutonomousDatabaseBackupName'", id.AutonomousDatabaseBackupName, "adbbackupid")
	}
}

func TestFormatAutonomousDatabaseBackupID(t *testing.T) {
	actual := NewAutonomousDatabaseBackupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "autonomousdatabasename", "adbbackupid").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/autonomousDatabases/autonomousdatabasename/autonomousDatabaseBackups/adbbackupid"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAutonomousDatabaseBackupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AutonomousDatabaseBackupId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/autonomousDatabases",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/autonomousDatabases/autonomousdatabasename",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/autonomousDatabases/autonomousdatabasename/autonomousDatabaseBackups",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/autonomousDatabases/autonomousdatabasename/autonomousDatabaseBackups/adbbackupid",
			Expected: &AutonomousDatabaseBackupId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:            "example-resource-group",
				AutonomousDatabaseName:       "autonomousdatabasename",
				AutonomousDatabaseBackupName: "adbbackupid",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/autonomousDatabases/autonomousdatabasename/autonomousDatabaseBackups/adbbackupid/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAutonomousDatabaseBackupID(v.Input)
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

		if actual.AutonomousDatabaseName != v.Expected.AutonomousDatabaseName {
			t.Fatalf("Expected %q but got %q for AutonomousDatabaseName", v.Expected.AutonomousDatabaseName, actual.AutonomousDatabaseName)
		}

		if actual.AutonomousDatabaseBackupName != v.Expected.AutonomousDatabaseBackupName {
			t.Fatalf("Expected %q but got %q for AutonomousDatabaseBackupName", v.Expected.AutonomousDatabaseBackupName, actual.AutonomousDatabaseBackupName)
		}

	}
}

func TestParseAutonomousDatabaseBackupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AutonomousDatabaseBackupId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/oRaClE.DaTaBaSe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/autonomousDatabases",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/oRaClE.DaTaBaSe/aUtOnOmOuSdAtAbAsEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/autonomousDatabases/autonomousdatabasename",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/oRaClE.DaTaBaSe/aUtOnOmOuSdAtAbAsEs/aUtOnOmOuSdAtAbAsEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/autonomousDatabases/autonomousdatabasename/autonomousDatabaseBackups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/oRaClE.DaTaBaSe/aUtOnOmOuSdAtAbAsEs/aUtOnOmOuSdAtAbAsEnAmE/aUtOnOmOuSdAtAbAsEbAcKuPs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/autonomousDatabases/autonomousdatabasename/autonomousDatabaseBackups/adbbackupid",
			Expected: &AutonomousDatabaseBackupId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:            "example-resource-group",
				AutonomousDatabaseName:       "autonomousdatabasename",
				AutonomousDatabaseBackupName: "adbbackupid",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/autonomousDatabases/autonomousdatabasename/autonomousDatabaseBackups/adbbackupid/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/oRaClE.DaTaBaSe/aUtOnOmOuSdAtAbAsEs/aUtOnOmOuSdAtAbAsEnAmE/aUtOnOmOuSdAtAbAsEbAcKuPs/aDbBaCkUpId",
			Expected: &AutonomousDatabaseBackupId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:            "eXaMpLe-rEsOuRcE-GrOuP",
				AutonomousDatabaseName:       "aUtOnOmOuSdAtAbAsEnAmE",
				AutonomousDatabaseBackupName: "aDbBaCkUpId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/oRaClE.DaTaBaSe/aUtOnOmOuSdAtAbAsEs/aUtOnOmOuSdAtAbAsEnAmE/aUtOnOmOuSdAtAbAsEbAcKuPs/aDbBaCkUpId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAutonomousDatabaseBackupIDInsensitively(v.Input)
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

		if actual.AutonomousDatabaseName != v.Expected.AutonomousDatabaseName {
			t.Fatalf("Expected %q but got %q for AutonomousDatabaseName", v.Expected.AutonomousDatabaseName, actual.AutonomousDatabaseName)
		}

		if actual.AutonomousDatabaseBackupName != v.Expected.AutonomousDatabaseBackupName {
			t.Fatalf("Expected %q but got %q for AutonomousDatabaseBackupName", v.Expected.AutonomousDatabaseBackupName, actual.AutonomousDatabaseBackupName)
		}

	}
}

func TestSegmentsForAutonomousDatabaseBackupId(t *testing.T) {
	segments := AutonomousDatabaseBackupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AutonomousDatabaseBackupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
