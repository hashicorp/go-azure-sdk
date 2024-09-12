package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDepOnboardingSettingIdEnrollmentProfileId{}

func TestNewDeviceManagementDepOnboardingSettingIdEnrollmentProfileID(t *testing.T) {
	id := NewDeviceManagementDepOnboardingSettingIdEnrollmentProfileID("depOnboardingSettingIdValue", "enrollmentProfileIdValue")

	if id.DepOnboardingSettingId != "depOnboardingSettingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DepOnboardingSettingId'", id.DepOnboardingSettingId, "depOnboardingSettingIdValue")
	}

	if id.EnrollmentProfileId != "enrollmentProfileIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EnrollmentProfileId'", id.EnrollmentProfileId, "enrollmentProfileIdValue")
	}
}

func TestFormatDeviceManagementDepOnboardingSettingIdEnrollmentProfileID(t *testing.T) {
	actual := NewDeviceManagementDepOnboardingSettingIdEnrollmentProfileID("depOnboardingSettingIdValue", "enrollmentProfileIdValue").ID()
	expected := "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue/enrollmentProfiles/enrollmentProfileIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDepOnboardingSettingIdEnrollmentProfileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDepOnboardingSettingIdEnrollmentProfileId
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
			Input: "/deviceManagement/depOnboardingSettings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue/enrollmentProfiles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue/enrollmentProfiles/enrollmentProfileIdValue",
			Expected: &DeviceManagementDepOnboardingSettingIdEnrollmentProfileId{
				DepOnboardingSettingId: "depOnboardingSettingIdValue",
				EnrollmentProfileId:    "enrollmentProfileIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue/enrollmentProfiles/enrollmentProfileIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDepOnboardingSettingIdEnrollmentProfileID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DepOnboardingSettingId != v.Expected.DepOnboardingSettingId {
			t.Fatalf("Expected %q but got %q for DepOnboardingSettingId", v.Expected.DepOnboardingSettingId, actual.DepOnboardingSettingId)
		}

		if actual.EnrollmentProfileId != v.Expected.EnrollmentProfileId {
			t.Fatalf("Expected %q but got %q for EnrollmentProfileId", v.Expected.EnrollmentProfileId, actual.EnrollmentProfileId)
		}

	}
}

func TestParseDeviceManagementDepOnboardingSettingIdEnrollmentProfileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDepOnboardingSettingIdEnrollmentProfileId
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
			Input: "/deviceManagement/depOnboardingSettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs/dEpOnBoArDiNgSeTtInGiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue/enrollmentProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs/dEpOnBoArDiNgSeTtInGiDvAlUe/eNrOlLmEnTpRoFiLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue/enrollmentProfiles/enrollmentProfileIdValue",
			Expected: &DeviceManagementDepOnboardingSettingIdEnrollmentProfileId{
				DepOnboardingSettingId: "depOnboardingSettingIdValue",
				EnrollmentProfileId:    "enrollmentProfileIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue/enrollmentProfiles/enrollmentProfileIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs/dEpOnBoArDiNgSeTtInGiDvAlUe/eNrOlLmEnTpRoFiLeS/eNrOlLmEnTpRoFiLeIdVaLuE",
			Expected: &DeviceManagementDepOnboardingSettingIdEnrollmentProfileId{
				DepOnboardingSettingId: "dEpOnBoArDiNgSeTtInGiDvAlUe",
				EnrollmentProfileId:    "eNrOlLmEnTpRoFiLeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs/dEpOnBoArDiNgSeTtInGiDvAlUe/eNrOlLmEnTpRoFiLeS/eNrOlLmEnTpRoFiLeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDepOnboardingSettingIdEnrollmentProfileIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DepOnboardingSettingId != v.Expected.DepOnboardingSettingId {
			t.Fatalf("Expected %q but got %q for DepOnboardingSettingId", v.Expected.DepOnboardingSettingId, actual.DepOnboardingSettingId)
		}

		if actual.EnrollmentProfileId != v.Expected.EnrollmentProfileId {
			t.Fatalf("Expected %q but got %q for EnrollmentProfileId", v.Expected.EnrollmentProfileId, actual.EnrollmentProfileId)
		}

	}
}

func TestSegmentsForDeviceManagementDepOnboardingSettingIdEnrollmentProfileId(t *testing.T) {
	segments := DeviceManagementDepOnboardingSettingIdEnrollmentProfileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDepOnboardingSettingIdEnrollmentProfileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
