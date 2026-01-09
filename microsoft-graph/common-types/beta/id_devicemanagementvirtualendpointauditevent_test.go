package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointAuditEventId{}

func TestNewDeviceManagementVirtualEndpointAuditEventID(t *testing.T) {
	id := NewDeviceManagementVirtualEndpointAuditEventID("cloudPCAuditEventId")

	if id.CloudPCAuditEventId != "cloudPCAuditEventId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCAuditEventId'", id.CloudPCAuditEventId, "cloudPCAuditEventId")
	}
}

func TestFormatDeviceManagementVirtualEndpointAuditEventID(t *testing.T) {
	actual := NewDeviceManagementVirtualEndpointAuditEventID("cloudPCAuditEventId").ID()
	expected := "/deviceManagement/virtualEndpoint/auditEvents/cloudPCAuditEventId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementVirtualEndpointAuditEventID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointAuditEventId
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
			Input: "/deviceManagement/virtualEndpoint/auditEvents",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/auditEvents/cloudPCAuditEventId",
			Expected: &DeviceManagementVirtualEndpointAuditEventId{
				CloudPCAuditEventId: "cloudPCAuditEventId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/auditEvents/cloudPCAuditEventId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointAuditEventID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCAuditEventId != v.Expected.CloudPCAuditEventId {
			t.Fatalf("Expected %q but got %q for CloudPCAuditEventId", v.Expected.CloudPCAuditEventId, actual.CloudPCAuditEventId)
		}

	}
}

func TestParseDeviceManagementVirtualEndpointAuditEventIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointAuditEventId
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
			Input: "/deviceManagement/virtualEndpoint/auditEvents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/aUdItEvEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/auditEvents/cloudPCAuditEventId",
			Expected: &DeviceManagementVirtualEndpointAuditEventId{
				CloudPCAuditEventId: "cloudPCAuditEventId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/auditEvents/cloudPCAuditEventId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/aUdItEvEnTs/cLoUdPcAuDiTeVeNtId",
			Expected: &DeviceManagementVirtualEndpointAuditEventId{
				CloudPCAuditEventId: "cLoUdPcAuDiTeVeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/aUdItEvEnTs/cLoUdPcAuDiTeVeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointAuditEventIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCAuditEventId != v.Expected.CloudPCAuditEventId {
			t.Fatalf("Expected %q but got %q for CloudPCAuditEventId", v.Expected.CloudPCAuditEventId, actual.CloudPCAuditEventId)
		}

	}
}

func TestSegmentsForDeviceManagementVirtualEndpointAuditEventId(t *testing.T) {
	segments := DeviceManagementVirtualEndpointAuditEventId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementVirtualEndpointAuditEventId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
