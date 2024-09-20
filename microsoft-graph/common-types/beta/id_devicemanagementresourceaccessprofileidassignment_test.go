package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementResourceAccessProfileIdAssignmentId{}

func TestNewDeviceManagementResourceAccessProfileIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementResourceAccessProfileIdAssignmentID("deviceManagementResourceAccessProfileBaseId", "deviceManagementResourceAccessProfileAssignmentId")

	if id.DeviceManagementResourceAccessProfileBaseId != "deviceManagementResourceAccessProfileBaseId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementResourceAccessProfileBaseId'", id.DeviceManagementResourceAccessProfileBaseId, "deviceManagementResourceAccessProfileBaseId")
	}

	if id.DeviceManagementResourceAccessProfileAssignmentId != "deviceManagementResourceAccessProfileAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementResourceAccessProfileAssignmentId'", id.DeviceManagementResourceAccessProfileAssignmentId, "deviceManagementResourceAccessProfileAssignmentId")
	}
}

func TestFormatDeviceManagementResourceAccessProfileIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementResourceAccessProfileIdAssignmentID("deviceManagementResourceAccessProfileBaseId", "deviceManagementResourceAccessProfileAssignmentId").ID()
	expected := "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseId/assignments/deviceManagementResourceAccessProfileAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementResourceAccessProfileIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementResourceAccessProfileIdAssignmentId
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
			Input: "/deviceManagement/resourceAccessProfiles",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseId/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseId/assignments/deviceManagementResourceAccessProfileAssignmentId",
			Expected: &DeviceManagementResourceAccessProfileIdAssignmentId{
				DeviceManagementResourceAccessProfileBaseId:       "deviceManagementResourceAccessProfileBaseId",
				DeviceManagementResourceAccessProfileAssignmentId: "deviceManagementResourceAccessProfileAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseId/assignments/deviceManagementResourceAccessProfileAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementResourceAccessProfileIdAssignmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementResourceAccessProfileBaseId != v.Expected.DeviceManagementResourceAccessProfileBaseId {
			t.Fatalf("Expected %q but got %q for DeviceManagementResourceAccessProfileBaseId", v.Expected.DeviceManagementResourceAccessProfileBaseId, actual.DeviceManagementResourceAccessProfileBaseId)
		}

		if actual.DeviceManagementResourceAccessProfileAssignmentId != v.Expected.DeviceManagementResourceAccessProfileAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementResourceAccessProfileAssignmentId", v.Expected.DeviceManagementResourceAccessProfileAssignmentId, actual.DeviceManagementResourceAccessProfileAssignmentId)
		}

	}
}

func TestParseDeviceManagementResourceAccessProfileIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementResourceAccessProfileIdAssignmentId
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
			Input: "/deviceManagement/resourceAccessProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEsOuRcEaCcEsSpRoFiLeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEsOuRcEaCcEsSpRoFiLeS/dEvIcEmAnAgEmEnTrEsOuRcEaCcEsSpRoFiLeBaSeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseId/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEsOuRcEaCcEsSpRoFiLeS/dEvIcEmAnAgEmEnTrEsOuRcEaCcEsSpRoFiLeBaSeId/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseId/assignments/deviceManagementResourceAccessProfileAssignmentId",
			Expected: &DeviceManagementResourceAccessProfileIdAssignmentId{
				DeviceManagementResourceAccessProfileBaseId:       "deviceManagementResourceAccessProfileBaseId",
				DeviceManagementResourceAccessProfileAssignmentId: "deviceManagementResourceAccessProfileAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseId/assignments/deviceManagementResourceAccessProfileAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEsOuRcEaCcEsSpRoFiLeS/dEvIcEmAnAgEmEnTrEsOuRcEaCcEsSpRoFiLeBaSeId/aSsIgNmEnTs/dEvIcEmAnAgEmEnTrEsOuRcEaCcEsSpRoFiLeAsSiGnMeNtId",
			Expected: &DeviceManagementResourceAccessProfileIdAssignmentId{
				DeviceManagementResourceAccessProfileBaseId:       "dEvIcEmAnAgEmEnTrEsOuRcEaCcEsSpRoFiLeBaSeId",
				DeviceManagementResourceAccessProfileAssignmentId: "dEvIcEmAnAgEmEnTrEsOuRcEaCcEsSpRoFiLeAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEsOuRcEaCcEsSpRoFiLeS/dEvIcEmAnAgEmEnTrEsOuRcEaCcEsSpRoFiLeBaSeId/aSsIgNmEnTs/dEvIcEmAnAgEmEnTrEsOuRcEaCcEsSpRoFiLeAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementResourceAccessProfileIdAssignmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementResourceAccessProfileBaseId != v.Expected.DeviceManagementResourceAccessProfileBaseId {
			t.Fatalf("Expected %q but got %q for DeviceManagementResourceAccessProfileBaseId", v.Expected.DeviceManagementResourceAccessProfileBaseId, actual.DeviceManagementResourceAccessProfileBaseId)
		}

		if actual.DeviceManagementResourceAccessProfileAssignmentId != v.Expected.DeviceManagementResourceAccessProfileAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementResourceAccessProfileAssignmentId", v.Expected.DeviceManagementResourceAccessProfileAssignmentId, actual.DeviceManagementResourceAccessProfileAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementResourceAccessProfileIdAssignmentId(t *testing.T) {
	segments := DeviceManagementResourceAccessProfileIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementResourceAccessProfileIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
