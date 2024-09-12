package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceRoleManagementAlertAlertDefinitionId{}

func TestNewIdentityGovernanceRoleManagementAlertAlertDefinitionID(t *testing.T) {
	id := NewIdentityGovernanceRoleManagementAlertAlertDefinitionID("unifiedRoleManagementAlertDefinitionIdValue")

	if id.UnifiedRoleManagementAlertDefinitionId != "unifiedRoleManagementAlertDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleManagementAlertDefinitionId'", id.UnifiedRoleManagementAlertDefinitionId, "unifiedRoleManagementAlertDefinitionIdValue")
	}
}

func TestFormatIdentityGovernanceRoleManagementAlertAlertDefinitionID(t *testing.T) {
	actual := NewIdentityGovernanceRoleManagementAlertAlertDefinitionID("unifiedRoleManagementAlertDefinitionIdValue").ID()
	expected := "/identityGovernance/roleManagementAlerts/alertDefinitions/unifiedRoleManagementAlertDefinitionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceRoleManagementAlertAlertDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceRoleManagementAlertAlertDefinitionId
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
			Input: "/identityGovernance/roleManagementAlerts/alertDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/roleManagementAlerts/alertDefinitions/unifiedRoleManagementAlertDefinitionIdValue",
			Expected: &IdentityGovernanceRoleManagementAlertAlertDefinitionId{
				UnifiedRoleManagementAlertDefinitionId: "unifiedRoleManagementAlertDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/roleManagementAlerts/alertDefinitions/unifiedRoleManagementAlertDefinitionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceRoleManagementAlertAlertDefinitionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleManagementAlertDefinitionId != v.Expected.UnifiedRoleManagementAlertDefinitionId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleManagementAlertDefinitionId", v.Expected.UnifiedRoleManagementAlertDefinitionId, actual.UnifiedRoleManagementAlertDefinitionId)
		}

	}
}

func TestParseIdentityGovernanceRoleManagementAlertAlertDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceRoleManagementAlertAlertDefinitionId
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
			Input: "/identityGovernance/roleManagementAlerts/alertDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/rOlEmAnAgEmEnTaLeRtS/aLeRtDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/roleManagementAlerts/alertDefinitions/unifiedRoleManagementAlertDefinitionIdValue",
			Expected: &IdentityGovernanceRoleManagementAlertAlertDefinitionId{
				UnifiedRoleManagementAlertDefinitionId: "unifiedRoleManagementAlertDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/roleManagementAlerts/alertDefinitions/unifiedRoleManagementAlertDefinitionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/rOlEmAnAgEmEnTaLeRtS/aLeRtDeFiNiTiOnS/uNiFiEdRoLeMaNaGeMeNtAlErTdEfInItIoNiDvAlUe",
			Expected: &IdentityGovernanceRoleManagementAlertAlertDefinitionId{
				UnifiedRoleManagementAlertDefinitionId: "uNiFiEdRoLeMaNaGeMeNtAlErTdEfInItIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/rOlEmAnAgEmEnTaLeRtS/aLeRtDeFiNiTiOnS/uNiFiEdRoLeMaNaGeMeNtAlErTdEfInItIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceRoleManagementAlertAlertDefinitionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleManagementAlertDefinitionId != v.Expected.UnifiedRoleManagementAlertDefinitionId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleManagementAlertDefinitionId", v.Expected.UnifiedRoleManagementAlertDefinitionId, actual.UnifiedRoleManagementAlertDefinitionId)
		}

	}
}

func TestSegmentsForIdentityGovernanceRoleManagementAlertAlertDefinitionId(t *testing.T) {
	segments := IdentityGovernanceRoleManagementAlertAlertDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceRoleManagementAlertAlertDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
