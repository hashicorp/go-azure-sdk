package rolemanagementcloudpcroleassignmentprincipal

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementCloudPCRoleAssignmentPrincipalId{}

func TestNewRoleManagementCloudPCRoleAssignmentPrincipalID(t *testing.T) {
	id := NewRoleManagementCloudPCRoleAssignmentPrincipalID("unifiedRoleAssignmentMultipleIdValue", "directoryObjectIdValue")

	if id.UnifiedRoleAssignmentMultipleId != "unifiedRoleAssignmentMultipleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleAssignmentMultipleId'", id.UnifiedRoleAssignmentMultipleId, "unifiedRoleAssignmentMultipleIdValue")
	}

	if id.DirectoryObjectId != "directoryObjectIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectIdValue")
	}
}

func TestFormatRoleManagementCloudPCRoleAssignmentPrincipalID(t *testing.T) {
	actual := NewRoleManagementCloudPCRoleAssignmentPrincipalID("unifiedRoleAssignmentMultipleIdValue", "directoryObjectIdValue").ID()
	expected := "/roleManagement/cloudPC/roleAssignments/unifiedRoleAssignmentMultipleIdValue/principals/directoryObjectIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementCloudPCRoleAssignmentPrincipalID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementCloudPCRoleAssignmentPrincipalId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/cloudPC",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/cloudPC/roleAssignments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/cloudPC/roleAssignments/unifiedRoleAssignmentMultipleIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/cloudPC/roleAssignments/unifiedRoleAssignmentMultipleIdValue/principals",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/cloudPC/roleAssignments/unifiedRoleAssignmentMultipleIdValue/principals/directoryObjectIdValue",
			Expected: &RoleManagementCloudPCRoleAssignmentPrincipalId{
				UnifiedRoleAssignmentMultipleId: "unifiedRoleAssignmentMultipleIdValue",
				DirectoryObjectId:               "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/cloudPC/roleAssignments/unifiedRoleAssignmentMultipleIdValue/principals/directoryObjectIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementCloudPCRoleAssignmentPrincipalID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleAssignmentMultipleId != v.Expected.UnifiedRoleAssignmentMultipleId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentMultipleId", v.Expected.UnifiedRoleAssignmentMultipleId, actual.UnifiedRoleAssignmentMultipleId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestParseRoleManagementCloudPCRoleAssignmentPrincipalIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementCloudPCRoleAssignmentPrincipalId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/cloudPC",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/cloudPC/roleAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rOlEaSsIgNmEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/cloudPC/roleAssignments/unifiedRoleAssignmentMultipleIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/cloudPC/roleAssignments/unifiedRoleAssignmentMultipleIdValue/principals",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeIdVaLuE/pRiNcIpAlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/cloudPC/roleAssignments/unifiedRoleAssignmentMultipleIdValue/principals/directoryObjectIdValue",
			Expected: &RoleManagementCloudPCRoleAssignmentPrincipalId{
				UnifiedRoleAssignmentMultipleId: "unifiedRoleAssignmentMultipleIdValue",
				DirectoryObjectId:               "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/cloudPC/roleAssignments/unifiedRoleAssignmentMultipleIdValue/principals/directoryObjectIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeIdVaLuE/pRiNcIpAlS/dIrEcToRyObJeCtIdVaLuE",
			Expected: &RoleManagementCloudPCRoleAssignmentPrincipalId{
				UnifiedRoleAssignmentMultipleId: "uNiFiEdRoLeAsSiGnMeNtMuLtIpLeIdVaLuE",
				DirectoryObjectId:               "dIrEcToRyObJeCtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeIdVaLuE/pRiNcIpAlS/dIrEcToRyObJeCtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementCloudPCRoleAssignmentPrincipalIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleAssignmentMultipleId != v.Expected.UnifiedRoleAssignmentMultipleId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentMultipleId", v.Expected.UnifiedRoleAssignmentMultipleId, actual.UnifiedRoleAssignmentMultipleId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestSegmentsForRoleManagementCloudPCRoleAssignmentPrincipalId(t *testing.T) {
	segments := RoleManagementCloudPCRoleAssignmentPrincipalId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementCloudPCRoleAssignmentPrincipalId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
