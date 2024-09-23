package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId{}

func TestNewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID(t *testing.T) {
	id := NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteId", "microsoftTunnelServerId")

	if id.MicrosoftTunnelSiteId != "microsoftTunnelSiteId" {
		t.Fatalf("Expected %q but got %q for Segment 'MicrosoftTunnelSiteId'", id.MicrosoftTunnelSiteId, "microsoftTunnelSiteId")
	}

	if id.MicrosoftTunnelServerId != "microsoftTunnelServerId" {
		t.Fatalf("Expected %q but got %q for Segment 'MicrosoftTunnelServerId'", id.MicrosoftTunnelServerId, "microsoftTunnelServerId")
	}
}

func TestFormatDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID(t *testing.T) {
	actual := NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteId", "microsoftTunnelServerId").ID()
	expected := "/deviceManagement/microsoftTunnelSites/microsoftTunnelSiteId/microsoftTunnelServers/microsoftTunnelServerId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId
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
			// Incomplete URI
			Input: "/deviceManagement/microsoftTunnelSites/microsoftTunnelSiteId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/microsoftTunnelSites/microsoftTunnelSiteId/microsoftTunnelServers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/microsoftTunnelSites/microsoftTunnelSiteId/microsoftTunnelServers/microsoftTunnelServerId",
			Expected: &DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId{
				MicrosoftTunnelSiteId:   "microsoftTunnelSiteId",
				MicrosoftTunnelServerId: "microsoftTunnelServerId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/microsoftTunnelSites/microsoftTunnelSiteId/microsoftTunnelServers/microsoftTunnelServerId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID(v.Input)
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

		if actual.MicrosoftTunnelServerId != v.Expected.MicrosoftTunnelServerId {
			t.Fatalf("Expected %q but got %q for MicrosoftTunnelServerId", v.Expected.MicrosoftTunnelServerId, actual.MicrosoftTunnelServerId)
		}

	}
}

func TestParseDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId
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
			// Incomplete URI
			Input: "/deviceManagement/microsoftTunnelSites/microsoftTunnelSiteId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mIcRoSoFtTuNnElSiTeS/mIcRoSoFtTuNnElSiTeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/microsoftTunnelSites/microsoftTunnelSiteId/microsoftTunnelServers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mIcRoSoFtTuNnElSiTeS/mIcRoSoFtTuNnElSiTeId/mIcRoSoFtTuNnElSeRvErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/microsoftTunnelSites/microsoftTunnelSiteId/microsoftTunnelServers/microsoftTunnelServerId",
			Expected: &DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId{
				MicrosoftTunnelSiteId:   "microsoftTunnelSiteId",
				MicrosoftTunnelServerId: "microsoftTunnelServerId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/microsoftTunnelSites/microsoftTunnelSiteId/microsoftTunnelServers/microsoftTunnelServerId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mIcRoSoFtTuNnElSiTeS/mIcRoSoFtTuNnElSiTeId/mIcRoSoFtTuNnElSeRvErS/mIcRoSoFtTuNnElSeRvErId",
			Expected: &DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId{
				MicrosoftTunnelSiteId:   "mIcRoSoFtTuNnElSiTeId",
				MicrosoftTunnelServerId: "mIcRoSoFtTuNnElSeRvErId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mIcRoSoFtTuNnElSiTeS/mIcRoSoFtTuNnElSiTeId/mIcRoSoFtTuNnElSeRvErS/mIcRoSoFtTuNnElSeRvErId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerIDInsensitively(v.Input)
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

		if actual.MicrosoftTunnelServerId != v.Expected.MicrosoftTunnelServerId {
			t.Fatalf("Expected %q but got %q for MicrosoftTunnelServerId", v.Expected.MicrosoftTunnelServerId, actual.MicrosoftTunnelServerId)
		}

	}
}

func TestSegmentsForDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId(t *testing.T) {
	segments := DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
