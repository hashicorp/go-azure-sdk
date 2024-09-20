package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId{}

func TestNewIdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingID("accessPackageAssignmentPolicyId", "customExtensionStageSettingId")

	if id.AccessPackageAssignmentPolicyId != "accessPackageAssignmentPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageAssignmentPolicyId'", id.AccessPackageAssignmentPolicyId, "accessPackageAssignmentPolicyId")
	}

	if id.CustomExtensionStageSettingId != "customExtensionStageSettingId" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomExtensionStageSettingId'", id.CustomExtensionStageSettingId, "customExtensionStageSettingId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingID("accessPackageAssignmentPolicyId", "customExtensionStageSettingId").ID()
	expected := "/identityGovernance/entitlementManagement/assignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings/customExtensionStageSettingId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId
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
			Input: "/identityGovernance/entitlementManagement/assignmentPolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/assignmentPolicies/accessPackageAssignmentPolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/assignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/assignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings/customExtensionStageSettingId",
			Expected: &IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId{
				AccessPackageAssignmentPolicyId: "accessPackageAssignmentPolicyId",
				CustomExtensionStageSettingId:   "customExtensionStageSettingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/assignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings/customExtensionStageSettingId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageAssignmentPolicyId != v.Expected.AccessPackageAssignmentPolicyId {
			t.Fatalf("Expected %q but got %q for AccessPackageAssignmentPolicyId", v.Expected.AccessPackageAssignmentPolicyId, actual.AccessPackageAssignmentPolicyId)
		}

		if actual.CustomExtensionStageSettingId != v.Expected.CustomExtensionStageSettingId {
			t.Fatalf("Expected %q but got %q for CustomExtensionStageSettingId", v.Expected.CustomExtensionStageSettingId, actual.CustomExtensionStageSettingId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId
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
			Input: "/identityGovernance/entitlementManagement/assignmentPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aSsIgNmEnTpOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/assignmentPolicies/accessPackageAssignmentPolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aSsIgNmEnTpOlIcIeS/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/assignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aSsIgNmEnTpOlIcIeS/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId/cUsToMeXtEnSiOnStAgEsEtTiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/assignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings/customExtensionStageSettingId",
			Expected: &IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId{
				AccessPackageAssignmentPolicyId: "accessPackageAssignmentPolicyId",
				CustomExtensionStageSettingId:   "customExtensionStageSettingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/assignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings/customExtensionStageSettingId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aSsIgNmEnTpOlIcIeS/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId/cUsToMeXtEnSiOnStAgEsEtTiNgS/cUsToMeXtEnSiOnStAgEsEtTiNgId",
			Expected: &IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId{
				AccessPackageAssignmentPolicyId: "aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId",
				CustomExtensionStageSettingId:   "cUsToMeXtEnSiOnStAgEsEtTiNgId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aSsIgNmEnTpOlIcIeS/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId/cUsToMeXtEnSiOnStAgEsEtTiNgS/cUsToMeXtEnSiOnStAgEsEtTiNgId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageAssignmentPolicyId != v.Expected.AccessPackageAssignmentPolicyId {
			t.Fatalf("Expected %q but got %q for AccessPackageAssignmentPolicyId", v.Expected.AccessPackageAssignmentPolicyId, actual.AccessPackageAssignmentPolicyId)
		}

		if actual.CustomExtensionStageSettingId != v.Expected.CustomExtensionStageSettingId {
			t.Fatalf("Expected %q but got %q for CustomExtensionStageSettingId", v.Expected.CustomExtensionStageSettingId, actual.CustomExtensionStageSettingId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
