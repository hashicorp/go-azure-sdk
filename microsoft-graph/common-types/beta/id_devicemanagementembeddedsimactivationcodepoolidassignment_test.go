package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId{}

func TestNewDeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentID("embeddedSIMActivationCodePoolId", "embeddedSIMActivationCodePoolAssignmentId")

	if id.EmbeddedSIMActivationCodePoolId != "embeddedSIMActivationCodePoolId" {
		t.Fatalf("Expected %q but got %q for Segment 'EmbeddedSIMActivationCodePoolId'", id.EmbeddedSIMActivationCodePoolId, "embeddedSIMActivationCodePoolId")
	}

	if id.EmbeddedSIMActivationCodePoolAssignmentId != "embeddedSIMActivationCodePoolAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'EmbeddedSIMActivationCodePoolAssignmentId'", id.EmbeddedSIMActivationCodePoolAssignmentId, "embeddedSIMActivationCodePoolAssignmentId")
	}
}

func TestFormatDeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentID("embeddedSIMActivationCodePoolId", "embeddedSIMActivationCodePoolAssignmentId").ID()
	expected := "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolId/assignments/embeddedSIMActivationCodePoolAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId
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
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolId/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolId/assignments/embeddedSIMActivationCodePoolAssignmentId",
			Expected: &DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId{
				EmbeddedSIMActivationCodePoolId:           "embeddedSIMActivationCodePoolId",
				EmbeddedSIMActivationCodePoolAssignmentId: "embeddedSIMActivationCodePoolAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolId/assignments/embeddedSIMActivationCodePoolAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentID(v.Input)
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

		if actual.EmbeddedSIMActivationCodePoolAssignmentId != v.Expected.EmbeddedSIMActivationCodePoolAssignmentId {
			t.Fatalf("Expected %q but got %q for EmbeddedSIMActivationCodePoolAssignmentId", v.Expected.EmbeddedSIMActivationCodePoolAssignmentId, actual.EmbeddedSIMActivationCodePoolAssignmentId)
		}

	}
}

func TestParseDeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId
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
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolId/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eMbEdDeDsImAcTiVaTiOnCoDePoOlS/eMbEdDeDsImAcTiVaTiOnCoDePoOlId/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolId/assignments/embeddedSIMActivationCodePoolAssignmentId",
			Expected: &DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId{
				EmbeddedSIMActivationCodePoolId:           "embeddedSIMActivationCodePoolId",
				EmbeddedSIMActivationCodePoolAssignmentId: "embeddedSIMActivationCodePoolAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/embeddedSIMActivationCodePools/embeddedSIMActivationCodePoolId/assignments/embeddedSIMActivationCodePoolAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eMbEdDeDsImAcTiVaTiOnCoDePoOlS/eMbEdDeDsImAcTiVaTiOnCoDePoOlId/aSsIgNmEnTs/eMbEdDeDsImAcTiVaTiOnCoDePoOlAsSiGnMeNtId",
			Expected: &DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId{
				EmbeddedSIMActivationCodePoolId:           "eMbEdDeDsImAcTiVaTiOnCoDePoOlId",
				EmbeddedSIMActivationCodePoolAssignmentId: "eMbEdDeDsImAcTiVaTiOnCoDePoOlAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eMbEdDeDsImAcTiVaTiOnCoDePoOlS/eMbEdDeDsImAcTiVaTiOnCoDePoOlId/aSsIgNmEnTs/eMbEdDeDsImAcTiVaTiOnCoDePoOlAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentIDInsensitively(v.Input)
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

		if actual.EmbeddedSIMActivationCodePoolAssignmentId != v.Expected.EmbeddedSIMActivationCodePoolAssignmentId {
			t.Fatalf("Expected %q but got %q for EmbeddedSIMActivationCodePoolAssignmentId", v.Expected.EmbeddedSIMActivationCodePoolAssignmentId, actual.EmbeddedSIMActivationCodePoolAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId(t *testing.T) {
	segments := DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
