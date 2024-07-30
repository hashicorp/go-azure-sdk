package devicecompliancepolicyscheduledactionsforrulescheduledactionconfiguration

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{}

func TestNewDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID(t *testing.T) {
	id := NewDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID("deviceCompliancePolicyIdValue", "deviceComplianceScheduledActionForRuleIdValue", "deviceComplianceActionItemIdValue")

	if id.DeviceCompliancePolicyId != "deviceCompliancePolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceCompliancePolicyId'", id.DeviceCompliancePolicyId, "deviceCompliancePolicyIdValue")
	}

	if id.DeviceComplianceScheduledActionForRuleId != "deviceComplianceScheduledActionForRuleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceComplianceScheduledActionForRuleId'", id.DeviceComplianceScheduledActionForRuleId, "deviceComplianceScheduledActionForRuleIdValue")
	}

	if id.DeviceComplianceActionItemId != "deviceComplianceActionItemIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceComplianceActionItemId'", id.DeviceComplianceActionItemId, "deviceComplianceActionItemIdValue")
	}
}

func TestFormatDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID(t *testing.T) {
	actual := NewDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID("deviceCompliancePolicyIdValue", "deviceComplianceScheduledActionForRuleIdValue", "deviceComplianceActionItemIdValue").ID()
	expected := "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/scheduledActionsForRule/deviceComplianceScheduledActionForRuleIdValue/scheduledActionConfigurations/deviceComplianceActionItemIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId
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
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/scheduledActionsForRule",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/scheduledActionsForRule/deviceComplianceScheduledActionForRuleIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/scheduledActionsForRule/deviceComplianceScheduledActionForRuleIdValue/scheduledActionConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/scheduledActionsForRule/deviceComplianceScheduledActionForRuleIdValue/scheduledActionConfigurations/deviceComplianceActionItemIdValue",
			Expected: &DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{
				DeviceCompliancePolicyId:                 "deviceCompliancePolicyIdValue",
				DeviceComplianceScheduledActionForRuleId: "deviceComplianceScheduledActionForRuleIdValue",
				DeviceComplianceActionItemId:             "deviceComplianceActionItemIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/scheduledActionsForRule/deviceComplianceScheduledActionForRuleIdValue/scheduledActionConfigurations/deviceComplianceActionItemIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID(v.Input)
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

		if actual.DeviceComplianceActionItemId != v.Expected.DeviceComplianceActionItemId {
			t.Fatalf("Expected %q but got %q for DeviceComplianceActionItemId", v.Expected.DeviceComplianceActionItemId, actual.DeviceComplianceActionItemId)
		}

	}
}

func TestParseDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId
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
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/scheduledActionsForRule",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiDvAlUe/sChEdUlEdAcTiOnSfOrRuLe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/scheduledActionsForRule/deviceComplianceScheduledActionForRuleIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiDvAlUe/sChEdUlEdAcTiOnSfOrRuLe/dEvIcEcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/scheduledActionsForRule/deviceComplianceScheduledActionForRuleIdValue/scheduledActionConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiDvAlUe/sChEdUlEdAcTiOnSfOrRuLe/dEvIcEcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiDvAlUe/sChEdUlEdAcTiOnCoNfIgUrAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/scheduledActionsForRule/deviceComplianceScheduledActionForRuleIdValue/scheduledActionConfigurations/deviceComplianceActionItemIdValue",
			Expected: &DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{
				DeviceCompliancePolicyId:                 "deviceCompliancePolicyIdValue",
				DeviceComplianceScheduledActionForRuleId: "deviceComplianceScheduledActionForRuleIdValue",
				DeviceComplianceActionItemId:             "deviceComplianceActionItemIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/scheduledActionsForRule/deviceComplianceScheduledActionForRuleIdValue/scheduledActionConfigurations/deviceComplianceActionItemIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiDvAlUe/sChEdUlEdAcTiOnSfOrRuLe/dEvIcEcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiDvAlUe/sChEdUlEdAcTiOnCoNfIgUrAtIoNs/dEvIcEcOmPlIaNcEaCtIoNiTeMiDvAlUe",
			Expected: &DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{
				DeviceCompliancePolicyId:                 "dEvIcEcOmPlIaNcEpOlIcYiDvAlUe",
				DeviceComplianceScheduledActionForRuleId: "dEvIcEcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiDvAlUe",
				DeviceComplianceActionItemId:             "dEvIcEcOmPlIaNcEaCtIoNiTeMiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiDvAlUe/sChEdUlEdAcTiOnSfOrRuLe/dEvIcEcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiDvAlUe/sChEdUlEdAcTiOnCoNfIgUrAtIoNs/dEvIcEcOmPlIaNcEaCtIoNiTeMiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationIDInsensitively(v.Input)
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

		if actual.DeviceComplianceActionItemId != v.Expected.DeviceComplianceActionItemId {
			t.Fatalf("Expected %q but got %q for DeviceComplianceActionItemId", v.Expected.DeviceComplianceActionItemId, actual.DeviceComplianceActionItemId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId(t *testing.T) {
	segments := DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
