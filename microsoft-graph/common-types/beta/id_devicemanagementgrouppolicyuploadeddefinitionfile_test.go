package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyUploadedDefinitionFileId{}

func TestNewDeviceManagementGroupPolicyUploadedDefinitionFileID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileId")

	if id.GroupPolicyUploadedDefinitionFileId != "groupPolicyUploadedDefinitionFileId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyUploadedDefinitionFileId'", id.GroupPolicyUploadedDefinitionFileId, "groupPolicyUploadedDefinitionFileId")
	}
}

func TestFormatDeviceManagementGroupPolicyUploadedDefinitionFileID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileId").ID()
	expected := "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementGroupPolicyUploadedDefinitionFileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyUploadedDefinitionFileId
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
			// Valid URI
			Input: "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId",
			Expected: &DeviceManagementGroupPolicyUploadedDefinitionFileId{
				GroupPolicyUploadedDefinitionFileId: "groupPolicyUploadedDefinitionFileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyUploadedDefinitionFileID(v.Input)
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

	}
}

func TestParseDeviceManagementGroupPolicyUploadedDefinitionFileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyUploadedDefinitionFileId
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
			// Valid URI
			Input: "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId",
			Expected: &DeviceManagementGroupPolicyUploadedDefinitionFileId{
				GroupPolicyUploadedDefinitionFileId: "groupPolicyUploadedDefinitionFileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyUploadedDefinitionFiles/groupPolicyUploadedDefinitionFileId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeS/gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeId",
			Expected: &DeviceManagementGroupPolicyUploadedDefinitionFileId{
				GroupPolicyUploadedDefinitionFileId: "gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeS/gRoUpPoLiCyUpLoAdEdDeFiNiTiOnFiLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyUploadedDefinitionFileIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementGroupPolicyUploadedDefinitionFileId(t *testing.T) {
	segments := DeviceManagementGroupPolicyUploadedDefinitionFileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementGroupPolicyUploadedDefinitionFileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
