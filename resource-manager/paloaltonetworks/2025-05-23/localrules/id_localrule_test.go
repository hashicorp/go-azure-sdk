package localrules

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &LocalRuleId{}

func TestNewLocalRuleID(t *testing.T) {
	id := NewLocalRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "localRulestackName", "localRuleName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.LocalRulestackName != "localRulestackName" {
		t.Fatalf("Expected %q but got %q for Segment 'LocalRulestackName'", id.LocalRulestackName, "localRulestackName")
	}

	if id.LocalRuleName != "localRuleName" {
		t.Fatalf("Expected %q but got %q for Segment 'LocalRuleName'", id.LocalRuleName, "localRuleName")
	}
}

func TestFormatLocalRuleID(t *testing.T) {
	actual := NewLocalRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "localRulestackName", "localRuleName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/localRulestackName/localRules/localRuleName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseLocalRuleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LocalRuleId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/localRulestackName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/localRulestackName/localRules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/localRulestackName/localRules/localRuleName",
			Expected: &LocalRuleId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:  "example-resource-group",
				LocalRulestackName: "localRulestackName",
				LocalRuleName:      "localRuleName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/localRulestackName/localRules/localRuleName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLocalRuleID(v.Input)
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

		if actual.LocalRulestackName != v.Expected.LocalRulestackName {
			t.Fatalf("Expected %q but got %q for LocalRulestackName", v.Expected.LocalRulestackName, actual.LocalRulestackName)
		}

		if actual.LocalRuleName != v.Expected.LocalRuleName {
			t.Fatalf("Expected %q but got %q for LocalRuleName", v.Expected.LocalRuleName, actual.LocalRuleName)
		}

	}
}

func TestParseLocalRuleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LocalRuleId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/lOcAlRuLeStAcKs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/localRulestackName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/lOcAlRuLeStAcKs/lOcAlRuLeStAcKnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/localRulestackName/localRules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/lOcAlRuLeStAcKs/lOcAlRuLeStAcKnAmE/lOcAlRuLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/localRulestackName/localRules/localRuleName",
			Expected: &LocalRuleId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:  "example-resource-group",
				LocalRulestackName: "localRulestackName",
				LocalRuleName:      "localRuleName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/localRulestackName/localRules/localRuleName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/lOcAlRuLeStAcKs/lOcAlRuLeStAcKnAmE/lOcAlRuLeS/lOcAlRuLeNaMe",
			Expected: &LocalRuleId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:  "eXaMpLe-rEsOuRcE-GrOuP",
				LocalRulestackName: "lOcAlRuLeStAcKnAmE",
				LocalRuleName:      "lOcAlRuLeNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/lOcAlRuLeStAcKs/lOcAlRuLeStAcKnAmE/lOcAlRuLeS/lOcAlRuLeNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLocalRuleIDInsensitively(v.Input)
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

		if actual.LocalRulestackName != v.Expected.LocalRulestackName {
			t.Fatalf("Expected %q but got %q for LocalRulestackName", v.Expected.LocalRulestackName, actual.LocalRulestackName)
		}

		if actual.LocalRuleName != v.Expected.LocalRuleName {
			t.Fatalf("Expected %q but got %q for LocalRuleName", v.Expected.LocalRuleName, actual.LocalRuleName)
		}

	}
}

func TestSegmentsForLocalRuleId(t *testing.T) {
	segments := LocalRuleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("LocalRuleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
