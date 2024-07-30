package autopiloteventpolicystatusdetail

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAutopilotEventIdPolicyStatusDetailId{}

func TestNewDeviceManagementAutopilotEventIdPolicyStatusDetailID(t *testing.T) {
	id := NewDeviceManagementAutopilotEventIdPolicyStatusDetailID("deviceManagementAutopilotEventIdValue", "deviceManagementAutopilotPolicyStatusDetailIdValue")

	if id.DeviceManagementAutopilotEventId != "deviceManagementAutopilotEventIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementAutopilotEventId'", id.DeviceManagementAutopilotEventId, "deviceManagementAutopilotEventIdValue")
	}

	if id.DeviceManagementAutopilotPolicyStatusDetailId != "deviceManagementAutopilotPolicyStatusDetailIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementAutopilotPolicyStatusDetailId'", id.DeviceManagementAutopilotPolicyStatusDetailId, "deviceManagementAutopilotPolicyStatusDetailIdValue")
	}
}

func TestFormatDeviceManagementAutopilotEventIdPolicyStatusDetailID(t *testing.T) {
	actual := NewDeviceManagementAutopilotEventIdPolicyStatusDetailID("deviceManagementAutopilotEventIdValue", "deviceManagementAutopilotPolicyStatusDetailIdValue").ID()
	expected := "/deviceManagement/autopilotEvents/deviceManagementAutopilotEventIdValue/policyStatusDetails/deviceManagementAutopilotPolicyStatusDetailIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementAutopilotEventIdPolicyStatusDetailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementAutopilotEventIdPolicyStatusDetailId
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
			Input: "/deviceManagement/autopilotEvents",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/autopilotEvents/deviceManagementAutopilotEventIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/autopilotEvents/deviceManagementAutopilotEventIdValue/policyStatusDetails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/autopilotEvents/deviceManagementAutopilotEventIdValue/policyStatusDetails/deviceManagementAutopilotPolicyStatusDetailIdValue",
			Expected: &DeviceManagementAutopilotEventIdPolicyStatusDetailId{
				DeviceManagementAutopilotEventId:              "deviceManagementAutopilotEventIdValue",
				DeviceManagementAutopilotPolicyStatusDetailId: "deviceManagementAutopilotPolicyStatusDetailIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/autopilotEvents/deviceManagementAutopilotEventIdValue/policyStatusDetails/deviceManagementAutopilotPolicyStatusDetailIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementAutopilotEventIdPolicyStatusDetailID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementAutopilotEventId != v.Expected.DeviceManagementAutopilotEventId {
			t.Fatalf("Expected %q but got %q for DeviceManagementAutopilotEventId", v.Expected.DeviceManagementAutopilotEventId, actual.DeviceManagementAutopilotEventId)
		}

		if actual.DeviceManagementAutopilotPolicyStatusDetailId != v.Expected.DeviceManagementAutopilotPolicyStatusDetailId {
			t.Fatalf("Expected %q but got %q for DeviceManagementAutopilotPolicyStatusDetailId", v.Expected.DeviceManagementAutopilotPolicyStatusDetailId, actual.DeviceManagementAutopilotPolicyStatusDetailId)
		}

	}
}

func TestParseDeviceManagementAutopilotEventIdPolicyStatusDetailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementAutopilotEventIdPolicyStatusDetailId
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
			Input: "/deviceManagement/autopilotEvents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aUtOpIlOtEvEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/autopilotEvents/deviceManagementAutopilotEventIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aUtOpIlOtEvEnTs/dEvIcEmAnAgEmEnTaUtOpIlOtEvEnTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/autopilotEvents/deviceManagementAutopilotEventIdValue/policyStatusDetails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aUtOpIlOtEvEnTs/dEvIcEmAnAgEmEnTaUtOpIlOtEvEnTiDvAlUe/pOlIcYsTaTuSdEtAiLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/autopilotEvents/deviceManagementAutopilotEventIdValue/policyStatusDetails/deviceManagementAutopilotPolicyStatusDetailIdValue",
			Expected: &DeviceManagementAutopilotEventIdPolicyStatusDetailId{
				DeviceManagementAutopilotEventId:              "deviceManagementAutopilotEventIdValue",
				DeviceManagementAutopilotPolicyStatusDetailId: "deviceManagementAutopilotPolicyStatusDetailIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/autopilotEvents/deviceManagementAutopilotEventIdValue/policyStatusDetails/deviceManagementAutopilotPolicyStatusDetailIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aUtOpIlOtEvEnTs/dEvIcEmAnAgEmEnTaUtOpIlOtEvEnTiDvAlUe/pOlIcYsTaTuSdEtAiLs/dEvIcEmAnAgEmEnTaUtOpIlOtPoLiCyStAtUsDeTaIlIdVaLuE",
			Expected: &DeviceManagementAutopilotEventIdPolicyStatusDetailId{
				DeviceManagementAutopilotEventId:              "dEvIcEmAnAgEmEnTaUtOpIlOtEvEnTiDvAlUe",
				DeviceManagementAutopilotPolicyStatusDetailId: "dEvIcEmAnAgEmEnTaUtOpIlOtPoLiCyStAtUsDeTaIlIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aUtOpIlOtEvEnTs/dEvIcEmAnAgEmEnTaUtOpIlOtEvEnTiDvAlUe/pOlIcYsTaTuSdEtAiLs/dEvIcEmAnAgEmEnTaUtOpIlOtPoLiCyStAtUsDeTaIlIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementAutopilotEventIdPolicyStatusDetailIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementAutopilotEventId != v.Expected.DeviceManagementAutopilotEventId {
			t.Fatalf("Expected %q but got %q for DeviceManagementAutopilotEventId", v.Expected.DeviceManagementAutopilotEventId, actual.DeviceManagementAutopilotEventId)
		}

		if actual.DeviceManagementAutopilotPolicyStatusDetailId != v.Expected.DeviceManagementAutopilotPolicyStatusDetailId {
			t.Fatalf("Expected %q but got %q for DeviceManagementAutopilotPolicyStatusDetailId", v.Expected.DeviceManagementAutopilotPolicyStatusDetailId, actual.DeviceManagementAutopilotPolicyStatusDetailId)
		}

	}
}

func TestSegmentsForDeviceManagementAutopilotEventIdPolicyStatusDetailId(t *testing.T) {
	segments := DeviceManagementAutopilotEventIdPolicyStatusDetailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementAutopilotEventIdPolicyStatusDetailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
