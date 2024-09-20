package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId{}

func TestNewIdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestID(t *testing.T) {
	id := NewIdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestID("privilegedAccessGroupEligibilityScheduleRequestId")

	if id.PrivilegedAccessGroupEligibilityScheduleRequestId != "privilegedAccessGroupEligibilityScheduleRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'PrivilegedAccessGroupEligibilityScheduleRequestId'", id.PrivilegedAccessGroupEligibilityScheduleRequestId, "privilegedAccessGroupEligibilityScheduleRequestId")
	}
}

func TestFormatIdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestID(t *testing.T) {
	actual := NewIdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestID("privilegedAccessGroupEligibilityScheduleRequestId").ID()
	expected := "/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests/privilegedAccessGroupEligibilityScheduleRequestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId
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
			Input: "/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests/privilegedAccessGroupEligibilityScheduleRequestId",
			Expected: &IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId{
				PrivilegedAccessGroupEligibilityScheduleRequestId: "privilegedAccessGroupEligibilityScheduleRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests/privilegedAccessGroupEligibilityScheduleRequestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrivilegedAccessGroupEligibilityScheduleRequestId != v.Expected.PrivilegedAccessGroupEligibilityScheduleRequestId {
			t.Fatalf("Expected %q but got %q for PrivilegedAccessGroupEligibilityScheduleRequestId", v.Expected.PrivilegedAccessGroupEligibilityScheduleRequestId, actual.PrivilegedAccessGroupEligibilityScheduleRequestId)
		}

	}
}

func TestParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId
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
			Input: "/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/eLiGiBiLiTyScHeDuLeReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests/privilegedAccessGroupEligibilityScheduleRequestId",
			Expected: &IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId{
				PrivilegedAccessGroupEligibilityScheduleRequestId: "privilegedAccessGroupEligibilityScheduleRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests/privilegedAccessGroupEligibilityScheduleRequestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/eLiGiBiLiTyScHeDuLeReQuEsTs/pRiViLeGeDaCcEsSgRoUpElIgIbIlItYsChEdUlErEqUeStId",
			Expected: &IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId{
				PrivilegedAccessGroupEligibilityScheduleRequestId: "pRiViLeGeDaCcEsSgRoUpElIgIbIlItYsChEdUlErEqUeStId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/eLiGiBiLiTyScHeDuLeReQuEsTs/pRiViLeGeDaCcEsSgRoUpElIgIbIlItYsChEdUlErEqUeStId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrivilegedAccessGroupEligibilityScheduleRequestId != v.Expected.PrivilegedAccessGroupEligibilityScheduleRequestId {
			t.Fatalf("Expected %q but got %q for PrivilegedAccessGroupEligibilityScheduleRequestId", v.Expected.PrivilegedAccessGroupEligibilityScheduleRequestId, actual.PrivilegedAccessGroupEligibilityScheduleRequestId)
		}

	}
}

func TestSegmentsForIdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId(t *testing.T) {
	segments := IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
