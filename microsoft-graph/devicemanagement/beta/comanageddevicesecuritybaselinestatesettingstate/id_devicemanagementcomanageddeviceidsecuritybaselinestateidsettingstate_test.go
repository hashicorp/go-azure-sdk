package comanageddevicesecuritybaselinestatesettingstate

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId{}

func TestNewDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateID(t *testing.T) {
	id := NewDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateID("managedDeviceIdValue", "securityBaselineStateIdValue", "securityBaselineSettingStateIdValue")

	if id.ManagedDeviceId != "managedDeviceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceIdValue")
	}

	if id.SecurityBaselineStateId != "securityBaselineStateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SecurityBaselineStateId'", id.SecurityBaselineStateId, "securityBaselineStateIdValue")
	}

	if id.SecurityBaselineSettingStateId != "securityBaselineSettingStateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SecurityBaselineSettingStateId'", id.SecurityBaselineSettingStateId, "securityBaselineSettingStateIdValue")
	}
}

func TestFormatDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateID(t *testing.T) {
	actual := NewDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateID("managedDeviceIdValue", "securityBaselineStateIdValue", "securityBaselineSettingStateIdValue").ID()
	expected := "/deviceManagement/comanagedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates/securityBaselineSettingStateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId
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
			Input: "/deviceManagement/comanagedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceIdValue/securityBaselineStates",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates/securityBaselineSettingStateIdValue",
			Expected: &DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId{
				ManagedDeviceId:                "managedDeviceIdValue",
				SecurityBaselineStateId:        "securityBaselineStateIdValue",
				SecurityBaselineSettingStateId: "securityBaselineSettingStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/comanagedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates/securityBaselineSettingStateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

		if actual.SecurityBaselineStateId != v.Expected.SecurityBaselineStateId {
			t.Fatalf("Expected %q but got %q for SecurityBaselineStateId", v.Expected.SecurityBaselineStateId, actual.SecurityBaselineStateId)
		}

		if actual.SecurityBaselineSettingStateId != v.Expected.SecurityBaselineSettingStateId {
			t.Fatalf("Expected %q but got %q for SecurityBaselineSettingStateId", v.Expected.SecurityBaselineSettingStateId, actual.SecurityBaselineSettingStateId)
		}

	}
}

func TestParseDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId
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
			Input: "/deviceManagement/comanagedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceIdValue/securityBaselineStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/sEcUrItYbAsElInEsTaTeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeIdVaLuE/sEtTiNgStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates/securityBaselineSettingStateIdValue",
			Expected: &DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId{
				ManagedDeviceId:                "managedDeviceIdValue",
				SecurityBaselineStateId:        "securityBaselineStateIdValue",
				SecurityBaselineSettingStateId: "securityBaselineSettingStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/comanagedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates/securityBaselineSettingStateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeIdVaLuE/sEtTiNgStAtEs/sEcUrItYbAsElInEsEtTiNgStAtEiDvAlUe",
			Expected: &DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId{
				ManagedDeviceId:                "mAnAgEdDeViCeIdVaLuE",
				SecurityBaselineStateId:        "sEcUrItYbAsElInEsTaTeIdVaLuE",
				SecurityBaselineSettingStateId: "sEcUrItYbAsElInEsEtTiNgStAtEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeIdVaLuE/sEtTiNgStAtEs/sEcUrItYbAsElInEsEtTiNgStAtEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

		if actual.SecurityBaselineStateId != v.Expected.SecurityBaselineStateId {
			t.Fatalf("Expected %q but got %q for SecurityBaselineStateId", v.Expected.SecurityBaselineStateId, actual.SecurityBaselineStateId)
		}

		if actual.SecurityBaselineSettingStateId != v.Expected.SecurityBaselineSettingStateId {
			t.Fatalf("Expected %q but got %q for SecurityBaselineSettingStateId", v.Expected.SecurityBaselineSettingStateId, actual.SecurityBaselineSettingStateId)
		}

	}
}

func TestSegmentsForDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId(t *testing.T) {
	segments := DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
