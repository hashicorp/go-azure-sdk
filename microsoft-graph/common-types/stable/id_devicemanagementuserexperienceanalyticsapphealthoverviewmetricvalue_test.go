package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId{}

func TestNewDeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueID("userExperienceAnalyticsMetricId")

	if id.UserExperienceAnalyticsMetricId != "userExperienceAnalyticsMetricId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsMetricId'", id.UserExperienceAnalyticsMetricId, "userExperienceAnalyticsMetricId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueID("userExperienceAnalyticsMetricId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsAppHealthOverview/metricValues/userExperienceAnalyticsMetricId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId
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
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthOverview",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthOverview/metricValues",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthOverview/metricValues/userExperienceAnalyticsMetricId",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId{
				UserExperienceAnalyticsMetricId: "userExperienceAnalyticsMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthOverview/metricValues/userExperienceAnalyticsMetricId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsMetricId != v.Expected.UserExperienceAnalyticsMetricId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsMetricId", v.Expected.UserExperienceAnalyticsMetricId, actual.UserExperienceAnalyticsMetricId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId
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
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthOverview",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHoVeRvIeW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthOverview/metricValues",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHoVeRvIeW/mEtRiCvAlUeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthOverview/metricValues/userExperienceAnalyticsMetricId",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId{
				UserExperienceAnalyticsMetricId: "userExperienceAnalyticsMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthOverview/metricValues/userExperienceAnalyticsMetricId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHoVeRvIeW/mEtRiCvAlUeS/uSeReXpErIeNcEaNaLyTiCsMeTrIcId",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId{
				UserExperienceAnalyticsMetricId: "uSeReXpErIeNcEaNaLyTiCsMeTrIcId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHoVeRvIeW/mEtRiCvAlUeS/uSeReXpErIeNcEaNaLyTiCsMeTrIcId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsMetricId != v.Expected.UserExperienceAnalyticsMetricId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsMetricId", v.Expected.UserExperienceAnalyticsMetricId, actual.UserExperienceAnalyticsMetricId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
