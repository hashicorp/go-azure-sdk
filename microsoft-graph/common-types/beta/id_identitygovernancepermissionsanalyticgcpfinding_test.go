package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePermissionsAnalyticGcpFindingId{}

func TestNewIdentityGovernancePermissionsAnalyticGcpFindingID(t *testing.T) {
	id := NewIdentityGovernancePermissionsAnalyticGcpFindingID("findingId")

	if id.FindingId != "findingId" {
		t.Fatalf("Expected %q but got %q for Segment 'FindingId'", id.FindingId, "findingId")
	}
}

func TestFormatIdentityGovernancePermissionsAnalyticGcpFindingID(t *testing.T) {
	actual := NewIdentityGovernancePermissionsAnalyticGcpFindingID("findingId").ID()
	expected := "/identityGovernance/permissionsAnalytics/gcp/findings/findingId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernancePermissionsAnalyticGcpFindingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePermissionsAnalyticGcpFindingId
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
			Input: "/identityGovernance/permissionsAnalytics/gcp/findings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/permissionsAnalytics/gcp/findings/findingId",
			Expected: &IdentityGovernancePermissionsAnalyticGcpFindingId{
				FindingId: "findingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/permissionsAnalytics/gcp/findings/findingId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePermissionsAnalyticGcpFindingID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.FindingId != v.Expected.FindingId {
			t.Fatalf("Expected %q but got %q for FindingId", v.Expected.FindingId, actual.FindingId)
		}

	}
}

func TestParseIdentityGovernancePermissionsAnalyticGcpFindingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePermissionsAnalyticGcpFindingId
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
			Input: "/identityGovernance/permissionsAnalytics/gcp/findings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsAnAlYtIcS/gCp/fInDiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/permissionsAnalytics/gcp/findings/findingId",
			Expected: &IdentityGovernancePermissionsAnalyticGcpFindingId{
				FindingId: "findingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/permissionsAnalytics/gcp/findings/findingId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsAnAlYtIcS/gCp/fInDiNgS/fInDiNgId",
			Expected: &IdentityGovernancePermissionsAnalyticGcpFindingId{
				FindingId: "fInDiNgId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsAnAlYtIcS/gCp/fInDiNgS/fInDiNgId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePermissionsAnalyticGcpFindingIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.FindingId != v.Expected.FindingId {
			t.Fatalf("Expected %q but got %q for FindingId", v.Expected.FindingId, actual.FindingId)
		}

	}
}

func TestSegmentsForIdentityGovernancePermissionsAnalyticGcpFindingId(t *testing.T) {
	segments := IdentityGovernancePermissionsAnalyticGcpFindingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernancePermissionsAnalyticGcpFindingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
