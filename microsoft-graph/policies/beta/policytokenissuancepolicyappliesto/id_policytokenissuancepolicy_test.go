package policytokenissuancepolicyappliesto

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyTokenIssuancePolicyId{}

func TestNewPolicyTokenIssuancePolicyID(t *testing.T) {
	id := NewPolicyTokenIssuancePolicyID("tokenIssuancePolicyIdValue")

	if id.TokenIssuancePolicyId != "tokenIssuancePolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TokenIssuancePolicyId'", id.TokenIssuancePolicyId, "tokenIssuancePolicyIdValue")
	}
}

func TestFormatPolicyTokenIssuancePolicyID(t *testing.T) {
	actual := NewPolicyTokenIssuancePolicyID("tokenIssuancePolicyIdValue").ID()
	expected := "/policies/tokenIssuancePolicies/tokenIssuancePolicyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyTokenIssuancePolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyTokenIssuancePolicyId
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
			Input: "/policies/tokenIssuancePolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/tokenIssuancePolicies/tokenIssuancePolicyIdValue",
			Expected: &PolicyTokenIssuancePolicyId{
				TokenIssuancePolicyId: "tokenIssuancePolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/tokenIssuancePolicies/tokenIssuancePolicyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyTokenIssuancePolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TokenIssuancePolicyId != v.Expected.TokenIssuancePolicyId {
			t.Fatalf("Expected %q but got %q for TokenIssuancePolicyId", v.Expected.TokenIssuancePolicyId, actual.TokenIssuancePolicyId)
		}

	}
}

func TestParsePolicyTokenIssuancePolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyTokenIssuancePolicyId
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
			Input: "/policies/tokenIssuancePolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/tOkEnIsSuAnCePoLiCiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/tokenIssuancePolicies/tokenIssuancePolicyIdValue",
			Expected: &PolicyTokenIssuancePolicyId{
				TokenIssuancePolicyId: "tokenIssuancePolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/tokenIssuancePolicies/tokenIssuancePolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/tOkEnIsSuAnCePoLiCiEs/tOkEnIsSuAnCePoLiCyIdVaLuE",
			Expected: &PolicyTokenIssuancePolicyId{
				TokenIssuancePolicyId: "tOkEnIsSuAnCePoLiCyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/tOkEnIsSuAnCePoLiCiEs/tOkEnIsSuAnCePoLiCyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyTokenIssuancePolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TokenIssuancePolicyId != v.Expected.TokenIssuancePolicyId {
			t.Fatalf("Expected %q but got %q for TokenIssuancePolicyId", v.Expected.TokenIssuancePolicyId, actual.TokenIssuancePolicyId)
		}

	}
}

func TestSegmentsForPolicyTokenIssuancePolicyId(t *testing.T) {
	segments := PolicyTokenIssuancePolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyTokenIssuancePolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
