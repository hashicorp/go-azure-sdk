package zebrafotadeployment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementZebraFotaDeploymentId{}

func TestNewDeviceManagementZebraFotaDeploymentID(t *testing.T) {
	id := NewDeviceManagementZebraFotaDeploymentID("zebraFotaDeploymentIdValue")

	if id.ZebraFotaDeploymentId != "zebraFotaDeploymentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ZebraFotaDeploymentId'", id.ZebraFotaDeploymentId, "zebraFotaDeploymentIdValue")
	}
}

func TestFormatDeviceManagementZebraFotaDeploymentID(t *testing.T) {
	actual := NewDeviceManagementZebraFotaDeploymentID("zebraFotaDeploymentIdValue").ID()
	expected := "/deviceManagement/zebraFotaDeployments/zebraFotaDeploymentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementZebraFotaDeploymentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementZebraFotaDeploymentId
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
			Input: "/deviceManagement/zebraFotaDeployments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/zebraFotaDeployments/zebraFotaDeploymentIdValue",
			Expected: &DeviceManagementZebraFotaDeploymentId{
				ZebraFotaDeploymentId: "zebraFotaDeploymentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/zebraFotaDeployments/zebraFotaDeploymentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementZebraFotaDeploymentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ZebraFotaDeploymentId != v.Expected.ZebraFotaDeploymentId {
			t.Fatalf("Expected %q but got %q for ZebraFotaDeploymentId", v.Expected.ZebraFotaDeploymentId, actual.ZebraFotaDeploymentId)
		}

	}
}

func TestParseDeviceManagementZebraFotaDeploymentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementZebraFotaDeploymentId
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
			Input: "/deviceManagement/zebraFotaDeployments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/zEbRaFoTaDePlOyMeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/zebraFotaDeployments/zebraFotaDeploymentIdValue",
			Expected: &DeviceManagementZebraFotaDeploymentId{
				ZebraFotaDeploymentId: "zebraFotaDeploymentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/zebraFotaDeployments/zebraFotaDeploymentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/zEbRaFoTaDePlOyMeNtS/zEbRaFoTaDePlOyMeNtIdVaLuE",
			Expected: &DeviceManagementZebraFotaDeploymentId{
				ZebraFotaDeploymentId: "zEbRaFoTaDePlOyMeNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/zEbRaFoTaDePlOyMeNtS/zEbRaFoTaDePlOyMeNtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementZebraFotaDeploymentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ZebraFotaDeploymentId != v.Expected.ZebraFotaDeploymentId {
			t.Fatalf("Expected %q but got %q for ZebraFotaDeploymentId", v.Expected.ZebraFotaDeploymentId, actual.ZebraFotaDeploymentId)
		}

	}
}

func TestSegmentsForDeviceManagementZebraFotaDeploymentId(t *testing.T) {
	segments := DeviceManagementZebraFotaDeploymentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementZebraFotaDeploymentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
