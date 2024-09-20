package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDepOnboardingSettingIdEnrollmentProfileId{}

func TestNewDeviceManagementDepOnboardingSettingIdEnrollmentProfileID(t *testing.T) {
	id := NewDeviceManagementDepOnboardingSettingIdEnrollmentProfileID("depOnboardingSettingId", "enrollmentProfileId")

	if id.DepOnboardingSettingId != "depOnboardingSettingId" {
		t.Fatalf("Expected %q but got %q for Segment 'DepOnboardingSettingId'", id.DepOnboardingSettingId, "depOnboardingSettingId")
	}

	if id.EnrollmentProfileId != "enrollmentProfileId" {
		t.Fatalf("Expected %q but got %q for Segment 'EnrollmentProfileId'", id.EnrollmentProfileId, "enrollmentProfileId")
	}
}

func TestFormatDeviceManagementDepOnboardingSettingIdEnrollmentProfileID(t *testing.T) {
	actual := NewDeviceManagementDepOnboardingSettingIdEnrollmentProfileID("depOnboardingSettingId", "enrollmentProfileId").ID()
	expected := "/deviceManagement/depOnboardingSettings/depOnboardingSettingId/enrollmentProfiles/enrollmentProfileId"
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
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingId/enrollmentProfiles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingId/enrollmentProfiles/enrollmentProfileId",
			Expected: &DeviceManagementDepOnboardingSettingIdEnrollmentProfileId{
				DepOnboardingSettingId: "depOnboardingSettingId",
				EnrollmentProfileId:    "enrollmentProfileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingId/enrollmentProfiles/enrollmentProfileId/extra",
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
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs/dEpOnBoArDiNgSeTtInGiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingId/enrollmentProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs/dEpOnBoArDiNgSeTtInGiD/eNrOlLmEnTpRoFiLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingId/enrollmentProfiles/enrollmentProfileId",
			Expected: &DeviceManagementDepOnboardingSettingIdEnrollmentProfileId{
				DepOnboardingSettingId: "depOnboardingSettingId",
				EnrollmentProfileId:    "enrollmentProfileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingId/enrollmentProfiles/enrollmentProfileId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs/dEpOnBoArDiNgSeTtInGiD/eNrOlLmEnTpRoFiLeS/eNrOlLmEnTpRoFiLeId",
			Expected: &DeviceManagementDepOnboardingSettingIdEnrollmentProfileId{
				DepOnboardingSettingId: "dEpOnBoArDiNgSeTtInGiD",
				EnrollmentProfileId:    "eNrOlLmEnTpRoFiLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs/dEpOnBoArDiNgSeTtInGiD/eNrOlLmEnTpRoFiLeS/eNrOlLmEnTpRoFiLeId/extra",
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
