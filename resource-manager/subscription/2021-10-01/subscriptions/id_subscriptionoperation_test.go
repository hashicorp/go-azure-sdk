package subscriptions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SubscriptionOperationId{}

func TestNewSubscriptionOperationID(t *testing.T) {
	id := NewSubscriptionOperationID("operationIdValue")

	if id.OperationId != "operationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OperationId'", id.OperationId, "operationIdValue")
	}
}

func TestFormatSubscriptionOperationID(t *testing.T) {
	actual := NewSubscriptionOperationID("operationIdValue").ID()
	expected := "/providers/Microsoft.Subscription/subscriptionOperations/operationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSubscriptionOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SubscriptionOperationId
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
			Input: "/providers/Microsoft.Subscription/subscriptionOperations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Subscription/subscriptionOperations/operationIdValue",
			Expected: &SubscriptionOperationId{
				OperationId: "operationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Subscription/subscriptionOperations/operationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSubscriptionOperationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OperationId != v.Expected.OperationId {
			t.Fatalf("Expected %q but got %q for OperationId", v.Expected.OperationId, actual.OperationId)
		}

	}
}

func TestParseSubscriptionOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SubscriptionOperationId
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
			Input: "/providers/Microsoft.Subscription/subscriptionOperations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.sUbScRiPtIoN/sUbScRiPtIoNoPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Subscription/subscriptionOperations/operationIdValue",
			Expected: &SubscriptionOperationId{
				OperationId: "operationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Subscription/subscriptionOperations/operationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.sUbScRiPtIoN/sUbScRiPtIoNoPeRaTiOnS/oPeRaTiOnIdVaLuE",
			Expected: &SubscriptionOperationId{
				OperationId: "oPeRaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.sUbScRiPtIoN/sUbScRiPtIoNoPeRaTiOnS/oPeRaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSubscriptionOperationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OperationId != v.Expected.OperationId {
			t.Fatalf("Expected %q but got %q for OperationId", v.Expected.OperationId, actual.OperationId)
		}

	}
}

func TestSegmentsForSubscriptionOperationId(t *testing.T) {
	segments := SubscriptionOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SubscriptionOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
