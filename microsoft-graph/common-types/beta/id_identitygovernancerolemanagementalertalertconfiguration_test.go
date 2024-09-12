package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceRoleManagementAlertAlertConfigurationId{}

func TestNewIdentityGovernanceRoleManagementAlertAlertConfigurationID(t *testing.T) {
	id := NewIdentityGovernanceRoleManagementAlertAlertConfigurationID("unifiedRoleManagementAlertConfigurationIdValue")

	if id.UnifiedRoleManagementAlertConfigurationId != "unifiedRoleManagementAlertConfigurationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleManagementAlertConfigurationId'", id.UnifiedRoleManagementAlertConfigurationId, "unifiedRoleManagementAlertConfigurationIdValue")
	}
}

func TestFormatIdentityGovernanceRoleManagementAlertAlertConfigurationID(t *testing.T) {
	actual := NewIdentityGovernanceRoleManagementAlertAlertConfigurationID("unifiedRoleManagementAlertConfigurationIdValue").ID()
	expected := "/identityGovernance/roleManagementAlerts/alertConfigurations/unifiedRoleManagementAlertConfigurationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceRoleManagementAlertAlertConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceRoleManagementAlertAlertConfigurationId
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
			Input: "/identityGovernance/roleManagementAlerts/alertConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/roleManagementAlerts/alertConfigurations/unifiedRoleManagementAlertConfigurationIdValue",
			Expected: &IdentityGovernanceRoleManagementAlertAlertConfigurationId{
				UnifiedRoleManagementAlertConfigurationId: "unifiedRoleManagementAlertConfigurationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/roleManagementAlerts/alertConfigurations/unifiedRoleManagementAlertConfigurationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceRoleManagementAlertAlertConfigurationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleManagementAlertConfigurationId != v.Expected.UnifiedRoleManagementAlertConfigurationId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleManagementAlertConfigurationId", v.Expected.UnifiedRoleManagementAlertConfigurationId, actual.UnifiedRoleManagementAlertConfigurationId)
		}

	}
}

func TestParseIdentityGovernanceRoleManagementAlertAlertConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceRoleManagementAlertAlertConfigurationId
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
			Input: "/identityGovernance/roleManagementAlerts/alertConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/rOlEmAnAgEmEnTaLeRtS/aLeRtCoNfIgUrAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/roleManagementAlerts/alertConfigurations/unifiedRoleManagementAlertConfigurationIdValue",
			Expected: &IdentityGovernanceRoleManagementAlertAlertConfigurationId{
				UnifiedRoleManagementAlertConfigurationId: "unifiedRoleManagementAlertConfigurationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/roleManagementAlerts/alertConfigurations/unifiedRoleManagementAlertConfigurationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/rOlEmAnAgEmEnTaLeRtS/aLeRtCoNfIgUrAtIoNs/uNiFiEdRoLeMaNaGeMeNtAlErTcOnFiGuRaTiOnIdVaLuE",
			Expected: &IdentityGovernanceRoleManagementAlertAlertConfigurationId{
				UnifiedRoleManagementAlertConfigurationId: "uNiFiEdRoLeMaNaGeMeNtAlErTcOnFiGuRaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/rOlEmAnAgEmEnTaLeRtS/aLeRtCoNfIgUrAtIoNs/uNiFiEdRoLeMaNaGeMeNtAlErTcOnFiGuRaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceRoleManagementAlertAlertConfigurationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleManagementAlertConfigurationId != v.Expected.UnifiedRoleManagementAlertConfigurationId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleManagementAlertConfigurationId", v.Expected.UnifiedRoleManagementAlertConfigurationId, actual.UnifiedRoleManagementAlertConfigurationId)
		}

	}
}

func TestSegmentsForIdentityGovernanceRoleManagementAlertAlertConfigurationId(t *testing.T) {
	segments := IdentityGovernanceRoleManagementAlertAlertConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceRoleManagementAlertAlertConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
