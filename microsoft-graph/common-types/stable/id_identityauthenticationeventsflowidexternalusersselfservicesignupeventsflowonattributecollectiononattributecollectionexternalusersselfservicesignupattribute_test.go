package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId{}

func TestNewIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeID(t *testing.T) {
	id := NewIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeID("authenticationEventsFlowId", "identityUserFlowAttributeId")

	if id.AuthenticationEventsFlowId != "authenticationEventsFlowId" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationEventsFlowId'", id.AuthenticationEventsFlowId, "authenticationEventsFlowId")
	}

	if id.IdentityUserFlowAttributeId != "identityUserFlowAttributeId" {
		t.Fatalf("Expected %q but got %q for Segment 'IdentityUserFlowAttributeId'", id.IdentityUserFlowAttributeId, "identityUserFlowAttributeId")
	}
}

func TestFormatIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeID(t *testing.T) {
	actual := NewIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeID("authenticationEventsFlowId", "identityUserFlowAttributeId").ID()
	expected := "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection/onAttributeCollectionExternalUsersSelfServiceSignUp/attributes/identityUserFlowAttributeId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId
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
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection/onAttributeCollectionExternalUsersSelfServiceSignUp",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection/onAttributeCollectionExternalUsersSelfServiceSignUp/attributes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection/onAttributeCollectionExternalUsersSelfServiceSignUp/attributes/identityUserFlowAttributeId",
			Expected: &IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId{
				AuthenticationEventsFlowId:  "authenticationEventsFlowId",
				IdentityUserFlowAttributeId: "identityUserFlowAttributeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection/onAttributeCollectionExternalUsersSelfServiceSignUp/attributes/identityUserFlowAttributeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeID(v.Input)
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

		if actual.IdentityUserFlowAttributeId != v.Expected.IdentityUserFlowAttributeId {
			t.Fatalf("Expected %q but got %q for IdentityUserFlowAttributeId", v.Expected.IdentityUserFlowAttributeId, actual.IdentityUserFlowAttributeId)
		}

	}
}

func TestParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId
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
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/eXtErNaLuSeRsSeLfSeRvIcEsIgNuPeVeNtSfLoW/oNaTtRiBuTeCoLlEcTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection/onAttributeCollectionExternalUsersSelfServiceSignUp",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/eXtErNaLuSeRsSeLfSeRvIcEsIgNuPeVeNtSfLoW/oNaTtRiBuTeCoLlEcTiOn/oNaTtRiBuTeCoLlEcTiOnExTeRnAlUsErSsElFsErViCeSiGnUp",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection/onAttributeCollectionExternalUsersSelfServiceSignUp/attributes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/eXtErNaLuSeRsSeLfSeRvIcEsIgNuPeVeNtSfLoW/oNaTtRiBuTeCoLlEcTiOn/oNaTtRiBuTeCoLlEcTiOnExTeRnAlUsErSsElFsErViCeSiGnUp/aTtRiBuTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection/onAttributeCollectionExternalUsersSelfServiceSignUp/attributes/identityUserFlowAttributeId",
			Expected: &IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId{
				AuthenticationEventsFlowId:  "authenticationEventsFlowId",
				IdentityUserFlowAttributeId: "identityUserFlowAttributeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection/onAttributeCollectionExternalUsersSelfServiceSignUp/attributes/identityUserFlowAttributeId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/eXtErNaLuSeRsSeLfSeRvIcEsIgNuPeVeNtSfLoW/oNaTtRiBuTeCoLlEcTiOn/oNaTtRiBuTeCoLlEcTiOnExTeRnAlUsErSsElFsErViCeSiGnUp/aTtRiBuTeS/iDeNtItYuSeRfLoWaTtRiBuTeId",
			Expected: &IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId{
				AuthenticationEventsFlowId:  "aUtHeNtIcAtIoNeVeNtSfLoWiD",
				IdentityUserFlowAttributeId: "iDeNtItYuSeRfLoWaTtRiBuTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/eXtErNaLuSeRsSeLfSeRvIcEsIgNuPeVeNtSfLoW/oNaTtRiBuTeCoLlEcTiOn/oNaTtRiBuTeCoLlEcTiOnExTeRnAlUsErSsElFsErViCeSiGnUp/aTtRiBuTeS/iDeNtItYuSeRfLoWaTtRiBuTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeIDInsensitively(v.Input)
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

		if actual.IdentityUserFlowAttributeId != v.Expected.IdentityUserFlowAttributeId {
			t.Fatalf("Expected %q but got %q for IdentityUserFlowAttributeId", v.Expected.IdentityUserFlowAttributeId, actual.IdentityUserFlowAttributeId)
		}

	}
}

func TestSegmentsForIdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId(t *testing.T) {
	segments := IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityAuthenticationEventsFlowIdExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
