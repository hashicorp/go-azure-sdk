package reservationdetails

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ReservationId{}

func TestNewReservationID(t *testing.T) {
	id := NewReservationID("reservationOrderIdValue", "reservationIdValue")

	if id.ReservationOrderId != "reservationOrderIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ReservationOrderId'", id.ReservationOrderId, "reservationOrderIdValue")
	}

	if id.ReservationId != "reservationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ReservationId'", id.ReservationId, "reservationIdValue")
	}
}

func TestFormatReservationID(t *testing.T) {
	actual := NewReservationID("reservationOrderIdValue", "reservationIdValue").ID()
	expected := "/providers/Microsoft.Capacity/reservationOrders/reservationOrderIdValue/reservations/reservationIdValue"
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
			Input: "/providers/Microsoft.Capacity/reservationOrders/reservationOrderIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Capacity/reservationOrders/reservationOrderIdValue/reservations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Capacity/reservationOrders/reservationOrderIdValue/reservations/reservationIdValue",
			Expected: &ReservationId{
				ReservationOrderId: "reservationOrderIdValue",
				ReservationId:      "reservationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Capacity/reservationOrders/reservationOrderIdValue/reservations/reservationIdValue/extra",
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
			Input: "/providers/Microsoft.Capacity/reservationOrders/reservationOrderIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cApAcItY/rEsErVaTiOnOrDeRs/rEsErVaTiOnOrDeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Capacity/reservationOrders/reservationOrderIdValue/reservations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cApAcItY/rEsErVaTiOnOrDeRs/rEsErVaTiOnOrDeRiDvAlUe/rEsErVaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Capacity/reservationOrders/reservationOrderIdValue/reservations/reservationIdValue",
			Expected: &ReservationId{
				ReservationOrderId: "reservationOrderIdValue",
				ReservationId:      "reservationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Capacity/reservationOrders/reservationOrderIdValue/reservations/reservationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cApAcItY/rEsErVaTiOnOrDeRs/rEsErVaTiOnOrDeRiDvAlUe/rEsErVaTiOnS/rEsErVaTiOnIdVaLuE",
			Expected: &ReservationId{
				ReservationOrderId: "rEsErVaTiOnOrDeRiDvAlUe",
				ReservationId:      "rEsErVaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cApAcItY/rEsErVaTiOnOrDeRs/rEsErVaTiOnOrDeRiDvAlUe/rEsErVaTiOnS/rEsErVaTiOnIdVaLuE/extra",
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
