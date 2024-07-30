package privilegedaccesgroupassignmentschedulegroupserviceprovisioningerror

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId{}

func TestNewIdentityGovernancePrivilegedAccesGroupAssignmentScheduleID(t *testing.T) {
	id := NewIdentityGovernancePrivilegedAccesGroupAssignmentScheduleID("privilegedAccessGroupAssignmentScheduleIdValue")

	if id.PrivilegedAccessGroupAssignmentScheduleId != "privilegedAccessGroupAssignmentScheduleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PrivilegedAccessGroupAssignmentScheduleId'", id.PrivilegedAccessGroupAssignmentScheduleId, "privilegedAccessGroupAssignmentScheduleIdValue")
	}
}

func TestFormatIdentityGovernancePrivilegedAccesGroupAssignmentScheduleID(t *testing.T) {
	actual := NewIdentityGovernancePrivilegedAccesGroupAssignmentScheduleID("privilegedAccessGroupAssignmentScheduleIdValue").ID()
	expected := "/identityGovernance/privilegedAccess/group/assignmentSchedules/privilegedAccessGroupAssignmentScheduleIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId
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
			Input: "/identityGovernance/privilegedAccess/group/assignmentSchedules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/privilegedAccess/group/assignmentSchedules/privilegedAccessGroupAssignmentScheduleIdValue",
			Expected: &IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId{
				PrivilegedAccessGroupAssignmentScheduleId: "privilegedAccessGroupAssignmentScheduleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/assignmentSchedules/privilegedAccessGroupAssignmentScheduleIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrivilegedAccessGroupAssignmentScheduleId != v.Expected.PrivilegedAccessGroupAssignmentScheduleId {
			t.Fatalf("Expected %q but got %q for PrivilegedAccessGroupAssignmentScheduleId", v.Expected.PrivilegedAccessGroupAssignmentScheduleId, actual.PrivilegedAccessGroupAssignmentScheduleId)
		}

	}
}

func TestParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId
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
			Input: "/identityGovernance/privilegedAccess/group/assignmentSchedules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/aSsIgNmEnTsChEdUlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/privilegedAccess/group/assignmentSchedules/privilegedAccessGroupAssignmentScheduleIdValue",
			Expected: &IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId{
				PrivilegedAccessGroupAssignmentScheduleId: "privilegedAccessGroupAssignmentScheduleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/assignmentSchedules/privilegedAccessGroupAssignmentScheduleIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/aSsIgNmEnTsChEdUlEs/pRiViLeGeDaCcEsSgRoUpAsSiGnMeNtScHeDuLeIdVaLuE",
			Expected: &IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId{
				PrivilegedAccessGroupAssignmentScheduleId: "pRiViLeGeDaCcEsSgRoUpAsSiGnMeNtScHeDuLeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/aSsIgNmEnTsChEdUlEs/pRiViLeGeDaCcEsSgRoUpAsSiGnMeNtScHeDuLeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrivilegedAccessGroupAssignmentScheduleId != v.Expected.PrivilegedAccessGroupAssignmentScheduleId {
			t.Fatalf("Expected %q but got %q for PrivilegedAccessGroupAssignmentScheduleId", v.Expected.PrivilegedAccessGroupAssignmentScheduleId, actual.PrivilegedAccessGroupAssignmentScheduleId)
		}

	}
}

func TestSegmentsForIdentityGovernancePrivilegedAccesGroupAssignmentScheduleId(t *testing.T) {
	segments := IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
