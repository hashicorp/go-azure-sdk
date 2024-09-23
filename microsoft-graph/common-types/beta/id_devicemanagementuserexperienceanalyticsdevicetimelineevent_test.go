package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId{}

func TestNewDeviceManagementUserExperienceAnalyticsDeviceTimelineEventID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsDeviceTimelineEventID("userExperienceAnalyticsDeviceTimelineEventId")

	if id.UserExperienceAnalyticsDeviceTimelineEventId != "userExperienceAnalyticsDeviceTimelineEventId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsDeviceTimelineEventId'", id.UserExperienceAnalyticsDeviceTimelineEventId, "userExperienceAnalyticsDeviceTimelineEventId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsDeviceTimelineEventID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsDeviceTimelineEventID("userExperienceAnalyticsDeviceTimelineEventId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsDeviceTimelineEvent/userExperienceAnalyticsDeviceTimelineEventId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsDeviceTimelineEventID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId
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
			Input: "/deviceManagement/userExperienceAnalyticsDeviceTimelineEvent",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsDeviceTimelineEvent/userExperienceAnalyticsDeviceTimelineEventId",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId{
				UserExperienceAnalyticsDeviceTimelineEventId: "userExperienceAnalyticsDeviceTimelineEventId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsDeviceTimelineEvent/userExperienceAnalyticsDeviceTimelineEventId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsDeviceTimelineEventID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsDeviceTimelineEventId != v.Expected.UserExperienceAnalyticsDeviceTimelineEventId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsDeviceTimelineEventId", v.Expected.UserExperienceAnalyticsDeviceTimelineEventId, actual.UserExperienceAnalyticsDeviceTimelineEventId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsDeviceTimelineEventIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId
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
			Input: "/deviceManagement/userExperienceAnalyticsDeviceTimelineEvent",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeTiMeLiNeEvEnT",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsDeviceTimelineEvent/userExperienceAnalyticsDeviceTimelineEventId",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId{
				UserExperienceAnalyticsDeviceTimelineEventId: "userExperienceAnalyticsDeviceTimelineEventId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsDeviceTimelineEvent/userExperienceAnalyticsDeviceTimelineEventId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeTiMeLiNeEvEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeTiMeLiNeEvEnTiD",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId{
				UserExperienceAnalyticsDeviceTimelineEventId: "uSeReXpErIeNcEaNaLyTiCsDeViCeTiMeLiNeEvEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeTiMeLiNeEvEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeTiMeLiNeEvEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsDeviceTimelineEventIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsDeviceTimelineEventId != v.Expected.UserExperienceAnalyticsDeviceTimelineEventId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsDeviceTimelineEventId", v.Expected.UserExperienceAnalyticsDeviceTimelineEventId, actual.UserExperienceAnalyticsDeviceTimelineEventId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsDeviceTimelineEventId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
