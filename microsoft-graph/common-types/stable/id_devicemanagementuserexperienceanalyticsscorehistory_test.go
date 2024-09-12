package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsScoreHistoryId{}

func TestNewDeviceManagementUserExperienceAnalyticsScoreHistoryID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsScoreHistoryID("userExperienceAnalyticsScoreHistoryIdValue")

	if id.UserExperienceAnalyticsScoreHistoryId != "userExperienceAnalyticsScoreHistoryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsScoreHistoryId'", id.UserExperienceAnalyticsScoreHistoryId, "userExperienceAnalyticsScoreHistoryIdValue")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsScoreHistoryID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsScoreHistoryID("userExperienceAnalyticsScoreHistoryIdValue").ID()
	expected := "/deviceManagement/userExperienceAnalyticsScoreHistory/userExperienceAnalyticsScoreHistoryIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsScoreHistoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsScoreHistoryId
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
			Input: "/deviceManagement/userExperienceAnalyticsScoreHistory",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsScoreHistory/userExperienceAnalyticsScoreHistoryIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsScoreHistoryId{
				UserExperienceAnalyticsScoreHistoryId: "userExperienceAnalyticsScoreHistoryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsScoreHistory/userExperienceAnalyticsScoreHistoryIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsScoreHistoryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsScoreHistoryId != v.Expected.UserExperienceAnalyticsScoreHistoryId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsScoreHistoryId", v.Expected.UserExperienceAnalyticsScoreHistoryId, actual.UserExperienceAnalyticsScoreHistoryId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsScoreHistoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsScoreHistoryId
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
			Input: "/deviceManagement/userExperienceAnalyticsScoreHistory",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsScOrEhIsToRy",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsScoreHistory/userExperienceAnalyticsScoreHistoryIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsScoreHistoryId{
				UserExperienceAnalyticsScoreHistoryId: "userExperienceAnalyticsScoreHistoryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsScoreHistory/userExperienceAnalyticsScoreHistoryIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsScOrEhIsToRy/uSeReXpErIeNcEaNaLyTiCsScOrEhIsToRyIdVaLuE",
			Expected: &DeviceManagementUserExperienceAnalyticsScoreHistoryId{
				UserExperienceAnalyticsScoreHistoryId: "uSeReXpErIeNcEaNaLyTiCsScOrEhIsToRyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsScOrEhIsToRy/uSeReXpErIeNcEaNaLyTiCsScOrEhIsToRyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsScoreHistoryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsScoreHistoryId != v.Expected.UserExperienceAnalyticsScoreHistoryId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsScoreHistoryId", v.Expected.UserExperienceAnalyticsScoreHistoryId, actual.UserExperienceAnalyticsScoreHistoryId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsScoreHistoryId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsScoreHistoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsScoreHistoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
