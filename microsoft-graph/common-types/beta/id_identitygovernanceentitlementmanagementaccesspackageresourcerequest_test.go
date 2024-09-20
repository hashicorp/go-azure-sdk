package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageResourceRequestID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageResourceRequestID("accessPackageResourceRequestId")

	if id.AccessPackageResourceRequestId != "accessPackageResourceRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRequestId'", id.AccessPackageResourceRequestId, "accessPackageResourceRequestId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageResourceRequestID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageResourceRequestID("accessPackageResourceRequestId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageResourceRequests/accessPackageResourceRequestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageResourceRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceRequests/accessPackageResourceRequestId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId{
				AccessPackageResourceRequestId: "accessPackageResourceRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceRequests/accessPackageResourceRequestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageResourceRequestID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageResourceRequestId != v.Expected.AccessPackageResourceRequestId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRequestId", v.Expected.AccessPackageResourceRequestId, actual.AccessPackageResourceRequestId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageResourceRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceRequests/accessPackageResourceRequestId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId{
				AccessPackageResourceRequestId: "accessPackageResourceRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceRequests/accessPackageResourceRequestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeReQuEsTs/aCcEsSpAcKaGeReSoUrCeReQuEsTiD",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId{
				AccessPackageResourceRequestId: "aCcEsSpAcKaGeReSoUrCeReQuEsTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeReQuEsTs/aCcEsSpAcKaGeReSoUrCeReQuEsTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageResourceRequestIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageResourceRequestId != v.Expected.AccessPackageResourceRequestId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRequestId", v.Expected.AccessPackageResourceRequestId, actual.AccessPackageResourceRequestId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageResourceRequestId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
