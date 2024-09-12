package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDepOnboardingSettingId{}

func TestNewDeviceManagementDepOnboardingSettingID(t *testing.T) {
	id := NewDeviceManagementDepOnboardingSettingID("depOnboardingSettingIdValue")

	if id.DepOnboardingSettingId != "depOnboardingSettingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DepOnboardingSettingId'", id.DepOnboardingSettingId, "depOnboardingSettingIdValue")
	}
}

func TestFormatDeviceManagementDepOnboardingSettingID(t *testing.T) {
	actual := NewDeviceManagementDepOnboardingSettingID("depOnboardingSettingIdValue").ID()
	expected := "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDepOnboardingSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDepOnboardingSettingId
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
			Input: "/deviceManagement/depOnboardingSettings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue",
			Expected: &DeviceManagementDepOnboardingSettingId{
				DepOnboardingSettingId: "depOnboardingSettingIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDepOnboardingSettingID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DepOnboardingSettingId != v.Expected.DepOnboardingSettingId {
			t.Fatalf("Expected %q but got %q for DepOnboardingSettingId", v.Expected.DepOnboardingSettingId, actual.DepOnboardingSettingId)
		}

	}
}

func TestParseDeviceManagementDepOnboardingSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDepOnboardingSettingId
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
			Input: "/deviceManagement/depOnboardingSettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue",
			Expected: &DeviceManagementDepOnboardingSettingId{
				DepOnboardingSettingId: "depOnboardingSettingIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs/dEpOnBoArDiNgSeTtInGiDvAlUe",
			Expected: &DeviceManagementDepOnboardingSettingId{
				DepOnboardingSettingId: "dEpOnBoArDiNgSeTtInGiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs/dEpOnBoArDiNgSeTtInGiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDepOnboardingSettingIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DepOnboardingSettingId != v.Expected.DepOnboardingSettingId {
			t.Fatalf("Expected %q but got %q for DepOnboardingSettingId", v.Expected.DepOnboardingSettingId, actual.DepOnboardingSettingId)
		}

	}
}

func TestSegmentsForDeviceManagementDepOnboardingSettingId(t *testing.T) {
	segments := DeviceManagementDepOnboardingSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDepOnboardingSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
