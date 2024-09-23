package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeManagedDeviceIdLogCollectionRequestId{}

func TestNewMeManagedDeviceIdLogCollectionRequestID(t *testing.T) {
	id := NewMeManagedDeviceIdLogCollectionRequestID("managedDeviceId", "deviceLogCollectionResponseId")

	if id.ManagedDeviceId != "managedDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceId")
	}

	if id.DeviceLogCollectionResponseId != "deviceLogCollectionResponseId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceLogCollectionResponseId'", id.DeviceLogCollectionResponseId, "deviceLogCollectionResponseId")
	}
}

func TestFormatMeManagedDeviceIdLogCollectionRequestID(t *testing.T) {
	actual := NewMeManagedDeviceIdLogCollectionRequestID("managedDeviceId", "deviceLogCollectionResponseId").ID()
	expected := "/me/managedDevices/managedDeviceId/logCollectionRequests/deviceLogCollectionResponseId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeManagedDeviceIdLogCollectionRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeManagedDeviceIdLogCollectionRequestId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceId/logCollectionRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/managedDevices/managedDeviceId/logCollectionRequests/deviceLogCollectionResponseId",
			Expected: &MeManagedDeviceIdLogCollectionRequestId{
				ManagedDeviceId:               "managedDeviceId",
				DeviceLogCollectionResponseId: "deviceLogCollectionResponseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/managedDevices/managedDeviceId/logCollectionRequests/deviceLogCollectionResponseId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeManagedDeviceIdLogCollectionRequestID(v.Input)
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

func TestParseMeManagedDeviceIdLogCollectionRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeManagedDeviceIdLogCollectionRequestId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceId/logCollectionRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeId/lOgCoLlEcTiOnReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/managedDevices/managedDeviceId/logCollectionRequests/deviceLogCollectionResponseId",
			Expected: &MeManagedDeviceIdLogCollectionRequestId{
				ManagedDeviceId:               "managedDeviceId",
				DeviceLogCollectionResponseId: "deviceLogCollectionResponseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/managedDevices/managedDeviceId/logCollectionRequests/deviceLogCollectionResponseId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeId/lOgCoLlEcTiOnReQuEsTs/dEvIcElOgCoLlEcTiOnReSpOnSeId",
			Expected: &MeManagedDeviceIdLogCollectionRequestId{
				ManagedDeviceId:               "mAnAgEdDeViCeId",
				DeviceLogCollectionResponseId: "dEvIcElOgCoLlEcTiOnReSpOnSeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeId/lOgCoLlEcTiOnReQuEsTs/dEvIcElOgCoLlEcTiOnReSpOnSeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeManagedDeviceIdLogCollectionRequestIDInsensitively(v.Input)
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

func TestSegmentsForMeManagedDeviceIdLogCollectionRequestId(t *testing.T) {
	segments := MeManagedDeviceIdLogCollectionRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeManagedDeviceIdLogCollectionRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
