package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityConditionalAccessPolicyId{}

func TestNewIdentityConditionalAccessPolicyID(t *testing.T) {
	id := NewIdentityConditionalAccessPolicyID("conditionalAccessPolicyId")

	if id.ConditionalAccessPolicyId != "conditionalAccessPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'ConditionalAccessPolicyId'", id.ConditionalAccessPolicyId, "conditionalAccessPolicyId")
	}
}

func TestFormatIdentityConditionalAccessPolicyID(t *testing.T) {
	actual := NewIdentityConditionalAccessPolicyID("conditionalAccessPolicyId").ID()
	expected := "/identity/conditionalAccess/policies/conditionalAccessPolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityConditionalAccessPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccessPolicyId
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
			Input: "/identity/conditionalAccess/policies/conditionalAccessPolicyId",
			Expected: &IdentityConditionalAccessPolicyId{
				ConditionalAccessPolicyId: "conditionalAccessPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/policies/conditionalAccessPolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccessPolicyID(v.Input)
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

func TestParseIdentityConditionalAccessPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccessPolicyId
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
			Input: "/identity/conditionalAccess/policies/conditionalAccessPolicyId",
			Expected: &IdentityConditionalAccessPolicyId{
				ConditionalAccessPolicyId: "conditionalAccessPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/policies/conditionalAccessPolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/pOlIcIeS/cOnDiTiOnAlAcCeSsPoLiCyId",
			Expected: &IdentityConditionalAccessPolicyId{
				ConditionalAccessPolicyId: "cOnDiTiOnAlAcCeSsPoLiCyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/pOlIcIeS/cOnDiTiOnAlAcCeSsPoLiCyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccessPolicyIDInsensitively(v.Input)
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

func TestSegmentsForIdentityConditionalAccessPolicyId(t *testing.T) {
	segments := IdentityConditionalAccessPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityConditionalAccessPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
