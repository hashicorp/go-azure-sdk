package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId{}

func TestNewDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationID("groupPolicyDefinitionId", "groupPolicyPresentationId")

	if id.GroupPolicyDefinitionId != "groupPolicyDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyDefinitionId'", id.GroupPolicyDefinitionId, "groupPolicyDefinitionId")
	}

	if id.GroupPolicyPresentationId != "groupPolicyPresentationId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyPresentationId'", id.GroupPolicyPresentationId, "groupPolicyPresentationId")
	}
}

func TestFormatDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationID("groupPolicyDefinitionId", "groupPolicyPresentationId").ID()
	expected := "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionId/nextVersionDefinition/previousVersionDefinition/presentations/groupPolicyPresentationId"
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
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionId/nextVersionDefinition",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionId/nextVersionDefinition/previousVersionDefinition",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionId/nextVersionDefinition/previousVersionDefinition/presentations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionId/nextVersionDefinition/previousVersionDefinition/presentations/groupPolicyPresentationId",
			Expected: &DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId{
				GroupPolicyDefinitionId:   "groupPolicyDefinitionId",
				GroupPolicyPresentationId: "groupPolicyPresentationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionId/nextVersionDefinition/previousVersionDefinition/presentations/groupPolicyPresentationId/extra",
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
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS/gRoUpPoLiCyDeFiNiTiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionId/nextVersionDefinition",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS/gRoUpPoLiCyDeFiNiTiOnId/nExTvErSiOnDeFiNiTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionId/nextVersionDefinition/previousVersionDefinition",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS/gRoUpPoLiCyDeFiNiTiOnId/nExTvErSiOnDeFiNiTiOn/pReViOuSvErSiOnDeFiNiTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionId/nextVersionDefinition/previousVersionDefinition/presentations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS/gRoUpPoLiCyDeFiNiTiOnId/nExTvErSiOnDeFiNiTiOn/pReViOuSvErSiOnDeFiNiTiOn/pReSeNtAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionId/nextVersionDefinition/previousVersionDefinition/presentations/groupPolicyPresentationId",
			Expected: &DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId{
				GroupPolicyDefinitionId:   "groupPolicyDefinitionId",
				GroupPolicyPresentationId: "groupPolicyPresentationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionId/nextVersionDefinition/previousVersionDefinition/presentations/groupPolicyPresentationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS/gRoUpPoLiCyDeFiNiTiOnId/nExTvErSiOnDeFiNiTiOn/pReViOuSvErSiOnDeFiNiTiOn/pReSeNtAtIoNs/gRoUpPoLiCyPrEsEnTaTiOnId",
			Expected: &DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId{
				GroupPolicyDefinitionId:   "gRoUpPoLiCyDeFiNiTiOnId",
				GroupPolicyPresentationId: "gRoUpPoLiCyPrEsEnTaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS/gRoUpPoLiCyDeFiNiTiOnId/nExTvErSiOnDeFiNiTiOn/pReViOuSvErSiOnDeFiNiTiOn/pReSeNtAtIoNs/gRoUpPoLiCyPrEsEnTaTiOnId/extra",
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
