package benefitutilizationsummariesasync

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ReservationOrderId{}

func TestNewReservationOrderID(t *testing.T) {
	id := NewReservationOrderID("reservationOrderIdValue")

	if id.ReservationOrderId != "reservationOrderIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ReservationOrderId'", id.ReservationOrderId, "reservationOrderIdValue")
	}
}

func TestFormatReservationOrderID(t *testing.T) {
	actual := NewReservationOrderID("reservationOrderIdValue").ID()
	expected := "/providers/Microsoft.Capacity/reservationOrders/reservationOrderIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReservationOrderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReservationOrderId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Capacity",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Capacity/reservationOrders",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Capacity/reservationOrders/reservationOrderIdValue",
			Expected: &ReservationOrderId{
				ReservationOrderId: "reservationOrderIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Capacity/reservationOrders/reservationOrderIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReservationOrderID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ReservationOrderId != v.Expected.ReservationOrderId {
			t.Fatalf("Expected %q but got %q for ReservationOrderId", v.Expected.ReservationOrderId, actual.ReservationOrderId)
		}

	}
}

func TestParseReservationOrderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReservationOrderId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Capacity",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cApAcItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Capacity/reservationOrders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cApAcItY/rEsErVaTiOnOrDeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Capacity/reservationOrders/reservationOrderIdValue",
			Expected: &ReservationOrderId{
				ReservationOrderId: "reservationOrderIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Capacity/reservationOrders/reservationOrderIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cApAcItY/rEsErVaTiOnOrDeRs/rEsErVaTiOnOrDeRiDvAlUe",
			Expected: &ReservationOrderId{
				ReservationOrderId: "rEsErVaTiOnOrDeRiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cApAcItY/rEsErVaTiOnOrDeRs/rEsErVaTiOnOrDeRiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReservationOrderIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ReservationOrderId != v.Expected.ReservationOrderId {
			t.Fatalf("Expected %q but got %q for ReservationOrderId", v.Expected.ReservationOrderId, actual.ReservationOrderId)
		}

	}
}

func TestSegmentsForReservationOrderId(t *testing.T) {
	segments := ReservationOrderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReservationOrderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
