package longtermretentionmanagedinstancebackups

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &LongTermRetentionManagedInstanceLongTermRetentionDatabaseId{}

func TestNewLongTermRetentionManagedInstanceLongTermRetentionDatabaseID(t *testing.T) {
	id := NewLongTermRetentionManagedInstanceLongTermRetentionDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue", "longTermRetentionManagedInstanceValue", "longTermRetentionDatabaseValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
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
}

func TestFormatLongTermRetentionManagedInstanceLongTermRetentionDatabaseID(t *testing.T) {
	actual := NewLongTermRetentionManagedInstanceLongTermRetentionDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue", "longTermRetentionManagedInstanceValue", "longTermRetentionDatabaseValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue/longTermRetentionDatabases/longTermRetentionDatabaseValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseLongTermRetentionManagedInstanceLongTermRetentionDatabaseID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LongTermRetentionManagedInstanceLongTermRetentionDatabaseId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue/longTermRetentionDatabases",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue/longTermRetentionDatabases/longTermRetentionDatabaseValue",
			Expected: &LongTermRetentionManagedInstanceLongTermRetentionDatabaseId{
				SubscriptionId:                       "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                    "example-resource-group",
				LocationName:                         "locationValue",
				LongTermRetentionManagedInstanceName: "longTermRetentionManagedInstanceValue",
				LongTermRetentionDatabaseName:        "longTermRetentionDatabaseValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue/longTermRetentionDatabases/longTermRetentionDatabaseValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLongTermRetentionManagedInstanceLongTermRetentionDatabaseID(v.Input)
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

		if actual.LongTermRetentionManagedInstanceName != v.Expected.LongTermRetentionManagedInstanceName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionManagedInstanceName", v.Expected.LongTermRetentionManagedInstanceName, actual.LongTermRetentionManagedInstanceName)
		}

		if actual.LongTermRetentionDatabaseName != v.Expected.LongTermRetentionDatabaseName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionDatabaseName", v.Expected.LongTermRetentionDatabaseName, actual.LongTermRetentionDatabaseName)
		}

	}
}

func TestParseLongTermRetentionManagedInstanceLongTermRetentionDatabaseIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LongTermRetentionManagedInstanceLongTermRetentionDatabaseId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEs/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue/longTermRetentionDatabases",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEs/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEvAlUe/lOnGtErMrEtEnTiOnDaTaBaSeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue/longTermRetentionDatabases/longTermRetentionDatabaseValue",
			Expected: &LongTermRetentionManagedInstanceLongTermRetentionDatabaseId{
				SubscriptionId:                       "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                    "example-resource-group",
				LocationName:                         "locationValue",
				LongTermRetentionManagedInstanceName: "longTermRetentionManagedInstanceValue",
				LongTermRetentionDatabaseName:        "longTermRetentionDatabaseValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/locations/locationValue/longTermRetentionManagedInstances/longTermRetentionManagedInstanceValue/longTermRetentionDatabases/longTermRetentionDatabaseValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEs/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEvAlUe/lOnGtErMrEtEnTiOnDaTaBaSeS/lOnGtErMrEtEnTiOnDaTaBaSeVaLuE",
			Expected: &LongTermRetentionManagedInstanceLongTermRetentionDatabaseId{
				SubscriptionId:                       "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                    "eXaMpLe-rEsOuRcE-GrOuP",
				LocationName:                         "lOcAtIoNvAlUe",
				LongTermRetentionManagedInstanceName: "lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEvAlUe",
				LongTermRetentionDatabaseName:        "lOnGtErMrEtEnTiOnDaTaBaSeVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNvAlUe/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEs/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEvAlUe/lOnGtErMrEtEnTiOnDaTaBaSeS/lOnGtErMrEtEnTiOnDaTaBaSeVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLongTermRetentionManagedInstanceLongTermRetentionDatabaseIDInsensitively(v.Input)
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

		if actual.LongTermRetentionManagedInstanceName != v.Expected.LongTermRetentionManagedInstanceName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionManagedInstanceName", v.Expected.LongTermRetentionManagedInstanceName, actual.LongTermRetentionManagedInstanceName)
		}

		if actual.LongTermRetentionDatabaseName != v.Expected.LongTermRetentionDatabaseName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionDatabaseName", v.Expected.LongTermRetentionDatabaseName, actual.LongTermRetentionDatabaseName)
		}

	}
}

func TestSegmentsForLongTermRetentionManagedInstanceLongTermRetentionDatabaseId(t *testing.T) {
	segments := LongTermRetentionManagedInstanceLongTermRetentionDatabaseId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("LongTermRetentionManagedInstanceLongTermRetentionDatabaseId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
