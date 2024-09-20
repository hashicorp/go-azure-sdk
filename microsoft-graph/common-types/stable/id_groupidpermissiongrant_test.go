package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdPermissionGrantId{}

func TestNewGroupIdPermissionGrantID(t *testing.T) {
	id := NewGroupIdPermissionGrantID("groupId", "resourceSpecificPermissionGrantId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.ResourceSpecificPermissionGrantId != "resourceSpecificPermissionGrantId" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceSpecificPermissionGrantId'", id.ResourceSpecificPermissionGrantId, "resourceSpecificPermissionGrantId")
	}
}

func TestFormatGroupIdPermissionGrantID(t *testing.T) {
	actual := NewGroupIdPermissionGrantID("groupId", "resourceSpecificPermissionGrantId").ID()
	expected := "/groups/groupId/permissionGrants/resourceSpecificPermissionGrantId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdPermissionGrantID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdPermissionGrantId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/permissionGrants",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/permissionGrants/resourceSpecificPermissionGrantId",
			Expected: &GroupIdPermissionGrantId{
				GroupId:                           "groupId",
				ResourceSpecificPermissionGrantId: "resourceSpecificPermissionGrantId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/permissionGrants/resourceSpecificPermissionGrantId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdPermissionGrantID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.ResourceSpecificPermissionGrantId != v.Expected.ResourceSpecificPermissionGrantId {
			t.Fatalf("Expected %q but got %q for ResourceSpecificPermissionGrantId", v.Expected.ResourceSpecificPermissionGrantId, actual.ResourceSpecificPermissionGrantId)
		}

	}
}

func TestParseGroupIdPermissionGrantIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdPermissionGrantId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/permissionGrants",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/pErMiSsIoNgRaNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/permissionGrants/resourceSpecificPermissionGrantId",
			Expected: &GroupIdPermissionGrantId{
				GroupId:                           "groupId",
				ResourceSpecificPermissionGrantId: "resourceSpecificPermissionGrantId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/permissionGrants/resourceSpecificPermissionGrantId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/pErMiSsIoNgRaNtS/rEsOuRcEsPeCiFiCpErMiSsIoNgRaNtId",
			Expected: &GroupIdPermissionGrantId{
				GroupId:                           "gRoUpId",
				ResourceSpecificPermissionGrantId: "rEsOuRcEsPeCiFiCpErMiSsIoNgRaNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/pErMiSsIoNgRaNtS/rEsOuRcEsPeCiFiCpErMiSsIoNgRaNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdPermissionGrantIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.ResourceSpecificPermissionGrantId != v.Expected.ResourceSpecificPermissionGrantId {
			t.Fatalf("Expected %q but got %q for ResourceSpecificPermissionGrantId", v.Expected.ResourceSpecificPermissionGrantId, actual.ResourceSpecificPermissionGrantId)
		}

	}
}

func TestSegmentsForGroupIdPermissionGrantId(t *testing.T) {
	segments := GroupIdPermissionGrantId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdPermissionGrantId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
