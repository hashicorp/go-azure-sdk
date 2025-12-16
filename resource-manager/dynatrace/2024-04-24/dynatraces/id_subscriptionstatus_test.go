package dynatraces

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SubscriptionStatusId{}

func TestNewSubscriptionStatusID(t *testing.T) {
	id := NewSubscriptionStatusID("12345678-1234-9876-4563-123456789012", "dynatraceEnvironmentId")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.DynatraceEnvironmentId != "dynatraceEnvironmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'DynatraceEnvironmentId'", id.DynatraceEnvironmentId, "dynatraceEnvironmentId")
	}
}

func TestFormatSubscriptionStatusID(t *testing.T) {
	actual := NewSubscriptionStatusID("12345678-1234-9876-4563-123456789012", "dynatraceEnvironmentId").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Dynatrace.Observability/subscriptionStatuses/dynatraceEnvironmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSubscriptionStatusID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SubscriptionStatusId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Dynatrace.Observability",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Dynatrace.Observability/subscriptionStatuses",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Dynatrace.Observability/subscriptionStatuses/dynatraceEnvironmentId",
			Expected: &SubscriptionStatusId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				DynatraceEnvironmentId: "dynatraceEnvironmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Dynatrace.Observability/subscriptionStatuses/dynatraceEnvironmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSubscriptionStatusID(v.Input)
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

		if actual.DynatraceEnvironmentId != v.Expected.DynatraceEnvironmentId {
			t.Fatalf("Expected %q but got %q for DynatraceEnvironmentId", v.Expected.DynatraceEnvironmentId, actual.DynatraceEnvironmentId)
		}

	}
}

func TestParseSubscriptionStatusIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SubscriptionStatusId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Dynatrace.Observability",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/dYnAtRaCe.oBsErVaBiLiTy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Dynatrace.Observability/subscriptionStatuses",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/dYnAtRaCe.oBsErVaBiLiTy/sUbScRiPtIoNsTaTuSeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Dynatrace.Observability/subscriptionStatuses/dynatraceEnvironmentId",
			Expected: &SubscriptionStatusId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				DynatraceEnvironmentId: "dynatraceEnvironmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Dynatrace.Observability/subscriptionStatuses/dynatraceEnvironmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/dYnAtRaCe.oBsErVaBiLiTy/sUbScRiPtIoNsTaTuSeS/dYnAtRaCeEnViRoNmEnTiD",
			Expected: &SubscriptionStatusId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				DynatraceEnvironmentId: "dYnAtRaCeEnViRoNmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/dYnAtRaCe.oBsErVaBiLiTy/sUbScRiPtIoNsTaTuSeS/dYnAtRaCeEnViRoNmEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSubscriptionStatusIDInsensitively(v.Input)
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

		if actual.DynatraceEnvironmentId != v.Expected.DynatraceEnvironmentId {
			t.Fatalf("Expected %q but got %q for DynatraceEnvironmentId", v.Expected.DynatraceEnvironmentId, actual.DynatraceEnvironmentId)
		}

	}
}

func TestSegmentsForSubscriptionStatusId(t *testing.T) {
	segments := SubscriptionStatusId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SubscriptionStatusId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
