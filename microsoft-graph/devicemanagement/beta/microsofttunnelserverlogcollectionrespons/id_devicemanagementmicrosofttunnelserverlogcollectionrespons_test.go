package microsofttunnelserverlogcollectionrespons

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMicrosoftTunnelServerLogCollectionResponsId{}

func TestNewDeviceManagementMicrosoftTunnelServerLogCollectionResponsID(t *testing.T) {
	id := NewDeviceManagementMicrosoftTunnelServerLogCollectionResponsID("microsoftTunnelServerLogCollectionResponseIdValue")

	if id.MicrosoftTunnelServerLogCollectionResponseId != "microsoftTunnelServerLogCollectionResponseIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MicrosoftTunnelServerLogCollectionResponseId'", id.MicrosoftTunnelServerLogCollectionResponseId, "microsoftTunnelServerLogCollectionResponseIdValue")
	}
}

func TestFormatDeviceManagementMicrosoftTunnelServerLogCollectionResponsID(t *testing.T) {
	actual := NewDeviceManagementMicrosoftTunnelServerLogCollectionResponsID("microsoftTunnelServerLogCollectionResponseIdValue").ID()
	expected := "/deviceManagement/microsoftTunnelServerLogCollectionResponses/microsoftTunnelServerLogCollectionResponseIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementMicrosoftTunnelServerLogCollectionResponsID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMicrosoftTunnelServerLogCollectionResponsId
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
			Input: "/deviceManagement/microsoftTunnelServerLogCollectionResponses/microsoftTunnelServerLogCollectionResponseIdValue",
			Expected: &DeviceManagementMicrosoftTunnelServerLogCollectionResponsId{
				MicrosoftTunnelServerLogCollectionResponseId: "microsoftTunnelServerLogCollectionResponseIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/microsoftTunnelServerLogCollectionResponses/microsoftTunnelServerLogCollectionResponseIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMicrosoftTunnelServerLogCollectionResponsID(v.Input)
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

func TestParseDeviceManagementMicrosoftTunnelServerLogCollectionResponsIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMicrosoftTunnelServerLogCollectionResponsId
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
			Input: "/deviceManagement/microsoftTunnelServerLogCollectionResponses/microsoftTunnelServerLogCollectionResponseIdValue",
			Expected: &DeviceManagementMicrosoftTunnelServerLogCollectionResponsId{
				MicrosoftTunnelServerLogCollectionResponseId: "microsoftTunnelServerLogCollectionResponseIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/microsoftTunnelServerLogCollectionResponses/microsoftTunnelServerLogCollectionResponseIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mIcRoSoFtTuNnElSeRvErLoGcOlLeCtIoNrEsPoNsEs/mIcRoSoFtTuNnElSeRvErLoGcOlLeCtIoNrEsPoNsEiDvAlUe",
			Expected: &DeviceManagementMicrosoftTunnelServerLogCollectionResponsId{
				MicrosoftTunnelServerLogCollectionResponseId: "mIcRoSoFtTuNnElSeRvErLoGcOlLeCtIoNrEsPoNsEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mIcRoSoFtTuNnElSeRvErLoGcOlLeCtIoNrEsPoNsEs/mIcRoSoFtTuNnElSeRvErLoGcOlLeCtIoNrEsPoNsEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMicrosoftTunnelServerLogCollectionResponsIDInsensitively(v.Input)
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

func TestSegmentsForDeviceManagementMicrosoftTunnelServerLogCollectionResponsId(t *testing.T) {
	segments := DeviceManagementMicrosoftTunnelServerLogCollectionResponsId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementMicrosoftTunnelServerLogCollectionResponsId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
