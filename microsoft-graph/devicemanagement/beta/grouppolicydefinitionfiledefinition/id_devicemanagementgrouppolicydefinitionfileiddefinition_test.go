package grouppolicydefinitionfiledefinition

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyDefinitionFileIdDefinitionId{}

func TestNewDeviceManagementGroupPolicyDefinitionFileIdDefinitionID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyDefinitionFileIdDefinitionID("groupPolicyDefinitionFileIdValue", "groupPolicyDefinitionIdValue")

	if id.GroupPolicyDefinitionFileId != "groupPolicyDefinitionFileIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyDefinitionFileId'", id.GroupPolicyDefinitionFileId, "groupPolicyDefinitionFileIdValue")
	}

	if id.GroupPolicyDefinitionId != "groupPolicyDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyDefinitionId'", id.GroupPolicyDefinitionId, "groupPolicyDefinitionIdValue")
	}
}

func TestFormatDeviceManagementGroupPolicyDefinitionFileIdDefinitionID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyDefinitionFileIdDefinitionID("groupPolicyDefinitionFileIdValue", "groupPolicyDefinitionIdValue").ID()
	expected := "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileIdValue/definitions/groupPolicyDefinitionIdValue"
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
			Input: "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileIdValue/definitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileIdValue/definitions/groupPolicyDefinitionIdValue",
			Expected: &DeviceManagementGroupPolicyDefinitionFileIdDefinitionId{
				GroupPolicyDefinitionFileId: "groupPolicyDefinitionFileIdValue",
				GroupPolicyDefinitionId:     "groupPolicyDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileIdValue/definitions/groupPolicyDefinitionIdValue/extra",
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
			Input: "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnFiLeS/gRoUpPoLiCyDeFiNiTiOnFiLeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileIdValue/definitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnFiLeS/gRoUpPoLiCyDeFiNiTiOnFiLeIdVaLuE/dEfInItIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileIdValue/definitions/groupPolicyDefinitionIdValue",
			Expected: &DeviceManagementGroupPolicyDefinitionFileIdDefinitionId{
				GroupPolicyDefinitionFileId: "groupPolicyDefinitionFileIdValue",
				GroupPolicyDefinitionId:     "groupPolicyDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyDefinitionFiles/groupPolicyDefinitionFileIdValue/definitions/groupPolicyDefinitionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnFiLeS/gRoUpPoLiCyDeFiNiTiOnFiLeIdVaLuE/dEfInItIoNs/gRoUpPoLiCyDeFiNiTiOnIdVaLuE",
			Expected: &DeviceManagementGroupPolicyDefinitionFileIdDefinitionId{
				GroupPolicyDefinitionFileId: "gRoUpPoLiCyDeFiNiTiOnFiLeIdVaLuE",
				GroupPolicyDefinitionId:     "gRoUpPoLiCyDeFiNiTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyDeFiNiTiOnFiLeS/gRoUpPoLiCyDeFiNiTiOnFiLeIdVaLuE/dEfInItIoNs/gRoUpPoLiCyDeFiNiTiOnIdVaLuE/extra",
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
