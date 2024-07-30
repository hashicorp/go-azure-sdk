package userexperienceanalyticsdeviceswithoutcloudidentity

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId{}

func TestNewDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityID("userExperienceAnalyticsDeviceWithoutCloudIdentityIdValue")

	if id.UserExperienceAnalyticsDeviceWithoutCloudIdentityId != "userExperienceAnalyticsDeviceWithoutCloudIdentityIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsDeviceWithoutCloudIdentityId'", id.UserExperienceAnalyticsDeviceWithoutCloudIdentityId, "userExperienceAnalyticsDeviceWithoutCloudIdentityIdValue")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityID("userExperienceAnalyticsDeviceWithoutCloudIdentityIdValue").ID()
	expected := "/deviceManagement/userExperienceAnalyticsDevicesWithoutCloudIdentity/userExperienceAnalyticsDeviceWithoutCloudIdentityIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId
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
			Input: "/deviceManagement/userExperienceAnalyticsDevicesWithoutCloudIdentity",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsDevicesWithoutCloudIdentity/userExperienceAnalyticsDeviceWithoutCloudIdentityIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId{
				UserExperienceAnalyticsDeviceWithoutCloudIdentityId: "userExperienceAnalyticsDeviceWithoutCloudIdentityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsDevicesWithoutCloudIdentity/userExperienceAnalyticsDeviceWithoutCloudIdentityIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsDeviceWithoutCloudIdentityId != v.Expected.UserExperienceAnalyticsDeviceWithoutCloudIdentityId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsDeviceWithoutCloudIdentityId", v.Expected.UserExperienceAnalyticsDeviceWithoutCloudIdentityId, actual.UserExperienceAnalyticsDeviceWithoutCloudIdentityId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId
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
			Input: "/deviceManagement/userExperienceAnalyticsDevicesWithoutCloudIdentity",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeSwItHoUtClOuDiDeNtItY",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsDevicesWithoutCloudIdentity/userExperienceAnalyticsDeviceWithoutCloudIdentityIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId{
				UserExperienceAnalyticsDeviceWithoutCloudIdentityId: "userExperienceAnalyticsDeviceWithoutCloudIdentityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsDevicesWithoutCloudIdentity/userExperienceAnalyticsDeviceWithoutCloudIdentityIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeSwItHoUtClOuDiDeNtItY/uSeReXpErIeNcEaNaLyTiCsDeViCeWiThOuTcLoUdIdEnTiTyIdVaLuE",
			Expected: &DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId{
				UserExperienceAnalyticsDeviceWithoutCloudIdentityId: "uSeReXpErIeNcEaNaLyTiCsDeViCeWiThOuTcLoUdIdEnTiTyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsDeViCeSwItHoUtClOuDiDeNtItY/uSeReXpErIeNcEaNaLyTiCsDeViCeWiThOuTcLoUdIdEnTiTyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsDeviceWithoutCloudIdentityId != v.Expected.UserExperienceAnalyticsDeviceWithoutCloudIdentityId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsDeviceWithoutCloudIdentityId", v.Expected.UserExperienceAnalyticsDeviceWithoutCloudIdentityId, actual.UserExperienceAnalyticsDeviceWithoutCloudIdentityId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
