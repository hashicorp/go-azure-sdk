package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsDeviceScopeId{}

func TestNewDeviceManagementUserExperienceAnalyticsDeviceScopeID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsDeviceScopeID("userExperienceAnalyticsDeviceScopeId")

	if id.UserExperienceAnalyticsDeviceScopeId != "userExperienceAnalyticsDeviceScopeId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsDeviceScopeId'", id.UserExperienceAnalyticsDeviceScopeId, "userExperienceAnalyticsDeviceScopeId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsDeviceScopeID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsDeviceScopeID("userExperienceAnalyticsDeviceScopeId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsDeviceScopes/userExperienceAnalyticsDeviceScopeId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsDeviceScopeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsDeviceScopeId
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
			Input: "/deviceManagement/userExperienceAnalyticsDeviceScopes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsDeviceScopes/userExperienceAnalyticsDeviceScopeId",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceScopeId{
				UserExperienceAnalyticsDeviceScopeId: "userExperienceAnalyticsDeviceScopeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsDeviceScopes/userExperienceAnalyticsDeviceScopeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsDeviceScopeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsDeviceScopeId != v.Expected.UserExperienceAnalyticsDeviceScopeId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsDeviceScopeId", v.Expected.UserExperienceAnalyticsDeviceScopeId, actual.UserExperienceAnalyticsDeviceScopeId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsDeviceScopeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsDeviceScopeId
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
			Input: "/deviceManagement/userExperienceAnalyticsDeviceScopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeScOpEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsDeviceScopes/userExperienceAnalyticsDeviceScopeId",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceScopeId{
				UserExperienceAnalyticsDeviceScopeId: "userExperienceAnalyticsDeviceScopeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsDeviceScopes/userExperienceAnalyticsDeviceScopeId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeScOpEs/uSeReXpErIeNcEaNaLyTiCsDeViCeScOpEiD",
			Expected: &DeviceManagementUserExperienceAnalyticsDeviceScopeId{
				UserExperienceAnalyticsDeviceScopeId: "uSeReXpErIeNcEaNaLyTiCsDeViCeScOpEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeScOpEs/uSeReXpErIeNcEaNaLyTiCsDeViCeScOpEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsDeviceScopeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsDeviceScopeId != v.Expected.UserExperienceAnalyticsDeviceScopeId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsDeviceScopeId", v.Expected.UserExperienceAnalyticsDeviceScopeId, actual.UserExperienceAnalyticsDeviceScopeId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsDeviceScopeId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsDeviceScopeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsDeviceScopeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
