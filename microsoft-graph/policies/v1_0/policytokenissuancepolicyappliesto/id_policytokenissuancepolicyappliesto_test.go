package policytokenissuancepolicyappliesto

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyTokenIssuancePolicyAppliesToId{}

func TestNewPolicyTokenIssuancePolicyAppliesToID(t *testing.T) {
	id := NewPolicyTokenIssuancePolicyAppliesToID("tokenIssuancePolicyIdValue", "directoryObjectIdValue")

	if id.TokenIssuancePolicyId != "tokenIssuancePolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TokenIssuancePolicyId'", id.TokenIssuancePolicyId, "tokenIssuancePolicyIdValue")
	}

	if id.DirectoryObjectId != "directoryObjectIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectIdValue")
	}
}

func TestFormatPolicyTokenIssuancePolicyAppliesToID(t *testing.T) {
	actual := NewPolicyTokenIssuancePolicyAppliesToID("tokenIssuancePolicyIdValue", "directoryObjectIdValue").ID()
	expected := "/policies/tokenIssuancePolicies/tokenIssuancePolicyIdValue/appliesTo/directoryObjectIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyTokenIssuancePolicyAppliesToID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyTokenIssuancePolicyAppliesToId
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
			Input: "/policies/tokenIssuancePolicies/tokenIssuancePolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/tokenIssuancePolicies/tokenIssuancePolicyIdValue/appliesTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/tokenIssuancePolicies/tokenIssuancePolicyIdValue/appliesTo/directoryObjectIdValue",
			Expected: &PolicyTokenIssuancePolicyAppliesToId{
				TokenIssuancePolicyId: "tokenIssuancePolicyIdValue",
				DirectoryObjectId:     "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/tokenIssuancePolicies/tokenIssuancePolicyIdValue/appliesTo/directoryObjectIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyTokenIssuancePolicyAppliesToID(v.Input)
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

func TestParsePolicyTokenIssuancePolicyAppliesToIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyTokenIssuancePolicyAppliesToId
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
			Input: "/policies/tokenIssuancePolicies/tokenIssuancePolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/tOkEnIsSuAnCePoLiCiEs/tOkEnIsSuAnCePoLiCyIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/tokenIssuancePolicies/tokenIssuancePolicyIdValue/appliesTo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/tOkEnIsSuAnCePoLiCiEs/tOkEnIsSuAnCePoLiCyIdVaLuE/aPpLiEsTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/tokenIssuancePolicies/tokenIssuancePolicyIdValue/appliesTo/directoryObjectIdValue",
			Expected: &PolicyTokenIssuancePolicyAppliesToId{
				TokenIssuancePolicyId: "tokenIssuancePolicyIdValue",
				DirectoryObjectId:     "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/tokenIssuancePolicies/tokenIssuancePolicyIdValue/appliesTo/directoryObjectIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/tOkEnIsSuAnCePoLiCiEs/tOkEnIsSuAnCePoLiCyIdVaLuE/aPpLiEsTo/dIrEcToRyObJeCtIdVaLuE",
			Expected: &PolicyTokenIssuancePolicyAppliesToId{
				TokenIssuancePolicyId: "tOkEnIsSuAnCePoLiCyIdVaLuE",
				DirectoryObjectId:     "dIrEcToRyObJeCtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/tOkEnIsSuAnCePoLiCiEs/tOkEnIsSuAnCePoLiCyIdVaLuE/aPpLiEsTo/dIrEcToRyObJeCtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyTokenIssuancePolicyAppliesToIDInsensitively(v.Input)
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

func TestSegmentsForPolicyTokenIssuancePolicyAppliesToId(t *testing.T) {
	segments := PolicyTokenIssuancePolicyAppliesToId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyTokenIssuancePolicyAppliesToId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
