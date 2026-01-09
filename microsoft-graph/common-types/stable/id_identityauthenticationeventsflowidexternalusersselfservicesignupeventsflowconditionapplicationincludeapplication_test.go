package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId{}

func TestNewIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationID(t *testing.T) {
	id := NewIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationID("authenticationEventsFlowId", "authenticationConditionApplicationAppId")

	if id.AuthenticationEventsFlowId != "authenticationEventsFlowId" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationEventsFlowId'", id.AuthenticationEventsFlowId, "authenticationEventsFlowId")
	}

	if id.AuthenticationConditionApplicationAppId != "authenticationConditionApplicationAppId" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationConditionApplicationAppId'", id.AuthenticationConditionApplicationAppId, "authenticationConditionApplicationAppId")
	}
}

func TestFormatIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationID(t *testing.T) {
	actual := NewIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationID("authenticationEventsFlowId", "authenticationConditionApplicationAppId").ID()
	expected := "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/conditions/applications/includeApplications/authenticationConditionApplicationAppId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId
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
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/conditions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/conditions/applications",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/conditions/applications/includeApplications",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/conditions/applications/includeApplications/authenticationConditionApplicationAppId",
			Expected: &IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId{
				AuthenticationEventsFlowId:              "authenticationEventsFlowId",
				AuthenticationConditionApplicationAppId: "authenticationConditionApplicationAppId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/conditions/applications/includeApplications/authenticationConditionApplicationAppId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationID(v.Input)
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

		if actual.AuthenticationConditionApplicationAppId != v.Expected.AuthenticationConditionApplicationAppId {
			t.Fatalf("Expected %q but got %q for AuthenticationConditionApplicationAppId", v.Expected.AuthenticationConditionApplicationAppId, actual.AuthenticationConditionApplicationAppId)
		}

	}
}

func TestParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId
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
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/conditions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/eXtErNaLuSeRsSeLfSeRvIcEsIgNuPeVeNtSfLoW/cOnDiTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/conditions/applications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/eXtErNaLuSeRsSeLfSeRvIcEsIgNuPeVeNtSfLoW/cOnDiTiOnS/aPpLiCaTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/conditions/applications/includeApplications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/eXtErNaLuSeRsSeLfSeRvIcEsIgNuPeVeNtSfLoW/cOnDiTiOnS/aPpLiCaTiOnS/iNcLuDeApPlIcAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/conditions/applications/includeApplications/authenticationConditionApplicationAppId",
			Expected: &IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId{
				AuthenticationEventsFlowId:              "authenticationEventsFlowId",
				AuthenticationConditionApplicationAppId: "authenticationConditionApplicationAppId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/conditions/applications/includeApplications/authenticationConditionApplicationAppId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/eXtErNaLuSeRsSeLfSeRvIcEsIgNuPeVeNtSfLoW/cOnDiTiOnS/aPpLiCaTiOnS/iNcLuDeApPlIcAtIoNs/aUtHeNtIcAtIoNcOnDiTiOnApPlIcAtIoNaPpId",
			Expected: &IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId{
				AuthenticationEventsFlowId:              "aUtHeNtIcAtIoNeVeNtSfLoWiD",
				AuthenticationConditionApplicationAppId: "aUtHeNtIcAtIoNcOnDiTiOnApPlIcAtIoNaPpId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/eXtErNaLuSeRsSeLfSeRvIcEsIgNuPeVeNtSfLoW/cOnDiTiOnS/aPpLiCaTiOnS/iNcLuDeApPlIcAtIoNs/aUtHeNtIcAtIoNcOnDiTiOnApPlIcAtIoNaPpId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationIDInsensitively(v.Input)
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

		if actual.AuthenticationConditionApplicationAppId != v.Expected.AuthenticationConditionApplicationAppId {
			t.Fatalf("Expected %q but got %q for AuthenticationConditionApplicationAppId", v.Expected.AuthenticationConditionApplicationAppId, actual.AuthenticationConditionApplicationAppId)
		}

	}
}

func TestSegmentsForIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId(t *testing.T) {
	segments := IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
