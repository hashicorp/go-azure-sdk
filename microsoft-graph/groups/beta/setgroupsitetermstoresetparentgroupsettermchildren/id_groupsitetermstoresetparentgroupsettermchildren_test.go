package setgroupsitetermstoresetparentgroupsettermchildren

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteTermStoreSetParentGroupSetTermChildrenId{}

func TestNewGroupSiteTermStoreSetParentGroupSetTermChildrenID(t *testing.T) {
	id := NewGroupSiteTermStoreSetParentGroupSetTermChildrenID("groupIdValue", "siteIdValue", "setIdValue", "setId1Value", "termIdValue", "termId1Value")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.SiteId != "siteIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteIdValue")
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

	if id.TermId1 != "termId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'TermId1'", id.TermId1, "termId1Value")
	}
}

func TestFormatGroupSiteTermStoreSetParentGroupSetTermChildrenID(t *testing.T) {
	actual := NewGroupSiteTermStoreSetParentGroupSetTermChildrenID("groupIdValue", "siteIdValue", "setIdValue", "setId1Value", "termIdValue", "termId1Value").ID()
	expected := "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/parentGroup/sets/setId1Value/terms/termIdValue/children/termId1Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupSiteTermStoreSetParentGroupSetTermChildrenID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteTermStoreSetParentGroupSetTermChildrenId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/parentGroup",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/parentGroup/sets",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/parentGroup/sets/setId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/parentGroup/sets/setId1Value/terms",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/parentGroup/sets/setId1Value/terms/termIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/parentGroup/sets/setId1Value/terms/termIdValue/children",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/parentGroup/sets/setId1Value/terms/termIdValue/children/termId1Value",
			Expected: &GroupSiteTermStoreSetParentGroupSetTermChildrenId{
				GroupId: "groupIdValue",
				SiteId:  "siteIdValue",
				SetId:   "setIdValue",
				SetId1:  "setId1Value",
				TermId:  "termIdValue",
				TermId1: "termId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/parentGroup/sets/setId1Value/terms/termIdValue/children/termId1Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteTermStoreSetParentGroupSetTermChildrenID(v.Input)
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

		if actual.SetId != v.Expected.SetId {
			t.Fatalf("Expected %q but got %q for SetId", v.Expected.SetId, actual.SetId)
		}

		if actual.SetId1 != v.Expected.SetId1 {
			t.Fatalf("Expected %q but got %q for SetId1", v.Expected.SetId1, actual.SetId1)
		}

		if actual.TermId != v.Expected.TermId {
			t.Fatalf("Expected %q but got %q for TermId", v.Expected.TermId, actual.TermId)
		}

		if actual.TermId1 != v.Expected.TermId1 {
			t.Fatalf("Expected %q but got %q for TermId1", v.Expected.TermId1, actual.TermId1)
		}

	}
}

func TestParseGroupSiteTermStoreSetParentGroupSetTermChildrenIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteTermStoreSetParentGroupSetTermChildrenId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/sEtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/sEtS/sEtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/parentGroup",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/sEtS/sEtIdVaLuE/pArEnTgRoUp",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/parentGroup/sets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/sEtS/sEtIdVaLuE/pArEnTgRoUp/sEtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/parentGroup/sets/setId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/sEtS/sEtIdVaLuE/pArEnTgRoUp/sEtS/sEtId1vAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/parentGroup/sets/setId1Value/terms",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/sEtS/sEtIdVaLuE/pArEnTgRoUp/sEtS/sEtId1vAlUe/tErMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/parentGroup/sets/setId1Value/terms/termIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/sEtS/sEtIdVaLuE/pArEnTgRoUp/sEtS/sEtId1vAlUe/tErMs/tErMiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/parentGroup/sets/setId1Value/terms/termIdValue/children",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/sEtS/sEtIdVaLuE/pArEnTgRoUp/sEtS/sEtId1vAlUe/tErMs/tErMiDvAlUe/cHiLdReN",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/parentGroup/sets/setId1Value/terms/termIdValue/children/termId1Value",
			Expected: &GroupSiteTermStoreSetParentGroupSetTermChildrenId{
				GroupId: "groupIdValue",
				SiteId:  "siteIdValue",
				SetId:   "setIdValue",
				SetId1:  "setId1Value",
				TermId:  "termIdValue",
				TermId1: "termId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/parentGroup/sets/setId1Value/terms/termIdValue/children/termId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/sEtS/sEtIdVaLuE/pArEnTgRoUp/sEtS/sEtId1vAlUe/tErMs/tErMiDvAlUe/cHiLdReN/tErMiD1VaLuE",
			Expected: &GroupSiteTermStoreSetParentGroupSetTermChildrenId{
				GroupId: "gRoUpIdVaLuE",
				SiteId:  "sItEiDvAlUe",
				SetId:   "sEtIdVaLuE",
				SetId1:  "sEtId1vAlUe",
				TermId:  "tErMiDvAlUe",
				TermId1: "tErMiD1VaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/sEtS/sEtIdVaLuE/pArEnTgRoUp/sEtS/sEtId1vAlUe/tErMs/tErMiDvAlUe/cHiLdReN/tErMiD1VaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteTermStoreSetParentGroupSetTermChildrenIDInsensitively(v.Input)
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

		if actual.SetId != v.Expected.SetId {
			t.Fatalf("Expected %q but got %q for SetId", v.Expected.SetId, actual.SetId)
		}

		if actual.SetId1 != v.Expected.SetId1 {
			t.Fatalf("Expected %q but got %q for SetId1", v.Expected.SetId1, actual.SetId1)
		}

		if actual.TermId != v.Expected.TermId {
			t.Fatalf("Expected %q but got %q for TermId", v.Expected.TermId, actual.TermId)
		}

		if actual.TermId1 != v.Expected.TermId1 {
			t.Fatalf("Expected %q but got %q for TermId1", v.Expected.TermId1, actual.TermId1)
		}

	}
}

func TestSegmentsForGroupSiteTermStoreSetParentGroupSetTermChildrenId(t *testing.T) {
	segments := GroupSiteTermStoreSetParentGroupSetTermChildrenId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupSiteTermStoreSetParentGroupSetTermChildrenId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
