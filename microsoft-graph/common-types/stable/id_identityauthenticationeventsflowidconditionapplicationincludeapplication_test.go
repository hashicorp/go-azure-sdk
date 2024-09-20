package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId{}

func TestNewIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationID(t *testing.T) {
	id := NewIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationID("authenticationEventsFlowId", "authenticationConditionApplicationAppId")

	if id.AuthenticationEventsFlowId != "authenticationEventsFlowId" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationEventsFlowId'", id.AuthenticationEventsFlowId, "authenticationEventsFlowId")
	}

	if id.AuthenticationConditionApplicationAppId != "authenticationConditionApplicationAppId" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationConditionApplicationAppId'", id.AuthenticationConditionApplicationAppId, "authenticationConditionApplicationAppId")
	}
}

func TestFormatIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationID(t *testing.T) {
	actual := NewIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationID("authenticationEventsFlowId", "authenticationConditionApplicationAppId").ID()
	expected := "/identity/authenticationEventsFlows/authenticationEventsFlowId/conditions/applications/includeApplications/authenticationConditionApplicationAppId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId
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
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/conditions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/conditions/applications",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/conditions/applications/includeApplications",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/conditions/applications/includeApplications/authenticationConditionApplicationAppId",
			Expected: &IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId{
				AuthenticationEventsFlowId:              "authenticationEventsFlowId",
				AuthenticationConditionApplicationAppId: "authenticationConditionApplicationAppId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/conditions/applications/includeApplications/authenticationConditionApplicationAppId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationID(v.Input)
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

func TestParseIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId
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
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/conditions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/cOnDiTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/conditions/applications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/cOnDiTiOnS/aPpLiCaTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/conditions/applications/includeApplications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/cOnDiTiOnS/aPpLiCaTiOnS/iNcLuDeApPlIcAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/conditions/applications/includeApplications/authenticationConditionApplicationAppId",
			Expected: &IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId{
				AuthenticationEventsFlowId:              "authenticationEventsFlowId",
				AuthenticationConditionApplicationAppId: "authenticationConditionApplicationAppId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowId/conditions/applications/includeApplications/authenticationConditionApplicationAppId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/cOnDiTiOnS/aPpLiCaTiOnS/iNcLuDeApPlIcAtIoNs/aUtHeNtIcAtIoNcOnDiTiOnApPlIcAtIoNaPpId",
			Expected: &IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId{
				AuthenticationEventsFlowId:              "aUtHeNtIcAtIoNeVeNtSfLoWiD",
				AuthenticationConditionApplicationAppId: "aUtHeNtIcAtIoNcOnDiTiOnApPlIcAtIoNaPpId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiD/cOnDiTiOnS/aPpLiCaTiOnS/iNcLuDeApPlIcAtIoNs/aUtHeNtIcAtIoNcOnDiTiOnApPlIcAtIoNaPpId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationIDInsensitively(v.Input)
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

func TestSegmentsForIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId(t *testing.T) {
	segments := IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
