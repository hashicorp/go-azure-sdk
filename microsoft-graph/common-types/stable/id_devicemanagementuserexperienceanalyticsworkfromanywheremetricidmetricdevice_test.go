package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId{}

func TestNewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceID("userExperienceAnalyticsWorkFromAnywhereMetricId", "userExperienceAnalyticsWorkFromAnywhereDeviceId")

	if id.UserExperienceAnalyticsWorkFromAnywhereMetricId != "userExperienceAnalyticsWorkFromAnywhereMetricId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsWorkFromAnywhereMetricId'", id.UserExperienceAnalyticsWorkFromAnywhereMetricId, "userExperienceAnalyticsWorkFromAnywhereMetricId")
	}

	if id.UserExperienceAnalyticsWorkFromAnywhereDeviceId != "userExperienceAnalyticsWorkFromAnywhereDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsWorkFromAnywhereDeviceId'", id.UserExperienceAnalyticsWorkFromAnywhereDeviceId, "userExperienceAnalyticsWorkFromAnywhereDeviceId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceID("userExperienceAnalyticsWorkFromAnywhereMetricId", "userExperienceAnalyticsWorkFromAnywhereDeviceId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricId/metricDevices/userExperienceAnalyticsWorkFromAnywhereDeviceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId
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
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricId/metricDevices",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricId/metricDevices/userExperienceAnalyticsWorkFromAnywhereDeviceId",
			Expected: &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId{
				UserExperienceAnalyticsWorkFromAnywhereMetricId: "userExperienceAnalyticsWorkFromAnywhereMetricId",
				UserExperienceAnalyticsWorkFromAnywhereDeviceId: "userExperienceAnalyticsWorkFromAnywhereDeviceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricId/metricDevices/userExperienceAnalyticsWorkFromAnywhereDeviceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsWorkFromAnywhereMetricId != v.Expected.UserExperienceAnalyticsWorkFromAnywhereMetricId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsWorkFromAnywhereMetricId", v.Expected.UserExperienceAnalyticsWorkFromAnywhereMetricId, actual.UserExperienceAnalyticsWorkFromAnywhereMetricId)
		}

		if actual.UserExperienceAnalyticsWorkFromAnywhereDeviceId != v.Expected.UserExperienceAnalyticsWorkFromAnywhereDeviceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsWorkFromAnywhereDeviceId", v.Expected.UserExperienceAnalyticsWorkFromAnywhereDeviceId, actual.UserExperienceAnalyticsWorkFromAnywhereDeviceId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId
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
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcS/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricId/metricDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcS/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcId/mEtRiCdEvIcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricId/metricDevices/userExperienceAnalyticsWorkFromAnywhereDeviceId",
			Expected: &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId{
				UserExperienceAnalyticsWorkFromAnywhereMetricId: "userExperienceAnalyticsWorkFromAnywhereMetricId",
				UserExperienceAnalyticsWorkFromAnywhereDeviceId: "userExperienceAnalyticsWorkFromAnywhereDeviceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricId/metricDevices/userExperienceAnalyticsWorkFromAnywhereDeviceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcS/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcId/mEtRiCdEvIcEs/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReDeViCeId",
			Expected: &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId{
				UserExperienceAnalyticsWorkFromAnywhereMetricId: "uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcId",
				UserExperienceAnalyticsWorkFromAnywhereDeviceId: "uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReDeViCeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcS/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcId/mEtRiCdEvIcEs/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReDeViCeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsWorkFromAnywhereMetricId != v.Expected.UserExperienceAnalyticsWorkFromAnywhereMetricId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsWorkFromAnywhereMetricId", v.Expected.UserExperienceAnalyticsWorkFromAnywhereMetricId, actual.UserExperienceAnalyticsWorkFromAnywhereMetricId)
		}

		if actual.UserExperienceAnalyticsWorkFromAnywhereDeviceId != v.Expected.UserExperienceAnalyticsWorkFromAnywhereDeviceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsWorkFromAnywhereDeviceId", v.Expected.UserExperienceAnalyticsWorkFromAnywhereDeviceId, actual.UserExperienceAnalyticsWorkFromAnywhereDeviceId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
