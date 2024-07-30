package privilegedaccesgroupeligibilityschedulerequestprincipal

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId{}

func TestNewIdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestID(t *testing.T) {
	id := NewIdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestID("privilegedAccessGroupEligibilityScheduleRequestIdValue")

	if id.PrivilegedAccessGroupEligibilityScheduleRequestId != "privilegedAccessGroupEligibilityScheduleRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PrivilegedAccessGroupEligibilityScheduleRequestId'", id.PrivilegedAccessGroupEligibilityScheduleRequestId, "privilegedAccessGroupEligibilityScheduleRequestIdValue")
	}
}

func TestFormatIdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestID(t *testing.T) {
	actual := NewIdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestID("privilegedAccessGroupEligibilityScheduleRequestIdValue").ID()
	expected := "/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests/privilegedAccessGroupEligibilityScheduleRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId
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
			Input: "/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests/privilegedAccessGroupEligibilityScheduleRequestIdValue",
			Expected: &IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId{
				PrivilegedAccessGroupEligibilityScheduleRequestId: "privilegedAccessGroupEligibilityScheduleRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests/privilegedAccessGroupEligibilityScheduleRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestID(v.Input)
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

func TestParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId
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
			Input: "/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests/privilegedAccessGroupEligibilityScheduleRequestIdValue",
			Expected: &IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId{
				PrivilegedAccessGroupEligibilityScheduleRequestId: "privilegedAccessGroupEligibilityScheduleRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests/privilegedAccessGroupEligibilityScheduleRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/eLiGiBiLiTyScHeDuLeReQuEsTs/pRiViLeGeDaCcEsSgRoUpElIgIbIlItYsChEdUlErEqUeStIdVaLuE",
			Expected: &IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId{
				PrivilegedAccessGroupEligibilityScheduleRequestId: "pRiViLeGeDaCcEsSgRoUpElIgIbIlItYsChEdUlErEqUeStIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/eLiGiBiLiTyScHeDuLeReQuEsTs/pRiViLeGeDaCcEsSgRoUpElIgIbIlItYsChEdUlErEqUeStIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestIDInsensitively(v.Input)
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

func TestSegmentsForIdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId(t *testing.T) {
	segments := IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
