package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateIdMigratableToId{}

func TestNewDeviceManagementTemplateIdMigratableToID(t *testing.T) {
	id := NewDeviceManagementTemplateIdMigratableToID("deviceManagementTemplateId", "deviceManagementTemplateId1")

	if id.DeviceManagementTemplateId != "deviceManagementTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateId'", id.DeviceManagementTemplateId, "deviceManagementTemplateId")
	}

	if id.DeviceManagementTemplateId1 != "deviceManagementTemplateId1" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateId1'", id.DeviceManagementTemplateId1, "deviceManagementTemplateId1")
	}
}

func TestFormatDeviceManagementTemplateIdMigratableToID(t *testing.T) {
	actual := NewDeviceManagementTemplateIdMigratableToID("deviceManagementTemplateId", "deviceManagementTemplateId1").ID()
	expected := "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementTemplateIdMigratableToID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTemplateIdMigratableToId
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
			Input: "/deviceManagement/templates",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1",
			Expected: &DeviceManagementTemplateIdMigratableToId{
				DeviceManagementTemplateId:  "deviceManagementTemplateId",
				DeviceManagementTemplateId1: "deviceManagementTemplateId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTemplateIdMigratableToID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementTemplateId != v.Expected.DeviceManagementTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementTemplateId", v.Expected.DeviceManagementTemplateId, actual.DeviceManagementTemplateId)
		}

		if actual.DeviceManagementTemplateId1 != v.Expected.DeviceManagementTemplateId1 {
			t.Fatalf("Expected %q but got %q for DeviceManagementTemplateId1", v.Expected.DeviceManagementTemplateId1, actual.DeviceManagementTemplateId1)
		}

	}
}

func TestParseDeviceManagementTemplateIdMigratableToIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTemplateIdMigratableToId
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
			Input: "/deviceManagement/templates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiD/mIgRaTaBlEtO",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1",
			Expected: &DeviceManagementTemplateIdMigratableToId{
				DeviceManagementTemplateId:  "deviceManagementTemplateId",
				DeviceManagementTemplateId1: "deviceManagementTemplateId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiD/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1",
			Expected: &DeviceManagementTemplateIdMigratableToId{
				DeviceManagementTemplateId:  "dEvIcEmAnAgEmEnTtEmPlAtEiD",
				DeviceManagementTemplateId1: "dEvIcEmAnAgEmEnTtEmPlAtEiD1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiD/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTemplateIdMigratableToIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementTemplateId != v.Expected.DeviceManagementTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementTemplateId", v.Expected.DeviceManagementTemplateId, actual.DeviceManagementTemplateId)
		}

		if actual.DeviceManagementTemplateId1 != v.Expected.DeviceManagementTemplateId1 {
			t.Fatalf("Expected %q but got %q for DeviceManagementTemplateId1", v.Expected.DeviceManagementTemplateId1, actual.DeviceManagementTemplateId1)
		}

	}
}

func TestSegmentsForDeviceManagementTemplateIdMigratableToId(t *testing.T) {
	segments := DeviceManagementTemplateIdMigratableToId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementTemplateIdMigratableToId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
