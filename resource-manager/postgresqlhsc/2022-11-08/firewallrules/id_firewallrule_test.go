package firewallrules

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &FirewallRuleId{}

func TestNewFirewallRuleID(t *testing.T) {
	id := NewFirewallRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverGroupsv2Value", "firewallRuleValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.ServerGroupsv2Name != "serverGroupsv2Value" {
		t.Fatalf("Expected %q but got %q for Segment 'ServerGroupsv2Name'", id.ServerGroupsv2Name, "serverGroupsv2Value")
	}

	if id.FirewallRuleName != "firewallRuleValue" {
		t.Fatalf("Expected %q but got %q for Segment 'FirewallRuleName'", id.FirewallRuleName, "firewallRuleValue")
	}
}

func TestFormatFirewallRuleID(t *testing.T) {
	actual := NewFirewallRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverGroupsv2Value", "firewallRuleValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/serverGroupsv2Value/firewallRules/firewallRuleValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseFirewallRuleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *FirewallRuleId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/serverGroupsv2",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/serverGroupsv2Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/serverGroupsv2Value/firewallRules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/serverGroupsv2Value/firewallRules/firewallRuleValue",
			Expected: &FirewallRuleId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:  "example-resource-group",
				ServerGroupsv2Name: "serverGroupsv2Value",
				FirewallRuleName:   "firewallRuleValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/serverGroupsv2Value/firewallRules/firewallRuleValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseFirewallRuleID(v.Input)
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

		if actual.ServerGroupsv2Name != v.Expected.ServerGroupsv2Name {
			t.Fatalf("Expected %q but got %q for ServerGroupsv2Name", v.Expected.ServerGroupsv2Name, actual.ServerGroupsv2Name)
		}

		if actual.FirewallRuleName != v.Expected.FirewallRuleName {
			t.Fatalf("Expected %q but got %q for FirewallRuleName", v.Expected.FirewallRuleName, actual.FirewallRuleName)
		}

	}
}

func TestParseFirewallRuleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *FirewallRuleId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dBfOrPoStGrEsQl",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/serverGroupsv2",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dBfOrPoStGrEsQl/sErVeRgRoUpSv2",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/serverGroupsv2Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dBfOrPoStGrEsQl/sErVeRgRoUpSv2/sErVeRgRoUpSv2vAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/serverGroupsv2Value/firewallRules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dBfOrPoStGrEsQl/sErVeRgRoUpSv2/sErVeRgRoUpSv2vAlUe/fIrEwAlLrUlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/serverGroupsv2Value/firewallRules/firewallRuleValue",
			Expected: &FirewallRuleId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:  "example-resource-group",
				ServerGroupsv2Name: "serverGroupsv2Value",
				FirewallRuleName:   "firewallRuleValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/serverGroupsv2Value/firewallRules/firewallRuleValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dBfOrPoStGrEsQl/sErVeRgRoUpSv2/sErVeRgRoUpSv2vAlUe/fIrEwAlLrUlEs/fIrEwAlLrUlEvAlUe",
			Expected: &FirewallRuleId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:  "eXaMpLe-rEsOuRcE-GrOuP",
				ServerGroupsv2Name: "sErVeRgRoUpSv2vAlUe",
				FirewallRuleName:   "fIrEwAlLrUlEvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dBfOrPoStGrEsQl/sErVeRgRoUpSv2/sErVeRgRoUpSv2vAlUe/fIrEwAlLrUlEs/fIrEwAlLrUlEvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseFirewallRuleIDInsensitively(v.Input)
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

		if actual.ServerGroupsv2Name != v.Expected.ServerGroupsv2Name {
			t.Fatalf("Expected %q but got %q for ServerGroupsv2Name", v.Expected.ServerGroupsv2Name, actual.ServerGroupsv2Name)
		}

		if actual.FirewallRuleName != v.Expected.FirewallRuleName {
			t.Fatalf("Expected %q but got %q for FirewallRuleName", v.Expected.FirewallRuleName, actual.FirewallRuleName)
		}

	}
}

func TestSegmentsForFirewallRuleId(t *testing.T) {
	segments := FirewallRuleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("FirewallRuleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
