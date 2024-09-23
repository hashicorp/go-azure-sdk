package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyHomeRealmDiscoveryPolicyId{}

func TestNewPolicyHomeRealmDiscoveryPolicyID(t *testing.T) {
	id := NewPolicyHomeRealmDiscoveryPolicyID("homeRealmDiscoveryPolicyId")

	if id.HomeRealmDiscoveryPolicyId != "homeRealmDiscoveryPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'HomeRealmDiscoveryPolicyId'", id.HomeRealmDiscoveryPolicyId, "homeRealmDiscoveryPolicyId")
	}
}

func TestFormatPolicyHomeRealmDiscoveryPolicyID(t *testing.T) {
	actual := NewPolicyHomeRealmDiscoveryPolicyID("homeRealmDiscoveryPolicyId").ID()
	expected := "/policies/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyHomeRealmDiscoveryPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyHomeRealmDiscoveryPolicyId
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
			Input: "/policies/homeRealmDiscoveryPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyId",
			Expected: &PolicyHomeRealmDiscoveryPolicyId{
				HomeRealmDiscoveryPolicyId: "homeRealmDiscoveryPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyHomeRealmDiscoveryPolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.HomeRealmDiscoveryPolicyId != v.Expected.HomeRealmDiscoveryPolicyId {
			t.Fatalf("Expected %q but got %q for HomeRealmDiscoveryPolicyId", v.Expected.HomeRealmDiscoveryPolicyId, actual.HomeRealmDiscoveryPolicyId)
		}

	}
}

func TestParsePolicyHomeRealmDiscoveryPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyHomeRealmDiscoveryPolicyId
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
			Input: "/policies/homeRealmDiscoveryPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/hOmErEaLmDiScOvErYpOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyId",
			Expected: &PolicyHomeRealmDiscoveryPolicyId{
				HomeRealmDiscoveryPolicyId: "homeRealmDiscoveryPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/hOmErEaLmDiScOvErYpOlIcIeS/hOmErEaLmDiScOvErYpOlIcYiD",
			Expected: &PolicyHomeRealmDiscoveryPolicyId{
				HomeRealmDiscoveryPolicyId: "hOmErEaLmDiScOvErYpOlIcYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/hOmErEaLmDiScOvErYpOlIcIeS/hOmErEaLmDiScOvErYpOlIcYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyHomeRealmDiscoveryPolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.HomeRealmDiscoveryPolicyId != v.Expected.HomeRealmDiscoveryPolicyId {
			t.Fatalf("Expected %q but got %q for HomeRealmDiscoveryPolicyId", v.Expected.HomeRealmDiscoveryPolicyId, actual.HomeRealmDiscoveryPolicyId)
		}

	}
}

func TestSegmentsForPolicyHomeRealmDiscoveryPolicyId(t *testing.T) {
	segments := PolicyHomeRealmDiscoveryPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyHomeRealmDiscoveryPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
