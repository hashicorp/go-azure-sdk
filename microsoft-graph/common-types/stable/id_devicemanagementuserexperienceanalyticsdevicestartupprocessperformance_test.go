package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId{}

func TestNewDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceID("userExperienceAnalyticsDeviceStartupProcessPerformanceIdValue")

	if id.UserExperienceAnalyticsDeviceStartupProcessPerformanceId != "userExperienceAnalyticsDeviceStartupProcessPerformanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsDeviceStartupProcessPerformanceId'", id.UserExperienceAnalyticsDeviceStartupProcessPerformanceId, "userExperienceAnalyticsDeviceStartupProcessPerformanceIdValue")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceID("userExperienceAnalyticsDeviceStartupProcessPerformanceIdValue").ID()
	expected := "/deviceManagement/userExperienceAnalyticsDeviceStartupProcessPerformance/userExperienceAnalyticsDeviceStartupProcessPerformanceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsDeviceStartupProcessPerformance",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsDeviceStartupProcessPerformance/userExperienceAnalyticsDeviceStartupProcessPerformanceIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId{
				UserExperienceAnalyticsDeviceStartupProcessPerformanceId: "userExperienceAnalyticsDeviceStartupProcessPerformanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsDeviceStartupProcessPerformance/userExperienceAnalyticsDeviceStartupProcessPerformanceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsDeviceStartupProcessPerformanceId != v.Expected.UserExperienceAnalyticsDeviceStartupProcessPerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsDeviceStartupProcessPerformanceId", v.Expected.UserExperienceAnalyticsDeviceStartupProcessPerformanceId, actual.UserExperienceAnalyticsDeviceStartupProcessPerformanceId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsDeviceStartupProcessPerformance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeStArTuPpRoCeSsPeRfOrMaNcE",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsDeviceStartupProcessPerformance/userExperienceAnalyticsDeviceStartupProcessPerformanceIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId{
				UserExperienceAnalyticsDeviceStartupProcessPerformanceId: "userExperienceAnalyticsDeviceStartupProcessPerformanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsDeviceStartupProcessPerformance/userExperienceAnalyticsDeviceStartupProcessPerformanceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeStArTuPpRoCeSsPeRfOrMaNcE/uSeReXpErIeNcEaNaLyTiCsDeViCeStArTuPpRoCeSsPeRfOrMaNcEiDvAlUe",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId{
				UserExperienceAnalyticsDeviceStartupProcessPerformanceId: "uSeReXpErIeNcEaNaLyTiCsDeViCeStArTuPpRoCeSsPeRfOrMaNcEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeStArTuPpRoCeSsPeRfOrMaNcE/uSeReXpErIeNcEaNaLyTiCsDeViCeStArTuPpRoCeSsPeRfOrMaNcEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsDeviceStartupProcessPerformanceId != v.Expected.UserExperienceAnalyticsDeviceStartupProcessPerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsDeviceStartupProcessPerformanceId", v.Expected.UserExperienceAnalyticsDeviceStartupProcessPerformanceId, actual.UserExperienceAnalyticsDeviceStartupProcessPerformanceId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
