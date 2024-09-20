package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId{}

func TestNewIdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionID(t *testing.T) {
	id := NewIdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionID("permissionsCreepIndexDistributionId")

	if id.PermissionsCreepIndexDistributionId != "permissionsCreepIndexDistributionId" {
		t.Fatalf("Expected %q but got %q for Segment 'PermissionsCreepIndexDistributionId'", id.PermissionsCreepIndexDistributionId, "permissionsCreepIndexDistributionId")
	}
}

func TestFormatIdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionID(t *testing.T) {
	actual := NewIdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionID("permissionsCreepIndexDistributionId").ID()
	expected := "/identityGovernance/permissionsAnalytics/azure/permissionsCreepIndexDistributions/permissionsCreepIndexDistributionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId
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
			Input: "/identityGovernance/permissionsAnalytics/azure",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/permissionsAnalytics/azure/permissionsCreepIndexDistributions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/permissionsAnalytics/azure/permissionsCreepIndexDistributions/permissionsCreepIndexDistributionId",
			Expected: &IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId{
				PermissionsCreepIndexDistributionId: "permissionsCreepIndexDistributionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/permissionsAnalytics/azure/permissionsCreepIndexDistributions/permissionsCreepIndexDistributionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionID(v.Input)
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

func TestParseIdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId
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
			Input: "/identityGovernance/permissionsAnalytics/azure",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsAnAlYtIcS/aZuRe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/permissionsAnalytics/azure/permissionsCreepIndexDistributions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsAnAlYtIcS/aZuRe/pErMiSsIoNsCrEePiNdExDiStRiBuTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/permissionsAnalytics/azure/permissionsCreepIndexDistributions/permissionsCreepIndexDistributionId",
			Expected: &IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId{
				PermissionsCreepIndexDistributionId: "permissionsCreepIndexDistributionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/permissionsAnalytics/azure/permissionsCreepIndexDistributions/permissionsCreepIndexDistributionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsAnAlYtIcS/aZuRe/pErMiSsIoNsCrEePiNdExDiStRiBuTiOnS/pErMiSsIoNsCrEePiNdExDiStRiBuTiOnId",
			Expected: &IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId{
				PermissionsCreepIndexDistributionId: "pErMiSsIoNsCrEePiNdExDiStRiBuTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsAnAlYtIcS/aZuRe/pErMiSsIoNsCrEePiNdExDiStRiBuTiOnS/pErMiSsIoNsCrEePiNdExDiStRiBuTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionIDInsensitively(v.Input)
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

func TestSegmentsForIdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId(t *testing.T) {
	segments := IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernancePermissionsAnalyticAzurePermissionsCreepIndexDistributionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
