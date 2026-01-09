package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId{}

func TestNewDeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewID("userExperienceAnalyticsAnomalyCorrelationGroupOverviewId")

	if id.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId != "userExperienceAnalyticsAnomalyCorrelationGroupOverviewId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId'", id.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId, "userExperienceAnalyticsAnomalyCorrelationGroupOverviewId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewID("userExperienceAnalyticsAnomalyCorrelationGroupOverviewId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsAnomalyCorrelationGroupOverview/userExperienceAnalyticsAnomalyCorrelationGroupOverviewId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId
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
			Input: "/deviceManagement/userExperienceAnalyticsAnomalyCorrelationGroupOverview",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAnomalyCorrelationGroupOverview/userExperienceAnalyticsAnomalyCorrelationGroupOverviewId",
			Expected: &DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId{
				UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId: "userExperienceAnalyticsAnomalyCorrelationGroupOverviewId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAnomalyCorrelationGroupOverview/userExperienceAnalyticsAnomalyCorrelationGroupOverviewId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId != v.Expected.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId", v.Expected.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId, actual.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId
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
			Input: "/deviceManagement/userExperienceAnalyticsAnomalyCorrelationGroupOverview",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsAnOmAlYcOrReLaTiOnGrOuPoVeRvIeW",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAnomalyCorrelationGroupOverview/userExperienceAnalyticsAnomalyCorrelationGroupOverviewId",
			Expected: &DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId{
				UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId: "userExperienceAnalyticsAnomalyCorrelationGroupOverviewId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAnomalyCorrelationGroupOverview/userExperienceAnalyticsAnomalyCorrelationGroupOverviewId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsAnOmAlYcOrReLaTiOnGrOuPoVeRvIeW/uSeReXpErIeNcEaNaLyTiCsAnOmAlYcOrReLaTiOnGrOuPoVeRvIeWiD",
			Expected: &DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId{
				UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId: "uSeReXpErIeNcEaNaLyTiCsAnOmAlYcOrReLaTiOnGrOuPoVeRvIeWiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsAnOmAlYcOrReLaTiOnGrOuPoVeRvIeW/uSeReXpErIeNcEaNaLyTiCsAnOmAlYcOrReLaTiOnGrOuPoVeRvIeWiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId != v.Expected.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId", v.Expected.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId, actual.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
