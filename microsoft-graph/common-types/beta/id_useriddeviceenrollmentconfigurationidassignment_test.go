package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDeviceEnrollmentConfigurationIdAssignmentId{}

func TestNewUserIdDeviceEnrollmentConfigurationIdAssignmentID(t *testing.T) {
	id := NewUserIdDeviceEnrollmentConfigurationIdAssignmentID("userId", "deviceEnrollmentConfigurationId", "enrollmentConfigurationAssignmentId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.DeviceEnrollmentConfigurationId != "deviceEnrollmentConfigurationId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceEnrollmentConfigurationId'", id.DeviceEnrollmentConfigurationId, "deviceEnrollmentConfigurationId")
	}

	if id.EnrollmentConfigurationAssignmentId != "enrollmentConfigurationAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'EnrollmentConfigurationAssignmentId'", id.EnrollmentConfigurationAssignmentId, "enrollmentConfigurationAssignmentId")
	}
}

func TestFormatUserIdDeviceEnrollmentConfigurationIdAssignmentID(t *testing.T) {
	actual := NewUserIdDeviceEnrollmentConfigurationIdAssignmentID("userId", "deviceEnrollmentConfigurationId", "enrollmentConfigurationAssignmentId").ID()
	expected := "/users/userId/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId/assignments/enrollmentConfigurationAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdDeviceEnrollmentConfigurationIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDeviceEnrollmentConfigurationIdAssignmentId
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
			// Incomplete URI
			Input: "/users/userId/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId/assignments/enrollmentConfigurationAssignmentId",
			Expected: &UserIdDeviceEnrollmentConfigurationIdAssignmentId{
				UserId:                              "userId",
				DeviceEnrollmentConfigurationId:     "deviceEnrollmentConfigurationId",
				EnrollmentConfigurationAssignmentId: "enrollmentConfigurationAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId/assignments/enrollmentConfigurationAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDeviceEnrollmentConfigurationIdAssignmentID(v.Input)
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

		if actual.EnrollmentConfigurationAssignmentId != v.Expected.EnrollmentConfigurationAssignmentId {
			t.Fatalf("Expected %q but got %q for EnrollmentConfigurationAssignmentId", v.Expected.EnrollmentConfigurationAssignmentId, actual.EnrollmentConfigurationAssignmentId)
		}

	}
}

func TestParseUserIdDeviceEnrollmentConfigurationIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDeviceEnrollmentConfigurationIdAssignmentId
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
			// Incomplete URI
			Input: "/users/userId/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnId/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId/assignments/enrollmentConfigurationAssignmentId",
			Expected: &UserIdDeviceEnrollmentConfigurationIdAssignmentId{
				UserId:                              "userId",
				DeviceEnrollmentConfigurationId:     "deviceEnrollmentConfigurationId",
				EnrollmentConfigurationAssignmentId: "enrollmentConfigurationAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId/assignments/enrollmentConfigurationAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnId/aSsIgNmEnTs/eNrOlLmEnTcOnFiGuRaTiOnAsSiGnMeNtId",
			Expected: &UserIdDeviceEnrollmentConfigurationIdAssignmentId{
				UserId:                              "uSeRiD",
				DeviceEnrollmentConfigurationId:     "dEvIcEeNrOlLmEnTcOnFiGuRaTiOnId",
				EnrollmentConfigurationAssignmentId: "eNrOlLmEnTcOnFiGuRaTiOnAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnId/aSsIgNmEnTs/eNrOlLmEnTcOnFiGuRaTiOnAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDeviceEnrollmentConfigurationIdAssignmentIDInsensitively(v.Input)
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

		if actual.EnrollmentConfigurationAssignmentId != v.Expected.EnrollmentConfigurationAssignmentId {
			t.Fatalf("Expected %q but got %q for EnrollmentConfigurationAssignmentId", v.Expected.EnrollmentConfigurationAssignmentId, actual.EnrollmentConfigurationAssignmentId)
		}

	}
}

func TestSegmentsForUserIdDeviceEnrollmentConfigurationIdAssignmentId(t *testing.T) {
	segments := UserIdDeviceEnrollmentConfigurationIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdDeviceEnrollmentConfigurationIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
