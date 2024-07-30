package userexperienceanalyticsanomaly

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAnomalyId{}

func TestNewDeviceManagementUserExperienceAnalyticsAnomalyID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsAnomalyID("userExperienceAnalyticsAnomalyIdValue")

	if id.UserExperienceAnalyticsAnomalyId != "userExperienceAnalyticsAnomalyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsAnomalyId'", id.UserExperienceAnalyticsAnomalyId, "userExperienceAnalyticsAnomalyIdValue")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsAnomalyID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsAnomalyID("userExperienceAnalyticsAnomalyIdValue").ID()
	expected := "/deviceManagement/userExperienceAnalyticsAnomaly/userExperienceAnalyticsAnomalyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAnomalyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAnomalyId
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
			Input: "/deviceManagement/userExperienceAnalyticsAnomaly",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAnomaly/userExperienceAnalyticsAnomalyIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsAnomalyId{
				UserExperienceAnalyticsAnomalyId: "userExperienceAnalyticsAnomalyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAnomaly/userExperienceAnalyticsAnomalyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAnomalyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAnomalyId != v.Expected.UserExperienceAnalyticsAnomalyId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAnomalyId", v.Expected.UserExperienceAnalyticsAnomalyId, actual.UserExperienceAnalyticsAnomalyId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAnomalyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAnomalyId
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
			Input: "/deviceManagement/userExperienceAnalyticsAnomaly",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsAnOmAlY",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAnomaly/userExperienceAnalyticsAnomalyIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsAnomalyId{
				UserExperienceAnalyticsAnomalyId: "userExperienceAnalyticsAnomalyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAnomaly/userExperienceAnalyticsAnomalyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsAnOmAlY/uSeReXpErIeNcEaNaLyTiCsAnOmAlYiDvAlUe",
			Expected: &DeviceManagementUserExperienceAnalyticsAnomalyId{
				UserExperienceAnalyticsAnomalyId: "uSeReXpErIeNcEaNaLyTiCsAnOmAlYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsAnOmAlY/uSeReXpErIeNcEaNaLyTiCsAnOmAlYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAnomalyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAnomalyId != v.Expected.UserExperienceAnalyticsAnomalyId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAnomalyId", v.Expected.UserExperienceAnalyticsAnomalyId, actual.UserExperienceAnalyticsAnomalyId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsAnomalyId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsAnomalyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsAnomalyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
