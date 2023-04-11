package alertssuppressionrules

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AlertsSuppressionRuleId{}

func TestNewAlertsSuppressionRuleID(t *testing.T) {
	id := NewAlertsSuppressionRuleID("12345678-1234-9876-4563-123456789012", "alertsSuppressionRuleValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.AlertsSuppressionRuleName != "alertsSuppressionRuleValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AlertsSuppressionRuleName'", id.AlertsSuppressionRuleName, "alertsSuppressionRuleValue")
	}
}

func TestFormatAlertsSuppressionRuleID(t *testing.T) {
	actual := NewAlertsSuppressionRuleID("12345678-1234-9876-4563-123456789012", "alertsSuppressionRuleValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/alertsSuppressionRules/alertsSuppressionRuleValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAlertsSuppressionRuleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AlertsSuppressionRuleId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/alertsSuppressionRules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/alertsSuppressionRules/alertsSuppressionRuleValue",
			Expected: &AlertsSuppressionRuleId{
				SubscriptionId:            "12345678-1234-9876-4563-123456789012",
				AlertsSuppressionRuleName: "alertsSuppressionRuleValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/alertsSuppressionRules/alertsSuppressionRuleValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAlertsSuppressionRuleID(v.Input)
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

		if actual.AlertsSuppressionRuleName != v.Expected.AlertsSuppressionRuleName {
			t.Fatalf("Expected %q but got %q for AlertsSuppressionRuleName", v.Expected.AlertsSuppressionRuleName, actual.AlertsSuppressionRuleName)
		}

	}
}

func TestParseAlertsSuppressionRuleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AlertsSuppressionRuleId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sEcUrItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/alertsSuppressionRules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sEcUrItY/aLeRtSsUpPrEsSiOnRuLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/alertsSuppressionRules/alertsSuppressionRuleValue",
			Expected: &AlertsSuppressionRuleId{
				SubscriptionId:            "12345678-1234-9876-4563-123456789012",
				AlertsSuppressionRuleName: "alertsSuppressionRuleValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/alertsSuppressionRules/alertsSuppressionRuleValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sEcUrItY/aLeRtSsUpPrEsSiOnRuLeS/aLeRtSsUpPrEsSiOnRuLeVaLuE",
			Expected: &AlertsSuppressionRuleId{
				SubscriptionId:            "12345678-1234-9876-4563-123456789012",
				AlertsSuppressionRuleName: "aLeRtSsUpPrEsSiOnRuLeVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sEcUrItY/aLeRtSsUpPrEsSiOnRuLeS/aLeRtSsUpPrEsSiOnRuLeVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAlertsSuppressionRuleIDInsensitively(v.Input)
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

		if actual.AlertsSuppressionRuleName != v.Expected.AlertsSuppressionRuleName {
			t.Fatalf("Expected %q but got %q for AlertsSuppressionRuleName", v.Expected.AlertsSuppressionRuleName, actual.AlertsSuppressionRuleName)
		}

	}
}

func TestSegmentsForAlertsSuppressionRuleId(t *testing.T) {
	segments := AlertsSuppressionRuleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AlertsSuppressionRuleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
