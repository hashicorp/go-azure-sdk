package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementChromeOSOnboardingSettingId{}

func TestNewDeviceManagementChromeOSOnboardingSettingID(t *testing.T) {
	id := NewDeviceManagementChromeOSOnboardingSettingID("chromeOSOnboardingSettingsIdValue")

	if id.ChromeOSOnboardingSettingsId != "chromeOSOnboardingSettingsIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChromeOSOnboardingSettingsId'", id.ChromeOSOnboardingSettingsId, "chromeOSOnboardingSettingsIdValue")
	}
}

func TestFormatDeviceManagementChromeOSOnboardingSettingID(t *testing.T) {
	actual := NewDeviceManagementChromeOSOnboardingSettingID("chromeOSOnboardingSettingsIdValue").ID()
	expected := "/deviceManagement/chromeOSOnboardingSettings/chromeOSOnboardingSettingsIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementChromeOSOnboardingSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementChromeOSOnboardingSettingId
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
			Input: "/deviceManagement/chromeOSOnboardingSettings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/chromeOSOnboardingSettings/chromeOSOnboardingSettingsIdValue",
			Expected: &DeviceManagementChromeOSOnboardingSettingId{
				ChromeOSOnboardingSettingsId: "chromeOSOnboardingSettingsIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/chromeOSOnboardingSettings/chromeOSOnboardingSettingsIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementChromeOSOnboardingSettingID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ChromeOSOnboardingSettingsId != v.Expected.ChromeOSOnboardingSettingsId {
			t.Fatalf("Expected %q but got %q for ChromeOSOnboardingSettingsId", v.Expected.ChromeOSOnboardingSettingsId, actual.ChromeOSOnboardingSettingsId)
		}

	}
}

func TestParseDeviceManagementChromeOSOnboardingSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementChromeOSOnboardingSettingId
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
			Input: "/deviceManagement/chromeOSOnboardingSettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cHrOmEoSoNbOaRdInGsEtTiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/chromeOSOnboardingSettings/chromeOSOnboardingSettingsIdValue",
			Expected: &DeviceManagementChromeOSOnboardingSettingId{
				ChromeOSOnboardingSettingsId: "chromeOSOnboardingSettingsIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/chromeOSOnboardingSettings/chromeOSOnboardingSettingsIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cHrOmEoSoNbOaRdInGsEtTiNgS/cHrOmEoSoNbOaRdInGsEtTiNgSiDvAlUe",
			Expected: &DeviceManagementChromeOSOnboardingSettingId{
				ChromeOSOnboardingSettingsId: "cHrOmEoSoNbOaRdInGsEtTiNgSiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cHrOmEoSoNbOaRdInGsEtTiNgS/cHrOmEoSoNbOaRdInGsEtTiNgSiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementChromeOSOnboardingSettingIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ChromeOSOnboardingSettingsId != v.Expected.ChromeOSOnboardingSettingsId {
			t.Fatalf("Expected %q but got %q for ChromeOSOnboardingSettingsId", v.Expected.ChromeOSOnboardingSettingsId, actual.ChromeOSOnboardingSettingsId)
		}

	}
}

func TestSegmentsForDeviceManagementChromeOSOnboardingSettingId(t *testing.T) {
	segments := DeviceManagementChromeOSOnboardingSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementChromeOSOnboardingSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
