package intentassignment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntentIdAssignmentId{}

func TestNewDeviceManagementIntentIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementIntentIdAssignmentID("deviceManagementIntentIdValue", "deviceManagementIntentAssignmentIdValue")

	if id.DeviceManagementIntentId != "deviceManagementIntentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementIntentId'", id.DeviceManagementIntentId, "deviceManagementIntentIdValue")
	}

	if id.DeviceManagementIntentAssignmentId != "deviceManagementIntentAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementIntentAssignmentId'", id.DeviceManagementIntentAssignmentId, "deviceManagementIntentAssignmentIdValue")
	}
}

func TestFormatDeviceManagementIntentIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementIntentIdAssignmentID("deviceManagementIntentIdValue", "deviceManagementIntentAssignmentIdValue").ID()
	expected := "/deviceManagement/intents/deviceManagementIntentIdValue/assignments/deviceManagementIntentAssignmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementIntentIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIntentIdAssignmentId
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
			Input: "/deviceManagement/intents",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/assignments/deviceManagementIntentAssignmentIdValue",
			Expected: &DeviceManagementIntentIdAssignmentId{
				DeviceManagementIntentId:           "deviceManagementIntentIdValue",
				DeviceManagementIntentAssignmentId: "deviceManagementIntentAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/assignments/deviceManagementIntentAssignmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIntentIdAssignmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementIntentId != v.Expected.DeviceManagementIntentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementIntentId", v.Expected.DeviceManagementIntentId, actual.DeviceManagementIntentId)
		}

		if actual.DeviceManagementIntentAssignmentId != v.Expected.DeviceManagementIntentAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementIntentAssignmentId", v.Expected.DeviceManagementIntentAssignmentId, actual.DeviceManagementIntentAssignmentId)
		}

	}
}

func TestParseDeviceManagementIntentIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIntentIdAssignmentId
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
			Input: "/deviceManagement/intents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/assignments/deviceManagementIntentAssignmentIdValue",
			Expected: &DeviceManagementIntentIdAssignmentId{
				DeviceManagementIntentId:           "deviceManagementIntentIdValue",
				DeviceManagementIntentAssignmentId: "deviceManagementIntentAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/assignments/deviceManagementIntentAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe/aSsIgNmEnTs/dEvIcEmAnAgEmEnTiNtEnTaSsIgNmEnTiDvAlUe",
			Expected: &DeviceManagementIntentIdAssignmentId{
				DeviceManagementIntentId:           "dEvIcEmAnAgEmEnTiNtEnTiDvAlUe",
				DeviceManagementIntentAssignmentId: "dEvIcEmAnAgEmEnTiNtEnTaSsIgNmEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe/aSsIgNmEnTs/dEvIcEmAnAgEmEnTiNtEnTaSsIgNmEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIntentIdAssignmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementIntentId != v.Expected.DeviceManagementIntentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementIntentId", v.Expected.DeviceManagementIntentId, actual.DeviceManagementIntentId)
		}

		if actual.DeviceManagementIntentAssignmentId != v.Expected.DeviceManagementIntentAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementIntentAssignmentId", v.Expected.DeviceManagementIntentAssignmentId, actual.DeviceManagementIntentAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementIntentIdAssignmentId(t *testing.T) {
	segments := DeviceManagementIntentIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementIntentIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
