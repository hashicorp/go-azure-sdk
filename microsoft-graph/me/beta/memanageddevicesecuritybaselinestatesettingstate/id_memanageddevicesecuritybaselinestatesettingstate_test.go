package memanageddevicesecuritybaselinestatesettingstate

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeManagedDeviceSecurityBaselineStateSettingStateId{}

func TestNewMeManagedDeviceSecurityBaselineStateSettingStateID(t *testing.T) {
	id := NewMeManagedDeviceSecurityBaselineStateSettingStateID("managedDeviceIdValue", "securityBaselineStateIdValue", "securityBaselineSettingStateIdValue")

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

func TestFormatMeManagedDeviceSecurityBaselineStateSettingStateID(t *testing.T) {
	actual := NewMeManagedDeviceSecurityBaselineStateSettingStateID("managedDeviceIdValue", "securityBaselineStateIdValue", "securityBaselineSettingStateIdValue").ID()
	expected := "/me/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates/securityBaselineSettingStateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeManagedDeviceSecurityBaselineStateSettingStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeManagedDeviceSecurityBaselineStateSettingStateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceIdValue/securityBaselineStates",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates/securityBaselineSettingStateIdValue",
			Expected: &MeManagedDeviceSecurityBaselineStateSettingStateId{
				ManagedDeviceId:                "managedDeviceIdValue",
				SecurityBaselineStateId:        "securityBaselineStateIdValue",
				SecurityBaselineSettingStateId: "securityBaselineSettingStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates/securityBaselineSettingStateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeManagedDeviceSecurityBaselineStateSettingStateID(v.Input)
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

func TestParseMeManagedDeviceSecurityBaselineStateSettingStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeManagedDeviceSecurityBaselineStateSettingStateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceIdValue/securityBaselineStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/sEcUrItYbAsElInEsTaTeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeIdVaLuE/sEtTiNgStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates/securityBaselineSettingStateIdValue",
			Expected: &MeManagedDeviceSecurityBaselineStateSettingStateId{
				ManagedDeviceId:                "managedDeviceIdValue",
				SecurityBaselineStateId:        "securityBaselineStateIdValue",
				SecurityBaselineSettingStateId: "securityBaselineSettingStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates/securityBaselineSettingStateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeIdVaLuE/sEtTiNgStAtEs/sEcUrItYbAsElInEsEtTiNgStAtEiDvAlUe",
			Expected: &MeManagedDeviceSecurityBaselineStateSettingStateId{
				ManagedDeviceId:                "mAnAgEdDeViCeIdVaLuE",
				SecurityBaselineStateId:        "sEcUrItYbAsElInEsTaTeIdVaLuE",
				SecurityBaselineSettingStateId: "sEcUrItYbAsElInEsEtTiNgStAtEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeIdVaLuE/sEtTiNgStAtEs/sEcUrItYbAsElInEsEtTiNgStAtEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeManagedDeviceSecurityBaselineStateSettingStateIDInsensitively(v.Input)
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

func TestSegmentsForMeManagedDeviceSecurityBaselineStateSettingStateId(t *testing.T) {
	segments := MeManagedDeviceSecurityBaselineStateSettingStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeManagedDeviceSecurityBaselineStateSettingStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
