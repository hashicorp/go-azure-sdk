package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId{}

func TestNewDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceID("userExperienceAnalyticsAppHealthDevicePerformanceId")

	if id.UserExperienceAnalyticsAppHealthDevicePerformanceId != "userExperienceAnalyticsAppHealthDevicePerformanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsAppHealthDevicePerformanceId'", id.UserExperienceAnalyticsAppHealthDevicePerformanceId, "userExperienceAnalyticsAppHealthDevicePerformanceId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceID("userExperienceAnalyticsAppHealthDevicePerformanceId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsAppHealthDevicePerformance/userExperienceAnalyticsAppHealthDevicePerformanceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthDevicePerformance",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthDevicePerformance/userExperienceAnalyticsAppHealthDevicePerformanceId",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId{
				UserExperienceAnalyticsAppHealthDevicePerformanceId: "userExperienceAnalyticsAppHealthDevicePerformanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthDevicePerformance/userExperienceAnalyticsAppHealthDevicePerformanceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAppHealthDevicePerformanceId != v.Expected.UserExperienceAnalyticsAppHealthDevicePerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAppHealthDevicePerformanceId", v.Expected.UserExperienceAnalyticsAppHealthDevicePerformanceId, actual.UserExperienceAnalyticsAppHealthDevicePerformanceId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthDevicePerformance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHdEvIcEpErFoRmAnCe",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthDevicePerformance/userExperienceAnalyticsAppHealthDevicePerformanceId",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId{
				UserExperienceAnalyticsAppHealthDevicePerformanceId: "userExperienceAnalyticsAppHealthDevicePerformanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthDevicePerformance/userExperienceAnalyticsAppHealthDevicePerformanceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHdEvIcEpErFoRmAnCe/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHdEvIcEpErFoRmAnCeId",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId{
				UserExperienceAnalyticsAppHealthDevicePerformanceId: "uSeReXpErIeNcEaNaLyTiCsApPhEaLtHdEvIcEpErFoRmAnCeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHdEvIcEpErFoRmAnCe/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHdEvIcEpErFoRmAnCeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAppHealthDevicePerformanceId != v.Expected.UserExperienceAnalyticsAppHealthDevicePerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAppHealthDevicePerformanceId", v.Expected.UserExperienceAnalyticsAppHealthDevicePerformanceId, actual.UserExperienceAnalyticsAppHealthDevicePerformanceId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
