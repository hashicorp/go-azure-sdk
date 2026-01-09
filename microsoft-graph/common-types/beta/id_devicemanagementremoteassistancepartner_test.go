package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementRemoteAssistancePartnerId{}

func TestNewDeviceManagementRemoteAssistancePartnerID(t *testing.T) {
	id := NewDeviceManagementRemoteAssistancePartnerID("remoteAssistancePartnerId")

	if id.RemoteAssistancePartnerId != "remoteAssistancePartnerId" {
		t.Fatalf("Expected %q but got %q for Segment 'RemoteAssistancePartnerId'", id.RemoteAssistancePartnerId, "remoteAssistancePartnerId")
	}
}

func TestFormatDeviceManagementRemoteAssistancePartnerID(t *testing.T) {
	actual := NewDeviceManagementRemoteAssistancePartnerID("remoteAssistancePartnerId").ID()
	expected := "/deviceManagement/remoteAssistancePartners/remoteAssistancePartnerId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementRemoteAssistancePartnerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementRemoteAssistancePartnerId
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
			Input: "/deviceManagement/remoteAssistancePartners",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/remoteAssistancePartners/remoteAssistancePartnerId",
			Expected: &DeviceManagementRemoteAssistancePartnerId{
				RemoteAssistancePartnerId: "remoteAssistancePartnerId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/remoteAssistancePartners/remoteAssistancePartnerId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementRemoteAssistancePartnerID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.RemoteAssistancePartnerId != v.Expected.RemoteAssistancePartnerId {
			t.Fatalf("Expected %q but got %q for RemoteAssistancePartnerId", v.Expected.RemoteAssistancePartnerId, actual.RemoteAssistancePartnerId)
		}

	}
}

func TestParseDeviceManagementRemoteAssistancePartnerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementRemoteAssistancePartnerId
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
			Input: "/deviceManagement/remoteAssistancePartners",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEmOtEaSsIsTaNcEpArTnErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/remoteAssistancePartners/remoteAssistancePartnerId",
			Expected: &DeviceManagementRemoteAssistancePartnerId{
				RemoteAssistancePartnerId: "remoteAssistancePartnerId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/remoteAssistancePartners/remoteAssistancePartnerId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEmOtEaSsIsTaNcEpArTnErS/rEmOtEaSsIsTaNcEpArTnErId",
			Expected: &DeviceManagementRemoteAssistancePartnerId{
				RemoteAssistancePartnerId: "rEmOtEaSsIsTaNcEpArTnErId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEmOtEaSsIsTaNcEpArTnErS/rEmOtEaSsIsTaNcEpArTnErId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementRemoteAssistancePartnerIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.RemoteAssistancePartnerId != v.Expected.RemoteAssistancePartnerId {
			t.Fatalf("Expected %q but got %q for RemoteAssistancePartnerId", v.Expected.RemoteAssistancePartnerId, actual.RemoteAssistancePartnerId)
		}

	}
}

func TestSegmentsForDeviceManagementRemoteAssistancePartnerId(t *testing.T) {
	segments := DeviceManagementRemoteAssistancePartnerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementRemoteAssistancePartnerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
