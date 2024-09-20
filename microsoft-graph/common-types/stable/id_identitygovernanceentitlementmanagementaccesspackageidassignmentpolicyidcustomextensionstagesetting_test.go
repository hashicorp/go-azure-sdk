package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingID("accessPackageId", "accessPackageAssignmentPolicyId", "customExtensionStageSettingId")

	if id.AccessPackageId != "accessPackageId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageId'", id.AccessPackageId, "accessPackageId")
	}

	if id.AccessPackageAssignmentPolicyId != "accessPackageAssignmentPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageAssignmentPolicyId'", id.AccessPackageAssignmentPolicyId, "accessPackageAssignmentPolicyId")
	}

	if id.CustomExtensionStageSettingId != "customExtensionStageSettingId" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomExtensionStageSettingId'", id.CustomExtensionStageSettingId, "customExtensionStageSettingId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingID("accessPackageId", "accessPackageAssignmentPolicyId", "customExtensionStageSettingId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/assignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings/customExtensionStageSettingId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId
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
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/assignmentPolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/assignmentPolicies/accessPackageAssignmentPolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/assignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/assignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings/customExtensionStageSettingId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId{
				AccessPackageId:                 "accessPackageId",
				AccessPackageAssignmentPolicyId: "accessPackageAssignmentPolicyId",
				CustomExtensionStageSettingId:   "customExtensionStageSettingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/assignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings/customExtensionStageSettingId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingID(v.Input)
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

		if actual.CustomExtensionStageSettingId != v.Expected.CustomExtensionStageSettingId {
			t.Fatalf("Expected %q but got %q for CustomExtensionStageSettingId", v.Expected.CustomExtensionStageSettingId, actual.CustomExtensionStageSettingId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId
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
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/assignmentPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/aSsIgNmEnTpOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/assignmentPolicies/accessPackageAssignmentPolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/aSsIgNmEnTpOlIcIeS/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/assignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/aSsIgNmEnTpOlIcIeS/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId/cUsToMeXtEnSiOnStAgEsEtTiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/assignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings/customExtensionStageSettingId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId{
				AccessPackageId:                 "accessPackageId",
				AccessPackageAssignmentPolicyId: "accessPackageAssignmentPolicyId",
				CustomExtensionStageSettingId:   "customExtensionStageSettingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/assignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings/customExtensionStageSettingId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/aSsIgNmEnTpOlIcIeS/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId/cUsToMeXtEnSiOnStAgEsEtTiNgS/cUsToMeXtEnSiOnStAgEsEtTiNgId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId{
				AccessPackageId:                 "aCcEsSpAcKaGeId",
				AccessPackageAssignmentPolicyId: "aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId",
				CustomExtensionStageSettingId:   "cUsToMeXtEnSiOnStAgEsEtTiNgId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/aSsIgNmEnTpOlIcIeS/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId/cUsToMeXtEnSiOnStAgEsEtTiNgS/cUsToMeXtEnSiOnStAgEsEtTiNgId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingIDInsensitively(v.Input)
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

		if actual.CustomExtensionStageSettingId != v.Expected.CustomExtensionStageSettingId {
			t.Fatalf("Expected %q but got %q for CustomExtensionStageSettingId", v.Expected.CustomExtensionStageSettingId, actual.CustomExtensionStageSettingId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
