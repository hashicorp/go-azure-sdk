package policymobiledevicemanagementpolicyincludedgroup

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyMobileDeviceManagementPolicyId{}

func TestNewPolicyMobileDeviceManagementPolicyID(t *testing.T) {
	id := NewPolicyMobileDeviceManagementPolicyID("mobilityManagementPolicyIdValue")

	if id.MobilityManagementPolicyId != "mobilityManagementPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MobilityManagementPolicyId'", id.MobilityManagementPolicyId, "mobilityManagementPolicyIdValue")
	}
}

func TestFormatPolicyMobileDeviceManagementPolicyID(t *testing.T) {
	actual := NewPolicyMobileDeviceManagementPolicyID("mobilityManagementPolicyIdValue").ID()
	expected := "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyMobileDeviceManagementPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyMobileDeviceManagementPolicyId
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
			// Valid URI
			Input: "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyIdValue",
			Expected: &PolicyMobileDeviceManagementPolicyId{
				MobilityManagementPolicyId: "mobilityManagementPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyMobileDeviceManagementPolicyID(v.Input)
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

func TestParsePolicyMobileDeviceManagementPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyMobileDeviceManagementPolicyId
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
			// Valid URI
			Input: "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyIdValue",
			Expected: &PolicyMobileDeviceManagementPolicyId{
				MobilityManagementPolicyId: "mobilityManagementPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/mobileDeviceManagementPolicies/mobilityManagementPolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEdEvIcEmAnAgEmEnTpOlIcIeS/mObIlItYmAnAgEmEnTpOlIcYiDvAlUe",
			Expected: &PolicyMobileDeviceManagementPolicyId{
				MobilityManagementPolicyId: "mObIlItYmAnAgEmEnTpOlIcYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/mObIlEdEvIcEmAnAgEmEnTpOlIcIeS/mObIlItYmAnAgEmEnTpOlIcYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyMobileDeviceManagementPolicyIDInsensitively(v.Input)
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

func TestSegmentsForPolicyMobileDeviceManagementPolicyId(t *testing.T) {
	segments := PolicyMobileDeviceManagementPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyMobileDeviceManagementPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
