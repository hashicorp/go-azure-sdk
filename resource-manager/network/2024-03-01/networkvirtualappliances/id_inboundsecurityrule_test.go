package networkvirtualappliances

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &InboundSecurityRuleId{}

func TestNewInboundSecurityRuleID(t *testing.T) {
	id := NewInboundSecurityRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "networkVirtualApplianceName", "inboundSecurityRuleName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.NetworkVirtualApplianceName != "networkVirtualApplianceName" {
		t.Fatalf("Expected %q but got %q for Segment 'NetworkVirtualApplianceName'", id.NetworkVirtualApplianceName, "networkVirtualApplianceName")
	}

	if id.InboundSecurityRuleName != "inboundSecurityRuleName" {
		t.Fatalf("Expected %q but got %q for Segment 'InboundSecurityRuleName'", id.InboundSecurityRuleName, "inboundSecurityRuleName")
	}
}

func TestFormatInboundSecurityRuleID(t *testing.T) {
	actual := NewInboundSecurityRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "networkVirtualApplianceName", "inboundSecurityRuleName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkVirtualAppliances/networkVirtualApplianceName/inboundSecurityRules/inboundSecurityRuleName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseInboundSecurityRuleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *InboundSecurityRuleId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkVirtualAppliances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkVirtualAppliances/networkVirtualApplianceName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkVirtualAppliances/networkVirtualApplianceName/inboundSecurityRules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkVirtualAppliances/networkVirtualApplianceName/inboundSecurityRules/inboundSecurityRuleName",
			Expected: &InboundSecurityRuleId{
				SubscriptionId:              "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:           "example-resource-group",
				NetworkVirtualApplianceName: "networkVirtualApplianceName",
				InboundSecurityRuleName:     "inboundSecurityRuleName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkVirtualAppliances/networkVirtualApplianceName/inboundSecurityRules/inboundSecurityRuleName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseInboundSecurityRuleID(v.Input)
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

		if actual.NetworkVirtualApplianceName != v.Expected.NetworkVirtualApplianceName {
			t.Fatalf("Expected %q but got %q for NetworkVirtualApplianceName", v.Expected.NetworkVirtualApplianceName, actual.NetworkVirtualApplianceName)
		}

		if actual.InboundSecurityRuleName != v.Expected.InboundSecurityRuleName {
			t.Fatalf("Expected %q but got %q for InboundSecurityRuleName", v.Expected.InboundSecurityRuleName, actual.InboundSecurityRuleName)
		}

	}
}

func TestParseInboundSecurityRuleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *InboundSecurityRuleId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkVirtualAppliances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkViRtUaLaPpLiAnCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkVirtualAppliances/networkVirtualApplianceName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkViRtUaLaPpLiAnCeS/nEtWoRkViRtUaLaPpLiAnCeNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkVirtualAppliances/networkVirtualApplianceName/inboundSecurityRules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkViRtUaLaPpLiAnCeS/nEtWoRkViRtUaLaPpLiAnCeNaMe/iNbOuNdSeCuRiTyRuLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkVirtualAppliances/networkVirtualApplianceName/inboundSecurityRules/inboundSecurityRuleName",
			Expected: &InboundSecurityRuleId{
				SubscriptionId:              "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:           "example-resource-group",
				NetworkVirtualApplianceName: "networkVirtualApplianceName",
				InboundSecurityRuleName:     "inboundSecurityRuleName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/networkVirtualAppliances/networkVirtualApplianceName/inboundSecurityRules/inboundSecurityRuleName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkViRtUaLaPpLiAnCeS/nEtWoRkViRtUaLaPpLiAnCeNaMe/iNbOuNdSeCuRiTyRuLeS/iNbOuNdSeCuRiTyRuLeNaMe",
			Expected: &InboundSecurityRuleId{
				SubscriptionId:              "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:           "eXaMpLe-rEsOuRcE-GrOuP",
				NetworkVirtualApplianceName: "nEtWoRkViRtUaLaPpLiAnCeNaMe",
				InboundSecurityRuleName:     "iNbOuNdSeCuRiTyRuLeNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkViRtUaLaPpLiAnCeS/nEtWoRkViRtUaLaPpLiAnCeNaMe/iNbOuNdSeCuRiTyRuLeS/iNbOuNdSeCuRiTyRuLeNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseInboundSecurityRuleIDInsensitively(v.Input)
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

		if actual.NetworkVirtualApplianceName != v.Expected.NetworkVirtualApplianceName {
			t.Fatalf("Expected %q but got %q for NetworkVirtualApplianceName", v.Expected.NetworkVirtualApplianceName, actual.NetworkVirtualApplianceName)
		}

		if actual.InboundSecurityRuleName != v.Expected.InboundSecurityRuleName {
			t.Fatalf("Expected %q but got %q for InboundSecurityRuleName", v.Expected.InboundSecurityRuleName, actual.InboundSecurityRuleName)
		}

	}
}

func TestSegmentsForInboundSecurityRuleId(t *testing.T) {
	segments := InboundSecurityRuleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("InboundSecurityRuleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
