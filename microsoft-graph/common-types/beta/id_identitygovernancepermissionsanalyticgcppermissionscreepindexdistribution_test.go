package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId{}

func TestNewIdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionID(t *testing.T) {
	id := NewIdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionID("permissionsCreepIndexDistributionId")

	if id.PermissionsCreepIndexDistributionId != "permissionsCreepIndexDistributionId" {
		t.Fatalf("Expected %q but got %q for Segment 'PermissionsCreepIndexDistributionId'", id.PermissionsCreepIndexDistributionId, "permissionsCreepIndexDistributionId")
	}
}

func TestFormatIdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionID(t *testing.T) {
	actual := NewIdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionID("permissionsCreepIndexDistributionId").ID()
	expected := "/identityGovernance/permissionsAnalytics/gcp/permissionsCreepIndexDistributions/permissionsCreepIndexDistributionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId
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
			Input: "/identityGovernance/permissionsAnalytics",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/permissionsAnalytics/gcp",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/permissionsAnalytics/gcp/permissionsCreepIndexDistributions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/permissionsAnalytics/gcp/permissionsCreepIndexDistributions/permissionsCreepIndexDistributionId",
			Expected: &IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId{
				PermissionsCreepIndexDistributionId: "permissionsCreepIndexDistributionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/permissionsAnalytics/gcp/permissionsCreepIndexDistributions/permissionsCreepIndexDistributionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PermissionsCreepIndexDistributionId != v.Expected.PermissionsCreepIndexDistributionId {
			t.Fatalf("Expected %q but got %q for PermissionsCreepIndexDistributionId", v.Expected.PermissionsCreepIndexDistributionId, actual.PermissionsCreepIndexDistributionId)
		}

	}
}

func TestParseIdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId
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
			Input: "/identityGovernance/permissionsAnalytics",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsAnAlYtIcS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/permissionsAnalytics/gcp",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsAnAlYtIcS/gCp",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/permissionsAnalytics/gcp/permissionsCreepIndexDistributions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsAnAlYtIcS/gCp/pErMiSsIoNsCrEePiNdExDiStRiBuTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/permissionsAnalytics/gcp/permissionsCreepIndexDistributions/permissionsCreepIndexDistributionId",
			Expected: &IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId{
				PermissionsCreepIndexDistributionId: "permissionsCreepIndexDistributionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/permissionsAnalytics/gcp/permissionsCreepIndexDistributions/permissionsCreepIndexDistributionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsAnAlYtIcS/gCp/pErMiSsIoNsCrEePiNdExDiStRiBuTiOnS/pErMiSsIoNsCrEePiNdExDiStRiBuTiOnId",
			Expected: &IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId{
				PermissionsCreepIndexDistributionId: "pErMiSsIoNsCrEePiNdExDiStRiBuTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsAnAlYtIcS/gCp/pErMiSsIoNsCrEePiNdExDiStRiBuTiOnS/pErMiSsIoNsCrEePiNdExDiStRiBuTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PermissionsCreepIndexDistributionId != v.Expected.PermissionsCreepIndexDistributionId {
			t.Fatalf("Expected %q but got %q for PermissionsCreepIndexDistributionId", v.Expected.PermissionsCreepIndexDistributionId, actual.PermissionsCreepIndexDistributionId)
		}

	}
}

func TestSegmentsForIdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId(t *testing.T) {
	segments := IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
