package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyPermissionGrantPolicyId{}

func TestNewPolicyPermissionGrantPolicyID(t *testing.T) {
	id := NewPolicyPermissionGrantPolicyID("permissionGrantPolicyId")

	if id.PermissionGrantPolicyId != "permissionGrantPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'PermissionGrantPolicyId'", id.PermissionGrantPolicyId, "permissionGrantPolicyId")
	}
}

func TestFormatPolicyPermissionGrantPolicyID(t *testing.T) {
	actual := NewPolicyPermissionGrantPolicyID("permissionGrantPolicyId").ID()
	expected := "/policies/permissionGrantPolicies/permissionGrantPolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyPermissionGrantPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyPermissionGrantPolicyId
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
			// Valid URI
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyId",
			Expected: &PolicyPermissionGrantPolicyId{
				PermissionGrantPolicyId: "permissionGrantPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyPermissionGrantPolicyID(v.Input)
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

	}
}

func TestParsePolicyPermissionGrantPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyPermissionGrantPolicyId
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
			// Valid URI
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyId",
			Expected: &PolicyPermissionGrantPolicyId{
				PermissionGrantPolicyId: "permissionGrantPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/pErMiSsIoNgRaNtPoLiCiEs/pErMiSsIoNgRaNtPoLiCyId",
			Expected: &PolicyPermissionGrantPolicyId{
				PermissionGrantPolicyId: "pErMiSsIoNgRaNtPoLiCyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/pErMiSsIoNgRaNtPoLiCiEs/pErMiSsIoNgRaNtPoLiCyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyPermissionGrantPolicyIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForPolicyPermissionGrantPolicyId(t *testing.T) {
	segments := PolicyPermissionGrantPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyPermissionGrantPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
