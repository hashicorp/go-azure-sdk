package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMicrosoftTunnelServerLogCollectionResponseId{}

func TestNewDeviceManagementMicrosoftTunnelServerLogCollectionResponseID(t *testing.T) {
	id := NewDeviceManagementMicrosoftTunnelServerLogCollectionResponseID("microsoftTunnelServerLogCollectionResponseId")

	if id.MicrosoftTunnelServerLogCollectionResponseId != "microsoftTunnelServerLogCollectionResponseId" {
		t.Fatalf("Expected %q but got %q for Segment 'MicrosoftTunnelServerLogCollectionResponseId'", id.MicrosoftTunnelServerLogCollectionResponseId, "microsoftTunnelServerLogCollectionResponseId")
	}
}

func TestFormatDeviceManagementMicrosoftTunnelServerLogCollectionResponseID(t *testing.T) {
	actual := NewDeviceManagementMicrosoftTunnelServerLogCollectionResponseID("microsoftTunnelServerLogCollectionResponseId").ID()
	expected := "/deviceManagement/microsoftTunnelServerLogCollectionResponses/microsoftTunnelServerLogCollectionResponseId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementMicrosoftTunnelServerLogCollectionResponseID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMicrosoftTunnelServerLogCollectionResponseId
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
			Input: "/deviceManagement/microsoftTunnelServerLogCollectionResponses",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/microsoftTunnelServerLogCollectionResponses/microsoftTunnelServerLogCollectionResponseId",
			Expected: &DeviceManagementMicrosoftTunnelServerLogCollectionResponseId{
				MicrosoftTunnelServerLogCollectionResponseId: "microsoftTunnelServerLogCollectionResponseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/microsoftTunnelServerLogCollectionResponses/microsoftTunnelServerLogCollectionResponseId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMicrosoftTunnelServerLogCollectionResponseID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MicrosoftTunnelServerLogCollectionResponseId != v.Expected.MicrosoftTunnelServerLogCollectionResponseId {
			t.Fatalf("Expected %q but got %q for MicrosoftTunnelServerLogCollectionResponseId", v.Expected.MicrosoftTunnelServerLogCollectionResponseId, actual.MicrosoftTunnelServerLogCollectionResponseId)
		}

	}
}

func TestParseDeviceManagementMicrosoftTunnelServerLogCollectionResponseIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMicrosoftTunnelServerLogCollectionResponseId
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
			Input: "/deviceManagement/microsoftTunnelServerLogCollectionResponses",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mIcRoSoFtTuNnElSeRvErLoGcOlLeCtIoNrEsPoNsEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/microsoftTunnelServerLogCollectionResponses/microsoftTunnelServerLogCollectionResponseId",
			Expected: &DeviceManagementMicrosoftTunnelServerLogCollectionResponseId{
				MicrosoftTunnelServerLogCollectionResponseId: "microsoftTunnelServerLogCollectionResponseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/microsoftTunnelServerLogCollectionResponses/microsoftTunnelServerLogCollectionResponseId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mIcRoSoFtTuNnElSeRvErLoGcOlLeCtIoNrEsPoNsEs/mIcRoSoFtTuNnElSeRvErLoGcOlLeCtIoNrEsPoNsEiD",
			Expected: &DeviceManagementMicrosoftTunnelServerLogCollectionResponseId{
				MicrosoftTunnelServerLogCollectionResponseId: "mIcRoSoFtTuNnElSeRvErLoGcOlLeCtIoNrEsPoNsEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mIcRoSoFtTuNnElSeRvErLoGcOlLeCtIoNrEsPoNsEs/mIcRoSoFtTuNnElSeRvErLoGcOlLeCtIoNrEsPoNsEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMicrosoftTunnelServerLogCollectionResponseIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MicrosoftTunnelServerLogCollectionResponseId != v.Expected.MicrosoftTunnelServerLogCollectionResponseId {
			t.Fatalf("Expected %q but got %q for MicrosoftTunnelServerLogCollectionResponseId", v.Expected.MicrosoftTunnelServerLogCollectionResponseId, actual.MicrosoftTunnelServerLogCollectionResponseId)
		}

	}
}

func TestSegmentsForDeviceManagementMicrosoftTunnelServerLogCollectionResponseId(t *testing.T) {
	segments := DeviceManagementMicrosoftTunnelServerLogCollectionResponseId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementMicrosoftTunnelServerLogCollectionResponseId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
