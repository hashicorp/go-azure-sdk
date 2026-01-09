package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceIdSecurityBaselineStateId{}

func TestNewDeviceManagementManagedDeviceIdSecurityBaselineStateID(t *testing.T) {
	id := NewDeviceManagementManagedDeviceIdSecurityBaselineStateID("managedDeviceId", "securityBaselineStateId")

	if id.ManagedDeviceId != "managedDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceId")
	}

	if id.SecurityBaselineStateId != "securityBaselineStateId" {
		t.Fatalf("Expected %q but got %q for Segment 'SecurityBaselineStateId'", id.SecurityBaselineStateId, "securityBaselineStateId")
	}
}

func TestFormatDeviceManagementManagedDeviceIdSecurityBaselineStateID(t *testing.T) {
	actual := NewDeviceManagementManagedDeviceIdSecurityBaselineStateID("managedDeviceId", "securityBaselineStateId").ID()
	expected := "/deviceManagement/managedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementManagedDeviceIdSecurityBaselineStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementManagedDeviceIdSecurityBaselineStateId
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
			Input: "/deviceManagement/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/managedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/managedDevices/managedDeviceId/securityBaselineStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/managedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId",
			Expected: &DeviceManagementManagedDeviceIdSecurityBaselineStateId{
				ManagedDeviceId:         "managedDeviceId",
				SecurityBaselineStateId: "securityBaselineStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/managedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementManagedDeviceIdSecurityBaselineStateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

		if actual.SecurityBaselineStateId != v.Expected.SecurityBaselineStateId {
			t.Fatalf("Expected %q but got %q for SecurityBaselineStateId", v.Expected.SecurityBaselineStateId, actual.SecurityBaselineStateId)
		}

	}
}

func TestParseDeviceManagementManagedDeviceIdSecurityBaselineStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementManagedDeviceIdSecurityBaselineStateId
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
			Input: "/deviceManagement/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/managedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/managedDevices/managedDeviceId/securityBaselineStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeId/sEcUrItYbAsElInEsTaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/managedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId",
			Expected: &DeviceManagementManagedDeviceIdSecurityBaselineStateId{
				ManagedDeviceId:         "managedDeviceId",
				SecurityBaselineStateId: "securityBaselineStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/managedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeId/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeId",
			Expected: &DeviceManagementManagedDeviceIdSecurityBaselineStateId{
				ManagedDeviceId:         "mAnAgEdDeViCeId",
				SecurityBaselineStateId: "sEcUrItYbAsElInEsTaTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeId/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementManagedDeviceIdSecurityBaselineStateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

		if actual.SecurityBaselineStateId != v.Expected.SecurityBaselineStateId {
			t.Fatalf("Expected %q but got %q for SecurityBaselineStateId", v.Expected.SecurityBaselineStateId, actual.SecurityBaselineStateId)
		}

	}
}

func TestSegmentsForDeviceManagementManagedDeviceIdSecurityBaselineStateId(t *testing.T) {
	segments := DeviceManagementManagedDeviceIdSecurityBaselineStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementManagedDeviceIdSecurityBaselineStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
