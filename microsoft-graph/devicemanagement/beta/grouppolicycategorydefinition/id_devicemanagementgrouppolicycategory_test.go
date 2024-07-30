package grouppolicycategorydefinition

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyCategoryId{}

func TestNewDeviceManagementGroupPolicyCategoryID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyCategoryID("groupPolicyCategoryIdValue")

	if id.GroupPolicyCategoryId != "groupPolicyCategoryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyCategoryId'", id.GroupPolicyCategoryId, "groupPolicyCategoryIdValue")
	}
}

func TestFormatDeviceManagementGroupPolicyCategoryID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyCategoryID("groupPolicyCategoryIdValue").ID()
	expected := "/deviceManagement/groupPolicyCategories/groupPolicyCategoryIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementGroupPolicyCategoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyCategoryId
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
			// Valid URI
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryIdValue",
			Expected: &DeviceManagementGroupPolicyCategoryId{
				GroupPolicyCategoryId: "groupPolicyCategoryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyCategoryID(v.Input)
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

	}
}

func TestParseDeviceManagementGroupPolicyCategoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyCategoryId
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
			// Valid URI
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryIdValue",
			Expected: &DeviceManagementGroupPolicyCategoryId{
				GroupPolicyCategoryId: "groupPolicyCategoryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCaTeGoRiEs/gRoUpPoLiCyCaTeGoRyIdVaLuE",
			Expected: &DeviceManagementGroupPolicyCategoryId{
				GroupPolicyCategoryId: "gRoUpPoLiCyCaTeGoRyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCaTeGoRiEs/gRoUpPoLiCyCaTeGoRyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyCategoryIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementGroupPolicyCategoryId(t *testing.T) {
	segments := DeviceManagementGroupPolicyCategoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementGroupPolicyCategoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
