package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyPermissionGrantPolicyIdIncludeId{}

func TestNewPolicyPermissionGrantPolicyIdIncludeID(t *testing.T) {
	id := NewPolicyPermissionGrantPolicyIdIncludeID("permissionGrantPolicyId", "permissionGrantConditionSetId")

	if id.PermissionGrantPolicyId != "permissionGrantPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'PermissionGrantPolicyId'", id.PermissionGrantPolicyId, "permissionGrantPolicyId")
	}

	if id.PermissionGrantConditionSetId != "permissionGrantConditionSetId" {
		t.Fatalf("Expected %q but got %q for Segment 'PermissionGrantConditionSetId'", id.PermissionGrantConditionSetId, "permissionGrantConditionSetId")
	}
}

func TestFormatPolicyPermissionGrantPolicyIdIncludeID(t *testing.T) {
	actual := NewPolicyPermissionGrantPolicyIdIncludeID("permissionGrantPolicyId", "permissionGrantConditionSetId").ID()
	expected := "/policies/permissionGrantPolicies/permissionGrantPolicyId/includes/permissionGrantConditionSetId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyPermissionGrantPolicyIdIncludeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyPermissionGrantPolicyIdIncludeId
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
			Input: "/policies/permissionGrantPolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyId/includes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyId/includes/permissionGrantConditionSetId",
			Expected: &PolicyPermissionGrantPolicyIdIncludeId{
				PermissionGrantPolicyId:       "permissionGrantPolicyId",
				PermissionGrantConditionSetId: "permissionGrantConditionSetId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyId/includes/permissionGrantConditionSetId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyPermissionGrantPolicyIdIncludeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PermissionGrantPolicyId != v.Expected.PermissionGrantPolicyId {
			t.Fatalf("Expected %q but got %q for PermissionGrantPolicyId", v.Expected.PermissionGrantPolicyId, actual.PermissionGrantPolicyId)
		}

		if actual.PermissionGrantConditionSetId != v.Expected.PermissionGrantConditionSetId {
			t.Fatalf("Expected %q but got %q for PermissionGrantConditionSetId", v.Expected.PermissionGrantConditionSetId, actual.PermissionGrantConditionSetId)
		}

	}
}

func TestParsePolicyPermissionGrantPolicyIdIncludeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyPermissionGrantPolicyIdIncludeId
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
			Input: "/policies/permissionGrantPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/pErMiSsIoNgRaNtPoLiCiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/pErMiSsIoNgRaNtPoLiCiEs/pErMiSsIoNgRaNtPoLiCyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyId/includes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/pErMiSsIoNgRaNtPoLiCiEs/pErMiSsIoNgRaNtPoLiCyId/iNcLuDeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyId/includes/permissionGrantConditionSetId",
			Expected: &PolicyPermissionGrantPolicyIdIncludeId{
				PermissionGrantPolicyId:       "permissionGrantPolicyId",
				PermissionGrantConditionSetId: "permissionGrantConditionSetId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyId/includes/permissionGrantConditionSetId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/pErMiSsIoNgRaNtPoLiCiEs/pErMiSsIoNgRaNtPoLiCyId/iNcLuDeS/pErMiSsIoNgRaNtCoNdItIoNsEtId",
			Expected: &PolicyPermissionGrantPolicyIdIncludeId{
				PermissionGrantPolicyId:       "pErMiSsIoNgRaNtPoLiCyId",
				PermissionGrantConditionSetId: "pErMiSsIoNgRaNtCoNdItIoNsEtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/pErMiSsIoNgRaNtPoLiCiEs/pErMiSsIoNgRaNtPoLiCyId/iNcLuDeS/pErMiSsIoNgRaNtCoNdItIoNsEtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyPermissionGrantPolicyIdIncludeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PermissionGrantPolicyId != v.Expected.PermissionGrantPolicyId {
			t.Fatalf("Expected %q but got %q for PermissionGrantPolicyId", v.Expected.PermissionGrantPolicyId, actual.PermissionGrantPolicyId)
		}

		if actual.PermissionGrantConditionSetId != v.Expected.PermissionGrantConditionSetId {
			t.Fatalf("Expected %q but got %q for PermissionGrantConditionSetId", v.Expected.PermissionGrantConditionSetId, actual.PermissionGrantConditionSetId)
		}

	}
}

func TestSegmentsForPolicyPermissionGrantPolicyIdIncludeId(t *testing.T) {
	segments := PolicyPermissionGrantPolicyIdIncludeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyPermissionGrantPolicyIdIncludeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
