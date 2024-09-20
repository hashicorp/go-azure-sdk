package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementReusablePolicySettingId{}

func TestNewDeviceManagementReusablePolicySettingID(t *testing.T) {
	id := NewDeviceManagementReusablePolicySettingID("deviceManagementReusablePolicySettingId")

	if id.DeviceManagementReusablePolicySettingId != "deviceManagementReusablePolicySettingId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementReusablePolicySettingId'", id.DeviceManagementReusablePolicySettingId, "deviceManagementReusablePolicySettingId")
	}
}

func TestFormatDeviceManagementReusablePolicySettingID(t *testing.T) {
	actual := NewDeviceManagementReusablePolicySettingID("deviceManagementReusablePolicySettingId").ID()
	expected := "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementReusablePolicySettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementReusablePolicySettingId
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
			Input: "/deviceManagement/reusablePolicySettings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId",
			Expected: &DeviceManagementReusablePolicySettingId{
				DeviceManagementReusablePolicySettingId: "deviceManagementReusablePolicySettingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementReusablePolicySettingID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementReusablePolicySettingId != v.Expected.DeviceManagementReusablePolicySettingId {
			t.Fatalf("Expected %q but got %q for DeviceManagementReusablePolicySettingId", v.Expected.DeviceManagementReusablePolicySettingId, actual.DeviceManagementReusablePolicySettingId)
		}

	}
}

func TestParseDeviceManagementReusablePolicySettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementReusablePolicySettingId
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
			Input: "/deviceManagement/reusablePolicySettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId",
			Expected: &DeviceManagementReusablePolicySettingId{
				DeviceManagementReusablePolicySettingId: "deviceManagementReusablePolicySettingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgId",
			Expected: &DeviceManagementReusablePolicySettingId{
				DeviceManagementReusablePolicySettingId: "dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementReusablePolicySettingIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementReusablePolicySettingId != v.Expected.DeviceManagementReusablePolicySettingId {
			t.Fatalf("Expected %q but got %q for DeviceManagementReusablePolicySettingId", v.Expected.DeviceManagementReusablePolicySettingId, actual.DeviceManagementReusablePolicySettingId)
		}

	}
}

func TestSegmentsForDeviceManagementReusablePolicySettingId(t *testing.T) {
	segments := DeviceManagementReusablePolicySettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementReusablePolicySettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
