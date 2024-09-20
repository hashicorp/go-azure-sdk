package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementNdesConnectorId{}

func TestNewDeviceManagementNdesConnectorID(t *testing.T) {
	id := NewDeviceManagementNdesConnectorID("ndesConnectorId")

	if id.NdesConnectorId != "ndesConnectorId" {
		t.Fatalf("Expected %q but got %q for Segment 'NdesConnectorId'", id.NdesConnectorId, "ndesConnectorId")
	}
}

func TestFormatDeviceManagementNdesConnectorID(t *testing.T) {
	actual := NewDeviceManagementNdesConnectorID("ndesConnectorId").ID()
	expected := "/deviceManagement/ndesConnectors/ndesConnectorId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementNdesConnectorID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementNdesConnectorId
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
			Input: "/deviceManagement/ndesConnectors",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/ndesConnectors/ndesConnectorId",
			Expected: &DeviceManagementNdesConnectorId{
				NdesConnectorId: "ndesConnectorId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/ndesConnectors/ndesConnectorId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementNdesConnectorID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.NdesConnectorId != v.Expected.NdesConnectorId {
			t.Fatalf("Expected %q but got %q for NdesConnectorId", v.Expected.NdesConnectorId, actual.NdesConnectorId)
		}

	}
}

func TestParseDeviceManagementNdesConnectorIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementNdesConnectorId
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
			Input: "/deviceManagement/ndesConnectors",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/nDeScOnNeCtOrS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/ndesConnectors/ndesConnectorId",
			Expected: &DeviceManagementNdesConnectorId{
				NdesConnectorId: "ndesConnectorId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/ndesConnectors/ndesConnectorId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/nDeScOnNeCtOrS/nDeScOnNeCtOrId",
			Expected: &DeviceManagementNdesConnectorId{
				NdesConnectorId: "nDeScOnNeCtOrId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/nDeScOnNeCtOrS/nDeScOnNeCtOrId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementNdesConnectorIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.NdesConnectorId != v.Expected.NdesConnectorId {
			t.Fatalf("Expected %q but got %q for NdesConnectorId", v.Expected.NdesConnectorId, actual.NdesConnectorId)
		}

	}
}

func TestSegmentsForDeviceManagementNdesConnectorId(t *testing.T) {
	segments := DeviceManagementNdesConnectorId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementNdesConnectorId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
