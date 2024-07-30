package userexperienceanalyticscategory

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsCategoryId{}

func TestNewDeviceManagementUserExperienceAnalyticsCategoryID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsCategoryID("userExperienceAnalyticsCategoryIdValue")

	if id.UserExperienceAnalyticsCategoryId != "userExperienceAnalyticsCategoryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsCategoryId'", id.UserExperienceAnalyticsCategoryId, "userExperienceAnalyticsCategoryIdValue")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsCategoryID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsCategoryID("userExperienceAnalyticsCategoryIdValue").ID()
	expected := "/deviceManagement/userExperienceAnalyticsCategories/userExperienceAnalyticsCategoryIdValue"
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
			Input: "/deviceManagement/userExperienceAnalyticsCategories/userExperienceAnalyticsCategoryIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsCategoryId{
				UserExperienceAnalyticsCategoryId: "userExperienceAnalyticsCategoryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsCategories/userExperienceAnalyticsCategoryIdValue/extra",
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
			Input: "/deviceManagement/userExperienceAnalyticsCategories/userExperienceAnalyticsCategoryIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsCategoryId{
				UserExperienceAnalyticsCategoryId: "userExperienceAnalyticsCategoryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsCategories/userExperienceAnalyticsCategoryIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsCaTeGoRiEs/uSeReXpErIeNcEaNaLyTiCsCaTeGoRyIdVaLuE",
			Expected: &DeviceManagementUserExperienceAnalyticsCategoryId{
				UserExperienceAnalyticsCategoryId: "uSeReXpErIeNcEaNaLyTiCsCaTeGoRyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsCaTeGoRiEs/uSeReXpErIeNcEaNaLyTiCsCaTeGoRyIdVaLuE/extra",
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
