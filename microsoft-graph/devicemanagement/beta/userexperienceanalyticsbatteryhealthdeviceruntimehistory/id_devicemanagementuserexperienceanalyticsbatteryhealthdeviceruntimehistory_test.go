package userexperienceanalyticsbatteryhealthdeviceruntimehistory

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId{}

func TestNewDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryID("userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryIdValue")

	if id.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId != "userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId'", id.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId, "userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryIdValue")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryID("userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryIdValue").ID()
	expected := "/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory/userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId
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
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory/userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId{
				UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId: "userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory/userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId != v.Expected.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId", v.Expected.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId, actual.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId
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
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHdEvIcErUnTiMeHiStOrY",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory/userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId{
				UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId: "userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory/userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHdEvIcErUnTiMeHiStOrY/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHdEvIcErUnTiMeHiStOrYiDvAlUe",
			Expected: &DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId{
				UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId: "uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHdEvIcErUnTiMeHiStOrYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHdEvIcErUnTiMeHiStOrY/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHdEvIcErUnTiMeHiStOrYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId != v.Expected.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId", v.Expected.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId, actual.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
