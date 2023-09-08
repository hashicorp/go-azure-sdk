package policyactivitybasedtimeoutpolicy

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyActivityBasedTimeoutPolicyId{}

func TestNewPolicyActivityBasedTimeoutPolicyID(t *testing.T) {
	id := NewPolicyActivityBasedTimeoutPolicyID("activityBasedTimeoutPolicyIdValue")

	if id.ActivityBasedTimeoutPolicyId != "activityBasedTimeoutPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ActivityBasedTimeoutPolicyId'", id.ActivityBasedTimeoutPolicyId, "activityBasedTimeoutPolicyIdValue")
	}
}

func TestFormatPolicyActivityBasedTimeoutPolicyID(t *testing.T) {
	actual := NewPolicyActivityBasedTimeoutPolicyID("activityBasedTimeoutPolicyIdValue").ID()
	expected := "/policies/activityBasedTimeoutPolicies/activityBasedTimeoutPolicyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyActivityBasedTimeoutPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyActivityBasedTimeoutPolicyId
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
			Input: "/policies/activityBasedTimeoutPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/activityBasedTimeoutPolicies/activityBasedTimeoutPolicyIdValue",
			Expected: &PolicyActivityBasedTimeoutPolicyId{
				ActivityBasedTimeoutPolicyId: "activityBasedTimeoutPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/activityBasedTimeoutPolicies/activityBasedTimeoutPolicyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyActivityBasedTimeoutPolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ActivityBasedTimeoutPolicyId != v.Expected.ActivityBasedTimeoutPolicyId {
			t.Fatalf("Expected %q but got %q for ActivityBasedTimeoutPolicyId", v.Expected.ActivityBasedTimeoutPolicyId, actual.ActivityBasedTimeoutPolicyId)
		}

	}
}

func TestParsePolicyActivityBasedTimeoutPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyActivityBasedTimeoutPolicyId
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
			Input: "/policies/activityBasedTimeoutPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aCtIvItYbAsEdTiMeOuTpOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/activityBasedTimeoutPolicies/activityBasedTimeoutPolicyIdValue",
			Expected: &PolicyActivityBasedTimeoutPolicyId{
				ActivityBasedTimeoutPolicyId: "activityBasedTimeoutPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/activityBasedTimeoutPolicies/activityBasedTimeoutPolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aCtIvItYbAsEdTiMeOuTpOlIcIeS/aCtIvItYbAsEdTiMeOuTpOlIcYiDvAlUe",
			Expected: &PolicyActivityBasedTimeoutPolicyId{
				ActivityBasedTimeoutPolicyId: "aCtIvItYbAsEdTiMeOuTpOlIcYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aCtIvItYbAsEdTiMeOuTpOlIcIeS/aCtIvItYbAsEdTiMeOuTpOlIcYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyActivityBasedTimeoutPolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ActivityBasedTimeoutPolicyId != v.Expected.ActivityBasedTimeoutPolicyId {
			t.Fatalf("Expected %q but got %q for ActivityBasedTimeoutPolicyId", v.Expected.ActivityBasedTimeoutPolicyId, actual.ActivityBasedTimeoutPolicyId)
		}

	}
}

func TestSegmentsForPolicyActivityBasedTimeoutPolicyId(t *testing.T) {
	segments := PolicyActivityBasedTimeoutPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyActivityBasedTimeoutPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
