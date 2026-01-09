package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceRoleManagementAlertOperationId{}

func TestNewIdentityGovernanceRoleManagementAlertOperationID(t *testing.T) {
	id := NewIdentityGovernanceRoleManagementAlertOperationID("longRunningOperationId")

	if id.LongRunningOperationId != "longRunningOperationId" {
		t.Fatalf("Expected %q but got %q for Segment 'LongRunningOperationId'", id.LongRunningOperationId, "longRunningOperationId")
	}
}

func TestFormatIdentityGovernanceRoleManagementAlertOperationID(t *testing.T) {
	actual := NewIdentityGovernanceRoleManagementAlertOperationID("longRunningOperationId").ID()
	expected := "/identityGovernance/roleManagementAlerts/operations/longRunningOperationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceRoleManagementAlertOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceRoleManagementAlertOperationId
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
			Input: "/identityGovernance/roleManagementAlerts/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/roleManagementAlerts/operations/longRunningOperationId",
			Expected: &IdentityGovernanceRoleManagementAlertOperationId{
				LongRunningOperationId: "longRunningOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/roleManagementAlerts/operations/longRunningOperationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceRoleManagementAlertOperationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.LongRunningOperationId != v.Expected.LongRunningOperationId {
			t.Fatalf("Expected %q but got %q for LongRunningOperationId", v.Expected.LongRunningOperationId, actual.LongRunningOperationId)
		}

	}
}

func TestParseIdentityGovernanceRoleManagementAlertOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceRoleManagementAlertOperationId
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
			Input: "/identityGovernance/roleManagementAlerts/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/rOlEmAnAgEmEnTaLeRtS/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/roleManagementAlerts/operations/longRunningOperationId",
			Expected: &IdentityGovernanceRoleManagementAlertOperationId{
				LongRunningOperationId: "longRunningOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/roleManagementAlerts/operations/longRunningOperationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/rOlEmAnAgEmEnTaLeRtS/oPeRaTiOnS/lOnGrUnNiNgOpErAtIoNiD",
			Expected: &IdentityGovernanceRoleManagementAlertOperationId{
				LongRunningOperationId: "lOnGrUnNiNgOpErAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/rOlEmAnAgEmEnTaLeRtS/oPeRaTiOnS/lOnGrUnNiNgOpErAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceRoleManagementAlertOperationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.LongRunningOperationId != v.Expected.LongRunningOperationId {
			t.Fatalf("Expected %q but got %q for LongRunningOperationId", v.Expected.LongRunningOperationId, actual.LongRunningOperationId)
		}

	}
}

func TestSegmentsForIdentityGovernanceRoleManagementAlertOperationId(t *testing.T) {
	segments := IdentityGovernanceRoleManagementAlertOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceRoleManagementAlertOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
