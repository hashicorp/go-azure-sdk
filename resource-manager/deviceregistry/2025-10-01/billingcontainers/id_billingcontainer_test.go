package billingcontainers

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &BillingContainerId{}

func TestNewBillingContainerID(t *testing.T) {
	id := NewBillingContainerID("12345678-1234-9876-4563-123456789012", "billingContainerName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.BillingContainerName != "billingContainerName" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingContainerName'", id.BillingContainerName, "billingContainerName")
	}
}

func TestFormatBillingContainerID(t *testing.T) {
	actual := NewBillingContainerID("12345678-1234-9876-4563-123456789012", "billingContainerName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.DeviceRegistry/billingContainers/billingContainerName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseBillingContainerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingContainerId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.DeviceRegistry",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.DeviceRegistry/billingContainers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.DeviceRegistry/billingContainers/billingContainerName",
			Expected: &BillingContainerId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				BillingContainerName: "billingContainerName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.DeviceRegistry/billingContainers/billingContainerName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingContainerID(v.Input)
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

		if actual.BillingContainerName != v.Expected.BillingContainerName {
			t.Fatalf("Expected %q but got %q for BillingContainerName", v.Expected.BillingContainerName, actual.BillingContainerName)
		}

	}
}

func TestParseBillingContainerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingContainerId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.DeviceRegistry",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.dEvIcErEgIsTrY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.DeviceRegistry/billingContainers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.dEvIcErEgIsTrY/bIlLiNgCoNtAiNeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.DeviceRegistry/billingContainers/billingContainerName",
			Expected: &BillingContainerId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				BillingContainerName: "billingContainerName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.DeviceRegistry/billingContainers/billingContainerName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.dEvIcErEgIsTrY/bIlLiNgCoNtAiNeRs/bIlLiNgCoNtAiNeRnAmE",
			Expected: &BillingContainerId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				BillingContainerName: "bIlLiNgCoNtAiNeRnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.dEvIcErEgIsTrY/bIlLiNgCoNtAiNeRs/bIlLiNgCoNtAiNeRnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingContainerIDInsensitively(v.Input)
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

		if actual.BillingContainerName != v.Expected.BillingContainerName {
			t.Fatalf("Expected %q but got %q for BillingContainerName", v.Expected.BillingContainerName, actual.BillingContainerName)
		}

	}
}

func TestSegmentsForBillingContainerId(t *testing.T) {
	segments := BillingContainerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("BillingContainerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
