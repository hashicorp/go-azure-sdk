package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId{}

func TestNewDeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceID("userExperienceAnalyticsAppHealthDeviceModelPerformanceId")

	if id.UserExperienceAnalyticsAppHealthDeviceModelPerformanceId != "userExperienceAnalyticsAppHealthDeviceModelPerformanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsAppHealthDeviceModelPerformanceId'", id.UserExperienceAnalyticsAppHealthDeviceModelPerformanceId, "userExperienceAnalyticsAppHealthDeviceModelPerformanceId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceID("userExperienceAnalyticsAppHealthDeviceModelPerformanceId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsAppHealthDeviceModelPerformance/userExperienceAnalyticsAppHealthDeviceModelPerformanceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthDeviceModelPerformance",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthDeviceModelPerformance/userExperienceAnalyticsAppHealthDeviceModelPerformanceId",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId{
				UserExperienceAnalyticsAppHealthDeviceModelPerformanceId: "userExperienceAnalyticsAppHealthDeviceModelPerformanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthDeviceModelPerformance/userExperienceAnalyticsAppHealthDeviceModelPerformanceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAppHealthDeviceModelPerformanceId != v.Expected.UserExperienceAnalyticsAppHealthDeviceModelPerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAppHealthDeviceModelPerformanceId", v.Expected.UserExperienceAnalyticsAppHealthDeviceModelPerformanceId, actual.UserExperienceAnalyticsAppHealthDeviceModelPerformanceId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthDeviceModelPerformance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHdEvIcEmOdElPeRfOrMaNcE",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthDeviceModelPerformance/userExperienceAnalyticsAppHealthDeviceModelPerformanceId",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId{
				UserExperienceAnalyticsAppHealthDeviceModelPerformanceId: "userExperienceAnalyticsAppHealthDeviceModelPerformanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthDeviceModelPerformance/userExperienceAnalyticsAppHealthDeviceModelPerformanceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHdEvIcEmOdElPeRfOrMaNcE/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHdEvIcEmOdElPeRfOrMaNcEiD",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId{
				UserExperienceAnalyticsAppHealthDeviceModelPerformanceId: "uSeReXpErIeNcEaNaLyTiCsApPhEaLtHdEvIcEmOdElPeRfOrMaNcEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHdEvIcEmOdElPeRfOrMaNcE/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHdEvIcEmOdElPeRfOrMaNcEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAppHealthDeviceModelPerformanceId != v.Expected.UserExperienceAnalyticsAppHealthDeviceModelPerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAppHealthDeviceModelPerformanceId", v.Expected.UserExperienceAnalyticsAppHealthDeviceModelPerformanceId, actual.UserExperienceAnalyticsAppHealthDeviceModelPerformanceId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
