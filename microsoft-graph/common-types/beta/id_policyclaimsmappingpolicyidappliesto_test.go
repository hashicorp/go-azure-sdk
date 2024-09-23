package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyClaimsMappingPolicyIdAppliesToId{}

func TestNewPolicyClaimsMappingPolicyIdAppliesToID(t *testing.T) {
	id := NewPolicyClaimsMappingPolicyIdAppliesToID("claimsMappingPolicyId", "directoryObjectId")

	if id.ClaimsMappingPolicyId != "claimsMappingPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'ClaimsMappingPolicyId'", id.ClaimsMappingPolicyId, "claimsMappingPolicyId")
	}

	if id.DirectoryObjectId != "directoryObjectId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectId")
	}
}

func TestFormatPolicyClaimsMappingPolicyIdAppliesToID(t *testing.T) {
	actual := NewPolicyClaimsMappingPolicyIdAppliesToID("claimsMappingPolicyId", "directoryObjectId").ID()
	expected := "/policies/claimsMappingPolicies/claimsMappingPolicyId/appliesTo/directoryObjectId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyClaimsMappingPolicyIdAppliesToID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyClaimsMappingPolicyIdAppliesToId
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
			// Incomplete URI
			Input: "/policies/claimsMappingPolicies/claimsMappingPolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/claimsMappingPolicies/claimsMappingPolicyId/appliesTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/claimsMappingPolicies/claimsMappingPolicyId/appliesTo/directoryObjectId",
			Expected: &PolicyClaimsMappingPolicyIdAppliesToId{
				ClaimsMappingPolicyId: "claimsMappingPolicyId",
				DirectoryObjectId:     "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/claimsMappingPolicies/claimsMappingPolicyId/appliesTo/directoryObjectId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyClaimsMappingPolicyIdAppliesToID(v.Input)
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

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestParsePolicyClaimsMappingPolicyIdAppliesToIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyClaimsMappingPolicyIdAppliesToId
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
			// Incomplete URI
			Input: "/policies/claimsMappingPolicies/claimsMappingPolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/cLaImSmApPiNgPoLiCiEs/cLaImSmApPiNgPoLiCyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/claimsMappingPolicies/claimsMappingPolicyId/appliesTo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/cLaImSmApPiNgPoLiCiEs/cLaImSmApPiNgPoLiCyId/aPpLiEsTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/claimsMappingPolicies/claimsMappingPolicyId/appliesTo/directoryObjectId",
			Expected: &PolicyClaimsMappingPolicyIdAppliesToId{
				ClaimsMappingPolicyId: "claimsMappingPolicyId",
				DirectoryObjectId:     "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/claimsMappingPolicies/claimsMappingPolicyId/appliesTo/directoryObjectId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/cLaImSmApPiNgPoLiCiEs/cLaImSmApPiNgPoLiCyId/aPpLiEsTo/dIrEcToRyObJeCtId",
			Expected: &PolicyClaimsMappingPolicyIdAppliesToId{
				ClaimsMappingPolicyId: "cLaImSmApPiNgPoLiCyId",
				DirectoryObjectId:     "dIrEcToRyObJeCtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/cLaImSmApPiNgPoLiCiEs/cLaImSmApPiNgPoLiCyId/aPpLiEsTo/dIrEcToRyObJeCtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyClaimsMappingPolicyIdAppliesToIDInsensitively(v.Input)
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

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestSegmentsForPolicyClaimsMappingPolicyIdAppliesToId(t *testing.T) {
	segments := PolicyClaimsMappingPolicyIdAppliesToId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyClaimsMappingPolicyIdAppliesToId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
