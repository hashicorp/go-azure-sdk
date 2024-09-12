package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId{}

func TestNewDeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceID("userExperienceAnalyticsAppHealthOSVersionPerformanceIdValue")

	if id.UserExperienceAnalyticsAppHealthOSVersionPerformanceId != "userExperienceAnalyticsAppHealthOSVersionPerformanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsAppHealthOSVersionPerformanceId'", id.UserExperienceAnalyticsAppHealthOSVersionPerformanceId, "userExperienceAnalyticsAppHealthOSVersionPerformanceIdValue")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceID("userExperienceAnalyticsAppHealthOSVersionPerformanceIdValue").ID()
	expected := "/deviceManagement/userExperienceAnalyticsAppHealthOSVersionPerformance/userExperienceAnalyticsAppHealthOSVersionPerformanceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthOSVersionPerformance",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthOSVersionPerformance/userExperienceAnalyticsAppHealthOSVersionPerformanceIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId{
				UserExperienceAnalyticsAppHealthOSVersionPerformanceId: "userExperienceAnalyticsAppHealthOSVersionPerformanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthOSVersionPerformance/userExperienceAnalyticsAppHealthOSVersionPerformanceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAppHealthOSVersionPerformanceId != v.Expected.UserExperienceAnalyticsAppHealthOSVersionPerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAppHealthOSVersionPerformanceId", v.Expected.UserExperienceAnalyticsAppHealthOSVersionPerformanceId, actual.UserExperienceAnalyticsAppHealthOSVersionPerformanceId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthOSVersionPerformance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHoSvErSiOnPeRfOrMaNcE",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthOSVersionPerformance/userExperienceAnalyticsAppHealthOSVersionPerformanceIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId{
				UserExperienceAnalyticsAppHealthOSVersionPerformanceId: "userExperienceAnalyticsAppHealthOSVersionPerformanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthOSVersionPerformance/userExperienceAnalyticsAppHealthOSVersionPerformanceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHoSvErSiOnPeRfOrMaNcE/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHoSvErSiOnPeRfOrMaNcEiDvAlUe",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId{
				UserExperienceAnalyticsAppHealthOSVersionPerformanceId: "uSeReXpErIeNcEaNaLyTiCsApPhEaLtHoSvErSiOnPeRfOrMaNcEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHoSvErSiOnPeRfOrMaNcE/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHoSvErSiOnPeRfOrMaNcEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAppHealthOSVersionPerformanceId != v.Expected.UserExperienceAnalyticsAppHealthOSVersionPerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAppHealthOSVersionPerformanceId", v.Expected.UserExperienceAnalyticsAppHealthOSVersionPerformanceId, actual.UserExperienceAnalyticsAppHealthOSVersionPerformanceId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
