package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId{}

func TestNewUserIdManagedDeviceIdSecurityBaselineStateIdSettingStateID(t *testing.T) {
	id := NewUserIdManagedDeviceIdSecurityBaselineStateIdSettingStateID("userId", "managedDeviceId", "securityBaselineStateId", "securityBaselineSettingStateId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

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

func TestFormatUserIdManagedDeviceIdSecurityBaselineStateIdSettingStateID(t *testing.T) {
	actual := NewUserIdManagedDeviceIdSecurityBaselineStateIdSettingStateID("userId", "managedDeviceId", "securityBaselineStateId", "securityBaselineSettingStateId").ID()
	expected := "/users/userId/managedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId/settingStates/securityBaselineSettingStateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdManagedDeviceIdSecurityBaselineStateIdSettingStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices/managedDeviceId/securityBaselineStates",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId/settingStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/managedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId/settingStates/securityBaselineSettingStateId",
			Expected: &UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId{
				UserId:                         "userId",
				ManagedDeviceId:                "managedDeviceId",
				SecurityBaselineStateId:        "securityBaselineStateId",
				SecurityBaselineSettingStateId: "securityBaselineSettingStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/managedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId/settingStates/securityBaselineSettingStateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdManagedDeviceIdSecurityBaselineStateIdSettingStateID(v.Input)
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

func TestParseUserIdManagedDeviceIdSecurityBaselineStateIdSettingStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices/managedDeviceId/securityBaselineStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId/sEcUrItYbAsElInEsTaTeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId/settingStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeId/sEtTiNgStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/managedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId/settingStates/securityBaselineSettingStateId",
			Expected: &UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId{
				UserId:                         "userId",
				ManagedDeviceId:                "managedDeviceId",
				SecurityBaselineStateId:        "securityBaselineStateId",
				SecurityBaselineSettingStateId: "securityBaselineSettingStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/managedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId/settingStates/securityBaselineSettingStateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeId/sEtTiNgStAtEs/sEcUrItYbAsElInEsEtTiNgStAtEiD",
			Expected: &UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId{
				UserId:                         "uSeRiD",
				ManagedDeviceId:                "mAnAgEdDeViCeId",
				SecurityBaselineStateId:        "sEcUrItYbAsElInEsTaTeId",
				SecurityBaselineSettingStateId: "sEcUrItYbAsElInEsEtTiNgStAtEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeId/sEtTiNgStAtEs/sEcUrItYbAsElInEsEtTiNgStAtEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdManagedDeviceIdSecurityBaselineStateIdSettingStateIDInsensitively(v.Input)
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

func TestSegmentsForUserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId(t *testing.T) {
	segments := UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdManagedDeviceIdSecurityBaselineStateIdSettingStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
