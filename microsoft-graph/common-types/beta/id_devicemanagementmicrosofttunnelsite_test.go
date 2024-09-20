package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMicrosoftTunnelSiteId{}

func TestNewDeviceManagementMicrosoftTunnelSiteID(t *testing.T) {
	id := NewDeviceManagementMicrosoftTunnelSiteID("microsoftTunnelSiteId")

	if id.MicrosoftTunnelSiteId != "microsoftTunnelSiteId" {
		t.Fatalf("Expected %q but got %q for Segment 'MicrosoftTunnelSiteId'", id.MicrosoftTunnelSiteId, "microsoftTunnelSiteId")
	}
}

func TestFormatDeviceManagementMicrosoftTunnelSiteID(t *testing.T) {
	actual := NewDeviceManagementMicrosoftTunnelSiteID("microsoftTunnelSiteId").ID()
	expected := "/deviceManagement/microsoftTunnelSites/microsoftTunnelSiteId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementMicrosoftTunnelSiteID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMicrosoftTunnelSiteId
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
			Input: "/deviceManagement/microsoftTunnelSites",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/microsoftTunnelSites/microsoftTunnelSiteId",
			Expected: &DeviceManagementMicrosoftTunnelSiteId{
				MicrosoftTunnelSiteId: "microsoftTunnelSiteId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/microsoftTunnelSites/microsoftTunnelSiteId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMicrosoftTunnelSiteID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MicrosoftTunnelSiteId != v.Expected.MicrosoftTunnelSiteId {
			t.Fatalf("Expected %q but got %q for MicrosoftTunnelSiteId", v.Expected.MicrosoftTunnelSiteId, actual.MicrosoftTunnelSiteId)
		}

	}
}

func TestParseDeviceManagementMicrosoftTunnelSiteIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMicrosoftTunnelSiteId
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
			Input: "/deviceManagement/microsoftTunnelSites",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mIcRoSoFtTuNnElSiTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/microsoftTunnelSites/microsoftTunnelSiteId",
			Expected: &DeviceManagementMicrosoftTunnelSiteId{
				MicrosoftTunnelSiteId: "microsoftTunnelSiteId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/microsoftTunnelSites/microsoftTunnelSiteId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mIcRoSoFtTuNnElSiTeS/mIcRoSoFtTuNnElSiTeId",
			Expected: &DeviceManagementMicrosoftTunnelSiteId{
				MicrosoftTunnelSiteId: "mIcRoSoFtTuNnElSiTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mIcRoSoFtTuNnElSiTeS/mIcRoSoFtTuNnElSiTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMicrosoftTunnelSiteIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MicrosoftTunnelSiteId != v.Expected.MicrosoftTunnelSiteId {
			t.Fatalf("Expected %q but got %q for MicrosoftTunnelSiteId", v.Expected.MicrosoftTunnelSiteId, actual.MicrosoftTunnelSiteId)
		}

	}
}

func TestSegmentsForDeviceManagementMicrosoftTunnelSiteId(t *testing.T) {
	segments := DeviceManagementMicrosoftTunnelSiteId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementMicrosoftTunnelSiteId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
