package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsDeviceScoreId{}

func TestNewDeviceManagementUserExperienceAnalyticsDeviceScoreID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsDeviceScoreID("userExperienceAnalyticsDeviceScoresId")

	if id.UserExperienceAnalyticsDeviceScoresId != "userExperienceAnalyticsDeviceScoresId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsDeviceScoresId'", id.UserExperienceAnalyticsDeviceScoresId, "userExperienceAnalyticsDeviceScoresId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsDeviceScoreID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsDeviceScoreID("userExperienceAnalyticsDeviceScoresId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsDeviceScores/userExperienceAnalyticsDeviceScoresId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsDeviceScoreID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsDeviceScoreId
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
			Input: "/deviceManagement/userExperienceAnalyticsDeviceScores",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsDeviceScores/userExperienceAnalyticsDeviceScoresId",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceScoreId{
				UserExperienceAnalyticsDeviceScoresId: "userExperienceAnalyticsDeviceScoresId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsDeviceScores/userExperienceAnalyticsDeviceScoresId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsDeviceScoreID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsDeviceScoresId != v.Expected.UserExperienceAnalyticsDeviceScoresId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsDeviceScoresId", v.Expected.UserExperienceAnalyticsDeviceScoresId, actual.UserExperienceAnalyticsDeviceScoresId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsDeviceScoreIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsDeviceScoreId
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
			Input: "/deviceManagement/userExperienceAnalyticsDeviceScores",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeScOrEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsDeviceScores/userExperienceAnalyticsDeviceScoresId",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceScoreId{
				UserExperienceAnalyticsDeviceScoresId: "userExperienceAnalyticsDeviceScoresId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsDeviceScores/userExperienceAnalyticsDeviceScoresId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeScOrEs/uSeReXpErIeNcEaNaLyTiCsDeViCeScOrEsId",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceScoreId{
				UserExperienceAnalyticsDeviceScoresId: "uSeReXpErIeNcEaNaLyTiCsDeViCeScOrEsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeScOrEs/uSeReXpErIeNcEaNaLyTiCsDeViCeScOrEsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsDeviceScoreIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsDeviceScoresId != v.Expected.UserExperienceAnalyticsDeviceScoresId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsDeviceScoresId", v.Expected.UserExperienceAnalyticsDeviceScoresId, actual.UserExperienceAnalyticsDeviceScoresId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsDeviceScoreId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsDeviceScoreId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsDeviceScoreId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
