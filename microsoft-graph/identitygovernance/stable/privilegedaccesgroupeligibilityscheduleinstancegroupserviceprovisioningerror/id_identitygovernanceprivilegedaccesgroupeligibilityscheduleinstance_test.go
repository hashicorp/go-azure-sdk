package privilegedaccesgroupeligibilityscheduleinstancegroupserviceprovisioningerror

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId{}

func TestNewIdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceID(t *testing.T) {
	id := NewIdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceID("privilegedAccessGroupEligibilityScheduleInstanceIdValue")

	if id.PrivilegedAccessGroupEligibilityScheduleInstanceId != "privilegedAccessGroupEligibilityScheduleInstanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PrivilegedAccessGroupEligibilityScheduleInstanceId'", id.PrivilegedAccessGroupEligibilityScheduleInstanceId, "privilegedAccessGroupEligibilityScheduleInstanceIdValue")
	}
}

func TestFormatIdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceID(t *testing.T) {
	actual := NewIdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceID("privilegedAccessGroupEligibilityScheduleInstanceIdValue").ID()
	expected := "/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances/privilegedAccessGroupEligibilityScheduleInstanceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId
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
			Input: "/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances/privilegedAccessGroupEligibilityScheduleInstanceIdValue",
			Expected: &IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId{
				PrivilegedAccessGroupEligibilityScheduleInstanceId: "privilegedAccessGroupEligibilityScheduleInstanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances/privilegedAccessGroupEligibilityScheduleInstanceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceID(v.Input)
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

func TestParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId
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
			Input: "/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances/privilegedAccessGroupEligibilityScheduleInstanceIdValue",
			Expected: &IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId{
				PrivilegedAccessGroupEligibilityScheduleInstanceId: "privilegedAccessGroupEligibilityScheduleInstanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances/privilegedAccessGroupEligibilityScheduleInstanceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/eLiGiBiLiTyScHeDuLeInStAnCeS/pRiViLeGeDaCcEsSgRoUpElIgIbIlItYsChEdUlEiNsTaNcEiDvAlUe",
			Expected: &IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId{
				PrivilegedAccessGroupEligibilityScheduleInstanceId: "pRiViLeGeDaCcEsSgRoUpElIgIbIlItYsChEdUlEiNsTaNcEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/eLiGiBiLiTyScHeDuLeInStAnCeS/pRiViLeGeDaCcEsSgRoUpElIgIbIlItYsChEdUlEiNsTaNcEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceIDInsensitively(v.Input)
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

func TestSegmentsForIdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId(t *testing.T) {
	segments := IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
