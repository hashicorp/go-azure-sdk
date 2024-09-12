package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementResourceOperationId{}

func TestNewDeviceManagementResourceOperationID(t *testing.T) {
	id := NewDeviceManagementResourceOperationID("resourceOperationIdValue")

	if id.ResourceOperationId != "resourceOperationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceOperationId'", id.ResourceOperationId, "resourceOperationIdValue")
	}
}

func TestFormatDeviceManagementResourceOperationID(t *testing.T) {
	actual := NewDeviceManagementResourceOperationID("resourceOperationIdValue").ID()
	expected := "/deviceManagement/resourceOperations/resourceOperationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementResourceOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementResourceOperationId
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
			Input: "/deviceManagement/resourceOperations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/resourceOperations/resourceOperationIdValue",
			Expected: &DeviceManagementResourceOperationId{
				ResourceOperationId: "resourceOperationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/resourceOperations/resourceOperationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementResourceOperationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ResourceOperationId != v.Expected.ResourceOperationId {
			t.Fatalf("Expected %q but got %q for ResourceOperationId", v.Expected.ResourceOperationId, actual.ResourceOperationId)
		}

	}
}

func TestParseDeviceManagementResourceOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementResourceOperationId
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
			Input: "/deviceManagement/resourceOperations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEsOuRcEoPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/resourceOperations/resourceOperationIdValue",
			Expected: &DeviceManagementResourceOperationId{
				ResourceOperationId: "resourceOperationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/resourceOperations/resourceOperationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEsOuRcEoPeRaTiOnS/rEsOuRcEoPeRaTiOnIdVaLuE",
			Expected: &DeviceManagementResourceOperationId{
				ResourceOperationId: "rEsOuRcEoPeRaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEsOuRcEoPeRaTiOnS/rEsOuRcEoPeRaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementResourceOperationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ResourceOperationId != v.Expected.ResourceOperationId {
			t.Fatalf("Expected %q but got %q for ResourceOperationId", v.Expected.ResourceOperationId, actual.ResourceOperationId)
		}

	}
}

func TestSegmentsForDeviceManagementResourceOperationId(t *testing.T) {
	segments := DeviceManagementResourceOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementResourceOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
