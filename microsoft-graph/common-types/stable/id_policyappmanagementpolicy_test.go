package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyAppManagementPolicyId{}

func TestNewPolicyAppManagementPolicyID(t *testing.T) {
	id := NewPolicyAppManagementPolicyID("appManagementPolicyId")

	if id.AppManagementPolicyId != "appManagementPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'AppManagementPolicyId'", id.AppManagementPolicyId, "appManagementPolicyId")
	}
}

func TestFormatPolicyAppManagementPolicyID(t *testing.T) {
	actual := NewPolicyAppManagementPolicyID("appManagementPolicyId").ID()
	expected := "/policies/appManagementPolicies/appManagementPolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyAppManagementPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyAppManagementPolicyId
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
			Input: "/policies/appManagementPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/appManagementPolicies/appManagementPolicyId",
			Expected: &PolicyAppManagementPolicyId{
				AppManagementPolicyId: "appManagementPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/appManagementPolicies/appManagementPolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyAppManagementPolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AppManagementPolicyId != v.Expected.AppManagementPolicyId {
			t.Fatalf("Expected %q but got %q for AppManagementPolicyId", v.Expected.AppManagementPolicyId, actual.AppManagementPolicyId)
		}

	}
}

func TestParsePolicyAppManagementPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyAppManagementPolicyId
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
			Input: "/policies/appManagementPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aPpMaNaGeMeNtPoLiCiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/appManagementPolicies/appManagementPolicyId",
			Expected: &PolicyAppManagementPolicyId{
				AppManagementPolicyId: "appManagementPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/appManagementPolicies/appManagementPolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aPpMaNaGeMeNtPoLiCiEs/aPpMaNaGeMeNtPoLiCyId",
			Expected: &PolicyAppManagementPolicyId{
				AppManagementPolicyId: "aPpMaNaGeMeNtPoLiCyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aPpMaNaGeMeNtPoLiCiEs/aPpMaNaGeMeNtPoLiCyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyAppManagementPolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AppManagementPolicyId != v.Expected.AppManagementPolicyId {
			t.Fatalf("Expected %q but got %q for AppManagementPolicyId", v.Expected.AppManagementPolicyId, actual.AppManagementPolicyId)
		}

	}
}

func TestSegmentsForPolicyAppManagementPolicyId(t *testing.T) {
	segments := PolicyAppManagementPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyAppManagementPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
