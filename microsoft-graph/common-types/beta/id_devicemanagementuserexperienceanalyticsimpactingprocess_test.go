package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsImpactingProcessId{}

func TestNewDeviceManagementUserExperienceAnalyticsImpactingProcessID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsImpactingProcessID("userExperienceAnalyticsImpactingProcessIdValue")

	if id.UserExperienceAnalyticsImpactingProcessId != "userExperienceAnalyticsImpactingProcessIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsImpactingProcessId'", id.UserExperienceAnalyticsImpactingProcessId, "userExperienceAnalyticsImpactingProcessIdValue")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsImpactingProcessID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsImpactingProcessID("userExperienceAnalyticsImpactingProcessIdValue").ID()
	expected := "/deviceManagement/userExperienceAnalyticsImpactingProcess/userExperienceAnalyticsImpactingProcessIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsImpactingProcessID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsImpactingProcessId
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
			Input: "/deviceManagement/userExperienceAnalyticsImpactingProcess",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsImpactingProcess/userExperienceAnalyticsImpactingProcessIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsImpactingProcessId{
				UserExperienceAnalyticsImpactingProcessId: "userExperienceAnalyticsImpactingProcessIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsImpactingProcess/userExperienceAnalyticsImpactingProcessIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsImpactingProcessID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsImpactingProcessId != v.Expected.UserExperienceAnalyticsImpactingProcessId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsImpactingProcessId", v.Expected.UserExperienceAnalyticsImpactingProcessId, actual.UserExperienceAnalyticsImpactingProcessId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsImpactingProcessIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsImpactingProcessId
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
			Input: "/deviceManagement/userExperienceAnalyticsImpactingProcess",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsImPaCtInGpRoCeSs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsImpactingProcess/userExperienceAnalyticsImpactingProcessIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsImpactingProcessId{
				UserExperienceAnalyticsImpactingProcessId: "userExperienceAnalyticsImpactingProcessIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsImpactingProcess/userExperienceAnalyticsImpactingProcessIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsImPaCtInGpRoCeSs/uSeReXpErIeNcEaNaLyTiCsImPaCtInGpRoCeSsIdVaLuE",
			Expected: &DeviceManagementUserExperienceAnalyticsImpactingProcessId{
				UserExperienceAnalyticsImpactingProcessId: "uSeReXpErIeNcEaNaLyTiCsImPaCtInGpRoCeSsIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsImPaCtInGpRoCeSs/uSeReXpErIeNcEaNaLyTiCsImPaCtInGpRoCeSsIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsImpactingProcessIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsImpactingProcessId != v.Expected.UserExperienceAnalyticsImpactingProcessId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsImpactingProcessId", v.Expected.UserExperienceAnalyticsImpactingProcessId, actual.UserExperienceAnalyticsImpactingProcessId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsImpactingProcessId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsImpactingProcessId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsImpactingProcessId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
