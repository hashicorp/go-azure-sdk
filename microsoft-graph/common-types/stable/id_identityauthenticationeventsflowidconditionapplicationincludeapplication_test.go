package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId{}

func TestNewIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationID(t *testing.T) {
	id := NewIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationID("authenticationEventsFlowIdValue", "authenticationConditionApplicationAppIdValue")

	if id.AuthenticationEventsFlowId != "authenticationEventsFlowIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationEventsFlowId'", id.AuthenticationEventsFlowId, "authenticationEventsFlowIdValue")
	}

	if id.AuthenticationConditionApplicationAppId != "authenticationConditionApplicationAppIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationConditionApplicationAppId'", id.AuthenticationConditionApplicationAppId, "authenticationConditionApplicationAppIdValue")
	}
}

func TestFormatIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationID(t *testing.T) {
	actual := NewIdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationID("authenticationEventsFlowIdValue", "authenticationConditionApplicationAppIdValue").ID()
	expected := "/identity/authenticationEventsFlows/authenticationEventsFlowIdValue/conditions/applications/includeApplications/authenticationConditionApplicationAppIdValue"
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
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowIdValue/conditions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowIdValue/conditions/applications",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowIdValue/conditions/applications/includeApplications",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowIdValue/conditions/applications/includeApplications/authenticationConditionApplicationAppIdValue",
			Expected: &IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId{
				AuthenticationEventsFlowId:              "authenticationEventsFlowIdValue",
				AuthenticationConditionApplicationAppId: "authenticationConditionApplicationAppIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowIdValue/conditions/applications/includeApplications/authenticationConditionApplicationAppIdValue/extra",
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
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowIdValue/conditions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiDvAlUe/cOnDiTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowIdValue/conditions/applications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiDvAlUe/cOnDiTiOnS/aPpLiCaTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowIdValue/conditions/applications/includeApplications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiDvAlUe/cOnDiTiOnS/aPpLiCaTiOnS/iNcLuDeApPlIcAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowIdValue/conditions/applications/includeApplications/authenticationConditionApplicationAppIdValue",
			Expected: &IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId{
				AuthenticationEventsFlowId:              "authenticationEventsFlowIdValue",
				AuthenticationConditionApplicationAppId: "authenticationConditionApplicationAppIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/authenticationEventsFlows/authenticationEventsFlowIdValue/conditions/applications/includeApplications/authenticationConditionApplicationAppIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiDvAlUe/cOnDiTiOnS/aPpLiCaTiOnS/iNcLuDeApPlIcAtIoNs/aUtHeNtIcAtIoNcOnDiTiOnApPlIcAtIoNaPpIdVaLuE",
			Expected: &IdentityAuthenticationEventsFlowIdConditionApplicationIncludeApplicationId{
				AuthenticationEventsFlowId:              "aUtHeNtIcAtIoNeVeNtSfLoWiDvAlUe",
				AuthenticationConditionApplicationAppId: "aUtHeNtIcAtIoNcOnDiTiOnApPlIcAtIoNaPpIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/aUtHeNtIcAtIoNeVeNtSfLoWs/aUtHeNtIcAtIoNeVeNtSfLoWiDvAlUe/cOnDiTiOnS/aPpLiCaTiOnS/iNcLuDeApPlIcAtIoNs/aUtHeNtIcAtIoNcOnDiTiOnApPlIcAtIoNaPpIdVaLuE/extra",
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
