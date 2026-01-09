package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId{}

func TestNewPolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationID(t *testing.T) {
	id := NewPolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationID("authenticationMethodConfigurationId")

	if id.AuthenticationMethodConfigurationId != "authenticationMethodConfigurationId" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationMethodConfigurationId'", id.AuthenticationMethodConfigurationId, "authenticationMethodConfigurationId")
	}
}

func TestFormatPolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationID(t *testing.T) {
	actual := NewPolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationID("authenticationMethodConfigurationId").ID()
	expected := "/policies/authenticationMethodsPolicy/authenticationMethodConfigurations/authenticationMethodConfigurationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId
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
			Input: "/policies/authenticationMethodsPolicy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/authenticationMethodsPolicy/authenticationMethodConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/authenticationMethodsPolicy/authenticationMethodConfigurations/authenticationMethodConfigurationId",
			Expected: &PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId{
				AuthenticationMethodConfigurationId: "authenticationMethodConfigurationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/authenticationMethodsPolicy/authenticationMethodConfigurations/authenticationMethodConfigurationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuthenticationMethodConfigurationId != v.Expected.AuthenticationMethodConfigurationId {
			t.Fatalf("Expected %q but got %q for AuthenticationMethodConfigurationId", v.Expected.AuthenticationMethodConfigurationId, actual.AuthenticationMethodConfigurationId)
		}

	}
}

func TestParsePolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId
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
			Input: "/policies/authenticationMethodsPolicy",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aUtHeNtIcAtIoNmEtHoDsPoLiCy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/authenticationMethodsPolicy/authenticationMethodConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aUtHeNtIcAtIoNmEtHoDsPoLiCy/aUtHeNtIcAtIoNmEtHoDcOnFiGuRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/authenticationMethodsPolicy/authenticationMethodConfigurations/authenticationMethodConfigurationId",
			Expected: &PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId{
				AuthenticationMethodConfigurationId: "authenticationMethodConfigurationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/authenticationMethodsPolicy/authenticationMethodConfigurations/authenticationMethodConfigurationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aUtHeNtIcAtIoNmEtHoDsPoLiCy/aUtHeNtIcAtIoNmEtHoDcOnFiGuRaTiOnS/aUtHeNtIcAtIoNmEtHoDcOnFiGuRaTiOnId",
			Expected: &PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId{
				AuthenticationMethodConfigurationId: "aUtHeNtIcAtIoNmEtHoDcOnFiGuRaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aUtHeNtIcAtIoNmEtHoDsPoLiCy/aUtHeNtIcAtIoNmEtHoDcOnFiGuRaTiOnS/aUtHeNtIcAtIoNmEtHoDcOnFiGuRaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuthenticationMethodConfigurationId != v.Expected.AuthenticationMethodConfigurationId {
			t.Fatalf("Expected %q but got %q for AuthenticationMethodConfigurationId", v.Expected.AuthenticationMethodConfigurationId, actual.AuthenticationMethodConfigurationId)
		}

	}
}

func TestSegmentsForPolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId(t *testing.T) {
	segments := PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
