package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId{}

func TestNewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdID("userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId")

	if id.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId != "userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId'", id.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId, "userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdID("userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId/userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId
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
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId/userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId{
				UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId: "userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId/userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId != v.Expected.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId", v.Expected.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId, actual.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId
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
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpLiCaTiOnPeRfOrMaNcEbYaPpVeRsIoNdEvIcEiD",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId/userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId{
				UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId: "userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId/userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpLiCaTiOnPeRfOrMaNcEbYaPpVeRsIoNdEvIcEiD/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpPeRfOrMaNcEbYaPpVeRsIoNdEvIcEiDiD",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId{
				UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId: "uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpPeRfOrMaNcEbYaPpVeRsIoNdEvIcEiDiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpLiCaTiOnPeRfOrMaNcEbYaPpVeRsIoNdEvIcEiD/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpPeRfOrMaNcEbYaPpVeRsIoNdEvIcEiDiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId != v.Expected.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId", v.Expected.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId, actual.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
