package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceConfigurationIdUserStatusId{}

func TestNewDeviceManagementDeviceConfigurationIdUserStatusID(t *testing.T) {
	id := NewDeviceManagementDeviceConfigurationIdUserStatusID("deviceConfigurationId", "deviceConfigurationUserStatusId")

	if id.DeviceConfigurationId != "deviceConfigurationId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceConfigurationId'", id.DeviceConfigurationId, "deviceConfigurationId")
	}

	if id.DeviceConfigurationUserStatusId != "deviceConfigurationUserStatusId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceConfigurationUserStatusId'", id.DeviceConfigurationUserStatusId, "deviceConfigurationUserStatusId")
	}
}

func TestFormatDeviceManagementDeviceConfigurationIdUserStatusID(t *testing.T) {
	actual := NewDeviceManagementDeviceConfigurationIdUserStatusID("deviceConfigurationId", "deviceConfigurationUserStatusId").ID()
	expected := "/deviceManagement/deviceConfigurations/deviceConfigurationId/userStatuses/deviceConfigurationUserStatusId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceConfigurationIdUserStatusID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceConfigurationIdUserStatusId
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
			Input: "/deviceManagement/deviceConfigurations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId/userStatuses",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId/userStatuses/deviceConfigurationUserStatusId",
			Expected: &DeviceManagementDeviceConfigurationIdUserStatusId{
				DeviceConfigurationId:           "deviceConfigurationId",
				DeviceConfigurationUserStatusId: "deviceConfigurationUserStatusId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId/userStatuses/deviceConfigurationUserStatusId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceConfigurationIdUserStatusID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceConfigurationId != v.Expected.DeviceConfigurationId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationId", v.Expected.DeviceConfigurationId, actual.DeviceConfigurationId)
		}

		if actual.DeviceConfigurationUserStatusId != v.Expected.DeviceConfigurationUserStatusId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationUserStatusId", v.Expected.DeviceConfigurationUserStatusId, actual.DeviceConfigurationUserStatusId)
		}

	}
}

func TestParseDeviceManagementDeviceConfigurationIdUserStatusIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceConfigurationIdUserStatusId
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
			Input: "/deviceManagement/deviceConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId/userStatuses",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnId/uSeRsTaTuSeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId/userStatuses/deviceConfigurationUserStatusId",
			Expected: &DeviceManagementDeviceConfigurationIdUserStatusId{
				DeviceConfigurationId:           "deviceConfigurationId",
				DeviceConfigurationUserStatusId: "deviceConfigurationUserStatusId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId/userStatuses/deviceConfigurationUserStatusId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnId/uSeRsTaTuSeS/dEvIcEcOnFiGuRaTiOnUsErStAtUsId",
			Expected: &DeviceManagementDeviceConfigurationIdUserStatusId{
				DeviceConfigurationId:           "dEvIcEcOnFiGuRaTiOnId",
				DeviceConfigurationUserStatusId: "dEvIcEcOnFiGuRaTiOnUsErStAtUsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnId/uSeRsTaTuSeS/dEvIcEcOnFiGuRaTiOnUsErStAtUsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceConfigurationIdUserStatusIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceConfigurationId != v.Expected.DeviceConfigurationId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationId", v.Expected.DeviceConfigurationId, actual.DeviceConfigurationId)
		}

		if actual.DeviceConfigurationUserStatusId != v.Expected.DeviceConfigurationUserStatusId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationUserStatusId", v.Expected.DeviceConfigurationUserStatusId, actual.DeviceConfigurationUserStatusId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceConfigurationIdUserStatusId(t *testing.T) {
	segments := DeviceManagementDeviceConfigurationIdUserStatusId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceConfigurationIdUserStatusId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
