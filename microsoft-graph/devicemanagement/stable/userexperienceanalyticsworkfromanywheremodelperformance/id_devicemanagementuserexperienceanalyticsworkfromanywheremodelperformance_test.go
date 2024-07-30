package userexperienceanalyticsworkfromanywheremodelperformance

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId{}

func TestNewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceID("userExperienceAnalyticsWorkFromAnywhereModelPerformanceIdValue")

	if id.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId != "userExperienceAnalyticsWorkFromAnywhereModelPerformanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId'", id.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId, "userExperienceAnalyticsWorkFromAnywhereModelPerformanceIdValue")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceID("userExperienceAnalyticsWorkFromAnywhereModelPerformanceIdValue").ID()
	expected := "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereModelPerformance/userExperienceAnalyticsWorkFromAnywhereModelPerformanceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereModelPerformance",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereModelPerformance/userExperienceAnalyticsWorkFromAnywhereModelPerformanceIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId{
				UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId: "userExperienceAnalyticsWorkFromAnywhereModelPerformanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereModelPerformance/userExperienceAnalyticsWorkFromAnywhereModelPerformanceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId != v.Expected.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId", v.Expected.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId, actual.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId
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
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereModelPerformance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMoDeLpErFoRmAnCe",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereModelPerformance/userExperienceAnalyticsWorkFromAnywhereModelPerformanceIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId{
				UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId: "userExperienceAnalyticsWorkFromAnywhereModelPerformanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereModelPerformance/userExperienceAnalyticsWorkFromAnywhereModelPerformanceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMoDeLpErFoRmAnCe/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMoDeLpErFoRmAnCeIdVaLuE",
			Expected: &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId{
				UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId: "uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMoDeLpErFoRmAnCeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMoDeLpErFoRmAnCe/uSeReXpErIeNcEaNaLyTiCsWoRkFrOmAnYwHeReMoDeLpErFoRmAnCeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId != v.Expected.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId", v.Expected.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId, actual.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
