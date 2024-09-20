package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdItemId{}

func TestNewGroupIdSiteIdItemID(t *testing.T) {
	id := NewGroupIdSiteIdItemID("groupId", "siteId", "baseItemId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.SiteId != "siteId" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteId")
	}

	if id.BaseItemId != "baseItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseItemId'", id.BaseItemId, "baseItemId")
	}
}

func TestFormatGroupIdSiteIdItemID(t *testing.T) {
	actual := NewGroupIdSiteIdItemID("groupId", "siteId", "baseItemId").ID()
	expected := "/groups/groupId/sites/siteId/items/baseItemId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdSiteIdItemID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdItemId
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
			Input: "/groups/groupId/sites/siteId/items",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/items/baseItemId",
			Expected: &GroupIdSiteIdItemId{
				GroupId:    "groupId",
				SiteId:     "siteId",
				BaseItemId: "baseItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/items/baseItemId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdItemID(v.Input)
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

		if actual.BaseItemId != v.Expected.BaseItemId {
			t.Fatalf("Expected %q but got %q for BaseItemId", v.Expected.BaseItemId, actual.BaseItemId)
		}

	}
}

func TestParseGroupIdSiteIdItemIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdItemId
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
			Input: "/groups/groupId/sites/siteId/items",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/iTeMs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/items/baseItemId",
			Expected: &GroupIdSiteIdItemId{
				GroupId:    "groupId",
				SiteId:     "siteId",
				BaseItemId: "baseItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/items/baseItemId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/iTeMs/bAsEiTeMiD",
			Expected: &GroupIdSiteIdItemId{
				GroupId:    "gRoUpId",
				SiteId:     "sItEiD",
				BaseItemId: "bAsEiTeMiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/iTeMs/bAsEiTeMiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdItemIDInsensitively(v.Input)
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

		if actual.BaseItemId != v.Expected.BaseItemId {
			t.Fatalf("Expected %q but got %q for BaseItemId", v.Expected.BaseItemId, actual.BaseItemId)
		}

	}
}

func TestSegmentsForGroupIdSiteIdItemId(t *testing.T) {
	segments := GroupIdSiteIdItemId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdSiteIdItemId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
