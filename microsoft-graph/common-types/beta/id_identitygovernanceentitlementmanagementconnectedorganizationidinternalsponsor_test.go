package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId{}

func TestNewIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorID("connectedOrganizationId", "directoryObjectId")

	if id.ConnectedOrganizationId != "connectedOrganizationId" {
		t.Fatalf("Expected %q but got %q for Segment 'ConnectedOrganizationId'", id.ConnectedOrganizationId, "connectedOrganizationId")
	}

	if id.DirectoryObjectId != "directoryObjectId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorID("connectedOrganizationId", "directoryObjectId").ID()
	expected := "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationId/internalSponsors/directoryObjectId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId
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
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationId/internalSponsors",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationId/internalSponsors/directoryObjectId",
			Expected: &IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId{
				ConnectedOrganizationId: "connectedOrganizationId",
				DirectoryObjectId:       "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationId/internalSponsors/directoryObjectId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ConnectedOrganizationId != v.Expected.ConnectedOrganizationId {
			t.Fatalf("Expected %q but got %q for ConnectedOrganizationId", v.Expected.ConnectedOrganizationId, actual.ConnectedOrganizationId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId
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
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cOnNeCtEdOrGaNiZaTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cOnNeCtEdOrGaNiZaTiOnS/cOnNeCtEdOrGaNiZaTiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationId/internalSponsors",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cOnNeCtEdOrGaNiZaTiOnS/cOnNeCtEdOrGaNiZaTiOnId/iNtErNaLsPoNsOrS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationId/internalSponsors/directoryObjectId",
			Expected: &IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId{
				ConnectedOrganizationId: "connectedOrganizationId",
				DirectoryObjectId:       "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationId/internalSponsors/directoryObjectId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cOnNeCtEdOrGaNiZaTiOnS/cOnNeCtEdOrGaNiZaTiOnId/iNtErNaLsPoNsOrS/dIrEcToRyObJeCtId",
			Expected: &IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId{
				ConnectedOrganizationId: "cOnNeCtEdOrGaNiZaTiOnId",
				DirectoryObjectId:       "dIrEcToRyObJeCtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cOnNeCtEdOrGaNiZaTiOnS/cOnNeCtEdOrGaNiZaTiOnId/iNtErNaLsPoNsOrS/dIrEcToRyObJeCtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ConnectedOrganizationId != v.Expected.ConnectedOrganizationId {
			t.Fatalf("Expected %q but got %q for ConnectedOrganizationId", v.Expected.ConnectedOrganizationId, actual.ConnectedOrganizationId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
