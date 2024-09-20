package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyObjectFileId{}

func TestNewDeviceManagementGroupPolicyObjectFileID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyObjectFileID("groupPolicyObjectFileId")

	if id.GroupPolicyObjectFileId != "groupPolicyObjectFileId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyObjectFileId'", id.GroupPolicyObjectFileId, "groupPolicyObjectFileId")
	}
}

func TestFormatDeviceManagementGroupPolicyObjectFileID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyObjectFileID("groupPolicyObjectFileId").ID()
	expected := "/deviceManagement/groupPolicyObjectFiles/groupPolicyObjectFileId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementGroupPolicyObjectFileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyObjectFileId
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
			Input: "/deviceManagement/groupPolicyObjectFiles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyObjectFiles/groupPolicyObjectFileId",
			Expected: &DeviceManagementGroupPolicyObjectFileId{
				GroupPolicyObjectFileId: "groupPolicyObjectFileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyObjectFiles/groupPolicyObjectFileId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyObjectFileID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupPolicyObjectFileId != v.Expected.GroupPolicyObjectFileId {
			t.Fatalf("Expected %q but got %q for GroupPolicyObjectFileId", v.Expected.GroupPolicyObjectFileId, actual.GroupPolicyObjectFileId)
		}

	}
}

func TestParseDeviceManagementGroupPolicyObjectFileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyObjectFileId
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
			Input: "/deviceManagement/groupPolicyObjectFiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyObJeCtFiLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyObjectFiles/groupPolicyObjectFileId",
			Expected: &DeviceManagementGroupPolicyObjectFileId{
				GroupPolicyObjectFileId: "groupPolicyObjectFileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyObjectFiles/groupPolicyObjectFileId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyObJeCtFiLeS/gRoUpPoLiCyObJeCtFiLeId",
			Expected: &DeviceManagementGroupPolicyObjectFileId{
				GroupPolicyObjectFileId: "gRoUpPoLiCyObJeCtFiLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyObJeCtFiLeS/gRoUpPoLiCyObJeCtFiLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyObjectFileIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupPolicyObjectFileId != v.Expected.GroupPolicyObjectFileId {
			t.Fatalf("Expected %q but got %q for GroupPolicyObjectFileId", v.Expected.GroupPolicyObjectFileId, actual.GroupPolicyObjectFileId)
		}

	}
}

func TestSegmentsForDeviceManagementGroupPolicyObjectFileId(t *testing.T) {
	segments := DeviceManagementGroupPolicyObjectFileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementGroupPolicyObjectFileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
