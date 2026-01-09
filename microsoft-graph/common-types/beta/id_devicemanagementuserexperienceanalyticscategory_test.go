package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsCategoryId{}

func TestNewDeviceManagementUserExperienceAnalyticsCategoryID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsCategoryID("userExperienceAnalyticsCategoryId")

	if id.UserExperienceAnalyticsCategoryId != "userExperienceAnalyticsCategoryId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsCategoryId'", id.UserExperienceAnalyticsCategoryId, "userExperienceAnalyticsCategoryId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsCategoryID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsCategoryID("userExperienceAnalyticsCategoryId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsCategories/userExperienceAnalyticsCategoryId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsCategoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsCategoryId
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
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsCategories/userExperienceAnalyticsCategoryId",
			Expected: &DeviceManagementUserExperienceAnalyticsCategoryId{
				UserExperienceAnalyticsCategoryId: "userExperienceAnalyticsCategoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsCategories/userExperienceAnalyticsCategoryId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsCategoryID(v.Input)
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

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsCategoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsCategoryId
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
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsCategories/userExperienceAnalyticsCategoryId",
			Expected: &DeviceManagementUserExperienceAnalyticsCategoryId{
				UserExperienceAnalyticsCategoryId: "userExperienceAnalyticsCategoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsCategories/userExperienceAnalyticsCategoryId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsCaTeGoRiEs/uSeReXpErIeNcEaNaLyTiCsCaTeGoRyId",
			Expected: &DeviceManagementUserExperienceAnalyticsCategoryId{
				UserExperienceAnalyticsCategoryId: "uSeReXpErIeNcEaNaLyTiCsCaTeGoRyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsCaTeGoRiEs/uSeReXpErIeNcEaNaLyTiCsCaTeGoRyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsCategoryIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsCategoryId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsCategoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsCategoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
