package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId{}

func TestNewPolicyAuthorizationPolicyIdDefaultUserRoleOverrideID(t *testing.T) {
	id := NewPolicyAuthorizationPolicyIdDefaultUserRoleOverrideID("authorizationPolicyId", "defaultUserRoleOverrideId")

	if id.AuthorizationPolicyId != "authorizationPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthorizationPolicyId'", id.AuthorizationPolicyId, "authorizationPolicyId")
	}

	if id.DefaultUserRoleOverrideId != "defaultUserRoleOverrideId" {
		t.Fatalf("Expected %q but got %q for Segment 'DefaultUserRoleOverrideId'", id.DefaultUserRoleOverrideId, "defaultUserRoleOverrideId")
	}
}

func TestFormatPolicyAuthorizationPolicyIdDefaultUserRoleOverrideID(t *testing.T) {
	actual := NewPolicyAuthorizationPolicyIdDefaultUserRoleOverrideID("authorizationPolicyId", "defaultUserRoleOverrideId").ID()
	expected := "/policies/authorizationPolicy/authorizationPolicyId/defaultUserRoleOverrides/defaultUserRoleOverrideId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyAuthorizationPolicyIdDefaultUserRoleOverrideID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId
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
			Input: "/policies/authorizationPolicy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/authorizationPolicy/authorizationPolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/authorizationPolicy/authorizationPolicyId/defaultUserRoleOverrides",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/authorizationPolicy/authorizationPolicyId/defaultUserRoleOverrides/defaultUserRoleOverrideId",
			Expected: &PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId{
				AuthorizationPolicyId:     "authorizationPolicyId",
				DefaultUserRoleOverrideId: "defaultUserRoleOverrideId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/authorizationPolicy/authorizationPolicyId/defaultUserRoleOverrides/defaultUserRoleOverrideId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyAuthorizationPolicyIdDefaultUserRoleOverrideID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuthorizationPolicyId != v.Expected.AuthorizationPolicyId {
			t.Fatalf("Expected %q but got %q for AuthorizationPolicyId", v.Expected.AuthorizationPolicyId, actual.AuthorizationPolicyId)
		}

		if actual.DefaultUserRoleOverrideId != v.Expected.DefaultUserRoleOverrideId {
			t.Fatalf("Expected %q but got %q for DefaultUserRoleOverrideId", v.Expected.DefaultUserRoleOverrideId, actual.DefaultUserRoleOverrideId)
		}

	}
}

func TestParsePolicyAuthorizationPolicyIdDefaultUserRoleOverrideIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId
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
			Input: "/policies/authorizationPolicy",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aUtHoRiZaTiOnPoLiCy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/authorizationPolicy/authorizationPolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aUtHoRiZaTiOnPoLiCy/aUtHoRiZaTiOnPoLiCyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/authorizationPolicy/authorizationPolicyId/defaultUserRoleOverrides",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aUtHoRiZaTiOnPoLiCy/aUtHoRiZaTiOnPoLiCyId/dEfAuLtUsErRoLeOvErRiDeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/authorizationPolicy/authorizationPolicyId/defaultUserRoleOverrides/defaultUserRoleOverrideId",
			Expected: &PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId{
				AuthorizationPolicyId:     "authorizationPolicyId",
				DefaultUserRoleOverrideId: "defaultUserRoleOverrideId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/authorizationPolicy/authorizationPolicyId/defaultUserRoleOverrides/defaultUserRoleOverrideId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aUtHoRiZaTiOnPoLiCy/aUtHoRiZaTiOnPoLiCyId/dEfAuLtUsErRoLeOvErRiDeS/dEfAuLtUsErRoLeOvErRiDeId",
			Expected: &PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId{
				AuthorizationPolicyId:     "aUtHoRiZaTiOnPoLiCyId",
				DefaultUserRoleOverrideId: "dEfAuLtUsErRoLeOvErRiDeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aUtHoRiZaTiOnPoLiCy/aUtHoRiZaTiOnPoLiCyId/dEfAuLtUsErRoLeOvErRiDeS/dEfAuLtUsErRoLeOvErRiDeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyAuthorizationPolicyIdDefaultUserRoleOverrideIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuthorizationPolicyId != v.Expected.AuthorizationPolicyId {
			t.Fatalf("Expected %q but got %q for AuthorizationPolicyId", v.Expected.AuthorizationPolicyId, actual.AuthorizationPolicyId)
		}

		if actual.DefaultUserRoleOverrideId != v.Expected.DefaultUserRoleOverrideId {
			t.Fatalf("Expected %q but got %q for DefaultUserRoleOverrideId", v.Expected.DefaultUserRoleOverrideId, actual.DefaultUserRoleOverrideId)
		}

	}
}

func TestSegmentsForPolicyAuthorizationPolicyIdDefaultUserRoleOverrideId(t *testing.T) {
	segments := PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyAuthorizationPolicyIdDefaultUserRoleOverrideId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
