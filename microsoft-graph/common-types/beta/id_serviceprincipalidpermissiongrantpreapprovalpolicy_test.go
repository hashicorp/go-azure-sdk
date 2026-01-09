package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdPermissionGrantPreApprovalPolicyId{}

func TestNewServicePrincipalIdPermissionGrantPreApprovalPolicyID(t *testing.T) {
	id := NewServicePrincipalIdPermissionGrantPreApprovalPolicyID("servicePrincipalId", "permissionGrantPreApprovalPolicyId")

	if id.ServicePrincipalId != "servicePrincipalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalId")
	}

	if id.PermissionGrantPreApprovalPolicyId != "permissionGrantPreApprovalPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'PermissionGrantPreApprovalPolicyId'", id.PermissionGrantPreApprovalPolicyId, "permissionGrantPreApprovalPolicyId")
	}
}

func TestFormatServicePrincipalIdPermissionGrantPreApprovalPolicyID(t *testing.T) {
	actual := NewServicePrincipalIdPermissionGrantPreApprovalPolicyID("servicePrincipalId", "permissionGrantPreApprovalPolicyId").ID()
	expected := "/servicePrincipals/servicePrincipalId/permissionGrantPreApprovalPolicies/permissionGrantPreApprovalPolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalIdPermissionGrantPreApprovalPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdPermissionGrantPreApprovalPolicyId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/permissionGrantPreApprovalPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/permissionGrantPreApprovalPolicies/permissionGrantPreApprovalPolicyId",
			Expected: &ServicePrincipalIdPermissionGrantPreApprovalPolicyId{
				ServicePrincipalId:                 "servicePrincipalId",
				PermissionGrantPreApprovalPolicyId: "permissionGrantPreApprovalPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/permissionGrantPreApprovalPolicies/permissionGrantPreApprovalPolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdPermissionGrantPreApprovalPolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ServicePrincipalId != v.Expected.ServicePrincipalId {
			t.Fatalf("Expected %q but got %q for ServicePrincipalId", v.Expected.ServicePrincipalId, actual.ServicePrincipalId)
		}

		if actual.PermissionGrantPreApprovalPolicyId != v.Expected.PermissionGrantPreApprovalPolicyId {
			t.Fatalf("Expected %q but got %q for PermissionGrantPreApprovalPolicyId", v.Expected.PermissionGrantPreApprovalPolicyId, actual.PermissionGrantPreApprovalPolicyId)
		}

	}
}

func TestParseServicePrincipalIdPermissionGrantPreApprovalPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdPermissionGrantPreApprovalPolicyId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/permissionGrantPreApprovalPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/pErMiSsIoNgRaNtPrEaPpRoVaLpOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/permissionGrantPreApprovalPolicies/permissionGrantPreApprovalPolicyId",
			Expected: &ServicePrincipalIdPermissionGrantPreApprovalPolicyId{
				ServicePrincipalId:                 "servicePrincipalId",
				PermissionGrantPreApprovalPolicyId: "permissionGrantPreApprovalPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/permissionGrantPreApprovalPolicies/permissionGrantPreApprovalPolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/pErMiSsIoNgRaNtPrEaPpRoVaLpOlIcIeS/pErMiSsIoNgRaNtPrEaPpRoVaLpOlIcYiD",
			Expected: &ServicePrincipalIdPermissionGrantPreApprovalPolicyId{
				ServicePrincipalId:                 "sErViCePrInCiPaLiD",
				PermissionGrantPreApprovalPolicyId: "pErMiSsIoNgRaNtPrEaPpRoVaLpOlIcYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/pErMiSsIoNgRaNtPrEaPpRoVaLpOlIcIeS/pErMiSsIoNgRaNtPrEaPpRoVaLpOlIcYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdPermissionGrantPreApprovalPolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ServicePrincipalId != v.Expected.ServicePrincipalId {
			t.Fatalf("Expected %q but got %q for ServicePrincipalId", v.Expected.ServicePrincipalId, actual.ServicePrincipalId)
		}

		if actual.PermissionGrantPreApprovalPolicyId != v.Expected.PermissionGrantPreApprovalPolicyId {
			t.Fatalf("Expected %q but got %q for PermissionGrantPreApprovalPolicyId", v.Expected.PermissionGrantPreApprovalPolicyId, actual.PermissionGrantPreApprovalPolicyId)
		}

	}
}

func TestSegmentsForServicePrincipalIdPermissionGrantPreApprovalPolicyId(t *testing.T) {
	segments := ServicePrincipalIdPermissionGrantPreApprovalPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalIdPermissionGrantPreApprovalPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
