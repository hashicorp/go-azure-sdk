package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceConfigurationRestrictedAppsViolationId{}

func TestNewDeviceManagementDeviceConfigurationRestrictedAppsViolationID(t *testing.T) {
	id := NewDeviceManagementDeviceConfigurationRestrictedAppsViolationID("restrictedAppsViolationId")

	if id.RestrictedAppsViolationId != "restrictedAppsViolationId" {
		t.Fatalf("Expected %q but got %q for Segment 'RestrictedAppsViolationId'", id.RestrictedAppsViolationId, "restrictedAppsViolationId")
	}
}

func TestFormatDeviceManagementDeviceConfigurationRestrictedAppsViolationID(t *testing.T) {
	actual := NewDeviceManagementDeviceConfigurationRestrictedAppsViolationID("restrictedAppsViolationId").ID()
	expected := "/deviceManagement/deviceConfigurationRestrictedAppsViolations/restrictedAppsViolationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceConfigurationRestrictedAppsViolationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceConfigurationRestrictedAppsViolationId
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
			Input: "/deviceManagement/deviceConfigurationRestrictedAppsViolations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceConfigurationRestrictedAppsViolations/restrictedAppsViolationId",
			Expected: &DeviceManagementDeviceConfigurationRestrictedAppsViolationId{
				RestrictedAppsViolationId: "restrictedAppsViolationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceConfigurationRestrictedAppsViolations/restrictedAppsViolationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceConfigurationRestrictedAppsViolationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.RestrictedAppsViolationId != v.Expected.RestrictedAppsViolationId {
			t.Fatalf("Expected %q but got %q for RestrictedAppsViolationId", v.Expected.RestrictedAppsViolationId, actual.RestrictedAppsViolationId)
		}

	}
}

func TestParseDeviceManagementDeviceConfigurationRestrictedAppsViolationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceConfigurationRestrictedAppsViolationId
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
			Input: "/deviceManagement/deviceConfigurationRestrictedAppsViolations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnReStRiCtEdApPsViOlAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceConfigurationRestrictedAppsViolations/restrictedAppsViolationId",
			Expected: &DeviceManagementDeviceConfigurationRestrictedAppsViolationId{
				RestrictedAppsViolationId: "restrictedAppsViolationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceConfigurationRestrictedAppsViolations/restrictedAppsViolationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnReStRiCtEdApPsViOlAtIoNs/rEsTrIcTeDaPpSvIoLaTiOnId",
			Expected: &DeviceManagementDeviceConfigurationRestrictedAppsViolationId{
				RestrictedAppsViolationId: "rEsTrIcTeDaPpSvIoLaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnReStRiCtEdApPsViOlAtIoNs/rEsTrIcTeDaPpSvIoLaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceConfigurationRestrictedAppsViolationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.RestrictedAppsViolationId != v.Expected.RestrictedAppsViolationId {
			t.Fatalf("Expected %q but got %q for RestrictedAppsViolationId", v.Expected.RestrictedAppsViolationId, actual.RestrictedAppsViolationId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceConfigurationRestrictedAppsViolationId(t *testing.T) {
	segments := DeviceManagementDeviceConfigurationRestrictedAppsViolationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceConfigurationRestrictedAppsViolationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
