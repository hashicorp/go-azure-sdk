package extensions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ExtensionTypeId{}

func TestNewExtensionTypeID(t *testing.T) {
	id := NewExtensionTypeID("12345678-1234-9876-4563-123456789012", "locationValue", "publisherValue", "extensionTypeValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.Location != "locationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'Location'", id.Location, "locationValue")
	}

	if id.Publisher != "publisherValue" {
		t.Fatalf("Expected %q but got %q for Segment 'Publisher'", id.Publisher, "publisherValue")
	}

	if id.ExtensionType != "extensionTypeValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionType'", id.ExtensionType, "extensionTypeValue")
	}
}

func TestFormatExtensionTypeID(t *testing.T) {
	actual := NewExtensionTypeID("12345678-1234-9876-4563-123456789012", "locationValue", "publisherValue", "extensionTypeValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HybridCompute/locations/locationValue/publishers/publisherValue/extensionTypes/extensionTypeValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseExtensionTypeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ExtensionTypeId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HybridCompute",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HybridCompute/locations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HybridCompute/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HybridCompute/locations/locationValue/publishers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HybridCompute/locations/locationValue/publishers/publisherValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HybridCompute/locations/locationValue/publishers/publisherValue/extensionTypes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HybridCompute/locations/locationValue/publishers/publisherValue/extensionTypes/extensionTypeValue",
			Expected: &ExtensionTypeId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				Location:       "locationValue",
				Publisher:      "publisherValue",
				ExtensionType:  "extensionTypeValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HybridCompute/locations/locationValue/publishers/publisherValue/extensionTypes/extensionTypeValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseExtensionTypeID(v.Input)
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

		if actual.Publisher != v.Expected.Publisher {
			t.Fatalf("Expected %q but got %q for Publisher", v.Expected.Publisher, actual.Publisher)
		}

		if actual.ExtensionType != v.Expected.ExtensionType {
			t.Fatalf("Expected %q but got %q for ExtensionType", v.Expected.ExtensionType, actual.ExtensionType)
		}

	}
}

func TestParseExtensionTypeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ExtensionTypeId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HybridCompute",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HybridCompute/locations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/lOcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HybridCompute/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HybridCompute/locations/locationValue/publishers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/pUbLiShErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HybridCompute/locations/locationValue/publishers/publisherValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/pUbLiShErS/pUbLiShErVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HybridCompute/locations/locationValue/publishers/publisherValue/extensionTypes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/pUbLiShErS/pUbLiShErVaLuE/eXtEnSiOnTyPeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HybridCompute/locations/locationValue/publishers/publisherValue/extensionTypes/extensionTypeValue",
			Expected: &ExtensionTypeId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				Location:       "locationValue",
				Publisher:      "publisherValue",
				ExtensionType:  "extensionTypeValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.HybridCompute/locations/locationValue/publishers/publisherValue/extensionTypes/extensionTypeValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/pUbLiShErS/pUbLiShErVaLuE/eXtEnSiOnTyPeS/eXtEnSiOnTyPeVaLuE",
			Expected: &ExtensionTypeId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				Location:       "lOcAtIoNvAlUe",
				Publisher:      "pUbLiShErVaLuE",
				ExtensionType:  "eXtEnSiOnTyPeVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/pUbLiShErS/pUbLiShErVaLuE/eXtEnSiOnTyPeS/eXtEnSiOnTyPeVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseExtensionTypeIDInsensitively(v.Input)
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

		if actual.Publisher != v.Expected.Publisher {
			t.Fatalf("Expected %q but got %q for Publisher", v.Expected.Publisher, actual.Publisher)
		}

		if actual.ExtensionType != v.Expected.ExtensionType {
			t.Fatalf("Expected %q but got %q for ExtensionType", v.Expected.ExtensionType, actual.ExtensionType)
		}

	}
}

func TestSegmentsForExtensionTypeId(t *testing.T) {
	segments := ExtensionTypeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ExtensionTypeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
