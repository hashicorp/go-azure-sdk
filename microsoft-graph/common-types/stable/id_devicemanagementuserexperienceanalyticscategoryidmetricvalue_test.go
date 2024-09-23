package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId{}

func TestNewDeviceManagementUserExperienceAnalyticsCategoryIdMetricValueID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsCategoryIdMetricValueID("userExperienceAnalyticsCategoryId", "userExperienceAnalyticsMetricId")

	if id.UserExperienceAnalyticsCategoryId != "userExperienceAnalyticsCategoryId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsCategoryId'", id.UserExperienceAnalyticsCategoryId, "userExperienceAnalyticsCategoryId")
	}

	if id.UserExperienceAnalyticsMetricId != "userExperienceAnalyticsMetricId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsMetricId'", id.UserExperienceAnalyticsMetricId, "userExperienceAnalyticsMetricId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsCategoryIdMetricValueID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsCategoryIdMetricValueID("userExperienceAnalyticsCategoryId", "userExperienceAnalyticsMetricId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsCategories/userExperienceAnalyticsCategoryId/metricValues/userExperienceAnalyticsMetricId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsCategoryIdMetricValueID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId
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
			Input: "/deviceManagement/userExperienceAnalyticsCategories",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/userExperienceAnalyticsCategories/userExperienceAnalyticsCategoryId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/userExperienceAnalyticsCategories/userExperienceAnalyticsCategoryId/metricValues",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsCategories/userExperienceAnalyticsCategoryId/metricValues/userExperienceAnalyticsMetricId",
			Expected: &DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId{
				UserExperienceAnalyticsCategoryId: "userExperienceAnalyticsCategoryId",
				UserExperienceAnalyticsMetricId:   "userExperienceAnalyticsMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsCategories/userExperienceAnalyticsCategoryId/metricValues/userExperienceAnalyticsMetricId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsCategoryIdMetricValueID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsCategoryId != v.Expected.UserExperienceAnalyticsCategoryId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsCategoryId", v.Expected.UserExperienceAnalyticsCategoryId, actual.UserExperienceAnalyticsCategoryId)
		}

		if actual.UserExperienceAnalyticsMetricId != v.Expected.UserExperienceAnalyticsMetricId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsMetricId", v.Expected.UserExperienceAnalyticsMetricId, actual.UserExperienceAnalyticsMetricId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsCategoryIdMetricValueIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId
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
			Input: "/deviceManagement/userExperienceAnalyticsCategories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsCaTeGoRiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/userExperienceAnalyticsCategories/userExperienceAnalyticsCategoryId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsCaTeGoRiEs/uSeReXpErIeNcEaNaLyTiCsCaTeGoRyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/userExperienceAnalyticsCategories/userExperienceAnalyticsCategoryId/metricValues",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsCaTeGoRiEs/uSeReXpErIeNcEaNaLyTiCsCaTeGoRyId/mEtRiCvAlUeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsCategories/userExperienceAnalyticsCategoryId/metricValues/userExperienceAnalyticsMetricId",
			Expected: &DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId{
				UserExperienceAnalyticsCategoryId: "userExperienceAnalyticsCategoryId",
				UserExperienceAnalyticsMetricId:   "userExperienceAnalyticsMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsCategories/userExperienceAnalyticsCategoryId/metricValues/userExperienceAnalyticsMetricId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsCaTeGoRiEs/uSeReXpErIeNcEaNaLyTiCsCaTeGoRyId/mEtRiCvAlUeS/uSeReXpErIeNcEaNaLyTiCsMeTrIcId",
			Expected: &DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId{
				UserExperienceAnalyticsCategoryId: "uSeReXpErIeNcEaNaLyTiCsCaTeGoRyId",
				UserExperienceAnalyticsMetricId:   "uSeReXpErIeNcEaNaLyTiCsMeTrIcId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsCaTeGoRiEs/uSeReXpErIeNcEaNaLyTiCsCaTeGoRyId/mEtRiCvAlUeS/uSeReXpErIeNcEaNaLyTiCsMeTrIcId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsCategoryIdMetricValueIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsCategoryId != v.Expected.UserExperienceAnalyticsCategoryId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsCategoryId", v.Expected.UserExperienceAnalyticsCategoryId, actual.UserExperienceAnalyticsCategoryId)
		}

		if actual.UserExperienceAnalyticsMetricId != v.Expected.UserExperienceAnalyticsMetricId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsMetricId", v.Expected.UserExperienceAnalyticsMetricId, actual.UserExperienceAnalyticsMetricId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
