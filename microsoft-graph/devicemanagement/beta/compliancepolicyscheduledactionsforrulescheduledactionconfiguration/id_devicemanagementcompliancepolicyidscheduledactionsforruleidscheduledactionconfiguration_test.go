package compliancepolicyscheduledactionsforrulescheduledactionconfiguration

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{}

func TestNewDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID(t *testing.T) {
	id := NewDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID("deviceManagementCompliancePolicyIdValue", "deviceManagementComplianceScheduledActionForRuleIdValue", "deviceManagementComplianceActionItemIdValue")

	if id.DeviceManagementCompliancePolicyId != "deviceManagementCompliancePolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementCompliancePolicyId'", id.DeviceManagementCompliancePolicyId, "deviceManagementCompliancePolicyIdValue")
	}

	if id.DeviceManagementComplianceScheduledActionForRuleId != "deviceManagementComplianceScheduledActionForRuleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementComplianceScheduledActionForRuleId'", id.DeviceManagementComplianceScheduledActionForRuleId, "deviceManagementComplianceScheduledActionForRuleIdValue")
	}

	if id.DeviceManagementComplianceActionItemId != "deviceManagementComplianceActionItemIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementComplianceActionItemId'", id.DeviceManagementComplianceActionItemId, "deviceManagementComplianceActionItemIdValue")
	}
}

func TestFormatDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID(t *testing.T) {
	actual := NewDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID("deviceManagementCompliancePolicyIdValue", "deviceManagementComplianceScheduledActionForRuleIdValue", "deviceManagementComplianceActionItemIdValue").ID()
	expected := "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleIdValue/scheduledActionConfigurations/deviceManagementComplianceActionItemIdValue"
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
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/scheduledActionsForRule",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleIdValue/scheduledActionConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleIdValue/scheduledActionConfigurations/deviceManagementComplianceActionItemIdValue",
			Expected: &DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{
				DeviceManagementCompliancePolicyId:                 "deviceManagementCompliancePolicyIdValue",
				DeviceManagementComplianceScheduledActionForRuleId: "deviceManagementComplianceScheduledActionForRuleIdValue",
				DeviceManagementComplianceActionItemId:             "deviceManagementComplianceActionItemIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleIdValue/scheduledActionConfigurations/deviceManagementComplianceActionItemIdValue/extra",
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
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/scheduledActionsForRule",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiDvAlUe/sChEdUlEdAcTiOnSfOrRuLe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiDvAlUe/sChEdUlEdAcTiOnSfOrRuLe/dEvIcEmAnAgEmEnTcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleIdValue/scheduledActionConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiDvAlUe/sChEdUlEdAcTiOnSfOrRuLe/dEvIcEmAnAgEmEnTcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiDvAlUe/sChEdUlEdAcTiOnCoNfIgUrAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleIdValue/scheduledActionConfigurations/deviceManagementComplianceActionItemIdValue",
			Expected: &DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{
				DeviceManagementCompliancePolicyId:                 "deviceManagementCompliancePolicyIdValue",
				DeviceManagementComplianceScheduledActionForRuleId: "deviceManagementComplianceScheduledActionForRuleIdValue",
				DeviceManagementComplianceActionItemId:             "deviceManagementComplianceActionItemIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleIdValue/scheduledActionConfigurations/deviceManagementComplianceActionItemIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiDvAlUe/sChEdUlEdAcTiOnSfOrRuLe/dEvIcEmAnAgEmEnTcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiDvAlUe/sChEdUlEdAcTiOnCoNfIgUrAtIoNs/dEvIcEmAnAgEmEnTcOmPlIaNcEaCtIoNiTeMiDvAlUe",
			Expected: &DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{
				DeviceManagementCompliancePolicyId:                 "dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiDvAlUe",
				DeviceManagementComplianceScheduledActionForRuleId: "dEvIcEmAnAgEmEnTcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiDvAlUe",
				DeviceManagementComplianceActionItemId:             "dEvIcEmAnAgEmEnTcOmPlIaNcEaCtIoNiTeMiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiDvAlUe/sChEdUlEdAcTiOnSfOrRuLe/dEvIcEmAnAgEmEnTcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiDvAlUe/sChEdUlEdAcTiOnCoNfIgUrAtIoNs/dEvIcEmAnAgEmEnTcOmPlIaNcEaCtIoNiTeMiDvAlUe/extra",
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
