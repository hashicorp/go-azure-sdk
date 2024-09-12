package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyConfigurationIdAssignmentId{}

func TestNewDeviceManagementGroupPolicyConfigurationIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyConfigurationIdAssignmentID("groupPolicyConfigurationIdValue", "groupPolicyConfigurationAssignmentIdValue")

	if id.GroupPolicyConfigurationId != "groupPolicyConfigurationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyConfigurationId'", id.GroupPolicyConfigurationId, "groupPolicyConfigurationIdValue")
	}

	if id.GroupPolicyConfigurationAssignmentId != "groupPolicyConfigurationAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyConfigurationAssignmentId'", id.GroupPolicyConfigurationAssignmentId, "groupPolicyConfigurationAssignmentIdValue")
	}
}

func TestFormatDeviceManagementGroupPolicyConfigurationIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyConfigurationIdAssignmentID("groupPolicyConfigurationIdValue", "groupPolicyConfigurationAssignmentIdValue").ID()
	expected := "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue/assignments/groupPolicyConfigurationAssignmentIdValue"
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
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue/assignments/groupPolicyConfigurationAssignmentIdValue",
			Expected: &DeviceManagementGroupPolicyConfigurationIdAssignmentId{
				GroupPolicyConfigurationId:           "groupPolicyConfigurationIdValue",
				GroupPolicyConfigurationAssignmentId: "groupPolicyConfigurationAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue/assignments/groupPolicyConfigurationAssignmentIdValue/extra",
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
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs/gRoUpPoLiCyCoNfIgUrAtIoNiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs/gRoUpPoLiCyCoNfIgUrAtIoNiDvAlUe/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue/assignments/groupPolicyConfigurationAssignmentIdValue",
			Expected: &DeviceManagementGroupPolicyConfigurationIdAssignmentId{
				GroupPolicyConfigurationId:           "groupPolicyConfigurationIdValue",
				GroupPolicyConfigurationAssignmentId: "groupPolicyConfigurationAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue/assignments/groupPolicyConfigurationAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs/gRoUpPoLiCyCoNfIgUrAtIoNiDvAlUe/aSsIgNmEnTs/gRoUpPoLiCyCoNfIgUrAtIoNaSsIgNmEnTiDvAlUe",
			Expected: &DeviceManagementGroupPolicyConfigurationIdAssignmentId{
				GroupPolicyConfigurationId:           "gRoUpPoLiCyCoNfIgUrAtIoNiDvAlUe",
				GroupPolicyConfigurationAssignmentId: "gRoUpPoLiCyCoNfIgUrAtIoNaSsIgNmEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs/gRoUpPoLiCyCoNfIgUrAtIoNiDvAlUe/aSsIgNmEnTs/gRoUpPoLiCyCoNfIgUrAtIoNaSsIgNmEnTiDvAlUe/extra",
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
