package fqdnlistglobalrulestack

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = FqdnListId{}

func TestNewFqdnListID(t *testing.T) {
	id := NewFqdnListID("globalRuleStackValue", "fqdnListValue")

	if id.GlobalRuleStackName != "globalRuleStackValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GlobalRuleStackName'", id.GlobalRuleStackName, "globalRuleStackValue")
	}

	if id.FqdnListName != "fqdnListValue" {
		t.Fatalf("Expected %q but got %q for Segment 'FqdnListName'", id.FqdnListName, "fqdnListValue")
	}
}

func TestFormatFqdnListID(t *testing.T) {
	actual := NewFqdnListID("globalRuleStackValue", "fqdnListValue").ID()
	expected := "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/fqdnLists/fqdnListValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseFqdnListID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *FqdnListId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/fqdnLists",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/fqdnLists/fqdnListValue",
			Expected: &FqdnListId{
				GlobalRuleStackName: "globalRuleStackValue",
				FqdnListName:        "fqdnListValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/fqdnLists/fqdnListValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseFqdnListID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GlobalRuleStackName != v.Expected.GlobalRuleStackName {
			t.Fatalf("Expected %q but got %q for GlobalRuleStackName", v.Expected.GlobalRuleStackName, actual.GlobalRuleStackName)
		}

		if actual.FqdnListName != v.Expected.FqdnListName {
			t.Fatalf("Expected %q but got %q for FqdnListName", v.Expected.FqdnListName, actual.FqdnListName)
		}

	}
}

func TestParseFqdnListIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *FqdnListId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/fqdnLists",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE/fQdNlIsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/fqdnLists/fqdnListValue",
			Expected: &FqdnListId{
				GlobalRuleStackName: "globalRuleStackValue",
				FqdnListName:        "fqdnListValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/fqdnLists/fqdnListValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE/fQdNlIsTs/fQdNlIsTvAlUe",
			Expected: &FqdnListId{
				GlobalRuleStackName: "gLoBaLrUlEsTaCkVaLuE",
				FqdnListName:        "fQdNlIsTvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE/fQdNlIsTs/fQdNlIsTvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseFqdnListIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GlobalRuleStackName != v.Expected.GlobalRuleStackName {
			t.Fatalf("Expected %q but got %q for GlobalRuleStackName", v.Expected.GlobalRuleStackName, actual.GlobalRuleStackName)
		}

		if actual.FqdnListName != v.Expected.FqdnListName {
			t.Fatalf("Expected %q but got %q for FqdnListName", v.Expected.FqdnListName, actual.FqdnListName)
		}

	}
}

func TestSegmentsForFqdnListId(t *testing.T) {
	segments := FqdnListId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("FqdnListId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
