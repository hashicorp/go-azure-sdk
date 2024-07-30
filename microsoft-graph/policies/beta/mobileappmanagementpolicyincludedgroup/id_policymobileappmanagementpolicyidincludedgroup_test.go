package mobileappmanagementpolicyincludedgroup

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyMobileAppManagementPolicyIdIncludedGroupId{}

func TestNewPolicyMobileAppManagementPolicyIdIncludedGroupID(t *testing.T) {
	id := NewPolicyMobileAppManagementPolicyIdIncludedGroupID("mobilityManagementPolicyIdValue", "groupIdValue")

	if id.MobilityManagementPolicyId != "mobilityManagementPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MobilityManagementPolicyId'", id.MobilityManagementPolicyId, "mobilityManagementPolicyIdValue")
	}

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}
}

func TestFormatPolicyMobileAppManagementPolicyIdIncludedGroupID(t *testing.T) {
	actual := NewPolicyMobileAppManagementPolicyIdIncludedGroupID("mobilityManagementPolicyIdValue", "groupIdValue").ID()
	expected := "/policies/mobileAppManagementPolicies/mobilityManagementPolicyIdValue/includedGroups/groupIdValue"
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
			Input: "/policies/mobileAppManagementPolicies/mobilityManagementPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/mobileAppManagementPolicies/mobilityManagementPolicyIdValue/includedGroups",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/mobileAppManagementPolicies/mobilityManagementPolicyIdValue/includedGroups/groupIdValue",
			Expected: &PolicyMobileAppManagementPolicyIdIncludedGroupId{
				MobilityManagementPolicyId: "mobilityManagementPolicyIdValue",
				GroupId:                    "groupIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/mobileAppManagementPolicies/mobilityManagementPolicyIdValue/includedGroups/groupIdValue/extra",
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
			Input: "/policies/mobileAppManagementPolicies/mobilityManagementPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEaPpMaNaGeMeNtPoLiCiEs/mObIlItYmAnAgEmEnTpOlIcYiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/mobileAppManagementPolicies/mobilityManagementPolicyIdValue/includedGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEaPpMaNaGeMeNtPoLiCiEs/mObIlItYmAnAgEmEnTpOlIcYiDvAlUe/iNcLuDeDgRoUpS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/mobileAppManagementPolicies/mobilityManagementPolicyIdValue/includedGroups/groupIdValue",
			Expected: &PolicyMobileAppManagementPolicyIdIncludedGroupId{
				MobilityManagementPolicyId: "mobilityManagementPolicyIdValue",
				GroupId:                    "groupIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/mobileAppManagementPolicies/mobilityManagementPolicyIdValue/includedGroups/groupIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEaPpMaNaGeMeNtPoLiCiEs/mObIlItYmAnAgEmEnTpOlIcYiDvAlUe/iNcLuDeDgRoUpS/gRoUpIdVaLuE",
			Expected: &PolicyMobileAppManagementPolicyIdIncludedGroupId{
				MobilityManagementPolicyId: "mObIlItYmAnAgEmEnTpOlIcYiDvAlUe",
				GroupId:                    "gRoUpIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEaPpMaNaGeMeNtPoLiCiEs/mObIlItYmAnAgEmEnTpOlIcYiDvAlUe/iNcLuDeDgRoUpS/gRoUpIdVaLuE/extra",
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
