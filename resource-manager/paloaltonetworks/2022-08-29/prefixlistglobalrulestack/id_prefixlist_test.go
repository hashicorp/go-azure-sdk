package prefixlistglobalrulestack

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PrefixListId{}

func TestNewPrefixListID(t *testing.T) {
	id := NewPrefixListID("globalRulestackValue", "prefixListValue")

	if id.GlobalRulestackName != "globalRulestackValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GlobalRulestackName'", id.GlobalRulestackName, "globalRulestackValue")
	}

	if id.PrefixListName != "prefixListValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PrefixListName'", id.PrefixListName, "prefixListValue")
	}
}

func TestFormatPrefixListID(t *testing.T) {
	actual := NewPrefixListID("globalRulestackValue", "prefixListValue").ID()
	expected := "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/globalRulestackValue/prefixLists/prefixListValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePrefixListID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PrefixListId
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
			Input: "/providers/PaloAltoNetworks.Cloudngfw",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/globalRulestackValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/globalRulestackValue/prefixLists",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/globalRulestackValue/prefixLists/prefixListValue",
			Expected: &PrefixListId{
				GlobalRulestackName: "globalRulestackValue",
				PrefixListName:      "prefixListValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/globalRulestackValue/prefixLists/prefixListValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePrefixListID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GlobalRulestackName != v.Expected.GlobalRulestackName {
			t.Fatalf("Expected %q but got %q for GlobalRulestackName", v.Expected.GlobalRulestackName, actual.GlobalRulestackName)
		}

		if actual.PrefixListName != v.Expected.PrefixListName {
			t.Fatalf("Expected %q but got %q for PrefixListName", v.Expected.PrefixListName, actual.PrefixListName)
		}

	}
}

func TestParsePrefixListIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PrefixListId
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
			Input: "/providers/PaloAltoNetworks.Cloudngfw",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/globalRulestackValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/globalRulestackValue/prefixLists",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE/pReFiXlIsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/globalRulestackValue/prefixLists/prefixListValue",
			Expected: &PrefixListId{
				GlobalRulestackName: "globalRulestackValue",
				PrefixListName:      "prefixListValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/globalRulestackValue/prefixLists/prefixListValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE/pReFiXlIsTs/pReFiXlIsTvAlUe",
			Expected: &PrefixListId{
				GlobalRulestackName: "gLoBaLrUlEsTaCkVaLuE",
				PrefixListName:      "pReFiXlIsTvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE/pReFiXlIsTs/pReFiXlIsTvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePrefixListIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GlobalRulestackName != v.Expected.GlobalRulestackName {
			t.Fatalf("Expected %q but got %q for GlobalRulestackName", v.Expected.GlobalRulestackName, actual.GlobalRulestackName)
		}

		if actual.PrefixListName != v.Expected.PrefixListName {
			t.Fatalf("Expected %q but got %q for PrefixListName", v.Expected.PrefixListName, actual.PrefixListName)
		}

	}
}

func TestSegmentsForPrefixListId(t *testing.T) {
	segments := PrefixListId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PrefixListId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
