package entitlementmanagementaccesspackageassignmentpolicyquestion

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionID("accessPackageIdValue", "accessPackageAssignmentPolicyIdValue", "accessPackageQuestionIdValue")

	if id.AccessPackageId != "accessPackageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageId'", id.AccessPackageId, "accessPackageIdValue")
	}

	if id.AccessPackageAssignmentPolicyId != "accessPackageAssignmentPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageAssignmentPolicyId'", id.AccessPackageAssignmentPolicyId, "accessPackageAssignmentPolicyIdValue")
	}

	if id.AccessPackageQuestionId != "accessPackageQuestionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageQuestionId'", id.AccessPackageQuestionId, "accessPackageQuestionIdValue")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionID("accessPackageIdValue", "accessPackageAssignmentPolicyIdValue", "accessPackageQuestionIdValue").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue/assignmentPolicies/accessPackageAssignmentPolicyIdValue/questions/accessPackageQuestionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId
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
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue/assignmentPolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue/assignmentPolicies/accessPackageAssignmentPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue/assignmentPolicies/accessPackageAssignmentPolicyIdValue/questions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue/assignmentPolicies/accessPackageAssignmentPolicyIdValue/questions/accessPackageQuestionIdValue",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId{
				AccessPackageId:                 "accessPackageIdValue",
				AccessPackageAssignmentPolicyId: "accessPackageAssignmentPolicyIdValue",
				AccessPackageQuestionId:         "accessPackageQuestionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue/assignmentPolicies/accessPackageAssignmentPolicyIdValue/questions/accessPackageQuestionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionID(v.Input)
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

		if actual.AccessPackageQuestionId != v.Expected.AccessPackageQuestionId {
			t.Fatalf("Expected %q but got %q for AccessPackageQuestionId", v.Expected.AccessPackageQuestionId, actual.AccessPackageQuestionId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId
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
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue/assignmentPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeIdVaLuE/aSsIgNmEnTpOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue/assignmentPolicies/accessPackageAssignmentPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeIdVaLuE/aSsIgNmEnTpOlIcIeS/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue/assignmentPolicies/accessPackageAssignmentPolicyIdValue/questions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeIdVaLuE/aSsIgNmEnTpOlIcIeS/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyIdVaLuE/qUeStIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue/assignmentPolicies/accessPackageAssignmentPolicyIdValue/questions/accessPackageQuestionIdValue",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId{
				AccessPackageId:                 "accessPackageIdValue",
				AccessPackageAssignmentPolicyId: "accessPackageAssignmentPolicyIdValue",
				AccessPackageQuestionId:         "accessPackageQuestionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue/assignmentPolicies/accessPackageAssignmentPolicyIdValue/questions/accessPackageQuestionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeIdVaLuE/aSsIgNmEnTpOlIcIeS/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyIdVaLuE/qUeStIoNs/aCcEsSpAcKaGeQuEsTiOnIdVaLuE",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId{
				AccessPackageId:                 "aCcEsSpAcKaGeIdVaLuE",
				AccessPackageAssignmentPolicyId: "aCcEsSpAcKaGeAsSiGnMeNtPoLiCyIdVaLuE",
				AccessPackageQuestionId:         "aCcEsSpAcKaGeQuEsTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeIdVaLuE/aSsIgNmEnTpOlIcIeS/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyIdVaLuE/qUeStIoNs/aCcEsSpAcKaGeQuEsTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionIDInsensitively(v.Input)
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

		if actual.AccessPackageQuestionId != v.Expected.AccessPackageQuestionId {
			t.Fatalf("Expected %q but got %q for AccessPackageQuestionId", v.Expected.AccessPackageQuestionId, actual.AccessPackageQuestionId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
