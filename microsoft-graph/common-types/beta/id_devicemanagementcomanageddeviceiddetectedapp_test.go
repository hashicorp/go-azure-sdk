package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComanagedDeviceIdDetectedAppId{}

func TestNewDeviceManagementComanagedDeviceIdDetectedAppID(t *testing.T) {
	id := NewDeviceManagementComanagedDeviceIdDetectedAppID("managedDeviceId", "detectedAppId")

	if id.ManagedDeviceId != "managedDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceId")
	}

	if id.DetectedAppId != "detectedAppId" {
		t.Fatalf("Expected %q but got %q for Segment 'DetectedAppId'", id.DetectedAppId, "detectedAppId")
	}
}

func TestFormatDeviceManagementComanagedDeviceIdDetectedAppID(t *testing.T) {
	actual := NewDeviceManagementComanagedDeviceIdDetectedAppID("managedDeviceId", "detectedAppId").ID()
	expected := "/deviceManagement/comanagedDevices/managedDeviceId/detectedApps/detectedAppId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementComanagedDeviceIdDetectedAppID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementComanagedDeviceIdDetectedAppId
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
			Input: "/deviceManagement/comanagedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/detectedApps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/detectedApps/detectedAppId",
			Expected: &DeviceManagementComanagedDeviceIdDetectedAppId{
				ManagedDeviceId: "managedDeviceId",
				DetectedAppId:   "detectedAppId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/detectedApps/detectedAppId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementComanagedDeviceIdDetectedAppID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

		if actual.DetectedAppId != v.Expected.DetectedAppId {
			t.Fatalf("Expected %q but got %q for DetectedAppId", v.Expected.DetectedAppId, actual.DetectedAppId)
		}

	}
}

func TestParseDeviceManagementComanagedDeviceIdDetectedAppIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementComanagedDeviceIdDetectedAppId
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
			Input: "/deviceManagement/comanagedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/detectedApps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeId/dEtEcTeDaPpS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/detectedApps/detectedAppId",
			Expected: &DeviceManagementComanagedDeviceIdDetectedAppId{
				ManagedDeviceId: "managedDeviceId",
				DetectedAppId:   "detectedAppId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/detectedApps/detectedAppId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeId/dEtEcTeDaPpS/dEtEcTeDaPpId",
			Expected: &DeviceManagementComanagedDeviceIdDetectedAppId{
				ManagedDeviceId: "mAnAgEdDeViCeId",
				DetectedAppId:   "dEtEcTeDaPpId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeId/dEtEcTeDaPpS/dEtEcTeDaPpId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementComanagedDeviceIdDetectedAppIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

		if actual.DetectedAppId != v.Expected.DetectedAppId {
			t.Fatalf("Expected %q but got %q for DetectedAppId", v.Expected.DetectedAppId, actual.DetectedAppId)
		}

	}
}

func TestSegmentsForDeviceManagementComanagedDeviceIdDetectedAppId(t *testing.T) {
	segments := DeviceManagementComanagedDeviceIdDetectedAppId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementComanagedDeviceIdDetectedAppId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
