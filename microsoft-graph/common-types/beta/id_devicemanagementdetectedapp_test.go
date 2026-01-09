package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDetectedAppId{}

func TestNewDeviceManagementDetectedAppID(t *testing.T) {
	id := NewDeviceManagementDetectedAppID("detectedAppId")

	if id.DetectedAppId != "detectedAppId" {
		t.Fatalf("Expected %q but got %q for Segment 'DetectedAppId'", id.DetectedAppId, "detectedAppId")
	}
}

func TestFormatDeviceManagementDetectedAppID(t *testing.T) {
	actual := NewDeviceManagementDetectedAppID("detectedAppId").ID()
	expected := "/deviceManagement/detectedApps/detectedAppId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDetectedAppID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDetectedAppId
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
			// Valid URI
			Input: "/deviceManagement/detectedApps/detectedAppId",
			Expected: &DeviceManagementDetectedAppId{
				DetectedAppId: "detectedAppId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/detectedApps/detectedAppId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDetectedAppID(v.Input)
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

	}
}

func TestParseDeviceManagementDetectedAppIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDetectedAppId
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
			// Valid URI
			Input: "/deviceManagement/detectedApps/detectedAppId",
			Expected: &DeviceManagementDetectedAppId{
				DetectedAppId: "detectedAppId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/detectedApps/detectedAppId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEtEcTeDaPpS/dEtEcTeDaPpId",
			Expected: &DeviceManagementDetectedAppId{
				DetectedAppId: "dEtEcTeDaPpId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEtEcTeDaPpS/dEtEcTeDaPpId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDetectedAppIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementDetectedAppId(t *testing.T) {
	segments := DeviceManagementDetectedAppId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDetectedAppId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
