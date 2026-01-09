package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointUserSettingId{}

func TestNewDeviceManagementVirtualEndpointUserSettingID(t *testing.T) {
	id := NewDeviceManagementVirtualEndpointUserSettingID("cloudPCUserSettingId")

	if id.CloudPCUserSettingId != "cloudPCUserSettingId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCUserSettingId'", id.CloudPCUserSettingId, "cloudPCUserSettingId")
	}
}

func TestFormatDeviceManagementVirtualEndpointUserSettingID(t *testing.T) {
	actual := NewDeviceManagementVirtualEndpointUserSettingID("cloudPCUserSettingId").ID()
	expected := "/deviceManagement/virtualEndpoint/userSettings/cloudPCUserSettingId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementVirtualEndpointUserSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointUserSettingId
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
			Input: "/deviceManagement/virtualEndpoint",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/virtualEndpoint/userSettings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/userSettings/cloudPCUserSettingId",
			Expected: &DeviceManagementVirtualEndpointUserSettingId{
				CloudPCUserSettingId: "cloudPCUserSettingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/userSettings/cloudPCUserSettingId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointUserSettingID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCUserSettingId != v.Expected.CloudPCUserSettingId {
			t.Fatalf("Expected %q but got %q for CloudPCUserSettingId", v.Expected.CloudPCUserSettingId, actual.CloudPCUserSettingId)
		}

	}
}

func TestParseDeviceManagementVirtualEndpointUserSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointUserSettingId
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
			Input: "/deviceManagement/virtualEndpoint",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/virtualEndpoint/userSettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/uSeRsEtTiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/userSettings/cloudPCUserSettingId",
			Expected: &DeviceManagementVirtualEndpointUserSettingId{
				CloudPCUserSettingId: "cloudPCUserSettingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/userSettings/cloudPCUserSettingId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/uSeRsEtTiNgS/cLoUdPcUsErSeTtInGiD",
			Expected: &DeviceManagementVirtualEndpointUserSettingId{
				CloudPCUserSettingId: "cLoUdPcUsErSeTtInGiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/uSeRsEtTiNgS/cLoUdPcUsErSeTtInGiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointUserSettingIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCUserSettingId != v.Expected.CloudPCUserSettingId {
			t.Fatalf("Expected %q but got %q for CloudPCUserSettingId", v.Expected.CloudPCUserSettingId, actual.CloudPCUserSettingId)
		}

	}
}

func TestSegmentsForDeviceManagementVirtualEndpointUserSettingId(t *testing.T) {
	segments := DeviceManagementVirtualEndpointUserSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementVirtualEndpointUserSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
