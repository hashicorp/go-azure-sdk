package zebrafotaartifact

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementZebraFotaArtifactId{}

func TestNewDeviceManagementZebraFotaArtifactID(t *testing.T) {
	id := NewDeviceManagementZebraFotaArtifactID("zebraFotaArtifactIdValue")

	if id.ZebraFotaArtifactId != "zebraFotaArtifactIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ZebraFotaArtifactId'", id.ZebraFotaArtifactId, "zebraFotaArtifactIdValue")
	}
}

func TestFormatDeviceManagementZebraFotaArtifactID(t *testing.T) {
	actual := NewDeviceManagementZebraFotaArtifactID("zebraFotaArtifactIdValue").ID()
	expected := "/deviceManagement/zebraFotaArtifacts/zebraFotaArtifactIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementZebraFotaArtifactID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementZebraFotaArtifactId
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
			Input: "/deviceManagement/zebraFotaArtifacts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/zebraFotaArtifacts/zebraFotaArtifactIdValue",
			Expected: &DeviceManagementZebraFotaArtifactId{
				ZebraFotaArtifactId: "zebraFotaArtifactIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/zebraFotaArtifacts/zebraFotaArtifactIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementZebraFotaArtifactID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ZebraFotaArtifactId != v.Expected.ZebraFotaArtifactId {
			t.Fatalf("Expected %q but got %q for ZebraFotaArtifactId", v.Expected.ZebraFotaArtifactId, actual.ZebraFotaArtifactId)
		}

	}
}

func TestParseDeviceManagementZebraFotaArtifactIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementZebraFotaArtifactId
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
			Input: "/deviceManagement/zebraFotaArtifacts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/zEbRaFoTaArTiFaCtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/zebraFotaArtifacts/zebraFotaArtifactIdValue",
			Expected: &DeviceManagementZebraFotaArtifactId{
				ZebraFotaArtifactId: "zebraFotaArtifactIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/zebraFotaArtifacts/zebraFotaArtifactIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/zEbRaFoTaArTiFaCtS/zEbRaFoTaArTiFaCtIdVaLuE",
			Expected: &DeviceManagementZebraFotaArtifactId{
				ZebraFotaArtifactId: "zEbRaFoTaArTiFaCtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/zEbRaFoTaArTiFaCtS/zEbRaFoTaArTiFaCtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementZebraFotaArtifactIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ZebraFotaArtifactId != v.Expected.ZebraFotaArtifactId {
			t.Fatalf("Expected %q but got %q for ZebraFotaArtifactId", v.Expected.ZebraFotaArtifactId, actual.ZebraFotaArtifactId)
		}

	}
}

func TestSegmentsForDeviceManagementZebraFotaArtifactId(t *testing.T) {
	segments := DeviceManagementZebraFotaArtifactId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementZebraFotaArtifactId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
