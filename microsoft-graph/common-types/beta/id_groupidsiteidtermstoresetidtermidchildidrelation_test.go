package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId{}

func TestNewGroupIdSiteIdTermStoreSetIdTermIdChildIdRelationID(t *testing.T) {
	id := NewGroupIdSiteIdTermStoreSetIdTermIdChildIdRelationID("groupIdValue", "siteIdValue", "setIdValue", "termIdValue", "termId1Value", "relationIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.SiteId != "siteIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteIdValue")
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

	if id.RelationId != "relationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RelationId'", id.RelationId, "relationIdValue")
	}
}

func TestFormatGroupIdSiteIdTermStoreSetIdTermIdChildIdRelationID(t *testing.T) {
	actual := NewGroupIdSiteIdTermStoreSetIdTermIdChildIdRelationID("groupIdValue", "siteIdValue", "setIdValue", "termIdValue", "termId1Value", "relationIdValue").ID()
	expected := "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/terms/termIdValue/children/termId1Value/relations/relationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdSiteIdTermStoreSetIdTermIdChildIdRelationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/terms",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/terms/termIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/terms/termIdValue/children",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/terms/termIdValue/children/termId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/terms/termIdValue/children/termId1Value/relations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/terms/termIdValue/children/termId1Value/relations/relationIdValue",
			Expected: &GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId{
				GroupId:    "groupIdValue",
				SiteId:     "siteIdValue",
				SetId:      "setIdValue",
				TermId:     "termIdValue",
				TermId1:    "termId1Value",
				RelationId: "relationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/terms/termIdValue/children/termId1Value/relations/relationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdTermStoreSetIdTermIdChildIdRelationID(v.Input)
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

		if actual.TermId != v.Expected.TermId {
			t.Fatalf("Expected %q but got %q for TermId", v.Expected.TermId, actual.TermId)
		}

		if actual.TermId1 != v.Expected.TermId1 {
			t.Fatalf("Expected %q but got %q for TermId1", v.Expected.TermId1, actual.TermId1)
		}

		if actual.RelationId != v.Expected.RelationId {
			t.Fatalf("Expected %q but got %q for RelationId", v.Expected.RelationId, actual.RelationId)
		}

	}
}

func TestParseGroupIdSiteIdTermStoreSetIdTermIdChildIdRelationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/terms",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/sEtS/sEtIdVaLuE/tErMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/terms/termIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/sEtS/sEtIdVaLuE/tErMs/tErMiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/terms/termIdValue/children",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/sEtS/sEtIdVaLuE/tErMs/tErMiDvAlUe/cHiLdReN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/terms/termIdValue/children/termId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/sEtS/sEtIdVaLuE/tErMs/tErMiDvAlUe/cHiLdReN/tErMiD1VaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/terms/termIdValue/children/termId1Value/relations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/sEtS/sEtIdVaLuE/tErMs/tErMiDvAlUe/cHiLdReN/tErMiD1VaLuE/rElAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/terms/termIdValue/children/termId1Value/relations/relationIdValue",
			Expected: &GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId{
				GroupId:    "groupIdValue",
				SiteId:     "siteIdValue",
				SetId:      "setIdValue",
				TermId:     "termIdValue",
				TermId1:    "termId1Value",
				RelationId: "relationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/sets/setIdValue/terms/termIdValue/children/termId1Value/relations/relationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/sEtS/sEtIdVaLuE/tErMs/tErMiDvAlUe/cHiLdReN/tErMiD1VaLuE/rElAtIoNs/rElAtIoNiDvAlUe",
			Expected: &GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId{
				GroupId:    "gRoUpIdVaLuE",
				SiteId:     "sItEiDvAlUe",
				SetId:      "sEtIdVaLuE",
				TermId:     "tErMiDvAlUe",
				TermId1:    "tErMiD1VaLuE",
				RelationId: "rElAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/sEtS/sEtIdVaLuE/tErMs/tErMiDvAlUe/cHiLdReN/tErMiD1VaLuE/rElAtIoNs/rElAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdTermStoreSetIdTermIdChildIdRelationIDInsensitively(v.Input)
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

		if actual.TermId != v.Expected.TermId {
			t.Fatalf("Expected %q but got %q for TermId", v.Expected.TermId, actual.TermId)
		}

		if actual.TermId1 != v.Expected.TermId1 {
			t.Fatalf("Expected %q but got %q for TermId1", v.Expected.TermId1, actual.TermId1)
		}

		if actual.RelationId != v.Expected.RelationId {
			t.Fatalf("Expected %q but got %q for RelationId", v.Expected.RelationId, actual.RelationId)
		}

	}
}

func TestSegmentsForGroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId(t *testing.T) {
	segments := GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdSiteIdTermStoreSetIdTermIdChildIdRelationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
