package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyMobileDeviceManagementPolicyIdIncludedGroupId{}

func TestNewPolicyMobileDeviceManagementPolicyIdIncludedGroupID(t *testing.T) {
	id := NewPolicyMobileDeviceManagementPolicyIdIncludedGroupID("mobilityManagementPolicyId", "groupId")

	if id.MobilityManagementPolicyId != "mobilityManagementPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'MobilityManagementPolicyId'", id.MobilityManagementPolicyId, "mobilityManagementPolicyId")
	}

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}
}

func TestFormatPolicyMobileDeviceManagementPolicyIdIncludedGroupID(t *testing.T) {
	actual := NewPolicyMobileDeviceManagementPolicyIdIncludedGroupID("mobilityManagementPolicyId", "groupId").ID()
	expected := "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyId/includedGroups/groupId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyMobileDeviceManagementPolicyIdIncludedGroupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyMobileDeviceManagementPolicyIdIncludedGroupId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/mobileDeviceManagementPolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyId/includedGroups",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyId/includedGroups/groupId",
			Expected: &PolicyMobileDeviceManagementPolicyIdIncludedGroupId{
				MobilityManagementPolicyId: "mobilityManagementPolicyId",
				GroupId:                    "groupId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyId/includedGroups/groupId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyMobileDeviceManagementPolicyIdIncludedGroupID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MobilityManagementPolicyId != v.Expected.MobilityManagementPolicyId {
			t.Fatalf("Expected %q but got %q for MobilityManagementPolicyId", v.Expected.MobilityManagementPolicyId, actual.MobilityManagementPolicyId)
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

	}
}

func TestParsePolicyMobileDeviceManagementPolicyIdIncludedGroupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyMobileDeviceManagementPolicyIdIncludedGroupId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/mobileDeviceManagementPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEdEvIcEmAnAgEmEnTpOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEdEvIcEmAnAgEmEnTpOlIcIeS/mObIlItYmAnAgEmEnTpOlIcYiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyId/includedGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEdEvIcEmAnAgEmEnTpOlIcIeS/mObIlItYmAnAgEmEnTpOlIcYiD/iNcLuDeDgRoUpS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyId/includedGroups/groupId",
			Expected: &PolicyMobileDeviceManagementPolicyIdIncludedGroupId{
				MobilityManagementPolicyId: "mobilityManagementPolicyId",
				GroupId:                    "groupId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyId/includedGroups/groupId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEdEvIcEmAnAgEmEnTpOlIcIeS/mObIlItYmAnAgEmEnTpOlIcYiD/iNcLuDeDgRoUpS/gRoUpId",
			Expected: &PolicyMobileDeviceManagementPolicyIdIncludedGroupId{
				MobilityManagementPolicyId: "mObIlItYmAnAgEmEnTpOlIcYiD",
				GroupId:                    "gRoUpId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEdEvIcEmAnAgEmEnTpOlIcIeS/mObIlItYmAnAgEmEnTpOlIcYiD/iNcLuDeDgRoUpS/gRoUpId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyMobileDeviceManagementPolicyIdIncludedGroupIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MobilityManagementPolicyId != v.Expected.MobilityManagementPolicyId {
			t.Fatalf("Expected %q but got %q for MobilityManagementPolicyId", v.Expected.MobilityManagementPolicyId, actual.MobilityManagementPolicyId)
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

	}
}

func TestSegmentsForPolicyMobileDeviceManagementPolicyIdIncludedGroupId(t *testing.T) {
	segments := PolicyMobileDeviceManagementPolicyIdIncludedGroupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyMobileDeviceManagementPolicyIdIncludedGroupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
