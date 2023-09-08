package policyfeaturerolloutpolicyappliesto

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyFeatureRolloutPolicyAppliesToId{}

func TestNewPolicyFeatureRolloutPolicyAppliesToID(t *testing.T) {
	id := NewPolicyFeatureRolloutPolicyAppliesToID("featureRolloutPolicyIdValue", "directoryObjectIdValue")

	if id.FeatureRolloutPolicyId != "featureRolloutPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'FeatureRolloutPolicyId'", id.FeatureRolloutPolicyId, "featureRolloutPolicyIdValue")
	}

	if id.DirectoryObjectId != "directoryObjectIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectIdValue")
	}
}

func TestFormatPolicyFeatureRolloutPolicyAppliesToID(t *testing.T) {
	actual := NewPolicyFeatureRolloutPolicyAppliesToID("featureRolloutPolicyIdValue", "directoryObjectIdValue").ID()
	expected := "/policies/featureRolloutPolicies/featureRolloutPolicyIdValue/appliesTo/directoryObjectIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyFeatureRolloutPolicyAppliesToID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyFeatureRolloutPolicyAppliesToId
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
			Input: "/policies/featureRolloutPolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/featureRolloutPolicies/featureRolloutPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/featureRolloutPolicies/featureRolloutPolicyIdValue/appliesTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/featureRolloutPolicies/featureRolloutPolicyIdValue/appliesTo/directoryObjectIdValue",
			Expected: &PolicyFeatureRolloutPolicyAppliesToId{
				FeatureRolloutPolicyId: "featureRolloutPolicyIdValue",
				DirectoryObjectId:      "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/featureRolloutPolicies/featureRolloutPolicyIdValue/appliesTo/directoryObjectIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyFeatureRolloutPolicyAppliesToID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.FeatureRolloutPolicyId != v.Expected.FeatureRolloutPolicyId {
			t.Fatalf("Expected %q but got %q for FeatureRolloutPolicyId", v.Expected.FeatureRolloutPolicyId, actual.FeatureRolloutPolicyId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestParsePolicyFeatureRolloutPolicyAppliesToIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyFeatureRolloutPolicyAppliesToId
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
			Input: "/policies/featureRolloutPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/fEaTuReRoLlOuTpOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/featureRolloutPolicies/featureRolloutPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/fEaTuReRoLlOuTpOlIcIeS/fEaTuReRoLlOuTpOlIcYiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/featureRolloutPolicies/featureRolloutPolicyIdValue/appliesTo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/fEaTuReRoLlOuTpOlIcIeS/fEaTuReRoLlOuTpOlIcYiDvAlUe/aPpLiEsTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/featureRolloutPolicies/featureRolloutPolicyIdValue/appliesTo/directoryObjectIdValue",
			Expected: &PolicyFeatureRolloutPolicyAppliesToId{
				FeatureRolloutPolicyId: "featureRolloutPolicyIdValue",
				DirectoryObjectId:      "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/featureRolloutPolicies/featureRolloutPolicyIdValue/appliesTo/directoryObjectIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/fEaTuReRoLlOuTpOlIcIeS/fEaTuReRoLlOuTpOlIcYiDvAlUe/aPpLiEsTo/dIrEcToRyObJeCtIdVaLuE",
			Expected: &PolicyFeatureRolloutPolicyAppliesToId{
				FeatureRolloutPolicyId: "fEaTuReRoLlOuTpOlIcYiDvAlUe",
				DirectoryObjectId:      "dIrEcToRyObJeCtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/fEaTuReRoLlOuTpOlIcIeS/fEaTuReRoLlOuTpOlIcYiDvAlUe/aPpLiEsTo/dIrEcToRyObJeCtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyFeatureRolloutPolicyAppliesToIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.FeatureRolloutPolicyId != v.Expected.FeatureRolloutPolicyId {
			t.Fatalf("Expected %q but got %q for FeatureRolloutPolicyId", v.Expected.FeatureRolloutPolicyId, actual.FeatureRolloutPolicyId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestSegmentsForPolicyFeatureRolloutPolicyAppliesToId(t *testing.T) {
	segments := PolicyFeatureRolloutPolicyAppliesToId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyFeatureRolloutPolicyAppliesToId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
