package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId{}

func TestNewDeviceManagementUserExperienceAnalyticsDeviceStartupHistoryID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsDeviceStartupHistoryID("userExperienceAnalyticsDeviceStartupHistoryIdValue")

	if id.UserExperienceAnalyticsDeviceStartupHistoryId != "userExperienceAnalyticsDeviceStartupHistoryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsDeviceStartupHistoryId'", id.UserExperienceAnalyticsDeviceStartupHistoryId, "userExperienceAnalyticsDeviceStartupHistoryIdValue")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsDeviceStartupHistoryID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsDeviceStartupHistoryID("userExperienceAnalyticsDeviceStartupHistoryIdValue").ID()
	expected := "/deviceManagement/userExperienceAnalyticsDeviceStartupHistory/userExperienceAnalyticsDeviceStartupHistoryIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsDeviceStartupHistoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId
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
			Input: "/deviceManagement/userExperienceAnalyticsDeviceStartupHistory",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsDeviceStartupHistory/userExperienceAnalyticsDeviceStartupHistoryIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId{
				UserExperienceAnalyticsDeviceStartupHistoryId: "userExperienceAnalyticsDeviceStartupHistoryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsDeviceStartupHistory/userExperienceAnalyticsDeviceStartupHistoryIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsDeviceStartupHistoryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsDeviceStartupHistoryId != v.Expected.UserExperienceAnalyticsDeviceStartupHistoryId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsDeviceStartupHistoryId", v.Expected.UserExperienceAnalyticsDeviceStartupHistoryId, actual.UserExperienceAnalyticsDeviceStartupHistoryId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsDeviceStartupHistoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId
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
			Input: "/deviceManagement/userExperienceAnalyticsDeviceStartupHistory",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeStArTuPhIsToRy",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsDeviceStartupHistory/userExperienceAnalyticsDeviceStartupHistoryIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId{
				UserExperienceAnalyticsDeviceStartupHistoryId: "userExperienceAnalyticsDeviceStartupHistoryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsDeviceStartupHistory/userExperienceAnalyticsDeviceStartupHistoryIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeStArTuPhIsToRy/uSeReXpErIeNcEaNaLyTiCsDeViCeStArTuPhIsToRyIdVaLuE",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId{
				UserExperienceAnalyticsDeviceStartupHistoryId: "uSeReXpErIeNcEaNaLyTiCsDeViCeStArTuPhIsToRyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeStArTuPhIsToRy/uSeReXpErIeNcEaNaLyTiCsDeViCeStArTuPhIsToRyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsDeviceStartupHistoryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsDeviceStartupHistoryId != v.Expected.UserExperienceAnalyticsDeviceStartupHistoryId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsDeviceStartupHistoryId", v.Expected.UserExperienceAnalyticsDeviceStartupHistoryId, actual.UserExperienceAnalyticsDeviceStartupHistoryId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
