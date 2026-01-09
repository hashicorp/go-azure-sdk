package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsRemoteConnectionId{}

func TestNewDeviceManagementUserExperienceAnalyticsRemoteConnectionID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsRemoteConnectionID("userExperienceAnalyticsRemoteConnectionId")

	if id.UserExperienceAnalyticsRemoteConnectionId != "userExperienceAnalyticsRemoteConnectionId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsRemoteConnectionId'", id.UserExperienceAnalyticsRemoteConnectionId, "userExperienceAnalyticsRemoteConnectionId")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsRemoteConnectionID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsRemoteConnectionID("userExperienceAnalyticsRemoteConnectionId").ID()
	expected := "/deviceManagement/userExperienceAnalyticsRemoteConnection/userExperienceAnalyticsRemoteConnectionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsRemoteConnectionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsRemoteConnectionId
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
			Input: "/deviceManagement/userExperienceAnalyticsRemoteConnection",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsRemoteConnection/userExperienceAnalyticsRemoteConnectionId",
			Expected: &DeviceManagementUserExperienceAnalyticsRemoteConnectionId{
				UserExperienceAnalyticsRemoteConnectionId: "userExperienceAnalyticsRemoteConnectionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsRemoteConnection/userExperienceAnalyticsRemoteConnectionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsRemoteConnectionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsRemoteConnectionId != v.Expected.UserExperienceAnalyticsRemoteConnectionId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsRemoteConnectionId", v.Expected.UserExperienceAnalyticsRemoteConnectionId, actual.UserExperienceAnalyticsRemoteConnectionId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsRemoteConnectionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsRemoteConnectionId
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
			Input: "/deviceManagement/userExperienceAnalyticsRemoteConnection",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsReMoTeCoNnEcTiOn",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsRemoteConnection/userExperienceAnalyticsRemoteConnectionId",
			Expected: &DeviceManagementUserExperienceAnalyticsRemoteConnectionId{
				UserExperienceAnalyticsRemoteConnectionId: "userExperienceAnalyticsRemoteConnectionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsRemoteConnection/userExperienceAnalyticsRemoteConnectionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsReMoTeCoNnEcTiOn/uSeReXpErIeNcEaNaLyTiCsReMoTeCoNnEcTiOnId",
			Expected: &DeviceManagementUserExperienceAnalyticsRemoteConnectionId{
				UserExperienceAnalyticsRemoteConnectionId: "uSeReXpErIeNcEaNaLyTiCsReMoTeCoNnEcTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsReMoTeCoNnEcTiOn/uSeReXpErIeNcEaNaLyTiCsReMoTeCoNnEcTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsRemoteConnectionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsRemoteConnectionId != v.Expected.UserExperienceAnalyticsRemoteConnectionId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsRemoteConnectionId", v.Expected.UserExperienceAnalyticsRemoteConnectionId, actual.UserExperienceAnalyticsRemoteConnectionId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsRemoteConnectionId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsRemoteConnectionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsRemoteConnectionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
