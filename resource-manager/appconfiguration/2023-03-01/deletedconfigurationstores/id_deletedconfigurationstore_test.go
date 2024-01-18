package deletedconfigurationstores

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeletedConfigurationStoreId{}

func TestNewDeletedConfigurationStoreID(t *testing.T) {
	id := NewDeletedConfigurationStoreID("12345678-1234-9876-4563-123456789012", "locationValue", "deletedConfigurationStoreValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.LocationName != "locationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "locationValue")
	}

	if id.DeletedConfigurationStoreName != "deletedConfigurationStoreValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeletedConfigurationStoreName'", id.DeletedConfigurationStoreName, "deletedConfigurationStoreValue")
	}
}

func TestFormatDeletedConfigurationStoreID(t *testing.T) {
	actual := NewDeletedConfigurationStoreID("12345678-1234-9876-4563-123456789012", "locationValue", "deletedConfigurationStoreValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.AppConfiguration/locations/locationValue/deletedConfigurationStores/deletedConfigurationStoreValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeletedConfigurationStoreID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeletedConfigurationStoreId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.AppConfiguration",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.AppConfiguration/locations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.AppConfiguration/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.AppConfiguration/locations/locationValue/deletedConfigurationStores",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.AppConfiguration/locations/locationValue/deletedConfigurationStores/deletedConfigurationStoreValue",
			Expected: &DeletedConfigurationStoreId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				LocationName:                  "locationValue",
				DeletedConfigurationStoreName: "deletedConfigurationStoreValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.AppConfiguration/locations/locationValue/deletedConfigurationStores/deletedConfigurationStoreValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeletedConfigurationStoreID(v.Input)
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

		if actual.DeletedConfigurationStoreName != v.Expected.DeletedConfigurationStoreName {
			t.Fatalf("Expected %q but got %q for DeletedConfigurationStoreName", v.Expected.DeletedConfigurationStoreName, actual.DeletedConfigurationStoreName)
		}

	}
}

func TestParseDeletedConfigurationStoreIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeletedConfigurationStoreId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.AppConfiguration",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aPpCoNfIgUrAtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.AppConfiguration/locations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aPpCoNfIgUrAtIoN/lOcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.AppConfiguration/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aPpCoNfIgUrAtIoN/lOcAtIoNs/lOcAtIoNvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.AppConfiguration/locations/locationValue/deletedConfigurationStores",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aPpCoNfIgUrAtIoN/lOcAtIoNs/lOcAtIoNvAlUe/dElEtEdCoNfIgUrAtIoNsToReS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.AppConfiguration/locations/locationValue/deletedConfigurationStores/deletedConfigurationStoreValue",
			Expected: &DeletedConfigurationStoreId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				LocationName:                  "locationValue",
				DeletedConfigurationStoreName: "deletedConfigurationStoreValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.AppConfiguration/locations/locationValue/deletedConfigurationStores/deletedConfigurationStoreValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aPpCoNfIgUrAtIoN/lOcAtIoNs/lOcAtIoNvAlUe/dElEtEdCoNfIgUrAtIoNsToReS/dElEtEdCoNfIgUrAtIoNsToReVaLuE",
			Expected: &DeletedConfigurationStoreId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				LocationName:                  "lOcAtIoNvAlUe",
				DeletedConfigurationStoreName: "dElEtEdCoNfIgUrAtIoNsToReVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aPpCoNfIgUrAtIoN/lOcAtIoNs/lOcAtIoNvAlUe/dElEtEdCoNfIgUrAtIoNsToReS/dElEtEdCoNfIgUrAtIoNsToReVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeletedConfigurationStoreIDInsensitively(v.Input)
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

		if actual.DeletedConfigurationStoreName != v.Expected.DeletedConfigurationStoreName {
			t.Fatalf("Expected %q but got %q for DeletedConfigurationStoreName", v.Expected.DeletedConfigurationStoreName, actual.DeletedConfigurationStoreName)
		}

	}
}

func TestSegmentsForDeletedConfigurationStoreId(t *testing.T) {
	segments := DeletedConfigurationStoreId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeletedConfigurationStoreId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
