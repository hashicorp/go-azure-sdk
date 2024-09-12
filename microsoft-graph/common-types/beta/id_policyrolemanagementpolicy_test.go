package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyRoleManagementPolicyId{}

func TestNewPolicyRoleManagementPolicyID(t *testing.T) {
	id := NewPolicyRoleManagementPolicyID("unifiedRoleManagementPolicyIdValue")

	if id.UnifiedRoleManagementPolicyId != "unifiedRoleManagementPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleManagementPolicyId'", id.UnifiedRoleManagementPolicyId, "unifiedRoleManagementPolicyIdValue")
	}
}

func TestFormatPolicyRoleManagementPolicyID(t *testing.T) {
	actual := NewPolicyRoleManagementPolicyID("unifiedRoleManagementPolicyIdValue").ID()
	expected := "/policies/roleManagementPolicies/unifiedRoleManagementPolicyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyRoleManagementPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyRoleManagementPolicyId
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
			Input: "/policies/roleManagementPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/roleManagementPolicies/unifiedRoleManagementPolicyIdValue",
			Expected: &PolicyRoleManagementPolicyId{
				UnifiedRoleManagementPolicyId: "unifiedRoleManagementPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/roleManagementPolicies/unifiedRoleManagementPolicyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyRoleManagementPolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleManagementPolicyId != v.Expected.UnifiedRoleManagementPolicyId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleManagementPolicyId", v.Expected.UnifiedRoleManagementPolicyId, actual.UnifiedRoleManagementPolicyId)
		}

	}
}

func TestParsePolicyRoleManagementPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyRoleManagementPolicyId
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
			Input: "/policies/roleManagementPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/rOlEmAnAgEmEnTpOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/roleManagementPolicies/unifiedRoleManagementPolicyIdValue",
			Expected: &PolicyRoleManagementPolicyId{
				UnifiedRoleManagementPolicyId: "unifiedRoleManagementPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/roleManagementPolicies/unifiedRoleManagementPolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/rOlEmAnAgEmEnTpOlIcIeS/uNiFiEdRoLeMaNaGeMeNtPoLiCyIdVaLuE",
			Expected: &PolicyRoleManagementPolicyId{
				UnifiedRoleManagementPolicyId: "uNiFiEdRoLeMaNaGeMeNtPoLiCyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/rOlEmAnAgEmEnTpOlIcIeS/uNiFiEdRoLeMaNaGeMeNtPoLiCyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyRoleManagementPolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleManagementPolicyId != v.Expected.UnifiedRoleManagementPolicyId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleManagementPolicyId", v.Expected.UnifiedRoleManagementPolicyId, actual.UnifiedRoleManagementPolicyId)
		}

	}
}

func TestSegmentsForPolicyRoleManagementPolicyId(t *testing.T) {
	segments := PolicyRoleManagementPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyRoleManagementPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
