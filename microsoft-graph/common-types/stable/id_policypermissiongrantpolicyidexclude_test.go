package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyPermissionGrantPolicyIdExcludeId{}

func TestNewPolicyPermissionGrantPolicyIdExcludeID(t *testing.T) {
	id := NewPolicyPermissionGrantPolicyIdExcludeID("permissionGrantPolicyId", "permissionGrantConditionSetId")

	if id.PermissionGrantPolicyId != "permissionGrantPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'PermissionGrantPolicyId'", id.PermissionGrantPolicyId, "permissionGrantPolicyId")
	}

	if id.PermissionGrantConditionSetId != "permissionGrantConditionSetId" {
		t.Fatalf("Expected %q but got %q for Segment 'PermissionGrantConditionSetId'", id.PermissionGrantConditionSetId, "permissionGrantConditionSetId")
	}
}

func TestFormatPolicyPermissionGrantPolicyIdExcludeID(t *testing.T) {
	actual := NewPolicyPermissionGrantPolicyIdExcludeID("permissionGrantPolicyId", "permissionGrantConditionSetId").ID()
	expected := "/policies/permissionGrantPolicies/permissionGrantPolicyId/excludes/permissionGrantConditionSetId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyPermissionGrantPolicyIdExcludeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyPermissionGrantPolicyIdExcludeId
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
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyId/excludes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyId/excludes/permissionGrantConditionSetId",
			Expected: &PolicyPermissionGrantPolicyIdExcludeId{
				PermissionGrantPolicyId:       "permissionGrantPolicyId",
				PermissionGrantConditionSetId: "permissionGrantConditionSetId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyId/excludes/permissionGrantConditionSetId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyPermissionGrantPolicyIdExcludeID(v.Input)
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

func TestParsePolicyPermissionGrantPolicyIdExcludeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyPermissionGrantPolicyIdExcludeId
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
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyId/excludes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/pErMiSsIoNgRaNtPoLiCiEs/pErMiSsIoNgRaNtPoLiCyId/eXcLuDeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyId/excludes/permissionGrantConditionSetId",
			Expected: &PolicyPermissionGrantPolicyIdExcludeId{
				PermissionGrantPolicyId:       "permissionGrantPolicyId",
				PermissionGrantConditionSetId: "permissionGrantConditionSetId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyId/excludes/permissionGrantConditionSetId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/pErMiSsIoNgRaNtPoLiCiEs/pErMiSsIoNgRaNtPoLiCyId/eXcLuDeS/pErMiSsIoNgRaNtCoNdItIoNsEtId",
			Expected: &PolicyPermissionGrantPolicyIdExcludeId{
				PermissionGrantPolicyId:       "pErMiSsIoNgRaNtPoLiCyId",
				PermissionGrantConditionSetId: "pErMiSsIoNgRaNtCoNdItIoNsEtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/pErMiSsIoNgRaNtPoLiCiEs/pErMiSsIoNgRaNtPoLiCyId/eXcLuDeS/pErMiSsIoNgRaNtCoNdItIoNsEtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyPermissionGrantPolicyIdExcludeIDInsensitively(v.Input)
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

func TestSegmentsForPolicyPermissionGrantPolicyIdExcludeId(t *testing.T) {
	segments := PolicyPermissionGrantPolicyIdExcludeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyPermissionGrantPolicyIdExcludeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
