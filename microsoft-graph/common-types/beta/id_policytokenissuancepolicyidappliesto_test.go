package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyTokenIssuancePolicyIdAppliesToId{}

func TestNewPolicyTokenIssuancePolicyIdAppliesToID(t *testing.T) {
	id := NewPolicyTokenIssuancePolicyIdAppliesToID("tokenIssuancePolicyId", "directoryObjectId")

	if id.TokenIssuancePolicyId != "tokenIssuancePolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'TokenIssuancePolicyId'", id.TokenIssuancePolicyId, "tokenIssuancePolicyId")
	}

	if id.DirectoryObjectId != "directoryObjectId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectId")
	}
}

func TestFormatPolicyTokenIssuancePolicyIdAppliesToID(t *testing.T) {
	actual := NewPolicyTokenIssuancePolicyIdAppliesToID("tokenIssuancePolicyId", "directoryObjectId").ID()
	expected := "/policies/tokenIssuancePolicies/tokenIssuancePolicyId/appliesTo/directoryObjectId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyTokenIssuancePolicyIdAppliesToID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyTokenIssuancePolicyIdAppliesToId
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
			Input: "/policies/tokenIssuancePolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/tokenIssuancePolicies/tokenIssuancePolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/tokenIssuancePolicies/tokenIssuancePolicyId/appliesTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/tokenIssuancePolicies/tokenIssuancePolicyId/appliesTo/directoryObjectId",
			Expected: &PolicyTokenIssuancePolicyIdAppliesToId{
				TokenIssuancePolicyId: "tokenIssuancePolicyId",
				DirectoryObjectId:     "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/tokenIssuancePolicies/tokenIssuancePolicyId/appliesTo/directoryObjectId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyTokenIssuancePolicyIdAppliesToID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TokenIssuancePolicyId != v.Expected.TokenIssuancePolicyId {
			t.Fatalf("Expected %q but got %q for TokenIssuancePolicyId", v.Expected.TokenIssuancePolicyId, actual.TokenIssuancePolicyId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestParsePolicyTokenIssuancePolicyIdAppliesToIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyTokenIssuancePolicyIdAppliesToId
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
			Input: "/policies/tokenIssuancePolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/tOkEnIsSuAnCePoLiCiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/tokenIssuancePolicies/tokenIssuancePolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/tOkEnIsSuAnCePoLiCiEs/tOkEnIsSuAnCePoLiCyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/tokenIssuancePolicies/tokenIssuancePolicyId/appliesTo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/tOkEnIsSuAnCePoLiCiEs/tOkEnIsSuAnCePoLiCyId/aPpLiEsTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/tokenIssuancePolicies/tokenIssuancePolicyId/appliesTo/directoryObjectId",
			Expected: &PolicyTokenIssuancePolicyIdAppliesToId{
				TokenIssuancePolicyId: "tokenIssuancePolicyId",
				DirectoryObjectId:     "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/tokenIssuancePolicies/tokenIssuancePolicyId/appliesTo/directoryObjectId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/tOkEnIsSuAnCePoLiCiEs/tOkEnIsSuAnCePoLiCyId/aPpLiEsTo/dIrEcToRyObJeCtId",
			Expected: &PolicyTokenIssuancePolicyIdAppliesToId{
				TokenIssuancePolicyId: "tOkEnIsSuAnCePoLiCyId",
				DirectoryObjectId:     "dIrEcToRyObJeCtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/tOkEnIsSuAnCePoLiCiEs/tOkEnIsSuAnCePoLiCyId/aPpLiEsTo/dIrEcToRyObJeCtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyTokenIssuancePolicyIdAppliesToIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TokenIssuancePolicyId != v.Expected.TokenIssuancePolicyId {
			t.Fatalf("Expected %q but got %q for TokenIssuancePolicyId", v.Expected.TokenIssuancePolicyId, actual.TokenIssuancePolicyId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestSegmentsForPolicyTokenIssuancePolicyIdAppliesToId(t *testing.T) {
	segments := PolicyTokenIssuancePolicyIdAppliesToId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyTokenIssuancePolicyIdAppliesToId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
