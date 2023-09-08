package groupsitetermstoresetparentgroupsettermrelation

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteTermStoreSetParentGroupSetTermRelationId{}

func TestNewGroupSiteTermStoreSetParentGroupSetTermRelationID(t *testing.T) {
	id := NewGroupSiteTermStoreSetParentGroupSetTermRelationID("groupIdValue", "siteIdValue", "storeIdValue", "setIdValue", "setId1Value", "termIdValue", "relationIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.SiteId != "siteIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteIdValue")
	}

	if id.StoreId != "storeIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'StoreId'", id.StoreId, "storeIdValue")
	}

	if id.SetId != "setIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SetId'", id.SetId, "setIdValue")
	}

	if id.SetId1 != "setId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'SetId1'", id.SetId1, "setId1Value")
	}

	if id.TermId != "termIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TermId'", id.TermId, "termIdValue")
	}

	if id.RelationId != "relationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RelationId'", id.RelationId, "relationIdValue")
	}
}

func TestFormatGroupSiteTermStoreSetParentGroupSetTermRelationID(t *testing.T) {
	actual := NewGroupSiteTermStoreSetParentGroupSetTermRelationID("groupIdValue", "siteIdValue", "storeIdValue", "setIdValue", "setId1Value", "termIdValue", "relationIdValue").ID()
	expected := "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets/setIdValue/parentGroup/sets/setId1Value/terms/termIdValue/relations/relationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupSiteTermStoreSetParentGroupSetTermRelationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteTermStoreSetParentGroupSetTermRelationId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets/setIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets/setIdValue/parentGroup",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets/setIdValue/parentGroup/sets",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets/setIdValue/parentGroup/sets/setId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets/setIdValue/parentGroup/sets/setId1Value/terms",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets/setIdValue/parentGroup/sets/setId1Value/terms/termIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets/setIdValue/parentGroup/sets/setId1Value/terms/termIdValue/relations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets/setIdValue/parentGroup/sets/setId1Value/terms/termIdValue/relations/relationIdValue",
			Expected: &GroupSiteTermStoreSetParentGroupSetTermRelationId{
				GroupId:    "groupIdValue",
				SiteId:     "siteIdValue",
				StoreId:    "storeIdValue",
				SetId:      "setIdValue",
				SetId1:     "setId1Value",
				TermId:     "termIdValue",
				RelationId: "relationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets/setIdValue/parentGroup/sets/setId1Value/terms/termIdValue/relations/relationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteTermStoreSetParentGroupSetTermRelationID(v.Input)
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

		if actual.StoreId != v.Expected.StoreId {
			t.Fatalf("Expected %q but got %q for StoreId", v.Expected.StoreId, actual.StoreId)
		}

		if actual.SetId != v.Expected.SetId {
			t.Fatalf("Expected %q but got %q for SetId", v.Expected.SetId, actual.SetId)
		}

		if actual.SetId1 != v.Expected.SetId1 {
			t.Fatalf("Expected %q but got %q for SetId1", v.Expected.SetId1, actual.SetId1)
		}

		if actual.TermId != v.Expected.TermId {
			t.Fatalf("Expected %q but got %q for TermId", v.Expected.TermId, actual.TermId)
		}

		if actual.RelationId != v.Expected.RelationId {
			t.Fatalf("Expected %q but got %q for RelationId", v.Expected.RelationId, actual.RelationId)
		}

	}
}

func TestParseGroupSiteTermStoreSetParentGroupSetTermRelationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteTermStoreSetParentGroupSetTermRelationId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS/sToReIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS/sToReIdVaLuE/sEtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets/setIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS/sToReIdVaLuE/sEtS/sEtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets/setIdValue/parentGroup",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS/sToReIdVaLuE/sEtS/sEtIdVaLuE/pArEnTgRoUp",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets/setIdValue/parentGroup/sets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS/sToReIdVaLuE/sEtS/sEtIdVaLuE/pArEnTgRoUp/sEtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets/setIdValue/parentGroup/sets/setId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS/sToReIdVaLuE/sEtS/sEtIdVaLuE/pArEnTgRoUp/sEtS/sEtId1vAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets/setIdValue/parentGroup/sets/setId1Value/terms",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS/sToReIdVaLuE/sEtS/sEtIdVaLuE/pArEnTgRoUp/sEtS/sEtId1vAlUe/tErMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets/setIdValue/parentGroup/sets/setId1Value/terms/termIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS/sToReIdVaLuE/sEtS/sEtIdVaLuE/pArEnTgRoUp/sEtS/sEtId1vAlUe/tErMs/tErMiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets/setIdValue/parentGroup/sets/setId1Value/terms/termIdValue/relations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS/sToReIdVaLuE/sEtS/sEtIdVaLuE/pArEnTgRoUp/sEtS/sEtId1vAlUe/tErMs/tErMiDvAlUe/rElAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets/setIdValue/parentGroup/sets/setId1Value/terms/termIdValue/relations/relationIdValue",
			Expected: &GroupSiteTermStoreSetParentGroupSetTermRelationId{
				GroupId:    "groupIdValue",
				SiteId:     "siteIdValue",
				StoreId:    "storeIdValue",
				SetId:      "setIdValue",
				SetId1:     "setId1Value",
				TermId:     "termIdValue",
				RelationId: "relationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/sets/setIdValue/parentGroup/sets/setId1Value/terms/termIdValue/relations/relationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS/sToReIdVaLuE/sEtS/sEtIdVaLuE/pArEnTgRoUp/sEtS/sEtId1vAlUe/tErMs/tErMiDvAlUe/rElAtIoNs/rElAtIoNiDvAlUe",
			Expected: &GroupSiteTermStoreSetParentGroupSetTermRelationId{
				GroupId:    "gRoUpIdVaLuE",
				SiteId:     "sItEiDvAlUe",
				StoreId:    "sToReIdVaLuE",
				SetId:      "sEtIdVaLuE",
				SetId1:     "sEtId1vAlUe",
				TermId:     "tErMiDvAlUe",
				RelationId: "rElAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS/sToReIdVaLuE/sEtS/sEtIdVaLuE/pArEnTgRoUp/sEtS/sEtId1vAlUe/tErMs/tErMiDvAlUe/rElAtIoNs/rElAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteTermStoreSetParentGroupSetTermRelationIDInsensitively(v.Input)
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

		if actual.StoreId != v.Expected.StoreId {
			t.Fatalf("Expected %q but got %q for StoreId", v.Expected.StoreId, actual.StoreId)
		}

		if actual.SetId != v.Expected.SetId {
			t.Fatalf("Expected %q but got %q for SetId", v.Expected.SetId, actual.SetId)
		}

		if actual.SetId1 != v.Expected.SetId1 {
			t.Fatalf("Expected %q but got %q for SetId1", v.Expected.SetId1, actual.SetId1)
		}

		if actual.TermId != v.Expected.TermId {
			t.Fatalf("Expected %q but got %q for TermId", v.Expected.TermId, actual.TermId)
		}

		if actual.RelationId != v.Expected.RelationId {
			t.Fatalf("Expected %q but got %q for RelationId", v.Expected.RelationId, actual.RelationId)
		}

	}
}

func TestSegmentsForGroupSiteTermStoreSetParentGroupSetTermRelationId(t *testing.T) {
	segments := GroupSiteTermStoreSetParentGroupSetTermRelationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupSiteTermStoreSetParentGroupSetTermRelationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
