package longtermretentionmanagedinstancebackups

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &LongTermRetentionDatabaseId{}

func TestNewLongTermRetentionDatabaseID(t *testing.T) {
	id := NewLongTermRetentionDatabaseID("12345678-1234-9876-4563-123456789012", "locationName", "managedInstanceName", "databaseName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.LocationName != "locationName" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "locationName")
	}

	if id.LongTermRetentionManagedInstanceName != "managedInstanceName" {
		t.Fatalf("Expected %q but got %q for Segment 'LongTermRetentionManagedInstanceName'", id.LongTermRetentionManagedInstanceName, "managedInstanceName")
	}

	if id.LongTermRetentionDatabaseName != "databaseName" {
		t.Fatalf("Expected %q but got %q for Segment 'LongTermRetentionDatabaseName'", id.LongTermRetentionDatabaseName, "databaseName")
	}
}

func TestFormatLongTermRetentionDatabaseID(t *testing.T) {
	actual := NewLongTermRetentionDatabaseID("12345678-1234-9876-4563-123456789012", "locationName", "managedInstanceName", "databaseName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName/longTermRetentionManagedInstances/managedInstanceName/longTermRetentionDatabases/databaseName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseLongTermRetentionDatabaseID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LongTermRetentionDatabaseId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName/longTermRetentionManagedInstances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName/longTermRetentionManagedInstances/managedInstanceName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName/longTermRetentionManagedInstances/managedInstanceName/longTermRetentionDatabases",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName/longTermRetentionManagedInstances/managedInstanceName/longTermRetentionDatabases/databaseName",
			Expected: &LongTermRetentionDatabaseId{
				SubscriptionId:                       "12345678-1234-9876-4563-123456789012",
				LocationName:                         "locationName",
				LongTermRetentionManagedInstanceName: "managedInstanceName",
				LongTermRetentionDatabaseName:        "databaseName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName/longTermRetentionManagedInstances/managedInstanceName/longTermRetentionDatabases/databaseName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLongTermRetentionDatabaseID(v.Input)
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

	}
}

func TestParseLongTermRetentionDatabaseIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LongTermRetentionDatabaseId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName/longTermRetentionManagedInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNnAmE/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName/longTermRetentionManagedInstances/managedInstanceName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNnAmE/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEs/mAnAgEdInStAnCeNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName/longTermRetentionManagedInstances/managedInstanceName/longTermRetentionDatabases",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNnAmE/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEs/mAnAgEdInStAnCeNaMe/lOnGtErMrEtEnTiOnDaTaBaSeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName/longTermRetentionManagedInstances/managedInstanceName/longTermRetentionDatabases/databaseName",
			Expected: &LongTermRetentionDatabaseId{
				SubscriptionId:                       "12345678-1234-9876-4563-123456789012",
				LocationName:                         "locationName",
				LongTermRetentionManagedInstanceName: "managedInstanceName",
				LongTermRetentionDatabaseName:        "databaseName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName/longTermRetentionManagedInstances/managedInstanceName/longTermRetentionDatabases/databaseName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNnAmE/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEs/mAnAgEdInStAnCeNaMe/lOnGtErMrEtEnTiOnDaTaBaSeS/dAtAbAsEnAmE",
			Expected: &LongTermRetentionDatabaseId{
				SubscriptionId:                       "12345678-1234-9876-4563-123456789012",
				LocationName:                         "lOcAtIoNnAmE",
				LongTermRetentionManagedInstanceName: "mAnAgEdInStAnCeNaMe",
				LongTermRetentionDatabaseName:        "dAtAbAsEnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNnAmE/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEs/mAnAgEdInStAnCeNaMe/lOnGtErMrEtEnTiOnDaTaBaSeS/dAtAbAsEnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLongTermRetentionDatabaseIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForLongTermRetentionDatabaseId(t *testing.T) {
	segments := LongTermRetentionDatabaseId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("LongTermRetentionDatabaseId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
