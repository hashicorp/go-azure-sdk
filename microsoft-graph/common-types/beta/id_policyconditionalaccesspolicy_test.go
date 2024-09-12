package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyConditionalAccessPolicyId{}

func TestNewPolicyConditionalAccessPolicyID(t *testing.T) {
	id := NewPolicyConditionalAccessPolicyID("conditionalAccessPolicyIdValue")

	if id.ConditionalAccessPolicyId != "conditionalAccessPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ConditionalAccessPolicyId'", id.ConditionalAccessPolicyId, "conditionalAccessPolicyIdValue")
	}
}

func TestFormatPolicyConditionalAccessPolicyID(t *testing.T) {
	actual := NewPolicyConditionalAccessPolicyID("conditionalAccessPolicyIdValue").ID()
	expected := "/policies/conditionalAccessPolicies/conditionalAccessPolicyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyConditionalAccessPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyConditionalAccessPolicyId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/conditionalAccessPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/conditionalAccessPolicies/conditionalAccessPolicyIdValue",
			Expected: &PolicyConditionalAccessPolicyId{
				ConditionalAccessPolicyId: "conditionalAccessPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/conditionalAccessPolicies/conditionalAccessPolicyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyConditionalAccessPolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ConditionalAccessPolicyId != v.Expected.ConditionalAccessPolicyId {
			t.Fatalf("Expected %q but got %q for ConditionalAccessPolicyId", v.Expected.ConditionalAccessPolicyId, actual.ConditionalAccessPolicyId)
		}

	}
}

func TestParsePolicyConditionalAccessPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyConditionalAccessPolicyId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/conditionalAccessPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/cOnDiTiOnAlAcCeSsPoLiCiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/conditionalAccessPolicies/conditionalAccessPolicyIdValue",
			Expected: &PolicyConditionalAccessPolicyId{
				ConditionalAccessPolicyId: "conditionalAccessPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/conditionalAccessPolicies/conditionalAccessPolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/cOnDiTiOnAlAcCeSsPoLiCiEs/cOnDiTiOnAlAcCeSsPoLiCyIdVaLuE",
			Expected: &PolicyConditionalAccessPolicyId{
				ConditionalAccessPolicyId: "cOnDiTiOnAlAcCeSsPoLiCyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/cOnDiTiOnAlAcCeSsPoLiCiEs/cOnDiTiOnAlAcCeSsPoLiCyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyConditionalAccessPolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ConditionalAccessPolicyId != v.Expected.ConditionalAccessPolicyId {
			t.Fatalf("Expected %q but got %q for ConditionalAccessPolicyId", v.Expected.ConditionalAccessPolicyId, actual.ConditionalAccessPolicyId)
		}

	}
}

func TestSegmentsForPolicyConditionalAccessPolicyId(t *testing.T) {
	segments := PolicyConditionalAccessPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyConditionalAccessPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
