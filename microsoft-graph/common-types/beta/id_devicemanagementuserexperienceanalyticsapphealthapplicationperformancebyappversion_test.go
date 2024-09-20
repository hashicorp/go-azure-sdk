package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId{}

func TestNewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionID("userExperienceAnalyticsAppHealthAppPerformanceByAppVersionId")

	if id.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId != "userExperienceAnalyticsAppHealthAppPerformanceByAppVersionId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId'", id.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId, "userExperienceAnalyticsAppHealthAppPerformanceByAppVersionId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionID("userExperienceAnalyticsAppHealthAppPerformanceByAppVersionId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion/userExperienceAnalyticsAppHealthAppPerformanceByAppVersionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId
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
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion/userExperienceAnalyticsAppHealthAppPerformanceByAppVersionId",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId{
				UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId: "userExperienceAnalyticsAppHealthAppPerformanceByAppVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion/userExperienceAnalyticsAppHealthAppPerformanceByAppVersionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId != v.Expected.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId", v.Expected.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId, actual.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId
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
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpLiCaTiOnPeRfOrMaNcEbYaPpVeRsIoN",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion/userExperienceAnalyticsAppHealthAppPerformanceByAppVersionId",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId{
				UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId: "userExperienceAnalyticsAppHealthAppPerformanceByAppVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion/userExperienceAnalyticsAppHealthAppPerformanceByAppVersionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpLiCaTiOnPeRfOrMaNcEbYaPpVeRsIoN/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpPeRfOrMaNcEbYaPpVeRsIoNiD",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId{
				UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId: "uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpPeRfOrMaNcEbYaPpVeRsIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpLiCaTiOnPeRfOrMaNcEbYaPpVeRsIoN/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpPeRfOrMaNcEbYaPpVeRsIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId != v.Expected.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId", v.Expected.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId, actual.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
