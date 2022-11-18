package providers

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = SubscriptionProviderId{}

func TestNewSubscriptionProviderID(t *testing.T) {
	id := NewSubscriptionProviderID("12345678-1234-9876-4563-123456789012", "resourceProviderNamespaceValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceProviderNamespace != "resourceProviderNamespaceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceProviderNamespace'", id.ResourceProviderNamespace, "resourceProviderNamespaceValue")
	}
}

func TestFormatSubscriptionProviderID(t *testing.T) {
	actual := NewSubscriptionProviderID("12345678-1234-9876-4563-123456789012", "resourceProviderNamespaceValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/resourceProviderNamespaceValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSubscriptionProviderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SubscriptionProviderId
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
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/resourceProviderNamespaceValue",
			Expected: &SubscriptionProviderId{
				SubscriptionId:            "12345678-1234-9876-4563-123456789012",
				ResourceProviderNamespace: "resourceProviderNamespaceValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/resourceProviderNamespaceValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSubscriptionProviderID(v.Input)
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

		if actual.ResourceProviderNamespace != v.Expected.ResourceProviderNamespace {
			t.Fatalf("Expected %q but got %q for ResourceProviderNamespace", v.Expected.ResourceProviderNamespace, actual.ResourceProviderNamespace)
		}

	}
}

func TestParseSubscriptionProviderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SubscriptionProviderId
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
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/resourceProviderNamespaceValue",
			Expected: &SubscriptionProviderId{
				SubscriptionId:            "12345678-1234-9876-4563-123456789012",
				ResourceProviderNamespace: "resourceProviderNamespaceValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/resourceProviderNamespaceValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/rEsOuRcEpRoViDeRnAmEsPaCeVaLuE",
			Expected: &SubscriptionProviderId{
				SubscriptionId:            "12345678-1234-9876-4563-123456789012",
				ResourceProviderNamespace: "rEsOuRcEpRoViDeRnAmEsPaCeVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/rEsOuRcEpRoViDeRnAmEsPaCeVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSubscriptionProviderIDInsensitively(v.Input)
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

		if actual.ResourceProviderNamespace != v.Expected.ResourceProviderNamespace {
			t.Fatalf("Expected %q but got %q for ResourceProviderNamespace", v.Expected.ResourceProviderNamespace, actual.ResourceProviderNamespace)
		}

	}
}

func TestSegmentsForSubscriptionProviderId(t *testing.T) {
	segments := SubscriptionProviderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SubscriptionProviderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
