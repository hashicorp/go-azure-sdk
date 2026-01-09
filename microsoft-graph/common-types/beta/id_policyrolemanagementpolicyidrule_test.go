package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyRoleManagementPolicyIdRuleId{}

func TestNewPolicyRoleManagementPolicyIdRuleID(t *testing.T) {
	id := NewPolicyRoleManagementPolicyIdRuleID("unifiedRoleManagementPolicyId", "unifiedRoleManagementPolicyRuleId")

	if id.UnifiedRoleManagementPolicyId != "unifiedRoleManagementPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleManagementPolicyId'", id.UnifiedRoleManagementPolicyId, "unifiedRoleManagementPolicyId")
	}

	if id.UnifiedRoleManagementPolicyRuleId != "unifiedRoleManagementPolicyRuleId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleManagementPolicyRuleId'", id.UnifiedRoleManagementPolicyRuleId, "unifiedRoleManagementPolicyRuleId")
	}
}

func TestFormatPolicyRoleManagementPolicyIdRuleID(t *testing.T) {
	actual := NewPolicyRoleManagementPolicyIdRuleID("unifiedRoleManagementPolicyId", "unifiedRoleManagementPolicyRuleId").ID()
	expected := "/policies/roleManagementPolicies/unifiedRoleManagementPolicyId/rules/unifiedRoleManagementPolicyRuleId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyRoleManagementPolicyIdRuleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyRoleManagementPolicyIdRuleId
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
			// Incomplete URI
			Input: "/policies/roleManagementPolicies/unifiedRoleManagementPolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/roleManagementPolicies/unifiedRoleManagementPolicyId/rules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/roleManagementPolicies/unifiedRoleManagementPolicyId/rules/unifiedRoleManagementPolicyRuleId",
			Expected: &PolicyRoleManagementPolicyIdRuleId{
				UnifiedRoleManagementPolicyId:     "unifiedRoleManagementPolicyId",
				UnifiedRoleManagementPolicyRuleId: "unifiedRoleManagementPolicyRuleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/roleManagementPolicies/unifiedRoleManagementPolicyId/rules/unifiedRoleManagementPolicyRuleId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyRoleManagementPolicyIdRuleID(v.Input)
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

		if actual.UnifiedRoleManagementPolicyRuleId != v.Expected.UnifiedRoleManagementPolicyRuleId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleManagementPolicyRuleId", v.Expected.UnifiedRoleManagementPolicyRuleId, actual.UnifiedRoleManagementPolicyRuleId)
		}

	}
}

func TestParsePolicyRoleManagementPolicyIdRuleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyRoleManagementPolicyIdRuleId
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
			// Incomplete URI
			Input: "/policies/roleManagementPolicies/unifiedRoleManagementPolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/rOlEmAnAgEmEnTpOlIcIeS/uNiFiEdRoLeMaNaGeMeNtPoLiCyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/roleManagementPolicies/unifiedRoleManagementPolicyId/rules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/rOlEmAnAgEmEnTpOlIcIeS/uNiFiEdRoLeMaNaGeMeNtPoLiCyId/rUlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/roleManagementPolicies/unifiedRoleManagementPolicyId/rules/unifiedRoleManagementPolicyRuleId",
			Expected: &PolicyRoleManagementPolicyIdRuleId{
				UnifiedRoleManagementPolicyId:     "unifiedRoleManagementPolicyId",
				UnifiedRoleManagementPolicyRuleId: "unifiedRoleManagementPolicyRuleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/roleManagementPolicies/unifiedRoleManagementPolicyId/rules/unifiedRoleManagementPolicyRuleId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/rOlEmAnAgEmEnTpOlIcIeS/uNiFiEdRoLeMaNaGeMeNtPoLiCyId/rUlEs/uNiFiEdRoLeMaNaGeMeNtPoLiCyRuLeId",
			Expected: &PolicyRoleManagementPolicyIdRuleId{
				UnifiedRoleManagementPolicyId:     "uNiFiEdRoLeMaNaGeMeNtPoLiCyId",
				UnifiedRoleManagementPolicyRuleId: "uNiFiEdRoLeMaNaGeMeNtPoLiCyRuLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/rOlEmAnAgEmEnTpOlIcIeS/uNiFiEdRoLeMaNaGeMeNtPoLiCyId/rUlEs/uNiFiEdRoLeMaNaGeMeNtPoLiCyRuLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyRoleManagementPolicyIdRuleIDInsensitively(v.Input)
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

		if actual.UnifiedRoleManagementPolicyRuleId != v.Expected.UnifiedRoleManagementPolicyRuleId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleManagementPolicyRuleId", v.Expected.UnifiedRoleManagementPolicyRuleId, actual.UnifiedRoleManagementPolicyRuleId)
		}

	}
}

func TestSegmentsForPolicyRoleManagementPolicyIdRuleId(t *testing.T) {
	segments := PolicyRoleManagementPolicyIdRuleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyRoleManagementPolicyIdRuleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
