package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId{}

func TestNewDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationID("groupPolicyDefinitionId", "groupPolicyPresentationId")

	if id.GroupPolicyDefinitionId != "groupPolicyDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyDefinitionId'", id.GroupPolicyDefinitionId, "groupPolicyDefinitionId")
	}

	if id.GroupPolicyPresentationId != "groupPolicyPresentationId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyPresentationId'", id.GroupPolicyPresentationId, "groupPolicyPresentationId")
	}
}

func TestFormatDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationID("groupPolicyDefinitionId", "groupPolicyPresentationId").ID()
	expected := "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionId/nextVersionDefinition/presentations/groupPolicyPresentationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId
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
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionId/nextVersionDefinition/presentations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionId/nextVersionDefinition/presentations/groupPolicyPresentationId",
			Expected: &DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId{
				GroupPolicyDefinitionId:   "groupPolicyDefinitionId",
				GroupPolicyPresentationId: "groupPolicyPresentationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionId/nextVersionDefinition/presentations/groupPolicyPresentationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationID(v.Input)
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

func TestParseDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId
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
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionId/nextVersionDefinition/presentations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS/gRoUpPoLiCyDeFiNiTiOnId/nExTvErSiOnDeFiNiTiOn/pReSeNtAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionId/nextVersionDefinition/presentations/groupPolicyPresentationId",
			Expected: &DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId{
				GroupPolicyDefinitionId:   "groupPolicyDefinitionId",
				GroupPolicyPresentationId: "groupPolicyPresentationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyDefinitions/groupPolicyDefinitionId/nextVersionDefinition/presentations/groupPolicyPresentationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS/gRoUpPoLiCyDeFiNiTiOnId/nExTvErSiOnDeFiNiTiOn/pReSeNtAtIoNs/gRoUpPoLiCyPrEsEnTaTiOnId",
			Expected: &DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId{
				GroupPolicyDefinitionId:   "gRoUpPoLiCyDeFiNiTiOnId",
				GroupPolicyPresentationId: "gRoUpPoLiCyPrEsEnTaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnS/gRoUpPoLiCyDeFiNiTiOnId/nExTvErSiOnDeFiNiTiOn/pReSeNtAtIoNs/gRoUpPoLiCyPrEsEnTaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationIDInsensitively(v.Input)
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

func TestSegmentsForDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId(t *testing.T) {
	segments := DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
