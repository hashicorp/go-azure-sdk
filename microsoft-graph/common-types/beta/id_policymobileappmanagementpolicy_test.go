package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyMobileAppManagementPolicyId{}

func TestNewPolicyMobileAppManagementPolicyID(t *testing.T) {
	id := NewPolicyMobileAppManagementPolicyID("mobilityManagementPolicyId")

	if id.MobilityManagementPolicyId != "mobilityManagementPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'MobilityManagementPolicyId'", id.MobilityManagementPolicyId, "mobilityManagementPolicyId")
	}
}

func TestFormatPolicyMobileAppManagementPolicyID(t *testing.T) {
	actual := NewPolicyMobileAppManagementPolicyID("mobilityManagementPolicyId").ID()
	expected := "/policies/mobileAppManagementPolicies/mobilityManagementPolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyMobileAppManagementPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyMobileAppManagementPolicyId
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
			// Valid URI
			Input: "/policies/mobileAppManagementPolicies/mobilityManagementPolicyId",
			Expected: &PolicyMobileAppManagementPolicyId{
				MobilityManagementPolicyId: "mobilityManagementPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/mobileAppManagementPolicies/mobilityManagementPolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyMobileAppManagementPolicyID(v.Input)
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

	}
}

func TestParsePolicyMobileAppManagementPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyMobileAppManagementPolicyId
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
			// Valid URI
			Input: "/policies/mobileAppManagementPolicies/mobilityManagementPolicyId",
			Expected: &PolicyMobileAppManagementPolicyId{
				MobilityManagementPolicyId: "mobilityManagementPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/mobileAppManagementPolicies/mobilityManagementPolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEaPpMaNaGeMeNtPoLiCiEs/mObIlItYmAnAgEmEnTpOlIcYiD",
			Expected: &PolicyMobileAppManagementPolicyId{
				MobilityManagementPolicyId: "mObIlItYmAnAgEmEnTpOlIcYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEaPpMaNaGeMeNtPoLiCiEs/mObIlItYmAnAgEmEnTpOlIcYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyMobileAppManagementPolicyIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForPolicyMobileAppManagementPolicyId(t *testing.T) {
	segments := PolicyMobileAppManagementPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyMobileAppManagementPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
