package eventsubscriptions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScopedEventSubscriptionId{}

func TestNewScopedEventSubscriptionID(t *testing.T) {
	id := NewScopedEventSubscriptionID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "eventSubscriptionName")

	if id.Scope != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'Scope'", id.Scope, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.EventSubscriptionName != "eventSubscriptionName" {
		t.Fatalf("Expected %q but got %q for Segment 'EventSubscriptionName'", id.EventSubscriptionName, "eventSubscriptionName")
	}
}

func TestFormatScopedEventSubscriptionID(t *testing.T) {
	actual := NewScopedEventSubscriptionID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "eventSubscriptionName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.EventGrid/eventSubscriptions/eventSubscriptionName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScopedEventSubscriptionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedEventSubscriptionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.EventGrid",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.EventGrid/eventSubscriptions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.EventGrid/eventSubscriptions/eventSubscriptionName",
			Expected: &ScopedEventSubscriptionId{
				Scope:                 "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				EventSubscriptionName: "eventSubscriptionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.EventGrid/eventSubscriptions/eventSubscriptionName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedEventSubscriptionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.Scope != v.Expected.Scope {
			t.Fatalf("Expected %q but got %q for Scope", v.Expected.Scope, actual.Scope)
		}

		if actual.EventSubscriptionName != v.Expected.EventSubscriptionName {
			t.Fatalf("Expected %q but got %q for EventSubscriptionName", v.Expected.EventSubscriptionName, actual.EventSubscriptionName)
		}

	}
}

func TestParseScopedEventSubscriptionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedEventSubscriptionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.EventGrid",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.eVeNtGrId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.EventGrid/eventSubscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.eVeNtGrId/eVeNtSuBsCrIpTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.EventGrid/eventSubscriptions/eventSubscriptionName",
			Expected: &ScopedEventSubscriptionId{
				Scope:                 "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				EventSubscriptionName: "eventSubscriptionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.EventGrid/eventSubscriptions/eventSubscriptionName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.eVeNtGrId/eVeNtSuBsCrIpTiOnS/eVeNtSuBsCrIpTiOnNaMe",
			Expected: &ScopedEventSubscriptionId{
				Scope:                 "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				EventSubscriptionName: "eVeNtSuBsCrIpTiOnNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.eVeNtGrId/eVeNtSuBsCrIpTiOnS/eVeNtSuBsCrIpTiOnNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedEventSubscriptionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.Scope != v.Expected.Scope {
			t.Fatalf("Expected %q but got %q for Scope", v.Expected.Scope, actual.Scope)
		}

		if actual.EventSubscriptionName != v.Expected.EventSubscriptionName {
			t.Fatalf("Expected %q but got %q for EventSubscriptionName", v.Expected.EventSubscriptionName, actual.EventSubscriptionName)
		}

	}
}

func TestSegmentsForScopedEventSubscriptionId(t *testing.T) {
	segments := ScopedEventSubscriptionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScopedEventSubscriptionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
