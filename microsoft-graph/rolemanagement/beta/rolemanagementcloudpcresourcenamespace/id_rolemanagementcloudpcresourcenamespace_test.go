package rolemanagementcloudpcresourcenamespace

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementCloudPCResourceNamespaceId{}

func TestNewRoleManagementCloudPCResourceNamespaceID(t *testing.T) {
	id := NewRoleManagementCloudPCResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

	if id.UnifiedRbacResourceNamespaceId != "unifiedRbacResourceNamespaceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRbacResourceNamespaceId'", id.UnifiedRbacResourceNamespaceId, "unifiedRbacResourceNamespaceIdValue")
	}
}

func TestFormatRoleManagementCloudPCResourceNamespaceID(t *testing.T) {
	actual := NewRoleManagementCloudPCResourceNamespaceID("unifiedRbacResourceNamespaceIdValue").ID()
	expected := "/roleManagement/cloudPC/resourceNamespaces/unifiedRbacResourceNamespaceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementCloudPCResourceNamespaceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementCloudPCResourceNamespaceId
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
			Input: "/roleManagement/cloudPC/resourceNamespaces",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/cloudPC/resourceNamespaces/unifiedRbacResourceNamespaceIdValue",
			Expected: &RoleManagementCloudPCResourceNamespaceId{
				UnifiedRbacResourceNamespaceId: "unifiedRbacResourceNamespaceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/cloudPC/resourceNamespaces/unifiedRbacResourceNamespaceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementCloudPCResourceNamespaceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRbacResourceNamespaceId != v.Expected.UnifiedRbacResourceNamespaceId {
			t.Fatalf("Expected %q but got %q for UnifiedRbacResourceNamespaceId", v.Expected.UnifiedRbacResourceNamespaceId, actual.UnifiedRbacResourceNamespaceId)
		}

	}
}

func TestParseRoleManagementCloudPCResourceNamespaceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementCloudPCResourceNamespaceId
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
			Input: "/roleManagement/cloudPC/resourceNamespaces",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rEsOuRcEnAmEsPaCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/cloudPC/resourceNamespaces/unifiedRbacResourceNamespaceIdValue",
			Expected: &RoleManagementCloudPCResourceNamespaceId{
				UnifiedRbacResourceNamespaceId: "unifiedRbacResourceNamespaceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/cloudPC/resourceNamespaces/unifiedRbacResourceNamespaceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rEsOuRcEnAmEsPaCeS/uNiFiEdRbAcReSoUrCeNaMeSpAcEiDvAlUe",
			Expected: &RoleManagementCloudPCResourceNamespaceId{
				UnifiedRbacResourceNamespaceId: "uNiFiEdRbAcReSoUrCeNaMeSpAcEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rEsOuRcEnAmEsPaCeS/uNiFiEdRbAcReSoUrCeNaMeSpAcEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementCloudPCResourceNamespaceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRbacResourceNamespaceId != v.Expected.UnifiedRbacResourceNamespaceId {
			t.Fatalf("Expected %q but got %q for UnifiedRbacResourceNamespaceId", v.Expected.UnifiedRbacResourceNamespaceId, actual.UnifiedRbacResourceNamespaceId)
		}

	}
}

func TestSegmentsForRoleManagementCloudPCResourceNamespaceId(t *testing.T) {
	segments := RoleManagementCloudPCResourceNamespaceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementCloudPCResourceNamespaceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
