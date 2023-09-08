package groupsitelistcontenttypecolumnlink

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteListContentTypeColumnLinkId{}

func TestNewGroupSiteListContentTypeColumnLinkID(t *testing.T) {
	id := NewGroupSiteListContentTypeColumnLinkID("groupIdValue", "siteIdValue", "listIdValue", "contentTypeIdValue", "columnLinkIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.SiteId != "siteIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteIdValue")
	}

	if id.ListId != "listIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ListId'", id.ListId, "listIdValue")
	}

	if id.ContentTypeId != "contentTypeIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ContentTypeId'", id.ContentTypeId, "contentTypeIdValue")
	}

	if id.ColumnLinkId != "columnLinkIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ColumnLinkId'", id.ColumnLinkId, "columnLinkIdValue")
	}
}

func TestFormatGroupSiteListContentTypeColumnLinkID(t *testing.T) {
	actual := NewGroupSiteListContentTypeColumnLinkID("groupIdValue", "siteIdValue", "listIdValue", "contentTypeIdValue", "columnLinkIdValue").ID()
	expected := "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue/contentTypes/contentTypeIdValue/columnLinks/columnLinkIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupSiteListContentTypeColumnLinkID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteListContentTypeColumnLinkId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/lists",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue/contentTypes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue/contentTypes/contentTypeIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue/contentTypes/contentTypeIdValue/columnLinks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue/contentTypes/contentTypeIdValue/columnLinks/columnLinkIdValue",
			Expected: &GroupSiteListContentTypeColumnLinkId{
				GroupId:       "groupIdValue",
				SiteId:        "siteIdValue",
				ListId:        "listIdValue",
				ContentTypeId: "contentTypeIdValue",
				ColumnLinkId:  "columnLinkIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue/contentTypes/contentTypeIdValue/columnLinks/columnLinkIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteListContentTypeColumnLinkID(v.Input)
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

		if actual.ColumnLinkId != v.Expected.ColumnLinkId {
			t.Fatalf("Expected %q but got %q for ColumnLinkId", v.Expected.ColumnLinkId, actual.ColumnLinkId)
		}

	}
}

func TestParseGroupSiteListContentTypeColumnLinkIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteListContentTypeColumnLinkId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/lists",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/lIsTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/lIsTs/lIsTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue/contentTypes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/lIsTs/lIsTiDvAlUe/cOnTeNtTyPeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue/contentTypes/contentTypeIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/lIsTs/lIsTiDvAlUe/cOnTeNtTyPeS/cOnTeNtTyPeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue/contentTypes/contentTypeIdValue/columnLinks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/lIsTs/lIsTiDvAlUe/cOnTeNtTyPeS/cOnTeNtTyPeIdVaLuE/cOlUmNlInKs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue/contentTypes/contentTypeIdValue/columnLinks/columnLinkIdValue",
			Expected: &GroupSiteListContentTypeColumnLinkId{
				GroupId:       "groupIdValue",
				SiteId:        "siteIdValue",
				ListId:        "listIdValue",
				ContentTypeId: "contentTypeIdValue",
				ColumnLinkId:  "columnLinkIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue/contentTypes/contentTypeIdValue/columnLinks/columnLinkIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/lIsTs/lIsTiDvAlUe/cOnTeNtTyPeS/cOnTeNtTyPeIdVaLuE/cOlUmNlInKs/cOlUmNlInKiDvAlUe",
			Expected: &GroupSiteListContentTypeColumnLinkId{
				GroupId:       "gRoUpIdVaLuE",
				SiteId:        "sItEiDvAlUe",
				ListId:        "lIsTiDvAlUe",
				ContentTypeId: "cOnTeNtTyPeIdVaLuE",
				ColumnLinkId:  "cOlUmNlInKiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/lIsTs/lIsTiDvAlUe/cOnTeNtTyPeS/cOnTeNtTyPeIdVaLuE/cOlUmNlInKs/cOlUmNlInKiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteListContentTypeColumnLinkIDInsensitively(v.Input)
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

		if actual.ColumnLinkId != v.Expected.ColumnLinkId {
			t.Fatalf("Expected %q but got %q for ColumnLinkId", v.Expected.ColumnLinkId, actual.ColumnLinkId)
		}

	}
}

func TestSegmentsForGroupSiteListContentTypeColumnLinkId(t *testing.T) {
	segments := GroupSiteListContentTypeColumnLinkId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupSiteListContentTypeColumnLinkId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
