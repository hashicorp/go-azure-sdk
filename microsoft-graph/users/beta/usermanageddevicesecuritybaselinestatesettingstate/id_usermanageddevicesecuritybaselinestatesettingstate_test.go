package usermanageddevicesecuritybaselinestatesettingstate

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserManagedDeviceSecurityBaselineStateSettingStateId{}

func TestNewUserManagedDeviceSecurityBaselineStateSettingStateID(t *testing.T) {
	id := NewUserManagedDeviceSecurityBaselineStateSettingStateID("userIdValue", "managedDeviceIdValue", "securityBaselineStateIdValue", "securityBaselineSettingStateIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

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

func TestFormatUserManagedDeviceSecurityBaselineStateSettingStateID(t *testing.T) {
	actual := NewUserManagedDeviceSecurityBaselineStateSettingStateID("userIdValue", "managedDeviceIdValue", "securityBaselineStateIdValue", "securityBaselineSettingStateIdValue").ID()
	expected := "/users/userIdValue/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates/securityBaselineSettingStateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserManagedDeviceSecurityBaselineStateSettingStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserManagedDeviceSecurityBaselineStateSettingStateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/securityBaselineStates",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates/securityBaselineSettingStateIdValue",
			Expected: &UserManagedDeviceSecurityBaselineStateSettingStateId{
				UserId:                         "userIdValue",
				ManagedDeviceId:                "managedDeviceIdValue",
				SecurityBaselineStateId:        "securityBaselineStateIdValue",
				SecurityBaselineSettingStateId: "securityBaselineSettingStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates/securityBaselineSettingStateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserManagedDeviceSecurityBaselineStateSettingStateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
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

func TestParseUserManagedDeviceSecurityBaselineStateSettingStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserManagedDeviceSecurityBaselineStateSettingStateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/securityBaselineStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/sEcUrItYbAsElInEsTaTeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeIdVaLuE/sEtTiNgStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates/securityBaselineSettingStateIdValue",
			Expected: &UserManagedDeviceSecurityBaselineStateSettingStateId{
				UserId:                         "userIdValue",
				ManagedDeviceId:                "managedDeviceIdValue",
				SecurityBaselineStateId:        "securityBaselineStateIdValue",
				SecurityBaselineSettingStateId: "securityBaselineSettingStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/settingStates/securityBaselineSettingStateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeIdVaLuE/sEtTiNgStAtEs/sEcUrItYbAsElInEsEtTiNgStAtEiDvAlUe",
			Expected: &UserManagedDeviceSecurityBaselineStateSettingStateId{
				UserId:                         "uSeRiDvAlUe",
				ManagedDeviceId:                "mAnAgEdDeViCeIdVaLuE",
				SecurityBaselineStateId:        "sEcUrItYbAsElInEsTaTeIdVaLuE",
				SecurityBaselineSettingStateId: "sEcUrItYbAsElInEsEtTiNgStAtEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeIdVaLuE/sEtTiNgStAtEs/sEcUrItYbAsElInEsEtTiNgStAtEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserManagedDeviceSecurityBaselineStateSettingStateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
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

func TestSegmentsForUserManagedDeviceSecurityBaselineStateSettingStateId(t *testing.T) {
	segments := UserManagedDeviceSecurityBaselineStateSettingStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserManagedDeviceSecurityBaselineStateSettingStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
