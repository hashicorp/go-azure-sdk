package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId{}

func TestNewGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnID(t *testing.T) {
	id := NewGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnID("groupId", "siteId", "baseSitePageId", "horizontalSectionId", "horizontalSectionColumnId")

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

	if id.HorizontalSectionColumnId != "horizontalSectionColumnId" {
		t.Fatalf("Expected %q but got %q for Segment 'HorizontalSectionColumnId'", id.HorizontalSectionColumnId, "horizontalSectionColumnId")
	}
}

func TestFormatGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnID(t *testing.T) {
	actual := NewGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnID("groupId", "siteId", "baseSitePageId", "horizontalSectionId", "horizontalSectionColumnId").ID()
	expected := "/groups/groupId/sites/siteId/pages/baseSitePageId/sitePage/canvasLayout/horizontalSections/horizontalSectionId/columns/horizontalSectionColumnId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId
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
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pages/baseSitePageId/sitePage/canvasLayout/horizontalSections/horizontalSectionId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pages/baseSitePageId/sitePage/canvasLayout/horizontalSections/horizontalSectionId/columns",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/pages/baseSitePageId/sitePage/canvasLayout/horizontalSections/horizontalSectionId/columns/horizontalSectionColumnId",
			Expected: &GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId{
				GroupId:                   "groupId",
				SiteId:                    "siteId",
				BaseSitePageId:            "baseSitePageId",
				HorizontalSectionId:       "horizontalSectionId",
				HorizontalSectionColumnId: "horizontalSectionColumnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/pages/baseSitePageId/sitePage/canvasLayout/horizontalSections/horizontalSectionId/columns/horizontalSectionColumnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnID(v.Input)
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

		if actual.HorizontalSectionColumnId != v.Expected.HorizontalSectionColumnId {
			t.Fatalf("Expected %q but got %q for HorizontalSectionColumnId", v.Expected.HorizontalSectionColumnId, actual.HorizontalSectionColumnId)
		}

	}
}

func TestParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId
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
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pages/baseSitePageId/sitePage/canvasLayout/horizontalSections/horizontalSectionId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEs/bAsEsItEpAgEiD/sItEpAgE/cAnVaSlAyOuT/hOrIzOnTaLsEcTiOnS/hOrIzOnTaLsEcTiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pages/baseSitePageId/sitePage/canvasLayout/horizontalSections/horizontalSectionId/columns",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEs/bAsEsItEpAgEiD/sItEpAgE/cAnVaSlAyOuT/hOrIzOnTaLsEcTiOnS/hOrIzOnTaLsEcTiOnId/cOlUmNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/pages/baseSitePageId/sitePage/canvasLayout/horizontalSections/horizontalSectionId/columns/horizontalSectionColumnId",
			Expected: &GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId{
				GroupId:                   "groupId",
				SiteId:                    "siteId",
				BaseSitePageId:            "baseSitePageId",
				HorizontalSectionId:       "horizontalSectionId",
				HorizontalSectionColumnId: "horizontalSectionColumnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/pages/baseSitePageId/sitePage/canvasLayout/horizontalSections/horizontalSectionId/columns/horizontalSectionColumnId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEs/bAsEsItEpAgEiD/sItEpAgE/cAnVaSlAyOuT/hOrIzOnTaLsEcTiOnS/hOrIzOnTaLsEcTiOnId/cOlUmNs/hOrIzOnTaLsEcTiOnCoLuMnId",
			Expected: &GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId{
				GroupId:                   "gRoUpId",
				SiteId:                    "sItEiD",
				BaseSitePageId:            "bAsEsItEpAgEiD",
				HorizontalSectionId:       "hOrIzOnTaLsEcTiOnId",
				HorizontalSectionColumnId: "hOrIzOnTaLsEcTiOnCoLuMnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEs/bAsEsItEpAgEiD/sItEpAgE/cAnVaSlAyOuT/hOrIzOnTaLsEcTiOnS/hOrIzOnTaLsEcTiOnId/cOlUmNs/hOrIzOnTaLsEcTiOnCoLuMnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIDInsensitively(v.Input)
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

		if actual.HorizontalSectionColumnId != v.Expected.HorizontalSectionColumnId {
			t.Fatalf("Expected %q but got %q for HorizontalSectionColumnId", v.Expected.HorizontalSectionColumnId, actual.HorizontalSectionColumnId)
		}

	}
}

func TestSegmentsForGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId(t *testing.T) {
	segments := GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
