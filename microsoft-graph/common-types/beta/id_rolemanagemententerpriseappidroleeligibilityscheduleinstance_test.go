package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId{}

func TestNewRoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceID(t *testing.T) {
	id := NewRoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceID("rbacApplicationId", "unifiedRoleEligibilityScheduleInstanceId")

	if id.RbacApplicationId != "rbacApplicationId" {
		t.Fatalf("Expected %q but got %q for Segment 'RbacApplicationId'", id.RbacApplicationId, "rbacApplicationId")
	}

	if id.UnifiedRoleEligibilityScheduleInstanceId != "unifiedRoleEligibilityScheduleInstanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleEligibilityScheduleInstanceId'", id.UnifiedRoleEligibilityScheduleInstanceId, "unifiedRoleEligibilityScheduleInstanceId")
	}
}

func TestFormatRoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceID(t *testing.T) {
	actual := NewRoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceID("rbacApplicationId", "unifiedRoleEligibilityScheduleInstanceId").ID()
	expected := "/roleManagement/enterpriseApps/rbacApplicationId/roleEligibilityScheduleInstances/unifiedRoleEligibilityScheduleInstanceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId
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
			Input: "/roleManagement/enterpriseApps/rbacApplicationId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/roleEligibilityScheduleInstances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/roleEligibilityScheduleInstances/unifiedRoleEligibilityScheduleInstanceId",
			Expected: &RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId{
				RbacApplicationId:                        "rbacApplicationId",
				UnifiedRoleEligibilityScheduleInstanceId: "unifiedRoleEligibilityScheduleInstanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/roleEligibilityScheduleInstances/unifiedRoleEligibilityScheduleInstanceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceID(v.Input)
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

		if actual.UnifiedRoleEligibilityScheduleInstanceId != v.Expected.UnifiedRoleEligibilityScheduleInstanceId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleEligibilityScheduleInstanceId", v.Expected.UnifiedRoleEligibilityScheduleInstanceId, actual.UnifiedRoleEligibilityScheduleInstanceId)
		}

	}
}

func TestParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId
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
			Input: "/roleManagement/enterpriseApps/rbacApplicationId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/roleEligibilityScheduleInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnId/rOlEeLiGiBiLiTyScHeDuLeInStAnCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/roleEligibilityScheduleInstances/unifiedRoleEligibilityScheduleInstanceId",
			Expected: &RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId{
				RbacApplicationId:                        "rbacApplicationId",
				UnifiedRoleEligibilityScheduleInstanceId: "unifiedRoleEligibilityScheduleInstanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/roleEligibilityScheduleInstances/unifiedRoleEligibilityScheduleInstanceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnId/rOlEeLiGiBiLiTyScHeDuLeInStAnCeS/uNiFiEdRoLeElIgIbIlItYsChEdUlEiNsTaNcEiD",
			Expected: &RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId{
				RbacApplicationId:                        "rBaCaPpLiCaTiOnId",
				UnifiedRoleEligibilityScheduleInstanceId: "uNiFiEdRoLeElIgIbIlItYsChEdUlEiNsTaNcEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnId/rOlEeLiGiBiLiTyScHeDuLeInStAnCeS/uNiFiEdRoLeElIgIbIlItYsChEdUlEiNsTaNcEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceIDInsensitively(v.Input)
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

		if actual.UnifiedRoleEligibilityScheduleInstanceId != v.Expected.UnifiedRoleEligibilityScheduleInstanceId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleEligibilityScheduleInstanceId", v.Expected.UnifiedRoleEligibilityScheduleInstanceId, actual.UnifiedRoleEligibilityScheduleInstanceId)
		}

	}
}

func TestSegmentsForRoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId(t *testing.T) {
	segments := RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
