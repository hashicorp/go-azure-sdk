package longtermretentionbackups

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &LocationLongTermRetentionServerLongTermRetentionDatabaseId{}

func TestNewLocationLongTermRetentionServerLongTermRetentionDatabaseID(t *testing.T) {
	id := NewLocationLongTermRetentionServerLongTermRetentionDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationName", "longTermRetentionServerName", "longTermRetentionDatabaseName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.LocationName != "locationName" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "locationName")
	}

	if id.LongTermRetentionServerName != "longTermRetentionServerName" {
		t.Fatalf("Expected %q but got %q for Segment 'LongTermRetentionServerName'", id.LongTermRetentionServerName, "longTermRetentionServerName")
	}

	if id.LongTermRetentionDatabaseName != "longTermRetentionDatabaseName" {
		t.Fatalf("Expected %q but got %q for Segment 'LongTermRetentionDatabaseName'", id.LongTermRetentionDatabaseName, "longTermRetentionDatabaseName")
	}
}

func TestFormatLocationLongTermRetentionServerLongTermRetentionDatabaseID(t *testing.T) {
	actual := NewLocationLongTermRetentionServerLongTermRetentionDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationName", "longTermRetentionServerName", "longTermRetentionDatabaseName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationName/longTermRetentionServers/longTermRetentionServerName/longTermRetentionDatabases/longTermRetentionDatabaseName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseLocationLongTermRetentionServerLongTermRetentionDatabaseID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LocationLongTermRetentionServerLongTermRetentionDatabaseId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationName/longTermRetentionServers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationName/longTermRetentionServers/longTermRetentionServerName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationName/longTermRetentionServers/longTermRetentionServerName/longTermRetentionDatabases",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationName/longTermRetentionServers/longTermRetentionServerName/longTermRetentionDatabases/longTermRetentionDatabaseName",
			Expected: &LocationLongTermRetentionServerLongTermRetentionDatabaseId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "example-resource-group",
				LocationName:                  "locationName",
				LongTermRetentionServerName:   "longTermRetentionServerName",
				LongTermRetentionDatabaseName: "longTermRetentionDatabaseName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationName/longTermRetentionServers/longTermRetentionServerName/longTermRetentionDatabases/longTermRetentionDatabaseName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLocationLongTermRetentionServerLongTermRetentionDatabaseID(v.Input)
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

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.LongTermRetentionServerName != v.Expected.LongTermRetentionServerName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionServerName", v.Expected.LongTermRetentionServerName, actual.LongTermRetentionServerName)
		}

		if actual.LongTermRetentionDatabaseName != v.Expected.LongTermRetentionDatabaseName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionDatabaseName", v.Expected.LongTermRetentionDatabaseName, actual.LongTermRetentionDatabaseName)
		}

	}
}

func TestParseLocationLongTermRetentionServerLongTermRetentionDatabaseIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LocationLongTermRetentionServerLongTermRetentionDatabaseId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationName/longTermRetentionServers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNnAmE/lOnGtErMrEtEnTiOnSeRvErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationName/longTermRetentionServers/longTermRetentionServerName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNnAmE/lOnGtErMrEtEnTiOnSeRvErS/lOnGtErMrEtEnTiOnSeRvErNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationName/longTermRetentionServers/longTermRetentionServerName/longTermRetentionDatabases",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNnAmE/lOnGtErMrEtEnTiOnSeRvErS/lOnGtErMrEtEnTiOnSeRvErNaMe/lOnGtErMrEtEnTiOnDaTaBaSeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationName/longTermRetentionServers/longTermRetentionServerName/longTermRetentionDatabases/longTermRetentionDatabaseName",
			Expected: &LocationLongTermRetentionServerLongTermRetentionDatabaseId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "example-resource-group",
				LocationName:                  "locationName",
				LongTermRetentionServerName:   "longTermRetentionServerName",
				LongTermRetentionDatabaseName: "longTermRetentionDatabaseName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationName/longTermRetentionServers/longTermRetentionServerName/longTermRetentionDatabases/longTermRetentionDatabaseName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNnAmE/lOnGtErMrEtEnTiOnSeRvErS/lOnGtErMrEtEnTiOnSeRvErNaMe/lOnGtErMrEtEnTiOnDaTaBaSeS/lOnGtErMrEtEnTiOnDaTaBaSeNaMe",
			Expected: &LocationLongTermRetentionServerLongTermRetentionDatabaseId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "eXaMpLe-rEsOuRcE-GrOuP",
				LocationName:                  "lOcAtIoNnAmE",
				LongTermRetentionServerName:   "lOnGtErMrEtEnTiOnSeRvErNaMe",
				LongTermRetentionDatabaseName: "lOnGtErMrEtEnTiOnDaTaBaSeNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNnAmE/lOnGtErMrEtEnTiOnSeRvErS/lOnGtErMrEtEnTiOnSeRvErNaMe/lOnGtErMrEtEnTiOnDaTaBaSeS/lOnGtErMrEtEnTiOnDaTaBaSeNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLocationLongTermRetentionServerLongTermRetentionDatabaseIDInsensitively(v.Input)
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

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.LongTermRetentionServerName != v.Expected.LongTermRetentionServerName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionServerName", v.Expected.LongTermRetentionServerName, actual.LongTermRetentionServerName)
		}

		if actual.LongTermRetentionDatabaseName != v.Expected.LongTermRetentionDatabaseName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionDatabaseName", v.Expected.LongTermRetentionDatabaseName, actual.LongTermRetentionDatabaseName)
		}

	}
}

func TestSegmentsForLocationLongTermRetentionServerLongTermRetentionDatabaseId(t *testing.T) {
	segments := LocationLongTermRetentionServerLongTermRetentionDatabaseId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("LocationLongTermRetentionServerLongTermRetentionDatabaseId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
