package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId{}

func TestNewDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateID(t *testing.T) {
	id := NewDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateID("advancedThreatProtectionOnboardingDeviceSettingStateIdValue")

	if id.AdvancedThreatProtectionOnboardingDeviceSettingStateId != "advancedThreatProtectionOnboardingDeviceSettingStateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AdvancedThreatProtectionOnboardingDeviceSettingStateId'", id.AdvancedThreatProtectionOnboardingDeviceSettingStateId, "advancedThreatProtectionOnboardingDeviceSettingStateIdValue")
	}
}

func TestFormatDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateID(t *testing.T) {
	actual := NewDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateID("advancedThreatProtectionOnboardingDeviceSettingStateIdValue").ID()
	expected := "/deviceManagement/advancedThreatProtectionOnboardingStateSummary/advancedThreatProtectionOnboardingDeviceSettingStates/advancedThreatProtectionOnboardingDeviceSettingStateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId
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
			Input: "/deviceManagement/advancedThreatProtectionOnboardingStateSummary",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/advancedThreatProtectionOnboardingStateSummary/advancedThreatProtectionOnboardingDeviceSettingStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/advancedThreatProtectionOnboardingStateSummary/advancedThreatProtectionOnboardingDeviceSettingStates/advancedThreatProtectionOnboardingDeviceSettingStateIdValue",
			Expected: &DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId{
				AdvancedThreatProtectionOnboardingDeviceSettingStateId: "advancedThreatProtectionOnboardingDeviceSettingStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/advancedThreatProtectionOnboardingStateSummary/advancedThreatProtectionOnboardingDeviceSettingStates/advancedThreatProtectionOnboardingDeviceSettingStateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AdvancedThreatProtectionOnboardingDeviceSettingStateId != v.Expected.AdvancedThreatProtectionOnboardingDeviceSettingStateId {
			t.Fatalf("Expected %q but got %q for AdvancedThreatProtectionOnboardingDeviceSettingStateId", v.Expected.AdvancedThreatProtectionOnboardingDeviceSettingStateId, actual.AdvancedThreatProtectionOnboardingDeviceSettingStateId)
		}

	}
}

func TestParseDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId
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
			Input: "/deviceManagement/advancedThreatProtectionOnboardingStateSummary",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aDvAnCeDtHrEaTpRoTeCtIoNoNbOaRdInGsTaTeSuMmArY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/advancedThreatProtectionOnboardingStateSummary/advancedThreatProtectionOnboardingDeviceSettingStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aDvAnCeDtHrEaTpRoTeCtIoNoNbOaRdInGsTaTeSuMmArY/aDvAnCeDtHrEaTpRoTeCtIoNoNbOaRdInGdEvIcEsEtTiNgStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/advancedThreatProtectionOnboardingStateSummary/advancedThreatProtectionOnboardingDeviceSettingStates/advancedThreatProtectionOnboardingDeviceSettingStateIdValue",
			Expected: &DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId{
				AdvancedThreatProtectionOnboardingDeviceSettingStateId: "advancedThreatProtectionOnboardingDeviceSettingStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/advancedThreatProtectionOnboardingStateSummary/advancedThreatProtectionOnboardingDeviceSettingStates/advancedThreatProtectionOnboardingDeviceSettingStateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aDvAnCeDtHrEaTpRoTeCtIoNoNbOaRdInGsTaTeSuMmArY/aDvAnCeDtHrEaTpRoTeCtIoNoNbOaRdInGdEvIcEsEtTiNgStAtEs/aDvAnCeDtHrEaTpRoTeCtIoNoNbOaRdInGdEvIcEsEtTiNgStAtEiDvAlUe",
			Expected: &DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId{
				AdvancedThreatProtectionOnboardingDeviceSettingStateId: "aDvAnCeDtHrEaTpRoTeCtIoNoNbOaRdInGdEvIcEsEtTiNgStAtEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aDvAnCeDtHrEaTpRoTeCtIoNoNbOaRdInGsTaTeSuMmArY/aDvAnCeDtHrEaTpRoTeCtIoNoNbOaRdInGdEvIcEsEtTiNgStAtEs/aDvAnCeDtHrEaTpRoTeCtIoNoNbOaRdInGdEvIcEsEtTiNgStAtEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AdvancedThreatProtectionOnboardingDeviceSettingStateId != v.Expected.AdvancedThreatProtectionOnboardingDeviceSettingStateId {
			t.Fatalf("Expected %q but got %q for AdvancedThreatProtectionOnboardingDeviceSettingStateId", v.Expected.AdvancedThreatProtectionOnboardingDeviceSettingStateId, actual.AdvancedThreatProtectionOnboardingDeviceSettingStateId)
		}

	}
}

func TestSegmentsForDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId(t *testing.T) {
	segments := DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
