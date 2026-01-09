package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementConfigurationPolicyIdAssignmentId{}

func TestNewDeviceManagementConfigurationPolicyIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementConfigurationPolicyIdAssignmentID("deviceManagementConfigurationPolicyId", "deviceManagementConfigurationPolicyAssignmentId")

	if id.DeviceManagementConfigurationPolicyId != "deviceManagementConfigurationPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationPolicyId'", id.DeviceManagementConfigurationPolicyId, "deviceManagementConfigurationPolicyId")
	}

	if id.DeviceManagementConfigurationPolicyAssignmentId != "deviceManagementConfigurationPolicyAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationPolicyAssignmentId'", id.DeviceManagementConfigurationPolicyAssignmentId, "deviceManagementConfigurationPolicyAssignmentId")
	}
}

func TestFormatDeviceManagementConfigurationPolicyIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementConfigurationPolicyIdAssignmentID("deviceManagementConfigurationPolicyId", "deviceManagementConfigurationPolicyAssignmentId").ID()
	expected := "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyId/assignments/deviceManagementConfigurationPolicyAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementConfigurationPolicyIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementConfigurationPolicyIdAssignmentId
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
			Input: "/deviceManagement/configurationPolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyId/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyId/assignments/deviceManagementConfigurationPolicyAssignmentId",
			Expected: &DeviceManagementConfigurationPolicyIdAssignmentId{
				DeviceManagementConfigurationPolicyId:           "deviceManagementConfigurationPolicyId",
				DeviceManagementConfigurationPolicyAssignmentId: "deviceManagementConfigurationPolicyAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyId/assignments/deviceManagementConfigurationPolicyAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementConfigurationPolicyIdAssignmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementConfigurationPolicyId != v.Expected.DeviceManagementConfigurationPolicyId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationPolicyId", v.Expected.DeviceManagementConfigurationPolicyId, actual.DeviceManagementConfigurationPolicyId)
		}

		if actual.DeviceManagementConfigurationPolicyAssignmentId != v.Expected.DeviceManagementConfigurationPolicyAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationPolicyAssignmentId", v.Expected.DeviceManagementConfigurationPolicyAssignmentId, actual.DeviceManagementConfigurationPolicyAssignmentId)
		}

	}
}

func TestParseDeviceManagementConfigurationPolicyIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementConfigurationPolicyIdAssignmentId
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
			Input: "/deviceManagement/configurationPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCiEs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyId/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCiEs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyId/assignments/deviceManagementConfigurationPolicyAssignmentId",
			Expected: &DeviceManagementConfigurationPolicyIdAssignmentId{
				DeviceManagementConfigurationPolicyId:           "deviceManagementConfigurationPolicyId",
				DeviceManagementConfigurationPolicyAssignmentId: "deviceManagementConfigurationPolicyAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyId/assignments/deviceManagementConfigurationPolicyAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCiEs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId/aSsIgNmEnTs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyAsSiGnMeNtId",
			Expected: &DeviceManagementConfigurationPolicyIdAssignmentId{
				DeviceManagementConfigurationPolicyId:           "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId",
				DeviceManagementConfigurationPolicyAssignmentId: "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCiEs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId/aSsIgNmEnTs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementConfigurationPolicyIdAssignmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementConfigurationPolicyId != v.Expected.DeviceManagementConfigurationPolicyId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationPolicyId", v.Expected.DeviceManagementConfigurationPolicyId, actual.DeviceManagementConfigurationPolicyId)
		}

		if actual.DeviceManagementConfigurationPolicyAssignmentId != v.Expected.DeviceManagementConfigurationPolicyAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationPolicyAssignmentId", v.Expected.DeviceManagementConfigurationPolicyAssignmentId, actual.DeviceManagementConfigurationPolicyAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementConfigurationPolicyIdAssignmentId(t *testing.T) {
	segments := DeviceManagementConfigurationPolicyIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementConfigurationPolicyIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
