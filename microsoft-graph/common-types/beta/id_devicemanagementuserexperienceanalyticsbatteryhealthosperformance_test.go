package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId{}

func TestNewDeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceID("userExperienceAnalyticsBatteryHealthOsPerformanceId")

	if id.UserExperienceAnalyticsBatteryHealthOsPerformanceId != "userExperienceAnalyticsBatteryHealthOsPerformanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsBatteryHealthOsPerformanceId'", id.UserExperienceAnalyticsBatteryHealthOsPerformanceId, "userExperienceAnalyticsBatteryHealthOsPerformanceId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceID("userExperienceAnalyticsBatteryHealthOsPerformanceId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsBatteryHealthOsPerformance/userExperienceAnalyticsBatteryHealthOsPerformanceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthOsPerformance",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthOsPerformance/userExperienceAnalyticsBatteryHealthOsPerformanceId",
			Expected: &DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId{
				UserExperienceAnalyticsBatteryHealthOsPerformanceId: "userExperienceAnalyticsBatteryHealthOsPerformanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthOsPerformance/userExperienceAnalyticsBatteryHealthOsPerformanceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsBatteryHealthOsPerformanceId != v.Expected.UserExperienceAnalyticsBatteryHealthOsPerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsBatteryHealthOsPerformanceId", v.Expected.UserExperienceAnalyticsBatteryHealthOsPerformanceId, actual.UserExperienceAnalyticsBatteryHealthOsPerformanceId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthOsPerformance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHoSpErFoRmAnCe",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthOsPerformance/userExperienceAnalyticsBatteryHealthOsPerformanceId",
			Expected: &DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId{
				UserExperienceAnalyticsBatteryHealthOsPerformanceId: "userExperienceAnalyticsBatteryHealthOsPerformanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthOsPerformance/userExperienceAnalyticsBatteryHealthOsPerformanceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHoSpErFoRmAnCe/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHoSpErFoRmAnCeId",
			Expected: &DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId{
				UserExperienceAnalyticsBatteryHealthOsPerformanceId: "uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHoSpErFoRmAnCeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHoSpErFoRmAnCe/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHoSpErFoRmAnCeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsBatteryHealthOsPerformanceId != v.Expected.UserExperienceAnalyticsBatteryHealthOsPerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsBatteryHealthOsPerformanceId", v.Expected.UserExperienceAnalyticsBatteryHealthOsPerformanceId, actual.UserExperienceAnalyticsBatteryHealthOsPerformanceId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
