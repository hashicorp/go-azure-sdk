package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyClaimsMappingPolicyId{}

func TestNewPolicyClaimsMappingPolicyID(t *testing.T) {
	id := NewPolicyClaimsMappingPolicyID("claimsMappingPolicyId")

	if id.ClaimsMappingPolicyId != "claimsMappingPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'ClaimsMappingPolicyId'", id.ClaimsMappingPolicyId, "claimsMappingPolicyId")
	}
}

func TestFormatPolicyClaimsMappingPolicyID(t *testing.T) {
	actual := NewPolicyClaimsMappingPolicyID("claimsMappingPolicyId").ID()
	expected := "/policies/claimsMappingPolicies/claimsMappingPolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyClaimsMappingPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyClaimsMappingPolicyId
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
			Input: "/policies/claimsMappingPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/claimsMappingPolicies/claimsMappingPolicyId",
			Expected: &PolicyClaimsMappingPolicyId{
				ClaimsMappingPolicyId: "claimsMappingPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/claimsMappingPolicies/claimsMappingPolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyClaimsMappingPolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ClaimsMappingPolicyId != v.Expected.ClaimsMappingPolicyId {
			t.Fatalf("Expected %q but got %q for ClaimsMappingPolicyId", v.Expected.ClaimsMappingPolicyId, actual.ClaimsMappingPolicyId)
		}

	}
}

func TestParsePolicyClaimsMappingPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyClaimsMappingPolicyId
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
			Input: "/policies/claimsMappingPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/cLaImSmApPiNgPoLiCiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/claimsMappingPolicies/claimsMappingPolicyId",
			Expected: &PolicyClaimsMappingPolicyId{
				ClaimsMappingPolicyId: "claimsMappingPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/claimsMappingPolicies/claimsMappingPolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/cLaImSmApPiNgPoLiCiEs/cLaImSmApPiNgPoLiCyId",
			Expected: &PolicyClaimsMappingPolicyId{
				ClaimsMappingPolicyId: "cLaImSmApPiNgPoLiCyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/cLaImSmApPiNgPoLiCiEs/cLaImSmApPiNgPoLiCyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyClaimsMappingPolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ClaimsMappingPolicyId != v.Expected.ClaimsMappingPolicyId {
			t.Fatalf("Expected %q but got %q for ClaimsMappingPolicyId", v.Expected.ClaimsMappingPolicyId, actual.ClaimsMappingPolicyId)
		}

	}
}

func TestSegmentsForPolicyClaimsMappingPolicyId(t *testing.T) {
	segments := PolicyClaimsMappingPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyClaimsMappingPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
