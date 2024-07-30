package userexperienceanalyticsdevicestartupprocess

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId{}

func TestNewDeviceManagementUserExperienceAnalyticsDeviceStartupProcessID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsDeviceStartupProcessID("userExperienceAnalyticsDeviceStartupProcessIdValue")

	if id.UserExperienceAnalyticsDeviceStartupProcessId != "userExperienceAnalyticsDeviceStartupProcessIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsDeviceStartupProcessId'", id.UserExperienceAnalyticsDeviceStartupProcessId, "userExperienceAnalyticsDeviceStartupProcessIdValue")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsDeviceStartupProcessID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsDeviceStartupProcessID("userExperienceAnalyticsDeviceStartupProcessIdValue").ID()
	expected := "/deviceManagement/userExperienceAnalyticsDeviceStartupProcesses/userExperienceAnalyticsDeviceStartupProcessIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsDeviceStartupProcessID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId
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
			Input: "/deviceManagement/userExperienceAnalyticsDeviceStartupProcesses",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsDeviceStartupProcesses/userExperienceAnalyticsDeviceStartupProcessIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId{
				UserExperienceAnalyticsDeviceStartupProcessId: "userExperienceAnalyticsDeviceStartupProcessIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsDeviceStartupProcesses/userExperienceAnalyticsDeviceStartupProcessIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsDeviceStartupProcessID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsDeviceStartupProcessId != v.Expected.UserExperienceAnalyticsDeviceStartupProcessId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsDeviceStartupProcessId", v.Expected.UserExperienceAnalyticsDeviceStartupProcessId, actual.UserExperienceAnalyticsDeviceStartupProcessId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsDeviceStartupProcessIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId
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
			Input: "/deviceManagement/userExperienceAnalyticsDeviceStartupProcesses",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeStArTuPpRoCeSsEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsDeviceStartupProcesses/userExperienceAnalyticsDeviceStartupProcessIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId{
				UserExperienceAnalyticsDeviceStartupProcessId: "userExperienceAnalyticsDeviceStartupProcessIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsDeviceStartupProcesses/userExperienceAnalyticsDeviceStartupProcessIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeStArTuPpRoCeSsEs/uSeReXpErIeNcEaNaLyTiCsDeViCeStArTuPpRoCeSsIdVaLuE",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId{
				UserExperienceAnalyticsDeviceStartupProcessId: "uSeReXpErIeNcEaNaLyTiCsDeViCeStArTuPpRoCeSsIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeStArTuPpRoCeSsEs/uSeReXpErIeNcEaNaLyTiCsDeViCeStArTuPpRoCeSsIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsDeviceStartupProcessIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsDeviceStartupProcessId != v.Expected.UserExperienceAnalyticsDeviceStartupProcessId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsDeviceStartupProcessId", v.Expected.UserExperienceAnalyticsDeviceStartupProcessId, actual.UserExperienceAnalyticsDeviceStartupProcessId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsDeviceStartupProcessId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
