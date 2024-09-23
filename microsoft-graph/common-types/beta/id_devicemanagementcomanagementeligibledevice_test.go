package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComanagementEligibleDeviceId{}

func TestNewDeviceManagementComanagementEligibleDeviceID(t *testing.T) {
	id := NewDeviceManagementComanagementEligibleDeviceID("comanagementEligibleDeviceId")

	if id.ComanagementEligibleDeviceId != "comanagementEligibleDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'ComanagementEligibleDeviceId'", id.ComanagementEligibleDeviceId, "comanagementEligibleDeviceId")
	}
}

func TestFormatDeviceManagementComanagementEligibleDeviceID(t *testing.T) {
	actual := NewDeviceManagementComanagementEligibleDeviceID("comanagementEligibleDeviceId").ID()
	expected := "/deviceManagement/comanagementEligibleDevices/comanagementEligibleDeviceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementComanagementEligibleDeviceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementComanagementEligibleDeviceId
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
			Input: "/deviceManagement/comanagementEligibleDevices",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/comanagementEligibleDevices/comanagementEligibleDeviceId",
			Expected: &DeviceManagementComanagementEligibleDeviceId{
				ComanagementEligibleDeviceId: "comanagementEligibleDeviceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/comanagementEligibleDevices/comanagementEligibleDeviceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementComanagementEligibleDeviceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ComanagementEligibleDeviceId != v.Expected.ComanagementEligibleDeviceId {
			t.Fatalf("Expected %q but got %q for ComanagementEligibleDeviceId", v.Expected.ComanagementEligibleDeviceId, actual.ComanagementEligibleDeviceId)
		}

	}
}

func TestParseDeviceManagementComanagementEligibleDeviceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementComanagementEligibleDeviceId
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
			Input: "/deviceManagement/comanagementEligibleDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEmEnTeLiGiBlEdEvIcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/comanagementEligibleDevices/comanagementEligibleDeviceId",
			Expected: &DeviceManagementComanagementEligibleDeviceId{
				ComanagementEligibleDeviceId: "comanagementEligibleDeviceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/comanagementEligibleDevices/comanagementEligibleDeviceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEmEnTeLiGiBlEdEvIcEs/cOmAnAgEmEnTeLiGiBlEdEvIcEiD",
			Expected: &DeviceManagementComanagementEligibleDeviceId{
				ComanagementEligibleDeviceId: "cOmAnAgEmEnTeLiGiBlEdEvIcEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEmEnTeLiGiBlEdEvIcEs/cOmAnAgEmEnTeLiGiBlEdEvIcEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementComanagementEligibleDeviceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ComanagementEligibleDeviceId != v.Expected.ComanagementEligibleDeviceId {
			t.Fatalf("Expected %q but got %q for ComanagementEligibleDeviceId", v.Expected.ComanagementEligibleDeviceId, actual.ComanagementEligibleDeviceId)
		}

	}
}

func TestSegmentsForDeviceManagementComanagementEligibleDeviceId(t *testing.T) {
	segments := DeviceManagementComanagementEligibleDeviceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementComanagementEligibleDeviceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
