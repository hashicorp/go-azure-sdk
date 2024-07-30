package userexperienceanalyticsbatteryhealthappimpact

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId{}

func TestNewDeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactID("userExperienceAnalyticsBatteryHealthAppImpactIdValue")

	if id.UserExperienceAnalyticsBatteryHealthAppImpactId != "userExperienceAnalyticsBatteryHealthAppImpactIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsBatteryHealthAppImpactId'", id.UserExperienceAnalyticsBatteryHealthAppImpactId, "userExperienceAnalyticsBatteryHealthAppImpactIdValue")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactID("userExperienceAnalyticsBatteryHealthAppImpactIdValue").ID()
	expected := "/deviceManagement/userExperienceAnalyticsBatteryHealthAppImpact/userExperienceAnalyticsBatteryHealthAppImpactIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId
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
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthAppImpact",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthAppImpact/userExperienceAnalyticsBatteryHealthAppImpactIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId{
				UserExperienceAnalyticsBatteryHealthAppImpactId: "userExperienceAnalyticsBatteryHealthAppImpactIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthAppImpact/userExperienceAnalyticsBatteryHealthAppImpactIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsBatteryHealthAppImpactId != v.Expected.UserExperienceAnalyticsBatteryHealthAppImpactId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsBatteryHealthAppImpactId", v.Expected.UserExperienceAnalyticsBatteryHealthAppImpactId, actual.UserExperienceAnalyticsBatteryHealthAppImpactId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId
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
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthAppImpact",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHaPpImPaCt",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthAppImpact/userExperienceAnalyticsBatteryHealthAppImpactIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId{
				UserExperienceAnalyticsBatteryHealthAppImpactId: "userExperienceAnalyticsBatteryHealthAppImpactIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsBatteryHealthAppImpact/userExperienceAnalyticsBatteryHealthAppImpactIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHaPpImPaCt/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHaPpImPaCtIdVaLuE",
			Expected: &DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId{
				UserExperienceAnalyticsBatteryHealthAppImpactId: "uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHaPpImPaCtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHaPpImPaCt/uSeReXpErIeNcEaNaLyTiCsBaTtErYhEaLtHaPpImPaCtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsBatteryHealthAppImpactId != v.Expected.UserExperienceAnalyticsBatteryHealthAppImpactId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsBatteryHealthAppImpactId", v.Expected.UserExperienceAnalyticsBatteryHealthAppImpactId, actual.UserExperienceAnalyticsBatteryHealthAppImpactId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
