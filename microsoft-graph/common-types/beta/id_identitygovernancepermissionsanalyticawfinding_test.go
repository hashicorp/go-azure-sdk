package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePermissionsAnalyticAwFindingId{}

func TestNewIdentityGovernancePermissionsAnalyticAwFindingID(t *testing.T) {
	id := NewIdentityGovernancePermissionsAnalyticAwFindingID("findingId")

	if id.FindingId != "findingId" {
		t.Fatalf("Expected %q but got %q for Segment 'FindingId'", id.FindingId, "findingId")
	}
}

func TestFormatIdentityGovernancePermissionsAnalyticAwFindingID(t *testing.T) {
	actual := NewIdentityGovernancePermissionsAnalyticAwFindingID("findingId").ID()
	expected := "/identityGovernance/permissionsAnalytics/aws/findings/findingId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernancePermissionsAnalyticAwFindingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePermissionsAnalyticAwFindingId
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
			Input: "/identityGovernance/permissionsAnalytics/aws/findings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/permissionsAnalytics/aws/findings/findingId",
			Expected: &IdentityGovernancePermissionsAnalyticAwFindingId{
				FindingId: "findingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/permissionsAnalytics/aws/findings/findingId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePermissionsAnalyticAwFindingID(v.Input)
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

func TestParseIdentityGovernancePermissionsAnalyticAwFindingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePermissionsAnalyticAwFindingId
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
			Input: "/identityGovernance/permissionsAnalytics/aws/findings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsAnAlYtIcS/aWs/fInDiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/permissionsAnalytics/aws/findings/findingId",
			Expected: &IdentityGovernancePermissionsAnalyticAwFindingId{
				FindingId: "findingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/permissionsAnalytics/aws/findings/findingId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsAnAlYtIcS/aWs/fInDiNgS/fInDiNgId",
			Expected: &IdentityGovernancePermissionsAnalyticAwFindingId{
				FindingId: "fInDiNgId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsAnAlYtIcS/aWs/fInDiNgS/fInDiNgId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePermissionsAnalyticAwFindingIDInsensitively(v.Input)
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

func TestSegmentsForIdentityGovernancePermissionsAnalyticAwFindingId(t *testing.T) {
	segments := IdentityGovernancePermissionsAnalyticAwFindingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernancePermissionsAnalyticAwFindingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
