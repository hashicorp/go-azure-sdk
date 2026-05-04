package raiexternalsafetyprovider

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RaiExternalSafetyProviderId{}

func TestNewRaiExternalSafetyProviderID(t *testing.T) {
	id := NewRaiExternalSafetyProviderID("12345678-1234-9876-4563-123456789012", "raiExternalSafetyProviderName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.RaiExternalSafetyProviderName != "raiExternalSafetyProviderName" {
		t.Fatalf("Expected %q but got %q for Segment 'RaiExternalSafetyProviderName'", id.RaiExternalSafetyProviderName, "raiExternalSafetyProviderName")
	}
}

func TestFormatRaiExternalSafetyProviderID(t *testing.T) {
	actual := NewRaiExternalSafetyProviderID("12345678-1234-9876-4563-123456789012", "raiExternalSafetyProviderName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/raiExternalSafetyProviders/raiExternalSafetyProviderName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRaiExternalSafetyProviderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RaiExternalSafetyProviderId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/raiExternalSafetyProviders",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/raiExternalSafetyProviders/raiExternalSafetyProviderName",
			Expected: &RaiExternalSafetyProviderId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				RaiExternalSafetyProviderName: "raiExternalSafetyProviderName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/raiExternalSafetyProviders/raiExternalSafetyProviderName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRaiExternalSafetyProviderID(v.Input)
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

		if actual.RaiExternalSafetyProviderName != v.Expected.RaiExternalSafetyProviderName {
			t.Fatalf("Expected %q but got %q for RaiExternalSafetyProviderName", v.Expected.RaiExternalSafetyProviderName, actual.RaiExternalSafetyProviderName)
		}

	}
}

func TestParseRaiExternalSafetyProviderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RaiExternalSafetyProviderId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOgNiTiVeSeRvIcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/raiExternalSafetyProviders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOgNiTiVeSeRvIcEs/rAiExTeRnAlSaFeTyPrOvIdErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/raiExternalSafetyProviders/raiExternalSafetyProviderName",
			Expected: &RaiExternalSafetyProviderId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				RaiExternalSafetyProviderName: "raiExternalSafetyProviderName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/raiExternalSafetyProviders/raiExternalSafetyProviderName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOgNiTiVeSeRvIcEs/rAiExTeRnAlSaFeTyPrOvIdErS/rAiExTeRnAlSaFeTyPrOvIdErNaMe",
			Expected: &RaiExternalSafetyProviderId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				RaiExternalSafetyProviderName: "rAiExTeRnAlSaFeTyPrOvIdErNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOgNiTiVeSeRvIcEs/rAiExTeRnAlSaFeTyPrOvIdErS/rAiExTeRnAlSaFeTyPrOvIdErNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRaiExternalSafetyProviderIDInsensitively(v.Input)
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

		if actual.RaiExternalSafetyProviderName != v.Expected.RaiExternalSafetyProviderName {
			t.Fatalf("Expected %q but got %q for RaiExternalSafetyProviderName", v.Expected.RaiExternalSafetyProviderName, actual.RaiExternalSafetyProviderName)
		}

	}
}

func TestSegmentsForRaiExternalSafetyProviderId(t *testing.T) {
	segments := RaiExternalSafetyProviderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RaiExternalSafetyProviderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
