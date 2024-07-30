package privilegedaccesgroupeligibilityschedulegroup

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId{}

func TestNewIdentityGovernancePrivilegedAccesGroupEligibilityScheduleID(t *testing.T) {
	id := NewIdentityGovernancePrivilegedAccesGroupEligibilityScheduleID("privilegedAccessGroupEligibilityScheduleIdValue")

	if id.PrivilegedAccessGroupEligibilityScheduleId != "privilegedAccessGroupEligibilityScheduleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PrivilegedAccessGroupEligibilityScheduleId'", id.PrivilegedAccessGroupEligibilityScheduleId, "privilegedAccessGroupEligibilityScheduleIdValue")
	}
}

func TestFormatIdentityGovernancePrivilegedAccesGroupEligibilityScheduleID(t *testing.T) {
	actual := NewIdentityGovernancePrivilegedAccesGroupEligibilityScheduleID("privilegedAccessGroupEligibilityScheduleIdValue").ID()
	expected := "/identityGovernance/privilegedAccess/group/eligibilitySchedules/privilegedAccessGroupEligibilityScheduleIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId
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
			Input: "/identityGovernance/privilegedAccess/group/eligibilitySchedules/privilegedAccessGroupEligibilityScheduleIdValue",
			Expected: &IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId{
				PrivilegedAccessGroupEligibilityScheduleId: "privilegedAccessGroupEligibilityScheduleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/eligibilitySchedules/privilegedAccessGroupEligibilityScheduleIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleID(v.Input)
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

func TestParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId
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
			Input: "/identityGovernance/privilegedAccess/group/eligibilitySchedules/privilegedAccessGroupEligibilityScheduleIdValue",
			Expected: &IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId{
				PrivilegedAccessGroupEligibilityScheduleId: "privilegedAccessGroupEligibilityScheduleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/eligibilitySchedules/privilegedAccessGroupEligibilityScheduleIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/eLiGiBiLiTyScHeDuLeS/pRiViLeGeDaCcEsSgRoUpElIgIbIlItYsChEdUlEiDvAlUe",
			Expected: &IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId{
				PrivilegedAccessGroupEligibilityScheduleId: "pRiViLeGeDaCcEsSgRoUpElIgIbIlItYsChEdUlEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/eLiGiBiLiTyScHeDuLeS/pRiViLeGeDaCcEsSgRoUpElIgIbIlItYsChEdUlEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleIDInsensitively(v.Input)
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

func TestSegmentsForIdentityGovernancePrivilegedAccesGroupEligibilityScheduleId(t *testing.T) {
	segments := IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
