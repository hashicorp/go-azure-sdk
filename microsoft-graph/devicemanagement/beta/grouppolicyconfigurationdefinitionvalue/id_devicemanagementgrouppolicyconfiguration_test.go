package grouppolicyconfigurationdefinitionvalue

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyConfigurationId{}

func TestNewDeviceManagementGroupPolicyConfigurationID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyConfigurationID("groupPolicyConfigurationIdValue")

	if id.GroupPolicyConfigurationId != "groupPolicyConfigurationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyConfigurationId'", id.GroupPolicyConfigurationId, "groupPolicyConfigurationIdValue")
	}
}

func TestFormatDeviceManagementGroupPolicyConfigurationID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyConfigurationID("groupPolicyConfigurationIdValue").ID()
	expected := "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementGroupPolicyConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyConfigurationId
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
			// Valid URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue",
			Expected: &DeviceManagementGroupPolicyConfigurationId{
				GroupPolicyConfigurationId: "groupPolicyConfigurationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyConfigurationID(v.Input)
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

	}
}

func TestParseDeviceManagementGroupPolicyConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyConfigurationId
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
			// Valid URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue",
			Expected: &DeviceManagementGroupPolicyConfigurationId{
				GroupPolicyConfigurationId: "groupPolicyConfigurationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs/gRoUpPoLiCyCoNfIgUrAtIoNiDvAlUe",
			Expected: &DeviceManagementGroupPolicyConfigurationId{
				GroupPolicyConfigurationId: "gRoUpPoLiCyCoNfIgUrAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs/gRoUpPoLiCyCoNfIgUrAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyConfigurationIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementGroupPolicyConfigurationId(t *testing.T) {
	segments := DeviceManagementGroupPolicyConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementGroupPolicyConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
