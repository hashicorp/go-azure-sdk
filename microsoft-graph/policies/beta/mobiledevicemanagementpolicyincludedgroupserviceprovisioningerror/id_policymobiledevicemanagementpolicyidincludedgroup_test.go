package mobiledevicemanagementpolicyincludedgroupserviceprovisioningerror

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyMobileDeviceManagementPolicyIdIncludedGroupId{}

func TestNewPolicyMobileDeviceManagementPolicyIdIncludedGroupID(t *testing.T) {
	id := NewPolicyMobileDeviceManagementPolicyIdIncludedGroupID("mobilityManagementPolicyIdValue", "groupIdValue")

	if id.MobilityManagementPolicyId != "mobilityManagementPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MobilityManagementPolicyId'", id.MobilityManagementPolicyId, "mobilityManagementPolicyIdValue")
	}

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}
}

func TestFormatPolicyMobileDeviceManagementPolicyIdIncludedGroupID(t *testing.T) {
	actual := NewPolicyMobileDeviceManagementPolicyIdIncludedGroupID("mobilityManagementPolicyIdValue", "groupIdValue").ID()
	expected := "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyIdValue/includedGroups/groupIdValue"
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
			Input: "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyIdValue/includedGroups",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyIdValue/includedGroups/groupIdValue",
			Expected: &PolicyMobileDeviceManagementPolicyIdIncludedGroupId{
				MobilityManagementPolicyId: "mobilityManagementPolicyIdValue",
				GroupId:                    "groupIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyIdValue/includedGroups/groupIdValue/extra",
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
			Input: "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEdEvIcEmAnAgEmEnTpOlIcIeS/mObIlItYmAnAgEmEnTpOlIcYiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyIdValue/includedGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEdEvIcEmAnAgEmEnTpOlIcIeS/mObIlItYmAnAgEmEnTpOlIcYiDvAlUe/iNcLuDeDgRoUpS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyIdValue/includedGroups/groupIdValue",
			Expected: &PolicyMobileDeviceManagementPolicyIdIncludedGroupId{
				MobilityManagementPolicyId: "mobilityManagementPolicyIdValue",
				GroupId:                    "groupIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyIdValue/includedGroups/groupIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEdEvIcEmAnAgEmEnTpOlIcIeS/mObIlItYmAnAgEmEnTpOlIcYiDvAlUe/iNcLuDeDgRoUpS/gRoUpIdVaLuE",
			Expected: &PolicyMobileDeviceManagementPolicyIdIncludedGroupId{
				MobilityManagementPolicyId: "mObIlItYmAnAgEmEnTpOlIcYiDvAlUe",
				GroupId:                    "gRoUpIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEdEvIcEmAnAgEmEnTpOlIcIeS/mObIlItYmAnAgEmEnTpOlIcYiDvAlUe/iNcLuDeDgRoUpS/gRoUpIdVaLuE/extra",
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
