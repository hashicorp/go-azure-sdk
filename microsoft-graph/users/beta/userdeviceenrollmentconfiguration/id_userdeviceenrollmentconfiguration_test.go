package userdeviceenrollmentconfiguration

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserDeviceEnrollmentConfigurationId{}

func TestNewUserDeviceEnrollmentConfigurationID(t *testing.T) {
	id := NewUserDeviceEnrollmentConfigurationID("userIdValue", "deviceEnrollmentConfigurationIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.DeviceEnrollmentConfigurationId != "deviceEnrollmentConfigurationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceEnrollmentConfigurationId'", id.DeviceEnrollmentConfigurationId, "deviceEnrollmentConfigurationIdValue")
	}
}

func TestFormatUserDeviceEnrollmentConfigurationID(t *testing.T) {
	actual := NewUserDeviceEnrollmentConfigurationID("userIdValue", "deviceEnrollmentConfigurationIdValue").ID()
	expected := "/users/userIdValue/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserDeviceEnrollmentConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserDeviceEnrollmentConfigurationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/deviceEnrollmentConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationIdValue",
			Expected: &UserDeviceEnrollmentConfigurationId{
				UserId:                          "userIdValue",
				DeviceEnrollmentConfigurationId: "deviceEnrollmentConfigurationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserDeviceEnrollmentConfigurationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.DeviceEnrollmentConfigurationId != v.Expected.DeviceEnrollmentConfigurationId {
			t.Fatalf("Expected %q but got %q for DeviceEnrollmentConfigurationId", v.Expected.DeviceEnrollmentConfigurationId, actual.DeviceEnrollmentConfigurationId)
		}

	}
}

func TestParseUserDeviceEnrollmentConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserDeviceEnrollmentConfigurationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/deviceEnrollmentConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationIdValue",
			Expected: &UserDeviceEnrollmentConfigurationId{
				UserId:                          "userIdValue",
				DeviceEnrollmentConfigurationId: "deviceEnrollmentConfigurationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnIdVaLuE",
			Expected: &UserDeviceEnrollmentConfigurationId{
				UserId:                          "uSeRiDvAlUe",
				DeviceEnrollmentConfigurationId: "dEvIcEeNrOlLmEnTcOnFiGuRaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserDeviceEnrollmentConfigurationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.DeviceEnrollmentConfigurationId != v.Expected.DeviceEnrollmentConfigurationId {
			t.Fatalf("Expected %q but got %q for DeviceEnrollmentConfigurationId", v.Expected.DeviceEnrollmentConfigurationId, actual.DeviceEnrollmentConfigurationId)
		}

	}
}

func TestSegmentsForUserDeviceEnrollmentConfigurationId(t *testing.T) {
	segments := UserDeviceEnrollmentConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserDeviceEnrollmentConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
