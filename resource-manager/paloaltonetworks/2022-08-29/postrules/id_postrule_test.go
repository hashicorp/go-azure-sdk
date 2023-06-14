package postrules

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PostRuleId{}

func TestNewPostRuleID(t *testing.T) {
	id := NewPostRuleID("globalRuleStackValue", "postRuleValue")

	if id.GlobalRuleStackName != "globalRuleStackValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GlobalRuleStackName'", id.GlobalRuleStackName, "globalRuleStackValue")
	}

	if id.PostRuleName != "postRuleValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PostRuleName'", id.PostRuleName, "postRuleValue")
	}
}

func TestFormatPostRuleID(t *testing.T) {
	actual := NewPostRuleID("globalRuleStackValue", "postRuleValue").ID()
	expected := "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/postRules/postRuleValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePostRuleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PostRuleId
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
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/postRules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/postRules/postRuleValue",
			Expected: &PostRuleId{
				GlobalRuleStackName: "globalRuleStackValue",
				PostRuleName:        "postRuleValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/postRules/postRuleValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePostRuleID(v.Input)
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

		if actual.PostRuleName != v.Expected.PostRuleName {
			t.Fatalf("Expected %q but got %q for PostRuleName", v.Expected.PostRuleName, actual.PostRuleName)
		}

	}
}

func TestParsePostRuleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PostRuleId
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
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/postRules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE/pOsTrUlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/postRules/postRuleValue",
			Expected: &PostRuleId{
				GlobalRuleStackName: "globalRuleStackValue",
				PostRuleName:        "postRuleValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/postRules/postRuleValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE/pOsTrUlEs/pOsTrUlEvAlUe",
			Expected: &PostRuleId{
				GlobalRuleStackName: "gLoBaLrUlEsTaCkVaLuE",
				PostRuleName:        "pOsTrUlEvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE/pOsTrUlEs/pOsTrUlEvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePostRuleIDInsensitively(v.Input)
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

		if actual.PostRuleName != v.Expected.PostRuleName {
			t.Fatalf("Expected %q but got %q for PostRuleName", v.Expected.PostRuleName, actual.PostRuleName)
		}

	}
}

func TestSegmentsForPostRuleId(t *testing.T) {
	segments := PostRuleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PostRuleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
