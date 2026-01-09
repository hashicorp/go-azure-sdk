package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId{}

func TestNewDeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationID("groupPolicyUploadedDefinitionFileId", "groupPolicyOperationId")

	if id.GroupPolicyUploadedDefinitionFileId != "groupPolicyUploadedDefinitionFileId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyUploadedDefinitionFileId'", id.GroupPolicyUploadedDefinitionFileId, "groupPolicyUploadedDefinitionFileId")
	}

	if id.GroupPolicyOperationId != "groupPolicyOperationId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyOperationId'", id.GroupPolicyOperationId, "groupPolicyOperationId")
	}
}

func TestFormatDeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationID("groupPolicyUploadedDefinitionFileId", "groupPolicyOperationId").ID()
	expected := "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId/groupPolicyOperations/groupPolicyOperationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId
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
			Input: "/deviceManagement/groupPolicyUploadedDefinitionFiles",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId/groupPolicyOperations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId/groupPolicyOperations/groupPolicyOperationId",
			Expected: &DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId{
				GroupPolicyUploadedDefinitionFileId: "groupPolicyUploadedDefinitionFileId",
				GroupPolicyOperationId:              "groupPolicyOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId/groupPolicyOperations/groupPolicyOperationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupPolicyUploadedDefinitionFileId != v.Expected.GroupPolicyUploadedDefinitionFileId {
			t.Fatalf("Expected %q but got %q for GroupPolicyUploadedDefinitionFileId", v.Expected.GroupPolicyUploadedDefinitionFileId, actual.GroupPolicyUploadedDefinitionFileId)
		}

		if actual.GroupPolicyOperationId != v.Expected.GroupPolicyOperationId {
			t.Fatalf("Expected %q but got %q for GroupPolicyOperationId", v.Expected.GroupPolicyOperationId, actual.GroupPolicyOperationId)
		}

	}
}

func TestParseDeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId
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
			Input: "/deviceManagement/groupPolicyUploadedDefinitionFiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeS/gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId/groupPolicyOperations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeS/gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeId/gRoUpPoLiCyOpErAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId/groupPolicyOperations/groupPolicyOperationId",
			Expected: &DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId{
				GroupPolicyUploadedDefinitionFileId: "groupPolicyUploadedDefinitionFileId",
				GroupPolicyOperationId:              "groupPolicyOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId/groupPolicyOperations/groupPolicyOperationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeS/gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeId/gRoUpPoLiCyOpErAtIoNs/gRoUpPoLiCyOpErAtIoNiD",
			Expected: &DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId{
				GroupPolicyUploadedDefinitionFileId: "gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeId",
				GroupPolicyOperationId:              "gRoUpPoLiCyOpErAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeS/gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeId/gRoUpPoLiCyOpErAtIoNs/gRoUpPoLiCyOpErAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupPolicyUploadedDefinitionFileId != v.Expected.GroupPolicyUploadedDefinitionFileId {
			t.Fatalf("Expected %q but got %q for GroupPolicyUploadedDefinitionFileId", v.Expected.GroupPolicyUploadedDefinitionFileId, actual.GroupPolicyUploadedDefinitionFileId)
		}

		if actual.GroupPolicyOperationId != v.Expected.GroupPolicyOperationId {
			t.Fatalf("Expected %q but got %q for GroupPolicyOperationId", v.Expected.GroupPolicyOperationId, actual.GroupPolicyOperationId)
		}

	}
}

func TestSegmentsForDeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId(t *testing.T) {
	segments := DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
