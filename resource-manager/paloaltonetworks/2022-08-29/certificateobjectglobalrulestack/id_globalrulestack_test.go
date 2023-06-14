package certificateobjectglobalrulestack

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GlobalRuleStackId{}

func TestNewGlobalRuleStackID(t *testing.T) {
	id := NewGlobalRuleStackID("globalRuleStackValue")

	if id.GlobalRuleStackName != "globalRuleStackValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GlobalRuleStackName'", id.GlobalRuleStackName, "globalRuleStackValue")
	}
}

func TestFormatGlobalRuleStackID(t *testing.T) {
	actual := NewGlobalRuleStackID("globalRuleStackValue").ID()
	expected := "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGlobalRuleStackID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GlobalRuleStackId
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
			// Valid URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue",
			Expected: &GlobalRuleStackId{
				GlobalRuleStackName: "globalRuleStackValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGlobalRuleStackID(v.Input)
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

	}
}

func TestParseGlobalRuleStackIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GlobalRuleStackId
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
			// Valid URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue",
			Expected: &GlobalRuleStackId{
				GlobalRuleStackName: "globalRuleStackValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE",
			Expected: &GlobalRuleStackId{
				GlobalRuleStackName: "gLoBaLrUlEsTaCkVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGlobalRuleStackIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForGlobalRuleStackId(t *testing.T) {
	segments := GlobalRuleStackId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GlobalRuleStackId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
