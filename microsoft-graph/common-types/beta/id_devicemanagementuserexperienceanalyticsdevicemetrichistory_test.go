package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId{}

func TestNewDeviceManagementUserExperienceAnalyticsDeviceMetricHistoryID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsDeviceMetricHistoryID("userExperienceAnalyticsMetricHistoryId")

	if id.UserExperienceAnalyticsMetricHistoryId != "userExperienceAnalyticsMetricHistoryId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsMetricHistoryId'", id.UserExperienceAnalyticsMetricHistoryId, "userExperienceAnalyticsMetricHistoryId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsDeviceMetricHistoryID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsDeviceMetricHistoryID("userExperienceAnalyticsMetricHistoryId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsDeviceMetricHistory/userExperienceAnalyticsMetricHistoryId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsDeviceMetricHistoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId
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
			Input: "/deviceManagement/userExperienceAnalyticsDeviceMetricHistory",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsDeviceMetricHistory/userExperienceAnalyticsMetricHistoryId",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId{
				UserExperienceAnalyticsMetricHistoryId: "userExperienceAnalyticsMetricHistoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsDeviceMetricHistory/userExperienceAnalyticsMetricHistoryId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsDeviceMetricHistoryID(v.Input)
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

func TestParseDeviceManagementUserExperienceAnalyticsDeviceMetricHistoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId
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
			Input: "/deviceManagement/userExperienceAnalyticsDeviceMetricHistory",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeMeTrIcHiStOrY",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsDeviceMetricHistory/userExperienceAnalyticsMetricHistoryId",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId{
				UserExperienceAnalyticsMetricHistoryId: "userExperienceAnalyticsMetricHistoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsDeviceMetricHistory/userExperienceAnalyticsMetricHistoryId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeMeTrIcHiStOrY/uSeReXpErIeNcEaNaLyTiCsMeTrIcHiStOrYiD",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId{
				UserExperienceAnalyticsMetricHistoryId: "uSeReXpErIeNcEaNaLyTiCsMeTrIcHiStOrYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeMeTrIcHiStOrY/uSeReXpErIeNcEaNaLyTiCsMeTrIcHiStOrYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsDeviceMetricHistoryIDInsensitively(v.Input)
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

func TestSegmentsForDeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
