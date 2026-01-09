package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId{}

func TestNewDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactID("userExperienceAnalyticsBatteryHealthDeviceAppImpactId")

	if id.UserExperienceAnalyticsBatteryHealthDeviceAppImpactId != "userExperienceAnalyticsBatteryHealthDeviceAppImpactId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsBatteryHealthDeviceAppImpactId'", id.UserExperienceAnalyticsBatteryHealthDeviceAppImpactId, "userExperienceAnalyticsBatteryHealthDeviceAppImpactId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactID("userExperienceAnalyticsBatteryHealthDeviceAppImpactId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceAppImpact/userExperienceAnalyticsBatteryHealthDeviceAppImpactId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId
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
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceAppImpact",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceAppImpact/userExperienceAnalyticsBatteryHealthDeviceAppImpactId",
			Expected: &DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId{
				UserExperienceAnalyticsBatteryHealthDeviceAppImpactId: "userExperienceAnalyticsBatteryHealthDeviceAppImpactId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceAppImpact/userExperienceAnalyticsBatteryHealthDeviceAppImpactId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsBatteryHealthDeviceAppImpactId != v.Expected.UserExperienceAnalyticsBatteryHealthDeviceAppImpactId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsBatteryHealthDeviceAppImpactId", v.Expected.UserExperienceAnalyticsBatteryHealthDeviceAppImpactId, actual.UserExperienceAnalyticsBatteryHealthDeviceAppImpactId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId
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
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceAppImpact",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHdEvIcEaPpImPaCt",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceAppImpact/userExperienceAnalyticsBatteryHealthDeviceAppImpactId",
			Expected: &DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId{
				UserExperienceAnalyticsBatteryHealthDeviceAppImpactId: "userExperienceAnalyticsBatteryHealthDeviceAppImpactId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceAppImpact/userExperienceAnalyticsBatteryHealthDeviceAppImpactId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHdEvIcEaPpImPaCt/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHdEvIcEaPpImPaCtId",
			Expected: &DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId{
				UserExperienceAnalyticsBatteryHealthDeviceAppImpactId: "uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHdEvIcEaPpImPaCtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHdEvIcEaPpImPaCt/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHdEvIcEaPpImPaCtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsBatteryHealthDeviceAppImpactId != v.Expected.UserExperienceAnalyticsBatteryHealthDeviceAppImpactId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsBatteryHealthDeviceAppImpactId", v.Expected.UserExperienceAnalyticsBatteryHealthDeviceAppImpactId, actual.UserExperienceAnalyticsBatteryHealthDeviceAppImpactId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
