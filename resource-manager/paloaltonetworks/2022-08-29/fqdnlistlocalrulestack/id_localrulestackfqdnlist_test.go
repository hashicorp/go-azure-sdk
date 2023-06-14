package fqdnlistlocalrulestack

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = LocalRuleStackFqdnListId{}

func TestNewLocalRuleStackFqdnListID(t *testing.T) {
	id := NewLocalRuleStackFqdnListID("12345678-1234-9876-4563-123456789012", "example-resource-group", "localRuleStackValue", "fqdnListValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.LocalRuleStackName != "localRuleStackValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LocalRuleStackName'", id.LocalRuleStackName, "localRuleStackValue")
	}

	if id.FqdnListName != "fqdnListValue" {
		t.Fatalf("Expected %q but got %q for Segment 'FqdnListName'", id.FqdnListName, "fqdnListValue")
	}
}

func TestFormatLocalRuleStackFqdnListID(t *testing.T) {
	actual := NewLocalRuleStackFqdnListID("12345678-1234-9876-4563-123456789012", "example-resource-group", "localRuleStackValue", "fqdnListValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.CloudNGFW/localRuleStacks/localRuleStackValue/fqdnLists/fqdnListValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseLocalRuleStackFqdnListID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LocalRuleStackFqdnListId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.CloudNGFW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.CloudNGFW/localRuleStacks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.CloudNGFW/localRuleStacks/localRuleStackValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.CloudNGFW/localRuleStacks/localRuleStackValue/fqdnLists",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.CloudNGFW/localRuleStacks/localRuleStackValue/fqdnLists/fqdnListValue",
			Expected: &LocalRuleStackFqdnListId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:  "example-resource-group",
				LocalRuleStackName: "localRuleStackValue",
				FqdnListName:       "fqdnListValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.CloudNGFW/localRuleStacks/localRuleStackValue/fqdnLists/fqdnListValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLocalRuleStackFqdnListID(v.Input)
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

		if actual.LocalRuleStackName != v.Expected.LocalRuleStackName {
			t.Fatalf("Expected %q but got %q for LocalRuleStackName", v.Expected.LocalRuleStackName, actual.LocalRuleStackName)
		}

		if actual.FqdnListName != v.Expected.FqdnListName {
			t.Fatalf("Expected %q but got %q for FqdnListName", v.Expected.FqdnListName, actual.FqdnListName)
		}

	}
}

func TestParseLocalRuleStackFqdnListIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LocalRuleStackFqdnListId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.CloudNGFW",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.CloudNGFW/localRuleStacks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/lOcAlRuLeStAcKs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.CloudNGFW/localRuleStacks/localRuleStackValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/lOcAlRuLeStAcKs/lOcAlRuLeStAcKvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.CloudNGFW/localRuleStacks/localRuleStackValue/fqdnLists",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/lOcAlRuLeStAcKs/lOcAlRuLeStAcKvAlUe/fQdNlIsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.CloudNGFW/localRuleStacks/localRuleStackValue/fqdnLists/fqdnListValue",
			Expected: &LocalRuleStackFqdnListId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:  "example-resource-group",
				LocalRuleStackName: "localRuleStackValue",
				FqdnListName:       "fqdnListValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.CloudNGFW/localRuleStacks/localRuleStackValue/fqdnLists/fqdnListValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/lOcAlRuLeStAcKs/lOcAlRuLeStAcKvAlUe/fQdNlIsTs/fQdNlIsTvAlUe",
			Expected: &LocalRuleStackFqdnListId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:  "eXaMpLe-rEsOuRcE-GrOuP",
				LocalRuleStackName: "lOcAlRuLeStAcKvAlUe",
				FqdnListName:       "fQdNlIsTvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/lOcAlRuLeStAcKs/lOcAlRuLeStAcKvAlUe/fQdNlIsTs/fQdNlIsTvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLocalRuleStackFqdnListIDInsensitively(v.Input)
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

		if actual.LocalRuleStackName != v.Expected.LocalRuleStackName {
			t.Fatalf("Expected %q but got %q for LocalRuleStackName", v.Expected.LocalRuleStackName, actual.LocalRuleStackName)
		}

		if actual.FqdnListName != v.Expected.FqdnListName {
			t.Fatalf("Expected %q but got %q for FqdnListName", v.Expected.FqdnListName, actual.FqdnListName)
		}

	}
}

func TestSegmentsForLocalRuleStackFqdnListId(t *testing.T) {
	segments := LocalRuleStackFqdnListId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("LocalRuleStackFqdnListId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
