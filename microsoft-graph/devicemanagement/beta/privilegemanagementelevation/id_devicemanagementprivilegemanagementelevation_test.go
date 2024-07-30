package privilegemanagementelevation

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementPrivilegeManagementElevationId{}

func TestNewDeviceManagementPrivilegeManagementElevationID(t *testing.T) {
	id := NewDeviceManagementPrivilegeManagementElevationID("privilegeManagementElevationIdValue")

	if id.PrivilegeManagementElevationId != "privilegeManagementElevationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PrivilegeManagementElevationId'", id.PrivilegeManagementElevationId, "privilegeManagementElevationIdValue")
	}
}

func TestFormatDeviceManagementPrivilegeManagementElevationID(t *testing.T) {
	actual := NewDeviceManagementPrivilegeManagementElevationID("privilegeManagementElevationIdValue").ID()
	expected := "/deviceManagement/privilegeManagementElevations/privilegeManagementElevationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementPrivilegeManagementElevationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementPrivilegeManagementElevationId
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
			Input: "/deviceManagement/privilegeManagementElevations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/privilegeManagementElevations/privilegeManagementElevationIdValue",
			Expected: &DeviceManagementPrivilegeManagementElevationId{
				PrivilegeManagementElevationId: "privilegeManagementElevationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/privilegeManagementElevations/privilegeManagementElevationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementPrivilegeManagementElevationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrivilegeManagementElevationId != v.Expected.PrivilegeManagementElevationId {
			t.Fatalf("Expected %q but got %q for PrivilegeManagementElevationId", v.Expected.PrivilegeManagementElevationId, actual.PrivilegeManagementElevationId)
		}

	}
}

func TestParseDeviceManagementPrivilegeManagementElevationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementPrivilegeManagementElevationId
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
			Input: "/deviceManagement/privilegeManagementElevations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/pRiViLeGeMaNaGeMeNtElEvAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/privilegeManagementElevations/privilegeManagementElevationIdValue",
			Expected: &DeviceManagementPrivilegeManagementElevationId{
				PrivilegeManagementElevationId: "privilegeManagementElevationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/privilegeManagementElevations/privilegeManagementElevationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/pRiViLeGeMaNaGeMeNtElEvAtIoNs/pRiViLeGeMaNaGeMeNtElEvAtIoNiDvAlUe",
			Expected: &DeviceManagementPrivilegeManagementElevationId{
				PrivilegeManagementElevationId: "pRiViLeGeMaNaGeMeNtElEvAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/pRiViLeGeMaNaGeMeNtElEvAtIoNs/pRiViLeGeMaNaGeMeNtElEvAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementPrivilegeManagementElevationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrivilegeManagementElevationId != v.Expected.PrivilegeManagementElevationId {
			t.Fatalf("Expected %q but got %q for PrivilegeManagementElevationId", v.Expected.PrivilegeManagementElevationId, actual.PrivilegeManagementElevationId)
		}

	}
}

func TestSegmentsForDeviceManagementPrivilegeManagementElevationId(t *testing.T) {
	segments := DeviceManagementPrivilegeManagementElevationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementPrivilegeManagementElevationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
