package subscriptionfeatureregistrations

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SubscriptionFeatureRegistrationId{}

func TestNewSubscriptionFeatureRegistrationID(t *testing.T) {
	id := NewSubscriptionFeatureRegistrationID("12345678-1234-9876-4563-123456789012", "featureProviderName", "subscriptionFeatureRegistrationName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.FeatureProviderName != "featureProviderName" {
		t.Fatalf("Expected %q but got %q for Segment 'FeatureProviderName'", id.FeatureProviderName, "featureProviderName")
	}

	if id.SubscriptionFeatureRegistrationName != "subscriptionFeatureRegistrationName" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionFeatureRegistrationName'", id.SubscriptionFeatureRegistrationName, "subscriptionFeatureRegistrationName")
	}
}

func TestFormatSubscriptionFeatureRegistrationID(t *testing.T) {
	actual := NewSubscriptionFeatureRegistrationID("12345678-1234-9876-4563-123456789012", "featureProviderName", "subscriptionFeatureRegistrationName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Features/featureProviders/featureProviderName/subscriptionFeatureRegistrations/subscriptionFeatureRegistrationName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSubscriptionFeatureRegistrationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SubscriptionFeatureRegistrationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Features",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Features/featureProviders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Features/featureProviders/featureProviderName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Features/featureProviders/featureProviderName/subscriptionFeatureRegistrations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Features/featureProviders/featureProviderName/subscriptionFeatureRegistrations/subscriptionFeatureRegistrationName",
			Expected: &SubscriptionFeatureRegistrationId{
				SubscriptionId:                      "12345678-1234-9876-4563-123456789012",
				FeatureProviderName:                 "featureProviderName",
				SubscriptionFeatureRegistrationName: "subscriptionFeatureRegistrationName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Features/featureProviders/featureProviderName/subscriptionFeatureRegistrations/subscriptionFeatureRegistrationName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSubscriptionFeatureRegistrationID(v.Input)
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

		if actual.FeatureProviderName != v.Expected.FeatureProviderName {
			t.Fatalf("Expected %q but got %q for FeatureProviderName", v.Expected.FeatureProviderName, actual.FeatureProviderName)
		}

		if actual.SubscriptionFeatureRegistrationName != v.Expected.SubscriptionFeatureRegistrationName {
			t.Fatalf("Expected %q but got %q for SubscriptionFeatureRegistrationName", v.Expected.SubscriptionFeatureRegistrationName, actual.SubscriptionFeatureRegistrationName)
		}

	}
}

func TestParseSubscriptionFeatureRegistrationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SubscriptionFeatureRegistrationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Features",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.fEaTuReS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Features/featureProviders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.fEaTuReS/fEaTuRePrOvIdErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Features/featureProviders/featureProviderName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.fEaTuReS/fEaTuRePrOvIdErS/fEaTuRePrOvIdErNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Features/featureProviders/featureProviderName/subscriptionFeatureRegistrations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.fEaTuReS/fEaTuRePrOvIdErS/fEaTuRePrOvIdErNaMe/sUbScRiPtIoNfEaTuReReGiStRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Features/featureProviders/featureProviderName/subscriptionFeatureRegistrations/subscriptionFeatureRegistrationName",
			Expected: &SubscriptionFeatureRegistrationId{
				SubscriptionId:                      "12345678-1234-9876-4563-123456789012",
				FeatureProviderName:                 "featureProviderName",
				SubscriptionFeatureRegistrationName: "subscriptionFeatureRegistrationName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Features/featureProviders/featureProviderName/subscriptionFeatureRegistrations/subscriptionFeatureRegistrationName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.fEaTuReS/fEaTuRePrOvIdErS/fEaTuRePrOvIdErNaMe/sUbScRiPtIoNfEaTuReReGiStRaTiOnS/sUbScRiPtIoNfEaTuReReGiStRaTiOnNaMe",
			Expected: &SubscriptionFeatureRegistrationId{
				SubscriptionId:                      "12345678-1234-9876-4563-123456789012",
				FeatureProviderName:                 "fEaTuRePrOvIdErNaMe",
				SubscriptionFeatureRegistrationName: "sUbScRiPtIoNfEaTuReReGiStRaTiOnNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.fEaTuReS/fEaTuRePrOvIdErS/fEaTuRePrOvIdErNaMe/sUbScRiPtIoNfEaTuReReGiStRaTiOnS/sUbScRiPtIoNfEaTuReReGiStRaTiOnNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSubscriptionFeatureRegistrationIDInsensitively(v.Input)
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

		if actual.FeatureProviderName != v.Expected.FeatureProviderName {
			t.Fatalf("Expected %q but got %q for FeatureProviderName", v.Expected.FeatureProviderName, actual.FeatureProviderName)
		}

		if actual.SubscriptionFeatureRegistrationName != v.Expected.SubscriptionFeatureRegistrationName {
			t.Fatalf("Expected %q but got %q for SubscriptionFeatureRegistrationName", v.Expected.SubscriptionFeatureRegistrationName, actual.SubscriptionFeatureRegistrationName)
		}

	}
}

func TestSegmentsForSubscriptionFeatureRegistrationId(t *testing.T) {
	segments := SubscriptionFeatureRegistrationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SubscriptionFeatureRegistrationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
