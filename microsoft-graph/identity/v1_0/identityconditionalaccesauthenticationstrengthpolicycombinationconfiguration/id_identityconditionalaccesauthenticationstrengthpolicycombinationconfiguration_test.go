package identityconditionalaccesauthenticationstrengthpolicycombinationconfiguration

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId{}

func TestNewIdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationID(t *testing.T) {
	id := NewIdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationID("authenticationStrengthPolicyIdValue", "authenticationCombinationConfigurationIdValue")

	if id.AuthenticationStrengthPolicyId != "authenticationStrengthPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationStrengthPolicyId'", id.AuthenticationStrengthPolicyId, "authenticationStrengthPolicyIdValue")
	}

	if id.AuthenticationCombinationConfigurationId != "authenticationCombinationConfigurationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationCombinationConfigurationId'", id.AuthenticationCombinationConfigurationId, "authenticationCombinationConfigurationIdValue")
	}
}

func TestFormatIdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationID(t *testing.T) {
	actual := NewIdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationID("authenticationStrengthPolicyIdValue", "authenticationCombinationConfigurationIdValue").ID()
	expected := "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyIdValue/combinationConfigurations/authenticationCombinationConfigurationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess/authenticationStrength",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess/authenticationStrength/policies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyIdValue/combinationConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyIdValue/combinationConfigurations/authenticationCombinationConfigurationIdValue",
			Expected: &IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId{
				AuthenticationStrengthPolicyId:           "authenticationStrengthPolicyIdValue",
				AuthenticationCombinationConfigurationId: "authenticationCombinationConfigurationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyIdValue/combinationConfigurations/authenticationCombinationConfigurationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationID(v.Input)
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

func TestParseIdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess/authenticationStrength",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/aUtHeNtIcAtIoNsTrEnGtH",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess/authenticationStrength/policies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/aUtHeNtIcAtIoNsTrEnGtH/pOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/aUtHeNtIcAtIoNsTrEnGtH/pOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcYiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyIdValue/combinationConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/aUtHeNtIcAtIoNsTrEnGtH/pOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcYiDvAlUe/cOmBiNaTiOnCoNfIgUrAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyIdValue/combinationConfigurations/authenticationCombinationConfigurationIdValue",
			Expected: &IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId{
				AuthenticationStrengthPolicyId:           "authenticationStrengthPolicyIdValue",
				AuthenticationCombinationConfigurationId: "authenticationCombinationConfigurationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyIdValue/combinationConfigurations/authenticationCombinationConfigurationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/aUtHeNtIcAtIoNsTrEnGtH/pOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcYiDvAlUe/cOmBiNaTiOnCoNfIgUrAtIoNs/aUtHeNtIcAtIoNcOmBiNaTiOnCoNfIgUrAtIoNiDvAlUe",
			Expected: &IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId{
				AuthenticationStrengthPolicyId:           "aUtHeNtIcAtIoNsTrEnGtHpOlIcYiDvAlUe",
				AuthenticationCombinationConfigurationId: "aUtHeNtIcAtIoNcOmBiNaTiOnCoNfIgUrAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/aUtHeNtIcAtIoNsTrEnGtH/pOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcYiDvAlUe/cOmBiNaTiOnCoNfIgUrAtIoNs/aUtHeNtIcAtIoNcOmBiNaTiOnCoNfIgUrAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationIDInsensitively(v.Input)
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

func TestSegmentsForIdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId(t *testing.T) {
	segments := IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
