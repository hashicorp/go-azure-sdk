package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntentIdSettingId{}

func TestNewDeviceManagementIntentIdSettingID(t *testing.T) {
	id := NewDeviceManagementIntentIdSettingID("deviceManagementIntentId", "deviceManagementSettingInstanceId")

	if id.DeviceManagementIntentId != "deviceManagementIntentId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementIntentId'", id.DeviceManagementIntentId, "deviceManagementIntentId")
	}

	if id.DeviceManagementSettingInstanceId != "deviceManagementSettingInstanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementSettingInstanceId'", id.DeviceManagementSettingInstanceId, "deviceManagementSettingInstanceId")
	}
}

func TestFormatDeviceManagementIntentIdSettingID(t *testing.T) {
	actual := NewDeviceManagementIntentIdSettingID("deviceManagementIntentId", "deviceManagementSettingInstanceId").ID()
	expected := "/deviceManagement/intents/deviceManagementIntentId/settings/deviceManagementSettingInstanceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementIntentIdSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIntentIdSettingId
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
			Input: "/deviceManagement/intents",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentId/settings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intents/deviceManagementIntentId/settings/deviceManagementSettingInstanceId",
			Expected: &DeviceManagementIntentIdSettingId{
				DeviceManagementIntentId:          "deviceManagementIntentId",
				DeviceManagementSettingInstanceId: "deviceManagementSettingInstanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intents/deviceManagementIntentId/settings/deviceManagementSettingInstanceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIntentIdSettingID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementIntentId != v.Expected.DeviceManagementIntentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementIntentId", v.Expected.DeviceManagementIntentId, actual.DeviceManagementIntentId)
		}

		if actual.DeviceManagementSettingInstanceId != v.Expected.DeviceManagementSettingInstanceId {
			t.Fatalf("Expected %q but got %q for DeviceManagementSettingInstanceId", v.Expected.DeviceManagementSettingInstanceId, actual.DeviceManagementSettingInstanceId)
		}

	}
}

func TestParseDeviceManagementIntentIdSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIntentIdSettingId
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
			Input: "/deviceManagement/intents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentId/settings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiD/sEtTiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intents/deviceManagementIntentId/settings/deviceManagementSettingInstanceId",
			Expected: &DeviceManagementIntentIdSettingId{
				DeviceManagementIntentId:          "deviceManagementIntentId",
				DeviceManagementSettingInstanceId: "deviceManagementSettingInstanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intents/deviceManagementIntentId/settings/deviceManagementSettingInstanceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiD/sEtTiNgS/dEvIcEmAnAgEmEnTsEtTiNgInStAnCeId",
			Expected: &DeviceManagementIntentIdSettingId{
				DeviceManagementIntentId:          "dEvIcEmAnAgEmEnTiNtEnTiD",
				DeviceManagementSettingInstanceId: "dEvIcEmAnAgEmEnTsEtTiNgInStAnCeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiD/sEtTiNgS/dEvIcEmAnAgEmEnTsEtTiNgInStAnCeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIntentIdSettingIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementIntentId != v.Expected.DeviceManagementIntentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementIntentId", v.Expected.DeviceManagementIntentId, actual.DeviceManagementIntentId)
		}

		if actual.DeviceManagementSettingInstanceId != v.Expected.DeviceManagementSettingInstanceId {
			t.Fatalf("Expected %q but got %q for DeviceManagementSettingInstanceId", v.Expected.DeviceManagementSettingInstanceId, actual.DeviceManagementSettingInstanceId)
		}

	}
}

func TestSegmentsForDeviceManagementIntentIdSettingId(t *testing.T) {
	segments := DeviceManagementIntentIdSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementIntentIdSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
