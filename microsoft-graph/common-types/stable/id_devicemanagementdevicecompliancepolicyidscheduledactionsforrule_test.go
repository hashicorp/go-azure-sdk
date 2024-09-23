package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId{}

func TestNewDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleID(t *testing.T) {
	id := NewDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleID("deviceCompliancePolicyId", "deviceComplianceScheduledActionForRuleId")

	if id.DeviceCompliancePolicyId != "deviceCompliancePolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceCompliancePolicyId'", id.DeviceCompliancePolicyId, "deviceCompliancePolicyId")
	}

	if id.DeviceComplianceScheduledActionForRuleId != "deviceComplianceScheduledActionForRuleId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceComplianceScheduledActionForRuleId'", id.DeviceComplianceScheduledActionForRuleId, "deviceComplianceScheduledActionForRuleId")
	}
}

func TestFormatDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleID(t *testing.T) {
	actual := NewDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleID("deviceCompliancePolicyId", "deviceComplianceScheduledActionForRuleId").ID()
	expected := "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/scheduledActionsForRule/deviceComplianceScheduledActionForRuleId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId
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
			Input: "/deviceManagement/deviceCompliancePolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/scheduledActionsForRule",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/scheduledActionsForRule/deviceComplianceScheduledActionForRuleId",
			Expected: &DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId{
				DeviceCompliancePolicyId:                 "deviceCompliancePolicyId",
				DeviceComplianceScheduledActionForRuleId: "deviceComplianceScheduledActionForRuleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/scheduledActionsForRule/deviceComplianceScheduledActionForRuleId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceCompliancePolicyId != v.Expected.DeviceCompliancePolicyId {
			t.Fatalf("Expected %q but got %q for DeviceCompliancePolicyId", v.Expected.DeviceCompliancePolicyId, actual.DeviceCompliancePolicyId)
		}

		if actual.DeviceComplianceScheduledActionForRuleId != v.Expected.DeviceComplianceScheduledActionForRuleId {
			t.Fatalf("Expected %q but got %q for DeviceComplianceScheduledActionForRuleId", v.Expected.DeviceComplianceScheduledActionForRuleId, actual.DeviceComplianceScheduledActionForRuleId)
		}

	}
}

func TestParseDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId
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
			Input: "/deviceManagement/deviceCompliancePolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/scheduledActionsForRule",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiD/sChEdUlEdAcTiOnSfOrRuLe",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/scheduledActionsForRule/deviceComplianceScheduledActionForRuleId",
			Expected: &DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId{
				DeviceCompliancePolicyId:                 "deviceCompliancePolicyId",
				DeviceComplianceScheduledActionForRuleId: "deviceComplianceScheduledActionForRuleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/scheduledActionsForRule/deviceComplianceScheduledActionForRuleId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiD/sChEdUlEdAcTiOnSfOrRuLe/dEvIcEcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiD",
			Expected: &DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId{
				DeviceCompliancePolicyId:                 "dEvIcEcOmPlIaNcEpOlIcYiD",
				DeviceComplianceScheduledActionForRuleId: "dEvIcEcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiD/sChEdUlEdAcTiOnSfOrRuLe/dEvIcEcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceCompliancePolicyId != v.Expected.DeviceCompliancePolicyId {
			t.Fatalf("Expected %q but got %q for DeviceCompliancePolicyId", v.Expected.DeviceCompliancePolicyId, actual.DeviceCompliancePolicyId)
		}

		if actual.DeviceComplianceScheduledActionForRuleId != v.Expected.DeviceComplianceScheduledActionForRuleId {
			t.Fatalf("Expected %q but got %q for DeviceComplianceScheduledActionForRuleId", v.Expected.DeviceComplianceScheduledActionForRuleId, actual.DeviceComplianceScheduledActionForRuleId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId(t *testing.T) {
	segments := DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
