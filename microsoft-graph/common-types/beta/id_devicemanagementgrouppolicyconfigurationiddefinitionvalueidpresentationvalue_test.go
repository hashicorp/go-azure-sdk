package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId{}

func TestNewDeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueID("groupPolicyConfigurationId", "groupPolicyDefinitionValueId", "groupPolicyPresentationValueId")

	if id.GroupPolicyConfigurationId != "groupPolicyConfigurationId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyConfigurationId'", id.GroupPolicyConfigurationId, "groupPolicyConfigurationId")
	}

	if id.GroupPolicyDefinitionValueId != "groupPolicyDefinitionValueId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyDefinitionValueId'", id.GroupPolicyDefinitionValueId, "groupPolicyDefinitionValueId")
	}

	if id.GroupPolicyPresentationValueId != "groupPolicyPresentationValueId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyPresentationValueId'", id.GroupPolicyPresentationValueId, "groupPolicyPresentationValueId")
	}
}

func TestFormatDeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueID("groupPolicyConfigurationId", "groupPolicyDefinitionValueId", "groupPolicyPresentationValueId").ID()
	expected := "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId/definitionValues/groupPolicyDefinitionValueId/presentationValues/groupPolicyPresentationValueId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId
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
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId/definitionValues",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId/definitionValues/groupPolicyDefinitionValueId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId/definitionValues/groupPolicyDefinitionValueId/presentationValues",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId/definitionValues/groupPolicyDefinitionValueId/presentationValues/groupPolicyPresentationValueId",
			Expected: &DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId{
				GroupPolicyConfigurationId:     "groupPolicyConfigurationId",
				GroupPolicyDefinitionValueId:   "groupPolicyDefinitionValueId",
				GroupPolicyPresentationValueId: "groupPolicyPresentationValueId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId/definitionValues/groupPolicyDefinitionValueId/presentationValues/groupPolicyPresentationValueId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueID(v.Input)
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

		if actual.GroupPolicyPresentationValueId != v.Expected.GroupPolicyPresentationValueId {
			t.Fatalf("Expected %q but got %q for GroupPolicyPresentationValueId", v.Expected.GroupPolicyPresentationValueId, actual.GroupPolicyPresentationValueId)
		}

	}
}

func TestParseDeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId
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
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs/gRoUpPoLiCyCoNfIgUrAtIoNiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId/definitionValues",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs/gRoUpPoLiCyCoNfIgUrAtIoNiD/dEfInItIoNvAlUeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId/definitionValues/groupPolicyDefinitionValueId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs/gRoUpPoLiCyCoNfIgUrAtIoNiD/dEfInItIoNvAlUeS/gRoUpPoLiCyDeFiNiTiOnVaLuEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId/definitionValues/groupPolicyDefinitionValueId/presentationValues",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs/gRoUpPoLiCyCoNfIgUrAtIoNiD/dEfInItIoNvAlUeS/gRoUpPoLiCyDeFiNiTiOnVaLuEiD/pReSeNtAtIoNvAlUeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId/definitionValues/groupPolicyDefinitionValueId/presentationValues/groupPolicyPresentationValueId",
			Expected: &DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId{
				GroupPolicyConfigurationId:     "groupPolicyConfigurationId",
				GroupPolicyDefinitionValueId:   "groupPolicyDefinitionValueId",
				GroupPolicyPresentationValueId: "groupPolicyPresentationValueId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyConfigurations/groupPolicyConfigurationId/definitionValues/groupPolicyDefinitionValueId/presentationValues/groupPolicyPresentationValueId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs/gRoUpPoLiCyCoNfIgUrAtIoNiD/dEfInItIoNvAlUeS/gRoUpPoLiCyDeFiNiTiOnVaLuEiD/pReSeNtAtIoNvAlUeS/gRoUpPoLiCyPrEsEnTaTiOnVaLuEiD",
			Expected: &DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId{
				GroupPolicyConfigurationId:     "gRoUpPoLiCyCoNfIgUrAtIoNiD",
				GroupPolicyDefinitionValueId:   "gRoUpPoLiCyDeFiNiTiOnVaLuEiD",
				GroupPolicyPresentationValueId: "gRoUpPoLiCyPrEsEnTaTiOnVaLuEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyCoNfIgUrAtIoNs/gRoUpPoLiCyCoNfIgUrAtIoNiD/dEfInItIoNvAlUeS/gRoUpPoLiCyDeFiNiTiOnVaLuEiD/pReSeNtAtIoNvAlUeS/gRoUpPoLiCyPrEsEnTaTiOnVaLuEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueIDInsensitively(v.Input)
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

		if actual.GroupPolicyPresentationValueId != v.Expected.GroupPolicyPresentationValueId {
			t.Fatalf("Expected %q but got %q for GroupPolicyPresentationValueId", v.Expected.GroupPolicyPresentationValueId, actual.GroupPolicyPresentationValueId)
		}

	}
}

func TestSegmentsForDeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId(t *testing.T) {
	segments := DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
