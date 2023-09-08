package groupsitepermission

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSitePermissionId{}

func TestNewGroupSitePermissionID(t *testing.T) {
	id := NewGroupSitePermissionID("groupIdValue", "siteIdValue", "permissionIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.SiteId != "siteIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteIdValue")
	}

	if id.PermissionId != "permissionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PermissionId'", id.PermissionId, "permissionIdValue")
	}
}

func TestFormatGroupSitePermissionID(t *testing.T) {
	actual := NewGroupSitePermissionID("groupIdValue", "siteIdValue", "permissionIdValue").ID()
	expected := "/groups/groupIdValue/sites/siteIdValue/permissions/permissionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupSitePermissionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSitePermissionId
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
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/permissions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/permissions/permissionIdValue",
			Expected: &GroupSitePermissionId{
				GroupId:      "groupIdValue",
				SiteId:       "siteIdValue",
				PermissionId: "permissionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/permissions/permissionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSitePermissionID(v.Input)
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

		if actual.PermissionId != v.Expected.PermissionId {
			t.Fatalf("Expected %q but got %q for PermissionId", v.Expected.PermissionId, actual.PermissionId)
		}

	}
}

func TestParseGroupSitePermissionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSitePermissionId
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
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/permissions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/pErMiSsIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/permissions/permissionIdValue",
			Expected: &GroupSitePermissionId{
				GroupId:      "groupIdValue",
				SiteId:       "siteIdValue",
				PermissionId: "permissionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/permissions/permissionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/pErMiSsIoNs/pErMiSsIoNiDvAlUe",
			Expected: &GroupSitePermissionId{
				GroupId:      "gRoUpIdVaLuE",
				SiteId:       "sItEiDvAlUe",
				PermissionId: "pErMiSsIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/pErMiSsIoNs/pErMiSsIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSitePermissionIDInsensitively(v.Input)
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

		if actual.PermissionId != v.Expected.PermissionId {
			t.Fatalf("Expected %q but got %q for PermissionId", v.Expected.PermissionId, actual.PermissionId)
		}

	}
}

func TestSegmentsForGroupSitePermissionId(t *testing.T) {
	segments := GroupSitePermissionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupSitePermissionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
