package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdRecycleBinItemId{}

func TestNewGroupIdSiteIdRecycleBinItemID(t *testing.T) {
	id := NewGroupIdSiteIdRecycleBinItemID("groupId", "siteId", "recycleBinItemId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.SiteId != "siteId" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteId")
	}

	if id.RecycleBinItemId != "recycleBinItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'RecycleBinItemId'", id.RecycleBinItemId, "recycleBinItemId")
	}
}

func TestFormatGroupIdSiteIdRecycleBinItemID(t *testing.T) {
	actual := NewGroupIdSiteIdRecycleBinItemID("groupId", "siteId", "recycleBinItemId").ID()
	expected := "/groups/groupId/sites/siteId/recycleBin/items/recycleBinItemId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdSiteIdRecycleBinItemID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdRecycleBinItemId
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
			Input: "/groups/groupId/sites/siteId/recycleBin",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/recycleBin/items",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/recycleBin/items/recycleBinItemId",
			Expected: &GroupIdSiteIdRecycleBinItemId{
				GroupId:          "groupId",
				SiteId:           "siteId",
				RecycleBinItemId: "recycleBinItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/recycleBin/items/recycleBinItemId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdRecycleBinItemID(v.Input)
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

		if actual.RecycleBinItemId != v.Expected.RecycleBinItemId {
			t.Fatalf("Expected %q but got %q for RecycleBinItemId", v.Expected.RecycleBinItemId, actual.RecycleBinItemId)
		}

	}
}

func TestParseGroupIdSiteIdRecycleBinItemIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdRecycleBinItemId
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
			Input: "/groups/groupId/sites/siteId/recycleBin",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/rEcYcLeBiN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/recycleBin/items",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/rEcYcLeBiN/iTeMs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/recycleBin/items/recycleBinItemId",
			Expected: &GroupIdSiteIdRecycleBinItemId{
				GroupId:          "groupId",
				SiteId:           "siteId",
				RecycleBinItemId: "recycleBinItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/recycleBin/items/recycleBinItemId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/rEcYcLeBiN/iTeMs/rEcYcLeBiNiTeMiD",
			Expected: &GroupIdSiteIdRecycleBinItemId{
				GroupId:          "gRoUpId",
				SiteId:           "sItEiD",
				RecycleBinItemId: "rEcYcLeBiNiTeMiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/rEcYcLeBiN/iTeMs/rEcYcLeBiNiTeMiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdRecycleBinItemIDInsensitively(v.Input)
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

		if actual.RecycleBinItemId != v.Expected.RecycleBinItemId {
			t.Fatalf("Expected %q but got %q for RecycleBinItemId", v.Expected.RecycleBinItemId, actual.RecycleBinItemId)
		}

	}
}

func TestSegmentsForGroupIdSiteIdRecycleBinItemId(t *testing.T) {
	segments := GroupIdSiteIdRecycleBinItemId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdSiteIdRecycleBinItemId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
