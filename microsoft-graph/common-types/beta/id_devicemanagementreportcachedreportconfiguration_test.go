package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementReportCachedReportConfigurationId{}

func TestNewDeviceManagementReportCachedReportConfigurationID(t *testing.T) {
	id := NewDeviceManagementReportCachedReportConfigurationID("deviceManagementCachedReportConfigurationId")

	if id.DeviceManagementCachedReportConfigurationId != "deviceManagementCachedReportConfigurationId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementCachedReportConfigurationId'", id.DeviceManagementCachedReportConfigurationId, "deviceManagementCachedReportConfigurationId")
	}
}

func TestFormatDeviceManagementReportCachedReportConfigurationID(t *testing.T) {
	actual := NewDeviceManagementReportCachedReportConfigurationID("deviceManagementCachedReportConfigurationId").ID()
	expected := "/deviceManagement/reports/cachedReportConfigurations/deviceManagementCachedReportConfigurationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementReportCachedReportConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementReportCachedReportConfigurationId
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
			Input: "/deviceManagement/reports",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reports/cachedReportConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/reports/cachedReportConfigurations/deviceManagementCachedReportConfigurationId",
			Expected: &DeviceManagementReportCachedReportConfigurationId{
				DeviceManagementCachedReportConfigurationId: "deviceManagementCachedReportConfigurationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/reports/cachedReportConfigurations/deviceManagementCachedReportConfigurationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementReportCachedReportConfigurationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementCachedReportConfigurationId != v.Expected.DeviceManagementCachedReportConfigurationId {
			t.Fatalf("Expected %q but got %q for DeviceManagementCachedReportConfigurationId", v.Expected.DeviceManagementCachedReportConfigurationId, actual.DeviceManagementCachedReportConfigurationId)
		}

	}
}

func TestParseDeviceManagementReportCachedReportConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementReportCachedReportConfigurationId
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
			Input: "/deviceManagement/reports",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEpOrTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reports/cachedReportConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEpOrTs/cAcHeDrEpOrTcOnFiGuRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/reports/cachedReportConfigurations/deviceManagementCachedReportConfigurationId",
			Expected: &DeviceManagementReportCachedReportConfigurationId{
				DeviceManagementCachedReportConfigurationId: "deviceManagementCachedReportConfigurationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/reports/cachedReportConfigurations/deviceManagementCachedReportConfigurationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEpOrTs/cAcHeDrEpOrTcOnFiGuRaTiOnS/dEvIcEmAnAgEmEnTcAcHeDrEpOrTcOnFiGuRaTiOnId",
			Expected: &DeviceManagementReportCachedReportConfigurationId{
				DeviceManagementCachedReportConfigurationId: "dEvIcEmAnAgEmEnTcAcHeDrEpOrTcOnFiGuRaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEpOrTs/cAcHeDrEpOrTcOnFiGuRaTiOnS/dEvIcEmAnAgEmEnTcAcHeDrEpOrTcOnFiGuRaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementReportCachedReportConfigurationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementCachedReportConfigurationId != v.Expected.DeviceManagementCachedReportConfigurationId {
			t.Fatalf("Expected %q but got %q for DeviceManagementCachedReportConfigurationId", v.Expected.DeviceManagementCachedReportConfigurationId, actual.DeviceManagementCachedReportConfigurationId)
		}

	}
}

func TestSegmentsForDeviceManagementReportCachedReportConfigurationId(t *testing.T) {
	segments := DeviceManagementReportCachedReportConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementReportCachedReportConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
