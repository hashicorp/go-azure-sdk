package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAnomalyDeviceId{}

func TestNewDeviceManagementUserExperienceAnalyticsAnomalyDeviceID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsAnomalyDeviceID("userExperienceAnalyticsAnomalyDeviceId")

	if id.UserExperienceAnalyticsAnomalyDeviceId != "userExperienceAnalyticsAnomalyDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsAnomalyDeviceId'", id.UserExperienceAnalyticsAnomalyDeviceId, "userExperienceAnalyticsAnomalyDeviceId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsAnomalyDeviceID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsAnomalyDeviceID("userExperienceAnalyticsAnomalyDeviceId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsAnomalyDevice/userExperienceAnalyticsAnomalyDeviceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAnomalyDeviceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAnomalyDeviceId
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
			Input: "/deviceManagement/userExperienceAnalyticsAnomalyDevice",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAnomalyDevice/userExperienceAnalyticsAnomalyDeviceId",
			Expected: &DeviceManagementUserExperienceAnalyticsAnomalyDeviceId{
				UserExperienceAnalyticsAnomalyDeviceId: "userExperienceAnalyticsAnomalyDeviceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAnomalyDevice/userExperienceAnalyticsAnomalyDeviceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAnomalyDeviceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAnomalyDeviceId != v.Expected.UserExperienceAnalyticsAnomalyDeviceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAnomalyDeviceId", v.Expected.UserExperienceAnalyticsAnomalyDeviceId, actual.UserExperienceAnalyticsAnomalyDeviceId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAnomalyDeviceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAnomalyDeviceId
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
			Input: "/deviceManagement/userExperienceAnalyticsAnomalyDevice",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsAnOmAlYdEvIcE",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAnomalyDevice/userExperienceAnalyticsAnomalyDeviceId",
			Expected: &DeviceManagementUserExperienceAnalyticsAnomalyDeviceId{
				UserExperienceAnalyticsAnomalyDeviceId: "userExperienceAnalyticsAnomalyDeviceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAnomalyDevice/userExperienceAnalyticsAnomalyDeviceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsAnOmAlYdEvIcE/uSeReXpErIeNcEaNaLyTiCsAnOmAlYdEvIcEiD",
			Expected: &DeviceManagementUserExperienceAnalyticsAnomalyDeviceId{
				UserExperienceAnalyticsAnomalyDeviceId: "uSeReXpErIeNcEaNaLyTiCsAnOmAlYdEvIcEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsAnOmAlYdEvIcE/uSeReXpErIeNcEaNaLyTiCsAnOmAlYdEvIcEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAnomalyDeviceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAnomalyDeviceId != v.Expected.UserExperienceAnalyticsAnomalyDeviceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAnomalyDeviceId", v.Expected.UserExperienceAnalyticsAnomalyDeviceId, actual.UserExperienceAnalyticsAnomalyDeviceId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsAnomalyDeviceId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsAnomalyDeviceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsAnomalyDeviceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
