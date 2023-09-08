package groupsitetermstoregroupsetrelation

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteTermStoreGroupSetId{}

func TestNewGroupSiteTermStoreGroupSetID(t *testing.T) {
	id := NewGroupSiteTermStoreGroupSetID("groupIdValue", "siteIdValue", "groupId1Value", "setIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.SiteId != "siteIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteIdValue")
	}

	if id.GroupId1 != "groupId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId1'", id.GroupId1, "groupId1Value")
	}

	if id.SetId != "setIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SetId'", id.SetId, "setIdValue")
	}
}

func TestFormatGroupSiteTermStoreGroupSetID(t *testing.T) {
	actual := NewGroupSiteTermStoreGroupSetID("groupIdValue", "siteIdValue", "groupId1Value", "setIdValue").ID()
	expected := "/groups/groupIdValue/sites/siteIdValue/termStore/groups/groupId1Value/sets/setIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupSiteTermStoreGroupSetID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteTermStoreGroupSetId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/groups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/groups/groupId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/groups/groupId1Value/sets",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/groups/groupId1Value/sets/setIdValue",
			Expected: &GroupSiteTermStoreGroupSetId{
				GroupId:  "groupIdValue",
				SiteId:   "siteIdValue",
				GroupId1: "groupId1Value",
				SetId:    "setIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/groups/groupId1Value/sets/setIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteTermStoreGroupSetID(v.Input)
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

		if actual.GroupId1 != v.Expected.GroupId1 {
			t.Fatalf("Expected %q but got %q for GroupId1", v.Expected.GroupId1, actual.GroupId1)
		}

		if actual.SetId != v.Expected.SetId {
			t.Fatalf("Expected %q but got %q for SetId", v.Expected.SetId, actual.SetId)
		}

	}
}

func TestParseGroupSiteTermStoreGroupSetIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteTermStoreGroupSetId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/groups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/gRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/groups/groupId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/gRoUpS/gRoUpId1vAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/groups/groupId1Value/sets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/gRoUpS/gRoUpId1vAlUe/sEtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/groups/groupId1Value/sets/setIdValue",
			Expected: &GroupSiteTermStoreGroupSetId{
				GroupId:  "groupIdValue",
				SiteId:   "siteIdValue",
				GroupId1: "groupId1Value",
				SetId:    "setIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/termStore/groups/groupId1Value/sets/setIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/gRoUpS/gRoUpId1vAlUe/sEtS/sEtIdVaLuE",
			Expected: &GroupSiteTermStoreGroupSetId{
				GroupId:  "gRoUpIdVaLuE",
				SiteId:   "sItEiDvAlUe",
				GroupId1: "gRoUpId1vAlUe",
				SetId:    "sEtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/tErMsToRe/gRoUpS/gRoUpId1vAlUe/sEtS/sEtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteTermStoreGroupSetIDInsensitively(v.Input)
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

		if actual.GroupId1 != v.Expected.GroupId1 {
			t.Fatalf("Expected %q but got %q for GroupId1", v.Expected.GroupId1, actual.GroupId1)
		}

		if actual.SetId != v.Expected.SetId {
			t.Fatalf("Expected %q but got %q for SetId", v.Expected.SetId, actual.SetId)
		}

	}
}

func TestSegmentsForGroupSiteTermStoreGroupSetId(t *testing.T) {
	segments := GroupSiteTermStoreGroupSetId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupSiteTermStoreGroupSetId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
