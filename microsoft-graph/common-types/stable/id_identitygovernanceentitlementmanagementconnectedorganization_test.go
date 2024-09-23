package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementConnectedOrganizationId{}

func TestNewIdentityGovernanceEntitlementManagementConnectedOrganizationID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementConnectedOrganizationID("connectedOrganizationId")

	if id.ConnectedOrganizationId != "connectedOrganizationId" {
		t.Fatalf("Expected %q but got %q for Segment 'ConnectedOrganizationId'", id.ConnectedOrganizationId, "connectedOrganizationId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementConnectedOrganizationID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementConnectedOrganizationID("connectedOrganizationId").ID()
	expected := "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementConnectedOrganizationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementConnectedOrganizationId
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
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationId",
			Expected: &IdentityGovernanceEntitlementManagementConnectedOrganizationId{
				ConnectedOrganizationId: "connectedOrganizationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementConnectedOrganizationID(v.Input)
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

	}
}

func TestParseIdentityGovernanceEntitlementManagementConnectedOrganizationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementConnectedOrganizationId
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
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationId",
			Expected: &IdentityGovernanceEntitlementManagementConnectedOrganizationId{
				ConnectedOrganizationId: "connectedOrganizationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/connectedOrganizations/connectedOrganizationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cOnNeCtEdOrGaNiZaTiOnS/cOnNeCtEdOrGaNiZaTiOnId",
			Expected: &IdentityGovernanceEntitlementManagementConnectedOrganizationId{
				ConnectedOrganizationId: "cOnNeCtEdOrGaNiZaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cOnNeCtEdOrGaNiZaTiOnS/cOnNeCtEdOrGaNiZaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementConnectedOrganizationIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementConnectedOrganizationId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementConnectedOrganizationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementConnectedOrganizationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
