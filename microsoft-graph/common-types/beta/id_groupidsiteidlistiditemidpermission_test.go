package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdListIdItemIdPermissionId{}

func TestNewGroupIdSiteIdListIdItemIdPermissionID(t *testing.T) {
	id := NewGroupIdSiteIdListIdItemIdPermissionID("groupId", "siteId", "listId", "listItemId", "permissionId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.SiteId != "siteId" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteId")
	}

	if id.ListId != "listId" {
		t.Fatalf("Expected %q but got %q for Segment 'ListId'", id.ListId, "listId")
	}

	if id.ListItemId != "listItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'ListItemId'", id.ListItemId, "listItemId")
	}

	if id.PermissionId != "permissionId" {
		t.Fatalf("Expected %q but got %q for Segment 'PermissionId'", id.PermissionId, "permissionId")
	}
}

func TestFormatGroupIdSiteIdListIdItemIdPermissionID(t *testing.T) {
	actual := NewGroupIdSiteIdListIdItemIdPermissionID("groupId", "siteId", "listId", "listItemId", "permissionId").ID()
	expected := "/groups/groupId/sites/siteId/lists/listId/items/listItemId/permissions/permissionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdSiteIdListIdItemIdPermissionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdListIdItemIdPermissionId
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
			Input: "/groups/groupId/sites",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/lists",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/lists/listId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/lists/listId/items",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/lists/listId/items/listItemId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/lists/listId/items/listItemId/permissions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/lists/listId/items/listItemId/permissions/permissionId",
			Expected: &GroupIdSiteIdListIdItemIdPermissionId{
				GroupId:      "groupId",
				SiteId:       "siteId",
				ListId:       "listId",
				ListItemId:   "listItemId",
				PermissionId: "permissionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/lists/listId/items/listItemId/permissions/permissionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdListIdItemIdPermissionID(v.Input)
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

		if actual.SiteId != v.Expected.SiteId {
			t.Fatalf("Expected %q but got %q for SiteId", v.Expected.SiteId, actual.SiteId)
		}

		if actual.ListId != v.Expected.ListId {
			t.Fatalf("Expected %q but got %q for ListId", v.Expected.ListId, actual.ListId)
		}

		if actual.ListItemId != v.Expected.ListItemId {
			t.Fatalf("Expected %q but got %q for ListItemId", v.Expected.ListItemId, actual.ListItemId)
		}

		if actual.PermissionId != v.Expected.PermissionId {
			t.Fatalf("Expected %q but got %q for PermissionId", v.Expected.PermissionId, actual.PermissionId)
		}

	}
}

func TestParseGroupIdSiteIdListIdItemIdPermissionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdListIdItemIdPermissionId
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
			Input: "/groups/groupId/sites",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/lists",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/lIsTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/lists/listId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/lIsTs/lIsTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/lists/listId/items",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/lIsTs/lIsTiD/iTeMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/lists/listId/items/listItemId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/lIsTs/lIsTiD/iTeMs/lIsTiTeMiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/lists/listId/items/listItemId/permissions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/lIsTs/lIsTiD/iTeMs/lIsTiTeMiD/pErMiSsIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/lists/listId/items/listItemId/permissions/permissionId",
			Expected: &GroupIdSiteIdListIdItemIdPermissionId{
				GroupId:      "groupId",
				SiteId:       "siteId",
				ListId:       "listId",
				ListItemId:   "listItemId",
				PermissionId: "permissionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/lists/listId/items/listItemId/permissions/permissionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/lIsTs/lIsTiD/iTeMs/lIsTiTeMiD/pErMiSsIoNs/pErMiSsIoNiD",
			Expected: &GroupIdSiteIdListIdItemIdPermissionId{
				GroupId:      "gRoUpId",
				SiteId:       "sItEiD",
				ListId:       "lIsTiD",
				ListItemId:   "lIsTiTeMiD",
				PermissionId: "pErMiSsIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/lIsTs/lIsTiD/iTeMs/lIsTiTeMiD/pErMiSsIoNs/pErMiSsIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdListIdItemIdPermissionIDInsensitively(v.Input)
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

		if actual.SiteId != v.Expected.SiteId {
			t.Fatalf("Expected %q but got %q for SiteId", v.Expected.SiteId, actual.SiteId)
		}

		if actual.ListId != v.Expected.ListId {
			t.Fatalf("Expected %q but got %q for ListId", v.Expected.ListId, actual.ListId)
		}

		if actual.ListItemId != v.Expected.ListItemId {
			t.Fatalf("Expected %q but got %q for ListItemId", v.Expected.ListItemId, actual.ListItemId)
		}

		if actual.PermissionId != v.Expected.PermissionId {
			t.Fatalf("Expected %q but got %q for PermissionId", v.Expected.PermissionId, actual.PermissionId)
		}

	}
}

func TestSegmentsForGroupIdSiteIdListIdItemIdPermissionId(t *testing.T) {
	segments := GroupIdSiteIdListIdItemIdPermissionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdSiteIdListIdItemIdPermissionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
