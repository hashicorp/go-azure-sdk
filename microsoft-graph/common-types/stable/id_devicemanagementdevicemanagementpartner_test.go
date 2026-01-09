package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceManagementPartnerId{}

func TestNewDeviceManagementDeviceManagementPartnerID(t *testing.T) {
	id := NewDeviceManagementDeviceManagementPartnerID("deviceManagementPartnerId")

	if id.DeviceManagementPartnerId != "deviceManagementPartnerId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementPartnerId'", id.DeviceManagementPartnerId, "deviceManagementPartnerId")
	}
}

func TestFormatDeviceManagementDeviceManagementPartnerID(t *testing.T) {
	actual := NewDeviceManagementDeviceManagementPartnerID("deviceManagementPartnerId").ID()
	expected := "/deviceManagement/deviceManagementPartners/deviceManagementPartnerId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceManagementPartnerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceManagementPartnerId
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
			Input: "/deviceManagement/deviceManagementPartners",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceManagementPartners/deviceManagementPartnerId",
			Expected: &DeviceManagementDeviceManagementPartnerId{
				DeviceManagementPartnerId: "deviceManagementPartnerId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceManagementPartners/deviceManagementPartnerId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceManagementPartnerID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementPartnerId != v.Expected.DeviceManagementPartnerId {
			t.Fatalf("Expected %q but got %q for DeviceManagementPartnerId", v.Expected.DeviceManagementPartnerId, actual.DeviceManagementPartnerId)
		}

	}
}

func TestParseDeviceManagementDeviceManagementPartnerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceManagementPartnerId
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
			Input: "/deviceManagement/deviceManagementPartners",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEmAnAgEmEnTpArTnErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceManagementPartners/deviceManagementPartnerId",
			Expected: &DeviceManagementDeviceManagementPartnerId{
				DeviceManagementPartnerId: "deviceManagementPartnerId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceManagementPartners/deviceManagementPartnerId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEmAnAgEmEnTpArTnErS/dEvIcEmAnAgEmEnTpArTnErId",
			Expected: &DeviceManagementDeviceManagementPartnerId{
				DeviceManagementPartnerId: "dEvIcEmAnAgEmEnTpArTnErId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEmAnAgEmEnTpArTnErS/dEvIcEmAnAgEmEnTpArTnErId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceManagementPartnerIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementPartnerId != v.Expected.DeviceManagementPartnerId {
			t.Fatalf("Expected %q but got %q for DeviceManagementPartnerId", v.Expected.DeviceManagementPartnerId, actual.DeviceManagementPartnerId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceManagementPartnerId(t *testing.T) {
	segments := DeviceManagementDeviceManagementPartnerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceManagementPartnerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
