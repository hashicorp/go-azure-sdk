package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId{}

func TestNewDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateID(t *testing.T) {
	id := NewDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateID("managedDeviceId", "securityBaselineStateId", "securityBaselineSettingStateId")

	if id.ManagedDeviceId != "managedDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceId")
	}

	if id.SecurityBaselineStateId != "securityBaselineStateId" {
		t.Fatalf("Expected %q but got %q for Segment 'SecurityBaselineStateId'", id.SecurityBaselineStateId, "securityBaselineStateId")
	}

	if id.SecurityBaselineSettingStateId != "securityBaselineSettingStateId" {
		t.Fatalf("Expected %q but got %q for Segment 'SecurityBaselineSettingStateId'", id.SecurityBaselineSettingStateId, "securityBaselineSettingStateId")
	}
}

func TestFormatDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateID(t *testing.T) {
	actual := NewDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateID("managedDeviceId", "securityBaselineStateId", "securityBaselineSettingStateId").ID()
	expected := "/deviceManagement/comanagedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId/settingStates/securityBaselineSettingStateId"
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
			Input: "/deviceManagement/comanagedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/securityBaselineStates",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId/settingStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId/settingStates/securityBaselineSettingStateId",
			Expected: &DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId{
				ManagedDeviceId:                "managedDeviceId",
				SecurityBaselineStateId:        "securityBaselineStateId",
				SecurityBaselineSettingStateId: "securityBaselineSettingStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId/settingStates/securityBaselineSettingStateId/extra",
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
			Input: "/deviceManagement/comanagedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/securityBaselineStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeId/sEcUrItYbAsElInEsTaTeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeId/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId/settingStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeId/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeId/sEtTiNgStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId/settingStates/securityBaselineSettingStateId",
			Expected: &DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId{
				ManagedDeviceId:                "managedDeviceId",
				SecurityBaselineStateId:        "securityBaselineStateId",
				SecurityBaselineSettingStateId: "securityBaselineSettingStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId/settingStates/securityBaselineSettingStateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeId/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeId/sEtTiNgStAtEs/sEcUrItYbAsElInEsEtTiNgStAtEiD",
			Expected: &DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId{
				ManagedDeviceId:                "mAnAgEdDeViCeId",
				SecurityBaselineStateId:        "sEcUrItYbAsElInEsTaTeId",
				SecurityBaselineSettingStateId: "sEcUrItYbAsElInEsEtTiNgStAtEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeId/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeId/sEtTiNgStAtEs/sEcUrItYbAsElInEsEtTiNgStAtEiD/extra",
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
