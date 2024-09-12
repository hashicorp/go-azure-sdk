package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId{}

func TestNewRoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestID(t *testing.T) {
	id := NewRoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestID("rbacApplicationIdValue", "unifiedRoleAssignmentScheduleRequestIdValue")

	if id.RbacApplicationId != "rbacApplicationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RbacApplicationId'", id.RbacApplicationId, "rbacApplicationIdValue")
	}

	if id.UnifiedRoleAssignmentScheduleRequestId != "unifiedRoleAssignmentScheduleRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleAssignmentScheduleRequestId'", id.UnifiedRoleAssignmentScheduleRequestId, "unifiedRoleAssignmentScheduleRequestIdValue")
	}
}

func TestFormatRoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestID(t *testing.T) {
	actual := NewRoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestID("rbacApplicationIdValue", "unifiedRoleAssignmentScheduleRequestIdValue").ID()
	expected := "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentScheduleRequests/unifiedRoleAssignmentScheduleRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentScheduleRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentScheduleRequests/unifiedRoleAssignmentScheduleRequestIdValue",
			Expected: &RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId{
				RbacApplicationId:                      "rbacApplicationIdValue",
				UnifiedRoleAssignmentScheduleRequestId: "unifiedRoleAssignmentScheduleRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentScheduleRequests/unifiedRoleAssignmentScheduleRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.RbacApplicationId != v.Expected.RbacApplicationId {
			t.Fatalf("Expected %q but got %q for RbacApplicationId", v.Expected.RbacApplicationId, actual.RbacApplicationId)
		}

		if actual.UnifiedRoleAssignmentScheduleRequestId != v.Expected.UnifiedRoleAssignmentScheduleRequestId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentScheduleRequestId", v.Expected.UnifiedRoleAssignmentScheduleRequestId, actual.UnifiedRoleAssignmentScheduleRequestId)
		}

	}
}

func TestParseRoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentScheduleRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEaSsIgNmEnTsChEdUlErEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentScheduleRequests/unifiedRoleAssignmentScheduleRequestIdValue",
			Expected: &RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId{
				RbacApplicationId:                      "rbacApplicationIdValue",
				UnifiedRoleAssignmentScheduleRequestId: "unifiedRoleAssignmentScheduleRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentScheduleRequests/unifiedRoleAssignmentScheduleRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEaSsIgNmEnTsChEdUlErEqUeStS/uNiFiEdRoLeAsSiGnMeNtScHeDuLeReQuEsTiDvAlUe",
			Expected: &RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId{
				RbacApplicationId:                      "rBaCaPpLiCaTiOnIdVaLuE",
				UnifiedRoleAssignmentScheduleRequestId: "uNiFiEdRoLeAsSiGnMeNtScHeDuLeReQuEsTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEaSsIgNmEnTsChEdUlErEqUeStS/uNiFiEdRoLeAsSiGnMeNtScHeDuLeReQuEsTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.RbacApplicationId != v.Expected.RbacApplicationId {
			t.Fatalf("Expected %q but got %q for RbacApplicationId", v.Expected.RbacApplicationId, actual.RbacApplicationId)
		}

		if actual.UnifiedRoleAssignmentScheduleRequestId != v.Expected.UnifiedRoleAssignmentScheduleRequestId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentScheduleRequestId", v.Expected.UnifiedRoleAssignmentScheduleRequestId, actual.UnifiedRoleAssignmentScheduleRequestId)
		}

	}
}

func TestSegmentsForRoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId(t *testing.T) {
	segments := RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
