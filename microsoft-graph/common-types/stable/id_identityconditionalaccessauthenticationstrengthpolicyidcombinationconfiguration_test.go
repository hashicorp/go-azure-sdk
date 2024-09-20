package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId{}

func TestNewIdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationID(t *testing.T) {
	id := NewIdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationID("authenticationStrengthPolicyId", "authenticationCombinationConfigurationId")

	if id.AuthenticationStrengthPolicyId != "authenticationStrengthPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationStrengthPolicyId'", id.AuthenticationStrengthPolicyId, "authenticationStrengthPolicyId")
	}

	if id.AuthenticationCombinationConfigurationId != "authenticationCombinationConfigurationId" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationCombinationConfigurationId'", id.AuthenticationCombinationConfigurationId, "authenticationCombinationConfigurationId")
	}
}

func TestFormatIdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationID(t *testing.T) {
	actual := NewIdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationID("authenticationStrengthPolicyId", "authenticationCombinationConfigurationId").ID()
	expected := "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyId/combinationConfigurations/authenticationCombinationConfigurationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId
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
			Input: "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyId/combinationConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyId/combinationConfigurations/authenticationCombinationConfigurationId",
			Expected: &IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId{
				AuthenticationStrengthPolicyId:           "authenticationStrengthPolicyId",
				AuthenticationCombinationConfigurationId: "authenticationCombinationConfigurationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyId/combinationConfigurations/authenticationCombinationConfigurationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationID(v.Input)
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

func TestParseIdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId
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
			Input: "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/aUtHeNtIcAtIoNsTrEnGtH/pOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcYiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyId/combinationConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/aUtHeNtIcAtIoNsTrEnGtH/pOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcYiD/cOmBiNaTiOnCoNfIgUrAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyId/combinationConfigurations/authenticationCombinationConfigurationId",
			Expected: &IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId{
				AuthenticationStrengthPolicyId:           "authenticationStrengthPolicyId",
				AuthenticationCombinationConfigurationId: "authenticationCombinationConfigurationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyId/combinationConfigurations/authenticationCombinationConfigurationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/aUtHeNtIcAtIoNsTrEnGtH/pOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcYiD/cOmBiNaTiOnCoNfIgUrAtIoNs/aUtHeNtIcAtIoNcOmBiNaTiOnCoNfIgUrAtIoNiD",
			Expected: &IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId{
				AuthenticationStrengthPolicyId:           "aUtHeNtIcAtIoNsTrEnGtHpOlIcYiD",
				AuthenticationCombinationConfigurationId: "aUtHeNtIcAtIoNcOmBiNaTiOnCoNfIgUrAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/aUtHeNtIcAtIoNsTrEnGtH/pOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcYiD/cOmBiNaTiOnCoNfIgUrAtIoNs/aUtHeNtIcAtIoNcOmBiNaTiOnCoNfIgUrAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationIDInsensitively(v.Input)
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

func TestSegmentsForIdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId(t *testing.T) {
	segments := IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
