package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId{}

func TestNewDeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceID("userExperienceAnalyticsNotAutopilotReadyDeviceId")

	if id.UserExperienceAnalyticsNotAutopilotReadyDeviceId != "userExperienceAnalyticsNotAutopilotReadyDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsNotAutopilotReadyDeviceId'", id.UserExperienceAnalyticsNotAutopilotReadyDeviceId, "userExperienceAnalyticsNotAutopilotReadyDeviceId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceID("userExperienceAnalyticsNotAutopilotReadyDeviceId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsNotAutopilotReadyDevice/userExperienceAnalyticsNotAutopilotReadyDeviceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId
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
			Input: "/deviceManagement/userExperienceAnalyticsNotAutopilotReadyDevice",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsNotAutopilotReadyDevice/userExperienceAnalyticsNotAutopilotReadyDeviceId",
			Expected: &DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId{
				UserExperienceAnalyticsNotAutopilotReadyDeviceId: "userExperienceAnalyticsNotAutopilotReadyDeviceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsNotAutopilotReadyDevice/userExperienceAnalyticsNotAutopilotReadyDeviceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsNotAutopilotReadyDeviceId != v.Expected.UserExperienceAnalyticsNotAutopilotReadyDeviceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsNotAutopilotReadyDeviceId", v.Expected.UserExperienceAnalyticsNotAutopilotReadyDeviceId, actual.UserExperienceAnalyticsNotAutopilotReadyDeviceId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId
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
			Input: "/deviceManagement/userExperienceAnalyticsNotAutopilotReadyDevice",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsNoTaUtOpIlOtReAdYdEvIcE",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsNotAutopilotReadyDevice/userExperienceAnalyticsNotAutopilotReadyDeviceId",
			Expected: &DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId{
				UserExperienceAnalyticsNotAutopilotReadyDeviceId: "userExperienceAnalyticsNotAutopilotReadyDeviceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsNotAutopilotReadyDevice/userExperienceAnalyticsNotAutopilotReadyDeviceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsNoTaUtOpIlOtReAdYdEvIcE/uSeReXpErIeNcEaNaLyTiCsNoTaUtOpIlOtReAdYdEvIcEiD",
			Expected: &DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId{
				UserExperienceAnalyticsNotAutopilotReadyDeviceId: "uSeReXpErIeNcEaNaLyTiCsNoTaUtOpIlOtReAdYdEvIcEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsNoTaUtOpIlOtReAdYdEvIcE/uSeReXpErIeNcEaNaLyTiCsNoTaUtOpIlOtReAdYdEvIcEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsNotAutopilotReadyDeviceId != v.Expected.UserExperienceAnalyticsNotAutopilotReadyDeviceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsNotAutopilotReadyDeviceId", v.Expected.UserExperienceAnalyticsNotAutopilotReadyDeviceId, actual.UserExperienceAnalyticsNotAutopilotReadyDeviceId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
