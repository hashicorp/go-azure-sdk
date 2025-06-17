package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId{}

func TestNewGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartID(t *testing.T) {
	id := NewGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartID("groupId", "siteId", "pageTemplateId", "horizontalSectionId", "horizontalSectionColumnId", "webPartId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.SiteId != "siteId" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteId")
	}

	if id.PageTemplateId != "pageTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'PageTemplateId'", id.PageTemplateId, "pageTemplateId")
	}

	if id.HorizontalSectionId != "horizontalSectionId" {
		t.Fatalf("Expected %q but got %q for Segment 'HorizontalSectionId'", id.HorizontalSectionId, "horizontalSectionId")
	}

	if id.HorizontalSectionColumnId != "horizontalSectionColumnId" {
		t.Fatalf("Expected %q but got %q for Segment 'HorizontalSectionColumnId'", id.HorizontalSectionColumnId, "horizontalSectionColumnId")
	}

	if id.WebPartId != "webPartId" {
		t.Fatalf("Expected %q but got %q for Segment 'WebPartId'", id.WebPartId, "webPartId")
	}
}

func TestFormatGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartID(t *testing.T) {
	actual := NewGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartID("groupId", "siteId", "pageTemplateId", "horizontalSectionId", "horizontalSectionColumnId", "webPartId").ID()
	expected := "/groups/groupId/sites/siteId/pageTemplates/pageTemplateId/canvasLayout/horizontalSections/horizontalSectionId/columns/horizontalSectionColumnId/webparts/webPartId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId
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
			Input: "/groups/groupId/sites/siteId/pageTemplates",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pageTemplates/pageTemplateId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pageTemplates/pageTemplateId/canvasLayout",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pageTemplates/pageTemplateId/canvasLayout/horizontalSections",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pageTemplates/pageTemplateId/canvasLayout/horizontalSections/horizontalSectionId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pageTemplates/pageTemplateId/canvasLayout/horizontalSections/horizontalSectionId/columns",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pageTemplates/pageTemplateId/canvasLayout/horizontalSections/horizontalSectionId/columns/horizontalSectionColumnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pageTemplates/pageTemplateId/canvasLayout/horizontalSections/horizontalSectionId/columns/horizontalSectionColumnId/webparts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/pageTemplates/pageTemplateId/canvasLayout/horizontalSections/horizontalSectionId/columns/horizontalSectionColumnId/webparts/webPartId",
			Expected: &GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId{
				GroupId:                   "groupId",
				SiteId:                    "siteId",
				PageTemplateId:            "pageTemplateId",
				HorizontalSectionId:       "horizontalSectionId",
				HorizontalSectionColumnId: "horizontalSectionColumnId",
				WebPartId:                 "webPartId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/pageTemplates/pageTemplateId/canvasLayout/horizontalSections/horizontalSectionId/columns/horizontalSectionColumnId/webparts/webPartId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartID(v.Input)
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

		if actual.PageTemplateId != v.Expected.PageTemplateId {
			t.Fatalf("Expected %q but got %q for PageTemplateId", v.Expected.PageTemplateId, actual.PageTemplateId)
		}

		if actual.HorizontalSectionId != v.Expected.HorizontalSectionId {
			t.Fatalf("Expected %q but got %q for HorizontalSectionId", v.Expected.HorizontalSectionId, actual.HorizontalSectionId)
		}

		if actual.HorizontalSectionColumnId != v.Expected.HorizontalSectionColumnId {
			t.Fatalf("Expected %q but got %q for HorizontalSectionColumnId", v.Expected.HorizontalSectionColumnId, actual.HorizontalSectionColumnId)
		}

		if actual.WebPartId != v.Expected.WebPartId {
			t.Fatalf("Expected %q but got %q for WebPartId", v.Expected.WebPartId, actual.WebPartId)
		}

	}
}

func TestParseGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId
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
			Input: "/groups/groupId/sites/siteId/pageTemplates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEtEmPlAtEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pageTemplates/pageTemplateId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEtEmPlAtEs/pAgEtEmPlAtEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pageTemplates/pageTemplateId/canvasLayout",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEtEmPlAtEs/pAgEtEmPlAtEiD/cAnVaSlAyOuT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pageTemplates/pageTemplateId/canvasLayout/horizontalSections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEtEmPlAtEs/pAgEtEmPlAtEiD/cAnVaSlAyOuT/hOrIzOnTaLsEcTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pageTemplates/pageTemplateId/canvasLayout/horizontalSections/horizontalSectionId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEtEmPlAtEs/pAgEtEmPlAtEiD/cAnVaSlAyOuT/hOrIzOnTaLsEcTiOnS/hOrIzOnTaLsEcTiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pageTemplates/pageTemplateId/canvasLayout/horizontalSections/horizontalSectionId/columns",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEtEmPlAtEs/pAgEtEmPlAtEiD/cAnVaSlAyOuT/hOrIzOnTaLsEcTiOnS/hOrIzOnTaLsEcTiOnId/cOlUmNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pageTemplates/pageTemplateId/canvasLayout/horizontalSections/horizontalSectionId/columns/horizontalSectionColumnId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEtEmPlAtEs/pAgEtEmPlAtEiD/cAnVaSlAyOuT/hOrIzOnTaLsEcTiOnS/hOrIzOnTaLsEcTiOnId/cOlUmNs/hOrIzOnTaLsEcTiOnCoLuMnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/pageTemplates/pageTemplateId/canvasLayout/horizontalSections/horizontalSectionId/columns/horizontalSectionColumnId/webparts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEtEmPlAtEs/pAgEtEmPlAtEiD/cAnVaSlAyOuT/hOrIzOnTaLsEcTiOnS/hOrIzOnTaLsEcTiOnId/cOlUmNs/hOrIzOnTaLsEcTiOnCoLuMnId/wEbPaRtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/pageTemplates/pageTemplateId/canvasLayout/horizontalSections/horizontalSectionId/columns/horizontalSectionColumnId/webparts/webPartId",
			Expected: &GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId{
				GroupId:                   "groupId",
				SiteId:                    "siteId",
				PageTemplateId:            "pageTemplateId",
				HorizontalSectionId:       "horizontalSectionId",
				HorizontalSectionColumnId: "horizontalSectionColumnId",
				WebPartId:                 "webPartId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/pageTemplates/pageTemplateId/canvasLayout/horizontalSections/horizontalSectionId/columns/horizontalSectionColumnId/webparts/webPartId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEtEmPlAtEs/pAgEtEmPlAtEiD/cAnVaSlAyOuT/hOrIzOnTaLsEcTiOnS/hOrIzOnTaLsEcTiOnId/cOlUmNs/hOrIzOnTaLsEcTiOnCoLuMnId/wEbPaRtS/wEbPaRtId",
			Expected: &GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId{
				GroupId:                   "gRoUpId",
				SiteId:                    "sItEiD",
				PageTemplateId:            "pAgEtEmPlAtEiD",
				HorizontalSectionId:       "hOrIzOnTaLsEcTiOnId",
				HorizontalSectionColumnId: "hOrIzOnTaLsEcTiOnCoLuMnId",
				WebPartId:                 "wEbPaRtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/pAgEtEmPlAtEs/pAgEtEmPlAtEiD/cAnVaSlAyOuT/hOrIzOnTaLsEcTiOnS/hOrIzOnTaLsEcTiOnId/cOlUmNs/hOrIzOnTaLsEcTiOnCoLuMnId/wEbPaRtS/wEbPaRtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartIDInsensitively(v.Input)
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

		if actual.PageTemplateId != v.Expected.PageTemplateId {
			t.Fatalf("Expected %q but got %q for PageTemplateId", v.Expected.PageTemplateId, actual.PageTemplateId)
		}

		if actual.HorizontalSectionId != v.Expected.HorizontalSectionId {
			t.Fatalf("Expected %q but got %q for HorizontalSectionId", v.Expected.HorizontalSectionId, actual.HorizontalSectionId)
		}

		if actual.HorizontalSectionColumnId != v.Expected.HorizontalSectionColumnId {
			t.Fatalf("Expected %q but got %q for HorizontalSectionColumnId", v.Expected.HorizontalSectionColumnId, actual.HorizontalSectionColumnId)
		}

		if actual.WebPartId != v.Expected.WebPartId {
			t.Fatalf("Expected %q but got %q for WebPartId", v.Expected.WebPartId, actual.WebPartId)
		}

	}
}

func TestSegmentsForGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId(t *testing.T) {
	segments := GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
