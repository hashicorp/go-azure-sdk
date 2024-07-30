package embeddedsimactivationcodepoolassignment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementEmbeddedSIMActivationCodePoolId{}

func TestNewDeviceManagementEmbeddedSIMActivationCodePoolID(t *testing.T) {
	id := NewDeviceManagementEmbeddedSIMActivationCodePoolID("embeddedSIMActivationCodePoolIdValue")

	if id.EmbeddedSIMActivationCodePoolId != "embeddedSIMActivationCodePoolIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EmbeddedSIMActivationCodePoolId'", id.EmbeddedSIMActivationCodePoolId, "embeddedSIMActivationCodePoolIdValue")
	}
}

func TestFormatDeviceManagementEmbeddedSIMActivationCodePoolID(t *testing.T) {
	actual := NewDeviceManagementEmbeddedSIMActivationCodePoolID("embeddedSIMActivationCodePoolIdValue").ID()
	expected := "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementEmbeddedSIMActivationCodePoolID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementEmbeddedSIMActivationCodePoolId
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
			// Valid URI
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolIdValue",
			Expected: &DeviceManagementEmbeddedSIMActivationCodePoolId{
				EmbeddedSIMActivationCodePoolId: "embeddedSIMActivationCodePoolIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementEmbeddedSIMActivationCodePoolID(v.Input)
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

	}
}

func TestParseDeviceManagementEmbeddedSIMActivationCodePoolIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementEmbeddedSIMActivationCodePoolId
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
			// Valid URI
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolIdValue",
			Expected: &DeviceManagementEmbeddedSIMActivationCodePoolId{
				EmbeddedSIMActivationCodePoolId: "embeddedSIMActivationCodePoolIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eMbEdDeDsImAcTiVaTiOnCoDePoOlS/eMbEdDeDsImAcTiVaTiOnCoDePoOlIdVaLuE",
			Expected: &DeviceManagementEmbeddedSIMActivationCodePoolId{
				EmbeddedSIMActivationCodePoolId: "eMbEdDeDsImAcTiVaTiOnCoDePoOlIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eMbEdDeDsImAcTiVaTiOnCoDePoOlS/eMbEdDeDsImAcTiVaTiOnCoDePoOlIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementEmbeddedSIMActivationCodePoolIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementEmbeddedSIMActivationCodePoolId(t *testing.T) {
	segments := DeviceManagementEmbeddedSIMActivationCodePoolId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementEmbeddedSIMActivationCodePoolId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
