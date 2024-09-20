package reservationdetails

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReservationId{}

func TestNewReservationID(t *testing.T) {
	id := NewReservationID("reservationOrderId", "reservationId")

	if id.ReservationOrderId != "reservationOrderId" {
		t.Fatalf("Expected %q but got %q for Segment 'ReservationOrderId'", id.ReservationOrderId, "reservationOrderId")
	}

	if id.ReservationId != "reservationId" {
		t.Fatalf("Expected %q but got %q for Segment 'ReservationId'", id.ReservationId, "reservationId")
	}
}

func TestFormatReservationID(t *testing.T) {
	actual := NewReservationID("reservationOrderId", "reservationId").ID()
	expected := "/providers/Microsoft.Capacity/reservationOrders/reservationOrderId/reservations/reservationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReservationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReservationId
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
			// Incomplete URI
			Input: "/providers/Microsoft.Capacity/reservationOrders/reservationOrderId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Capacity/reservationOrders/reservationOrderId/reservations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Capacity/reservationOrders/reservationOrderId/reservations/reservationId",
			Expected: &ReservationId{
				ReservationOrderId: "reservationOrderId",
				ReservationId:      "reservationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Capacity/reservationOrders/reservationOrderId/reservations/reservationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReservationID(v.Input)
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

		if actual.ReservationId != v.Expected.ReservationId {
			t.Fatalf("Expected %q but got %q for ReservationId", v.Expected.ReservationId, actual.ReservationId)
		}

	}
}

func TestParseReservationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReservationId
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
			// Incomplete URI
			Input: "/providers/Microsoft.Capacity/reservationOrders/reservationOrderId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cApAcItY/rEsErVaTiOnOrDeRs/rEsErVaTiOnOrDeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Capacity/reservationOrders/reservationOrderId/reservations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cApAcItY/rEsErVaTiOnOrDeRs/rEsErVaTiOnOrDeRiD/rEsErVaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Capacity/reservationOrders/reservationOrderId/reservations/reservationId",
			Expected: &ReservationId{
				ReservationOrderId: "reservationOrderId",
				ReservationId:      "reservationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Capacity/reservationOrders/reservationOrderId/reservations/reservationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cApAcItY/rEsErVaTiOnOrDeRs/rEsErVaTiOnOrDeRiD/rEsErVaTiOnS/rEsErVaTiOnId",
			Expected: &ReservationId{
				ReservationOrderId: "rEsErVaTiOnOrDeRiD",
				ReservationId:      "rEsErVaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cApAcItY/rEsErVaTiOnOrDeRs/rEsErVaTiOnOrDeRiD/rEsErVaTiOnS/rEsErVaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReservationIDInsensitively(v.Input)
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

		if actual.ReservationId != v.Expected.ReservationId {
			t.Fatalf("Expected %q but got %q for ReservationId", v.Expected.ReservationId, actual.ReservationId)
		}

	}
}

func TestSegmentsForReservationId(t *testing.T) {
	segments := ReservationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReservationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
