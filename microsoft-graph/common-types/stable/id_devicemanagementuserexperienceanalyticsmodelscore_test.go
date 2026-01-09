package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsModelScoreId{}

func TestNewDeviceManagementUserExperienceAnalyticsModelScoreID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsModelScoreID("userExperienceAnalyticsModelScoresId")

	if id.UserExperienceAnalyticsModelScoresId != "userExperienceAnalyticsModelScoresId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsModelScoresId'", id.UserExperienceAnalyticsModelScoresId, "userExperienceAnalyticsModelScoresId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsModelScoreID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsModelScoreID("userExperienceAnalyticsModelScoresId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsModelScores/userExperienceAnalyticsModelScoresId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsModelScoreID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsModelScoreId
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
			Input: "/deviceManagement/userExperienceAnalyticsModelScores",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsModelScores/userExperienceAnalyticsModelScoresId",
			Expected: &DeviceManagementUserExperienceAnalyticsModelScoreId{
				UserExperienceAnalyticsModelScoresId: "userExperienceAnalyticsModelScoresId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsModelScores/userExperienceAnalyticsModelScoresId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsModelScoreID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsModelScoresId != v.Expected.UserExperienceAnalyticsModelScoresId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsModelScoresId", v.Expected.UserExperienceAnalyticsModelScoresId, actual.UserExperienceAnalyticsModelScoresId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsModelScoreIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsModelScoreId
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
			Input: "/deviceManagement/userExperienceAnalyticsModelScores",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsMoDeLsCoReS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsModelScores/userExperienceAnalyticsModelScoresId",
			Expected: &DeviceManagementUserExperienceAnalyticsModelScoreId{
				UserExperienceAnalyticsModelScoresId: "userExperienceAnalyticsModelScoresId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsModelScores/userExperienceAnalyticsModelScoresId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsMoDeLsCoReS/uSeReXpErIeNcEaNaLyTiCsMoDeLsCoReSiD",
			Expected: &DeviceManagementUserExperienceAnalyticsModelScoreId{
				UserExperienceAnalyticsModelScoresId: "uSeReXpErIeNcEaNaLyTiCsMoDeLsCoReSiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsMoDeLsCoReS/uSeReXpErIeNcEaNaLyTiCsMoDeLsCoReSiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsModelScoreIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsModelScoresId != v.Expected.UserExperienceAnalyticsModelScoresId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsModelScoresId", v.Expected.UserExperienceAnalyticsModelScoresId, actual.UserExperienceAnalyticsModelScoresId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsModelScoreId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsModelScoreId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsModelScoreId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
