package policyauthorizationpolicydefaultuserroleoverride

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyAuthorizationPolicyId{}

func TestNewPolicyAuthorizationPolicyID(t *testing.T) {
	id := NewPolicyAuthorizationPolicyID("authorizationPolicyIdValue")

	if id.AuthorizationPolicyId != "authorizationPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthorizationPolicyId'", id.AuthorizationPolicyId, "authorizationPolicyIdValue")
	}
}

func TestFormatPolicyAuthorizationPolicyID(t *testing.T) {
	actual := NewPolicyAuthorizationPolicyID("authorizationPolicyIdValue").ID()
	expected := "/policies/authorizationPolicy/authorizationPolicyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyAuthorizationPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyAuthorizationPolicyId
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
			// Valid URI
			Input: "/policies/authorizationPolicy/authorizationPolicyIdValue",
			Expected: &PolicyAuthorizationPolicyId{
				AuthorizationPolicyId: "authorizationPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/authorizationPolicy/authorizationPolicyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyAuthorizationPolicyID(v.Input)
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

	}
}

func TestParsePolicyAuthorizationPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyAuthorizationPolicyId
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
			// Valid URI
			Input: "/policies/authorizationPolicy/authorizationPolicyIdValue",
			Expected: &PolicyAuthorizationPolicyId{
				AuthorizationPolicyId: "authorizationPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/authorizationPolicy/authorizationPolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aUtHoRiZaTiOnPoLiCy/aUtHoRiZaTiOnPoLiCyIdVaLuE",
			Expected: &PolicyAuthorizationPolicyId{
				AuthorizationPolicyId: "aUtHoRiZaTiOnPoLiCyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aUtHoRiZaTiOnPoLiCy/aUtHoRiZaTiOnPoLiCyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyAuthorizationPolicyIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForPolicyAuthorizationPolicyId(t *testing.T) {
	segments := PolicyAuthorizationPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyAuthorizationPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
