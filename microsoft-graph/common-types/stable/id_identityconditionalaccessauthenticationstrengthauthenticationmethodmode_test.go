package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId{}

func TestNewIdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeID(t *testing.T) {
	id := NewIdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeID("authenticationMethodModeDetailIdValue")

	if id.AuthenticationMethodModeDetailId != "authenticationMethodModeDetailIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationMethodModeDetailId'", id.AuthenticationMethodModeDetailId, "authenticationMethodModeDetailIdValue")
	}
}

func TestFormatIdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeID(t *testing.T) {
	actual := NewIdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeID("authenticationMethodModeDetailIdValue").ID()
	expected := "/identity/conditionalAccess/authenticationStrength/authenticationMethodModes/authenticationMethodModeDetailIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId
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
			Input: "/identity/conditionalAccess/authenticationStrength/authenticationMethodModes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/conditionalAccess/authenticationStrength/authenticationMethodModes/authenticationMethodModeDetailIdValue",
			Expected: &IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId{
				AuthenticationMethodModeDetailId: "authenticationMethodModeDetailIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/authenticationStrength/authenticationMethodModes/authenticationMethodModeDetailIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuthenticationMethodModeDetailId != v.Expected.AuthenticationMethodModeDetailId {
			t.Fatalf("Expected %q but got %q for AuthenticationMethodModeDetailId", v.Expected.AuthenticationMethodModeDetailId, actual.AuthenticationMethodModeDetailId)
		}

	}
}

func TestParseIdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId
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
			Input: "/identity/conditionalAccess/authenticationStrength/authenticationMethodModes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/aUtHeNtIcAtIoNsTrEnGtH/aUtHeNtIcAtIoNmEtHoDmOdEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/conditionalAccess/authenticationStrength/authenticationMethodModes/authenticationMethodModeDetailIdValue",
			Expected: &IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId{
				AuthenticationMethodModeDetailId: "authenticationMethodModeDetailIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/authenticationStrength/authenticationMethodModes/authenticationMethodModeDetailIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/aUtHeNtIcAtIoNsTrEnGtH/aUtHeNtIcAtIoNmEtHoDmOdEs/aUtHeNtIcAtIoNmEtHoDmOdEdEtAiLiDvAlUe",
			Expected: &IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId{
				AuthenticationMethodModeDetailId: "aUtHeNtIcAtIoNmEtHoDmOdEdEtAiLiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/aUtHeNtIcAtIoNsTrEnGtH/aUtHeNtIcAtIoNmEtHoDmOdEs/aUtHeNtIcAtIoNmEtHoDmOdEdEtAiLiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuthenticationMethodModeDetailId != v.Expected.AuthenticationMethodModeDetailId {
			t.Fatalf("Expected %q but got %q for AuthenticationMethodModeDetailId", v.Expected.AuthenticationMethodModeDetailId, actual.AuthenticationMethodModeDetailId)
		}

	}
}

func TestSegmentsForIdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId(t *testing.T) {
	segments := IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
