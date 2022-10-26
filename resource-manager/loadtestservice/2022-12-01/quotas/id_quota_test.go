package quotas

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = QuotaId{}

func TestNewQuotaID(t *testing.T) {
	id := NewQuotaID("12345678-1234-9876-4563-123456789012", "locationValue", "quotaBucketValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.Location != "locationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'Location'", id.Location, "locationValue")
	}

	if id.QuotaBucketName != "quotaBucketValue" {
		t.Fatalf("Expected %q but got %q for Segment 'QuotaBucketName'", id.QuotaBucketName, "quotaBucketValue")
	}
}

func TestFormatQuotaID(t *testing.T) {
	actual := NewQuotaID("12345678-1234-9876-4563-123456789012", "locationValue", "quotaBucketValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.LoadTestService/locations/locationValue/quotas/quotaBucketValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseQuotaID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *QuotaId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.LoadTestService",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.LoadTestService/locations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.LoadTestService/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.LoadTestService/locations/locationValue/quotas",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.LoadTestService/locations/locationValue/quotas/quotaBucketValue",
			Expected: &QuotaId{
				SubscriptionId:  "12345678-1234-9876-4563-123456789012",
				Location:        "locationValue",
				QuotaBucketName: "quotaBucketValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.LoadTestService/locations/locationValue/quotas/quotaBucketValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseQuotaID(v.Input)
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

		if actual.QuotaBucketName != v.Expected.QuotaBucketName {
			t.Fatalf("Expected %q but got %q for QuotaBucketName", v.Expected.QuotaBucketName, actual.QuotaBucketName)
		}

	}
}

func TestParseQuotaIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *QuotaId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.LoadTestService",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.lOaDtEsTsErViCe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.LoadTestService/locations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.lOaDtEsTsErViCe/lOcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.LoadTestService/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.lOaDtEsTsErViCe/lOcAtIoNs/lOcAtIoNvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.LoadTestService/locations/locationValue/quotas",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.lOaDtEsTsErViCe/lOcAtIoNs/lOcAtIoNvAlUe/qUoTaS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.LoadTestService/locations/locationValue/quotas/quotaBucketValue",
			Expected: &QuotaId{
				SubscriptionId:  "12345678-1234-9876-4563-123456789012",
				Location:        "locationValue",
				QuotaBucketName: "quotaBucketValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.LoadTestService/locations/locationValue/quotas/quotaBucketValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.lOaDtEsTsErViCe/lOcAtIoNs/lOcAtIoNvAlUe/qUoTaS/qUoTaBuCkEtVaLuE",
			Expected: &QuotaId{
				SubscriptionId:  "12345678-1234-9876-4563-123456789012",
				Location:        "lOcAtIoNvAlUe",
				QuotaBucketName: "qUoTaBuCkEtVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.lOaDtEsTsErViCe/lOcAtIoNs/lOcAtIoNvAlUe/qUoTaS/qUoTaBuCkEtVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseQuotaIDInsensitively(v.Input)
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

		if actual.QuotaBucketName != v.Expected.QuotaBucketName {
			t.Fatalf("Expected %q but got %q for QuotaBucketName", v.Expected.QuotaBucketName, actual.QuotaBucketName)
		}

	}
}

func TestSegmentsForQuotaId(t *testing.T) {
	segments := QuotaId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("QuotaId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
