package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceConfigurationIdUserStatusId{}

func TestNewDeviceManagementDeviceConfigurationIdUserStatusID(t *testing.T) {
	id := NewDeviceManagementDeviceConfigurationIdUserStatusID("deviceConfigurationIdValue", "deviceConfigurationUserStatusIdValue")

	if id.DeviceConfigurationId != "deviceConfigurationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceConfigurationId'", id.DeviceConfigurationId, "deviceConfigurationIdValue")
	}

	if id.DeviceConfigurationUserStatusId != "deviceConfigurationUserStatusIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceConfigurationUserStatusId'", id.DeviceConfigurationUserStatusId, "deviceConfigurationUserStatusIdValue")
	}
}

func TestFormatDeviceManagementDeviceConfigurationIdUserStatusID(t *testing.T) {
	actual := NewDeviceManagementDeviceConfigurationIdUserStatusID("deviceConfigurationIdValue", "deviceConfigurationUserStatusIdValue").ID()
	expected := "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/userStatuses/deviceConfigurationUserStatusIdValue"
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
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/userStatuses",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/userStatuses/deviceConfigurationUserStatusIdValue",
			Expected: &DeviceManagementDeviceConfigurationIdUserStatusId{
				DeviceConfigurationId:           "deviceConfigurationIdValue",
				DeviceConfigurationUserStatusId: "deviceConfigurationUserStatusIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/userStatuses/deviceConfigurationUserStatusIdValue/extra",
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
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/userStatuses",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnIdVaLuE/uSeRsTaTuSeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/userStatuses/deviceConfigurationUserStatusIdValue",
			Expected: &DeviceManagementDeviceConfigurationIdUserStatusId{
				DeviceConfigurationId:           "deviceConfigurationIdValue",
				DeviceConfigurationUserStatusId: "deviceConfigurationUserStatusIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/userStatuses/deviceConfigurationUserStatusIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnIdVaLuE/uSeRsTaTuSeS/dEvIcEcOnFiGuRaTiOnUsErStAtUsIdVaLuE",
			Expected: &DeviceManagementDeviceConfigurationIdUserStatusId{
				DeviceConfigurationId:           "dEvIcEcOnFiGuRaTiOnIdVaLuE",
				DeviceConfigurationUserStatusId: "dEvIcEcOnFiGuRaTiOnUsErStAtUsIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnIdVaLuE/uSeRsTaTuSeS/dEvIcEcOnFiGuRaTiOnUsErStAtUsIdVaLuE/extra",
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
