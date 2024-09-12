package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntentIdUserStateId{}

func TestNewDeviceManagementIntentIdUserStateID(t *testing.T) {
	id := NewDeviceManagementIntentIdUserStateID("deviceManagementIntentIdValue", "deviceManagementIntentUserStateIdValue")

	if id.DeviceManagementIntentId != "deviceManagementIntentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementIntentId'", id.DeviceManagementIntentId, "deviceManagementIntentIdValue")
	}

	if id.DeviceManagementIntentUserStateId != "deviceManagementIntentUserStateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementIntentUserStateId'", id.DeviceManagementIntentUserStateId, "deviceManagementIntentUserStateIdValue")
	}
}

func TestFormatDeviceManagementIntentIdUserStateID(t *testing.T) {
	actual := NewDeviceManagementIntentIdUserStateID("deviceManagementIntentIdValue", "deviceManagementIntentUserStateIdValue").ID()
	expected := "/deviceManagement/intents/deviceManagementIntentIdValue/userStates/deviceManagementIntentUserStateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementIntentIdUserStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIntentIdUserStateId
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
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/userStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/userStates/deviceManagementIntentUserStateIdValue",
			Expected: &DeviceManagementIntentIdUserStateId{
				DeviceManagementIntentId:          "deviceManagementIntentIdValue",
				DeviceManagementIntentUserStateId: "deviceManagementIntentUserStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/userStates/deviceManagementIntentUserStateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIntentIdUserStateID(v.Input)
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

		if actual.DeviceManagementIntentUserStateId != v.Expected.DeviceManagementIntentUserStateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementIntentUserStateId", v.Expected.DeviceManagementIntentUserStateId, actual.DeviceManagementIntentUserStateId)
		}

	}
}

func TestParseDeviceManagementIntentIdUserStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIntentIdUserStateId
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
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/userStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe/uSeRsTaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/userStates/deviceManagementIntentUserStateIdValue",
			Expected: &DeviceManagementIntentIdUserStateId{
				DeviceManagementIntentId:          "deviceManagementIntentIdValue",
				DeviceManagementIntentUserStateId: "deviceManagementIntentUserStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/userStates/deviceManagementIntentUserStateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe/uSeRsTaTeS/dEvIcEmAnAgEmEnTiNtEnTuSeRsTaTeIdVaLuE",
			Expected: &DeviceManagementIntentIdUserStateId{
				DeviceManagementIntentId:          "dEvIcEmAnAgEmEnTiNtEnTiDvAlUe",
				DeviceManagementIntentUserStateId: "dEvIcEmAnAgEmEnTiNtEnTuSeRsTaTeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe/uSeRsTaTeS/dEvIcEmAnAgEmEnTiNtEnTuSeRsTaTeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIntentIdUserStateIDInsensitively(v.Input)
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

		if actual.DeviceManagementIntentUserStateId != v.Expected.DeviceManagementIntentUserStateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementIntentUserStateId", v.Expected.DeviceManagementIntentUserStateId, actual.DeviceManagementIntentUserStateId)
		}

	}
}

func TestSegmentsForDeviceManagementIntentIdUserStateId(t *testing.T) {
	segments := DeviceManagementIntentIdUserStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementIntentIdUserStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
