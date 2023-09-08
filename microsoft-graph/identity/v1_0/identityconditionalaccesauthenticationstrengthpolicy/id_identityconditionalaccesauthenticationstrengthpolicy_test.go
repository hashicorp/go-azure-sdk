package identityconditionalaccesauthenticationstrengthpolicy

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityConditionalAccesAuthenticationStrengthPolicyId{}

func TestNewIdentityConditionalAccesAuthenticationStrengthPolicyID(t *testing.T) {
	id := NewIdentityConditionalAccesAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue")

	if id.AuthenticationStrengthPolicyId != "authenticationStrengthPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationStrengthPolicyId'", id.AuthenticationStrengthPolicyId, "authenticationStrengthPolicyIdValue")
	}
}

func TestFormatIdentityConditionalAccesAuthenticationStrengthPolicyID(t *testing.T) {
	actual := NewIdentityConditionalAccesAuthenticationStrengthPolicyID("authenticationStrengthPolicyIdValue").ID()
	expected := "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityConditionalAccesAuthenticationStrengthPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccesAuthenticationStrengthPolicyId
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
			// Valid URI
			Input: "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyIdValue",
			Expected: &IdentityConditionalAccesAuthenticationStrengthPolicyId{
				AuthenticationStrengthPolicyId: "authenticationStrengthPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccesAuthenticationStrengthPolicyID(v.Input)
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

func TestParseIdentityConditionalAccesAuthenticationStrengthPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccesAuthenticationStrengthPolicyId
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
			// Valid URI
			Input: "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyIdValue",
			Expected: &IdentityConditionalAccesAuthenticationStrengthPolicyId{
				AuthenticationStrengthPolicyId: "authenticationStrengthPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/authenticationStrength/policies/authenticationStrengthPolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/aUtHeNtIcAtIoNsTrEnGtH/pOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcYiDvAlUe",
			Expected: &IdentityConditionalAccesAuthenticationStrengthPolicyId{
				AuthenticationStrengthPolicyId: "aUtHeNtIcAtIoNsTrEnGtHpOlIcYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/aUtHeNtIcAtIoNsTrEnGtH/pOlIcIeS/aUtHeNtIcAtIoNsTrEnGtHpOlIcYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccesAuthenticationStrengthPolicyIDInsensitively(v.Input)
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

func TestSegmentsForIdentityConditionalAccesAuthenticationStrengthPolicyId(t *testing.T) {
	segments := IdentityConditionalAccesAuthenticationStrengthPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityConditionalAccesAuthenticationStrengthPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
