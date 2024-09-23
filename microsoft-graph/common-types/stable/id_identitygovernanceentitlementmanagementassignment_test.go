package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAssignmentId{}

func TestNewIdentityGovernanceEntitlementManagementAssignmentID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAssignmentID("accessPackageAssignmentId")

	if id.AccessPackageAssignmentId != "accessPackageAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageAssignmentId'", id.AccessPackageAssignmentId, "accessPackageAssignmentId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAssignmentID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAssignmentID("accessPackageAssignmentId").ID()
	expected := "/identityGovernance/entitlementManagement/assignments/accessPackageAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAssignmentId
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
			Input: "/identityGovernance/entitlementManagement/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/assignments/accessPackageAssignmentId",
			Expected: &IdentityGovernanceEntitlementManagementAssignmentId{
				AccessPackageAssignmentId: "accessPackageAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/assignments/accessPackageAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAssignmentID(v.Input)
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

	}
}

func TestParseIdentityGovernanceEntitlementManagementAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAssignmentId
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
			Input: "/identityGovernance/entitlementManagement/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/assignments/accessPackageAssignmentId",
			Expected: &IdentityGovernanceEntitlementManagementAssignmentId{
				AccessPackageAssignmentId: "accessPackageAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/assignments/accessPackageAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aSsIgNmEnTs/aCcEsSpAcKaGeAsSiGnMeNtId",
			Expected: &IdentityGovernanceEntitlementManagementAssignmentId{
				AccessPackageAssignmentId: "aCcEsSpAcKaGeAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aSsIgNmEnTs/aCcEsSpAcKaGeAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAssignmentIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAssignmentId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
