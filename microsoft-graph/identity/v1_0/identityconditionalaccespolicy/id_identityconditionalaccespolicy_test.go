package identityconditionalaccespolicy

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityConditionalAccesPolicyId{}

func TestNewIdentityConditionalAccesPolicyID(t *testing.T) {
	id := NewIdentityConditionalAccesPolicyID("conditionalAccessPolicyIdValue")

	if id.ConditionalAccessPolicyId != "conditionalAccessPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ConditionalAccessPolicyId'", id.ConditionalAccessPolicyId, "conditionalAccessPolicyIdValue")
	}
}

func TestFormatIdentityConditionalAccesPolicyID(t *testing.T) {
	actual := NewIdentityConditionalAccesPolicyID("conditionalAccessPolicyIdValue").ID()
	expected := "/identity/conditionalAccess/policies/conditionalAccessPolicyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityConditionalAccesPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccesPolicyId
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
			Input: "/identity/conditionalAccess/policies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/conditionalAccess/policies/conditionalAccessPolicyIdValue",
			Expected: &IdentityConditionalAccesPolicyId{
				ConditionalAccessPolicyId: "conditionalAccessPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/policies/conditionalAccessPolicyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccesPolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ConditionalAccessPolicyId != v.Expected.ConditionalAccessPolicyId {
			t.Fatalf("Expected %q but got %q for ConditionalAccessPolicyId", v.Expected.ConditionalAccessPolicyId, actual.ConditionalAccessPolicyId)
		}

	}
}

func TestParseIdentityConditionalAccesPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccesPolicyId
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
			Input: "/identity/conditionalAccess/policies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/pOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/conditionalAccess/policies/conditionalAccessPolicyIdValue",
			Expected: &IdentityConditionalAccesPolicyId{
				ConditionalAccessPolicyId: "conditionalAccessPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/policies/conditionalAccessPolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/pOlIcIeS/cOnDiTiOnAlAcCeSsPoLiCyIdVaLuE",
			Expected: &IdentityConditionalAccesPolicyId{
				ConditionalAccessPolicyId: "cOnDiTiOnAlAcCeSsPoLiCyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/pOlIcIeS/cOnDiTiOnAlAcCeSsPoLiCyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccesPolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ConditionalAccessPolicyId != v.Expected.ConditionalAccessPolicyId {
			t.Fatalf("Expected %q but got %q for ConditionalAccessPolicyId", v.Expected.ConditionalAccessPolicyId, actual.ConditionalAccessPolicyId)
		}

	}
}

func TestSegmentsForIdentityConditionalAccesPolicyId(t *testing.T) {
	segments := IdentityConditionalAccesPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityConditionalAccesPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
