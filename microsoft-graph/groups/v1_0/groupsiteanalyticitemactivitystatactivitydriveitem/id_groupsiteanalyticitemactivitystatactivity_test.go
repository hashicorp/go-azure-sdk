package groupsiteanalyticitemactivitystatactivitydriveitem

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteAnalyticItemActivityStatActivityId{}

func TestNewGroupSiteAnalyticItemActivityStatActivityID(t *testing.T) {
	id := NewGroupSiteAnalyticItemActivityStatActivityID("groupIdValue", "siteIdValue", "itemActivityStatIdValue", "itemActivityIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.SiteId != "siteIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteIdValue")
	}

	if id.ItemActivityStatId != "itemActivityStatIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemActivityStatId'", id.ItemActivityStatId, "itemActivityStatIdValue")
	}

	if id.ItemActivityId != "itemActivityIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemActivityId'", id.ItemActivityId, "itemActivityIdValue")
	}
}

func TestFormatGroupSiteAnalyticItemActivityStatActivityID(t *testing.T) {
	actual := NewGroupSiteAnalyticItemActivityStatActivityID("groupIdValue", "siteIdValue", "itemActivityStatIdValue", "itemActivityIdValue").ID()
	expected := "/groups/groupIdValue/sites/siteIdValue/analytics/itemActivityStats/itemActivityStatIdValue/activities/itemActivityIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupSiteAnalyticItemActivityStatActivityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteAnalyticItemActivityStatActivityId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/analytics",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/analytics/itemActivityStats",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/analytics/itemActivityStats/itemActivityStatIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/analytics/itemActivityStats/itemActivityStatIdValue/activities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/analytics/itemActivityStats/itemActivityStatIdValue/activities/itemActivityIdValue",
			Expected: &GroupSiteAnalyticItemActivityStatActivityId{
				GroupId:            "groupIdValue",
				SiteId:             "siteIdValue",
				ItemActivityStatId: "itemActivityStatIdValue",
				ItemActivityId:     "itemActivityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/analytics/itemActivityStats/itemActivityStatIdValue/activities/itemActivityIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteAnalyticItemActivityStatActivityID(v.Input)
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

func TestParseGroupSiteAnalyticItemActivityStatActivityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteAnalyticItemActivityStatActivityId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/analytics",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/aNaLyTiCs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/analytics/itemActivityStats",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/aNaLyTiCs/iTeMaCtIvItYsTaTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/analytics/itemActivityStats/itemActivityStatIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/analytics/itemActivityStats/itemActivityStatIdValue/activities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiDvAlUe/aCtIvItIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/analytics/itemActivityStats/itemActivityStatIdValue/activities/itemActivityIdValue",
			Expected: &GroupSiteAnalyticItemActivityStatActivityId{
				GroupId:            "groupIdValue",
				SiteId:             "siteIdValue",
				ItemActivityStatId: "itemActivityStatIdValue",
				ItemActivityId:     "itemActivityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/analytics/itemActivityStats/itemActivityStatIdValue/activities/itemActivityIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiDvAlUe/aCtIvItIeS/iTeMaCtIvItYiDvAlUe",
			Expected: &GroupSiteAnalyticItemActivityStatActivityId{
				GroupId:            "gRoUpIdVaLuE",
				SiteId:             "sItEiDvAlUe",
				ItemActivityStatId: "iTeMaCtIvItYsTaTiDvAlUe",
				ItemActivityId:     "iTeMaCtIvItYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/aNaLyTiCs/iTeMaCtIvItYsTaTs/iTeMaCtIvItYsTaTiDvAlUe/aCtIvItIeS/iTeMaCtIvItYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteAnalyticItemActivityStatActivityIDInsensitively(v.Input)
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

func TestSegmentsForGroupSiteAnalyticItemActivityStatActivityId(t *testing.T) {
	segments := GroupSiteAnalyticItemActivityStatActivityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupSiteAnalyticItemActivityStatActivityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
