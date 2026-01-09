package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId{}

func TestNewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionID("userExperienceAnalyticsAppHealthAppPerformanceByOSVersionId")

	if id.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId != "userExperienceAnalyticsAppHealthAppPerformanceByOSVersionId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId'", id.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId, "userExperienceAnalyticsAppHealthAppPerformanceByOSVersionId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionID("userExperienceAnalyticsAppHealthAppPerformanceByOSVersionId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion/userExperienceAnalyticsAppHealthAppPerformanceByOSVersionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId
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
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion/userExperienceAnalyticsAppHealthAppPerformanceByOSVersionId",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId{
				UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId: "userExperienceAnalyticsAppHealthAppPerformanceByOSVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion/userExperienceAnalyticsAppHealthAppPerformanceByOSVersionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId != v.Expected.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId", v.Expected.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId, actual.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId
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
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpLiCaTiOnPeRfOrMaNcEbYoSvErSiOn",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion/userExperienceAnalyticsAppHealthAppPerformanceByOSVersionId",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId{
				UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId: "userExperienceAnalyticsAppHealthAppPerformanceByOSVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion/userExperienceAnalyticsAppHealthAppPerformanceByOSVersionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpLiCaTiOnPeRfOrMaNcEbYoSvErSiOn/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpPeRfOrMaNcEbYoSvErSiOnId",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId{
				UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId: "uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpPeRfOrMaNcEbYoSvErSiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpLiCaTiOnPeRfOrMaNcEbYoSvErSiOn/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpPeRfOrMaNcEbYoSvErSiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId != v.Expected.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId", v.Expected.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId, actual.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
