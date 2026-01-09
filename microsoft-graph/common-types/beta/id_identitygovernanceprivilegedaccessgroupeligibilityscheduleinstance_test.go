package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId{}

func TestNewIdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceID(t *testing.T) {
	id := NewIdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceID("privilegedAccessGroupEligibilityScheduleInstanceId")

	if id.PrivilegedAccessGroupEligibilityScheduleInstanceId != "privilegedAccessGroupEligibilityScheduleInstanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'PrivilegedAccessGroupEligibilityScheduleInstanceId'", id.PrivilegedAccessGroupEligibilityScheduleInstanceId, "privilegedAccessGroupEligibilityScheduleInstanceId")
	}
}

func TestFormatIdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceID(t *testing.T) {
	actual := NewIdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceID("privilegedAccessGroupEligibilityScheduleInstanceId").ID()
	expected := "/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances/privilegedAccessGroupEligibilityScheduleInstanceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId
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
			Input: "/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances/privilegedAccessGroupEligibilityScheduleInstanceId",
			Expected: &IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId{
				PrivilegedAccessGroupEligibilityScheduleInstanceId: "privilegedAccessGroupEligibilityScheduleInstanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances/privilegedAccessGroupEligibilityScheduleInstanceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrivilegedAccessGroupEligibilityScheduleInstanceId != v.Expected.PrivilegedAccessGroupEligibilityScheduleInstanceId {
			t.Fatalf("Expected %q but got %q for PrivilegedAccessGroupEligibilityScheduleInstanceId", v.Expected.PrivilegedAccessGroupEligibilityScheduleInstanceId, actual.PrivilegedAccessGroupEligibilityScheduleInstanceId)
		}

	}
}

func TestParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId
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
			Input: "/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/eLiGiBiLiTyScHeDuLeInStAnCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances/privilegedAccessGroupEligibilityScheduleInstanceId",
			Expected: &IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId{
				PrivilegedAccessGroupEligibilityScheduleInstanceId: "privilegedAccessGroupEligibilityScheduleInstanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances/privilegedAccessGroupEligibilityScheduleInstanceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/eLiGiBiLiTyScHeDuLeInStAnCeS/pRiViLeGeDaCcEsSgRoUpElIgIbIlItYsChEdUlEiNsTaNcEiD",
			Expected: &IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId{
				PrivilegedAccessGroupEligibilityScheduleInstanceId: "pRiViLeGeDaCcEsSgRoUpElIgIbIlItYsChEdUlEiNsTaNcEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/eLiGiBiLiTyScHeDuLeInStAnCeS/pRiViLeGeDaCcEsSgRoUpElIgIbIlItYsChEdUlEiNsTaNcEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrivilegedAccessGroupEligibilityScheduleInstanceId != v.Expected.PrivilegedAccessGroupEligibilityScheduleInstanceId {
			t.Fatalf("Expected %q but got %q for PrivilegedAccessGroupEligibilityScheduleInstanceId", v.Expected.PrivilegedAccessGroupEligibilityScheduleInstanceId, actual.PrivilegedAccessGroupEligibilityScheduleInstanceId)
		}

	}
}

func TestSegmentsForIdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId(t *testing.T) {
	segments := IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
