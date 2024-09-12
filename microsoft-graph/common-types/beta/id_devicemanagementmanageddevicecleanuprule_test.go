package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceCleanupRuleId{}

func TestNewDeviceManagementManagedDeviceCleanupRuleID(t *testing.T) {
	id := NewDeviceManagementManagedDeviceCleanupRuleID("managedDeviceCleanupRuleIdValue")

	if id.ManagedDeviceCleanupRuleId != "managedDeviceCleanupRuleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceCleanupRuleId'", id.ManagedDeviceCleanupRuleId, "managedDeviceCleanupRuleIdValue")
	}
}

func TestFormatDeviceManagementManagedDeviceCleanupRuleID(t *testing.T) {
	actual := NewDeviceManagementManagedDeviceCleanupRuleID("managedDeviceCleanupRuleIdValue").ID()
	expected := "/deviceManagement/managedDeviceCleanupRules/managedDeviceCleanupRuleIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementManagedDeviceCleanupRuleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementManagedDeviceCleanupRuleId
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
			Input: "/deviceManagement/managedDeviceCleanupRules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/managedDeviceCleanupRules/managedDeviceCleanupRuleIdValue",
			Expected: &DeviceManagementManagedDeviceCleanupRuleId{
				ManagedDeviceCleanupRuleId: "managedDeviceCleanupRuleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/managedDeviceCleanupRules/managedDeviceCleanupRuleIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementManagedDeviceCleanupRuleID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedDeviceCleanupRuleId != v.Expected.ManagedDeviceCleanupRuleId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceCleanupRuleId", v.Expected.ManagedDeviceCleanupRuleId, actual.ManagedDeviceCleanupRuleId)
		}

	}
}

func TestParseDeviceManagementManagedDeviceCleanupRuleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementManagedDeviceCleanupRuleId
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
			Input: "/deviceManagement/managedDeviceCleanupRules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeClEaNuPrUlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/managedDeviceCleanupRules/managedDeviceCleanupRuleIdValue",
			Expected: &DeviceManagementManagedDeviceCleanupRuleId{
				ManagedDeviceCleanupRuleId: "managedDeviceCleanupRuleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/managedDeviceCleanupRules/managedDeviceCleanupRuleIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeClEaNuPrUlEs/mAnAgEdDeViCeClEaNuPrUlEiDvAlUe",
			Expected: &DeviceManagementManagedDeviceCleanupRuleId{
				ManagedDeviceCleanupRuleId: "mAnAgEdDeViCeClEaNuPrUlEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeClEaNuPrUlEs/mAnAgEdDeViCeClEaNuPrUlEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementManagedDeviceCleanupRuleIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedDeviceCleanupRuleId != v.Expected.ManagedDeviceCleanupRuleId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceCleanupRuleId", v.Expected.ManagedDeviceCleanupRuleId, actual.ManagedDeviceCleanupRuleId)
		}

	}
}

func TestSegmentsForDeviceManagementManagedDeviceCleanupRuleId(t *testing.T) {
	segments := DeviceManagementManagedDeviceCleanupRuleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementManagedDeviceCleanupRuleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
