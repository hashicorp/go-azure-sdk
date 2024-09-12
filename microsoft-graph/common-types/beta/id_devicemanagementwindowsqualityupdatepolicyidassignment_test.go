package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId{}

func TestNewDeviceManagementWindowsQualityUpdatePolicyIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementWindowsQualityUpdatePolicyIdAssignmentID("windowsQualityUpdatePolicyIdValue", "windowsQualityUpdatePolicyAssignmentIdValue")

	if id.WindowsQualityUpdatePolicyId != "windowsQualityUpdatePolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsQualityUpdatePolicyId'", id.WindowsQualityUpdatePolicyId, "windowsQualityUpdatePolicyIdValue")
	}

	if id.WindowsQualityUpdatePolicyAssignmentId != "windowsQualityUpdatePolicyAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsQualityUpdatePolicyAssignmentId'", id.WindowsQualityUpdatePolicyAssignmentId, "windowsQualityUpdatePolicyAssignmentIdValue")
	}
}

func TestFormatDeviceManagementWindowsQualityUpdatePolicyIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementWindowsQualityUpdatePolicyIdAssignmentID("windowsQualityUpdatePolicyIdValue", "windowsQualityUpdatePolicyAssignmentIdValue").ID()
	expected := "/deviceManagement/windowsQualityUpdatePolicies/windowsQualityUpdatePolicyIdValue/assignments/windowsQualityUpdatePolicyAssignmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementWindowsQualityUpdatePolicyIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId
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
			Input: "/deviceManagement/windowsQualityUpdatePolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/windowsQualityUpdatePolicies/windowsQualityUpdatePolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/windowsQualityUpdatePolicies/windowsQualityUpdatePolicyIdValue/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsQualityUpdatePolicies/windowsQualityUpdatePolicyIdValue/assignments/windowsQualityUpdatePolicyAssignmentIdValue",
			Expected: &DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId{
				WindowsQualityUpdatePolicyId:           "windowsQualityUpdatePolicyIdValue",
				WindowsQualityUpdatePolicyAssignmentId: "windowsQualityUpdatePolicyAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsQualityUpdatePolicies/windowsQualityUpdatePolicyIdValue/assignments/windowsQualityUpdatePolicyAssignmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsQualityUpdatePolicyIdAssignmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsQualityUpdatePolicyId != v.Expected.WindowsQualityUpdatePolicyId {
			t.Fatalf("Expected %q but got %q for WindowsQualityUpdatePolicyId", v.Expected.WindowsQualityUpdatePolicyId, actual.WindowsQualityUpdatePolicyId)
		}

		if actual.WindowsQualityUpdatePolicyAssignmentId != v.Expected.WindowsQualityUpdatePolicyAssignmentId {
			t.Fatalf("Expected %q but got %q for WindowsQualityUpdatePolicyAssignmentId", v.Expected.WindowsQualityUpdatePolicyAssignmentId, actual.WindowsQualityUpdatePolicyAssignmentId)
		}

	}
}

func TestParseDeviceManagementWindowsQualityUpdatePolicyIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId
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
			Input: "/deviceManagement/windowsQualityUpdatePolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsQuAlItYuPdAtEpOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/windowsQualityUpdatePolicies/windowsQualityUpdatePolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsQuAlItYuPdAtEpOlIcIeS/wInDoWsQuAlItYuPdAtEpOlIcYiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/windowsQualityUpdatePolicies/windowsQualityUpdatePolicyIdValue/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsQuAlItYuPdAtEpOlIcIeS/wInDoWsQuAlItYuPdAtEpOlIcYiDvAlUe/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsQualityUpdatePolicies/windowsQualityUpdatePolicyIdValue/assignments/windowsQualityUpdatePolicyAssignmentIdValue",
			Expected: &DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId{
				WindowsQualityUpdatePolicyId:           "windowsQualityUpdatePolicyIdValue",
				WindowsQualityUpdatePolicyAssignmentId: "windowsQualityUpdatePolicyAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsQualityUpdatePolicies/windowsQualityUpdatePolicyIdValue/assignments/windowsQualityUpdatePolicyAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsQuAlItYuPdAtEpOlIcIeS/wInDoWsQuAlItYuPdAtEpOlIcYiDvAlUe/aSsIgNmEnTs/wInDoWsQuAlItYuPdAtEpOlIcYaSsIgNmEnTiDvAlUe",
			Expected: &DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId{
				WindowsQualityUpdatePolicyId:           "wInDoWsQuAlItYuPdAtEpOlIcYiDvAlUe",
				WindowsQualityUpdatePolicyAssignmentId: "wInDoWsQuAlItYuPdAtEpOlIcYaSsIgNmEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsQuAlItYuPdAtEpOlIcIeS/wInDoWsQuAlItYuPdAtEpOlIcYiDvAlUe/aSsIgNmEnTs/wInDoWsQuAlItYuPdAtEpOlIcYaSsIgNmEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsQualityUpdatePolicyIdAssignmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsQualityUpdatePolicyId != v.Expected.WindowsQualityUpdatePolicyId {
			t.Fatalf("Expected %q but got %q for WindowsQualityUpdatePolicyId", v.Expected.WindowsQualityUpdatePolicyId, actual.WindowsQualityUpdatePolicyId)
		}

		if actual.WindowsQualityUpdatePolicyAssignmentId != v.Expected.WindowsQualityUpdatePolicyAssignmentId {
			t.Fatalf("Expected %q but got %q for WindowsQualityUpdatePolicyAssignmentId", v.Expected.WindowsQualityUpdatePolicyAssignmentId, actual.WindowsQualityUpdatePolicyAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementWindowsQualityUpdatePolicyIdAssignmentId(t *testing.T) {
	segments := DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
