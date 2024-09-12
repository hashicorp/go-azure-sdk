package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId{}

func TestNewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricID("userExperienceAnalyticsWorkFromAnywhereMetricIdValue")

	if id.UserExperienceAnalyticsWorkFromAnywhereMetricId != "userExperienceAnalyticsWorkFromAnywhereMetricIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsWorkFromAnywhereMetricId'", id.UserExperienceAnalyticsWorkFromAnywhereMetricId, "userExperienceAnalyticsWorkFromAnywhereMetricIdValue")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricID("userExperienceAnalyticsWorkFromAnywhereMetricIdValue").ID()
	expected := "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId
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
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId{
				UserExperienceAnalyticsWorkFromAnywhereMetricId: "userExperienceAnalyticsWorkFromAnywhereMetricIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricID(v.Input)
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

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId
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
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId{
				UserExperienceAnalyticsWorkFromAnywhereMetricId: "userExperienceAnalyticsWorkFromAnywhereMetricIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/userExperienceAnalyticsWorkFromAnywhereMetricIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcS/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcIdVaLuE",
			Expected: &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId{
				UserExperienceAnalyticsWorkFromAnywhereMetricId: "uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcS/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMeTrIcIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
