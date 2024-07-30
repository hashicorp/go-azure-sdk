package grouppolicydefinitionpreviousversiondefinitionnextversiondefinitionpresentation

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId{}

func TestNewDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationID("groupPolicyDefinitionIdValue", "groupPolicyPresentationIdValue")

	if id.GroupPolicyDefinitionId != "groupPolicyDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyDefinitionId'", id.GroupPolicyDefinitionId, "groupPolicyDefinitionIdValue")
	}

	if id.GroupPolicyPresentationId != "groupPolicyPresentationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyPresentationId'", id.GroupPolicyPresentationId, "groupPolicyPresentationIdValue")
	}
}

func TestFormatDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationID("groupPolicyDefinitionIdValue", "groupPolicyPresentationIdValue").ID()
	expected := "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/previousVersionDefinition/nextVersionDefinition/presentations/groupPolicyPresentationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId
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
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/previousVersionDefinition",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/previousVersionDefinition/nextVersionDefinition",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/previousVersionDefinition/nextVersionDefinition/presentations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/previousVersionDefinition/nextVersionDefinition/presentations/groupPolicyPresentationIdValue",
			Expected: &DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId{
				GroupPolicyDefinitionId:   "groupPolicyDefinitionIdValue",
				GroupPolicyPresentationId: "groupPolicyPresentationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/previousVersionDefinition/nextVersionDefinition/presentations/groupPolicyPresentationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationID(v.Input)
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

func TestParseDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId
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
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/previousVersionDefinition",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS/gRoUpPoLiCyDeFiNiTiOnIdVaLuE/pReViOuSvErSiOnDeFiNiTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/previousVersionDefinition/nextVersionDefinition",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS/gRoUpPoLiCyDeFiNiTiOnIdVaLuE/pReViOuSvErSiOnDeFiNiTiOn/nExTvErSiOnDeFiNiTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/previousVersionDefinition/nextVersionDefinition/presentations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS/gRoUpPoLiCyDeFiNiTiOnIdVaLuE/pReViOuSvErSiOnDeFiNiTiOn/nExTvErSiOnDeFiNiTiOn/pReSeNtAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/previousVersionDefinition/nextVersionDefinition/presentations/groupPolicyPresentationIdValue",
			Expected: &DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId{
				GroupPolicyDefinitionId:   "groupPolicyDefinitionIdValue",
				GroupPolicyPresentationId: "groupPolicyPresentationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionIdValue/previousVersionDefinition/nextVersionDefinition/presentations/groupPolicyPresentationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS/gRoUpPoLiCyDeFiNiTiOnIdVaLuE/pReViOuSvErSiOnDeFiNiTiOn/nExTvErSiOnDeFiNiTiOn/pReSeNtAtIoNs/gRoUpPoLiCyPrEsEnTaTiOnIdVaLuE",
			Expected: &DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId{
				GroupPolicyDefinitionId:   "gRoUpPoLiCyDeFiNiTiOnIdVaLuE",
				GroupPolicyPresentationId: "gRoUpPoLiCyPrEsEnTaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS/gRoUpPoLiCyDeFiNiTiOnIdVaLuE/pReViOuSvErSiOnDeFiNiTiOn/nExTvErSiOnDeFiNiTiOn/pReSeNtAtIoNs/gRoUpPoLiCyPrEsEnTaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationIDInsensitively(v.Input)
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

func TestSegmentsForDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId(t *testing.T) {
	segments := DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
