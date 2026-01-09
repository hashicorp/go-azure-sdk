package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyCategoryIdChildId{}

func TestNewDeviceManagementGroupPolicyCategoryIdChildID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyCategoryIdChildID("groupPolicyCategoryId", "groupPolicyCategoryId1")

	if id.GroupPolicyCategoryId != "groupPolicyCategoryId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyCategoryId'", id.GroupPolicyCategoryId, "groupPolicyCategoryId")
	}

	if id.GroupPolicyCategoryId1 != "groupPolicyCategoryId1" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyCategoryId1'", id.GroupPolicyCategoryId1, "groupPolicyCategoryId1")
	}
}

func TestFormatDeviceManagementGroupPolicyCategoryIdChildID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyCategoryIdChildID("groupPolicyCategoryId", "groupPolicyCategoryId1").ID()
	expected := "/deviceManagement/groupPolicyCategories/groupPolicyCategoryId/children/groupPolicyCategoryId1"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementGroupPolicyCategoryIdChildID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyCategoryIdChildId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyCategories",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryId/children",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryId/children/groupPolicyCategoryId1",
			Expected: &DeviceManagementGroupPolicyCategoryIdChildId{
				GroupPolicyCategoryId:  "groupPolicyCategoryId",
				GroupPolicyCategoryId1: "groupPolicyCategoryId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryId/children/groupPolicyCategoryId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyCategoryIdChildID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupPolicyCategoryId != v.Expected.GroupPolicyCategoryId {
			t.Fatalf("Expected %q but got %q for GroupPolicyCategoryId", v.Expected.GroupPolicyCategoryId, actual.GroupPolicyCategoryId)
		}

		if actual.GroupPolicyCategoryId1 != v.Expected.GroupPolicyCategoryId1 {
			t.Fatalf("Expected %q but got %q for GroupPolicyCategoryId1", v.Expected.GroupPolicyCategoryId1, actual.GroupPolicyCategoryId1)
		}

	}
}

func TestParseDeviceManagementGroupPolicyCategoryIdChildIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyCategoryIdChildId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyCategories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCaTeGoRiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCaTeGoRiEs/gRoUpPoLiCyCaTeGoRyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryId/children",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCaTeGoRiEs/gRoUpPoLiCyCaTeGoRyId/cHiLdReN",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryId/children/groupPolicyCategoryId1",
			Expected: &DeviceManagementGroupPolicyCategoryIdChildId{
				GroupPolicyCategoryId:  "groupPolicyCategoryId",
				GroupPolicyCategoryId1: "groupPolicyCategoryId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryId/children/groupPolicyCategoryId1/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCaTeGoRiEs/gRoUpPoLiCyCaTeGoRyId/cHiLdReN/gRoUpPoLiCyCaTeGoRyId1",
			Expected: &DeviceManagementGroupPolicyCategoryIdChildId{
				GroupPolicyCategoryId:  "gRoUpPoLiCyCaTeGoRyId",
				GroupPolicyCategoryId1: "gRoUpPoLiCyCaTeGoRyId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCaTeGoRiEs/gRoUpPoLiCyCaTeGoRyId/cHiLdReN/gRoUpPoLiCyCaTeGoRyId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyCategoryIdChildIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupPolicyCategoryId != v.Expected.GroupPolicyCategoryId {
			t.Fatalf("Expected %q but got %q for GroupPolicyCategoryId", v.Expected.GroupPolicyCategoryId, actual.GroupPolicyCategoryId)
		}

		if actual.GroupPolicyCategoryId1 != v.Expected.GroupPolicyCategoryId1 {
			t.Fatalf("Expected %q but got %q for GroupPolicyCategoryId1", v.Expected.GroupPolicyCategoryId1, actual.GroupPolicyCategoryId1)
		}

	}
}

func TestSegmentsForDeviceManagementGroupPolicyCategoryIdChildId(t *testing.T) {
	segments := DeviceManagementGroupPolicyCategoryIdChildId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementGroupPolicyCategoryIdChildId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
