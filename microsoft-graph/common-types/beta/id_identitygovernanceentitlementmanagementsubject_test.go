package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementSubjectId{}

func TestNewIdentityGovernanceEntitlementManagementSubjectID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementSubjectID("accessPackageSubjectIdValue")

	if id.AccessPackageSubjectId != "accessPackageSubjectIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageSubjectId'", id.AccessPackageSubjectId, "accessPackageSubjectIdValue")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementSubjectID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementSubjectID("accessPackageSubjectIdValue").ID()
	expected := "/identityGovernance/entitlementManagement/subjects/accessPackageSubjectIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementSubjectID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementSubjectId
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
			Input: "/identityGovernance/entitlementManagement/subjects",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/subjects/accessPackageSubjectIdValue",
			Expected: &IdentityGovernanceEntitlementManagementSubjectId{
				AccessPackageSubjectId: "accessPackageSubjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/subjects/accessPackageSubjectIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementSubjectID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageSubjectId != v.Expected.AccessPackageSubjectId {
			t.Fatalf("Expected %q but got %q for AccessPackageSubjectId", v.Expected.AccessPackageSubjectId, actual.AccessPackageSubjectId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementSubjectIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementSubjectId
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
			Input: "/identityGovernance/entitlementManagement/subjects",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/sUbJeCtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/subjects/accessPackageSubjectIdValue",
			Expected: &IdentityGovernanceEntitlementManagementSubjectId{
				AccessPackageSubjectId: "accessPackageSubjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/subjects/accessPackageSubjectIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/sUbJeCtS/aCcEsSpAcKaGeSuBjEcTiDvAlUe",
			Expected: &IdentityGovernanceEntitlementManagementSubjectId{
				AccessPackageSubjectId: "aCcEsSpAcKaGeSuBjEcTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/sUbJeCtS/aCcEsSpAcKaGeSuBjEcTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementSubjectIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageSubjectId != v.Expected.AccessPackageSubjectId {
			t.Fatalf("Expected %q but got %q for AccessPackageSubjectId", v.Expected.AccessPackageSubjectId, actual.AccessPackageSubjectId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementSubjectId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementSubjectId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementSubjectId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
