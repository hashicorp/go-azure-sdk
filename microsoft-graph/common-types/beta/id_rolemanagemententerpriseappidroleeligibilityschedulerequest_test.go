package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId{}

func TestNewRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestID(t *testing.T) {
	id := NewRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestID("rbacApplicationIdValue", "unifiedRoleEligibilityScheduleRequestIdValue")

	if id.RbacApplicationId != "rbacApplicationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RbacApplicationId'", id.RbacApplicationId, "rbacApplicationIdValue")
	}

	if id.UnifiedRoleEligibilityScheduleRequestId != "unifiedRoleEligibilityScheduleRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleEligibilityScheduleRequestId'", id.UnifiedRoleEligibilityScheduleRequestId, "unifiedRoleEligibilityScheduleRequestIdValue")
	}
}

func TestFormatRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestID(t *testing.T) {
	actual := NewRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestID("rbacApplicationIdValue", "unifiedRoleEligibilityScheduleRequestIdValue").ID()
	expected := "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleEligibilityScheduleRequests/unifiedRoleEligibilityScheduleRequestIdValue"
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
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleEligibilityScheduleRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleEligibilityScheduleRequests/unifiedRoleEligibilityScheduleRequestIdValue",
			Expected: &RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId{
				RbacApplicationId:                       "rbacApplicationIdValue",
				UnifiedRoleEligibilityScheduleRequestId: "unifiedRoleEligibilityScheduleRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleEligibilityScheduleRequests/unifiedRoleEligibilityScheduleRequestIdValue/extra",
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
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleEligibilityScheduleRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEeLiGiBiLiTyScHeDuLeReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleEligibilityScheduleRequests/unifiedRoleEligibilityScheduleRequestIdValue",
			Expected: &RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId{
				RbacApplicationId:                       "rbacApplicationIdValue",
				UnifiedRoleEligibilityScheduleRequestId: "unifiedRoleEligibilityScheduleRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleEligibilityScheduleRequests/unifiedRoleEligibilityScheduleRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEeLiGiBiLiTyScHeDuLeReQuEsTs/uNiFiEdRoLeElIgIbIlItYsChEdUlErEqUeStIdVaLuE",
			Expected: &RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId{
				RbacApplicationId:                       "rBaCaPpLiCaTiOnIdVaLuE",
				UnifiedRoleEligibilityScheduleRequestId: "uNiFiEdRoLeElIgIbIlItYsChEdUlErEqUeStIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEeLiGiBiLiTyScHeDuLeReQuEsTs/uNiFiEdRoLeElIgIbIlItYsChEdUlErEqUeStIdVaLuE/extra",
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
