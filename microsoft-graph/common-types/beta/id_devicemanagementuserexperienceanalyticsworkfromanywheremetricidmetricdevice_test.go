package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId{}

func TestNewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceID("userExperienceAnalyticsWorkFromAnywhereMetricIdValue", "userExperienceAnalyticsWorkFromAnywhereDeviceIdValue")

	if id.UserExperienceAnalyticsWorkFromAnywhereMetricId != "userExperienceAnalyticsWorkFromAnywhereMetricIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsWorkFromAnywhereMetricId'", id.UserExperienceAnalyticsWorkFromAnywhereMetricId, "userExperienceAnalyticsWorkFromAnywhereMetricIdValue")
	}

	if id.UserExperienceAnalyticsWorkFromAnywhereDeviceId != "userExperienceAnalyticsWorkFromAnywhereDeviceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsWorkFromAnywhereDeviceId'", id.UserExperienceAnalyticsWorkFromAnywhereDeviceId, "userExperienceAnalyticsWorkFromAnywhereDeviceIdValue")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceID("userExperienceAnalyticsWorkFromAnywhereMetricIdValue", "userExperienceAnalyticsWorkFromAnywhereDeviceIdValue").ID()
	expected := "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricIdValue/metricDevices/userExperienceAnalyticsWorkFromAnywhereDeviceIdValue"
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
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricIdValue/metricDevices",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricIdValue/metricDevices/userExperienceAnalyticsWorkFromAnywhereDeviceIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId{
				UserExperienceAnalyticsWorkFromAnywhereMetricId: "userExperienceAnalyticsWorkFromAnywhereMetricIdValue",
				UserExperienceAnalyticsWorkFromAnywhereDeviceId: "userExperienceAnalyticsWorkFromAnywhereDeviceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricIdValue/metricDevices/userExperienceAnalyticsWorkFromAnywhereDeviceIdValue/extra",
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
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcS/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricIdValue/metricDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcS/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcIdVaLuE/mEtRiCdEvIcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricIdValue/metricDevices/userExperienceAnalyticsWorkFromAnywhereDeviceIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId{
				UserExperienceAnalyticsWorkFromAnywhereMetricId: "userExperienceAnalyticsWorkFromAnywhereMetricIdValue",
				UserExperienceAnalyticsWorkFromAnywhereDeviceId: "userExperienceAnalyticsWorkFromAnywhereDeviceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricIdValue/metricDevices/userExperienceAnalyticsWorkFromAnywhereDeviceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcS/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcIdVaLuE/mEtRiCdEvIcEs/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReDeViCeIdVaLuE",
			Expected: &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId{
				UserExperienceAnalyticsWorkFromAnywhereMetricId: "uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcIdVaLuE",
				UserExperienceAnalyticsWorkFromAnywhereDeviceId: "uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReDeViCeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcS/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcIdVaLuE/mEtRiCdEvIcEs/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReDeViCeIdVaLuE/extra",
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
