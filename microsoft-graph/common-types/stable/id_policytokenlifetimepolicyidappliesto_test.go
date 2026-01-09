package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyTokenLifetimePolicyIdAppliesToId{}

func TestNewPolicyTokenLifetimePolicyIdAppliesToID(t *testing.T) {
	id := NewPolicyTokenLifetimePolicyIdAppliesToID("tokenLifetimePolicyId", "directoryObjectId")

	if id.TokenLifetimePolicyId != "tokenLifetimePolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'TokenLifetimePolicyId'", id.TokenLifetimePolicyId, "tokenLifetimePolicyId")
	}

	if id.DirectoryObjectId != "directoryObjectId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectId")
	}
}

func TestFormatPolicyTokenLifetimePolicyIdAppliesToID(t *testing.T) {
	actual := NewPolicyTokenLifetimePolicyIdAppliesToID("tokenLifetimePolicyId", "directoryObjectId").ID()
	expected := "/policies/tokenLifetimePolicies/tokenLifetimePolicyId/appliesTo/directoryObjectId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyTokenLifetimePolicyIdAppliesToID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyTokenLifetimePolicyIdAppliesToId
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
			Input: "/policies/tokenLifetimePolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/tokenLifetimePolicies/tokenLifetimePolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/tokenLifetimePolicies/tokenLifetimePolicyId/appliesTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/tokenLifetimePolicies/tokenLifetimePolicyId/appliesTo/directoryObjectId",
			Expected: &PolicyTokenLifetimePolicyIdAppliesToId{
				TokenLifetimePolicyId: "tokenLifetimePolicyId",
				DirectoryObjectId:     "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/tokenLifetimePolicies/tokenLifetimePolicyId/appliesTo/directoryObjectId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyTokenLifetimePolicyIdAppliesToID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TokenLifetimePolicyId != v.Expected.TokenLifetimePolicyId {
			t.Fatalf("Expected %q but got %q for TokenLifetimePolicyId", v.Expected.TokenLifetimePolicyId, actual.TokenLifetimePolicyId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestParsePolicyTokenLifetimePolicyIdAppliesToIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyTokenLifetimePolicyIdAppliesToId
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
			Input: "/policies/tokenLifetimePolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/tOkEnLiFeTiMePoLiCiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/tokenLifetimePolicies/tokenLifetimePolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/tOkEnLiFeTiMePoLiCiEs/tOkEnLiFeTiMePoLiCyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/tokenLifetimePolicies/tokenLifetimePolicyId/appliesTo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/tOkEnLiFeTiMePoLiCiEs/tOkEnLiFeTiMePoLiCyId/aPpLiEsTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/tokenLifetimePolicies/tokenLifetimePolicyId/appliesTo/directoryObjectId",
			Expected: &PolicyTokenLifetimePolicyIdAppliesToId{
				TokenLifetimePolicyId: "tokenLifetimePolicyId",
				DirectoryObjectId:     "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/tokenLifetimePolicies/tokenLifetimePolicyId/appliesTo/directoryObjectId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/tOkEnLiFeTiMePoLiCiEs/tOkEnLiFeTiMePoLiCyId/aPpLiEsTo/dIrEcToRyObJeCtId",
			Expected: &PolicyTokenLifetimePolicyIdAppliesToId{
				TokenLifetimePolicyId: "tOkEnLiFeTiMePoLiCyId",
				DirectoryObjectId:     "dIrEcToRyObJeCtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/tOkEnLiFeTiMePoLiCiEs/tOkEnLiFeTiMePoLiCyId/aPpLiEsTo/dIrEcToRyObJeCtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyTokenLifetimePolicyIdAppliesToIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TokenLifetimePolicyId != v.Expected.TokenLifetimePolicyId {
			t.Fatalf("Expected %q but got %q for TokenLifetimePolicyId", v.Expected.TokenLifetimePolicyId, actual.TokenLifetimePolicyId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestSegmentsForPolicyTokenLifetimePolicyIdAppliesToId(t *testing.T) {
	segments := PolicyTokenLifetimePolicyIdAppliesToId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyTokenLifetimePolicyIdAppliesToId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
