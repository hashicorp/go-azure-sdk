package longtermretentionmanagedinstancebackups

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = LongTermRetentionManagedInstanceBackupId{}

func TestNewLongTermRetentionManagedInstanceBackupID(t *testing.T) {
	id := NewLongTermRetentionManagedInstanceBackupID("12345678-1234-9876-4563-123456789012", "locationValue", "longTermRetentionManagedInstanceValue", "longTermRetentionDatabaseValue", "longTermRetentionManagedInstanceBackupValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.LocationName != "locationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "locationValue")
	}

	if id.LongTermRetentionManagedInstanceName != "longTermRetentionManagedInstanceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LongTermRetentionManagedInstanceName'", id.LongTermRetentionManagedInstanceName, "longTermRetentionManagedInstanceValue")
	}

	if id.LongTermRetentionDatabaseName != "longTermRetentionDatabaseValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LongTermRetentionDatabaseName'", id.LongTermRetentionDatabaseName, "longTermRetentionDatabaseValue")
	}

	if id.LongTermRetentionManagedInstanceBackupName != "longTermRetentionManagedInstanceBackupValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LongTermRetentionManagedInstanceBackupName'", id.LongTermRetentionManagedInstanceBackupName, "longTermRetentionManagedInstanceBackupValue")
	}
}

func TestFormatLongTermRetentionManagedInstanceBackupID(t *testing.T) {
	actual := NewLongTermRetentionManagedInstanceBackupID("12345678-1234-9876-4563-123456789012", "locationValue", "longTermRetentionManagedInstanceValue", "longTermRetentionDatabaseValue", "longTermRetentionManagedInstanceBackupValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue/longTermRetentionDatabases/longTermRetentionDatabaseValue/longTermRetentionManagedInstanceBackups/longTermRetentionManagedInstanceBackupValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseLongTermRetentionManagedInstanceBackupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LongTermRetentionManagedInstanceBackupId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue/longTermRetentionDatabases",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue/longTermRetentionDatabases/longTermRetentionDatabaseValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue/longTermRetentionDatabases/longTermRetentionDatabaseValue/longTermRetentionManagedInstanceBackups",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue/longTermRetentionDatabases/longTermRetentionDatabaseValue/longTermRetentionManagedInstanceBackups/longTermRetentionManagedInstanceBackupValue",
			Expected: &LongTermRetentionManagedInstanceBackupId{
				SubscriptionId:                             "12345678-1234-9876-4563-123456789012",
				LocationName:                               "locationValue",
				LongTermRetentionManagedInstanceName:       "longTermRetentionManagedInstanceValue",
				LongTermRetentionDatabaseName:              "longTermRetentionDatabaseValue",
				LongTermRetentionManagedInstanceBackupName: "longTermRetentionManagedInstanceBackupValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue/longTermRetentionDatabases/longTermRetentionDatabaseValue/longTermRetentionManagedInstanceBackups/longTermRetentionManagedInstanceBackupValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLongTermRetentionManagedInstanceBackupID(v.Input)
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

		if actual.LongTermRetentionManagedInstanceName != v.Expected.LongTermRetentionManagedInstanceName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionManagedInstanceName", v.Expected.LongTermRetentionManagedInstanceName, actual.LongTermRetentionManagedInstanceName)
		}

		if actual.LongTermRetentionDatabaseName != v.Expected.LongTermRetentionDatabaseName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionDatabaseName", v.Expected.LongTermRetentionDatabaseName, actual.LongTermRetentionDatabaseName)
		}

		if actual.LongTermRetentionManagedInstanceBackupName != v.Expected.LongTermRetentionManagedInstanceBackupName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionManagedInstanceBackupName", v.Expected.LongTermRetentionManagedInstanceBackupName, actual.LongTermRetentionManagedInstanceBackupName)
		}

	}
}

func TestParseLongTermRetentionManagedInstanceBackupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LongTermRetentionManagedInstanceBackupId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEs/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue/longTermRetentionDatabases",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEs/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEvAlUe/lOnGtErMrEtEnTiOnDaTaBaSeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue/longTermRetentionDatabases/longTermRetentionDatabaseValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEs/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEvAlUe/lOnGtErMrEtEnTiOnDaTaBaSeS/lOnGtErMrEtEnTiOnDaTaBaSeVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue/longTermRetentionDatabases/longTermRetentionDatabaseValue/longTermRetentionManagedInstanceBackups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEs/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEvAlUe/lOnGtErMrEtEnTiOnDaTaBaSeS/lOnGtErMrEtEnTiOnDaTaBaSeVaLuE/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEbAcKuPs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue/longTermRetentionDatabases/longTermRetentionDatabaseValue/longTermRetentionManagedInstanceBackups/longTermRetentionManagedInstanceBackupValue",
			Expected: &LongTermRetentionManagedInstanceBackupId{
				SubscriptionId:                             "12345678-1234-9876-4563-123456789012",
				LocationName:                               "locationValue",
				LongTermRetentionManagedInstanceName:       "longTermRetentionManagedInstanceValue",
				LongTermRetentionDatabaseName:              "longTermRetentionDatabaseValue",
				LongTermRetentionManagedInstanceBackupName: "longTermRetentionManagedInstanceBackupValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue/longTermRetentionDatabases/longTermRetentionDatabaseValue/longTermRetentionManagedInstanceBackups/longTermRetentionManagedInstanceBackupValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEs/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEvAlUe/lOnGtErMrEtEnTiOnDaTaBaSeS/lOnGtErMrEtEnTiOnDaTaBaSeVaLuE/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEbAcKuPs/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEbAcKuPvAlUe",
			Expected: &LongTermRetentionManagedInstanceBackupId{
				SubscriptionId:                             "12345678-1234-9876-4563-123456789012",
				LocationName:                               "lOcAtIoNvAlUe",
				LongTermRetentionManagedInstanceName:       "lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEvAlUe",
				LongTermRetentionDatabaseName:              "lOnGtErMrEtEnTiOnDaTaBaSeVaLuE",
				LongTermRetentionManagedInstanceBackupName: "lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEbAcKuPvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEs/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEvAlUe/lOnGtErMrEtEnTiOnDaTaBaSeS/lOnGtErMrEtEnTiOnDaTaBaSeVaLuE/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEbAcKuPs/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEbAcKuPvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLongTermRetentionManagedInstanceBackupIDInsensitively(v.Input)
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

		if actual.LongTermRetentionManagedInstanceName != v.Expected.LongTermRetentionManagedInstanceName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionManagedInstanceName", v.Expected.LongTermRetentionManagedInstanceName, actual.LongTermRetentionManagedInstanceName)
		}

		if actual.LongTermRetentionDatabaseName != v.Expected.LongTermRetentionDatabaseName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionDatabaseName", v.Expected.LongTermRetentionDatabaseName, actual.LongTermRetentionDatabaseName)
		}

		if actual.LongTermRetentionManagedInstanceBackupName != v.Expected.LongTermRetentionManagedInstanceBackupName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionManagedInstanceBackupName", v.Expected.LongTermRetentionManagedInstanceBackupName, actual.LongTermRetentionManagedInstanceBackupName)
		}

	}
}

func TestSegmentsForLongTermRetentionManagedInstanceBackupId(t *testing.T) {
	segments := LongTermRetentionManagedInstanceBackupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("LongTermRetentionManagedInstanceBackupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
