package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsMetricHistoryId{}

func TestNewDeviceManagementUserExperienceAnalyticsMetricHistoryID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsMetricHistoryID("userExperienceAnalyticsMetricHistoryId")

	if id.UserExperienceAnalyticsMetricHistoryId != "userExperienceAnalyticsMetricHistoryId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsMetricHistoryId'", id.UserExperienceAnalyticsMetricHistoryId, "userExperienceAnalyticsMetricHistoryId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsMetricHistoryID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsMetricHistoryID("userExperienceAnalyticsMetricHistoryId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsMetricHistory/userExperienceAnalyticsMetricHistoryId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsMetricHistoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsMetricHistoryId
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
			Input: "/deviceManagement/userExperienceAnalyticsMetricHistory",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsMetricHistory/userExperienceAnalyticsMetricHistoryId",
			Expected: &DeviceManagementUserExperienceAnalyticsMetricHistoryId{
				UserExperienceAnalyticsMetricHistoryId: "userExperienceAnalyticsMetricHistoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsMetricHistory/userExperienceAnalyticsMetricHistoryId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsMetricHistoryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsMetricHistoryId != v.Expected.UserExperienceAnalyticsMetricHistoryId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsMetricHistoryId", v.Expected.UserExperienceAnalyticsMetricHistoryId, actual.UserExperienceAnalyticsMetricHistoryId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsMetricHistoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsMetricHistoryId
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
			Input: "/deviceManagement/userExperienceAnalyticsMetricHistory",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsMeTrIcHiStOrY",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsMetricHistory/userExperienceAnalyticsMetricHistoryId",
			Expected: &DeviceManagementUserExperienceAnalyticsMetricHistoryId{
				UserExperienceAnalyticsMetricHistoryId: "userExperienceAnalyticsMetricHistoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsMetricHistory/userExperienceAnalyticsMetricHistoryId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsMeTrIcHiStOrY/uSeReXpErIeNcEaNaLyTiCsMeTrIcHiStOrYiD",
			Expected: &DeviceManagementUserExperienceAnalyticsMetricHistoryId{
				UserExperienceAnalyticsMetricHistoryId: "uSeReXpErIeNcEaNaLyTiCsMeTrIcHiStOrYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsMeTrIcHiStOrY/uSeReXpErIeNcEaNaLyTiCsMeTrIcHiStOrYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsMetricHistoryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsMetricHistoryId != v.Expected.UserExperienceAnalyticsMetricHistoryId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsMetricHistoryId", v.Expected.UserExperienceAnalyticsMetricHistoryId, actual.UserExperienceAnalyticsMetricHistoryId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsMetricHistoryId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsMetricHistoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsMetricHistoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
