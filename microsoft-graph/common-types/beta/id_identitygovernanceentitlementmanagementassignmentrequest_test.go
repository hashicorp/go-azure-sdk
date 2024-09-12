package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAssignmentRequestId{}

func TestNewIdentityGovernanceEntitlementManagementAssignmentRequestID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAssignmentRequestID("accessPackageAssignmentRequestIdValue")

	if id.AccessPackageAssignmentRequestId != "accessPackageAssignmentRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageAssignmentRequestId'", id.AccessPackageAssignmentRequestId, "accessPackageAssignmentRequestIdValue")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAssignmentRequestID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAssignmentRequestID("accessPackageAssignmentRequestIdValue").ID()
	expected := "/identityGovernance/entitlementManagement/assignmentRequests/accessPackageAssignmentRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAssignmentRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAssignmentRequestId
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
			Input: "/identityGovernance/entitlementManagement/assignmentRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/assignmentRequests/accessPackageAssignmentRequestIdValue",
			Expected: &IdentityGovernanceEntitlementManagementAssignmentRequestId{
				AccessPackageAssignmentRequestId: "accessPackageAssignmentRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/assignmentRequests/accessPackageAssignmentRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAssignmentRequestID(v.Input)
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

func TestParseIdentityGovernanceEntitlementManagementAssignmentRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAssignmentRequestId
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
			Input: "/identityGovernance/entitlementManagement/assignmentRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aSsIgNmEnTrEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/assignmentRequests/accessPackageAssignmentRequestIdValue",
			Expected: &IdentityGovernanceEntitlementManagementAssignmentRequestId{
				AccessPackageAssignmentRequestId: "accessPackageAssignmentRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/assignmentRequests/accessPackageAssignmentRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aSsIgNmEnTrEqUeStS/aCcEsSpAcKaGeAsSiGnMeNtReQuEsTiDvAlUe",
			Expected: &IdentityGovernanceEntitlementManagementAssignmentRequestId{
				AccessPackageAssignmentRequestId: "aCcEsSpAcKaGeAsSiGnMeNtReQuEsTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aSsIgNmEnTrEqUeStS/aCcEsSpAcKaGeAsSiGnMeNtReQuEsTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAssignmentRequestIDInsensitively(v.Input)
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

func TestSegmentsForIdentityGovernanceEntitlementManagementAssignmentRequestId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAssignmentRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAssignmentRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
