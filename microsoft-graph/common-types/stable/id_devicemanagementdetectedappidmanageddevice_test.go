package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDetectedAppIdManagedDeviceId{}

func TestNewDeviceManagementDetectedAppIdManagedDeviceID(t *testing.T) {
	id := NewDeviceManagementDetectedAppIdManagedDeviceID("detectedAppId", "managedDeviceId")

	if id.DetectedAppId != "detectedAppId" {
		t.Fatalf("Expected %q but got %q for Segment 'DetectedAppId'", id.DetectedAppId, "detectedAppId")
	}

	if id.ManagedDeviceId != "managedDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceId")
	}
}

func TestFormatDeviceManagementDetectedAppIdManagedDeviceID(t *testing.T) {
	actual := NewDeviceManagementDetectedAppIdManagedDeviceID("detectedAppId", "managedDeviceId").ID()
	expected := "/deviceManagement/detectedApps/detectedAppId/managedDevices/managedDeviceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDetectedAppIdManagedDeviceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDetectedAppIdManagedDeviceId
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
			Input: "/deviceManagement/detectedApps",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/detectedApps/detectedAppId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/detectedApps/detectedAppId/managedDevices",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/detectedApps/detectedAppId/managedDevices/managedDeviceId",
			Expected: &DeviceManagementDetectedAppIdManagedDeviceId{
				DetectedAppId:   "detectedAppId",
				ManagedDeviceId: "managedDeviceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/detectedApps/detectedAppId/managedDevices/managedDeviceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDetectedAppIdManagedDeviceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DetectedAppId != v.Expected.DetectedAppId {
			t.Fatalf("Expected %q but got %q for DetectedAppId", v.Expected.DetectedAppId, actual.DetectedAppId)
		}

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

	}
}

func TestParseDeviceManagementDetectedAppIdManagedDeviceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDetectedAppIdManagedDeviceId
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
			Input: "/deviceManagement/detectedApps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEtEcTeDaPpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/detectedApps/detectedAppId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEtEcTeDaPpS/dEtEcTeDaPpId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/detectedApps/detectedAppId/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEtEcTeDaPpS/dEtEcTeDaPpId/mAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/detectedApps/detectedAppId/managedDevices/managedDeviceId",
			Expected: &DeviceManagementDetectedAppIdManagedDeviceId{
				DetectedAppId:   "detectedAppId",
				ManagedDeviceId: "managedDeviceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/detectedApps/detectedAppId/managedDevices/managedDeviceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEtEcTeDaPpS/dEtEcTeDaPpId/mAnAgEdDeViCeS/mAnAgEdDeViCeId",
			Expected: &DeviceManagementDetectedAppIdManagedDeviceId{
				DetectedAppId:   "dEtEcTeDaPpId",
				ManagedDeviceId: "mAnAgEdDeViCeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEtEcTeDaPpS/dEtEcTeDaPpId/mAnAgEdDeViCeS/mAnAgEdDeViCeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDetectedAppIdManagedDeviceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DetectedAppId != v.Expected.DetectedAppId {
			t.Fatalf("Expected %q but got %q for DetectedAppId", v.Expected.DetectedAppId, actual.DetectedAppId)
		}

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

	}
}

func TestSegmentsForDeviceManagementDetectedAppIdManagedDeviceId(t *testing.T) {
	segments := DeviceManagementDetectedAppIdManagedDeviceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDetectedAppIdManagedDeviceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
