package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAuditEventId{}

func TestNewDeviceManagementAuditEventID(t *testing.T) {
	id := NewDeviceManagementAuditEventID("auditEventId")

	if id.AuditEventId != "auditEventId" {
		t.Fatalf("Expected %q but got %q for Segment 'AuditEventId'", id.AuditEventId, "auditEventId")
	}
}

func TestFormatDeviceManagementAuditEventID(t *testing.T) {
	actual := NewDeviceManagementAuditEventID("auditEventId").ID()
	expected := "/deviceManagement/auditEvents/auditEventId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementAuditEventID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementAuditEventId
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
			Input: "/deviceManagement/auditEvents",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/auditEvents/auditEventId",
			Expected: &DeviceManagementAuditEventId{
				AuditEventId: "auditEventId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/auditEvents/auditEventId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementAuditEventID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuditEventId != v.Expected.AuditEventId {
			t.Fatalf("Expected %q but got %q for AuditEventId", v.Expected.AuditEventId, actual.AuditEventId)
		}

	}
}

func TestParseDeviceManagementAuditEventIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementAuditEventId
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
			Input: "/deviceManagement/auditEvents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aUdItEvEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/auditEvents/auditEventId",
			Expected: &DeviceManagementAuditEventId{
				AuditEventId: "auditEventId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/auditEvents/auditEventId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aUdItEvEnTs/aUdItEvEnTiD",
			Expected: &DeviceManagementAuditEventId{
				AuditEventId: "aUdItEvEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aUdItEvEnTs/aUdItEvEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementAuditEventIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuditEventId != v.Expected.AuditEventId {
			t.Fatalf("Expected %q but got %q for AuditEventId", v.Expected.AuditEventId, actual.AuditEventId)
		}

	}
}

func TestSegmentsForDeviceManagementAuditEventId(t *testing.T) {
	segments := DeviceManagementAuditEventId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementAuditEventId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
