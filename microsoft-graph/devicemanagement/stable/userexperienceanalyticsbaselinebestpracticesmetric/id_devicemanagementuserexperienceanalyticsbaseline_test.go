package userexperienceanalyticsbaselinebestpracticesmetric

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsBaselineId{}

func TestNewDeviceManagementUserExperienceAnalyticsBaselineID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsBaselineID("userExperienceAnalyticsBaselineIdValue")

	if id.UserExperienceAnalyticsBaselineId != "userExperienceAnalyticsBaselineIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsBaselineId'", id.UserExperienceAnalyticsBaselineId, "userExperienceAnalyticsBaselineIdValue")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsBaselineID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsBaselineID("userExperienceAnalyticsBaselineIdValue").ID()
	expected := "/deviceManagement/userExperienceAnalyticsBaselines/userExperienceAnalyticsBaselineIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsBaselineID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsBaselineId
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
			Input: "/deviceManagement/userExperienceAnalyticsBaselines",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsBaselines/userExperienceAnalyticsBaselineIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsBaselineId{
				UserExperienceAnalyticsBaselineId: "userExperienceAnalyticsBaselineIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsBaselines/userExperienceAnalyticsBaselineIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsBaselineID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsBaselineId != v.Expected.UserExperienceAnalyticsBaselineId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsBaselineId", v.Expected.UserExperienceAnalyticsBaselineId, actual.UserExperienceAnalyticsBaselineId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsBaselineIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsBaselineId
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
			Input: "/deviceManagement/userExperienceAnalyticsBaselines",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaSeLiNeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsBaselines/userExperienceAnalyticsBaselineIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsBaselineId{
				UserExperienceAnalyticsBaselineId: "userExperienceAnalyticsBaselineIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsBaselines/userExperienceAnalyticsBaselineIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaSeLiNeS/uSeReXpErIeNcEaNaLyTiCsBaSeLiNeIdVaLuE",
			Expected: &DeviceManagementUserExperienceAnalyticsBaselineId{
				UserExperienceAnalyticsBaselineId: "uSeReXpErIeNcEaNaLyTiCsBaSeLiNeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaSeLiNeS/uSeReXpErIeNcEaNaLyTiCsBaSeLiNeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsBaselineIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsBaselineId != v.Expected.UserExperienceAnalyticsBaselineId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsBaselineId", v.Expected.UserExperienceAnalyticsBaselineId, actual.UserExperienceAnalyticsBaselineId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsBaselineId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsBaselineId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsBaselineId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
