package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDeviceEnrollmentConfigurationIdAssignmentId{}

func TestNewMeDeviceEnrollmentConfigurationIdAssignmentID(t *testing.T) {
	id := NewMeDeviceEnrollmentConfigurationIdAssignmentID("deviceEnrollmentConfigurationId", "enrollmentConfigurationAssignmentId")

	if id.DeviceEnrollmentConfigurationId != "deviceEnrollmentConfigurationId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceEnrollmentConfigurationId'", id.DeviceEnrollmentConfigurationId, "deviceEnrollmentConfigurationId")
	}

	if id.EnrollmentConfigurationAssignmentId != "enrollmentConfigurationAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'EnrollmentConfigurationAssignmentId'", id.EnrollmentConfigurationAssignmentId, "enrollmentConfigurationAssignmentId")
	}
}

func TestFormatMeDeviceEnrollmentConfigurationIdAssignmentID(t *testing.T) {
	actual := NewMeDeviceEnrollmentConfigurationIdAssignmentID("deviceEnrollmentConfigurationId", "enrollmentConfigurationAssignmentId").ID()
	expected := "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId/assignments/enrollmentConfigurationAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDeviceEnrollmentConfigurationIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDeviceEnrollmentConfigurationIdAssignmentId
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
			// Incomplete URI
			Input: "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId/assignments/enrollmentConfigurationAssignmentId",
			Expected: &MeDeviceEnrollmentConfigurationIdAssignmentId{
				DeviceEnrollmentConfigurationId:     "deviceEnrollmentConfigurationId",
				EnrollmentConfigurationAssignmentId: "enrollmentConfigurationAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId/assignments/enrollmentConfigurationAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDeviceEnrollmentConfigurationIdAssignmentID(v.Input)
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

		if actual.EnrollmentConfigurationAssignmentId != v.Expected.EnrollmentConfigurationAssignmentId {
			t.Fatalf("Expected %q but got %q for EnrollmentConfigurationAssignmentId", v.Expected.EnrollmentConfigurationAssignmentId, actual.EnrollmentConfigurationAssignmentId)
		}

	}
}

func TestParseMeDeviceEnrollmentConfigurationIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDeviceEnrollmentConfigurationIdAssignmentId
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
			// Incomplete URI
			Input: "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnId/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId/assignments/enrollmentConfigurationAssignmentId",
			Expected: &MeDeviceEnrollmentConfigurationIdAssignmentId{
				DeviceEnrollmentConfigurationId:     "deviceEnrollmentConfigurationId",
				EnrollmentConfigurationAssignmentId: "enrollmentConfigurationAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationId/assignments/enrollmentConfigurationAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnId/aSsIgNmEnTs/eNrOlLmEnTcOnFiGuRaTiOnAsSiGnMeNtId",
			Expected: &MeDeviceEnrollmentConfigurationIdAssignmentId{
				DeviceEnrollmentConfigurationId:     "dEvIcEeNrOlLmEnTcOnFiGuRaTiOnId",
				EnrollmentConfigurationAssignmentId: "eNrOlLmEnTcOnFiGuRaTiOnAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnId/aSsIgNmEnTs/eNrOlLmEnTcOnFiGuRaTiOnAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDeviceEnrollmentConfigurationIdAssignmentIDInsensitively(v.Input)
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

		if actual.EnrollmentConfigurationAssignmentId != v.Expected.EnrollmentConfigurationAssignmentId {
			t.Fatalf("Expected %q but got %q for EnrollmentConfigurationAssignmentId", v.Expected.EnrollmentConfigurationAssignmentId, actual.EnrollmentConfigurationAssignmentId)
		}

	}
}

func TestSegmentsForMeDeviceEnrollmentConfigurationIdAssignmentId(t *testing.T) {
	segments := MeDeviceEnrollmentConfigurationIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDeviceEnrollmentConfigurationIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
