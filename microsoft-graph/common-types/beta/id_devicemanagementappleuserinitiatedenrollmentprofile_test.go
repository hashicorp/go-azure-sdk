package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAppleUserInitiatedEnrollmentProfileId{}

func TestNewDeviceManagementAppleUserInitiatedEnrollmentProfileID(t *testing.T) {
	id := NewDeviceManagementAppleUserInitiatedEnrollmentProfileID("appleUserInitiatedEnrollmentProfileId")

	if id.AppleUserInitiatedEnrollmentProfileId != "appleUserInitiatedEnrollmentProfileId" {
		t.Fatalf("Expected %q but got %q for Segment 'AppleUserInitiatedEnrollmentProfileId'", id.AppleUserInitiatedEnrollmentProfileId, "appleUserInitiatedEnrollmentProfileId")
	}
}

func TestFormatDeviceManagementAppleUserInitiatedEnrollmentProfileID(t *testing.T) {
	actual := NewDeviceManagementAppleUserInitiatedEnrollmentProfileID("appleUserInitiatedEnrollmentProfileId").ID()
	expected := "/deviceManagement/appleUserInitiatedEnrollmentProfiles/appleUserInitiatedEnrollmentProfileId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementAppleUserInitiatedEnrollmentProfileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementAppleUserInitiatedEnrollmentProfileId
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
			Input: "/deviceManagement/appleUserInitiatedEnrollmentProfiles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/appleUserInitiatedEnrollmentProfiles/appleUserInitiatedEnrollmentProfileId",
			Expected: &DeviceManagementAppleUserInitiatedEnrollmentProfileId{
				AppleUserInitiatedEnrollmentProfileId: "appleUserInitiatedEnrollmentProfileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/appleUserInitiatedEnrollmentProfiles/appleUserInitiatedEnrollmentProfileId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementAppleUserInitiatedEnrollmentProfileID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AppleUserInitiatedEnrollmentProfileId != v.Expected.AppleUserInitiatedEnrollmentProfileId {
			t.Fatalf("Expected %q but got %q for AppleUserInitiatedEnrollmentProfileId", v.Expected.AppleUserInitiatedEnrollmentProfileId, actual.AppleUserInitiatedEnrollmentProfileId)
		}

	}
}

func TestParseDeviceManagementAppleUserInitiatedEnrollmentProfileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementAppleUserInitiatedEnrollmentProfileId
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
			Input: "/deviceManagement/appleUserInitiatedEnrollmentProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aPpLeUsErInItIaTeDeNrOlLmEnTpRoFiLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/appleUserInitiatedEnrollmentProfiles/appleUserInitiatedEnrollmentProfileId",
			Expected: &DeviceManagementAppleUserInitiatedEnrollmentProfileId{
				AppleUserInitiatedEnrollmentProfileId: "appleUserInitiatedEnrollmentProfileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/appleUserInitiatedEnrollmentProfiles/appleUserInitiatedEnrollmentProfileId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aPpLeUsErInItIaTeDeNrOlLmEnTpRoFiLeS/aPpLeUsErInItIaTeDeNrOlLmEnTpRoFiLeId",
			Expected: &DeviceManagementAppleUserInitiatedEnrollmentProfileId{
				AppleUserInitiatedEnrollmentProfileId: "aPpLeUsErInItIaTeDeNrOlLmEnTpRoFiLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aPpLeUsErInItIaTeDeNrOlLmEnTpRoFiLeS/aPpLeUsErInItIaTeDeNrOlLmEnTpRoFiLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementAppleUserInitiatedEnrollmentProfileIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AppleUserInitiatedEnrollmentProfileId != v.Expected.AppleUserInitiatedEnrollmentProfileId {
			t.Fatalf("Expected %q but got %q for AppleUserInitiatedEnrollmentProfileId", v.Expected.AppleUserInitiatedEnrollmentProfileId, actual.AppleUserInitiatedEnrollmentProfileId)
		}

	}
}

func TestSegmentsForDeviceManagementAppleUserInitiatedEnrollmentProfileId(t *testing.T) {
	segments := DeviceManagementAppleUserInitiatedEnrollmentProfileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementAppleUserInitiatedEnrollmentProfileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
