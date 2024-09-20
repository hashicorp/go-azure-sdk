package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyTokenLifetimePolicyId{}

func TestNewPolicyTokenLifetimePolicyID(t *testing.T) {
	id := NewPolicyTokenLifetimePolicyID("tokenLifetimePolicyId")

	if id.TokenLifetimePolicyId != "tokenLifetimePolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'TokenLifetimePolicyId'", id.TokenLifetimePolicyId, "tokenLifetimePolicyId")
	}
}

func TestFormatPolicyTokenLifetimePolicyID(t *testing.T) {
	actual := NewPolicyTokenLifetimePolicyID("tokenLifetimePolicyId").ID()
	expected := "/policies/tokenLifetimePolicies/tokenLifetimePolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyTokenLifetimePolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyTokenLifetimePolicyId
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
			// Valid URI
			Input: "/policies/tokenLifetimePolicies/tokenLifetimePolicyId",
			Expected: &PolicyTokenLifetimePolicyId{
				TokenLifetimePolicyId: "tokenLifetimePolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/tokenLifetimePolicies/tokenLifetimePolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyTokenLifetimePolicyID(v.Input)
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

	}
}

func TestParsePolicyTokenLifetimePolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyTokenLifetimePolicyId
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
			// Valid URI
			Input: "/policies/tokenLifetimePolicies/tokenLifetimePolicyId",
			Expected: &PolicyTokenLifetimePolicyId{
				TokenLifetimePolicyId: "tokenLifetimePolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/tokenLifetimePolicies/tokenLifetimePolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/tOkEnLiFeTiMePoLiCiEs/tOkEnLiFeTiMePoLiCyId",
			Expected: &PolicyTokenLifetimePolicyId{
				TokenLifetimePolicyId: "tOkEnLiFeTiMePoLiCyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/tOkEnLiFeTiMePoLiCiEs/tOkEnLiFeTiMePoLiCyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyTokenLifetimePolicyIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForPolicyTokenLifetimePolicyId(t *testing.T) {
	segments := PolicyTokenLifetimePolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyTokenLifetimePolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
