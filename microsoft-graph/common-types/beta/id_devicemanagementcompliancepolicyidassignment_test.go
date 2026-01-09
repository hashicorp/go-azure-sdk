package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCompliancePolicyIdAssignmentId{}

func TestNewDeviceManagementCompliancePolicyIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementCompliancePolicyIdAssignmentID("deviceManagementCompliancePolicyId", "deviceManagementConfigurationPolicyAssignmentId")

	if id.DeviceManagementCompliancePolicyId != "deviceManagementCompliancePolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementCompliancePolicyId'", id.DeviceManagementCompliancePolicyId, "deviceManagementCompliancePolicyId")
	}

	if id.DeviceManagementConfigurationPolicyAssignmentId != "deviceManagementConfigurationPolicyAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationPolicyAssignmentId'", id.DeviceManagementConfigurationPolicyAssignmentId, "deviceManagementConfigurationPolicyAssignmentId")
	}
}

func TestFormatDeviceManagementCompliancePolicyIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementCompliancePolicyIdAssignmentID("deviceManagementCompliancePolicyId", "deviceManagementConfigurationPolicyAssignmentId").ID()
	expected := "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId/assignments/deviceManagementConfigurationPolicyAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementCompliancePolicyIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCompliancePolicyIdAssignmentId
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
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId/assignments/deviceManagementConfigurationPolicyAssignmentId",
			Expected: &DeviceManagementCompliancePolicyIdAssignmentId{
				DeviceManagementCompliancePolicyId:              "deviceManagementCompliancePolicyId",
				DeviceManagementConfigurationPolicyAssignmentId: "deviceManagementConfigurationPolicyAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId/assignments/deviceManagementConfigurationPolicyAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCompliancePolicyIdAssignmentID(v.Input)
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

		if actual.DeviceManagementConfigurationPolicyAssignmentId != v.Expected.DeviceManagementConfigurationPolicyAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationPolicyAssignmentId", v.Expected.DeviceManagementConfigurationPolicyAssignmentId, actual.DeviceManagementConfigurationPolicyAssignmentId)
		}

	}
}

func TestParseDeviceManagementCompliancePolicyIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCompliancePolicyIdAssignmentId
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
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiD/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId/assignments/deviceManagementConfigurationPolicyAssignmentId",
			Expected: &DeviceManagementCompliancePolicyIdAssignmentId{
				DeviceManagementCompliancePolicyId:              "deviceManagementCompliancePolicyId",
				DeviceManagementConfigurationPolicyAssignmentId: "deviceManagementConfigurationPolicyAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId/assignments/deviceManagementConfigurationPolicyAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiD/aSsIgNmEnTs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyAsSiGnMeNtId",
			Expected: &DeviceManagementCompliancePolicyIdAssignmentId{
				DeviceManagementCompliancePolicyId:              "dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiD",
				DeviceManagementConfigurationPolicyAssignmentId: "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiD/aSsIgNmEnTs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCompliancePolicyIdAssignmentIDInsensitively(v.Input)
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

		if actual.DeviceManagementConfigurationPolicyAssignmentId != v.Expected.DeviceManagementConfigurationPolicyAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationPolicyAssignmentId", v.Expected.DeviceManagementConfigurationPolicyAssignmentId, actual.DeviceManagementConfigurationPolicyAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementCompliancePolicyIdAssignmentId(t *testing.T) {
	segments := DeviceManagementCompliancePolicyIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementCompliancePolicyIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
