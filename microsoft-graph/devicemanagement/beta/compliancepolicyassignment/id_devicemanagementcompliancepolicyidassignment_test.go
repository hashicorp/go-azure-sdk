package compliancepolicyassignment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCompliancePolicyIdAssignmentId{}

func TestNewDeviceManagementCompliancePolicyIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementCompliancePolicyIdAssignmentID("deviceManagementCompliancePolicyIdValue", "deviceManagementConfigurationPolicyAssignmentIdValue")

	if id.DeviceManagementCompliancePolicyId != "deviceManagementCompliancePolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementCompliancePolicyId'", id.DeviceManagementCompliancePolicyId, "deviceManagementCompliancePolicyIdValue")
	}

	if id.DeviceManagementConfigurationPolicyAssignmentId != "deviceManagementConfigurationPolicyAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationPolicyAssignmentId'", id.DeviceManagementConfigurationPolicyAssignmentId, "deviceManagementConfigurationPolicyAssignmentIdValue")
	}
}

func TestFormatDeviceManagementCompliancePolicyIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementCompliancePolicyIdAssignmentID("deviceManagementCompliancePolicyIdValue", "deviceManagementConfigurationPolicyAssignmentIdValue").ID()
	expected := "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/assignments/deviceManagementConfigurationPolicyAssignmentIdValue"
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
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/assignments/deviceManagementConfigurationPolicyAssignmentIdValue",
			Expected: &DeviceManagementCompliancePolicyIdAssignmentId{
				DeviceManagementCompliancePolicyId:              "deviceManagementCompliancePolicyIdValue",
				DeviceManagementConfigurationPolicyAssignmentId: "deviceManagementConfigurationPolicyAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/assignments/deviceManagementConfigurationPolicyAssignmentIdValue/extra",
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
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiDvAlUe/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/assignments/deviceManagementConfigurationPolicyAssignmentIdValue",
			Expected: &DeviceManagementCompliancePolicyIdAssignmentId{
				DeviceManagementCompliancePolicyId:              "deviceManagementCompliancePolicyIdValue",
				DeviceManagementConfigurationPolicyAssignmentId: "deviceManagementConfigurationPolicyAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/assignments/deviceManagementConfigurationPolicyAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiDvAlUe/aSsIgNmEnTs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyAsSiGnMeNtIdVaLuE",
			Expected: &DeviceManagementCompliancePolicyIdAssignmentId{
				DeviceManagementCompliancePolicyId:              "dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiDvAlUe",
				DeviceManagementConfigurationPolicyAssignmentId: "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyAsSiGnMeNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiDvAlUe/aSsIgNmEnTs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyAsSiGnMeNtIdVaLuE/extra",
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
