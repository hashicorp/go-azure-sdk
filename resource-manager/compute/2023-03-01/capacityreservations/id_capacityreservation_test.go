package capacityreservations

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = CapacityReservationId{}

func TestNewCapacityReservationID(t *testing.T) {
	id := NewCapacityReservationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "capacityReservationGroupValue", "capacityReservationValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.CapacityReservationGroupName != "capacityReservationGroupValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CapacityReservationGroupName'", id.CapacityReservationGroupName, "capacityReservationGroupValue")
	}

	if id.CapacityReservationName != "capacityReservationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CapacityReservationName'", id.CapacityReservationName, "capacityReservationValue")
	}
}

func TestFormatCapacityReservationID(t *testing.T) {
	actual := NewCapacityReservationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "capacityReservationGroupValue", "capacityReservationValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/capacityReservationGroups/capacityReservationGroupValue/capacityReservations/capacityReservationValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseCapacityReservationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CapacityReservationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/capacityReservationGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/capacityReservationGroups/capacityReservationGroupValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/capacityReservationGroups/capacityReservationGroupValue/capacityReservations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/capacityReservationGroups/capacityReservationGroupValue/capacityReservations/capacityReservationValue",
			Expected: &CapacityReservationId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:            "example-resource-group",
				CapacityReservationGroupName: "capacityReservationGroupValue",
				CapacityReservationName:      "capacityReservationValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/capacityReservationGroups/capacityReservationGroupValue/capacityReservations/capacityReservationValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCapacityReservationID(v.Input)
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

		if actual.CapacityReservationGroupName != v.Expected.CapacityReservationGroupName {
			t.Fatalf("Expected %q but got %q for CapacityReservationGroupName", v.Expected.CapacityReservationGroupName, actual.CapacityReservationGroupName)
		}

		if actual.CapacityReservationName != v.Expected.CapacityReservationName {
			t.Fatalf("Expected %q but got %q for CapacityReservationName", v.Expected.CapacityReservationName, actual.CapacityReservationName)
		}

	}
}

func TestParseCapacityReservationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CapacityReservationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/capacityReservationGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/cApAcItYrEsErVaTiOnGrOuPs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/capacityReservationGroups/capacityReservationGroupValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/cApAcItYrEsErVaTiOnGrOuPs/cApAcItYrEsErVaTiOnGrOuPvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/capacityReservationGroups/capacityReservationGroupValue/capacityReservations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/cApAcItYrEsErVaTiOnGrOuPs/cApAcItYrEsErVaTiOnGrOuPvAlUe/cApAcItYrEsErVaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/capacityReservationGroups/capacityReservationGroupValue/capacityReservations/capacityReservationValue",
			Expected: &CapacityReservationId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:            "example-resource-group",
				CapacityReservationGroupName: "capacityReservationGroupValue",
				CapacityReservationName:      "capacityReservationValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Compute/capacityReservationGroups/capacityReservationGroupValue/capacityReservations/capacityReservationValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/cApAcItYrEsErVaTiOnGrOuPs/cApAcItYrEsErVaTiOnGrOuPvAlUe/cApAcItYrEsErVaTiOnS/cApAcItYrEsErVaTiOnVaLuE",
			Expected: &CapacityReservationId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:            "eXaMpLe-rEsOuRcE-GrOuP",
				CapacityReservationGroupName: "cApAcItYrEsErVaTiOnGrOuPvAlUe",
				CapacityReservationName:      "cApAcItYrEsErVaTiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOmPuTe/cApAcItYrEsErVaTiOnGrOuPs/cApAcItYrEsErVaTiOnGrOuPvAlUe/cApAcItYrEsErVaTiOnS/cApAcItYrEsErVaTiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCapacityReservationIDInsensitively(v.Input)
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

		if actual.CapacityReservationGroupName != v.Expected.CapacityReservationGroupName {
			t.Fatalf("Expected %q but got %q for CapacityReservationGroupName", v.Expected.CapacityReservationGroupName, actual.CapacityReservationGroupName)
		}

		if actual.CapacityReservationName != v.Expected.CapacityReservationName {
			t.Fatalf("Expected %q but got %q for CapacityReservationName", v.Expected.CapacityReservationName, actual.CapacityReservationName)
		}

	}
}

func TestSegmentsForCapacityReservationId(t *testing.T) {
	segments := CapacityReservationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("CapacityReservationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
