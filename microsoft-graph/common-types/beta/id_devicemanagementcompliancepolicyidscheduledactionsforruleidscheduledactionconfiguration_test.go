package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{}

func TestNewDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID(t *testing.T) {
	id := NewDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID("deviceManagementCompliancePolicyId", "deviceManagementComplianceScheduledActionForRuleId", "deviceManagementComplianceActionItemId")

	if id.DeviceManagementCompliancePolicyId != "deviceManagementCompliancePolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementCompliancePolicyId'", id.DeviceManagementCompliancePolicyId, "deviceManagementCompliancePolicyId")
	}

	if id.DeviceManagementComplianceScheduledActionForRuleId != "deviceManagementComplianceScheduledActionForRuleId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementComplianceScheduledActionForRuleId'", id.DeviceManagementComplianceScheduledActionForRuleId, "deviceManagementComplianceScheduledActionForRuleId")
	}

	if id.DeviceManagementComplianceActionItemId != "deviceManagementComplianceActionItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementComplianceActionItemId'", id.DeviceManagementComplianceActionItemId, "deviceManagementComplianceActionItemId")
	}
}

func TestFormatDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID(t *testing.T) {
	actual := NewDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID("deviceManagementCompliancePolicyId", "deviceManagementComplianceScheduledActionForRuleId", "deviceManagementComplianceActionItemId").ID()
	expected := "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleId/scheduledActionConfigurations/deviceManagementComplianceActionItemId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId
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
			Input: "/deviceManagement/compliancePolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId/scheduledActionsForRule",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleId/scheduledActionConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleId/scheduledActionConfigurations/deviceManagementComplianceActionItemId",
			Expected: &DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{
				DeviceManagementCompliancePolicyId:                 "deviceManagementCompliancePolicyId",
				DeviceManagementComplianceScheduledActionForRuleId: "deviceManagementComplianceScheduledActionForRuleId",
				DeviceManagementComplianceActionItemId:             "deviceManagementComplianceActionItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleId/scheduledActionConfigurations/deviceManagementComplianceActionItemId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementCompliancePolicyId != v.Expected.DeviceManagementCompliancePolicyId {
			t.Fatalf("Expected %q but got %q for DeviceManagementCompliancePolicyId", v.Expected.DeviceManagementCompliancePolicyId, actual.DeviceManagementCompliancePolicyId)
		}

		if actual.DeviceManagementComplianceScheduledActionForRuleId != v.Expected.DeviceManagementComplianceScheduledActionForRuleId {
			t.Fatalf("Expected %q but got %q for DeviceManagementComplianceScheduledActionForRuleId", v.Expected.DeviceManagementComplianceScheduledActionForRuleId, actual.DeviceManagementComplianceScheduledActionForRuleId)
		}

		if actual.DeviceManagementComplianceActionItemId != v.Expected.DeviceManagementComplianceActionItemId {
			t.Fatalf("Expected %q but got %q for DeviceManagementComplianceActionItemId", v.Expected.DeviceManagementComplianceActionItemId, actual.DeviceManagementComplianceActionItemId)
		}

	}
}

func TestParseDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId
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
			Input: "/deviceManagement/compliancePolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId/scheduledActionsForRule",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiD/sChEdUlEdAcTiOnSfOrRuLe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiD/sChEdUlEdAcTiOnSfOrRuLe/dEvIcEmAnAgEmEnTcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleId/scheduledActionConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiD/sChEdUlEdAcTiOnSfOrRuLe/dEvIcEmAnAgEmEnTcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiD/sChEdUlEdAcTiOnCoNfIgUrAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleId/scheduledActionConfigurations/deviceManagementComplianceActionItemId",
			Expected: &DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{
				DeviceManagementCompliancePolicyId:                 "deviceManagementCompliancePolicyId",
				DeviceManagementComplianceScheduledActionForRuleId: "deviceManagementComplianceScheduledActionForRuleId",
				DeviceManagementComplianceActionItemId:             "deviceManagementComplianceActionItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleId/scheduledActionConfigurations/deviceManagementComplianceActionItemId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiD/sChEdUlEdAcTiOnSfOrRuLe/dEvIcEmAnAgEmEnTcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiD/sChEdUlEdAcTiOnCoNfIgUrAtIoNs/dEvIcEmAnAgEmEnTcOmPlIaNcEaCtIoNiTeMiD",
			Expected: &DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{
				DeviceManagementCompliancePolicyId:                 "dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiD",
				DeviceManagementComplianceScheduledActionForRuleId: "dEvIcEmAnAgEmEnTcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiD",
				DeviceManagementComplianceActionItemId:             "dEvIcEmAnAgEmEnTcOmPlIaNcEaCtIoNiTeMiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiD/sChEdUlEdAcTiOnSfOrRuLe/dEvIcEmAnAgEmEnTcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiD/sChEdUlEdAcTiOnCoNfIgUrAtIoNs/dEvIcEmAnAgEmEnTcOmPlIaNcEaCtIoNiTeMiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementCompliancePolicyId != v.Expected.DeviceManagementCompliancePolicyId {
			t.Fatalf("Expected %q but got %q for DeviceManagementCompliancePolicyId", v.Expected.DeviceManagementCompliancePolicyId, actual.DeviceManagementCompliancePolicyId)
		}

		if actual.DeviceManagementComplianceScheduledActionForRuleId != v.Expected.DeviceManagementComplianceScheduledActionForRuleId {
			t.Fatalf("Expected %q but got %q for DeviceManagementComplianceScheduledActionForRuleId", v.Expected.DeviceManagementComplianceScheduledActionForRuleId, actual.DeviceManagementComplianceScheduledActionForRuleId)
		}

		if actual.DeviceManagementComplianceActionItemId != v.Expected.DeviceManagementComplianceActionItemId {
			t.Fatalf("Expected %q but got %q for DeviceManagementComplianceActionItemId", v.Expected.DeviceManagementComplianceActionItemId, actual.DeviceManagementComplianceActionItemId)
		}

	}
}

func TestSegmentsForDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId(t *testing.T) {
	segments := DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
