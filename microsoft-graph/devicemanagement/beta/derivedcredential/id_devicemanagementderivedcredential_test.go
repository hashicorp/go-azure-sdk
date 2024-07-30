package derivedcredential

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDerivedCredentialId{}

func TestNewDeviceManagementDerivedCredentialID(t *testing.T) {
	id := NewDeviceManagementDerivedCredentialID("deviceManagementDerivedCredentialSettingsIdValue")

	if id.DeviceManagementDerivedCredentialSettingsId != "deviceManagementDerivedCredentialSettingsIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementDerivedCredentialSettingsId'", id.DeviceManagementDerivedCredentialSettingsId, "deviceManagementDerivedCredentialSettingsIdValue")
	}
}

func TestFormatDeviceManagementDerivedCredentialID(t *testing.T) {
	actual := NewDeviceManagementDerivedCredentialID("deviceManagementDerivedCredentialSettingsIdValue").ID()
	expected := "/deviceManagement/derivedCredentials/deviceManagementDerivedCredentialSettingsIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDerivedCredentialID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDerivedCredentialId
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
			Input: "/deviceManagement/derivedCredentials",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/derivedCredentials/deviceManagementDerivedCredentialSettingsIdValue",
			Expected: &DeviceManagementDerivedCredentialId{
				DeviceManagementDerivedCredentialSettingsId: "deviceManagementDerivedCredentialSettingsIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/derivedCredentials/deviceManagementDerivedCredentialSettingsIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDerivedCredentialID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementDerivedCredentialSettingsId != v.Expected.DeviceManagementDerivedCredentialSettingsId {
			t.Fatalf("Expected %q but got %q for DeviceManagementDerivedCredentialSettingsId", v.Expected.DeviceManagementDerivedCredentialSettingsId, actual.DeviceManagementDerivedCredentialSettingsId)
		}

	}
}

func TestParseDeviceManagementDerivedCredentialIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDerivedCredentialId
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
			Input: "/deviceManagement/derivedCredentials",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dErIvEdCrEdEnTiAlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/derivedCredentials/deviceManagementDerivedCredentialSettingsIdValue",
			Expected: &DeviceManagementDerivedCredentialId{
				DeviceManagementDerivedCredentialSettingsId: "deviceManagementDerivedCredentialSettingsIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/derivedCredentials/deviceManagementDerivedCredentialSettingsIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dErIvEdCrEdEnTiAlS/dEvIcEmAnAgEmEnTdErIvEdCrEdEnTiAlSeTtInGsIdVaLuE",
			Expected: &DeviceManagementDerivedCredentialId{
				DeviceManagementDerivedCredentialSettingsId: "dEvIcEmAnAgEmEnTdErIvEdCrEdEnTiAlSeTtInGsIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dErIvEdCrEdEnTiAlS/dEvIcEmAnAgEmEnTdErIvEdCrEdEnTiAlSeTtInGsIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDerivedCredentialIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementDerivedCredentialSettingsId != v.Expected.DeviceManagementDerivedCredentialSettingsId {
			t.Fatalf("Expected %q but got %q for DeviceManagementDerivedCredentialSettingsId", v.Expected.DeviceManagementDerivedCredentialSettingsId, actual.DeviceManagementDerivedCredentialSettingsId)
		}

	}
}

func TestSegmentsForDeviceManagementDerivedCredentialId(t *testing.T) {
	segments := DeviceManagementDerivedCredentialId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDerivedCredentialId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
