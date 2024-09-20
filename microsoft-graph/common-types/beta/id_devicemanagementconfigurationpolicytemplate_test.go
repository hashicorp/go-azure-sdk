package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementConfigurationPolicyTemplateId{}

func TestNewDeviceManagementConfigurationPolicyTemplateID(t *testing.T) {
	id := NewDeviceManagementConfigurationPolicyTemplateID("deviceManagementConfigurationPolicyTemplateId")

	if id.DeviceManagementConfigurationPolicyTemplateId != "deviceManagementConfigurationPolicyTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationPolicyTemplateId'", id.DeviceManagementConfigurationPolicyTemplateId, "deviceManagementConfigurationPolicyTemplateId")
	}
}

func TestFormatDeviceManagementConfigurationPolicyTemplateID(t *testing.T) {
	actual := NewDeviceManagementConfigurationPolicyTemplateID("deviceManagementConfigurationPolicyTemplateId").ID()
	expected := "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementConfigurationPolicyTemplateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementConfigurationPolicyTemplateId
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
			Input: "/deviceManagement/configurationPolicyTemplates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId",
			Expected: &DeviceManagementConfigurationPolicyTemplateId{
				DeviceManagementConfigurationPolicyTemplateId: "deviceManagementConfigurationPolicyTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementConfigurationPolicyTemplateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementConfigurationPolicyTemplateId != v.Expected.DeviceManagementConfigurationPolicyTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationPolicyTemplateId", v.Expected.DeviceManagementConfigurationPolicyTemplateId, actual.DeviceManagementConfigurationPolicyTemplateId)
		}

	}
}

func TestParseDeviceManagementConfigurationPolicyTemplateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementConfigurationPolicyTemplateId
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
			Input: "/deviceManagement/configurationPolicyTemplates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCyTeMpLaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId",
			Expected: &DeviceManagementConfigurationPolicyTemplateId{
				DeviceManagementConfigurationPolicyTemplateId: "deviceManagementConfigurationPolicyTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCyTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyTeMpLaTeId",
			Expected: &DeviceManagementConfigurationPolicyTemplateId{
				DeviceManagementConfigurationPolicyTemplateId: "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyTeMpLaTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCyTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyTeMpLaTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementConfigurationPolicyTemplateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementConfigurationPolicyTemplateId != v.Expected.DeviceManagementConfigurationPolicyTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationPolicyTemplateId", v.Expected.DeviceManagementConfigurationPolicyTemplateId, actual.DeviceManagementConfigurationPolicyTemplateId)
		}

	}
}

func TestSegmentsForDeviceManagementConfigurationPolicyTemplateId(t *testing.T) {
	segments := DeviceManagementConfigurationPolicyTemplateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementConfigurationPolicyTemplateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
