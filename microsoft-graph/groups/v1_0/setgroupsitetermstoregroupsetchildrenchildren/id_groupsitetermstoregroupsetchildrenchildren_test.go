package setgroupsitetermstoregroupsetchildrenchildren

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteTermStoreGroupSetChildrenChildrenId{}

func TestNewGroupSiteTermStoreGroupSetChildrenChildrenID(t *testing.T) {
	id := NewGroupSiteTermStoreGroupSetChildrenChildrenID("groupIdValue", "siteIdValue", "storeIdValue", "groupId1Value", "setIdValue", "termIdValue", "termId1Value")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.SiteId != "siteIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteIdValue")
	}

	if id.StoreId != "storeIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'StoreId'", id.StoreId, "storeIdValue")
	}

	if id.GroupId1 != "groupId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId1'", id.GroupId1, "groupId1Value")
	}

	if id.SetId != "setIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SetId'", id.SetId, "setIdValue")
	}

	if id.TermId != "termIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TermId'", id.TermId, "termIdValue")
	}

	if id.TermId1 != "termId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'TermId1'", id.TermId1, "termId1Value")
	}
}

func TestFormatGroupSiteTermStoreGroupSetChildrenChildrenID(t *testing.T) {
	actual := NewGroupSiteTermStoreGroupSetChildrenChildrenID("groupIdValue", "siteIdValue", "storeIdValue", "groupId1Value", "setIdValue", "termIdValue", "termId1Value").ID()
	expected := "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/groups/groupId1Value/sets/setIdValue/children/termIdValue/children/termId1Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupSiteTermStoreGroupSetChildrenChildrenID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteTermStoreGroupSetChildrenChildrenId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/groups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/groups/groupId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/groups/groupId1Value/sets",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/groups/groupId1Value/sets/setIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/groups/groupId1Value/sets/setIdValue/children",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/groups/groupId1Value/sets/setIdValue/children/termIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/groups/groupId1Value/sets/setIdValue/children/termIdValue/children",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/groups/groupId1Value/sets/setIdValue/children/termIdValue/children/termId1Value",
			Expected: &GroupSiteTermStoreGroupSetChildrenChildrenId{
				GroupId:  "groupIdValue",
				SiteId:   "siteIdValue",
				StoreId:  "storeIdValue",
				GroupId1: "groupId1Value",
				SetId:    "setIdValue",
				TermId:   "termIdValue",
				TermId1:  "termId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/groups/groupId1Value/sets/setIdValue/children/termIdValue/children/termId1Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteTermStoreGroupSetChildrenChildrenID(v.Input)
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

		if actual.GroupId1 != v.Expected.GroupId1 {
			t.Fatalf("Expected %q but got %q for GroupId1", v.Expected.GroupId1, actual.GroupId1)
		}

		if actual.SetId != v.Expected.SetId {
			t.Fatalf("Expected %q but got %q for SetId", v.Expected.SetId, actual.SetId)
		}

		if actual.TermId != v.Expected.TermId {
			t.Fatalf("Expected %q but got %q for TermId", v.Expected.TermId, actual.TermId)
		}

		if actual.TermId1 != v.Expected.TermId1 {
			t.Fatalf("Expected %q but got %q for TermId1", v.Expected.TermId1, actual.TermId1)
		}

	}
}

func TestParseGroupSiteTermStoreGroupSetChildrenChildrenIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteTermStoreGroupSetChildrenChildrenId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/groups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS/sToReIdVaLuE/gRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/groups/groupId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS/sToReIdVaLuE/gRoUpS/gRoUpId1vAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/groups/groupId1Value/sets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS/sToReIdVaLuE/gRoUpS/gRoUpId1vAlUe/sEtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/groups/groupId1Value/sets/setIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS/sToReIdVaLuE/gRoUpS/gRoUpId1vAlUe/sEtS/sEtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/groups/groupId1Value/sets/setIdValue/children",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS/sToReIdVaLuE/gRoUpS/gRoUpId1vAlUe/sEtS/sEtIdVaLuE/cHiLdReN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/groups/groupId1Value/sets/setIdValue/children/termIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS/sToReIdVaLuE/gRoUpS/gRoUpId1vAlUe/sEtS/sEtIdVaLuE/cHiLdReN/tErMiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/groups/groupId1Value/sets/setIdValue/children/termIdValue/children",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS/sToReIdVaLuE/gRoUpS/gRoUpId1vAlUe/sEtS/sEtIdVaLuE/cHiLdReN/tErMiDvAlUe/cHiLdReN",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/groups/groupId1Value/sets/setIdValue/children/termIdValue/children/termId1Value",
			Expected: &GroupSiteTermStoreGroupSetChildrenChildrenId{
				GroupId:  "groupIdValue",
				SiteId:   "siteIdValue",
				StoreId:  "storeIdValue",
				GroupId1: "groupId1Value",
				SetId:    "setIdValue",
				TermId:   "termIdValue",
				TermId1:  "termId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/termStores/storeIdValue/groups/groupId1Value/sets/setIdValue/children/termIdValue/children/termId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS/sToReIdVaLuE/gRoUpS/gRoUpId1vAlUe/sEtS/sEtIdVaLuE/cHiLdReN/tErMiDvAlUe/cHiLdReN/tErMiD1VaLuE",
			Expected: &GroupSiteTermStoreGroupSetChildrenChildrenId{
				GroupId:  "gRoUpIdVaLuE",
				SiteId:   "sItEiDvAlUe",
				StoreId:  "sToReIdVaLuE",
				GroupId1: "gRoUpId1vAlUe",
				SetId:    "sEtIdVaLuE",
				TermId:   "tErMiDvAlUe",
				TermId1:  "tErMiD1VaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToReS/sToReIdVaLuE/gRoUpS/gRoUpId1vAlUe/sEtS/sEtIdVaLuE/cHiLdReN/tErMiDvAlUe/cHiLdReN/tErMiD1VaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteTermStoreGroupSetChildrenChildrenIDInsensitively(v.Input)
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

		if actual.GroupId1 != v.Expected.GroupId1 {
			t.Fatalf("Expected %q but got %q for GroupId1", v.Expected.GroupId1, actual.GroupId1)
		}

		if actual.SetId != v.Expected.SetId {
			t.Fatalf("Expected %q but got %q for SetId", v.Expected.SetId, actual.SetId)
		}

		if actual.TermId != v.Expected.TermId {
			t.Fatalf("Expected %q but got %q for TermId", v.Expected.TermId, actual.TermId)
		}

		if actual.TermId1 != v.Expected.TermId1 {
			t.Fatalf("Expected %q but got %q for TermId1", v.Expected.TermId1, actual.TermId1)
		}

	}
}

func TestSegmentsForGroupSiteTermStoreGroupSetChildrenChildrenId(t *testing.T) {
	segments := GroupSiteTermStoreGroupSetChildrenChildrenId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupSiteTermStoreGroupSetChildrenChildrenId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
