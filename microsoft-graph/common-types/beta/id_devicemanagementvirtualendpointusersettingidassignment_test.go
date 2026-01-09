package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointUserSettingIdAssignmentId{}

func TestNewDeviceManagementVirtualEndpointUserSettingIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementVirtualEndpointUserSettingIdAssignmentID("cloudPCUserSettingId", "cloudPCUserSettingAssignmentId")

	if id.CloudPCUserSettingId != "cloudPCUserSettingId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCUserSettingId'", id.CloudPCUserSettingId, "cloudPCUserSettingId")
	}

	if id.CloudPCUserSettingAssignmentId != "cloudPCUserSettingAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCUserSettingAssignmentId'", id.CloudPCUserSettingAssignmentId, "cloudPCUserSettingAssignmentId")
	}
}

func TestFormatDeviceManagementVirtualEndpointUserSettingIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementVirtualEndpointUserSettingIdAssignmentID("cloudPCUserSettingId", "cloudPCUserSettingAssignmentId").ID()
	expected := "/deviceManagement/virtualEndpoint/userSettings/cloudPCUserSettingId/assignments/cloudPCUserSettingAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementVirtualEndpointUserSettingIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointUserSettingIdAssignmentId
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
			// Incomplete URI
			Input: "/deviceManagement/virtualEndpoint/userSettings/cloudPCUserSettingId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/virtualEndpoint/userSettings/cloudPCUserSettingId/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/userSettings/cloudPCUserSettingId/assignments/cloudPCUserSettingAssignmentId",
			Expected: &DeviceManagementVirtualEndpointUserSettingIdAssignmentId{
				CloudPCUserSettingId:           "cloudPCUserSettingId",
				CloudPCUserSettingAssignmentId: "cloudPCUserSettingAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/userSettings/cloudPCUserSettingId/assignments/cloudPCUserSettingAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointUserSettingIdAssignmentID(v.Input)
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

		if actual.CloudPCUserSettingAssignmentId != v.Expected.CloudPCUserSettingAssignmentId {
			t.Fatalf("Expected %q but got %q for CloudPCUserSettingAssignmentId", v.Expected.CloudPCUserSettingAssignmentId, actual.CloudPCUserSettingAssignmentId)
		}

	}
}

func TestParseDeviceManagementVirtualEndpointUserSettingIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointUserSettingIdAssignmentId
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
			// Incomplete URI
			Input: "/deviceManagement/virtualEndpoint/userSettings/cloudPCUserSettingId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/uSeRsEtTiNgS/cLoUdPcUsErSeTtInGiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/virtualEndpoint/userSettings/cloudPCUserSettingId/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/uSeRsEtTiNgS/cLoUdPcUsErSeTtInGiD/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/userSettings/cloudPCUserSettingId/assignments/cloudPCUserSettingAssignmentId",
			Expected: &DeviceManagementVirtualEndpointUserSettingIdAssignmentId{
				CloudPCUserSettingId:           "cloudPCUserSettingId",
				CloudPCUserSettingAssignmentId: "cloudPCUserSettingAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/userSettings/cloudPCUserSettingId/assignments/cloudPCUserSettingAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/uSeRsEtTiNgS/cLoUdPcUsErSeTtInGiD/aSsIgNmEnTs/cLoUdPcUsErSeTtInGaSsIgNmEnTiD",
			Expected: &DeviceManagementVirtualEndpointUserSettingIdAssignmentId{
				CloudPCUserSettingId:           "cLoUdPcUsErSeTtInGiD",
				CloudPCUserSettingAssignmentId: "cLoUdPcUsErSeTtInGaSsIgNmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/uSeRsEtTiNgS/cLoUdPcUsErSeTtInGiD/aSsIgNmEnTs/cLoUdPcUsErSeTtInGaSsIgNmEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointUserSettingIdAssignmentIDInsensitively(v.Input)
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

		if actual.CloudPCUserSettingAssignmentId != v.Expected.CloudPCUserSettingAssignmentId {
			t.Fatalf("Expected %q but got %q for CloudPCUserSettingAssignmentId", v.Expected.CloudPCUserSettingAssignmentId, actual.CloudPCUserSettingAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementVirtualEndpointUserSettingIdAssignmentId(t *testing.T) {
	segments := DeviceManagementVirtualEndpointUserSettingIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementVirtualEndpointUserSettingIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
