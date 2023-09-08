package groupsitelistoperation

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteListOperationId{}

func TestNewGroupSiteListOperationID(t *testing.T) {
	id := NewGroupSiteListOperationID("groupIdValue", "siteIdValue", "listIdValue", "richLongRunningOperationIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.SiteId != "siteIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteIdValue")
	}

	if id.ListId != "listIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ListId'", id.ListId, "listIdValue")
	}

	if id.RichLongRunningOperationId != "richLongRunningOperationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RichLongRunningOperationId'", id.RichLongRunningOperationId, "richLongRunningOperationIdValue")
	}
}

func TestFormatGroupSiteListOperationID(t *testing.T) {
	actual := NewGroupSiteListOperationID("groupIdValue", "siteIdValue", "listIdValue", "richLongRunningOperationIdValue").ID()
	expected := "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue/operations/richLongRunningOperationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupSiteListOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteListOperationId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/lists",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue/operations/richLongRunningOperationIdValue",
			Expected: &GroupSiteListOperationId{
				GroupId:                    "groupIdValue",
				SiteId:                     "siteIdValue",
				ListId:                     "listIdValue",
				RichLongRunningOperationId: "richLongRunningOperationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue/operations/richLongRunningOperationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteListOperationID(v.Input)
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

		if actual.ListId != v.Expected.ListId {
			t.Fatalf("Expected %q but got %q for ListId", v.Expected.ListId, actual.ListId)
		}

		if actual.RichLongRunningOperationId != v.Expected.RichLongRunningOperationId {
			t.Fatalf("Expected %q but got %q for RichLongRunningOperationId", v.Expected.RichLongRunningOperationId, actual.RichLongRunningOperationId)
		}

	}
}

func TestParseGroupSiteListOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteListOperationId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/lists",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/lIsTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/lIsTs/lIsTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/lIsTs/lIsTiDvAlUe/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue/operations/richLongRunningOperationIdValue",
			Expected: &GroupSiteListOperationId{
				GroupId:                    "groupIdValue",
				SiteId:                     "siteIdValue",
				ListId:                     "listIdValue",
				RichLongRunningOperationId: "richLongRunningOperationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/lists/listIdValue/operations/richLongRunningOperationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/lIsTs/lIsTiDvAlUe/oPeRaTiOnS/rIcHlOnGrUnNiNgOpErAtIoNiDvAlUe",
			Expected: &GroupSiteListOperationId{
				GroupId:                    "gRoUpIdVaLuE",
				SiteId:                     "sItEiDvAlUe",
				ListId:                     "lIsTiDvAlUe",
				RichLongRunningOperationId: "rIcHlOnGrUnNiNgOpErAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/lIsTs/lIsTiDvAlUe/oPeRaTiOnS/rIcHlOnGrUnNiNgOpErAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteListOperationIDInsensitively(v.Input)
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

		if actual.ListId != v.Expected.ListId {
			t.Fatalf("Expected %q but got %q for ListId", v.Expected.ListId, actual.ListId)
		}

		if actual.RichLongRunningOperationId != v.Expected.RichLongRunningOperationId {
			t.Fatalf("Expected %q but got %q for RichLongRunningOperationId", v.Expected.RichLongRunningOperationId, actual.RichLongRunningOperationId)
		}

	}
}

func TestSegmentsForGroupSiteListOperationId(t *testing.T) {
	segments := GroupSiteListOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupSiteListOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
