package rolemanagemententerpriseapproleassignmentscheduleinstanceactivatedusing

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId{}

func TestNewRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceID(t *testing.T) {
	id := NewRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceID("rbacApplicationIdValue", "unifiedRoleAssignmentScheduleInstanceIdValue")

	if id.RbacApplicationId != "rbacApplicationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RbacApplicationId'", id.RbacApplicationId, "rbacApplicationIdValue")
	}

	if id.UnifiedRoleAssignmentScheduleInstanceId != "unifiedRoleAssignmentScheduleInstanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleAssignmentScheduleInstanceId'", id.UnifiedRoleAssignmentScheduleInstanceId, "unifiedRoleAssignmentScheduleInstanceIdValue")
	}
}

func TestFormatRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceID(t *testing.T) {
	actual := NewRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceID("rbacApplicationIdValue", "unifiedRoleAssignmentScheduleInstanceIdValue").ID()
	expected := "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentScheduleInstances/unifiedRoleAssignmentScheduleInstanceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId
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
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentScheduleInstances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentScheduleInstances/unifiedRoleAssignmentScheduleInstanceIdValue",
			Expected: &RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId{
				RbacApplicationId:                       "rbacApplicationIdValue",
				UnifiedRoleAssignmentScheduleInstanceId: "unifiedRoleAssignmentScheduleInstanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentScheduleInstances/unifiedRoleAssignmentScheduleInstanceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceID(v.Input)
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

		if actual.UnifiedRoleAssignmentScheduleInstanceId != v.Expected.UnifiedRoleAssignmentScheduleInstanceId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentScheduleInstanceId", v.Expected.UnifiedRoleAssignmentScheduleInstanceId, actual.UnifiedRoleAssignmentScheduleInstanceId)
		}

	}
}

func TestParseRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId
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
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentScheduleInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEaSsIgNmEnTsChEdUlEiNsTaNcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentScheduleInstances/unifiedRoleAssignmentScheduleInstanceIdValue",
			Expected: &RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId{
				RbacApplicationId:                       "rbacApplicationIdValue",
				UnifiedRoleAssignmentScheduleInstanceId: "unifiedRoleAssignmentScheduleInstanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentScheduleInstances/unifiedRoleAssignmentScheduleInstanceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEaSsIgNmEnTsChEdUlEiNsTaNcEs/uNiFiEdRoLeAsSiGnMeNtScHeDuLeInStAnCeIdVaLuE",
			Expected: &RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId{
				RbacApplicationId:                       "rBaCaPpLiCaTiOnIdVaLuE",
				UnifiedRoleAssignmentScheduleInstanceId: "uNiFiEdRoLeAsSiGnMeNtScHeDuLeInStAnCeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEaSsIgNmEnTsChEdUlEiNsTaNcEs/uNiFiEdRoLeAsSiGnMeNtScHeDuLeInStAnCeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceIDInsensitively(v.Input)
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

		if actual.UnifiedRoleAssignmentScheduleInstanceId != v.Expected.UnifiedRoleAssignmentScheduleInstanceId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentScheduleInstanceId", v.Expected.UnifiedRoleAssignmentScheduleInstanceId, actual.UnifiedRoleAssignmentScheduleInstanceId)
		}

	}
}

func TestSegmentsForRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId(t *testing.T) {
	segments := RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
