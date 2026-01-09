package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyDefinitionFileIdDefinitionId{}

func TestNewDeviceManagementGroupPolicyDefinitionFileIdDefinitionID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyDefinitionFileIdDefinitionID("groupPolicyDefinitionFileId", "groupPolicyDefinitionId")

	if id.GroupPolicyDefinitionFileId != "groupPolicyDefinitionFileId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyDefinitionFileId'", id.GroupPolicyDefinitionFileId, "groupPolicyDefinitionFileId")
	}

	if id.GroupPolicyDefinitionId != "groupPolicyDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyDefinitionId'", id.GroupPolicyDefinitionId, "groupPolicyDefinitionId")
	}
}

func TestFormatDeviceManagementGroupPolicyDefinitionFileIdDefinitionID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyDefinitionFileIdDefinitionID("groupPolicyDefinitionFileId", "groupPolicyDefinitionId").ID()
	expected := "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileId/definitions/groupPolicyDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementGroupPolicyDefinitionFileIdDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyDefinitionFileIdDefinitionId
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
			Input: "/deviceManagement/groupPolicyDefinitionFiles",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileId/definitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileId/definitions/groupPolicyDefinitionId",
			Expected: &DeviceManagementGroupPolicyDefinitionFileIdDefinitionId{
				GroupPolicyDefinitionFileId: "groupPolicyDefinitionFileId",
				GroupPolicyDefinitionId:     "groupPolicyDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileId/definitions/groupPolicyDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyDefinitionFileIdDefinitionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupPolicyDefinitionFileId != v.Expected.GroupPolicyDefinitionFileId {
			t.Fatalf("Expected %q but got %q for GroupPolicyDefinitionFileId", v.Expected.GroupPolicyDefinitionFileId, actual.GroupPolicyDefinitionFileId)
		}

		if actual.GroupPolicyDefinitionId != v.Expected.GroupPolicyDefinitionId {
			t.Fatalf("Expected %q but got %q for GroupPolicyDefinitionId", v.Expected.GroupPolicyDefinitionId, actual.GroupPolicyDefinitionId)
		}

	}
}

func TestParseDeviceManagementGroupPolicyDefinitionFileIdDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyDefinitionFileIdDefinitionId
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
			Input: "/deviceManagement/groupPolicyDefinitionFiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnFiLeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnFiLeS/gRoUpPoLiCyDeFiNiTiOnFiLeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileId/definitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnFiLeS/gRoUpPoLiCyDeFiNiTiOnFiLeId/dEfInItIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileId/definitions/groupPolicyDefinitionId",
			Expected: &DeviceManagementGroupPolicyDefinitionFileIdDefinitionId{
				GroupPolicyDefinitionFileId: "groupPolicyDefinitionFileId",
				GroupPolicyDefinitionId:     "groupPolicyDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileId/definitions/groupPolicyDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnFiLeS/gRoUpPoLiCyDeFiNiTiOnFiLeId/dEfInItIoNs/gRoUpPoLiCyDeFiNiTiOnId",
			Expected: &DeviceManagementGroupPolicyDefinitionFileIdDefinitionId{
				GroupPolicyDefinitionFileId: "gRoUpPoLiCyDeFiNiTiOnFiLeId",
				GroupPolicyDefinitionId:     "gRoUpPoLiCyDeFiNiTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnFiLeS/gRoUpPoLiCyDeFiNiTiOnFiLeId/dEfInItIoNs/gRoUpPoLiCyDeFiNiTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyDefinitionFileIdDefinitionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupPolicyDefinitionFileId != v.Expected.GroupPolicyDefinitionFileId {
			t.Fatalf("Expected %q but got %q for GroupPolicyDefinitionFileId", v.Expected.GroupPolicyDefinitionFileId, actual.GroupPolicyDefinitionFileId)
		}

		if actual.GroupPolicyDefinitionId != v.Expected.GroupPolicyDefinitionId {
			t.Fatalf("Expected %q but got %q for GroupPolicyDefinitionId", v.Expected.GroupPolicyDefinitionId, actual.GroupPolicyDefinitionId)
		}

	}
}

func TestSegmentsForDeviceManagementGroupPolicyDefinitionFileIdDefinitionId(t *testing.T) {
	segments := DeviceManagementGroupPolicyDefinitionFileIdDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementGroupPolicyDefinitionFileIdDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
