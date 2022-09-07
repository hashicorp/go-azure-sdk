package operationsstatus

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = OperationsStatusId{}

func TestNewOperationsStatusID(t *testing.T) {
	id := NewOperationsStatusID("12345678-1234-9876-4563-123456789012", "locationValue", "operationIdValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.Location != "locationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'Location'", id.Location, "locationValue")
	}

	if id.OperationId != "operationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OperationId'", id.OperationId, "operationIdValue")
	}
}

func TestFormatOperationsStatusID(t *testing.T) {
	actual := NewOperationsStatusID("12345678-1234-9876-4563-123456789012", "locationValue", "operationIdValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Cache/locations/locationValue/operationsStatus/operationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseOperationsStatusID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *OperationsStatusId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Cache",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Cache/locations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Cache/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Cache/locations/locationValue/operationsStatus",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Cache/locations/locationValue/operationsStatus/operationIdValue",
			Expected: &OperationsStatusId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				Location:       "locationValue",
				OperationId:    "operationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Cache/locations/locationValue/operationsStatus/operationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseOperationsStatusID(v.Input)
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

		if actual.Location != v.Expected.Location {
			t.Fatalf("Expected %q but got %q for Location", v.Expected.Location, actual.Location)
		}

		if actual.OperationId != v.Expected.OperationId {
			t.Fatalf("Expected %q but got %q for OperationId", v.Expected.OperationId, actual.OperationId)
		}

	}
}

func TestParseOperationsStatusIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *OperationsStatusId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Cache",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cAcHe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Cache/locations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cAcHe/lOcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Cache/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cAcHe/lOcAtIoNs/lOcAtIoNvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Cache/locations/locationValue/operationsStatus",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cAcHe/lOcAtIoNs/lOcAtIoNvAlUe/oPeRaTiOnSsTaTuS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Cache/locations/locationValue/operationsStatus/operationIdValue",
			Expected: &OperationsStatusId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				Location:       "locationValue",
				OperationId:    "operationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Cache/locations/locationValue/operationsStatus/operationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cAcHe/lOcAtIoNs/lOcAtIoNvAlUe/oPeRaTiOnSsTaTuS/oPeRaTiOnIdVaLuE",
			Expected: &OperationsStatusId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				Location:       "lOcAtIoNvAlUe",
				OperationId:    "oPeRaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cAcHe/lOcAtIoNs/lOcAtIoNvAlUe/oPeRaTiOnSsTaTuS/oPeRaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseOperationsStatusIDInsensitively(v.Input)
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

		if actual.Location != v.Expected.Location {
			t.Fatalf("Expected %q but got %q for Location", v.Expected.Location, actual.Location)
		}

		if actual.OperationId != v.Expected.OperationId {
			t.Fatalf("Expected %q but got %q for OperationId", v.Expected.OperationId, actual.OperationId)
		}

	}
}

func TestSegmentsForOperationsStatusId(t *testing.T) {
	segments := OperationsStatusId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("OperationsStatusId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
