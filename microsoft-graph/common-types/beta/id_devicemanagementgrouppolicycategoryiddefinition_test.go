package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyCategoryIdDefinitionId{}

func TestNewDeviceManagementGroupPolicyCategoryIdDefinitionID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyCategoryIdDefinitionID("groupPolicyCategoryId", "groupPolicyDefinitionId")

	if id.GroupPolicyCategoryId != "groupPolicyCategoryId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyCategoryId'", id.GroupPolicyCategoryId, "groupPolicyCategoryId")
	}

	if id.GroupPolicyDefinitionId != "groupPolicyDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyDefinitionId'", id.GroupPolicyDefinitionId, "groupPolicyDefinitionId")
	}
}

func TestFormatDeviceManagementGroupPolicyCategoryIdDefinitionID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyCategoryIdDefinitionID("groupPolicyCategoryId", "groupPolicyDefinitionId").ID()
	expected := "/deviceManagement/groupPolicyCategories/groupPolicyCategoryId/definitions/groupPolicyDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementGroupPolicyCategoryIdDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyCategoryIdDefinitionId
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
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryId/definitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryId/definitions/groupPolicyDefinitionId",
			Expected: &DeviceManagementGroupPolicyCategoryIdDefinitionId{
				GroupPolicyCategoryId:   "groupPolicyCategoryId",
				GroupPolicyDefinitionId: "groupPolicyDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryId/definitions/groupPolicyDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyCategoryIdDefinitionID(v.Input)
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

		if actual.GroupPolicyDefinitionId != v.Expected.GroupPolicyDefinitionId {
			t.Fatalf("Expected %q but got %q for GroupPolicyDefinitionId", v.Expected.GroupPolicyDefinitionId, actual.GroupPolicyDefinitionId)
		}

	}
}

func TestParseDeviceManagementGroupPolicyCategoryIdDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyCategoryIdDefinitionId
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
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryId/definitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCaTeGoRiEs/gRoUpPoLiCyCaTeGoRyId/dEfInItIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryId/definitions/groupPolicyDefinitionId",
			Expected: &DeviceManagementGroupPolicyCategoryIdDefinitionId{
				GroupPolicyCategoryId:   "groupPolicyCategoryId",
				GroupPolicyDefinitionId: "groupPolicyDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyCategories/groupPolicyCategoryId/definitions/groupPolicyDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCaTeGoRiEs/gRoUpPoLiCyCaTeGoRyId/dEfInItIoNs/gRoUpPoLiCyDeFiNiTiOnId",
			Expected: &DeviceManagementGroupPolicyCategoryIdDefinitionId{
				GroupPolicyCategoryId:   "gRoUpPoLiCyCaTeGoRyId",
				GroupPolicyDefinitionId: "gRoUpPoLiCyDeFiNiTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCaTeGoRiEs/gRoUpPoLiCyCaTeGoRyId/dEfInItIoNs/gRoUpPoLiCyDeFiNiTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyCategoryIdDefinitionIDInsensitively(v.Input)
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

		if actual.GroupPolicyDefinitionId != v.Expected.GroupPolicyDefinitionId {
			t.Fatalf("Expected %q but got %q for GroupPolicyDefinitionId", v.Expected.GroupPolicyDefinitionId, actual.GroupPolicyDefinitionId)
		}

	}
}

func TestSegmentsForDeviceManagementGroupPolicyCategoryIdDefinitionId(t *testing.T) {
	segments := DeviceManagementGroupPolicyCategoryIdDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementGroupPolicyCategoryIdDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
