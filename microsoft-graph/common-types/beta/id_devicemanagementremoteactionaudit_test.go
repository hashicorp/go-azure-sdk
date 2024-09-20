package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementRemoteActionAuditId{}

func TestNewDeviceManagementRemoteActionAuditID(t *testing.T) {
	id := NewDeviceManagementRemoteActionAuditID("remoteActionAuditId")

	if id.RemoteActionAuditId != "remoteActionAuditId" {
		t.Fatalf("Expected %q but got %q for Segment 'RemoteActionAuditId'", id.RemoteActionAuditId, "remoteActionAuditId")
	}
}

func TestFormatDeviceManagementRemoteActionAuditID(t *testing.T) {
	actual := NewDeviceManagementRemoteActionAuditID("remoteActionAuditId").ID()
	expected := "/deviceManagement/remoteActionAudits/remoteActionAuditId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementRemoteActionAuditID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementRemoteActionAuditId
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
			Input: "/deviceManagement/remoteActionAudits",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/remoteActionAudits/remoteActionAuditId",
			Expected: &DeviceManagementRemoteActionAuditId{
				RemoteActionAuditId: "remoteActionAuditId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/remoteActionAudits/remoteActionAuditId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementRemoteActionAuditID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.RemoteActionAuditId != v.Expected.RemoteActionAuditId {
			t.Fatalf("Expected %q but got %q for RemoteActionAuditId", v.Expected.RemoteActionAuditId, actual.RemoteActionAuditId)
		}

	}
}

func TestParseDeviceManagementRemoteActionAuditIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementRemoteActionAuditId
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
			Input: "/deviceManagement/remoteActionAudits",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEmOtEaCtIoNaUdItS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/remoteActionAudits/remoteActionAuditId",
			Expected: &DeviceManagementRemoteActionAuditId{
				RemoteActionAuditId: "remoteActionAuditId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/remoteActionAudits/remoteActionAuditId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEmOtEaCtIoNaUdItS/rEmOtEaCtIoNaUdItId",
			Expected: &DeviceManagementRemoteActionAuditId{
				RemoteActionAuditId: "rEmOtEaCtIoNaUdItId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEmOtEaCtIoNaUdItS/rEmOtEaCtIoNaUdItId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementRemoteActionAuditIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.RemoteActionAuditId != v.Expected.RemoteActionAuditId {
			t.Fatalf("Expected %q but got %q for RemoteActionAuditId", v.Expected.RemoteActionAuditId, actual.RemoteActionAuditId)
		}

	}
}

func TestSegmentsForDeviceManagementRemoteActionAuditId(t *testing.T) {
	segments := DeviceManagementRemoteActionAuditId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementRemoteActionAuditId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
