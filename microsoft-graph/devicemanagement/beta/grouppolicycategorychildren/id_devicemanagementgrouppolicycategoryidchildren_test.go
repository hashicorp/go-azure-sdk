package grouppolicycategorychildren

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyCategoryIdChildrenId{}

func TestNewDeviceManagementGroupPolicyCategoryIdChildrenID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyCategoryIdChildrenID("groupPolicyCategoryIdValue", "groupPolicyCategoryId1Value")

	if id.GroupPolicyCategoryId != "groupPolicyCategoryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyCategoryId'", id.GroupPolicyCategoryId, "groupPolicyCategoryIdValue")
	}

	if id.GroupPolicyCategoryId1 != "groupPolicyCategoryId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyCategoryId1'", id.GroupPolicyCategoryId1, "groupPolicyCategoryId1Value")
	}
}

func TestFormatDeviceManagementGroupPolicyCategoryIdChildrenID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyCategoryIdChildrenID("groupPolicyCategoryIdValue", "groupPolicyCategoryId1Value").ID()
	expected := "/deviceManagement/groupPolicyCategories/groupPolicyCategoryIdValue/children/groupPolicyCategoryId1Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementGroupPolicyCategoryIdChildrenID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyCategoryIdChildrenId
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
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryIdValue/children",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryIdValue/children/groupPolicyCategoryId1Value",
			Expected: &DeviceManagementGroupPolicyCategoryIdChildrenId{
				GroupPolicyCategoryId:  "groupPolicyCategoryIdValue",
				GroupPolicyCategoryId1: "groupPolicyCategoryId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryIdValue/children/groupPolicyCategoryId1Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyCategoryIdChildrenID(v.Input)
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

func TestParseDeviceManagementGroupPolicyCategoryIdChildrenIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyCategoryIdChildrenId
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
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCaTeGoRiEs/gRoUpPoLiCyCaTeGoRyIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryIdValue/children",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCaTeGoRiEs/gRoUpPoLiCyCaTeGoRyIdVaLuE/cHiLdReN",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryIdValue/children/groupPolicyCategoryId1Value",
			Expected: &DeviceManagementGroupPolicyCategoryIdChildrenId{
				GroupPolicyCategoryId:  "groupPolicyCategoryIdValue",
				GroupPolicyCategoryId1: "groupPolicyCategoryId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryIdValue/children/groupPolicyCategoryId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCaTeGoRiEs/gRoUpPoLiCyCaTeGoRyIdVaLuE/cHiLdReN/gRoUpPoLiCyCaTeGoRyId1vAlUe",
			Expected: &DeviceManagementGroupPolicyCategoryIdChildrenId{
				GroupPolicyCategoryId:  "gRoUpPoLiCyCaTeGoRyIdVaLuE",
				GroupPolicyCategoryId1: "gRoUpPoLiCyCaTeGoRyId1vAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCaTeGoRiEs/gRoUpPoLiCyCaTeGoRyIdVaLuE/cHiLdReN/gRoUpPoLiCyCaTeGoRyId1vAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyCategoryIdChildrenIDInsensitively(v.Input)
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

func TestSegmentsForDeviceManagementGroupPolicyCategoryIdChildrenId(t *testing.T) {
	segments := DeviceManagementGroupPolicyCategoryIdChildrenId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementGroupPolicyCategoryIdChildrenId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
