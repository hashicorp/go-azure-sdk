package iotsecuritysolutionsanalytics

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AggregatedAlertId{}

func TestNewAggregatedAlertID(t *testing.T) {
	id := NewAggregatedAlertID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue", "aggregatedAlertValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.IotSecuritySolutionName != "iotSecuritySolutionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'IotSecuritySolutionName'", id.IotSecuritySolutionName, "iotSecuritySolutionValue")
	}

	if id.AggregatedAlertName != "aggregatedAlertValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AggregatedAlertName'", id.AggregatedAlertName, "aggregatedAlertValue")
	}
}

func TestFormatAggregatedAlertID(t *testing.T) {
	actual := NewAggregatedAlertID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue", "aggregatedAlertValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions/iotSecuritySolutionValue/analyticsModels/default/aggregatedAlerts/aggregatedAlertValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAggregatedAlertID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AggregatedAlertId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions/iotSecuritySolutionValue/analyticsModels/default/aggregatedAlerts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions/iotSecuritySolutionValue/analyticsModels/default/aggregatedAlerts/aggregatedAlertValue",
			Expected: &AggregatedAlertId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:       "example-resource-group",
				IotSecuritySolutionName: "iotSecuritySolutionValue",
				AggregatedAlertName:     "aggregatedAlertValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions/iotSecuritySolutionValue/analyticsModels/default/aggregatedAlerts/aggregatedAlertValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAggregatedAlertID(v.Input)
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

		if actual.AggregatedAlertName != v.Expected.AggregatedAlertName {
			t.Fatalf("Expected %q but got %q for AggregatedAlertName", v.Expected.AggregatedAlertName, actual.AggregatedAlertName)
		}

	}
}

func TestParseAggregatedAlertIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AggregatedAlertId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions/iotSecuritySolutionValue/analyticsModels/default/aggregatedAlerts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/iOtSeCuRiTySoLuTiOnS/iOtSeCuRiTySoLuTiOnVaLuE/aNaLyTiCsMoDeLs/dEfAuLt/aGgReGaTeDaLeRtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions/iotSecuritySolutionValue/analyticsModels/default/aggregatedAlerts/aggregatedAlertValue",
			Expected: &AggregatedAlertId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:       "example-resource-group",
				IotSecuritySolutionName: "iotSecuritySolutionValue",
				AggregatedAlertName:     "aggregatedAlertValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/iotSecuritySolutions/iotSecuritySolutionValue/analyticsModels/default/aggregatedAlerts/aggregatedAlertValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/iOtSeCuRiTySoLuTiOnS/iOtSeCuRiTySoLuTiOnVaLuE/aNaLyTiCsMoDeLs/dEfAuLt/aGgReGaTeDaLeRtS/aGgReGaTeDaLeRtVaLuE",
			Expected: &AggregatedAlertId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:       "eXaMpLe-rEsOuRcE-GrOuP",
				IotSecuritySolutionName: "iOtSeCuRiTySoLuTiOnVaLuE",
				AggregatedAlertName:     "aGgReGaTeDaLeRtVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/iOtSeCuRiTySoLuTiOnS/iOtSeCuRiTySoLuTiOnVaLuE/aNaLyTiCsMoDeLs/dEfAuLt/aGgReGaTeDaLeRtS/aGgReGaTeDaLeRtVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAggregatedAlertIDInsensitively(v.Input)
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

		if actual.AggregatedAlertName != v.Expected.AggregatedAlertName {
			t.Fatalf("Expected %q but got %q for AggregatedAlertName", v.Expected.AggregatedAlertName, actual.AggregatedAlertName)
		}

	}
}

func TestSegmentsForAggregatedAlertId(t *testing.T) {
	segments := AggregatedAlertId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AggregatedAlertId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
