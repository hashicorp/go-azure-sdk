package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCompliancePolicyIdScheduledActionsForRuleId{}

func TestNewDeviceManagementCompliancePolicyIdScheduledActionsForRuleID(t *testing.T) {
	id := NewDeviceManagementCompliancePolicyIdScheduledActionsForRuleID("deviceManagementCompliancePolicyIdValue", "deviceManagementComplianceScheduledActionForRuleIdValue")

	if id.DeviceManagementCompliancePolicyId != "deviceManagementCompliancePolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementCompliancePolicyId'", id.DeviceManagementCompliancePolicyId, "deviceManagementCompliancePolicyIdValue")
	}

	if id.DeviceManagementComplianceScheduledActionForRuleId != "deviceManagementComplianceScheduledActionForRuleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementComplianceScheduledActionForRuleId'", id.DeviceManagementComplianceScheduledActionForRuleId, "deviceManagementComplianceScheduledActionForRuleIdValue")
	}
}

func TestFormatDeviceManagementCompliancePolicyIdScheduledActionsForRuleID(t *testing.T) {
	actual := NewDeviceManagementCompliancePolicyIdScheduledActionsForRuleID("deviceManagementCompliancePolicyIdValue", "deviceManagementComplianceScheduledActionForRuleIdValue").ID()
	expected := "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementCompliancePolicyIdScheduledActionsForRuleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCompliancePolicyIdScheduledActionsForRuleId
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
			// Valid URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleIdValue",
			Expected: &DeviceManagementCompliancePolicyIdScheduledActionsForRuleId{
				DeviceManagementCompliancePolicyId:                 "deviceManagementCompliancePolicyIdValue",
				DeviceManagementComplianceScheduledActionForRuleId: "deviceManagementComplianceScheduledActionForRuleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCompliancePolicyIdScheduledActionsForRuleID(v.Input)
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

	}
}

func TestParseDeviceManagementCompliancePolicyIdScheduledActionsForRuleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCompliancePolicyIdScheduledActionsForRuleId
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
			// Valid URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleIdValue",
			Expected: &DeviceManagementCompliancePolicyIdScheduledActionsForRuleId{
				DeviceManagementCompliancePolicyId:                 "deviceManagementCompliancePolicyIdValue",
				DeviceManagementComplianceScheduledActionForRuleId: "deviceManagementComplianceScheduledActionForRuleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/scheduledActionsForRule/deviceManagementComplianceScheduledActionForRuleIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiDvAlUe/sChEdUlEdAcTiOnSfOrRuLe/dEvIcEmAnAgEmEnTcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiDvAlUe",
			Expected: &DeviceManagementCompliancePolicyIdScheduledActionsForRuleId{
				DeviceManagementCompliancePolicyId:                 "dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiDvAlUe",
				DeviceManagementComplianceScheduledActionForRuleId: "dEvIcEmAnAgEmEnTcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiDvAlUe/sChEdUlEdAcTiOnSfOrRuLe/dEvIcEmAnAgEmEnTcOmPlIaNcEsChEdUlEdAcTiOnFoRrUlEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCompliancePolicyIdScheduledActionsForRuleIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementCompliancePolicyIdScheduledActionsForRuleId(t *testing.T) {
	segments := DeviceManagementCompliancePolicyIdScheduledActionsForRuleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementCompliancePolicyIdScheduledActionsForRuleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
