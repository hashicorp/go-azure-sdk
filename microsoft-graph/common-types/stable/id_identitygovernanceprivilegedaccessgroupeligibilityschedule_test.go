package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId{}

func TestNewIdentityGovernancePrivilegedAccessGroupEligibilityScheduleID(t *testing.T) {
	id := NewIdentityGovernancePrivilegedAccessGroupEligibilityScheduleID("privilegedAccessGroupEligibilityScheduleId")

	if id.PrivilegedAccessGroupEligibilityScheduleId != "privilegedAccessGroupEligibilityScheduleId" {
		t.Fatalf("Expected %q but got %q for Segment 'PrivilegedAccessGroupEligibilityScheduleId'", id.PrivilegedAccessGroupEligibilityScheduleId, "privilegedAccessGroupEligibilityScheduleId")
	}
}

func TestFormatIdentityGovernancePrivilegedAccessGroupEligibilityScheduleID(t *testing.T) {
	actual := NewIdentityGovernancePrivilegedAccessGroupEligibilityScheduleID("privilegedAccessGroupEligibilityScheduleId").ID()
	expected := "/identityGovernance/privilegedAccess/group/eligibilitySchedules/privilegedAccessGroupEligibilityScheduleId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId
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
			Input: "/identityGovernance/privilegedAccess",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/privilegedAccess/group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/privilegedAccess/group/eligibilitySchedules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/privilegedAccess/group/eligibilitySchedules/privilegedAccessGroupEligibilityScheduleId",
			Expected: &IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId{
				PrivilegedAccessGroupEligibilityScheduleId: "privilegedAccessGroupEligibilityScheduleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/eligibilitySchedules/privilegedAccessGroupEligibilityScheduleId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrivilegedAccessGroupEligibilityScheduleId != v.Expected.PrivilegedAccessGroupEligibilityScheduleId {
			t.Fatalf("Expected %q but got %q for PrivilegedAccessGroupEligibilityScheduleId", v.Expected.PrivilegedAccessGroupEligibilityScheduleId, actual.PrivilegedAccessGroupEligibilityScheduleId)
		}

	}
}

func TestParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId
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
			Input: "/identityGovernance/privilegedAccess",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/privilegedAccess/group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/privilegedAccess/group/eligibilitySchedules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/eLiGiBiLiTyScHeDuLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/privilegedAccess/group/eligibilitySchedules/privilegedAccessGroupEligibilityScheduleId",
			Expected: &IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId{
				PrivilegedAccessGroupEligibilityScheduleId: "privilegedAccessGroupEligibilityScheduleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/eligibilitySchedules/privilegedAccessGroupEligibilityScheduleId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/eLiGiBiLiTyScHeDuLeS/pRiViLeGeDaCcEsSgRoUpElIgIbIlItYsChEdUlEiD",
			Expected: &IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId{
				PrivilegedAccessGroupEligibilityScheduleId: "pRiViLeGeDaCcEsSgRoUpElIgIbIlItYsChEdUlEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/eLiGiBiLiTyScHeDuLeS/pRiViLeGeDaCcEsSgRoUpElIgIbIlItYsChEdUlEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrivilegedAccessGroupEligibilityScheduleId != v.Expected.PrivilegedAccessGroupEligibilityScheduleId {
			t.Fatalf("Expected %q but got %q for PrivilegedAccessGroupEligibilityScheduleId", v.Expected.PrivilegedAccessGroupEligibilityScheduleId, actual.PrivilegedAccessGroupEligibilityScheduleId)
		}

	}
}

func TestSegmentsForIdentityGovernancePrivilegedAccessGroupEligibilityScheduleId(t *testing.T) {
	segments := IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
