package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeFollowedSiteId{}

func TestNewMeFollowedSiteID(t *testing.T) {
	id := NewMeFollowedSiteID("siteIdValue")

	if id.SiteId != "siteIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteIdValue")
	}
}

func TestFormatMeFollowedSiteID(t *testing.T) {
	actual := NewMeFollowedSiteID("siteIdValue").ID()
	expected := "/me/followedSites/siteIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeFollowedSiteID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeFollowedSiteId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/followedSites",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/followedSites/siteIdValue",
			Expected: &MeFollowedSiteId{
				SiteId: "siteIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/followedSites/siteIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeFollowedSiteID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SiteId != v.Expected.SiteId {
			t.Fatalf("Expected %q but got %q for SiteId", v.Expected.SiteId, actual.SiteId)
		}

	}
}

func TestParseMeFollowedSiteIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeFollowedSiteId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/followedSites",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/fOlLoWeDsItEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/followedSites/siteIdValue",
			Expected: &MeFollowedSiteId{
				SiteId: "siteIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/followedSites/siteIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/fOlLoWeDsItEs/sItEiDvAlUe",
			Expected: &MeFollowedSiteId{
				SiteId: "sItEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/fOlLoWeDsItEs/sItEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeFollowedSiteIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SiteId != v.Expected.SiteId {
			t.Fatalf("Expected %q but got %q for SiteId", v.Expected.SiteId, actual.SiteId)
		}

	}
}

func TestSegmentsForMeFollowedSiteId(t *testing.T) {
	segments := MeFollowedSiteId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeFollowedSiteId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
