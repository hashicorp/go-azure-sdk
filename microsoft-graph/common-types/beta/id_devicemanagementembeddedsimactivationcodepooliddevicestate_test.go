package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId{}

func TestNewDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateID(t *testing.T) {
	id := NewDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateID("embeddedSIMActivationCodePoolId", "embeddedSIMDeviceStateId")

	if id.EmbeddedSIMActivationCodePoolId != "embeddedSIMActivationCodePoolId" {
		t.Fatalf("Expected %q but got %q for Segment 'EmbeddedSIMActivationCodePoolId'", id.EmbeddedSIMActivationCodePoolId, "embeddedSIMActivationCodePoolId")
	}

	if id.EmbeddedSIMDeviceStateId != "embeddedSIMDeviceStateId" {
		t.Fatalf("Expected %q but got %q for Segment 'EmbeddedSIMDeviceStateId'", id.EmbeddedSIMDeviceStateId, "embeddedSIMDeviceStateId")
	}
}

func TestFormatDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateID(t *testing.T) {
	actual := NewDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateID("embeddedSIMActivationCodePoolId", "embeddedSIMDeviceStateId").ID()
	expected := "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolId/deviceStates/embeddedSIMDeviceStateId"
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
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolId/deviceStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolId/deviceStates/embeddedSIMDeviceStateId",
			Expected: &DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId{
				EmbeddedSIMActivationCodePoolId: "embeddedSIMActivationCodePoolId",
				EmbeddedSIMDeviceStateId:        "embeddedSIMDeviceStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolId/deviceStates/embeddedSIMDeviceStateId/extra",
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
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eMbEdDeDsImAcTiVaTiOnCoDePoOlS/eMbEdDeDsImAcTiVaTiOnCoDePoOlId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolId/deviceStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eMbEdDeDsImAcTiVaTiOnCoDePoOlS/eMbEdDeDsImAcTiVaTiOnCoDePoOlId/dEvIcEsTaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolId/deviceStates/embeddedSIMDeviceStateId",
			Expected: &DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId{
				EmbeddedSIMActivationCodePoolId: "embeddedSIMActivationCodePoolId",
				EmbeddedSIMDeviceStateId:        "embeddedSIMDeviceStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolId/deviceStates/embeddedSIMDeviceStateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eMbEdDeDsImAcTiVaTiOnCoDePoOlS/eMbEdDeDsImAcTiVaTiOnCoDePoOlId/dEvIcEsTaTeS/eMbEdDeDsImDeViCeStAtEiD",
			Expected: &DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId{
				EmbeddedSIMActivationCodePoolId: "eMbEdDeDsImAcTiVaTiOnCoDePoOlId",
				EmbeddedSIMDeviceStateId:        "eMbEdDeDsImDeViCeStAtEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eMbEdDeDsImAcTiVaTiOnCoDePoOlS/eMbEdDeDsImAcTiVaTiOnCoDePoOlId/dEvIcEsTaTeS/eMbEdDeDsImDeViCeStAtEiD/extra",
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
