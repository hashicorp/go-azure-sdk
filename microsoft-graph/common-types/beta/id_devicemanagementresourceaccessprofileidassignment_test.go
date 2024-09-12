package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementResourceAccessProfileIdAssignmentId{}

func TestNewDeviceManagementResourceAccessProfileIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementResourceAccessProfileIdAssignmentID("deviceManagementResourceAccessProfileBaseIdValue", "deviceManagementResourceAccessProfileAssignmentIdValue")

	if id.DeviceManagementResourceAccessProfileBaseId != "deviceManagementResourceAccessProfileBaseIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementResourceAccessProfileBaseId'", id.DeviceManagementResourceAccessProfileBaseId, "deviceManagementResourceAccessProfileBaseIdValue")
	}

	if id.DeviceManagementResourceAccessProfileAssignmentId != "deviceManagementResourceAccessProfileAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementResourceAccessProfileAssignmentId'", id.DeviceManagementResourceAccessProfileAssignmentId, "deviceManagementResourceAccessProfileAssignmentIdValue")
	}
}

func TestFormatDeviceManagementResourceAccessProfileIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementResourceAccessProfileIdAssignmentID("deviceManagementResourceAccessProfileBaseIdValue", "deviceManagementResourceAccessProfileAssignmentIdValue").ID()
	expected := "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseIdValue/assignments/deviceManagementResourceAccessProfileAssignmentIdValue"
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
			Input: "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseIdValue/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseIdValue/assignments/deviceManagementResourceAccessProfileAssignmentIdValue",
			Expected: &DeviceManagementResourceAccessProfileIdAssignmentId{
				DeviceManagementResourceAccessProfileBaseId:       "deviceManagementResourceAccessProfileBaseIdValue",
				DeviceManagementResourceAccessProfileAssignmentId: "deviceManagementResourceAccessProfileAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseIdValue/assignments/deviceManagementResourceAccessProfileAssignmentIdValue/extra",
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
			Input: "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEsOuRcEaCcEsSpRoFiLeS/dEvIcEmAnAgEmEnTrEsOuRcEaCcEsSpRoFiLeBaSeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseIdValue/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEsOuRcEaCcEsSpRoFiLeS/dEvIcEmAnAgEmEnTrEsOuRcEaCcEsSpRoFiLeBaSeIdVaLuE/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseIdValue/assignments/deviceManagementResourceAccessProfileAssignmentIdValue",
			Expected: &DeviceManagementResourceAccessProfileIdAssignmentId{
				DeviceManagementResourceAccessProfileBaseId:       "deviceManagementResourceAccessProfileBaseIdValue",
				DeviceManagementResourceAccessProfileAssignmentId: "deviceManagementResourceAccessProfileAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseIdValue/assignments/deviceManagementResourceAccessProfileAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEsOuRcEaCcEsSpRoFiLeS/dEvIcEmAnAgEmEnTrEsOuRcEaCcEsSpRoFiLeBaSeIdVaLuE/aSsIgNmEnTs/dEvIcEmAnAgEmEnTrEsOuRcEaCcEsSpRoFiLeAsSiGnMeNtIdVaLuE",
			Expected: &DeviceManagementResourceAccessProfileIdAssignmentId{
				DeviceManagementResourceAccessProfileBaseId:       "dEvIcEmAnAgEmEnTrEsOuRcEaCcEsSpRoFiLeBaSeIdVaLuE",
				DeviceManagementResourceAccessProfileAssignmentId: "dEvIcEmAnAgEmEnTrEsOuRcEaCcEsSpRoFiLeAsSiGnMeNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEsOuRcEaCcEsSpRoFiLeS/dEvIcEmAnAgEmEnTrEsOuRcEaCcEsSpRoFiLeBaSeIdVaLuE/aSsIgNmEnTs/dEvIcEmAnAgEmEnTrEsOuRcEaCcEsSpRoFiLeAsSiGnMeNtIdVaLuE/extra",
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
