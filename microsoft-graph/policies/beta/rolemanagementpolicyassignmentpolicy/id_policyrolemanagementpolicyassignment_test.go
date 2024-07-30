package rolemanagementpolicyassignmentpolicy

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyRoleManagementPolicyAssignmentId{}

func TestNewPolicyRoleManagementPolicyAssignmentID(t *testing.T) {
	id := NewPolicyRoleManagementPolicyAssignmentID("unifiedRoleManagementPolicyAssignmentIdValue")

	if id.UnifiedRoleManagementPolicyAssignmentId != "unifiedRoleManagementPolicyAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleManagementPolicyAssignmentId'", id.UnifiedRoleManagementPolicyAssignmentId, "unifiedRoleManagementPolicyAssignmentIdValue")
	}
}

func TestFormatPolicyRoleManagementPolicyAssignmentID(t *testing.T) {
	actual := NewPolicyRoleManagementPolicyAssignmentID("unifiedRoleManagementPolicyAssignmentIdValue").ID()
	expected := "/policies/roleManagementPolicyAssignments/unifiedRoleManagementPolicyAssignmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyRoleManagementPolicyAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyRoleManagementPolicyAssignmentId
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
			Input: "/policies/roleManagementPolicyAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/roleManagementPolicyAssignments/unifiedRoleManagementPolicyAssignmentIdValue",
			Expected: &PolicyRoleManagementPolicyAssignmentId{
				UnifiedRoleManagementPolicyAssignmentId: "unifiedRoleManagementPolicyAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/roleManagementPolicyAssignments/unifiedRoleManagementPolicyAssignmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyRoleManagementPolicyAssignmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleManagementPolicyAssignmentId != v.Expected.UnifiedRoleManagementPolicyAssignmentId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleManagementPolicyAssignmentId", v.Expected.UnifiedRoleManagementPolicyAssignmentId, actual.UnifiedRoleManagementPolicyAssignmentId)
		}

	}
}

func TestParsePolicyRoleManagementPolicyAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyRoleManagementPolicyAssignmentId
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
			Input: "/policies/roleManagementPolicyAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/rOlEmAnAgEmEnTpOlIcYaSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/roleManagementPolicyAssignments/unifiedRoleManagementPolicyAssignmentIdValue",
			Expected: &PolicyRoleManagementPolicyAssignmentId{
				UnifiedRoleManagementPolicyAssignmentId: "unifiedRoleManagementPolicyAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/roleManagementPolicyAssignments/unifiedRoleManagementPolicyAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/rOlEmAnAgEmEnTpOlIcYaSsIgNmEnTs/uNiFiEdRoLeMaNaGeMeNtPoLiCyAsSiGnMeNtIdVaLuE",
			Expected: &PolicyRoleManagementPolicyAssignmentId{
				UnifiedRoleManagementPolicyAssignmentId: "uNiFiEdRoLeMaNaGeMeNtPoLiCyAsSiGnMeNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/rOlEmAnAgEmEnTpOlIcYaSsIgNmEnTs/uNiFiEdRoLeMaNaGeMeNtPoLiCyAsSiGnMeNtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyRoleManagementPolicyAssignmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleManagementPolicyAssignmentId != v.Expected.UnifiedRoleManagementPolicyAssignmentId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleManagementPolicyAssignmentId", v.Expected.UnifiedRoleManagementPolicyAssignmentId, actual.UnifiedRoleManagementPolicyAssignmentId)
		}

	}
}

func TestSegmentsForPolicyRoleManagementPolicyAssignmentId(t *testing.T) {
	segments := PolicyRoleManagementPolicyAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyRoleManagementPolicyAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
