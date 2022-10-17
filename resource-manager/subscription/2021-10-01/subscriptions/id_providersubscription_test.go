package subscriptions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ProviderSubscriptionId{}

func TestNewProviderSubscriptionID(t *testing.T) {
	id := NewProviderSubscriptionID("12345678-1234-9876-4563-123456789012")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}
}

func TestFormatProviderSubscriptionID(t *testing.T) {
	actual := NewProviderSubscriptionID("12345678-1234-9876-4563-123456789012").ID()
	expected := "/providers/Microsoft.Subscription/subscriptions/12345678-1234-9876-4563-123456789012"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProviderSubscriptionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ProviderSubscriptionId
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
			Input: "/providers/Microsoft.Subscription",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Subscription/subscriptions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Subscription/subscriptions/12345678-1234-9876-4563-123456789012",
			Expected: &ProviderSubscriptionId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Subscription/subscriptions/12345678-1234-9876-4563-123456789012/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviderSubscriptionID(v.Input)
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

	}
}

func TestParseProviderSubscriptionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ProviderSubscriptionId
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
			Input: "/providers/Microsoft.Subscription",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.sUbScRiPtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Subscription/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.sUbScRiPtIoN/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Subscription/subscriptions/12345678-1234-9876-4563-123456789012",
			Expected: &ProviderSubscriptionId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Subscription/subscriptions/12345678-1234-9876-4563-123456789012/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.sUbScRiPtIoN/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Expected: &ProviderSubscriptionId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.sUbScRiPtIoN/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviderSubscriptionIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForProviderSubscriptionId(t *testing.T) {
	segments := ProviderSubscriptionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ProviderSubscriptionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
