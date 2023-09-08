package policyrolemanagementpolicyrule

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyRoleManagementPolicyRuleId{}

func TestNewPolicyRoleManagementPolicyRuleID(t *testing.T) {
	id := NewPolicyRoleManagementPolicyRuleID("unifiedRoleManagementPolicyIdValue", "unifiedRoleManagementPolicyRuleIdValue")

	if id.UnifiedRoleManagementPolicyId != "unifiedRoleManagementPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleManagementPolicyId'", id.UnifiedRoleManagementPolicyId, "unifiedRoleManagementPolicyIdValue")
	}

	if id.UnifiedRoleManagementPolicyRuleId != "unifiedRoleManagementPolicyRuleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleManagementPolicyRuleId'", id.UnifiedRoleManagementPolicyRuleId, "unifiedRoleManagementPolicyRuleIdValue")
	}
}

func TestFormatPolicyRoleManagementPolicyRuleID(t *testing.T) {
	actual := NewPolicyRoleManagementPolicyRuleID("unifiedRoleManagementPolicyIdValue", "unifiedRoleManagementPolicyRuleIdValue").ID()
	expected := "/policies/roleManagementPolicies/unifiedRoleManagementPolicyIdValue/rules/unifiedRoleManagementPolicyRuleIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyRoleManagementPolicyRuleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyRoleManagementPolicyRuleId
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
			Input: "/policies/roleManagementPolicies/unifiedRoleManagementPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/roleManagementPolicies/unifiedRoleManagementPolicyIdValue/rules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/roleManagementPolicies/unifiedRoleManagementPolicyIdValue/rules/unifiedRoleManagementPolicyRuleIdValue",
			Expected: &PolicyRoleManagementPolicyRuleId{
				UnifiedRoleManagementPolicyId:     "unifiedRoleManagementPolicyIdValue",
				UnifiedRoleManagementPolicyRuleId: "unifiedRoleManagementPolicyRuleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/roleManagementPolicies/unifiedRoleManagementPolicyIdValue/rules/unifiedRoleManagementPolicyRuleIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyRoleManagementPolicyRuleID(v.Input)
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

func TestParsePolicyRoleManagementPolicyRuleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyRoleManagementPolicyRuleId
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
			Input: "/policies/roleManagementPolicies/unifiedRoleManagementPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/rOlEmAnAgEmEnTpOlIcIeS/uNiFiEdRoLeMaNaGeMeNtPoLiCyIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/roleManagementPolicies/unifiedRoleManagementPolicyIdValue/rules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/rOlEmAnAgEmEnTpOlIcIeS/uNiFiEdRoLeMaNaGeMeNtPoLiCyIdVaLuE/rUlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/roleManagementPolicies/unifiedRoleManagementPolicyIdValue/rules/unifiedRoleManagementPolicyRuleIdValue",
			Expected: &PolicyRoleManagementPolicyRuleId{
				UnifiedRoleManagementPolicyId:     "unifiedRoleManagementPolicyIdValue",
				UnifiedRoleManagementPolicyRuleId: "unifiedRoleManagementPolicyRuleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/roleManagementPolicies/unifiedRoleManagementPolicyIdValue/rules/unifiedRoleManagementPolicyRuleIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/rOlEmAnAgEmEnTpOlIcIeS/uNiFiEdRoLeMaNaGeMeNtPoLiCyIdVaLuE/rUlEs/uNiFiEdRoLeMaNaGeMeNtPoLiCyRuLeIdVaLuE",
			Expected: &PolicyRoleManagementPolicyRuleId{
				UnifiedRoleManagementPolicyId:     "uNiFiEdRoLeMaNaGeMeNtPoLiCyIdVaLuE",
				UnifiedRoleManagementPolicyRuleId: "uNiFiEdRoLeMaNaGeMeNtPoLiCyRuLeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/rOlEmAnAgEmEnTpOlIcIeS/uNiFiEdRoLeMaNaGeMeNtPoLiCyIdVaLuE/rUlEs/uNiFiEdRoLeMaNaGeMeNtPoLiCyRuLeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyRoleManagementPolicyRuleIDInsensitively(v.Input)
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

func TestSegmentsForPolicyRoleManagementPolicyRuleId(t *testing.T) {
	segments := PolicyRoleManagementPolicyRuleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyRoleManagementPolicyRuleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
