package roleassignmentschedulerequests

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScopedRoleAssignmentScheduleRequestId{}

func TestNewScopedRoleAssignmentScheduleRequestID(t *testing.T) {
	id := NewScopedRoleAssignmentScheduleRequestID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "roleAssignmentScheduleRequestValue")

	if id.Scope != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'Scope'", id.Scope, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.RoleAssignmentScheduleRequestName != "roleAssignmentScheduleRequestValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleAssignmentScheduleRequestName'", id.RoleAssignmentScheduleRequestName, "roleAssignmentScheduleRequestValue")
	}
}

func TestFormatScopedRoleAssignmentScheduleRequestID(t *testing.T) {
	actual := NewScopedRoleAssignmentScheduleRequestID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "roleAssignmentScheduleRequestValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleAssignmentScheduleRequests/roleAssignmentScheduleRequestValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScopedRoleAssignmentScheduleRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedRoleAssignmentScheduleRequestId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleAssignmentScheduleRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleAssignmentScheduleRequests/roleAssignmentScheduleRequestValue",
			Expected: &ScopedRoleAssignmentScheduleRequestId{
				Scope:                             "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				RoleAssignmentScheduleRequestName: "roleAssignmentScheduleRequestValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleAssignmentScheduleRequests/roleAssignmentScheduleRequestValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedRoleAssignmentScheduleRequestID(v.Input)
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

		if actual.RoleAssignmentScheduleRequestName != v.Expected.RoleAssignmentScheduleRequestName {
			t.Fatalf("Expected %q but got %q for RoleAssignmentScheduleRequestName", v.Expected.RoleAssignmentScheduleRequestName, actual.RoleAssignmentScheduleRequestName)
		}

	}
}

func TestParseScopedRoleAssignmentScheduleRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedRoleAssignmentScheduleRequestId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleAssignmentScheduleRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/rOlEaSsIgNmEnTsChEdUlErEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleAssignmentScheduleRequests/roleAssignmentScheduleRequestValue",
			Expected: &ScopedRoleAssignmentScheduleRequestId{
				Scope:                             "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				RoleAssignmentScheduleRequestName: "roleAssignmentScheduleRequestValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleAssignmentScheduleRequests/roleAssignmentScheduleRequestValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/rOlEaSsIgNmEnTsChEdUlErEqUeStS/rOlEaSsIgNmEnTsChEdUlErEqUeStVaLuE",
			Expected: &ScopedRoleAssignmentScheduleRequestId{
				Scope:                             "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				RoleAssignmentScheduleRequestName: "rOlEaSsIgNmEnTsChEdUlErEqUeStVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/rOlEaSsIgNmEnTsChEdUlErEqUeStS/rOlEaSsIgNmEnTsChEdUlErEqUeStVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedRoleAssignmentScheduleRequestIDInsensitively(v.Input)
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

		if actual.RoleAssignmentScheduleRequestName != v.Expected.RoleAssignmentScheduleRequestName {
			t.Fatalf("Expected %q but got %q for RoleAssignmentScheduleRequestName", v.Expected.RoleAssignmentScheduleRequestName, actual.RoleAssignmentScheduleRequestName)
		}

	}
}

func TestSegmentsForScopedRoleAssignmentScheduleRequestId(t *testing.T) {
	segments := ScopedRoleAssignmentScheduleRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScopedRoleAssignmentScheduleRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
