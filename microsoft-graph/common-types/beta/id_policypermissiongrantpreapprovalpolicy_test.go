package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyPermissionGrantPreApprovalPolicyId{}

func TestNewPolicyPermissionGrantPreApprovalPolicyID(t *testing.T) {
	id := NewPolicyPermissionGrantPreApprovalPolicyID("permissionGrantPreApprovalPolicyIdValue")

	if id.PermissionGrantPreApprovalPolicyId != "permissionGrantPreApprovalPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PermissionGrantPreApprovalPolicyId'", id.PermissionGrantPreApprovalPolicyId, "permissionGrantPreApprovalPolicyIdValue")
	}
}

func TestFormatPolicyPermissionGrantPreApprovalPolicyID(t *testing.T) {
	actual := NewPolicyPermissionGrantPreApprovalPolicyID("permissionGrantPreApprovalPolicyIdValue").ID()
	expected := "/policies/permissionGrantPreApprovalPolicies/permissionGrantPreApprovalPolicyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyPermissionGrantPreApprovalPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyPermissionGrantPreApprovalPolicyId
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
			Input: "/policies/permissionGrantPreApprovalPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/permissionGrantPreApprovalPolicies/permissionGrantPreApprovalPolicyIdValue",
			Expected: &PolicyPermissionGrantPreApprovalPolicyId{
				PermissionGrantPreApprovalPolicyId: "permissionGrantPreApprovalPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/permissionGrantPreApprovalPolicies/permissionGrantPreApprovalPolicyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyPermissionGrantPreApprovalPolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PermissionGrantPreApprovalPolicyId != v.Expected.PermissionGrantPreApprovalPolicyId {
			t.Fatalf("Expected %q but got %q for PermissionGrantPreApprovalPolicyId", v.Expected.PermissionGrantPreApprovalPolicyId, actual.PermissionGrantPreApprovalPolicyId)
		}

	}
}

func TestParsePolicyPermissionGrantPreApprovalPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyPermissionGrantPreApprovalPolicyId
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
			Input: "/policies/permissionGrantPreApprovalPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/pErMiSsIoNgRaNtPrEaPpRoVaLpOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/permissionGrantPreApprovalPolicies/permissionGrantPreApprovalPolicyIdValue",
			Expected: &PolicyPermissionGrantPreApprovalPolicyId{
				PermissionGrantPreApprovalPolicyId: "permissionGrantPreApprovalPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/permissionGrantPreApprovalPolicies/permissionGrantPreApprovalPolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/pErMiSsIoNgRaNtPrEaPpRoVaLpOlIcIeS/pErMiSsIoNgRaNtPrEaPpRoVaLpOlIcYiDvAlUe",
			Expected: &PolicyPermissionGrantPreApprovalPolicyId{
				PermissionGrantPreApprovalPolicyId: "pErMiSsIoNgRaNtPrEaPpRoVaLpOlIcYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/pErMiSsIoNgRaNtPrEaPpRoVaLpOlIcIeS/pErMiSsIoNgRaNtPrEaPpRoVaLpOlIcYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyPermissionGrantPreApprovalPolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PermissionGrantPreApprovalPolicyId != v.Expected.PermissionGrantPreApprovalPolicyId {
			t.Fatalf("Expected %q but got %q for PermissionGrantPreApprovalPolicyId", v.Expected.PermissionGrantPreApprovalPolicyId, actual.PermissionGrantPreApprovalPolicyId)
		}

	}
}

func TestSegmentsForPolicyPermissionGrantPreApprovalPolicyId(t *testing.T) {
	segments := PolicyPermissionGrantPreApprovalPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyPermissionGrantPreApprovalPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
