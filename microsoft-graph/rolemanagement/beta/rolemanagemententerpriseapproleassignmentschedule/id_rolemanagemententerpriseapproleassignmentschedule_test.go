package rolemanagemententerpriseapproleassignmentschedule

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppRoleAssignmentScheduleId{}

func TestNewRoleManagementEnterpriseAppRoleAssignmentScheduleID(t *testing.T) {
	id := NewRoleManagementEnterpriseAppRoleAssignmentScheduleID("rbacApplicationIdValue", "unifiedRoleAssignmentScheduleIdValue")

	if id.RbacApplicationId != "rbacApplicationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RbacApplicationId'", id.RbacApplicationId, "rbacApplicationIdValue")
	}

	if id.UnifiedRoleAssignmentScheduleId != "unifiedRoleAssignmentScheduleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleAssignmentScheduleId'", id.UnifiedRoleAssignmentScheduleId, "unifiedRoleAssignmentScheduleIdValue")
	}
}

func TestFormatRoleManagementEnterpriseAppRoleAssignmentScheduleID(t *testing.T) {
	actual := NewRoleManagementEnterpriseAppRoleAssignmentScheduleID("rbacApplicationIdValue", "unifiedRoleAssignmentScheduleIdValue").ID()
	expected := "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentSchedules/unifiedRoleAssignmentScheduleIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEnterpriseAppRoleAssignmentScheduleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppRoleAssignmentScheduleId
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
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentSchedules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentSchedules/unifiedRoleAssignmentScheduleIdValue",
			Expected: &RoleManagementEnterpriseAppRoleAssignmentScheduleId{
				RbacApplicationId:               "rbacApplicationIdValue",
				UnifiedRoleAssignmentScheduleId: "unifiedRoleAssignmentScheduleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentSchedules/unifiedRoleAssignmentScheduleIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppRoleAssignmentScheduleID(v.Input)
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

		if actual.UnifiedRoleAssignmentScheduleId != v.Expected.UnifiedRoleAssignmentScheduleId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentScheduleId", v.Expected.UnifiedRoleAssignmentScheduleId, actual.UnifiedRoleAssignmentScheduleId)
		}

	}
}

func TestParseRoleManagementEnterpriseAppRoleAssignmentScheduleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppRoleAssignmentScheduleId
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
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentSchedules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEaSsIgNmEnTsChEdUlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentSchedules/unifiedRoleAssignmentScheduleIdValue",
			Expected: &RoleManagementEnterpriseAppRoleAssignmentScheduleId{
				RbacApplicationId:               "rbacApplicationIdValue",
				UnifiedRoleAssignmentScheduleId: "unifiedRoleAssignmentScheduleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentSchedules/unifiedRoleAssignmentScheduleIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEaSsIgNmEnTsChEdUlEs/uNiFiEdRoLeAsSiGnMeNtScHeDuLeIdVaLuE",
			Expected: &RoleManagementEnterpriseAppRoleAssignmentScheduleId{
				RbacApplicationId:               "rBaCaPpLiCaTiOnIdVaLuE",
				UnifiedRoleAssignmentScheduleId: "uNiFiEdRoLeAsSiGnMeNtScHeDuLeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEaSsIgNmEnTsChEdUlEs/uNiFiEdRoLeAsSiGnMeNtScHeDuLeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppRoleAssignmentScheduleIDInsensitively(v.Input)
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

		if actual.UnifiedRoleAssignmentScheduleId != v.Expected.UnifiedRoleAssignmentScheduleId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentScheduleId", v.Expected.UnifiedRoleAssignmentScheduleId, actual.UnifiedRoleAssignmentScheduleId)
		}

	}
}

func TestSegmentsForRoleManagementEnterpriseAppRoleAssignmentScheduleId(t *testing.T) {
	segments := RoleManagementEnterpriseAppRoleAssignmentScheduleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEnterpriseAppRoleAssignmentScheduleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
