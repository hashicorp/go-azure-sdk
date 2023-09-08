package medeviceenrollmentconfigurationassignment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeDeviceEnrollmentConfigurationAssignmentId{}

func TestNewMeDeviceEnrollmentConfigurationAssignmentID(t *testing.T) {
	id := NewMeDeviceEnrollmentConfigurationAssignmentID("deviceEnrollmentConfigurationIdValue", "enrollmentConfigurationAssignmentIdValue")

	if id.DeviceEnrollmentConfigurationId != "deviceEnrollmentConfigurationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceEnrollmentConfigurationId'", id.DeviceEnrollmentConfigurationId, "deviceEnrollmentConfigurationIdValue")
	}

	if id.EnrollmentConfigurationAssignmentId != "enrollmentConfigurationAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EnrollmentConfigurationAssignmentId'", id.EnrollmentConfigurationAssignmentId, "enrollmentConfigurationAssignmentIdValue")
	}
}

func TestFormatMeDeviceEnrollmentConfigurationAssignmentID(t *testing.T) {
	actual := NewMeDeviceEnrollmentConfigurationAssignmentID("deviceEnrollmentConfigurationIdValue", "enrollmentConfigurationAssignmentIdValue").ID()
	expected := "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationIdValue/assignments/enrollmentConfigurationAssignmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDeviceEnrollmentConfigurationAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDeviceEnrollmentConfigurationAssignmentId
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
			Input: "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationIdValue/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationIdValue/assignments/enrollmentConfigurationAssignmentIdValue",
			Expected: &MeDeviceEnrollmentConfigurationAssignmentId{
				DeviceEnrollmentConfigurationId:     "deviceEnrollmentConfigurationIdValue",
				EnrollmentConfigurationAssignmentId: "enrollmentConfigurationAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationIdValue/assignments/enrollmentConfigurationAssignmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDeviceEnrollmentConfigurationAssignmentID(v.Input)
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

func TestParseMeDeviceEnrollmentConfigurationAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDeviceEnrollmentConfigurationAssignmentId
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
			Input: "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationIdValue/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnIdVaLuE/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationIdValue/assignments/enrollmentConfigurationAssignmentIdValue",
			Expected: &MeDeviceEnrollmentConfigurationAssignmentId{
				DeviceEnrollmentConfigurationId:     "deviceEnrollmentConfigurationIdValue",
				EnrollmentConfigurationAssignmentId: "enrollmentConfigurationAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/deviceEnrollmentConfigurations/deviceEnrollmentConfigurationIdValue/assignments/enrollmentConfigurationAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnIdVaLuE/aSsIgNmEnTs/eNrOlLmEnTcOnFiGuRaTiOnAsSiGnMeNtIdVaLuE",
			Expected: &MeDeviceEnrollmentConfigurationAssignmentId{
				DeviceEnrollmentConfigurationId:     "dEvIcEeNrOlLmEnTcOnFiGuRaTiOnIdVaLuE",
				EnrollmentConfigurationAssignmentId: "eNrOlLmEnTcOnFiGuRaTiOnAsSiGnMeNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnS/dEvIcEeNrOlLmEnTcOnFiGuRaTiOnIdVaLuE/aSsIgNmEnTs/eNrOlLmEnTcOnFiGuRaTiOnAsSiGnMeNtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDeviceEnrollmentConfigurationAssignmentIDInsensitively(v.Input)
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

func TestSegmentsForMeDeviceEnrollmentConfigurationAssignmentId(t *testing.T) {
	segments := MeDeviceEnrollmentConfigurationAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDeviceEnrollmentConfigurationAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
