package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyIdCustomExtensionStageSettingID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyIdCustomExtensionStageSettingID("accessPackageAssignmentId", "accessPackageAssignmentPolicyId", "customExtensionStageSettingId")

	if id.AccessPackageAssignmentId != "accessPackageAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageAssignmentId'", id.AccessPackageAssignmentId, "accessPackageAssignmentId")
	}

	if id.AccessPackageAssignmentPolicyId != "accessPackageAssignmentPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageAssignmentPolicyId'", id.AccessPackageAssignmentPolicyId, "accessPackageAssignmentPolicyId")
	}

	if id.CustomExtensionStageSettingId != "customExtensionStageSettingId" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomExtensionStageSettingId'", id.CustomExtensionStageSettingId, "customExtensionStageSettingId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyIdCustomExtensionStageSettingID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyIdCustomExtensionStageSettingID("accessPackageAssignmentId", "accessPackageAssignmentPolicyId", "customExtensionStageSettingId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings/customExtensionStageSettingId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyIdCustomExtensionStageSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageAssignmentPolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings/customExtensionStageSettingId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId{
				AccessPackageAssignmentId:       "accessPackageAssignmentId",
				AccessPackageAssignmentPolicyId: "accessPackageAssignmentPolicyId",
				CustomExtensionStageSettingId:   "customExtensionStageSettingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings/customExtensionStageSettingId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyIdCustomExtensionStageSettingID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageAssignmentId != v.Expected.AccessPackageAssignmentId {
			t.Fatalf("Expected %q but got %q for AccessPackageAssignmentId", v.Expected.AccessPackageAssignmentId, actual.AccessPackageAssignmentId)
		}

		if actual.AccessPackageAssignmentPolicyId != v.Expected.AccessPackageAssignmentPolicyId {
			t.Fatalf("Expected %q but got %q for AccessPackageAssignmentPolicyId", v.Expected.AccessPackageAssignmentPolicyId, actual.AccessPackageAssignmentPolicyId)
		}

		if actual.CustomExtensionStageSettingId != v.Expected.CustomExtensionStageSettingId {
			t.Fatalf("Expected %q but got %q for CustomExtensionStageSettingId", v.Expected.CustomExtensionStageSettingId, actual.CustomExtensionStageSettingId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyIdCustomExtensionStageSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId/aCcEsSpAcKaGe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageAssignmentPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId/aCcEsSpAcKaGe/aCcEsSpAcKaGeAsSiGnMeNtPoLiCiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId/aCcEsSpAcKaGe/aCcEsSpAcKaGeAsSiGnMeNtPoLiCiEs/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId/aCcEsSpAcKaGe/aCcEsSpAcKaGeAsSiGnMeNtPoLiCiEs/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId/cUsToMeXtEnSiOnStAgEsEtTiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings/customExtensionStageSettingId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId{
				AccessPackageAssignmentId:       "accessPackageAssignmentId",
				AccessPackageAssignmentPolicyId: "accessPackageAssignmentPolicyId",
				CustomExtensionStageSettingId:   "customExtensionStageSettingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId/customExtensionStageSettings/customExtensionStageSettingId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId/aCcEsSpAcKaGe/aCcEsSpAcKaGeAsSiGnMeNtPoLiCiEs/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId/cUsToMeXtEnSiOnStAgEsEtTiNgS/cUsToMeXtEnSiOnStAgEsEtTiNgId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId{
				AccessPackageAssignmentId:       "aCcEsSpAcKaGeAsSiGnMeNtId",
				AccessPackageAssignmentPolicyId: "aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId",
				CustomExtensionStageSettingId:   "cUsToMeXtEnSiOnStAgEsEtTiNgId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId/aCcEsSpAcKaGe/aCcEsSpAcKaGeAsSiGnMeNtPoLiCiEs/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId/cUsToMeXtEnSiOnStAgEsEtTiNgS/cUsToMeXtEnSiOnStAgEsEtTiNgId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyIdCustomExtensionStageSettingIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageAssignmentId != v.Expected.AccessPackageAssignmentId {
			t.Fatalf("Expected %q but got %q for AccessPackageAssignmentId", v.Expected.AccessPackageAssignmentId, actual.AccessPackageAssignmentId)
		}

		if actual.AccessPackageAssignmentPolicyId != v.Expected.AccessPackageAssignmentPolicyId {
			t.Fatalf("Expected %q but got %q for AccessPackageAssignmentPolicyId", v.Expected.AccessPackageAssignmentPolicyId, actual.AccessPackageAssignmentPolicyId)
		}

		if actual.CustomExtensionStageSettingId != v.Expected.CustomExtensionStageSettingId {
			t.Fatalf("Expected %q but got %q for CustomExtensionStageSettingId", v.Expected.CustomExtensionStageSettingId, actual.CustomExtensionStageSettingId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
