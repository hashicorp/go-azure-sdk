package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId{}

func TestNewIdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionID(t *testing.T) {
	id := NewIdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionID("permissionsCreepIndexDistributionIdValue")

	if id.PermissionsCreepIndexDistributionId != "permissionsCreepIndexDistributionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PermissionsCreepIndexDistributionId'", id.PermissionsCreepIndexDistributionId, "permissionsCreepIndexDistributionIdValue")
	}
}

func TestFormatIdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionID(t *testing.T) {
	actual := NewIdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionID("permissionsCreepIndexDistributionIdValue").ID()
	expected := "/identityGovernance/permissionsAnalytics/aws/permissionsCreepIndexDistributions/permissionsCreepIndexDistributionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId
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
			Input: "/identityGovernance/permissionsAnalytics/aws",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/permissionsAnalytics/aws/permissionsCreepIndexDistributions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/permissionsAnalytics/aws/permissionsCreepIndexDistributions/permissionsCreepIndexDistributionIdValue",
			Expected: &IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId{
				PermissionsCreepIndexDistributionId: "permissionsCreepIndexDistributionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/permissionsAnalytics/aws/permissionsCreepIndexDistributions/permissionsCreepIndexDistributionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionID(v.Input)
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

func TestParseIdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId
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
			Input: "/identityGovernance/permissionsAnalytics/aws",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsAnAlYtIcS/aWs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/permissionsAnalytics/aws/permissionsCreepIndexDistributions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsAnAlYtIcS/aWs/pErMiSsIoNsCrEePiNdExDiStRiBuTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/permissionsAnalytics/aws/permissionsCreepIndexDistributions/permissionsCreepIndexDistributionIdValue",
			Expected: &IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId{
				PermissionsCreepIndexDistributionId: "permissionsCreepIndexDistributionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/permissionsAnalytics/aws/permissionsCreepIndexDistributions/permissionsCreepIndexDistributionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsAnAlYtIcS/aWs/pErMiSsIoNsCrEePiNdExDiStRiBuTiOnS/pErMiSsIoNsCrEePiNdExDiStRiBuTiOnIdVaLuE",
			Expected: &IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId{
				PermissionsCreepIndexDistributionId: "pErMiSsIoNsCrEePiNdExDiStRiBuTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsAnAlYtIcS/aWs/pErMiSsIoNsCrEePiNdExDiStRiBuTiOnS/pErMiSsIoNsCrEePiNdExDiStRiBuTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionIDInsensitively(v.Input)
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

func TestSegmentsForIdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId(t *testing.T) {
	segments := IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernancePermissionsAnalyticAwPermissionsCreepIndexDistributionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
