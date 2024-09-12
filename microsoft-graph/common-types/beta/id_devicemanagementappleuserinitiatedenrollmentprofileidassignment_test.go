package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId{}

func TestNewDeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentID("appleUserInitiatedEnrollmentProfileIdValue", "appleEnrollmentProfileAssignmentIdValue")

	if id.AppleUserInitiatedEnrollmentProfileId != "appleUserInitiatedEnrollmentProfileIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AppleUserInitiatedEnrollmentProfileId'", id.AppleUserInitiatedEnrollmentProfileId, "appleUserInitiatedEnrollmentProfileIdValue")
	}

	if id.AppleEnrollmentProfileAssignmentId != "appleEnrollmentProfileAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AppleEnrollmentProfileAssignmentId'", id.AppleEnrollmentProfileAssignmentId, "appleEnrollmentProfileAssignmentIdValue")
	}
}

func TestFormatDeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentID("appleUserInitiatedEnrollmentProfileIdValue", "appleEnrollmentProfileAssignmentIdValue").ID()
	expected := "/deviceManagement/appleUserInitiatedEnrollmentProfiles/appleUserInitiatedEnrollmentProfileIdValue/assignments/appleEnrollmentProfileAssignmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId
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
			// Incomplete URI
			Input: "/deviceManagement/appleUserInitiatedEnrollmentProfiles/appleUserInitiatedEnrollmentProfileIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/appleUserInitiatedEnrollmentProfiles/appleUserInitiatedEnrollmentProfileIdValue/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/appleUserInitiatedEnrollmentProfiles/appleUserInitiatedEnrollmentProfileIdValue/assignments/appleEnrollmentProfileAssignmentIdValue",
			Expected: &DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId{
				AppleUserInitiatedEnrollmentProfileId: "appleUserInitiatedEnrollmentProfileIdValue",
				AppleEnrollmentProfileAssignmentId:    "appleEnrollmentProfileAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/appleUserInitiatedEnrollmentProfiles/appleUserInitiatedEnrollmentProfileIdValue/assignments/appleEnrollmentProfileAssignmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentID(v.Input)
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

		if actual.AppleEnrollmentProfileAssignmentId != v.Expected.AppleEnrollmentProfileAssignmentId {
			t.Fatalf("Expected %q but got %q for AppleEnrollmentProfileAssignmentId", v.Expected.AppleEnrollmentProfileAssignmentId, actual.AppleEnrollmentProfileAssignmentId)
		}

	}
}

func TestParseDeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId
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
			// Incomplete URI
			Input: "/deviceManagement/appleUserInitiatedEnrollmentProfiles/appleUserInitiatedEnrollmentProfileIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aPpLeUsErInItIaTeDeNrOlLmEnTpRoFiLeS/aPpLeUsErInItIaTeDeNrOlLmEnTpRoFiLeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/appleUserInitiatedEnrollmentProfiles/appleUserInitiatedEnrollmentProfileIdValue/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aPpLeUsErInItIaTeDeNrOlLmEnTpRoFiLeS/aPpLeUsErInItIaTeDeNrOlLmEnTpRoFiLeIdVaLuE/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/appleUserInitiatedEnrollmentProfiles/appleUserInitiatedEnrollmentProfileIdValue/assignments/appleEnrollmentProfileAssignmentIdValue",
			Expected: &DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId{
				AppleUserInitiatedEnrollmentProfileId: "appleUserInitiatedEnrollmentProfileIdValue",
				AppleEnrollmentProfileAssignmentId:    "appleEnrollmentProfileAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/appleUserInitiatedEnrollmentProfiles/appleUserInitiatedEnrollmentProfileIdValue/assignments/appleEnrollmentProfileAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aPpLeUsErInItIaTeDeNrOlLmEnTpRoFiLeS/aPpLeUsErInItIaTeDeNrOlLmEnTpRoFiLeIdVaLuE/aSsIgNmEnTs/aPpLeEnRoLlMeNtPrOfIlEaSsIgNmEnTiDvAlUe",
			Expected: &DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId{
				AppleUserInitiatedEnrollmentProfileId: "aPpLeUsErInItIaTeDeNrOlLmEnTpRoFiLeIdVaLuE",
				AppleEnrollmentProfileAssignmentId:    "aPpLeEnRoLlMeNtPrOfIlEaSsIgNmEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aPpLeUsErInItIaTeDeNrOlLmEnTpRoFiLeS/aPpLeUsErInItIaTeDeNrOlLmEnTpRoFiLeIdVaLuE/aSsIgNmEnTs/aPpLeEnRoLlMeNtPrOfIlEaSsIgNmEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentIDInsensitively(v.Input)
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

		if actual.AppleEnrollmentProfileAssignmentId != v.Expected.AppleEnrollmentProfileAssignmentId {
			t.Fatalf("Expected %q but got %q for AppleEnrollmentProfileAssignmentId", v.Expected.AppleEnrollmentProfileAssignmentId, actual.AppleEnrollmentProfileAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId(t *testing.T) {
	segments := DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
