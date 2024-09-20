package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementServiceNowConnectionId{}

func TestNewDeviceManagementServiceNowConnectionID(t *testing.T) {
	id := NewDeviceManagementServiceNowConnectionID("serviceNowConnectionId")

	if id.ServiceNowConnectionId != "serviceNowConnectionId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServiceNowConnectionId'", id.ServiceNowConnectionId, "serviceNowConnectionId")
	}
}

func TestFormatDeviceManagementServiceNowConnectionID(t *testing.T) {
	actual := NewDeviceManagementServiceNowConnectionID("serviceNowConnectionId").ID()
	expected := "/deviceManagement/serviceNowConnections/serviceNowConnectionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementServiceNowConnectionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementServiceNowConnectionId
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
			Input: "/deviceManagement/serviceNowConnections",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/serviceNowConnections/serviceNowConnectionId",
			Expected: &DeviceManagementServiceNowConnectionId{
				ServiceNowConnectionId: "serviceNowConnectionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/serviceNowConnections/serviceNowConnectionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementServiceNowConnectionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ServiceNowConnectionId != v.Expected.ServiceNowConnectionId {
			t.Fatalf("Expected %q but got %q for ServiceNowConnectionId", v.Expected.ServiceNowConnectionId, actual.ServiceNowConnectionId)
		}

	}
}

func TestParseDeviceManagementServiceNowConnectionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementServiceNowConnectionId
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
			Input: "/deviceManagement/serviceNowConnections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/sErViCeNoWcOnNeCtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/serviceNowConnections/serviceNowConnectionId",
			Expected: &DeviceManagementServiceNowConnectionId{
				ServiceNowConnectionId: "serviceNowConnectionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/serviceNowConnections/serviceNowConnectionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/sErViCeNoWcOnNeCtIoNs/sErViCeNoWcOnNeCtIoNiD",
			Expected: &DeviceManagementServiceNowConnectionId{
				ServiceNowConnectionId: "sErViCeNoWcOnNeCtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/sErViCeNoWcOnNeCtIoNs/sErViCeNoWcOnNeCtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementServiceNowConnectionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ServiceNowConnectionId != v.Expected.ServiceNowConnectionId {
			t.Fatalf("Expected %q but got %q for ServiceNowConnectionId", v.Expected.ServiceNowConnectionId, actual.ServiceNowConnectionId)
		}

	}
}

func TestSegmentsForDeviceManagementServiceNowConnectionId(t *testing.T) {
	segments := DeviceManagementServiceNowConnectionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementServiceNowConnectionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
