package longtermretentionbackups

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &LongTermRetentionBackupId{}

func TestNewLongTermRetentionBackupID(t *testing.T) {
	id := NewLongTermRetentionBackupID("12345678-1234-9876-4563-123456789012", "locationValue", "longTermRetentionServerValue", "longTermRetentionDatabaseValue", "longTermRetentionBackupValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.LocationName != "locationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "locationValue")
	}

	if id.LongTermRetentionServerName != "longTermRetentionServerValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LongTermRetentionServerName'", id.LongTermRetentionServerName, "longTermRetentionServerValue")
	}

	if id.LongTermRetentionDatabaseName != "longTermRetentionDatabaseValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LongTermRetentionDatabaseName'", id.LongTermRetentionDatabaseName, "longTermRetentionDatabaseValue")
	}

	if id.LongTermRetentionBackupName != "longTermRetentionBackupValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LongTermRetentionBackupName'", id.LongTermRetentionBackupName, "longTermRetentionBackupValue")
	}
}

func TestFormatLongTermRetentionBackupID(t *testing.T) {
	actual := NewLongTermRetentionBackupID("12345678-1234-9876-4563-123456789012", "locationValue", "longTermRetentionServerValue", "longTermRetentionDatabaseValue", "longTermRetentionBackupValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionServers/longTermRetentionServerValue/longTermRetentionDatabases/longTermRetentionDatabaseValue/longTermRetentionBackups/longTermRetentionBackupValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseLongTermRetentionBackupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LongTermRetentionBackupId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionServers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionServers/longTermRetentionServerValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionServers/longTermRetentionServerValue/longTermRetentionDatabases",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionServers/longTermRetentionServerValue/longTermRetentionDatabases/longTermRetentionDatabaseValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionServers/longTermRetentionServerValue/longTermRetentionDatabases/longTermRetentionDatabaseValue/longTermRetentionBackups",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionServers/longTermRetentionServerValue/longTermRetentionDatabases/longTermRetentionDatabaseValue/longTermRetentionBackups/longTermRetentionBackupValue",
			Expected: &LongTermRetentionBackupId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				LocationName:                  "locationValue",
				LongTermRetentionServerName:   "longTermRetentionServerValue",
				LongTermRetentionDatabaseName: "longTermRetentionDatabaseValue",
				LongTermRetentionBackupName:   "longTermRetentionBackupValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionServers/longTermRetentionServerValue/longTermRetentionDatabases/longTermRetentionDatabaseValue/longTermRetentionBackups/longTermRetentionBackupValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLongTermRetentionBackupID(v.Input)
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

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.LongTermRetentionServerName != v.Expected.LongTermRetentionServerName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionServerName", v.Expected.LongTermRetentionServerName, actual.LongTermRetentionServerName)
		}

		if actual.LongTermRetentionDatabaseName != v.Expected.LongTermRetentionDatabaseName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionDatabaseName", v.Expected.LongTermRetentionDatabaseName, actual.LongTermRetentionDatabaseName)
		}

		if actual.LongTermRetentionBackupName != v.Expected.LongTermRetentionBackupName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionBackupName", v.Expected.LongTermRetentionBackupName, actual.LongTermRetentionBackupName)
		}

	}
}

func TestParseLongTermRetentionBackupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LongTermRetentionBackupId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionServers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe/lOnGtErMrEtEnTiOnSeRvErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionServers/longTermRetentionServerValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe/lOnGtErMrEtEnTiOnSeRvErS/lOnGtErMrEtEnTiOnSeRvErVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionServers/longTermRetentionServerValue/longTermRetentionDatabases",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe/lOnGtErMrEtEnTiOnSeRvErS/lOnGtErMrEtEnTiOnSeRvErVaLuE/lOnGtErMrEtEnTiOnDaTaBaSeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionServers/longTermRetentionServerValue/longTermRetentionDatabases/longTermRetentionDatabaseValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe/lOnGtErMrEtEnTiOnSeRvErS/lOnGtErMrEtEnTiOnSeRvErVaLuE/lOnGtErMrEtEnTiOnDaTaBaSeS/lOnGtErMrEtEnTiOnDaTaBaSeVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionServers/longTermRetentionServerValue/longTermRetentionDatabases/longTermRetentionDatabaseValue/longTermRetentionBackups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe/lOnGtErMrEtEnTiOnSeRvErS/lOnGtErMrEtEnTiOnSeRvErVaLuE/lOnGtErMrEtEnTiOnDaTaBaSeS/lOnGtErMrEtEnTiOnDaTaBaSeVaLuE/lOnGtErMrEtEnTiOnBaCkUpS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionServers/longTermRetentionServerValue/longTermRetentionDatabases/longTermRetentionDatabaseValue/longTermRetentionBackups/longTermRetentionBackupValue",
			Expected: &LongTermRetentionBackupId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				LocationName:                  "locationValue",
				LongTermRetentionServerName:   "longTermRetentionServerValue",
				LongTermRetentionDatabaseName: "longTermRetentionDatabaseValue",
				LongTermRetentionBackupName:   "longTermRetentionBackupValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionServers/longTermRetentionServerValue/longTermRetentionDatabases/longTermRetentionDatabaseValue/longTermRetentionBackups/longTermRetentionBackupValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe/lOnGtErMrEtEnTiOnSeRvErS/lOnGtErMrEtEnTiOnSeRvErVaLuE/lOnGtErMrEtEnTiOnDaTaBaSeS/lOnGtErMrEtEnTiOnDaTaBaSeVaLuE/lOnGtErMrEtEnTiOnBaCkUpS/lOnGtErMrEtEnTiOnBaCkUpVaLuE",
			Expected: &LongTermRetentionBackupId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				LocationName:                  "lOcAtIoNvAlUe",
				LongTermRetentionServerName:   "lOnGtErMrEtEnTiOnSeRvErVaLuE",
				LongTermRetentionDatabaseName: "lOnGtErMrEtEnTiOnDaTaBaSeVaLuE",
				LongTermRetentionBackupName:   "lOnGtErMrEtEnTiOnBaCkUpVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe/lOnGtErMrEtEnTiOnSeRvErS/lOnGtErMrEtEnTiOnSeRvErVaLuE/lOnGtErMrEtEnTiOnDaTaBaSeS/lOnGtErMrEtEnTiOnDaTaBaSeVaLuE/lOnGtErMrEtEnTiOnBaCkUpS/lOnGtErMrEtEnTiOnBaCkUpVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLongTermRetentionBackupIDInsensitively(v.Input)
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

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.LongTermRetentionServerName != v.Expected.LongTermRetentionServerName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionServerName", v.Expected.LongTermRetentionServerName, actual.LongTermRetentionServerName)
		}

		if actual.LongTermRetentionDatabaseName != v.Expected.LongTermRetentionDatabaseName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionDatabaseName", v.Expected.LongTermRetentionDatabaseName, actual.LongTermRetentionDatabaseName)
		}

		if actual.LongTermRetentionBackupName != v.Expected.LongTermRetentionBackupName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionBackupName", v.Expected.LongTermRetentionBackupName, actual.LongTermRetentionBackupName)
		}

	}
}

func TestSegmentsForLongTermRetentionBackupId(t *testing.T) {
	segments := LongTermRetentionBackupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("LongTermRetentionBackupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
