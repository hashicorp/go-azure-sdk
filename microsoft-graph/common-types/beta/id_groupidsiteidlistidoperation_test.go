package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdListIdOperationId{}

func TestNewGroupIdSiteIdListIdOperationID(t *testing.T) {
	id := NewGroupIdSiteIdListIdOperationID("groupId", "siteId", "listId", "richLongRunningOperationId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.SiteId != "siteId" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteId")
	}

	if id.ListId != "listId" {
		t.Fatalf("Expected %q but got %q for Segment 'ListId'", id.ListId, "listId")
	}

	if id.RichLongRunningOperationId != "richLongRunningOperationId" {
		t.Fatalf("Expected %q but got %q for Segment 'RichLongRunningOperationId'", id.RichLongRunningOperationId, "richLongRunningOperationId")
	}
}

func TestFormatGroupIdSiteIdListIdOperationID(t *testing.T) {
	actual := NewGroupIdSiteIdListIdOperationID("groupId", "siteId", "listId", "richLongRunningOperationId").ID()
	expected := "/groups/groupId/sites/siteId/lists/listId/operations/richLongRunningOperationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdSiteIdListIdOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdListIdOperationId
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
			Input: "/groups/groupId/sites/siteId/lists/listId/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/lists/listId/operations/richLongRunningOperationId",
			Expected: &GroupIdSiteIdListIdOperationId{
				GroupId:                    "groupId",
				SiteId:                     "siteId",
				ListId:                     "listId",
				RichLongRunningOperationId: "richLongRunningOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/lists/listId/operations/richLongRunningOperationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdListIdOperationID(v.Input)
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

		if actual.RichLongRunningOperationId != v.Expected.RichLongRunningOperationId {
			t.Fatalf("Expected %q but got %q for RichLongRunningOperationId", v.Expected.RichLongRunningOperationId, actual.RichLongRunningOperationId)
		}

	}
}

func TestParseGroupIdSiteIdListIdOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdListIdOperationId
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
			Input: "/groups/groupId/sites/siteId/lists/listId/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/lIsTs/lIsTiD/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/lists/listId/operations/richLongRunningOperationId",
			Expected: &GroupIdSiteIdListIdOperationId{
				GroupId:                    "groupId",
				SiteId:                     "siteId",
				ListId:                     "listId",
				RichLongRunningOperationId: "richLongRunningOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/lists/listId/operations/richLongRunningOperationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/lIsTs/lIsTiD/oPeRaTiOnS/rIcHlOnGrUnNiNgOpErAtIoNiD",
			Expected: &GroupIdSiteIdListIdOperationId{
				GroupId:                    "gRoUpId",
				SiteId:                     "sItEiD",
				ListId:                     "lIsTiD",
				RichLongRunningOperationId: "rIcHlOnGrUnNiNgOpErAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/lIsTs/lIsTiD/oPeRaTiOnS/rIcHlOnGrUnNiNgOpErAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdListIdOperationIDInsensitively(v.Input)
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

		if actual.RichLongRunningOperationId != v.Expected.RichLongRunningOperationId {
			t.Fatalf("Expected %q but got %q for RichLongRunningOperationId", v.Expected.RichLongRunningOperationId, actual.RichLongRunningOperationId)
		}

	}
}

func TestSegmentsForGroupIdSiteIdListIdOperationId(t *testing.T) {
	segments := GroupIdSiteIdListIdOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdSiteIdListIdOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
