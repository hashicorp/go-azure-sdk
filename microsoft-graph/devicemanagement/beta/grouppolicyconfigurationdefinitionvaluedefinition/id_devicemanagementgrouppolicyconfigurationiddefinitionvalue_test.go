package grouppolicyconfigurationdefinitionvaluedefinition

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyConfigurationIdDefinitionValueId{}

func TestNewDeviceManagementGroupPolicyConfigurationIdDefinitionValueID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyConfigurationIdDefinitionValueID("groupPolicyConfigurationIdValue", "groupPolicyDefinitionValueIdValue")

	if id.GroupPolicyConfigurationId != "groupPolicyConfigurationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyConfigurationId'", id.GroupPolicyConfigurationId, "groupPolicyConfigurationIdValue")
	}

	if id.GroupPolicyDefinitionValueId != "groupPolicyDefinitionValueIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyDefinitionValueId'", id.GroupPolicyDefinitionValueId, "groupPolicyDefinitionValueIdValue")
	}
}

func TestFormatDeviceManagementGroupPolicyConfigurationIdDefinitionValueID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyConfigurationIdDefinitionValueID("groupPolicyConfigurationIdValue", "groupPolicyDefinitionValueIdValue").ID()
	expected := "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue/definitionValues/groupPolicyDefinitionValueIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementGroupPolicyConfigurationIdDefinitionValueID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyConfigurationIdDefinitionValueId
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
			Input: "/deviceManagement/groupPolicyConfigurations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue/definitionValues",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue/definitionValues/groupPolicyDefinitionValueIdValue",
			Expected: &DeviceManagementGroupPolicyConfigurationIdDefinitionValueId{
				GroupPolicyConfigurationId:   "groupPolicyConfigurationIdValue",
				GroupPolicyDefinitionValueId: "groupPolicyDefinitionValueIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue/definitionValues/groupPolicyDefinitionValueIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyConfigurationIdDefinitionValueID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupPolicyConfigurationId != v.Expected.GroupPolicyConfigurationId {
			t.Fatalf("Expected %q but got %q for GroupPolicyConfigurationId", v.Expected.GroupPolicyConfigurationId, actual.GroupPolicyConfigurationId)
		}

		if actual.GroupPolicyDefinitionValueId != v.Expected.GroupPolicyDefinitionValueId {
			t.Fatalf("Expected %q but got %q for GroupPolicyDefinitionValueId", v.Expected.GroupPolicyDefinitionValueId, actual.GroupPolicyDefinitionValueId)
		}

	}
}

func TestParseDeviceManagementGroupPolicyConfigurationIdDefinitionValueIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyConfigurationIdDefinitionValueId
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
			Input: "/deviceManagement/groupPolicyConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs/gRoUpPoLiCyCoNfIgUrAtIoNiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue/definitionValues",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs/gRoUpPoLiCyCoNfIgUrAtIoNiDvAlUe/dEfInItIoNvAlUeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue/definitionValues/groupPolicyDefinitionValueIdValue",
			Expected: &DeviceManagementGroupPolicyConfigurationIdDefinitionValueId{
				GroupPolicyConfigurationId:   "groupPolicyConfigurationIdValue",
				GroupPolicyDefinitionValueId: "groupPolicyDefinitionValueIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationIdValue/definitionValues/groupPolicyDefinitionValueIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs/gRoUpPoLiCyCoNfIgUrAtIoNiDvAlUe/dEfInItIoNvAlUeS/gRoUpPoLiCyDeFiNiTiOnVaLuEiDvAlUe",
			Expected: &DeviceManagementGroupPolicyConfigurationIdDefinitionValueId{
				GroupPolicyConfigurationId:   "gRoUpPoLiCyCoNfIgUrAtIoNiDvAlUe",
				GroupPolicyDefinitionValueId: "gRoUpPoLiCyDeFiNiTiOnVaLuEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs/gRoUpPoLiCyCoNfIgUrAtIoNiDvAlUe/dEfInItIoNvAlUeS/gRoUpPoLiCyDeFiNiTiOnVaLuEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyConfigurationIdDefinitionValueIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupPolicyConfigurationId != v.Expected.GroupPolicyConfigurationId {
			t.Fatalf("Expected %q but got %q for GroupPolicyConfigurationId", v.Expected.GroupPolicyConfigurationId, actual.GroupPolicyConfigurationId)
		}

		if actual.GroupPolicyDefinitionValueId != v.Expected.GroupPolicyDefinitionValueId {
			t.Fatalf("Expected %q but got %q for GroupPolicyDefinitionValueId", v.Expected.GroupPolicyDefinitionValueId, actual.GroupPolicyDefinitionValueId)
		}

	}
}

func TestSegmentsForDeviceManagementGroupPolicyConfigurationIdDefinitionValueId(t *testing.T) {
	segments := DeviceManagementGroupPolicyConfigurationIdDefinitionValueId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementGroupPolicyConfigurationIdDefinitionValueId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
