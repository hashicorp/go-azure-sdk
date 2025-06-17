package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId{}

func TestNewGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionID(t *testing.T) {
	id := NewGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionID("groupId", "siteId", "baseSitePageId", "horizontalSectionId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.SiteId != "siteId" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteId")
	}

	if id.BaseSitePageId != "baseSitePageId" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseSitePageId'", id.BaseSitePageId, "baseSitePageId")
	}

	if id.HorizontalSectionId != "horizontalSectionId" {
		t.Fatalf("Expected %q but got %q for Segment 'HorizontalSectionId'", id.HorizontalSectionId, "horizontalSectionId")
	}
}

func TestFormatGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionID(t *testing.T) {
	actual := NewGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionID("groupId", "siteId", "baseSitePageId", "horizontalSectionId").ID()
	expected := "/groups/groupId/sites/siteId/pages/baseSitePageId/sitePage/canvasLayout/horizontalSections/horizontalSectionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId
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
			Input: "/groups/groupId/sites/siteId/pages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pages/baseSitePageId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pages/baseSitePageId/sitePage",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pages/baseSitePageId/sitePage/canvasLayout",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pages/baseSitePageId/sitePage/canvasLayout/horizontalSections",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/pages/baseSitePageId/sitePage/canvasLayout/horizontalSections/horizontalSectionId",
			Expected: &GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId{
				GroupId:             "groupId",
				SiteId:              "siteId",
				BaseSitePageId:      "baseSitePageId",
				HorizontalSectionId: "horizontalSectionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/pages/baseSitePageId/sitePage/canvasLayout/horizontalSections/horizontalSectionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionID(v.Input)
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

		if actual.BaseSitePageId != v.Expected.BaseSitePageId {
			t.Fatalf("Expected %q but got %q for BaseSitePageId", v.Expected.BaseSitePageId, actual.BaseSitePageId)
		}

		if actual.HorizontalSectionId != v.Expected.HorizontalSectionId {
			t.Fatalf("Expected %q but got %q for HorizontalSectionId", v.Expected.HorizontalSectionId, actual.HorizontalSectionId)
		}

	}
}

func TestParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId
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
			Input: "/groups/groupId/sites/siteId/pages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pages/baseSitePageId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEs/bAsEsItEpAgEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pages/baseSitePageId/sitePage",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEs/bAsEsItEpAgEiD/sItEpAgE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pages/baseSitePageId/sitePage/canvasLayout",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEs/bAsEsItEpAgEiD/sItEpAgE/cAnVaSlAyOuT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pages/baseSitePageId/sitePage/canvasLayout/horizontalSections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEs/bAsEsItEpAgEiD/sItEpAgE/cAnVaSlAyOuT/hOrIzOnTaLsEcTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/pages/baseSitePageId/sitePage/canvasLayout/horizontalSections/horizontalSectionId",
			Expected: &GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId{
				GroupId:             "groupId",
				SiteId:              "siteId",
				BaseSitePageId:      "baseSitePageId",
				HorizontalSectionId: "horizontalSectionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/pages/baseSitePageId/sitePage/canvasLayout/horizontalSections/horizontalSectionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEs/bAsEsItEpAgEiD/sItEpAgE/cAnVaSlAyOuT/hOrIzOnTaLsEcTiOnS/hOrIzOnTaLsEcTiOnId",
			Expected: &GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId{
				GroupId:             "gRoUpId",
				SiteId:              "sItEiD",
				BaseSitePageId:      "bAsEsItEpAgEiD",
				HorizontalSectionId: "hOrIzOnTaLsEcTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEs/bAsEsItEpAgEiD/sItEpAgE/cAnVaSlAyOuT/hOrIzOnTaLsEcTiOnS/hOrIzOnTaLsEcTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIDInsensitively(v.Input)
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

		if actual.BaseSitePageId != v.Expected.BaseSitePageId {
			t.Fatalf("Expected %q but got %q for BaseSitePageId", v.Expected.BaseSitePageId, actual.BaseSitePageId)
		}

		if actual.HorizontalSectionId != v.Expected.HorizontalSectionId {
			t.Fatalf("Expected %q but got %q for HorizontalSectionId", v.Expected.HorizontalSectionId, actual.HorizontalSectionId)
		}

	}
}

func TestSegmentsForGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId(t *testing.T) {
	segments := GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
