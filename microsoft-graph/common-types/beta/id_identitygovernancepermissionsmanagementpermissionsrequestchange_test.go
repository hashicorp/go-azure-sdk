package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePermissionsManagementPermissionsRequestChangeId{}

func TestNewIdentityGovernancePermissionsManagementPermissionsRequestChangeID(t *testing.T) {
	id := NewIdentityGovernancePermissionsManagementPermissionsRequestChangeID("permissionsRequestChangeId")

	if id.PermissionsRequestChangeId != "permissionsRequestChangeId" {
		t.Fatalf("Expected %q but got %q for Segment 'PermissionsRequestChangeId'", id.PermissionsRequestChangeId, "permissionsRequestChangeId")
	}
}

func TestFormatIdentityGovernancePermissionsManagementPermissionsRequestChangeID(t *testing.T) {
	actual := NewIdentityGovernancePermissionsManagementPermissionsRequestChangeID("permissionsRequestChangeId").ID()
	expected := "/identityGovernance/permissionsManagement/permissionsRequestChanges/permissionsRequestChangeId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernancePermissionsManagementPermissionsRequestChangeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePermissionsManagementPermissionsRequestChangeId
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
			Input: "/identityGovernance/permissionsManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/permissionsManagement/permissionsRequestChanges",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/permissionsManagement/permissionsRequestChanges/permissionsRequestChangeId",
			Expected: &IdentityGovernancePermissionsManagementPermissionsRequestChangeId{
				PermissionsRequestChangeId: "permissionsRequestChangeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/permissionsManagement/permissionsRequestChanges/permissionsRequestChangeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePermissionsManagementPermissionsRequestChangeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PermissionsRequestChangeId != v.Expected.PermissionsRequestChangeId {
			t.Fatalf("Expected %q but got %q for PermissionsRequestChangeId", v.Expected.PermissionsRequestChangeId, actual.PermissionsRequestChangeId)
		}

	}
}

func TestParseIdentityGovernancePermissionsManagementPermissionsRequestChangeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePermissionsManagementPermissionsRequestChangeId
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
			Input: "/identityGovernance/permissionsManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsMaNaGeMeNt",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/permissionsManagement/permissionsRequestChanges",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsMaNaGeMeNt/pErMiSsIoNsReQuEsTcHaNgEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/permissionsManagement/permissionsRequestChanges/permissionsRequestChangeId",
			Expected: &IdentityGovernancePermissionsManagementPermissionsRequestChangeId{
				PermissionsRequestChangeId: "permissionsRequestChangeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/permissionsManagement/permissionsRequestChanges/permissionsRequestChangeId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsMaNaGeMeNt/pErMiSsIoNsReQuEsTcHaNgEs/pErMiSsIoNsReQuEsTcHaNgEiD",
			Expected: &IdentityGovernancePermissionsManagementPermissionsRequestChangeId{
				PermissionsRequestChangeId: "pErMiSsIoNsReQuEsTcHaNgEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsMaNaGeMeNt/pErMiSsIoNsReQuEsTcHaNgEs/pErMiSsIoNsReQuEsTcHaNgEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePermissionsManagementPermissionsRequestChangeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PermissionsRequestChangeId != v.Expected.PermissionsRequestChangeId {
			t.Fatalf("Expected %q but got %q for PermissionsRequestChangeId", v.Expected.PermissionsRequestChangeId, actual.PermissionsRequestChangeId)
		}

	}
}

func TestSegmentsForIdentityGovernancePermissionsManagementPermissionsRequestChangeId(t *testing.T) {
	segments := IdentityGovernancePermissionsManagementPermissionsRequestChangeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernancePermissionsManagementPermissionsRequestChangeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
