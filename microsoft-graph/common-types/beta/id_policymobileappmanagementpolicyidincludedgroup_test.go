package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyMobileAppManagementPolicyIdIncludedGroupId{}

func TestNewPolicyMobileAppManagementPolicyIdIncludedGroupID(t *testing.T) {
	id := NewPolicyMobileAppManagementPolicyIdIncludedGroupID("mobilityManagementPolicyId", "groupId")

	if id.MobilityManagementPolicyId != "mobilityManagementPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'MobilityManagementPolicyId'", id.MobilityManagementPolicyId, "mobilityManagementPolicyId")
	}

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}
}

func TestFormatPolicyMobileAppManagementPolicyIdIncludedGroupID(t *testing.T) {
	actual := NewPolicyMobileAppManagementPolicyIdIncludedGroupID("mobilityManagementPolicyId", "groupId").ID()
	expected := "/policies/mobileAppManagementPolicies/mobilityManagementPolicyId/includedGroups/groupId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyMobileAppManagementPolicyIdIncludedGroupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyMobileAppManagementPolicyIdIncludedGroupId
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
			Input: "/policies/mobileAppManagementPolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/mobileAppManagementPolicies/mobilityManagementPolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/mobileAppManagementPolicies/mobilityManagementPolicyId/includedGroups",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/mobileAppManagementPolicies/mobilityManagementPolicyId/includedGroups/groupId",
			Expected: &PolicyMobileAppManagementPolicyIdIncludedGroupId{
				MobilityManagementPolicyId: "mobilityManagementPolicyId",
				GroupId:                    "groupId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/mobileAppManagementPolicies/mobilityManagementPolicyId/includedGroups/groupId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyMobileAppManagementPolicyIdIncludedGroupID(v.Input)
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

func TestParsePolicyMobileAppManagementPolicyIdIncludedGroupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyMobileAppManagementPolicyIdIncludedGroupId
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
			Input: "/policies/mobileAppManagementPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEaPpMaNaGeMeNtPoLiCiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/mobileAppManagementPolicies/mobilityManagementPolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEaPpMaNaGeMeNtPoLiCiEs/mObIlItYmAnAgEmEnTpOlIcYiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/mobileAppManagementPolicies/mobilityManagementPolicyId/includedGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEaPpMaNaGeMeNtPoLiCiEs/mObIlItYmAnAgEmEnTpOlIcYiD/iNcLuDeDgRoUpS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/mobileAppManagementPolicies/mobilityManagementPolicyId/includedGroups/groupId",
			Expected: &PolicyMobileAppManagementPolicyIdIncludedGroupId{
				MobilityManagementPolicyId: "mobilityManagementPolicyId",
				GroupId:                    "groupId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/mobileAppManagementPolicies/mobilityManagementPolicyId/includedGroups/groupId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEaPpMaNaGeMeNtPoLiCiEs/mObIlItYmAnAgEmEnTpOlIcYiD/iNcLuDeDgRoUpS/gRoUpId",
			Expected: &PolicyMobileAppManagementPolicyIdIncludedGroupId{
				MobilityManagementPolicyId: "mObIlItYmAnAgEmEnTpOlIcYiD",
				GroupId:                    "gRoUpId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEaPpMaNaGeMeNtPoLiCiEs/mObIlItYmAnAgEmEnTpOlIcYiD/iNcLuDeDgRoUpS/gRoUpId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyMobileAppManagementPolicyIdIncludedGroupIDInsensitively(v.Input)
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

func TestSegmentsForPolicyMobileAppManagementPolicyIdIncludedGroupId(t *testing.T) {
	segments := PolicyMobileAppManagementPolicyIdIncludedGroupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyMobileAppManagementPolicyIdIncludedGroupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
