package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerID("accessPackageId", "accessPackageAssignmentPolicyId", "customExtensionHandlerId")

	if id.AccessPackageId != "accessPackageId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageId'", id.AccessPackageId, "accessPackageId")
	}

	if id.AccessPackageAssignmentPolicyId != "accessPackageAssignmentPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageAssignmentPolicyId'", id.AccessPackageAssignmentPolicyId, "accessPackageAssignmentPolicyId")
	}

	if id.CustomExtensionHandlerId != "customExtensionHandlerId" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomExtensionHandlerId'", id.CustomExtensionHandlerId, "customExtensionHandlerId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerID("accessPackageId", "accessPackageAssignmentPolicyId", "customExtensionHandlerId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId/customExtensionHandlers/customExtensionHandlerId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageAssignmentPolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId/customExtensionHandlers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId/customExtensionHandlers/customExtensionHandlerId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId{
				AccessPackageId:                 "accessPackageId",
				AccessPackageAssignmentPolicyId: "accessPackageAssignmentPolicyId",
				CustomExtensionHandlerId:        "customExtensionHandlerId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId/customExtensionHandlers/customExtensionHandlerId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageId != v.Expected.AccessPackageId {
			t.Fatalf("Expected %q but got %q for AccessPackageId", v.Expected.AccessPackageId, actual.AccessPackageId)
		}

		if actual.AccessPackageAssignmentPolicyId != v.Expected.AccessPackageAssignmentPolicyId {
			t.Fatalf("Expected %q but got %q for AccessPackageAssignmentPolicyId", v.Expected.AccessPackageAssignmentPolicyId, actual.AccessPackageAssignmentPolicyId)
		}

		if actual.CustomExtensionHandlerId != v.Expected.CustomExtensionHandlerId {
			t.Fatalf("Expected %q but got %q for CustomExtensionHandlerId", v.Expected.CustomExtensionHandlerId, actual.CustomExtensionHandlerId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageAssignmentPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/aCcEsSpAcKaGeAsSiGnMeNtPoLiCiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/aCcEsSpAcKaGeAsSiGnMeNtPoLiCiEs/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId/customExtensionHandlers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/aCcEsSpAcKaGeAsSiGnMeNtPoLiCiEs/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId/cUsToMeXtEnSiOnHaNdLeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId/customExtensionHandlers/customExtensionHandlerId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId{
				AccessPackageId:                 "accessPackageId",
				AccessPackageAssignmentPolicyId: "accessPackageAssignmentPolicyId",
				CustomExtensionHandlerId:        "customExtensionHandlerId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId/customExtensionHandlers/customExtensionHandlerId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/aCcEsSpAcKaGeAsSiGnMeNtPoLiCiEs/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId/cUsToMeXtEnSiOnHaNdLeRs/cUsToMeXtEnSiOnHaNdLeRiD",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId{
				AccessPackageId:                 "aCcEsSpAcKaGeId",
				AccessPackageAssignmentPolicyId: "aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId",
				CustomExtensionHandlerId:        "cUsToMeXtEnSiOnHaNdLeRiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/aCcEsSpAcKaGeAsSiGnMeNtPoLiCiEs/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId/cUsToMeXtEnSiOnHaNdLeRs/cUsToMeXtEnSiOnHaNdLeRiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageId != v.Expected.AccessPackageId {
			t.Fatalf("Expected %q but got %q for AccessPackageId", v.Expected.AccessPackageId, actual.AccessPackageId)
		}

		if actual.AccessPackageAssignmentPolicyId != v.Expected.AccessPackageAssignmentPolicyId {
			t.Fatalf("Expected %q but got %q for AccessPackageAssignmentPolicyId", v.Expected.AccessPackageAssignmentPolicyId, actual.AccessPackageAssignmentPolicyId)
		}

		if actual.CustomExtensionHandlerId != v.Expected.CustomExtensionHandlerId {
			t.Fatalf("Expected %q but got %q for CustomExtensionHandlerId", v.Expected.CustomExtensionHandlerId, actual.CustomExtensionHandlerId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
