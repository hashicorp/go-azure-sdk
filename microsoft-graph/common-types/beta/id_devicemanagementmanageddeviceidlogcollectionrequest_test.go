package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceIdLogCollectionRequestId{}

func TestNewDeviceManagementManagedDeviceIdLogCollectionRequestID(t *testing.T) {
	id := NewDeviceManagementManagedDeviceIdLogCollectionRequestID("managedDeviceId", "deviceLogCollectionResponseId")

	if id.ManagedDeviceId != "managedDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceId")
	}

	if id.DeviceLogCollectionResponseId != "deviceLogCollectionResponseId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceLogCollectionResponseId'", id.DeviceLogCollectionResponseId, "deviceLogCollectionResponseId")
	}
}

func TestFormatDeviceManagementManagedDeviceIdLogCollectionRequestID(t *testing.T) {
	actual := NewDeviceManagementManagedDeviceIdLogCollectionRequestID("managedDeviceId", "deviceLogCollectionResponseId").ID()
	expected := "/deviceManagement/managedDevices/managedDeviceId/logCollectionRequests/deviceLogCollectionResponseId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementManagedDeviceIdLogCollectionRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementManagedDeviceIdLogCollectionRequestId
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
			Input: "/deviceManagement/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/managedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/managedDevices/managedDeviceId/logCollectionRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/managedDevices/managedDeviceId/logCollectionRequests/deviceLogCollectionResponseId",
			Expected: &DeviceManagementManagedDeviceIdLogCollectionRequestId{
				ManagedDeviceId:               "managedDeviceId",
				DeviceLogCollectionResponseId: "deviceLogCollectionResponseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/managedDevices/managedDeviceId/logCollectionRequests/deviceLogCollectionResponseId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementManagedDeviceIdLogCollectionRequestID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

		if actual.DeviceLogCollectionResponseId != v.Expected.DeviceLogCollectionResponseId {
			t.Fatalf("Expected %q but got %q for DeviceLogCollectionResponseId", v.Expected.DeviceLogCollectionResponseId, actual.DeviceLogCollectionResponseId)
		}

	}
}

func TestParseDeviceManagementManagedDeviceIdLogCollectionRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementManagedDeviceIdLogCollectionRequestId
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
			Input: "/deviceManagement/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/managedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/managedDevices/managedDeviceId/logCollectionRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeId/lOgCoLlEcTiOnReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/managedDevices/managedDeviceId/logCollectionRequests/deviceLogCollectionResponseId",
			Expected: &DeviceManagementManagedDeviceIdLogCollectionRequestId{
				ManagedDeviceId:               "managedDeviceId",
				DeviceLogCollectionResponseId: "deviceLogCollectionResponseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/managedDevices/managedDeviceId/logCollectionRequests/deviceLogCollectionResponseId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeId/lOgCoLlEcTiOnReQuEsTs/dEvIcElOgCoLlEcTiOnReSpOnSeId",
			Expected: &DeviceManagementManagedDeviceIdLogCollectionRequestId{
				ManagedDeviceId:               "mAnAgEdDeViCeId",
				DeviceLogCollectionResponseId: "dEvIcElOgCoLlEcTiOnReSpOnSeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeId/lOgCoLlEcTiOnReQuEsTs/dEvIcElOgCoLlEcTiOnReSpOnSeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementManagedDeviceIdLogCollectionRequestIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

		if actual.DeviceLogCollectionResponseId != v.Expected.DeviceLogCollectionResponseId {
			t.Fatalf("Expected %q but got %q for DeviceLogCollectionResponseId", v.Expected.DeviceLogCollectionResponseId, actual.DeviceLogCollectionResponseId)
		}

	}
}

func TestSegmentsForDeviceManagementManagedDeviceIdLogCollectionRequestId(t *testing.T) {
	segments := DeviceManagementManagedDeviceIdLogCollectionRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementManagedDeviceIdLogCollectionRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
