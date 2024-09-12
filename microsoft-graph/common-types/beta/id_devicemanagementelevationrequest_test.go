package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementElevationRequestId{}

func TestNewDeviceManagementElevationRequestID(t *testing.T) {
	id := NewDeviceManagementElevationRequestID("privilegeManagementElevationRequestIdValue")

	if id.PrivilegeManagementElevationRequestId != "privilegeManagementElevationRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PrivilegeManagementElevationRequestId'", id.PrivilegeManagementElevationRequestId, "privilegeManagementElevationRequestIdValue")
	}
}

func TestFormatDeviceManagementElevationRequestID(t *testing.T) {
	actual := NewDeviceManagementElevationRequestID("privilegeManagementElevationRequestIdValue").ID()
	expected := "/deviceManagement/elevationRequests/privilegeManagementElevationRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementElevationRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementElevationRequestId
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
			Input: "/deviceManagement/elevationRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/elevationRequests/privilegeManagementElevationRequestIdValue",
			Expected: &DeviceManagementElevationRequestId{
				PrivilegeManagementElevationRequestId: "privilegeManagementElevationRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/elevationRequests/privilegeManagementElevationRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementElevationRequestID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrivilegeManagementElevationRequestId != v.Expected.PrivilegeManagementElevationRequestId {
			t.Fatalf("Expected %q but got %q for PrivilegeManagementElevationRequestId", v.Expected.PrivilegeManagementElevationRequestId, actual.PrivilegeManagementElevationRequestId)
		}

	}
}

func TestParseDeviceManagementElevationRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementElevationRequestId
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
			Input: "/deviceManagement/elevationRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eLeVaTiOnReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/elevationRequests/privilegeManagementElevationRequestIdValue",
			Expected: &DeviceManagementElevationRequestId{
				PrivilegeManagementElevationRequestId: "privilegeManagementElevationRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/elevationRequests/privilegeManagementElevationRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eLeVaTiOnReQuEsTs/pRiViLeGeMaNaGeMeNtElEvAtIoNrEqUeStIdVaLuE",
			Expected: &DeviceManagementElevationRequestId{
				PrivilegeManagementElevationRequestId: "pRiViLeGeMaNaGeMeNtElEvAtIoNrEqUeStIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eLeVaTiOnReQuEsTs/pRiViLeGeMaNaGeMeNtElEvAtIoNrEqUeStIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementElevationRequestIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrivilegeManagementElevationRequestId != v.Expected.PrivilegeManagementElevationRequestId {
			t.Fatalf("Expected %q but got %q for PrivilegeManagementElevationRequestId", v.Expected.PrivilegeManagementElevationRequestId, actual.PrivilegeManagementElevationRequestId)
		}

	}
}

func TestSegmentsForDeviceManagementElevationRequestId(t *testing.T) {
	segments := DeviceManagementElevationRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementElevationRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
