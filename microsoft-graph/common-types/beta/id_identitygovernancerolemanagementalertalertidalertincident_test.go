package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId{}

func TestNewIdentityGovernanceRoleManagementAlertAlertIdAlertIncidentID(t *testing.T) {
	id := NewIdentityGovernanceRoleManagementAlertAlertIdAlertIncidentID("unifiedRoleManagementAlertIdValue", "unifiedRoleManagementAlertIncidentIdValue")

	if id.UnifiedRoleManagementAlertId != "unifiedRoleManagementAlertIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleManagementAlertId'", id.UnifiedRoleManagementAlertId, "unifiedRoleManagementAlertIdValue")
	}

	if id.UnifiedRoleManagementAlertIncidentId != "unifiedRoleManagementAlertIncidentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleManagementAlertIncidentId'", id.UnifiedRoleManagementAlertIncidentId, "unifiedRoleManagementAlertIncidentIdValue")
	}
}

func TestFormatIdentityGovernanceRoleManagementAlertAlertIdAlertIncidentID(t *testing.T) {
	actual := NewIdentityGovernanceRoleManagementAlertAlertIdAlertIncidentID("unifiedRoleManagementAlertIdValue", "unifiedRoleManagementAlertIncidentIdValue").ID()
	expected := "/identityGovernance/roleManagementAlerts/alerts/unifiedRoleManagementAlertIdValue/alertIncidents/unifiedRoleManagementAlertIncidentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceRoleManagementAlertAlertIdAlertIncidentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId
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
			// Incomplete URI
			Input: "/identityGovernance/roleManagementAlerts/alerts/unifiedRoleManagementAlertIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/roleManagementAlerts/alerts/unifiedRoleManagementAlertIdValue/alertIncidents",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/roleManagementAlerts/alerts/unifiedRoleManagementAlertIdValue/alertIncidents/unifiedRoleManagementAlertIncidentIdValue",
			Expected: &IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId{
				UnifiedRoleManagementAlertId:         "unifiedRoleManagementAlertIdValue",
				UnifiedRoleManagementAlertIncidentId: "unifiedRoleManagementAlertIncidentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/roleManagementAlerts/alerts/unifiedRoleManagementAlertIdValue/alertIncidents/unifiedRoleManagementAlertIncidentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceRoleManagementAlertAlertIdAlertIncidentID(v.Input)
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

		if actual.UnifiedRoleManagementAlertIncidentId != v.Expected.UnifiedRoleManagementAlertIncidentId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleManagementAlertIncidentId", v.Expected.UnifiedRoleManagementAlertIncidentId, actual.UnifiedRoleManagementAlertIncidentId)
		}

	}
}

func TestParseIdentityGovernanceRoleManagementAlertAlertIdAlertIncidentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId
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
			// Incomplete URI
			Input: "/identityGovernance/roleManagementAlerts/alerts/unifiedRoleManagementAlertIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/rOlEmAnAgEmEnTaLeRtS/aLeRtS/uNiFiEdRoLeMaNaGeMeNtAlErTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/roleManagementAlerts/alerts/unifiedRoleManagementAlertIdValue/alertIncidents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/rOlEmAnAgEmEnTaLeRtS/aLeRtS/uNiFiEdRoLeMaNaGeMeNtAlErTiDvAlUe/aLeRtInCiDeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/roleManagementAlerts/alerts/unifiedRoleManagementAlertIdValue/alertIncidents/unifiedRoleManagementAlertIncidentIdValue",
			Expected: &IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId{
				UnifiedRoleManagementAlertId:         "unifiedRoleManagementAlertIdValue",
				UnifiedRoleManagementAlertIncidentId: "unifiedRoleManagementAlertIncidentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/roleManagementAlerts/alerts/unifiedRoleManagementAlertIdValue/alertIncidents/unifiedRoleManagementAlertIncidentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/rOlEmAnAgEmEnTaLeRtS/aLeRtS/uNiFiEdRoLeMaNaGeMeNtAlErTiDvAlUe/aLeRtInCiDeNtS/uNiFiEdRoLeMaNaGeMeNtAlErTiNcIdEnTiDvAlUe",
			Expected: &IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId{
				UnifiedRoleManagementAlertId:         "uNiFiEdRoLeMaNaGeMeNtAlErTiDvAlUe",
				UnifiedRoleManagementAlertIncidentId: "uNiFiEdRoLeMaNaGeMeNtAlErTiNcIdEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/rOlEmAnAgEmEnTaLeRtS/aLeRtS/uNiFiEdRoLeMaNaGeMeNtAlErTiDvAlUe/aLeRtInCiDeNtS/uNiFiEdRoLeMaNaGeMeNtAlErTiNcIdEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceRoleManagementAlertAlertIdAlertIncidentIDInsensitively(v.Input)
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

		if actual.UnifiedRoleManagementAlertIncidentId != v.Expected.UnifiedRoleManagementAlertIncidentId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleManagementAlertIncidentId", v.Expected.UnifiedRoleManagementAlertIncidentId, actual.UnifiedRoleManagementAlertIncidentId)
		}

	}
}

func TestSegmentsForIdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId(t *testing.T) {
	segments := IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
