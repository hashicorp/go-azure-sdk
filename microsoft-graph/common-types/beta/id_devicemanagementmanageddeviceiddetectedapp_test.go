package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceIdDetectedAppId{}

func TestNewDeviceManagementManagedDeviceIdDetectedAppID(t *testing.T) {
	id := NewDeviceManagementManagedDeviceIdDetectedAppID("managedDeviceId", "detectedAppId")

	if id.ManagedDeviceId != "managedDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceId")
	}

	if id.DetectedAppId != "detectedAppId" {
		t.Fatalf("Expected %q but got %q for Segment 'DetectedAppId'", id.DetectedAppId, "detectedAppId")
	}
}

func TestFormatDeviceManagementManagedDeviceIdDetectedAppID(t *testing.T) {
	actual := NewDeviceManagementManagedDeviceIdDetectedAppID("managedDeviceId", "detectedAppId").ID()
	expected := "/deviceManagement/managedDevices/managedDeviceId/detectedApps/detectedAppId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementManagedDeviceIdDetectedAppID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementManagedDeviceIdDetectedAppId
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
			Input: "/deviceManagement/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/managedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/managedDevices/managedDeviceId/detectedApps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/managedDevices/managedDeviceId/detectedApps/detectedAppId",
			Expected: &DeviceManagementManagedDeviceIdDetectedAppId{
				ManagedDeviceId: "managedDeviceId",
				DetectedAppId:   "detectedAppId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/managedDevices/managedDeviceId/detectedApps/detectedAppId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementManagedDeviceIdDetectedAppID(v.Input)
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

func TestParseDeviceManagementManagedDeviceIdDetectedAppIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementManagedDeviceIdDetectedAppId
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
			Input: "/deviceManagement/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/managedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/managedDevices/managedDeviceId/detectedApps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeId/dEtEcTeDaPpS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/managedDevices/managedDeviceId/detectedApps/detectedAppId",
			Expected: &DeviceManagementManagedDeviceIdDetectedAppId{
				ManagedDeviceId: "managedDeviceId",
				DetectedAppId:   "detectedAppId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/managedDevices/managedDeviceId/detectedApps/detectedAppId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeId/dEtEcTeDaPpS/dEtEcTeDaPpId",
			Expected: &DeviceManagementManagedDeviceIdDetectedAppId{
				ManagedDeviceId: "mAnAgEdDeViCeId",
				DetectedAppId:   "dEtEcTeDaPpId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeId/dEtEcTeDaPpS/dEtEcTeDaPpId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementManagedDeviceIdDetectedAppIDInsensitively(v.Input)
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

func TestSegmentsForDeviceManagementManagedDeviceIdDetectedAppId(t *testing.T) {
	segments := DeviceManagementManagedDeviceIdDetectedAppId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementManagedDeviceIdDetectedAppId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
