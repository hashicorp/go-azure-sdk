package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDomainJoinConnectorId{}

func TestNewDeviceManagementDomainJoinConnectorID(t *testing.T) {
	id := NewDeviceManagementDomainJoinConnectorID("deviceManagementDomainJoinConnectorId")

	if id.DeviceManagementDomainJoinConnectorId != "deviceManagementDomainJoinConnectorId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementDomainJoinConnectorId'", id.DeviceManagementDomainJoinConnectorId, "deviceManagementDomainJoinConnectorId")
	}
}

func TestFormatDeviceManagementDomainJoinConnectorID(t *testing.T) {
	actual := NewDeviceManagementDomainJoinConnectorID("deviceManagementDomainJoinConnectorId").ID()
	expected := "/deviceManagement/domainJoinConnectors/deviceManagementDomainJoinConnectorId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDomainJoinConnectorID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDomainJoinConnectorId
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
			Input: "/deviceManagement/domainJoinConnectors",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/domainJoinConnectors/deviceManagementDomainJoinConnectorId",
			Expected: &DeviceManagementDomainJoinConnectorId{
				DeviceManagementDomainJoinConnectorId: "deviceManagementDomainJoinConnectorId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/domainJoinConnectors/deviceManagementDomainJoinConnectorId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDomainJoinConnectorID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementDomainJoinConnectorId != v.Expected.DeviceManagementDomainJoinConnectorId {
			t.Fatalf("Expected %q but got %q for DeviceManagementDomainJoinConnectorId", v.Expected.DeviceManagementDomainJoinConnectorId, actual.DeviceManagementDomainJoinConnectorId)
		}

	}
}

func TestParseDeviceManagementDomainJoinConnectorIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDomainJoinConnectorId
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
			Input: "/deviceManagement/domainJoinConnectors",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dOmAiNjOiNcOnNeCtOrS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/domainJoinConnectors/deviceManagementDomainJoinConnectorId",
			Expected: &DeviceManagementDomainJoinConnectorId{
				DeviceManagementDomainJoinConnectorId: "deviceManagementDomainJoinConnectorId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/domainJoinConnectors/deviceManagementDomainJoinConnectorId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dOmAiNjOiNcOnNeCtOrS/dEvIcEmAnAgEmEnTdOmAiNjOiNcOnNeCtOrId",
			Expected: &DeviceManagementDomainJoinConnectorId{
				DeviceManagementDomainJoinConnectorId: "dEvIcEmAnAgEmEnTdOmAiNjOiNcOnNeCtOrId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dOmAiNjOiNcOnNeCtOrS/dEvIcEmAnAgEmEnTdOmAiNjOiNcOnNeCtOrId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDomainJoinConnectorIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementDomainJoinConnectorId != v.Expected.DeviceManagementDomainJoinConnectorId {
			t.Fatalf("Expected %q but got %q for DeviceManagementDomainJoinConnectorId", v.Expected.DeviceManagementDomainJoinConnectorId, actual.DeviceManagementDomainJoinConnectorId)
		}

	}
}

func TestSegmentsForDeviceManagementDomainJoinConnectorId(t *testing.T) {
	segments := DeviceManagementDomainJoinConnectorId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDomainJoinConnectorId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
