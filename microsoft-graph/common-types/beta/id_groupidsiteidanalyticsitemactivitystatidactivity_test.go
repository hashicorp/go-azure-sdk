package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdAnalyticsItemActivityStatIdActivityId{}

func TestNewGroupIdSiteIdAnalyticsItemActivityStatIdActivityID(t *testing.T) {
	id := NewGroupIdSiteIdAnalyticsItemActivityStatIdActivityID("groupId", "siteId", "itemActivityStatId", "itemActivityId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.SiteId != "siteId" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteId")
	}

	if id.ItemActivityStatId != "itemActivityStatId" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemActivityStatId'", id.ItemActivityStatId, "itemActivityStatId")
	}

	if id.ItemActivityId != "itemActivityId" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemActivityId'", id.ItemActivityId, "itemActivityId")
	}
}

func TestFormatGroupIdSiteIdAnalyticsItemActivityStatIdActivityID(t *testing.T) {
	actual := NewGroupIdSiteIdAnalyticsItemActivityStatIdActivityID("groupId", "siteId", "itemActivityStatId", "itemActivityId").ID()
	expected := "/groups/groupId/sites/siteId/analytics/itemActivityStats/itemActivityStatId/activities/itemActivityId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdSiteIdAnalyticsItemActivityStatIdActivityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdAnalyticsItemActivityStatIdActivityId
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
			Input: "/groups/groupId/sites/siteId/analytics",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/analytics/itemActivityStats",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/analytics/itemActivityStats/itemActivityStatId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/analytics/itemActivityStats/itemActivityStatId/activities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/analytics/itemActivityStats/itemActivityStatId/activities/itemActivityId",
			Expected: &GroupIdSiteIdAnalyticsItemActivityStatIdActivityId{
				GroupId:            "groupId",
				SiteId:             "siteId",
				ItemActivityStatId: "itemActivityStatId",
				ItemActivityId:     "itemActivityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/analytics/itemActivityStats/itemActivityStatId/activities/itemActivityId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdAnalyticsItemActivityStatIdActivityID(v.Input)
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

		if actual.ItemActivityStatId != v.Expected.ItemActivityStatId {
			t.Fatalf("Expected %q but got %q for ItemActivityStatId", v.Expected.ItemActivityStatId, actual.ItemActivityStatId)
		}

		if actual.ItemActivityId != v.Expected.ItemActivityId {
			t.Fatalf("Expected %q but got %q for ItemActivityId", v.Expected.ItemActivityId, actual.ItemActivityId)
		}

	}
}

func TestParseGroupIdSiteIdAnalyticsItemActivityStatIdActivityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdAnalyticsItemActivityStatIdActivityId
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
			Input: "/groups/groupId/sites/siteId/analytics",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/aNaLyTiCs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/analytics/itemActivityStats",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/aNaLyTiCs/iTeMaCtIvItYsTaTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/analytics/itemActivityStats/itemActivityStatId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/analytics/itemActivityStats/itemActivityStatId/activities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiD/aCtIvItIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/analytics/itemActivityStats/itemActivityStatId/activities/itemActivityId",
			Expected: &GroupIdSiteIdAnalyticsItemActivityStatIdActivityId{
				GroupId:            "groupId",
				SiteId:             "siteId",
				ItemActivityStatId: "itemActivityStatId",
				ItemActivityId:     "itemActivityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/analytics/itemActivityStats/itemActivityStatId/activities/itemActivityId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiD/aCtIvItIeS/iTeMaCtIvItYiD",
			Expected: &GroupIdSiteIdAnalyticsItemActivityStatIdActivityId{
				GroupId:            "gRoUpId",
				SiteId:             "sItEiD",
				ItemActivityStatId: "iTeMaCtIvItYsTaTiD",
				ItemActivityId:     "iTeMaCtIvItYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiD/aCtIvItIeS/iTeMaCtIvItYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdAnalyticsItemActivityStatIdActivityIDInsensitively(v.Input)
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

		if actual.ItemActivityStatId != v.Expected.ItemActivityStatId {
			t.Fatalf("Expected %q but got %q for ItemActivityStatId", v.Expected.ItemActivityStatId, actual.ItemActivityStatId)
		}

		if actual.ItemActivityId != v.Expected.ItemActivityId {
			t.Fatalf("Expected %q but got %q for ItemActivityId", v.Expected.ItemActivityId, actual.ItemActivityId)
		}

	}
}

func TestSegmentsForGroupIdSiteIdAnalyticsItemActivityStatIdActivityId(t *testing.T) {
	segments := GroupIdSiteIdAnalyticsItemActivityStatIdActivityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdSiteIdAnalyticsItemActivityStatIdActivityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
