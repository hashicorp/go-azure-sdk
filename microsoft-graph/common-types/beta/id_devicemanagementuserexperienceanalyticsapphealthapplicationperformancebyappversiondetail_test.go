package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId{}

func TestNewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailID("userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId")

	if id.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId != "userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId'", id.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId, "userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailID("userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails/userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId
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
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails/userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId{
				UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId: "userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails/userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId != v.Expected.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId", v.Expected.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId, actual.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId
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
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpLiCaTiOnPeRfOrMaNcEbYaPpVeRsIoNdEtAiLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails/userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId{
				UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId: "userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails/userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpLiCaTiOnPeRfOrMaNcEbYaPpVeRsIoNdEtAiLs/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpPeRfOrMaNcEbYaPpVeRsIoNdEtAiLsId",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId{
				UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId: "uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpPeRfOrMaNcEbYaPpVeRsIoNdEtAiLsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpLiCaTiOnPeRfOrMaNcEbYaPpVeRsIoNdEtAiLs/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHaPpPeRfOrMaNcEbYaPpVeRsIoNdEtAiLsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId != v.Expected.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId", v.Expected.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId, actual.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
