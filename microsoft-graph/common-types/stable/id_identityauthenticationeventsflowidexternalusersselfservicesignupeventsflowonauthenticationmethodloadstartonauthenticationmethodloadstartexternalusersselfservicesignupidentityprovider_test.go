package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId{}

func TestNewIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderID(t *testing.T) {
	id := NewIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderID("authenticationEventsFlowId", "identityProviderBaseId")

	if id.AuthenticationEventsFlowId != "authenticationEventsFlowId" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationEventsFlowId'", id.AuthenticationEventsFlowId, "authenticationEventsFlowId")
	}

	if id.IdentityProviderBaseId != "identityProviderBaseId" {
		t.Fatalf("Expected %q but got %q for Segment 'IdentityProviderBaseId'", id.IdentityProviderBaseId, "identityProviderBaseId")
	}
}

func TestFormatIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderID(t *testing.T) {
	actual := NewIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderID("authenticationEventsFlowId", "identityProviderBaseId").ID()
	expected := "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart/onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp/identityProviders/identityProviderBaseId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId
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
			Input: "/identity/authenticationEventsFlows",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart/onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart/onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp/identityProviders",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart/onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp/identityProviders/identityProviderBaseId",
			Expected: &IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId{
				AuthenticationEventsFlowId: "authenticationEventsFlowId",
				IdentityProviderBaseId:     "identityProviderBaseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart/onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp/identityProviders/identityProviderBaseId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuthenticationEventsFlowId != v.Expected.AuthenticationEventsFlowId {
			t.Fatalf("Expected %q but got %q for AuthenticationEventsFlowId", v.Expected.AuthenticationEventsFlowId, actual.AuthenticationEventsFlowId)
		}

		if actual.IdentityProviderBaseId != v.Expected.IdentityProviderBaseId {
			t.Fatalf("Expected %q but got %q for IdentityProviderBaseId", v.Expected.IdentityProviderBaseId, actual.IdentityProviderBaseId)
		}

	}
}

func TestParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId
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
			Input: "/identity/authenticationEventsFlows",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/eXtErNaLuSeRsSeLfSeRvIcEsIgNuPeVeNtSfLoW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/eXtErNaLuSeRsSeLfSeRvIcEsIgNuPeVeNtSfLoW/oNaUtHeNtIcAtIoNmEtHoDlOaDsTaRt",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart/onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/eXtErNaLuSeRsSeLfSeRvIcEsIgNuPeVeNtSfLoW/oNaUtHeNtIcAtIoNmEtHoDlOaDsTaRt/oNaUtHeNtIcAtIoNmEtHoDlOaDsTaRtExTeRnAlUsErSsElFsErViCeSiGnUp",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart/onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp/identityProviders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/eXtErNaLuSeRsSeLfSeRvIcEsIgNuPeVeNtSfLoW/oNaUtHeNtIcAtIoNmEtHoDlOaDsTaRt/oNaUtHeNtIcAtIoNmEtHoDlOaDsTaRtExTeRnAlUsErSsElFsErViCeSiGnUp/iDeNtItYpRoViDeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart/onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp/identityProviders/identityProviderBaseId",
			Expected: &IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId{
				AuthenticationEventsFlowId: "authenticationEventsFlowId",
				IdentityProviderBaseId:     "identityProviderBaseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart/onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp/identityProviders/identityProviderBaseId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/eXtErNaLuSeRsSeLfSeRvIcEsIgNuPeVeNtSfLoW/oNaUtHeNtIcAtIoNmEtHoDlOaDsTaRt/oNaUtHeNtIcAtIoNmEtHoDlOaDsTaRtExTeRnAlUsErSsElFsErViCeSiGnUp/iDeNtItYpRoViDeRs/iDeNtItYpRoViDeRbAsEiD",
			Expected: &IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId{
				AuthenticationEventsFlowId: "aUtHeNtIcAtIoNeVeNtSfLoWiD",
				IdentityProviderBaseId:     "iDeNtItYpRoViDeRbAsEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/eXtErNaLuSeRsSeLfSeRvIcEsIgNuPeVeNtSfLoW/oNaUtHeNtIcAtIoNmEtHoDlOaDsTaRt/oNaUtHeNtIcAtIoNmEtHoDlOaDsTaRtExTeRnAlUsErSsElFsErViCeSiGnUp/iDeNtItYpRoViDeRs/iDeNtItYpRoViDeRbAsEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuthenticationEventsFlowId != v.Expected.AuthenticationEventsFlowId {
			t.Fatalf("Expected %q but got %q for AuthenticationEventsFlowId", v.Expected.AuthenticationEventsFlowId, actual.AuthenticationEventsFlowId)
		}

		if actual.IdentityProviderBaseId != v.Expected.IdentityProviderBaseId {
			t.Fatalf("Expected %q but got %q for IdentityProviderBaseId", v.Expected.IdentityProviderBaseId, actual.IdentityProviderBaseId)
		}

	}
}

func TestSegmentsForIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId(t *testing.T) {
	segments := IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
