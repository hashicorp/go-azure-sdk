package privilegedaccesgroupassignmentscheduleinstanceactivatedusing

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId{}

func TestNewIdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceID(t *testing.T) {
	id := NewIdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceID("privilegedAccessGroupAssignmentScheduleInstanceIdValue")

	if id.PrivilegedAccessGroupAssignmentScheduleInstanceId != "privilegedAccessGroupAssignmentScheduleInstanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PrivilegedAccessGroupAssignmentScheduleInstanceId'", id.PrivilegedAccessGroupAssignmentScheduleInstanceId, "privilegedAccessGroupAssignmentScheduleInstanceIdValue")
	}
}

func TestFormatIdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceID(t *testing.T) {
	actual := NewIdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceID("privilegedAccessGroupAssignmentScheduleInstanceIdValue").ID()
	expected := "/identityGovernance/privilegedAccess/group/assignmentScheduleInstances/privilegedAccessGroupAssignmentScheduleInstanceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId
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
			Input: "/identityGovernance/privilegedAccess/group/assignmentScheduleInstances/privilegedAccessGroupAssignmentScheduleInstanceIdValue",
			Expected: &IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId{
				PrivilegedAccessGroupAssignmentScheduleInstanceId: "privilegedAccessGroupAssignmentScheduleInstanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/assignmentScheduleInstances/privilegedAccessGroupAssignmentScheduleInstanceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceID(v.Input)
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

func TestParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId
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
			Input: "/identityGovernance/privilegedAccess/group/assignmentScheduleInstances/privilegedAccessGroupAssignmentScheduleInstanceIdValue",
			Expected: &IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId{
				PrivilegedAccessGroupAssignmentScheduleInstanceId: "privilegedAccessGroupAssignmentScheduleInstanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/assignmentScheduleInstances/privilegedAccessGroupAssignmentScheduleInstanceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/aSsIgNmEnTsChEdUlEiNsTaNcEs/pRiViLeGeDaCcEsSgRoUpAsSiGnMeNtScHeDuLeInStAnCeIdVaLuE",
			Expected: &IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId{
				PrivilegedAccessGroupAssignmentScheduleInstanceId: "pRiViLeGeDaCcEsSgRoUpAsSiGnMeNtScHeDuLeInStAnCeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/aSsIgNmEnTsChEdUlEiNsTaNcEs/pRiViLeGeDaCcEsSgRoUpAsSiGnMeNtScHeDuLeInStAnCeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceIDInsensitively(v.Input)
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

func TestSegmentsForIdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId(t *testing.T) {
	segments := IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
