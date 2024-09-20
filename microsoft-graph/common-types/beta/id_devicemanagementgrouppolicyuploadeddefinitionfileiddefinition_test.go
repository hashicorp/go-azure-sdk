package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId{}

func TestNewDeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionID("groupPolicyUploadedDefinitionFileId", "groupPolicyDefinitionId")

	if id.GroupPolicyUploadedDefinitionFileId != "groupPolicyUploadedDefinitionFileId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyUploadedDefinitionFileId'", id.GroupPolicyUploadedDefinitionFileId, "groupPolicyUploadedDefinitionFileId")
	}

	if id.GroupPolicyDefinitionId != "groupPolicyDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyDefinitionId'", id.GroupPolicyDefinitionId, "groupPolicyDefinitionId")
	}
}

func TestFormatDeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionID("groupPolicyUploadedDefinitionFileId", "groupPolicyDefinitionId").ID()
	expected := "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId/definitions/groupPolicyDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId
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
			Input: "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId/definitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId/definitions/groupPolicyDefinitionId",
			Expected: &DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId{
				GroupPolicyUploadedDefinitionFileId: "groupPolicyUploadedDefinitionFileId",
				GroupPolicyDefinitionId:             "groupPolicyDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId/definitions/groupPolicyDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionID(v.Input)
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

		if actual.GroupPolicyDefinitionId != v.Expected.GroupPolicyDefinitionId {
			t.Fatalf("Expected %q but got %q for GroupPolicyDefinitionId", v.Expected.GroupPolicyDefinitionId, actual.GroupPolicyDefinitionId)
		}

	}
}

func TestParseDeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId
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
			Input: "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId/definitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeS/gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeId/dEfInItIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId/definitions/groupPolicyDefinitionId",
			Expected: &DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId{
				GroupPolicyUploadedDefinitionFileId: "groupPolicyUploadedDefinitionFileId",
				GroupPolicyDefinitionId:             "groupPolicyDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId/definitions/groupPolicyDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeS/gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeId/dEfInItIoNs/gRoUpPoLiCyDeFiNiTiOnId",
			Expected: &DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId{
				GroupPolicyUploadedDefinitionFileId: "gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeId",
				GroupPolicyDefinitionId:             "gRoUpPoLiCyDeFiNiTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeS/gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeId/dEfInItIoNs/gRoUpPoLiCyDeFiNiTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionIDInsensitively(v.Input)
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

		if actual.GroupPolicyDefinitionId != v.Expected.GroupPolicyDefinitionId {
			t.Fatalf("Expected %q but got %q for GroupPolicyDefinitionId", v.Expected.GroupPolicyDefinitionId, actual.GroupPolicyDefinitionId)
		}

	}
}

func TestSegmentsForDeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId(t *testing.T) {
	segments := DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
