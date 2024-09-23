package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId{}

func TestNewIdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceID(t *testing.T) {
	id := NewIdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceID("privilegedAccessGroupAssignmentScheduleInstanceId")

	if id.PrivilegedAccessGroupAssignmentScheduleInstanceId != "privilegedAccessGroupAssignmentScheduleInstanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'PrivilegedAccessGroupAssignmentScheduleInstanceId'", id.PrivilegedAccessGroupAssignmentScheduleInstanceId, "privilegedAccessGroupAssignmentScheduleInstanceId")
	}
}

func TestFormatIdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceID(t *testing.T) {
	actual := NewIdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceID("privilegedAccessGroupAssignmentScheduleInstanceId").ID()
	expected := "/identityGovernance/privilegedAccess/group/assignmentScheduleInstances/privilegedAccessGroupAssignmentScheduleInstanceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId
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
			Input: "/identityGovernance/privilegedAccess/group/assignmentScheduleInstances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/privilegedAccess/group/assignmentScheduleInstances/privilegedAccessGroupAssignmentScheduleInstanceId",
			Expected: &IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId{
				PrivilegedAccessGroupAssignmentScheduleInstanceId: "privilegedAccessGroupAssignmentScheduleInstanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/assignmentScheduleInstances/privilegedAccessGroupAssignmentScheduleInstanceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrivilegedAccessGroupAssignmentScheduleInstanceId != v.Expected.PrivilegedAccessGroupAssignmentScheduleInstanceId {
			t.Fatalf("Expected %q but got %q for PrivilegedAccessGroupAssignmentScheduleInstanceId", v.Expected.PrivilegedAccessGroupAssignmentScheduleInstanceId, actual.PrivilegedAccessGroupAssignmentScheduleInstanceId)
		}

	}
}

func TestParseIdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId
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
			Input: "/identityGovernance/privilegedAccess/group/assignmentScheduleInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/aSsIgNmEnTsChEdUlEiNsTaNcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/privilegedAccess/group/assignmentScheduleInstances/privilegedAccessGroupAssignmentScheduleInstanceId",
			Expected: &IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId{
				PrivilegedAccessGroupAssignmentScheduleInstanceId: "privilegedAccessGroupAssignmentScheduleInstanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/assignmentScheduleInstances/privilegedAccessGroupAssignmentScheduleInstanceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/aSsIgNmEnTsChEdUlEiNsTaNcEs/pRiViLeGeDaCcEsSgRoUpAsSiGnMeNtScHeDuLeInStAnCeId",
			Expected: &IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId{
				PrivilegedAccessGroupAssignmentScheduleInstanceId: "pRiViLeGeDaCcEsSgRoUpAsSiGnMeNtScHeDuLeInStAnCeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/aSsIgNmEnTsChEdUlEiNsTaNcEs/pRiViLeGeDaCcEsSgRoUpAsSiGnMeNtScHeDuLeInStAnCeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrivilegedAccessGroupAssignmentScheduleInstanceId != v.Expected.PrivilegedAccessGroupAssignmentScheduleInstanceId {
			t.Fatalf("Expected %q but got %q for PrivilegedAccessGroupAssignmentScheduleInstanceId", v.Expected.PrivilegedAccessGroupAssignmentScheduleInstanceId, actual.PrivilegedAccessGroupAssignmentScheduleInstanceId)
		}

	}
}

func TestSegmentsForIdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId(t *testing.T) {
	segments := IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
