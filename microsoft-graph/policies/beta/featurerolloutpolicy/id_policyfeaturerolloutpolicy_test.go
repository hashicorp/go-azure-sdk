package featurerolloutpolicy

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyFeatureRolloutPolicyId{}

func TestNewPolicyFeatureRolloutPolicyID(t *testing.T) {
	id := NewPolicyFeatureRolloutPolicyID("featureRolloutPolicyIdValue")

	if id.FeatureRolloutPolicyId != "featureRolloutPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'FeatureRolloutPolicyId'", id.FeatureRolloutPolicyId, "featureRolloutPolicyIdValue")
	}
}

func TestFormatPolicyFeatureRolloutPolicyID(t *testing.T) {
	actual := NewPolicyFeatureRolloutPolicyID("featureRolloutPolicyIdValue").ID()
	expected := "/policies/featureRolloutPolicies/featureRolloutPolicyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyFeatureRolloutPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyFeatureRolloutPolicyId
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
			// Valid URI
			Input: "/policies/featureRolloutPolicies/featureRolloutPolicyIdValue",
			Expected: &PolicyFeatureRolloutPolicyId{
				FeatureRolloutPolicyId: "featureRolloutPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/featureRolloutPolicies/featureRolloutPolicyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyFeatureRolloutPolicyID(v.Input)
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

	}
}

func TestParsePolicyFeatureRolloutPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyFeatureRolloutPolicyId
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
			// Valid URI
			Input: "/policies/featureRolloutPolicies/featureRolloutPolicyIdValue",
			Expected: &PolicyFeatureRolloutPolicyId{
				FeatureRolloutPolicyId: "featureRolloutPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/featureRolloutPolicies/featureRolloutPolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/fEaTuReRoLlOuTpOlIcIeS/fEaTuReRoLlOuTpOlIcYiDvAlUe",
			Expected: &PolicyFeatureRolloutPolicyId{
				FeatureRolloutPolicyId: "fEaTuReRoLlOuTpOlIcYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/fEaTuReRoLlOuTpOlIcIeS/fEaTuReRoLlOuTpOlIcYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyFeatureRolloutPolicyIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForPolicyFeatureRolloutPolicyId(t *testing.T) {
	segments := PolicyFeatureRolloutPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyFeatureRolloutPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
