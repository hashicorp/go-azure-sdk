package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID("accessPackageAssignmentRequestId")

	if id.AccessPackageAssignmentRequestId != "accessPackageAssignmentRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageAssignmentRequestId'", id.AccessPackageAssignmentRequestId, "accessPackageAssignmentRequestId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID("accessPackageAssignmentRequestId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageAssignmentRequests/accessPackageAssignmentRequestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentRequests/accessPackageAssignmentRequestId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId{
				AccessPackageAssignmentRequestId: "accessPackageAssignmentRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentRequests/accessPackageAssignmentRequestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageAssignmentRequestId != v.Expected.AccessPackageAssignmentRequestId {
			t.Fatalf("Expected %q but got %q for AccessPackageAssignmentRequestId", v.Expected.AccessPackageAssignmentRequestId, actual.AccessPackageAssignmentRequestId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentRequests/accessPackageAssignmentRequestId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId{
				AccessPackageAssignmentRequestId: "accessPackageAssignmentRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentRequests/accessPackageAssignmentRequestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtReQuEsTs/aCcEsSpAcKaGeAsSiGnMeNtReQuEsTiD",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId{
				AccessPackageAssignmentRequestId: "aCcEsSpAcKaGeAsSiGnMeNtReQuEsTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtReQuEsTs/aCcEsSpAcKaGeAsSiGnMeNtReQuEsTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageAssignmentRequestId != v.Expected.AccessPackageAssignmentRequestId {
			t.Fatalf("Expected %q but got %q for AccessPackageAssignmentRequestId", v.Expected.AccessPackageAssignmentRequestId, actual.AccessPackageAssignmentRequestId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
