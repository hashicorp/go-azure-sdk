package policyauthenticationstrengthpolicycombinationconfiguration

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyAuthenticationStrengthPolicyCombinationConfigurationId{}

func TestNewPolicyAuthenticationStrengthPolicyCombinationConfigurationID(t *testing.T) {
	id := NewPolicyAuthenticationStrengthPolicyCombinationConfigurationID("authenticationStrengthPolicyIdValue", "authenticationCombinationConfigurationIdValue")

	if id.AuthenticationStrengthPolicyId != "authenticationStrengthPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationStrengthPolicyId'", id.AuthenticationStrengthPolicyId, "authenticationStrengthPolicyIdValue")
	}

	if id.AuthenticationCombinationConfigurationId != "authenticationCombinationConfigurationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationCombinationConfigurationId'", id.AuthenticationCombinationConfigurationId, "authenticationCombinationConfigurationIdValue")
	}
}

func TestFormatPolicyAuthenticationStrengthPolicyCombinationConfigurationID(t *testing.T) {
	actual := NewPolicyAuthenticationStrengthPolicyCombinationConfigurationID("authenticationStrengthPolicyIdValue", "authenticationCombinationConfigurationIdValue").ID()
	expected := "/policies/authenticationStrengthPolicies/authenticationStrengthPolicyIdValue/combinationConfigurations/authenticationCombinationConfigurationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyAuthenticationStrengthPolicyCombinationConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyAuthenticationStrengthPolicyCombinationConfigurationId
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
			// Incomplete URI
			Input: "/policies/authenticationStrengthPolicies/authenticationStrengthPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/authenticationStrengthPolicies/authenticationStrengthPolicyIdValue/combinationConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/authenticationStrengthPolicies/authenticationStrengthPolicyIdValue/combinationConfigurations/authenticationCombinationConfigurationIdValue",
			Expected: &PolicyAuthenticationStrengthPolicyCombinationConfigurationId{
				AuthenticationStrengthPolicyId:           "authenticationStrengthPolicyIdValue",
				AuthenticationCombinationConfigurationId: "authenticationCombinationConfigurationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/authenticationStrengthPolicies/authenticationStrengthPolicyIdValue/combinationConfigurations/authenticationCombinationConfigurationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyAuthenticationStrengthPolicyCombinationConfigurationID(v.Input)
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

		if actual.AuthenticationCombinationConfigurationId != v.Expected.AuthenticationCombinationConfigurationId {
			t.Fatalf("Expected %q but got %q for AuthenticationCombinationConfigurationId", v.Expected.AuthenticationCombinationConfigurationId, actual.AuthenticationCombinationConfigurationId)
		}

	}
}

func TestParsePolicyAuthenticationStrengthPolicyCombinationConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyAuthenticationStrengthPolicyCombinationConfigurationId
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
			// Incomplete URI
			Input: "/policies/authenticationStrengthPolicies/authenticationStrengthPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcYiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/authenticationStrengthPolicies/authenticationStrengthPolicyIdValue/combinationConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcYiDvAlUe/cOmBiNaTiOnCoNfIgUrAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/authenticationStrengthPolicies/authenticationStrengthPolicyIdValue/combinationConfigurations/authenticationCombinationConfigurationIdValue",
			Expected: &PolicyAuthenticationStrengthPolicyCombinationConfigurationId{
				AuthenticationStrengthPolicyId:           "authenticationStrengthPolicyIdValue",
				AuthenticationCombinationConfigurationId: "authenticationCombinationConfigurationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/authenticationStrengthPolicies/authenticationStrengthPolicyIdValue/combinationConfigurations/authenticationCombinationConfigurationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcYiDvAlUe/cOmBiNaTiOnCoNfIgUrAtIoNs/aUtHeNtIcAtIoNcOmBiNaTiOnCoNfIgUrAtIoNiDvAlUe",
			Expected: &PolicyAuthenticationStrengthPolicyCombinationConfigurationId{
				AuthenticationStrengthPolicyId:           "aUtHeNtIcAtIoNsTrEnGtHpOlIcYiDvAlUe",
				AuthenticationCombinationConfigurationId: "aUtHeNtIcAtIoNcOmBiNaTiOnCoNfIgUrAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcYiDvAlUe/cOmBiNaTiOnCoNfIgUrAtIoNs/aUtHeNtIcAtIoNcOmBiNaTiOnCoNfIgUrAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyAuthenticationStrengthPolicyCombinationConfigurationIDInsensitively(v.Input)
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

		if actual.AuthenticationCombinationConfigurationId != v.Expected.AuthenticationCombinationConfigurationId {
			t.Fatalf("Expected %q but got %q for AuthenticationCombinationConfigurationId", v.Expected.AuthenticationCombinationConfigurationId, actual.AuthenticationCombinationConfigurationId)
		}

	}
}

func TestSegmentsForPolicyAuthenticationStrengthPolicyCombinationConfigurationId(t *testing.T) {
	segments := PolicyAuthenticationStrengthPolicyCombinationConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyAuthenticationStrengthPolicyCombinationConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
