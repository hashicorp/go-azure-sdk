package privilegedaccesgroupassignmentschedulerequesttargetschedule

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId{}

func TestNewIdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestID(t *testing.T) {
	id := NewIdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestID("privilegedAccessGroupAssignmentScheduleRequestIdValue")

	if id.PrivilegedAccessGroupAssignmentScheduleRequestId != "privilegedAccessGroupAssignmentScheduleRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PrivilegedAccessGroupAssignmentScheduleRequestId'", id.PrivilegedAccessGroupAssignmentScheduleRequestId, "privilegedAccessGroupAssignmentScheduleRequestIdValue")
	}
}

func TestFormatIdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestID(t *testing.T) {
	actual := NewIdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestID("privilegedAccessGroupAssignmentScheduleRequestIdValue").ID()
	expected := "/identityGovernance/privilegedAccess/group/assignmentScheduleRequests/privilegedAccessGroupAssignmentScheduleRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId
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
			Input: "/identityGovernance/privilegedAccess/group/assignmentScheduleRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/privilegedAccess/group/assignmentScheduleRequests/privilegedAccessGroupAssignmentScheduleRequestIdValue",
			Expected: &IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId{
				PrivilegedAccessGroupAssignmentScheduleRequestId: "privilegedAccessGroupAssignmentScheduleRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/assignmentScheduleRequests/privilegedAccessGroupAssignmentScheduleRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrivilegedAccessGroupAssignmentScheduleRequestId != v.Expected.PrivilegedAccessGroupAssignmentScheduleRequestId {
			t.Fatalf("Expected %q but got %q for PrivilegedAccessGroupAssignmentScheduleRequestId", v.Expected.PrivilegedAccessGroupAssignmentScheduleRequestId, actual.PrivilegedAccessGroupAssignmentScheduleRequestId)
		}

	}
}

func TestParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId
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
			Input: "/identityGovernance/privilegedAccess/group/assignmentScheduleRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/aSsIgNmEnTsChEdUlErEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/privilegedAccess/group/assignmentScheduleRequests/privilegedAccessGroupAssignmentScheduleRequestIdValue",
			Expected: &IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId{
				PrivilegedAccessGroupAssignmentScheduleRequestId: "privilegedAccessGroupAssignmentScheduleRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/assignmentScheduleRequests/privilegedAccessGroupAssignmentScheduleRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/aSsIgNmEnTsChEdUlErEqUeStS/pRiViLeGeDaCcEsSgRoUpAsSiGnMeNtScHeDuLeReQuEsTiDvAlUe",
			Expected: &IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId{
				PrivilegedAccessGroupAssignmentScheduleRequestId: "pRiViLeGeDaCcEsSgRoUpAsSiGnMeNtScHeDuLeReQuEsTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/aSsIgNmEnTsChEdUlErEqUeStS/pRiViLeGeDaCcEsSgRoUpAsSiGnMeNtScHeDuLeReQuEsTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrivilegedAccessGroupAssignmentScheduleRequestId != v.Expected.PrivilegedAccessGroupAssignmentScheduleRequestId {
			t.Fatalf("Expected %q but got %q for PrivilegedAccessGroupAssignmentScheduleRequestId", v.Expected.PrivilegedAccessGroupAssignmentScheduleRequestId, actual.PrivilegedAccessGroupAssignmentScheduleRequestId)
		}

	}
}

func TestSegmentsForIdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId(t *testing.T) {
	segments := IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
