package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceRoleManagementAlertAlertId{}

func TestNewIdentityGovernanceRoleManagementAlertAlertID(t *testing.T) {
	id := NewIdentityGovernanceRoleManagementAlertAlertID("unifiedRoleManagementAlertIdValue")

	if id.UnifiedRoleManagementAlertId != "unifiedRoleManagementAlertIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleManagementAlertId'", id.UnifiedRoleManagementAlertId, "unifiedRoleManagementAlertIdValue")
	}
}

func TestFormatIdentityGovernanceRoleManagementAlertAlertID(t *testing.T) {
	actual := NewIdentityGovernanceRoleManagementAlertAlertID("unifiedRoleManagementAlertIdValue").ID()
	expected := "/identityGovernance/roleManagementAlerts/alerts/unifiedRoleManagementAlertIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceRoleManagementAlertAlertID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceRoleManagementAlertAlertId
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
			Input: "/identityGovernance/roleManagementAlerts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/roleManagementAlerts/alerts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/roleManagementAlerts/alerts/unifiedRoleManagementAlertIdValue",
			Expected: &IdentityGovernanceRoleManagementAlertAlertId{
				UnifiedRoleManagementAlertId: "unifiedRoleManagementAlertIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/roleManagementAlerts/alerts/unifiedRoleManagementAlertIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceRoleManagementAlertAlertID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleManagementAlertId != v.Expected.UnifiedRoleManagementAlertId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleManagementAlertId", v.Expected.UnifiedRoleManagementAlertId, actual.UnifiedRoleManagementAlertId)
		}

	}
}

func TestParseIdentityGovernanceRoleManagementAlertAlertIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceRoleManagementAlertAlertId
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
			Input: "/identityGovernance/roleManagementAlerts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/rOlEmAnAgEmEnTaLeRtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/roleManagementAlerts/alerts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/rOlEmAnAgEmEnTaLeRtS/aLeRtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/roleManagementAlerts/alerts/unifiedRoleManagementAlertIdValue",
			Expected: &IdentityGovernanceRoleManagementAlertAlertId{
				UnifiedRoleManagementAlertId: "unifiedRoleManagementAlertIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/roleManagementAlerts/alerts/unifiedRoleManagementAlertIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/rOlEmAnAgEmEnTaLeRtS/aLeRtS/uNiFiEdRoLeMaNaGeMeNtAlErTiDvAlUe",
			Expected: &IdentityGovernanceRoleManagementAlertAlertId{
				UnifiedRoleManagementAlertId: "uNiFiEdRoLeMaNaGeMeNtAlErTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/rOlEmAnAgEmEnTaLeRtS/aLeRtS/uNiFiEdRoLeMaNaGeMeNtAlErTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceRoleManagementAlertAlertIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleManagementAlertId != v.Expected.UnifiedRoleManagementAlertId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleManagementAlertId", v.Expected.UnifiedRoleManagementAlertId, actual.UnifiedRoleManagementAlertId)
		}

	}
}

func TestSegmentsForIdentityGovernanceRoleManagementAlertAlertId(t *testing.T) {
	segments := IdentityGovernanceRoleManagementAlertAlertId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceRoleManagementAlertAlertId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
