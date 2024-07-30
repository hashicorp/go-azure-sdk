package userexperienceanalyticsapphealthapplicationperformance

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId{}

func TestNewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceID("userExperienceAnalyticsAppHealthApplicationPerformanceIdValue")

	if id.UserExperienceAnalyticsAppHealthApplicationPerformanceId != "userExperienceAnalyticsAppHealthApplicationPerformanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsAppHealthApplicationPerformanceId'", id.UserExperienceAnalyticsAppHealthApplicationPerformanceId, "userExperienceAnalyticsAppHealthApplicationPerformanceIdValue")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceID("userExperienceAnalyticsAppHealthApplicationPerformanceIdValue").ID()
	expected := "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformance/userExperienceAnalyticsAppHealthApplicationPerformanceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformance",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformance/userExperienceAnalyticsAppHealthApplicationPerformanceIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId{
				UserExperienceAnalyticsAppHealthApplicationPerformanceId: "userExperienceAnalyticsAppHealthApplicationPerformanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformance/userExperienceAnalyticsAppHealthApplicationPerformanceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAppHealthApplicationPerformanceId != v.Expected.UserExperienceAnalyticsAppHealthApplicationPerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAppHealthApplicationPerformanceId", v.Expected.UserExperienceAnalyticsAppHealthApplicationPerformanceId, actual.UserExperienceAnalyticsAppHealthApplicationPerformanceId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpLiCaTiOnPeRfOrMaNcE",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformance/userExperienceAnalyticsAppHealthApplicationPerformanceIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId{
				UserExperienceAnalyticsAppHealthApplicationPerformanceId: "userExperienceAnalyticsAppHealthApplicationPerformanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformance/userExperienceAnalyticsAppHealthApplicationPerformanceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpLiCaTiOnPeRfOrMaNcE/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpLiCaTiOnPeRfOrMaNcEiDvAlUe",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId{
				UserExperienceAnalyticsAppHealthApplicationPerformanceId: "uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpLiCaTiOnPeRfOrMaNcEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpLiCaTiOnPeRfOrMaNcE/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpLiCaTiOnPeRfOrMaNcEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAppHealthApplicationPerformanceId != v.Expected.UserExperienceAnalyticsAppHealthApplicationPerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAppHealthApplicationPerformanceId", v.Expected.UserExperienceAnalyticsAppHealthApplicationPerformanceId, actual.UserExperienceAnalyticsAppHealthApplicationPerformanceId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
