package deviceenrollmentconfigurationassignment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDeviceEnrollmentConfigurationId{}

func TestNewMeDeviceEnrollmentConfigurationID(t *testing.T) {
	id := NewMeDeviceEnrollmentConfigurationID("deviceEnrollmentConfigurationIdValue")

	if id.DeviceEnrollmentConfigurationId != "deviceEnrollmentConfigurationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceEnrollmentConfigurationId'", id.DeviceEnrollmentConfigurationId, "deviceEnrollmentConfigurationIdValue")
	}
}

func TestFormatMeDeviceEnrollmentConfigurationID(t *testing.T) {
	actual := NewMeDeviceEnrollmentConfigurationID("deviceEnrollmentConfigurationIdValue").ID()
	expected := "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDeviceEnrollmentConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDeviceEnrollmentConfigurationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/deviceEnrollmentConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationIdValue",
			Expected: &MeDeviceEnrollmentConfigurationId{
				DeviceEnrollmentConfigurationId: "deviceEnrollmentConfigurationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDeviceEnrollmentConfigurationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceEnrollmentConfigurationId != v.Expected.DeviceEnrollmentConfigurationId {
			t.Fatalf("Expected %q but got %q for DeviceEnrollmentConfigurationId", v.Expected.DeviceEnrollmentConfigurationId, actual.DeviceEnrollmentConfigurationId)
		}

	}
}

func TestParseMeDeviceEnrollmentConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDeviceEnrollmentConfigurationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/deviceEnrollmentConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationIdValue",
			Expected: &MeDeviceEnrollmentConfigurationId{
				DeviceEnrollmentConfigurationId: "deviceEnrollmentConfigurationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnIdVaLuE",
			Expected: &MeDeviceEnrollmentConfigurationId{
				DeviceEnrollmentConfigurationId: "dEvIcEeNrOlLmEnTcOnFiGuRaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDeviceEnrollmentConfigurationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceEnrollmentConfigurationId != v.Expected.DeviceEnrollmentConfigurationId {
			t.Fatalf("Expected %q but got %q for DeviceEnrollmentConfigurationId", v.Expected.DeviceEnrollmentConfigurationId, actual.DeviceEnrollmentConfigurationId)
		}

	}
}

func TestSegmentsForMeDeviceEnrollmentConfigurationId(t *testing.T) {
	segments := MeDeviceEnrollmentConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDeviceEnrollmentConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
