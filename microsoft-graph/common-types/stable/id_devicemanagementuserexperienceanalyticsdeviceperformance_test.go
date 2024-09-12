package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsDevicePerformanceId{}

func TestNewDeviceManagementUserExperienceAnalyticsDevicePerformanceID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsDevicePerformanceID("userExperienceAnalyticsDevicePerformanceIdValue")

	if id.UserExperienceAnalyticsDevicePerformanceId != "userExperienceAnalyticsDevicePerformanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsDevicePerformanceId'", id.UserExperienceAnalyticsDevicePerformanceId, "userExperienceAnalyticsDevicePerformanceIdValue")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsDevicePerformanceID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsDevicePerformanceID("userExperienceAnalyticsDevicePerformanceIdValue").ID()
	expected := "/deviceManagement/userExperienceAnalyticsDevicePerformance/userExperienceAnalyticsDevicePerformanceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsDevicePerformanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsDevicePerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsDevicePerformance",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsDevicePerformance/userExperienceAnalyticsDevicePerformanceIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsDevicePerformanceId{
				UserExperienceAnalyticsDevicePerformanceId: "userExperienceAnalyticsDevicePerformanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsDevicePerformance/userExperienceAnalyticsDevicePerformanceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsDevicePerformanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsDevicePerformanceId != v.Expected.UserExperienceAnalyticsDevicePerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsDevicePerformanceId", v.Expected.UserExperienceAnalyticsDevicePerformanceId, actual.UserExperienceAnalyticsDevicePerformanceId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsDevicePerformanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsDevicePerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsDevicePerformance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCePeRfOrMaNcE",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsDevicePerformance/userExperienceAnalyticsDevicePerformanceIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsDevicePerformanceId{
				UserExperienceAnalyticsDevicePerformanceId: "userExperienceAnalyticsDevicePerformanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsDevicePerformance/userExperienceAnalyticsDevicePerformanceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCePeRfOrMaNcE/uSeReXpErIeNcEaNaLyTiCsDeViCePeRfOrMaNcEiDvAlUe",
			Expected: &DeviceManagementUserExperienceAnalyticsDevicePerformanceId{
				UserExperienceAnalyticsDevicePerformanceId: "uSeReXpErIeNcEaNaLyTiCsDeViCePeRfOrMaNcEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCePeRfOrMaNcE/uSeReXpErIeNcEaNaLyTiCsDeViCePeRfOrMaNcEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsDevicePerformanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsDevicePerformanceId != v.Expected.UserExperienceAnalyticsDevicePerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsDevicePerformanceId", v.Expected.UserExperienceAnalyticsDevicePerformanceId, actual.UserExperienceAnalyticsDevicePerformanceId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsDevicePerformanceId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsDevicePerformanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsDevicePerformanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
