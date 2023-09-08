package groupsitecontenttypecolumnposition

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteContentTypeColumnPositionId{}

func TestNewGroupSiteContentTypeColumnPositionID(t *testing.T) {
	id := NewGroupSiteContentTypeColumnPositionID("groupIdValue", "siteIdValue", "contentTypeIdValue", "columnDefinitionIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.SiteId != "siteIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteIdValue")
	}

	if id.ContentTypeId != "contentTypeIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ContentTypeId'", id.ContentTypeId, "contentTypeIdValue")
	}

	if id.ColumnDefinitionId != "columnDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ColumnDefinitionId'", id.ColumnDefinitionId, "columnDefinitionIdValue")
	}
}

func TestFormatGroupSiteContentTypeColumnPositionID(t *testing.T) {
	actual := NewGroupSiteContentTypeColumnPositionID("groupIdValue", "siteIdValue", "contentTypeIdValue", "columnDefinitionIdValue").ID()
	expected := "/groups/groupIdValue/sites/siteIdValue/contentTypes/contentTypeIdValue/columnPositions/columnDefinitionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupSiteContentTypeColumnPositionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteContentTypeColumnPositionId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/contentTypes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/contentTypes/contentTypeIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/contentTypes/contentTypeIdValue/columnPositions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/contentTypes/contentTypeIdValue/columnPositions/columnDefinitionIdValue",
			Expected: &GroupSiteContentTypeColumnPositionId{
				GroupId:            "groupIdValue",
				SiteId:             "siteIdValue",
				ContentTypeId:      "contentTypeIdValue",
				ColumnDefinitionId: "columnDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/contentTypes/contentTypeIdValue/columnPositions/columnDefinitionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteContentTypeColumnPositionID(v.Input)
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

		if actual.ContentTypeId != v.Expected.ContentTypeId {
			t.Fatalf("Expected %q but got %q for ContentTypeId", v.Expected.ContentTypeId, actual.ContentTypeId)
		}

		if actual.ColumnDefinitionId != v.Expected.ColumnDefinitionId {
			t.Fatalf("Expected %q but got %q for ColumnDefinitionId", v.Expected.ColumnDefinitionId, actual.ColumnDefinitionId)
		}

	}
}

func TestParseGroupSiteContentTypeColumnPositionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteContentTypeColumnPositionId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/contentTypes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/cOnTeNtTyPeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/contentTypes/contentTypeIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/cOnTeNtTyPeS/cOnTeNtTyPeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/contentTypes/contentTypeIdValue/columnPositions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/cOnTeNtTyPeS/cOnTeNtTyPeIdVaLuE/cOlUmNpOsItIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/contentTypes/contentTypeIdValue/columnPositions/columnDefinitionIdValue",
			Expected: &GroupSiteContentTypeColumnPositionId{
				GroupId:            "groupIdValue",
				SiteId:             "siteIdValue",
				ContentTypeId:      "contentTypeIdValue",
				ColumnDefinitionId: "columnDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/contentTypes/contentTypeIdValue/columnPositions/columnDefinitionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/cOnTeNtTyPeS/cOnTeNtTyPeIdVaLuE/cOlUmNpOsItIoNs/cOlUmNdEfInItIoNiDvAlUe",
			Expected: &GroupSiteContentTypeColumnPositionId{
				GroupId:            "gRoUpIdVaLuE",
				SiteId:             "sItEiDvAlUe",
				ContentTypeId:      "cOnTeNtTyPeIdVaLuE",
				ColumnDefinitionId: "cOlUmNdEfInItIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/cOnTeNtTyPeS/cOnTeNtTyPeIdVaLuE/cOlUmNpOsItIoNs/cOlUmNdEfInItIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteContentTypeColumnPositionIDInsensitively(v.Input)
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

		if actual.ContentTypeId != v.Expected.ContentTypeId {
			t.Fatalf("Expected %q but got %q for ContentTypeId", v.Expected.ContentTypeId, actual.ContentTypeId)
		}

		if actual.ColumnDefinitionId != v.Expected.ColumnDefinitionId {
			t.Fatalf("Expected %q but got %q for ColumnDefinitionId", v.Expected.ColumnDefinitionId, actual.ColumnDefinitionId)
		}

	}
}

func TestSegmentsForGroupSiteContentTypeColumnPositionId(t *testing.T) {
	segments := GroupSiteContentTypeColumnPositionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupSiteContentTypeColumnPositionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
