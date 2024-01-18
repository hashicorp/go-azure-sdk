package topleveldomains

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &TopLevelDomainId{}

func TestNewTopLevelDomainID(t *testing.T) {
	id := NewTopLevelDomainID("12345678-1234-9876-4563-123456789012", "topLevelDomainValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.TopLevelDomainName != "topLevelDomainValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TopLevelDomainName'", id.TopLevelDomainName, "topLevelDomainValue")
	}
}

func TestFormatTopLevelDomainID(t *testing.T) {
	actual := NewTopLevelDomainID("12345678-1234-9876-4563-123456789012", "topLevelDomainValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.DomainRegistration/topLevelDomains/topLevelDomainValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseTopLevelDomainID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TopLevelDomainId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.DomainRegistration",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.DomainRegistration/topLevelDomains",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.DomainRegistration/topLevelDomains/topLevelDomainValue",
			Expected: &TopLevelDomainId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				TopLevelDomainName: "topLevelDomainValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.DomainRegistration/topLevelDomains/topLevelDomainValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTopLevelDomainID(v.Input)
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

		if actual.TopLevelDomainName != v.Expected.TopLevelDomainName {
			t.Fatalf("Expected %q but got %q for TopLevelDomainName", v.Expected.TopLevelDomainName, actual.TopLevelDomainName)
		}

	}
}

func TestParseTopLevelDomainIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TopLevelDomainId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.DomainRegistration",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.dOmAiNrEgIsTrAtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.DomainRegistration/topLevelDomains",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.dOmAiNrEgIsTrAtIoN/tOpLeVeLdOmAiNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.DomainRegistration/topLevelDomains/topLevelDomainValue",
			Expected: &TopLevelDomainId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				TopLevelDomainName: "topLevelDomainValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.DomainRegistration/topLevelDomains/topLevelDomainValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.dOmAiNrEgIsTrAtIoN/tOpLeVeLdOmAiNs/tOpLeVeLdOmAiNvAlUe",
			Expected: &TopLevelDomainId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				TopLevelDomainName: "tOpLeVeLdOmAiNvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.dOmAiNrEgIsTrAtIoN/tOpLeVeLdOmAiNs/tOpLeVeLdOmAiNvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTopLevelDomainIDInsensitively(v.Input)
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

		if actual.TopLevelDomainName != v.Expected.TopLevelDomainName {
			t.Fatalf("Expected %q but got %q for TopLevelDomainName", v.Expected.TopLevelDomainName, actual.TopLevelDomainName)
		}

	}
}

func TestSegmentsForTopLevelDomainId(t *testing.T) {
	segments := TopLevelDomainId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("TopLevelDomainId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
