package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId{}

func TestNewDeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceID("userExperienceAnalyticsBatteryHealthModelPerformanceId")

	if id.UserExperienceAnalyticsBatteryHealthModelPerformanceId != "userExperienceAnalyticsBatteryHealthModelPerformanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsBatteryHealthModelPerformanceId'", id.UserExperienceAnalyticsBatteryHealthModelPerformanceId, "userExperienceAnalyticsBatteryHealthModelPerformanceId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceID("userExperienceAnalyticsBatteryHealthModelPerformanceId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsBatteryHealthModelPerformance/userExperienceAnalyticsBatteryHealthModelPerformanceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthModelPerformance",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthModelPerformance/userExperienceAnalyticsBatteryHealthModelPerformanceId",
			Expected: &DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId{
				UserExperienceAnalyticsBatteryHealthModelPerformanceId: "userExperienceAnalyticsBatteryHealthModelPerformanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthModelPerformance/userExperienceAnalyticsBatteryHealthModelPerformanceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsBatteryHealthModelPerformanceId != v.Expected.UserExperienceAnalyticsBatteryHealthModelPerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsBatteryHealthModelPerformanceId", v.Expected.UserExperienceAnalyticsBatteryHealthModelPerformanceId, actual.UserExperienceAnalyticsBatteryHealthModelPerformanceId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthModelPerformance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHmOdElPeRfOrMaNcE",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthModelPerformance/userExperienceAnalyticsBatteryHealthModelPerformanceId",
			Expected: &DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId{
				UserExperienceAnalyticsBatteryHealthModelPerformanceId: "userExperienceAnalyticsBatteryHealthModelPerformanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthModelPerformance/userExperienceAnalyticsBatteryHealthModelPerformanceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHmOdElPeRfOrMaNcE/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHmOdElPeRfOrMaNcEiD",
			Expected: &DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId{
				UserExperienceAnalyticsBatteryHealthModelPerformanceId: "uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHmOdElPeRfOrMaNcEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHmOdElPeRfOrMaNcE/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHmOdElPeRfOrMaNcEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsBatteryHealthModelPerformanceId != v.Expected.UserExperienceAnalyticsBatteryHealthModelPerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsBatteryHealthModelPerformanceId", v.Expected.UserExperienceAnalyticsBatteryHealthModelPerformanceId, actual.UserExperienceAnalyticsBatteryHealthModelPerformanceId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
