package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyDefinitionFileId{}

func TestNewDeviceManagementGroupPolicyDefinitionFileID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyDefinitionFileID("groupPolicyDefinitionFileId")

	if id.GroupPolicyDefinitionFileId != "groupPolicyDefinitionFileId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyDefinitionFileId'", id.GroupPolicyDefinitionFileId, "groupPolicyDefinitionFileId")
	}
}

func TestFormatDeviceManagementGroupPolicyDefinitionFileID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyDefinitionFileID("groupPolicyDefinitionFileId").ID()
	expected := "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementGroupPolicyDefinitionFileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyDefinitionFileId
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
			Input: "/deviceManagement/groupPolicyDefinitionFiles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileId",
			Expected: &DeviceManagementGroupPolicyDefinitionFileId{
				GroupPolicyDefinitionFileId: "groupPolicyDefinitionFileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyDefinitionFileID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupPolicyDefinitionFileId != v.Expected.GroupPolicyDefinitionFileId {
			t.Fatalf("Expected %q but got %q for GroupPolicyDefinitionFileId", v.Expected.GroupPolicyDefinitionFileId, actual.GroupPolicyDefinitionFileId)
		}

	}
}

func TestParseDeviceManagementGroupPolicyDefinitionFileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyDefinitionFileId
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
			Input: "/deviceManagement/groupPolicyDefinitionFiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnFiLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileId",
			Expected: &DeviceManagementGroupPolicyDefinitionFileId{
				GroupPolicyDefinitionFileId: "groupPolicyDefinitionFileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnFiLeS/gRoUpPoLiCyDeFiNiTiOnFiLeId",
			Expected: &DeviceManagementGroupPolicyDefinitionFileId{
				GroupPolicyDefinitionFileId: "gRoUpPoLiCyDeFiNiTiOnFiLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnFiLeS/gRoUpPoLiCyDeFiNiTiOnFiLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyDefinitionFileIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupPolicyDefinitionFileId != v.Expected.GroupPolicyDefinitionFileId {
			t.Fatalf("Expected %q but got %q for GroupPolicyDefinitionFileId", v.Expected.GroupPolicyDefinitionFileId, actual.GroupPolicyDefinitionFileId)
		}

	}
}

func TestSegmentsForDeviceManagementGroupPolicyDefinitionFileId(t *testing.T) {
	segments := DeviceManagementGroupPolicyDefinitionFileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementGroupPolicyDefinitionFileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
