package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId{}

func TestNewDeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceID("userExperienceAnalyticsBatteryHealthDevicePerformanceId")

	if id.UserExperienceAnalyticsBatteryHealthDevicePerformanceId != "userExperienceAnalyticsBatteryHealthDevicePerformanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsBatteryHealthDevicePerformanceId'", id.UserExperienceAnalyticsBatteryHealthDevicePerformanceId, "userExperienceAnalyticsBatteryHealthDevicePerformanceId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceID("userExperienceAnalyticsBatteryHealthDevicePerformanceId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsBatteryHealthDevicePerformance/userExperienceAnalyticsBatteryHealthDevicePerformanceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthDevicePerformance",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthDevicePerformance/userExperienceAnalyticsBatteryHealthDevicePerformanceId",
			Expected: &DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId{
				UserExperienceAnalyticsBatteryHealthDevicePerformanceId: "userExperienceAnalyticsBatteryHealthDevicePerformanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthDevicePerformance/userExperienceAnalyticsBatteryHealthDevicePerformanceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsBatteryHealthDevicePerformanceId != v.Expected.UserExperienceAnalyticsBatteryHealthDevicePerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsBatteryHealthDevicePerformanceId", v.Expected.UserExperienceAnalyticsBatteryHealthDevicePerformanceId, actual.UserExperienceAnalyticsBatteryHealthDevicePerformanceId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthDevicePerformance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHdEvIcEpErFoRmAnCe",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthDevicePerformance/userExperienceAnalyticsBatteryHealthDevicePerformanceId",
			Expected: &DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId{
				UserExperienceAnalyticsBatteryHealthDevicePerformanceId: "userExperienceAnalyticsBatteryHealthDevicePerformanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthDevicePerformance/userExperienceAnalyticsBatteryHealthDevicePerformanceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHdEvIcEpErFoRmAnCe/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHdEvIcEpErFoRmAnCeId",
			Expected: &DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId{
				UserExperienceAnalyticsBatteryHealthDevicePerformanceId: "uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHdEvIcEpErFoRmAnCeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHdEvIcEpErFoRmAnCe/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHdEvIcEpErFoRmAnCeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsBatteryHealthDevicePerformanceId != v.Expected.UserExperienceAnalyticsBatteryHealthDevicePerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsBatteryHealthDevicePerformanceId", v.Expected.UserExperienceAnalyticsBatteryHealthDevicePerformanceId, actual.UserExperienceAnalyticsBatteryHealthDevicePerformanceId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
