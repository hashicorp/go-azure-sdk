package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyAuthenticationStrengthPolicyId{}

func TestNewPolicyAuthenticationStrengthPolicyID(t *testing.T) {
	id := NewPolicyAuthenticationStrengthPolicyID("authenticationStrengthPolicyId")

	if id.AuthenticationStrengthPolicyId != "authenticationStrengthPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationStrengthPolicyId'", id.AuthenticationStrengthPolicyId, "authenticationStrengthPolicyId")
	}
}

func TestFormatPolicyAuthenticationStrengthPolicyID(t *testing.T) {
	actual := NewPolicyAuthenticationStrengthPolicyID("authenticationStrengthPolicyId").ID()
	expected := "/policies/authenticationStrengthPolicies/authenticationStrengthPolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyAuthenticationStrengthPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyAuthenticationStrengthPolicyId
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
			Input: "/policies/authenticationStrengthPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/authenticationStrengthPolicies/authenticationStrengthPolicyId",
			Expected: &PolicyAuthenticationStrengthPolicyId{
				AuthenticationStrengthPolicyId: "authenticationStrengthPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/authenticationStrengthPolicies/authenticationStrengthPolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyAuthenticationStrengthPolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuthenticationStrengthPolicyId != v.Expected.AuthenticationStrengthPolicyId {
			t.Fatalf("Expected %q but got %q for AuthenticationStrengthPolicyId", v.Expected.AuthenticationStrengthPolicyId, actual.AuthenticationStrengthPolicyId)
		}

	}
}

func TestParsePolicyAuthenticationStrengthPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyAuthenticationStrengthPolicyId
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
			Input: "/policies/authenticationStrengthPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/authenticationStrengthPolicies/authenticationStrengthPolicyId",
			Expected: &PolicyAuthenticationStrengthPolicyId{
				AuthenticationStrengthPolicyId: "authenticationStrengthPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/authenticationStrengthPolicies/authenticationStrengthPolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcYiD",
			Expected: &PolicyAuthenticationStrengthPolicyId{
				AuthenticationStrengthPolicyId: "aUtHeNtIcAtIoNsTrEnGtHpOlIcYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyAuthenticationStrengthPolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuthenticationStrengthPolicyId != v.Expected.AuthenticationStrengthPolicyId {
			t.Fatalf("Expected %q but got %q for AuthenticationStrengthPolicyId", v.Expected.AuthenticationStrengthPolicyId, actual.AuthenticationStrengthPolicyId)
		}

	}
}

func TestSegmentsForPolicyAuthenticationStrengthPolicyId(t *testing.T) {
	segments := PolicyAuthenticationStrengthPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyAuthenticationStrengthPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
