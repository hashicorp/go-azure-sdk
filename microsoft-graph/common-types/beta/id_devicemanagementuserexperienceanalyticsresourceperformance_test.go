package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsResourcePerformanceId{}

func TestNewDeviceManagementUserExperienceAnalyticsResourcePerformanceID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsResourcePerformanceID("userExperienceAnalyticsResourcePerformanceId")

	if id.UserExperienceAnalyticsResourcePerformanceId != "userExperienceAnalyticsResourcePerformanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsResourcePerformanceId'", id.UserExperienceAnalyticsResourcePerformanceId, "userExperienceAnalyticsResourcePerformanceId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsResourcePerformanceID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsResourcePerformanceID("userExperienceAnalyticsResourcePerformanceId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsResourcePerformance/userExperienceAnalyticsResourcePerformanceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsResourcePerformanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsResourcePerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsResourcePerformance",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsResourcePerformance/userExperienceAnalyticsResourcePerformanceId",
			Expected: &DeviceManagementUserExperienceAnalyticsResourcePerformanceId{
				UserExperienceAnalyticsResourcePerformanceId: "userExperienceAnalyticsResourcePerformanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsResourcePerformance/userExperienceAnalyticsResourcePerformanceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsResourcePerformanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsResourcePerformanceId != v.Expected.UserExperienceAnalyticsResourcePerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsResourcePerformanceId", v.Expected.UserExperienceAnalyticsResourcePerformanceId, actual.UserExperienceAnalyticsResourcePerformanceId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsResourcePerformanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsResourcePerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsResourcePerformance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsReSoUrCePeRfOrMaNcE",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsResourcePerformance/userExperienceAnalyticsResourcePerformanceId",
			Expected: &DeviceManagementUserExperienceAnalyticsResourcePerformanceId{
				UserExperienceAnalyticsResourcePerformanceId: "userExperienceAnalyticsResourcePerformanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsResourcePerformance/userExperienceAnalyticsResourcePerformanceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsReSoUrCePeRfOrMaNcE/uSeReXpErIeNcEaNaLyTiCsReSoUrCePeRfOrMaNcEiD",
			Expected: &DeviceManagementUserExperienceAnalyticsResourcePerformanceId{
				UserExperienceAnalyticsResourcePerformanceId: "uSeReXpErIeNcEaNaLyTiCsReSoUrCePeRfOrMaNcEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsReSoUrCePeRfOrMaNcE/uSeReXpErIeNcEaNaLyTiCsReSoUrCePeRfOrMaNcEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsResourcePerformanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsResourcePerformanceId != v.Expected.UserExperienceAnalyticsResourcePerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsResourcePerformanceId", v.Expected.UserExperienceAnalyticsResourcePerformanceId, actual.UserExperienceAnalyticsResourcePerformanceId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsResourcePerformanceId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsResourcePerformanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsResourcePerformanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
