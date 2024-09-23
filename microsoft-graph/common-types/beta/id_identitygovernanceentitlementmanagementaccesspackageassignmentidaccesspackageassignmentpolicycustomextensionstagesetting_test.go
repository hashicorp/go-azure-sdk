package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingID("accessPackageAssignmentId", "customExtensionStageSettingId")

	if id.AccessPackageAssignmentId != "accessPackageAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageAssignmentId'", id.AccessPackageAssignmentId, "accessPackageAssignmentId")
	}

	if id.CustomExtensionStageSettingId != "customExtensionStageSettingId" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomExtensionStageSettingId'", id.CustomExtensionStageSettingId, "customExtensionStageSettingId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingID("accessPackageAssignmentId", "customExtensionStageSettingId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackageAssignmentPolicy/customExtensionStageSettings/customExtensionStageSettingId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackageAssignmentPolicy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackageAssignmentPolicy/customExtensionStageSettings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackageAssignmentPolicy/customExtensionStageSettings/customExtensionStageSettingId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId{
				AccessPackageAssignmentId:     "accessPackageAssignmentId",
				CustomExtensionStageSettingId: "customExtensionStageSettingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackageAssignmentPolicy/customExtensionStageSettings/customExtensionStageSettingId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingID(v.Input)
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

		if actual.CustomExtensionStageSettingId != v.Expected.CustomExtensionStageSettingId {
			t.Fatalf("Expected %q but got %q for CustomExtensionStageSettingId", v.Expected.CustomExtensionStageSettingId, actual.CustomExtensionStageSettingId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackageAssignmentPolicy",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId/aCcEsSpAcKaGeAsSiGnMeNtPoLiCy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackageAssignmentPolicy/customExtensionStageSettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId/aCcEsSpAcKaGeAsSiGnMeNtPoLiCy/cUsToMeXtEnSiOnStAgEsEtTiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackageAssignmentPolicy/customExtensionStageSettings/customExtensionStageSettingId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId{
				AccessPackageAssignmentId:     "accessPackageAssignmentId",
				CustomExtensionStageSettingId: "customExtensionStageSettingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackageAssignmentPolicy/customExtensionStageSettings/customExtensionStageSettingId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId/aCcEsSpAcKaGeAsSiGnMeNtPoLiCy/cUsToMeXtEnSiOnStAgEsEtTiNgS/cUsToMeXtEnSiOnStAgEsEtTiNgId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId{
				AccessPackageAssignmentId:     "aCcEsSpAcKaGeAsSiGnMeNtId",
				CustomExtensionStageSettingId: "cUsToMeXtEnSiOnStAgEsEtTiNgId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId/aCcEsSpAcKaGeAsSiGnMeNtPoLiCy/cUsToMeXtEnSiOnStAgEsEtTiNgS/cUsToMeXtEnSiOnStAgEsEtTiNgId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingIDInsensitively(v.Input)
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

		if actual.CustomExtensionStageSettingId != v.Expected.CustomExtensionStageSettingId {
			t.Fatalf("Expected %q but got %q for CustomExtensionStageSettingId", v.Expected.CustomExtensionStageSettingId, actual.CustomExtensionStageSettingId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
