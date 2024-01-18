package managedhsms

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeletedManagedHSMId{}

func TestNewDeletedManagedHSMID(t *testing.T) {
	id := NewDeletedManagedHSMID("12345678-1234-9876-4563-123456789012", "locationValue", "deletedManagedHSMValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.LocationName != "locationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "locationValue")
	}

	if id.DeletedManagedHSMName != "deletedManagedHSMValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeletedManagedHSMName'", id.DeletedManagedHSMName, "deletedManagedHSMValue")
	}
}

func TestFormatDeletedManagedHSMID(t *testing.T) {
	actual := NewDeletedManagedHSMID("12345678-1234-9876-4563-123456789012", "locationValue", "deletedManagedHSMValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.KeyVault/locations/locationValue/deletedManagedHSMs/deletedManagedHSMValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeletedManagedHSMID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeletedManagedHSMId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.KeyVault",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.KeyVault/locations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.KeyVault/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.KeyVault/locations/locationValue/deletedManagedHSMs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.KeyVault/locations/locationValue/deletedManagedHSMs/deletedManagedHSMValue",
			Expected: &DeletedManagedHSMId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				LocationName:          "locationValue",
				DeletedManagedHSMName: "deletedManagedHSMValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.KeyVault/locations/locationValue/deletedManagedHSMs/deletedManagedHSMValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeletedManagedHSMID(v.Input)
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

		if actual.DeletedManagedHSMName != v.Expected.DeletedManagedHSMName {
			t.Fatalf("Expected %q but got %q for DeletedManagedHSMName", v.Expected.DeletedManagedHSMName, actual.DeletedManagedHSMName)
		}

	}
}

func TestParseDeletedManagedHSMIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeletedManagedHSMId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.KeyVault",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.kEyVaUlT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.KeyVault/locations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.kEyVaUlT/lOcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.KeyVault/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.kEyVaUlT/lOcAtIoNs/lOcAtIoNvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.KeyVault/locations/locationValue/deletedManagedHSMs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.kEyVaUlT/lOcAtIoNs/lOcAtIoNvAlUe/dElEtEdMaNaGeDhSmS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.KeyVault/locations/locationValue/deletedManagedHSMs/deletedManagedHSMValue",
			Expected: &DeletedManagedHSMId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				LocationName:          "locationValue",
				DeletedManagedHSMName: "deletedManagedHSMValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.KeyVault/locations/locationValue/deletedManagedHSMs/deletedManagedHSMValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.kEyVaUlT/lOcAtIoNs/lOcAtIoNvAlUe/dElEtEdMaNaGeDhSmS/dElEtEdMaNaGeDhSmVaLuE",
			Expected: &DeletedManagedHSMId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				LocationName:          "lOcAtIoNvAlUe",
				DeletedManagedHSMName: "dElEtEdMaNaGeDhSmVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.kEyVaUlT/lOcAtIoNs/lOcAtIoNvAlUe/dElEtEdMaNaGeDhSmS/dElEtEdMaNaGeDhSmVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeletedManagedHSMIDInsensitively(v.Input)
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

		if actual.DeletedManagedHSMName != v.Expected.DeletedManagedHSMName {
			t.Fatalf("Expected %q but got %q for DeletedManagedHSMName", v.Expected.DeletedManagedHSMName, actual.DeletedManagedHSMName)
		}

	}
}

func TestSegmentsForDeletedManagedHSMId(t *testing.T) {
	segments := DeletedManagedHSMId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeletedManagedHSMId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
