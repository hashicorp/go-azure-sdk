package grouppolicydefinitionnextversiondefinitionpreviousversiondefinitionpresentation

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId{}

func TestNewDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationID("groupPolicyDefinitionIdValue", "groupPolicyPresentationIdValue")

	if id.GroupPolicyDefinitionId != "groupPolicyDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyDefinitionId'", id.GroupPolicyDefinitionId, "groupPolicyDefinitionIdValue")
	}

	if id.GroupPolicyPresentationId != "groupPolicyPresentationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyPresentationId'", id.GroupPolicyPresentationId, "groupPolicyPresentationIdValue")
	}
}

func TestFormatDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationID("groupPolicyDefinitionIdValue", "groupPolicyPresentationIdValue").ID()
	expected := "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/nextVersionDefinition/previousVersionDefinition/presentations/groupPolicyPresentationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId
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
			Input: "/deviceManagement/groupPolicyDefinitions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/nextVersionDefinition",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/nextVersionDefinition/previousVersionDefinition",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/nextVersionDefinition/previousVersionDefinition/presentations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/nextVersionDefinition/previousVersionDefinition/presentations/groupPolicyPresentationIdValue",
			Expected: &DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId{
				GroupPolicyDefinitionId:   "groupPolicyDefinitionIdValue",
				GroupPolicyPresentationId: "groupPolicyPresentationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/nextVersionDefinition/previousVersionDefinition/presentations/groupPolicyPresentationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupPolicyDefinitionId != v.Expected.GroupPolicyDefinitionId {
			t.Fatalf("Expected %q but got %q for GroupPolicyDefinitionId", v.Expected.GroupPolicyDefinitionId, actual.GroupPolicyDefinitionId)
		}

		if actual.GroupPolicyPresentationId != v.Expected.GroupPolicyPresentationId {
			t.Fatalf("Expected %q but got %q for GroupPolicyPresentationId", v.Expected.GroupPolicyPresentationId, actual.GroupPolicyPresentationId)
		}

	}
}

func TestParseDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId
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
			Input: "/deviceManagement/groupPolicyDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS/gRoUpPoLiCyDeFiNiTiOnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/nextVersionDefinition",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS/gRoUpPoLiCyDeFiNiTiOnIdVaLuE/nExTvErSiOnDeFiNiTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/nextVersionDefinition/previousVersionDefinition",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS/gRoUpPoLiCyDeFiNiTiOnIdVaLuE/nExTvErSiOnDeFiNiTiOn/pReViOuSvErSiOnDeFiNiTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/nextVersionDefinition/previousVersionDefinition/presentations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS/gRoUpPoLiCyDeFiNiTiOnIdVaLuE/nExTvErSiOnDeFiNiTiOn/pReViOuSvErSiOnDeFiNiTiOn/pReSeNtAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/nextVersionDefinition/previousVersionDefinition/presentations/groupPolicyPresentationIdValue",
			Expected: &DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId{
				GroupPolicyDefinitionId:   "groupPolicyDefinitionIdValue",
				GroupPolicyPresentationId: "groupPolicyPresentationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/nextVersionDefinition/previousVersionDefinition/presentations/groupPolicyPresentationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS/gRoUpPoLiCyDeFiNiTiOnIdVaLuE/nExTvErSiOnDeFiNiTiOn/pReViOuSvErSiOnDeFiNiTiOn/pReSeNtAtIoNs/gRoUpPoLiCyPrEsEnTaTiOnIdVaLuE",
			Expected: &DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId{
				GroupPolicyDefinitionId:   "gRoUpPoLiCyDeFiNiTiOnIdVaLuE",
				GroupPolicyPresentationId: "gRoUpPoLiCyPrEsEnTaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS/gRoUpPoLiCyDeFiNiTiOnIdVaLuE/nExTvErSiOnDeFiNiTiOn/pReViOuSvErSiOnDeFiNiTiOn/pReSeNtAtIoNs/gRoUpPoLiCyPrEsEnTaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupPolicyDefinitionId != v.Expected.GroupPolicyDefinitionId {
			t.Fatalf("Expected %q but got %q for GroupPolicyDefinitionId", v.Expected.GroupPolicyDefinitionId, actual.GroupPolicyDefinitionId)
		}

		if actual.GroupPolicyPresentationId != v.Expected.GroupPolicyPresentationId {
			t.Fatalf("Expected %q but got %q for GroupPolicyPresentationId", v.Expected.GroupPolicyPresentationId, actual.GroupPolicyPresentationId)
		}

	}
}

func TestSegmentsForDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId(t *testing.T) {
	segments := DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
