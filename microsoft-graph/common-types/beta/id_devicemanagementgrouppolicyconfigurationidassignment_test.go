package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyConfigurationIdAssignmentId{}

func TestNewDeviceManagementGroupPolicyConfigurationIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyConfigurationIdAssignmentID("groupPolicyConfigurationId", "groupPolicyConfigurationAssignmentId")

	if id.GroupPolicyConfigurationId != "groupPolicyConfigurationId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyConfigurationId'", id.GroupPolicyConfigurationId, "groupPolicyConfigurationId")
	}

	if id.GroupPolicyConfigurationAssignmentId != "groupPolicyConfigurationAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyConfigurationAssignmentId'", id.GroupPolicyConfigurationAssignmentId, "groupPolicyConfigurationAssignmentId")
	}
}

func TestFormatDeviceManagementGroupPolicyConfigurationIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyConfigurationIdAssignmentID("groupPolicyConfigurationId", "groupPolicyConfigurationAssignmentId").ID()
	expected := "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId/assignments/groupPolicyConfigurationAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementGroupPolicyConfigurationIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyConfigurationIdAssignmentId
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
			Input: "/deviceManagement/groupPolicyConfigurations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId/assignments/groupPolicyConfigurationAssignmentId",
			Expected: &DeviceManagementGroupPolicyConfigurationIdAssignmentId{
				GroupPolicyConfigurationId:           "groupPolicyConfigurationId",
				GroupPolicyConfigurationAssignmentId: "groupPolicyConfigurationAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId/assignments/groupPolicyConfigurationAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyConfigurationIdAssignmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupPolicyConfigurationId != v.Expected.GroupPolicyConfigurationId {
			t.Fatalf("Expected %q but got %q for GroupPolicyConfigurationId", v.Expected.GroupPolicyConfigurationId, actual.GroupPolicyConfigurationId)
		}

		if actual.GroupPolicyConfigurationAssignmentId != v.Expected.GroupPolicyConfigurationAssignmentId {
			t.Fatalf("Expected %q but got %q for GroupPolicyConfigurationAssignmentId", v.Expected.GroupPolicyConfigurationAssignmentId, actual.GroupPolicyConfigurationAssignmentId)
		}

	}
}

func TestParseDeviceManagementGroupPolicyConfigurationIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyConfigurationIdAssignmentId
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
			Input: "/deviceManagement/groupPolicyConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs/gRoUpPoLiCyCoNfIgUrAtIoNiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs/gRoUpPoLiCyCoNfIgUrAtIoNiD/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId/assignments/groupPolicyConfigurationAssignmentId",
			Expected: &DeviceManagementGroupPolicyConfigurationIdAssignmentId{
				GroupPolicyConfigurationId:           "groupPolicyConfigurationId",
				GroupPolicyConfigurationAssignmentId: "groupPolicyConfigurationAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId/assignments/groupPolicyConfigurationAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs/gRoUpPoLiCyCoNfIgUrAtIoNiD/aSsIgNmEnTs/gRoUpPoLiCyCoNfIgUrAtIoNaSsIgNmEnTiD",
			Expected: &DeviceManagementGroupPolicyConfigurationIdAssignmentId{
				GroupPolicyConfigurationId:           "gRoUpPoLiCyCoNfIgUrAtIoNiD",
				GroupPolicyConfigurationAssignmentId: "gRoUpPoLiCyCoNfIgUrAtIoNaSsIgNmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs/gRoUpPoLiCyCoNfIgUrAtIoNiD/aSsIgNmEnTs/gRoUpPoLiCyCoNfIgUrAtIoNaSsIgNmEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyConfigurationIdAssignmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupPolicyConfigurationId != v.Expected.GroupPolicyConfigurationId {
			t.Fatalf("Expected %q but got %q for GroupPolicyConfigurationId", v.Expected.GroupPolicyConfigurationId, actual.GroupPolicyConfigurationId)
		}

		if actual.GroupPolicyConfigurationAssignmentId != v.Expected.GroupPolicyConfigurationAssignmentId {
			t.Fatalf("Expected %q but got %q for GroupPolicyConfigurationAssignmentId", v.Expected.GroupPolicyConfigurationAssignmentId, actual.GroupPolicyConfigurationAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementGroupPolicyConfigurationIdAssignmentId(t *testing.T) {
	segments := DeviceManagementGroupPolicyConfigurationIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementGroupPolicyConfigurationIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
