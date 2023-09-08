package rolemanagemententerpriseapproleeligibilityscheduleprincipal

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppRoleEligibilityScheduleId{}

func TestNewRoleManagementEnterpriseAppRoleEligibilityScheduleID(t *testing.T) {
	id := NewRoleManagementEnterpriseAppRoleEligibilityScheduleID("rbacApplicationIdValue", "unifiedRoleEligibilityScheduleIdValue")

	if id.RbacApplicationId != "rbacApplicationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RbacApplicationId'", id.RbacApplicationId, "rbacApplicationIdValue")
	}

	if id.UnifiedRoleEligibilityScheduleId != "unifiedRoleEligibilityScheduleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleEligibilityScheduleId'", id.UnifiedRoleEligibilityScheduleId, "unifiedRoleEligibilityScheduleIdValue")
	}
}

func TestFormatRoleManagementEnterpriseAppRoleEligibilityScheduleID(t *testing.T) {
	actual := NewRoleManagementEnterpriseAppRoleEligibilityScheduleID("rbacApplicationIdValue", "unifiedRoleEligibilityScheduleIdValue").ID()
	expected := "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleEligibilitySchedules/unifiedRoleEligibilityScheduleIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEnterpriseAppRoleEligibilityScheduleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppRoleEligibilityScheduleId
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
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleEligibilitySchedules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleEligibilitySchedules/unifiedRoleEligibilityScheduleIdValue",
			Expected: &RoleManagementEnterpriseAppRoleEligibilityScheduleId{
				RbacApplicationId:                "rbacApplicationIdValue",
				UnifiedRoleEligibilityScheduleId: "unifiedRoleEligibilityScheduleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleEligibilitySchedules/unifiedRoleEligibilityScheduleIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppRoleEligibilityScheduleID(v.Input)
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

		if actual.UnifiedRoleEligibilityScheduleId != v.Expected.UnifiedRoleEligibilityScheduleId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleEligibilityScheduleId", v.Expected.UnifiedRoleEligibilityScheduleId, actual.UnifiedRoleEligibilityScheduleId)
		}

	}
}

func TestParseRoleManagementEnterpriseAppRoleEligibilityScheduleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppRoleEligibilityScheduleId
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
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleEligibilitySchedules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEeLiGiBiLiTyScHeDuLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleEligibilitySchedules/unifiedRoleEligibilityScheduleIdValue",
			Expected: &RoleManagementEnterpriseAppRoleEligibilityScheduleId{
				RbacApplicationId:                "rbacApplicationIdValue",
				UnifiedRoleEligibilityScheduleId: "unifiedRoleEligibilityScheduleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleEligibilitySchedules/unifiedRoleEligibilityScheduleIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEeLiGiBiLiTyScHeDuLeS/uNiFiEdRoLeElIgIbIlItYsChEdUlEiDvAlUe",
			Expected: &RoleManagementEnterpriseAppRoleEligibilityScheduleId{
				RbacApplicationId:                "rBaCaPpLiCaTiOnIdVaLuE",
				UnifiedRoleEligibilityScheduleId: "uNiFiEdRoLeElIgIbIlItYsChEdUlEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEeLiGiBiLiTyScHeDuLeS/uNiFiEdRoLeElIgIbIlItYsChEdUlEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppRoleEligibilityScheduleIDInsensitively(v.Input)
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

		if actual.UnifiedRoleEligibilityScheduleId != v.Expected.UnifiedRoleEligibilityScheduleId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleEligibilityScheduleId", v.Expected.UnifiedRoleEligibilityScheduleId, actual.UnifiedRoleEligibilityScheduleId)
		}

	}
}

func TestSegmentsForRoleManagementEnterpriseAppRoleEligibilityScheduleId(t *testing.T) {
	segments := RoleManagementEnterpriseAppRoleEligibilityScheduleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEnterpriseAppRoleEligibilityScheduleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
