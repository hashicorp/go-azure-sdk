package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId{}

func TestNewIdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorID("connectedOrganizationIdValue", "directoryObjectIdValue")

	if id.ConnectedOrganizationId != "connectedOrganizationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ConnectedOrganizationId'", id.ConnectedOrganizationId, "connectedOrganizationIdValue")
	}

	if id.DirectoryObjectId != "directoryObjectIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectIdValue")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorID("connectedOrganizationIdValue", "directoryObjectIdValue").ID()
	expected := "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationIdValue/externalSponsors/directoryObjectIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId
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
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationIdValue/externalSponsors",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationIdValue/externalSponsors/directoryObjectIdValue",
			Expected: &IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId{
				ConnectedOrganizationId: "connectedOrganizationIdValue",
				DirectoryObjectId:       "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationIdValue/externalSponsors/directoryObjectIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorID(v.Input)
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

func TestParseIdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId
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
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationIdValue/externalSponsors",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cOnNeCtEdOrGaNiZaTiOnS/cOnNeCtEdOrGaNiZaTiOnIdVaLuE/eXtErNaLsPoNsOrS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationIdValue/externalSponsors/directoryObjectIdValue",
			Expected: &IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId{
				ConnectedOrganizationId: "connectedOrganizationIdValue",
				DirectoryObjectId:       "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationIdValue/externalSponsors/directoryObjectIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cOnNeCtEdOrGaNiZaTiOnS/cOnNeCtEdOrGaNiZaTiOnIdVaLuE/eXtErNaLsPoNsOrS/dIrEcToRyObJeCtIdVaLuE",
			Expected: &IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId{
				ConnectedOrganizationId: "cOnNeCtEdOrGaNiZaTiOnIdVaLuE",
				DirectoryObjectId:       "dIrEcToRyObJeCtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cOnNeCtEdOrGaNiZaTiOnS/cOnNeCtEdOrGaNiZaTiOnIdVaLuE/eXtErNaLsPoNsOrS/dIrEcToRyObJeCtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorIDInsensitively(v.Input)
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

func TestSegmentsForIdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
