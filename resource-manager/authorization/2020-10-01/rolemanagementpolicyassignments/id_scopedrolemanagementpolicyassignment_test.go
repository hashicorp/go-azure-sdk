// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolemanagementpolicyassignments

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ScopedRoleManagementPolicyAssignmentId{}

func TestNewScopedRoleManagementPolicyAssignmentID(t *testing.T) {
	id := NewScopedRoleManagementPolicyAssignmentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "roleManagementPolicyAssignmentValue")

	if id.Scope != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'Scope'", id.Scope, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.RoleManagementPolicyAssignmentName != "roleManagementPolicyAssignmentValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleManagementPolicyAssignmentName'", id.RoleManagementPolicyAssignmentName, "roleManagementPolicyAssignmentValue")
	}
}

func TestFormatScopedRoleManagementPolicyAssignmentID(t *testing.T) {
	actual := NewScopedRoleManagementPolicyAssignmentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "roleManagementPolicyAssignmentValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleManagementPolicyAssignments/roleManagementPolicyAssignmentValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScopedRoleManagementPolicyAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedRoleManagementPolicyAssignmentId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleManagementPolicyAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleManagementPolicyAssignments/roleManagementPolicyAssignmentValue",
			Expected: &ScopedRoleManagementPolicyAssignmentId{
				Scope:                              "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				RoleManagementPolicyAssignmentName: "roleManagementPolicyAssignmentValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleManagementPolicyAssignments/roleManagementPolicyAssignmentValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedRoleManagementPolicyAssignmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.Scope != v.Expected.Scope {
			t.Fatalf("Expected %q but got %q for Scope", v.Expected.Scope, actual.Scope)
		}

		if actual.RoleManagementPolicyAssignmentName != v.Expected.RoleManagementPolicyAssignmentName {
			t.Fatalf("Expected %q but got %q for RoleManagementPolicyAssignmentName", v.Expected.RoleManagementPolicyAssignmentName, actual.RoleManagementPolicyAssignmentName)
		}

	}
}

func TestParseScopedRoleManagementPolicyAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedRoleManagementPolicyAssignmentId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleManagementPolicyAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/rOlEmAnAgEmEnTpOlIcYaSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleManagementPolicyAssignments/roleManagementPolicyAssignmentValue",
			Expected: &ScopedRoleManagementPolicyAssignmentId{
				Scope:                              "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				RoleManagementPolicyAssignmentName: "roleManagementPolicyAssignmentValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleManagementPolicyAssignments/roleManagementPolicyAssignmentValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/rOlEmAnAgEmEnTpOlIcYaSsIgNmEnTs/rOlEmAnAgEmEnTpOlIcYaSsIgNmEnTvAlUe",
			Expected: &ScopedRoleManagementPolicyAssignmentId{
				Scope:                              "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				RoleManagementPolicyAssignmentName: "rOlEmAnAgEmEnTpOlIcYaSsIgNmEnTvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/rOlEmAnAgEmEnTpOlIcYaSsIgNmEnTs/rOlEmAnAgEmEnTpOlIcYaSsIgNmEnTvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedRoleManagementPolicyAssignmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.Scope != v.Expected.Scope {
			t.Fatalf("Expected %q but got %q for Scope", v.Expected.Scope, actual.Scope)
		}

		if actual.RoleManagementPolicyAssignmentName != v.Expected.RoleManagementPolicyAssignmentName {
			t.Fatalf("Expected %q but got %q for RoleManagementPolicyAssignmentName", v.Expected.RoleManagementPolicyAssignmentName, actual.RoleManagementPolicyAssignmentName)
		}

	}
}

func TestSegmentsForScopedRoleManagementPolicyAssignmentId(t *testing.T) {
	segments := ScopedRoleManagementPolicyAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScopedRoleManagementPolicyAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
