package virtualendpointexternalpartnersetting

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointExternalPartnerSettingId{}

func TestNewDeviceManagementVirtualEndpointExternalPartnerSettingID(t *testing.T) {
	id := NewDeviceManagementVirtualEndpointExternalPartnerSettingID("cloudPCExternalPartnerSettingIdValue")

	if id.CloudPCExternalPartnerSettingId != "cloudPCExternalPartnerSettingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCExternalPartnerSettingId'", id.CloudPCExternalPartnerSettingId, "cloudPCExternalPartnerSettingIdValue")
	}
}

func TestFormatDeviceManagementVirtualEndpointExternalPartnerSettingID(t *testing.T) {
	actual := NewDeviceManagementVirtualEndpointExternalPartnerSettingID("cloudPCExternalPartnerSettingIdValue").ID()
	expected := "/deviceManagement/virtualEndpoint/externalPartnerSettings/cloudPCExternalPartnerSettingIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementVirtualEndpointExternalPartnerSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointExternalPartnerSettingId
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
			Input: "/deviceManagement/virtualEndpoint/externalPartnerSettings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/externalPartnerSettings/cloudPCExternalPartnerSettingIdValue",
			Expected: &DeviceManagementVirtualEndpointExternalPartnerSettingId{
				CloudPCExternalPartnerSettingId: "cloudPCExternalPartnerSettingIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/externalPartnerSettings/cloudPCExternalPartnerSettingIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointExternalPartnerSettingID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCExternalPartnerSettingId != v.Expected.CloudPCExternalPartnerSettingId {
			t.Fatalf("Expected %q but got %q for CloudPCExternalPartnerSettingId", v.Expected.CloudPCExternalPartnerSettingId, actual.CloudPCExternalPartnerSettingId)
		}

	}
}

func TestParseDeviceManagementVirtualEndpointExternalPartnerSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointExternalPartnerSettingId
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
			Input: "/deviceManagement/virtualEndpoint/externalPartnerSettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/eXtErNaLpArTnErSeTtInGs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/externalPartnerSettings/cloudPCExternalPartnerSettingIdValue",
			Expected: &DeviceManagementVirtualEndpointExternalPartnerSettingId{
				CloudPCExternalPartnerSettingId: "cloudPCExternalPartnerSettingIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/externalPartnerSettings/cloudPCExternalPartnerSettingIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/eXtErNaLpArTnErSeTtInGs/cLoUdPcExTeRnAlPaRtNeRsEtTiNgIdVaLuE",
			Expected: &DeviceManagementVirtualEndpointExternalPartnerSettingId{
				CloudPCExternalPartnerSettingId: "cLoUdPcExTeRnAlPaRtNeRsEtTiNgIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/eXtErNaLpArTnErSeTtInGs/cLoUdPcExTeRnAlPaRtNeRsEtTiNgIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointExternalPartnerSettingIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCExternalPartnerSettingId != v.Expected.CloudPCExternalPartnerSettingId {
			t.Fatalf("Expected %q but got %q for CloudPCExternalPartnerSettingId", v.Expected.CloudPCExternalPartnerSettingId, actual.CloudPCExternalPartnerSettingId)
		}

	}
}

func TestSegmentsForDeviceManagementVirtualEndpointExternalPartnerSettingId(t *testing.T) {
	segments := DeviceManagementVirtualEndpointExternalPartnerSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementVirtualEndpointExternalPartnerSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
