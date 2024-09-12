package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId{}

func TestNewIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorID("connectedOrganizationIdValue", "directoryObjectIdValue")

	if id.ConnectedOrganizationId != "connectedOrganizationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ConnectedOrganizationId'", id.ConnectedOrganizationId, "connectedOrganizationIdValue")
	}

	if id.DirectoryObjectId != "directoryObjectIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectIdValue")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorID("connectedOrganizationIdValue", "directoryObjectIdValue").ID()
	expected := "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationIdValue/internalSponsors/directoryObjectIdValue"
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
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationIdValue/internalSponsors",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationIdValue/internalSponsors/directoryObjectIdValue",
			Expected: &IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId{
				ConnectedOrganizationId: "connectedOrganizationIdValue",
				DirectoryObjectId:       "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationIdValue/internalSponsors/directoryObjectIdValue/extra",
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
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cOnNeCtEdOrGaNiZaTiOnS/cOnNeCtEdOrGaNiZaTiOnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationIdValue/internalSponsors",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cOnNeCtEdOrGaNiZaTiOnS/cOnNeCtEdOrGaNiZaTiOnIdVaLuE/iNtErNaLsPoNsOrS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationIdValue/internalSponsors/directoryObjectIdValue",
			Expected: &IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId{
				ConnectedOrganizationId: "connectedOrganizationIdValue",
				DirectoryObjectId:       "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationIdValue/internalSponsors/directoryObjectIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cOnNeCtEdOrGaNiZaTiOnS/cOnNeCtEdOrGaNiZaTiOnIdVaLuE/iNtErNaLsPoNsOrS/dIrEcToRyObJeCtIdVaLuE",
			Expected: &IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId{
				ConnectedOrganizationId: "cOnNeCtEdOrGaNiZaTiOnIdVaLuE",
				DirectoryObjectId:       "dIrEcToRyObJeCtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cOnNeCtEdOrGaNiZaTiOnS/cOnNeCtEdOrGaNiZaTiOnIdVaLuE/iNtErNaLsPoNsOrS/dIrEcToRyObJeCtIdVaLuE/extra",
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
