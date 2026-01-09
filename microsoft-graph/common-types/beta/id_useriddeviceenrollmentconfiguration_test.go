package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDeviceEnrollmentConfigurationId{}

func TestNewUserIdDeviceEnrollmentConfigurationID(t *testing.T) {
	id := NewUserIdDeviceEnrollmentConfigurationID("userId", "deviceEnrollmentConfigurationId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.DeviceEnrollmentConfigurationId != "deviceEnrollmentConfigurationId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceEnrollmentConfigurationId'", id.DeviceEnrollmentConfigurationId, "deviceEnrollmentConfigurationId")
	}
}

func TestFormatUserIdDeviceEnrollmentConfigurationID(t *testing.T) {
	actual := NewUserIdDeviceEnrollmentConfigurationID("userId", "deviceEnrollmentConfigurationId").ID()
	expected := "/users/userId/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdDeviceEnrollmentConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDeviceEnrollmentConfigurationId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/deviceEnrollmentConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId",
			Expected: &UserIdDeviceEnrollmentConfigurationId{
				UserId:                          "userId",
				DeviceEnrollmentConfigurationId: "deviceEnrollmentConfigurationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDeviceEnrollmentConfigurationID(v.Input)
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

func TestParseUserIdDeviceEnrollmentConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDeviceEnrollmentConfigurationId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/deviceEnrollmentConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId",
			Expected: &UserIdDeviceEnrollmentConfigurationId{
				UserId:                          "userId",
				DeviceEnrollmentConfigurationId: "deviceEnrollmentConfigurationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnId",
			Expected: &UserIdDeviceEnrollmentConfigurationId{
				UserId:                          "uSeRiD",
				DeviceEnrollmentConfigurationId: "dEvIcEeNrOlLmEnTcOnFiGuRaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDeviceEnrollmentConfigurationIDInsensitively(v.Input)
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

func TestSegmentsForUserIdDeviceEnrollmentConfigurationId(t *testing.T) {
	segments := UserIdDeviceEnrollmentConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdDeviceEnrollmentConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
