package policyclaimsmappingpolicyappliesto

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyClaimsMappingPolicyAppliesToId{}

func TestNewPolicyClaimsMappingPolicyAppliesToID(t *testing.T) {
	id := NewPolicyClaimsMappingPolicyAppliesToID("claimsMappingPolicyIdValue", "directoryObjectIdValue")

	if id.ClaimsMappingPolicyId != "claimsMappingPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ClaimsMappingPolicyId'", id.ClaimsMappingPolicyId, "claimsMappingPolicyIdValue")
	}

	if id.DirectoryObjectId != "directoryObjectIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectIdValue")
	}
}

func TestFormatPolicyClaimsMappingPolicyAppliesToID(t *testing.T) {
	actual := NewPolicyClaimsMappingPolicyAppliesToID("claimsMappingPolicyIdValue", "directoryObjectIdValue").ID()
	expected := "/policies/claimsMappingPolicies/claimsMappingPolicyIdValue/appliesTo/directoryObjectIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyClaimsMappingPolicyAppliesToID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyClaimsMappingPolicyAppliesToId
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
			Input: "/policies/claimsMappingPolicies/claimsMappingPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/claimsMappingPolicies/claimsMappingPolicyIdValue/appliesTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/claimsMappingPolicies/claimsMappingPolicyIdValue/appliesTo/directoryObjectIdValue",
			Expected: &PolicyClaimsMappingPolicyAppliesToId{
				ClaimsMappingPolicyId: "claimsMappingPolicyIdValue",
				DirectoryObjectId:     "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/claimsMappingPolicies/claimsMappingPolicyIdValue/appliesTo/directoryObjectIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyClaimsMappingPolicyAppliesToID(v.Input)
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

func TestParsePolicyClaimsMappingPolicyAppliesToIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyClaimsMappingPolicyAppliesToId
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
			Input: "/policies/claimsMappingPolicies/claimsMappingPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/cLaImSmApPiNgPoLiCiEs/cLaImSmApPiNgPoLiCyIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/claimsMappingPolicies/claimsMappingPolicyIdValue/appliesTo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/cLaImSmApPiNgPoLiCiEs/cLaImSmApPiNgPoLiCyIdVaLuE/aPpLiEsTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/claimsMappingPolicies/claimsMappingPolicyIdValue/appliesTo/directoryObjectIdValue",
			Expected: &PolicyClaimsMappingPolicyAppliesToId{
				ClaimsMappingPolicyId: "claimsMappingPolicyIdValue",
				DirectoryObjectId:     "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/claimsMappingPolicies/claimsMappingPolicyIdValue/appliesTo/directoryObjectIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/cLaImSmApPiNgPoLiCiEs/cLaImSmApPiNgPoLiCyIdVaLuE/aPpLiEsTo/dIrEcToRyObJeCtIdVaLuE",
			Expected: &PolicyClaimsMappingPolicyAppliesToId{
				ClaimsMappingPolicyId: "cLaImSmApPiNgPoLiCyIdVaLuE",
				DirectoryObjectId:     "dIrEcToRyObJeCtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/cLaImSmApPiNgPoLiCiEs/cLaImSmApPiNgPoLiCyIdVaLuE/aPpLiEsTo/dIrEcToRyObJeCtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyClaimsMappingPolicyAppliesToIDInsensitively(v.Input)
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

func TestSegmentsForPolicyClaimsMappingPolicyAppliesToId(t *testing.T) {
	segments := PolicyClaimsMappingPolicyAppliesToId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyClaimsMappingPolicyAppliesToId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
