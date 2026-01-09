package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdListIdContentTypeIdColumnPositionId{}

func TestNewGroupIdSiteIdListIdContentTypeIdColumnPositionID(t *testing.T) {
	id := NewGroupIdSiteIdListIdContentTypeIdColumnPositionID("groupId", "siteId", "listId", "contentTypeId", "columnDefinitionId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.SiteId != "siteId" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteId")
	}

	if id.ListId != "listId" {
		t.Fatalf("Expected %q but got %q for Segment 'ListId'", id.ListId, "listId")
	}

	if id.ContentTypeId != "contentTypeId" {
		t.Fatalf("Expected %q but got %q for Segment 'ContentTypeId'", id.ContentTypeId, "contentTypeId")
	}

	if id.ColumnDefinitionId != "columnDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'ColumnDefinitionId'", id.ColumnDefinitionId, "columnDefinitionId")
	}
}

func TestFormatGroupIdSiteIdListIdContentTypeIdColumnPositionID(t *testing.T) {
	actual := NewGroupIdSiteIdListIdContentTypeIdColumnPositionID("groupId", "siteId", "listId", "contentTypeId", "columnDefinitionId").ID()
	expected := "/groups/groupId/sites/siteId/lists/listId/contentTypes/contentTypeId/columnPositions/columnDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdSiteIdListIdContentTypeIdColumnPositionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdListIdContentTypeIdColumnPositionId
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
			Input: "/groups/groupId/sites/siteId/lists/listId/contentTypes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/lists/listId/contentTypes/contentTypeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/lists/listId/contentTypes/contentTypeId/columnPositions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/lists/listId/contentTypes/contentTypeId/columnPositions/columnDefinitionId",
			Expected: &GroupIdSiteIdListIdContentTypeIdColumnPositionId{
				GroupId:            "groupId",
				SiteId:             "siteId",
				ListId:             "listId",
				ContentTypeId:      "contentTypeId",
				ColumnDefinitionId: "columnDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/lists/listId/contentTypes/contentTypeId/columnPositions/columnDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdListIdContentTypeIdColumnPositionID(v.Input)
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

		if actual.ContentTypeId != v.Expected.ContentTypeId {
			t.Fatalf("Expected %q but got %q for ContentTypeId", v.Expected.ContentTypeId, actual.ContentTypeId)
		}

		if actual.ColumnDefinitionId != v.Expected.ColumnDefinitionId {
			t.Fatalf("Expected %q but got %q for ColumnDefinitionId", v.Expected.ColumnDefinitionId, actual.ColumnDefinitionId)
		}

	}
}

func TestParseGroupIdSiteIdListIdContentTypeIdColumnPositionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdListIdContentTypeIdColumnPositionId
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
			Input: "/groups/groupId/sites/siteId/lists/listId/contentTypes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/lIsTs/lIsTiD/cOnTeNtTyPeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/lists/listId/contentTypes/contentTypeId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/lIsTs/lIsTiD/cOnTeNtTyPeS/cOnTeNtTyPeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/lists/listId/contentTypes/contentTypeId/columnPositions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/lIsTs/lIsTiD/cOnTeNtTyPeS/cOnTeNtTyPeId/cOlUmNpOsItIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/lists/listId/contentTypes/contentTypeId/columnPositions/columnDefinitionId",
			Expected: &GroupIdSiteIdListIdContentTypeIdColumnPositionId{
				GroupId:            "groupId",
				SiteId:             "siteId",
				ListId:             "listId",
				ContentTypeId:      "contentTypeId",
				ColumnDefinitionId: "columnDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/lists/listId/contentTypes/contentTypeId/columnPositions/columnDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/lIsTs/lIsTiD/cOnTeNtTyPeS/cOnTeNtTyPeId/cOlUmNpOsItIoNs/cOlUmNdEfInItIoNiD",
			Expected: &GroupIdSiteIdListIdContentTypeIdColumnPositionId{
				GroupId:            "gRoUpId",
				SiteId:             "sItEiD",
				ListId:             "lIsTiD",
				ContentTypeId:      "cOnTeNtTyPeId",
				ColumnDefinitionId: "cOlUmNdEfInItIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/lIsTs/lIsTiD/cOnTeNtTyPeS/cOnTeNtTyPeId/cOlUmNpOsItIoNs/cOlUmNdEfInItIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdListIdContentTypeIdColumnPositionIDInsensitively(v.Input)
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

		if actual.ContentTypeId != v.Expected.ContentTypeId {
			t.Fatalf("Expected %q but got %q for ContentTypeId", v.Expected.ContentTypeId, actual.ContentTypeId)
		}

		if actual.ColumnDefinitionId != v.Expected.ColumnDefinitionId {
			t.Fatalf("Expected %q but got %q for ColumnDefinitionId", v.Expected.ColumnDefinitionId, actual.ColumnDefinitionId)
		}

	}
}

func TestSegmentsForGroupIdSiteIdListIdContentTypeIdColumnPositionId(t *testing.T) {
	segments := GroupIdSiteIdListIdContentTypeIdColumnPositionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdSiteIdListIdContentTypeIdColumnPositionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
