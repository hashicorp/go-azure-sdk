package iotsecuritysolutionsanalytics

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AggregatedRecommendationId{}

func TestNewAggregatedRecommendationID(t *testing.T) {
	id := NewAggregatedRecommendationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue", "aggregatedRecommendationValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.IotSecuritySolutionName != "iotSecuritySolutionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'IotSecuritySolutionName'", id.IotSecuritySolutionName, "iotSecuritySolutionValue")
	}

	if id.AggregatedRecommendationName != "aggregatedRecommendationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AggregatedRecommendationName'", id.AggregatedRecommendationName, "aggregatedRecommendationValue")
	}
}

func TestFormatAggregatedRecommendationID(t *testing.T) {
	actual := NewAggregatedRecommendationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue", "aggregatedRecommendationValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions/iotSecuritySolutionValue/analyticsModels/default/aggregatedRecommendations/aggregatedRecommendationValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAggregatedRecommendationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AggregatedRecommendationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions/iotSecuritySolutionValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions/iotSecuritySolutionValue/analyticsModels",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions/iotSecuritySolutionValue/analyticsModels/default",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions/iotSecuritySolutionValue/analyticsModels/default/aggregatedRecommendations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions/iotSecuritySolutionValue/analyticsModels/default/aggregatedRecommendations/aggregatedRecommendationValue",
			Expected: &AggregatedRecommendationId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:            "example-resource-group",
				IotSecuritySolutionName:      "iotSecuritySolutionValue",
				AggregatedRecommendationName: "aggregatedRecommendationValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions/iotSecuritySolutionValue/analyticsModels/default/aggregatedRecommendations/aggregatedRecommendationValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAggregatedRecommendationID(v.Input)
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

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.IotSecuritySolutionName != v.Expected.IotSecuritySolutionName {
			t.Fatalf("Expected %q but got %q for IotSecuritySolutionName", v.Expected.IotSecuritySolutionName, actual.IotSecuritySolutionName)
		}

		if actual.AggregatedRecommendationName != v.Expected.AggregatedRecommendationName {
			t.Fatalf("Expected %q but got %q for AggregatedRecommendationName", v.Expected.AggregatedRecommendationName, actual.AggregatedRecommendationName)
		}

	}
}

func TestParseAggregatedRecommendationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AggregatedRecommendationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/iOtSeCuRiTySoLuTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions/iotSecuritySolutionValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/iOtSeCuRiTySoLuTiOnS/iOtSeCuRiTySoLuTiOnVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions/iotSecuritySolutionValue/analyticsModels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/iOtSeCuRiTySoLuTiOnS/iOtSeCuRiTySoLuTiOnVaLuE/aNaLyTiCsMoDeLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions/iotSecuritySolutionValue/analyticsModels/default",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/iOtSeCuRiTySoLuTiOnS/iOtSeCuRiTySoLuTiOnVaLuE/aNaLyTiCsMoDeLs/dEfAuLt",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions/iotSecuritySolutionValue/analyticsModels/default/aggregatedRecommendations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/iOtSeCuRiTySoLuTiOnS/iOtSeCuRiTySoLuTiOnVaLuE/aNaLyTiCsMoDeLs/dEfAuLt/aGgReGaTeDrEcOmMeNdAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions/iotSecuritySolutionValue/analyticsModels/default/aggregatedRecommendations/aggregatedRecommendationValue",
			Expected: &AggregatedRecommendationId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:            "example-resource-group",
				IotSecuritySolutionName:      "iotSecuritySolutionValue",
				AggregatedRecommendationName: "aggregatedRecommendationValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions/iotSecuritySolutionValue/analyticsModels/default/aggregatedRecommendations/aggregatedRecommendationValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/iOtSeCuRiTySoLuTiOnS/iOtSeCuRiTySoLuTiOnVaLuE/aNaLyTiCsMoDeLs/dEfAuLt/aGgReGaTeDrEcOmMeNdAtIoNs/aGgReGaTeDrEcOmMeNdAtIoNvAlUe",
			Expected: &AggregatedRecommendationId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:            "eXaMpLe-rEsOuRcE-GrOuP",
				IotSecuritySolutionName:      "iOtSeCuRiTySoLuTiOnVaLuE",
				AggregatedRecommendationName: "aGgReGaTeDrEcOmMeNdAtIoNvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/iOtSeCuRiTySoLuTiOnS/iOtSeCuRiTySoLuTiOnVaLuE/aNaLyTiCsMoDeLs/dEfAuLt/aGgReGaTeDrEcOmMeNdAtIoNs/aGgReGaTeDrEcOmMeNdAtIoNvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAggregatedRecommendationIDInsensitively(v.Input)
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

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.IotSecuritySolutionName != v.Expected.IotSecuritySolutionName {
			t.Fatalf("Expected %q but got %q for IotSecuritySolutionName", v.Expected.IotSecuritySolutionName, actual.IotSecuritySolutionName)
		}

		if actual.AggregatedRecommendationName != v.Expected.AggregatedRecommendationName {
			t.Fatalf("Expected %q but got %q for AggregatedRecommendationName", v.Expected.AggregatedRecommendationName, actual.AggregatedRecommendationName)
		}

	}
}

func TestSegmentsForAggregatedRecommendationId(t *testing.T) {
	segments := AggregatedRecommendationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AggregatedRecommendationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
