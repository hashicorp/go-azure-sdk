package embeddedsimactivationcodepooldevicestate

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId{}

func TestNewDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateID(t *testing.T) {
	id := NewDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateID("embeddedSIMActivationCodePoolIdValue", "embeddedSIMDeviceStateIdValue")

	if id.EmbeddedSIMActivationCodePoolId != "embeddedSIMActivationCodePoolIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EmbeddedSIMActivationCodePoolId'", id.EmbeddedSIMActivationCodePoolId, "embeddedSIMActivationCodePoolIdValue")
	}

	if id.EmbeddedSIMDeviceStateId != "embeddedSIMDeviceStateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EmbeddedSIMDeviceStateId'", id.EmbeddedSIMDeviceStateId, "embeddedSIMDeviceStateIdValue")
	}
}

func TestFormatDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateID(t *testing.T) {
	actual := NewDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateID("embeddedSIMActivationCodePoolIdValue", "embeddedSIMDeviceStateIdValue").ID()
	expected := "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolIdValue/deviceStates/embeddedSIMDeviceStateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId
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
			Input: "/deviceManagement/embeddedSIMActivationCodePools",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolIdValue/deviceStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolIdValue/deviceStates/embeddedSIMDeviceStateIdValue",
			Expected: &DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId{
				EmbeddedSIMActivationCodePoolId: "embeddedSIMActivationCodePoolIdValue",
				EmbeddedSIMDeviceStateId:        "embeddedSIMDeviceStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolIdValue/deviceStates/embeddedSIMDeviceStateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.EmbeddedSIMActivationCodePoolId != v.Expected.EmbeddedSIMActivationCodePoolId {
			t.Fatalf("Expected %q but got %q for EmbeddedSIMActivationCodePoolId", v.Expected.EmbeddedSIMActivationCodePoolId, actual.EmbeddedSIMActivationCodePoolId)
		}

		if actual.EmbeddedSIMDeviceStateId != v.Expected.EmbeddedSIMDeviceStateId {
			t.Fatalf("Expected %q but got %q for EmbeddedSIMDeviceStateId", v.Expected.EmbeddedSIMDeviceStateId, actual.EmbeddedSIMDeviceStateId)
		}

	}
}

func TestParseDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId
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
			Input: "/deviceManagement/embeddedSIMActivationCodePools",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eMbEdDeDsImAcTiVaTiOnCoDePoOlS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eMbEdDeDsImAcTiVaTiOnCoDePoOlS/eMbEdDeDsImAcTiVaTiOnCoDePoOlIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolIdValue/deviceStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eMbEdDeDsImAcTiVaTiOnCoDePoOlS/eMbEdDeDsImAcTiVaTiOnCoDePoOlIdVaLuE/dEvIcEsTaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolIdValue/deviceStates/embeddedSIMDeviceStateIdValue",
			Expected: &DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId{
				EmbeddedSIMActivationCodePoolId: "embeddedSIMActivationCodePoolIdValue",
				EmbeddedSIMDeviceStateId:        "embeddedSIMDeviceStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolIdValue/deviceStates/embeddedSIMDeviceStateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eMbEdDeDsImAcTiVaTiOnCoDePoOlS/eMbEdDeDsImAcTiVaTiOnCoDePoOlIdVaLuE/dEvIcEsTaTeS/eMbEdDeDsImDeViCeStAtEiDvAlUe",
			Expected: &DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId{
				EmbeddedSIMActivationCodePoolId: "eMbEdDeDsImAcTiVaTiOnCoDePoOlIdVaLuE",
				EmbeddedSIMDeviceStateId:        "eMbEdDeDsImDeViCeStAtEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eMbEdDeDsImAcTiVaTiOnCoDePoOlS/eMbEdDeDsImAcTiVaTiOnCoDePoOlIdVaLuE/dEvIcEsTaTeS/eMbEdDeDsImDeViCeStAtEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.EmbeddedSIMActivationCodePoolId != v.Expected.EmbeddedSIMActivationCodePoolId {
			t.Fatalf("Expected %q but got %q for EmbeddedSIMActivationCodePoolId", v.Expected.EmbeddedSIMActivationCodePoolId, actual.EmbeddedSIMActivationCodePoolId)
		}

		if actual.EmbeddedSIMDeviceStateId != v.Expected.EmbeddedSIMDeviceStateId {
			t.Fatalf("Expected %q but got %q for EmbeddedSIMDeviceStateId", v.Expected.EmbeddedSIMDeviceStateId, actual.EmbeddedSIMDeviceStateId)
		}

	}
}

func TestSegmentsForDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId(t *testing.T) {
	segments := DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
