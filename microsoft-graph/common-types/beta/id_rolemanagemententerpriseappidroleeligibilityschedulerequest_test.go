package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId{}

func TestNewRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestID(t *testing.T) {
	id := NewRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestID("rbacApplicationId", "unifiedRoleEligibilityScheduleRequestId")

	if id.RbacApplicationId != "rbacApplicationId" {
		t.Fatalf("Expected %q but got %q for Segment 'RbacApplicationId'", id.RbacApplicationId, "rbacApplicationId")
	}

	if id.UnifiedRoleEligibilityScheduleRequestId != "unifiedRoleEligibilityScheduleRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleEligibilityScheduleRequestId'", id.UnifiedRoleEligibilityScheduleRequestId, "unifiedRoleEligibilityScheduleRequestId")
	}
}

func TestFormatRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestID(t *testing.T) {
	actual := NewRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestID("rbacApplicationId", "unifiedRoleEligibilityScheduleRequestId").ID()
	expected := "/roleManagement/enterpriseApps/rbacApplicationId/roleEligibilityScheduleRequests/unifiedRoleEligibilityScheduleRequestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId
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
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/roleEligibilityScheduleRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/roleEligibilityScheduleRequests/unifiedRoleEligibilityScheduleRequestId",
			Expected: &RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId{
				RbacApplicationId:                       "rbacApplicationId",
				UnifiedRoleEligibilityScheduleRequestId: "unifiedRoleEligibilityScheduleRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/roleEligibilityScheduleRequests/unifiedRoleEligibilityScheduleRequestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestID(v.Input)
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

		if actual.UnifiedRoleEligibilityScheduleRequestId != v.Expected.UnifiedRoleEligibilityScheduleRequestId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleEligibilityScheduleRequestId", v.Expected.UnifiedRoleEligibilityScheduleRequestId, actual.UnifiedRoleEligibilityScheduleRequestId)
		}

	}
}

func TestParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId
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
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/roleEligibilityScheduleRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnId/rOlEeLiGiBiLiTyScHeDuLeReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/roleEligibilityScheduleRequests/unifiedRoleEligibilityScheduleRequestId",
			Expected: &RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId{
				RbacApplicationId:                       "rbacApplicationId",
				UnifiedRoleEligibilityScheduleRequestId: "unifiedRoleEligibilityScheduleRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/roleEligibilityScheduleRequests/unifiedRoleEligibilityScheduleRequestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnId/rOlEeLiGiBiLiTyScHeDuLeReQuEsTs/uNiFiEdRoLeElIgIbIlItYsChEdUlErEqUeStId",
			Expected: &RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId{
				RbacApplicationId:                       "rBaCaPpLiCaTiOnId",
				UnifiedRoleEligibilityScheduleRequestId: "uNiFiEdRoLeElIgIbIlItYsChEdUlErEqUeStId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnId/rOlEeLiGiBiLiTyScHeDuLeReQuEsTs/uNiFiEdRoLeElIgIbIlItYsChEdUlErEqUeStId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestIDInsensitively(v.Input)
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

		if actual.UnifiedRoleEligibilityScheduleRequestId != v.Expected.UnifiedRoleEligibilityScheduleRequestId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleEligibilityScheduleRequestId", v.Expected.UnifiedRoleEligibilityScheduleRequestId, actual.UnifiedRoleEligibilityScheduleRequestId)
		}

	}
}

func TestSegmentsForRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId(t *testing.T) {
	segments := RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
