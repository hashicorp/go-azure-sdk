package grouponenotesectiongroupsectionpagecontent

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupOnenoteSectionGroupSectionPageId{}

func TestNewGroupOnenoteSectionGroupSectionPageID(t *testing.T) {
	id := NewGroupOnenoteSectionGroupSectionPageID("groupIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.SectionGroupId != "sectionGroupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SectionGroupId'", id.SectionGroupId, "sectionGroupIdValue")
	}

	if id.OnenoteSectionId != "onenoteSectionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenoteSectionId'", id.OnenoteSectionId, "onenoteSectionIdValue")
	}

	if id.OnenotePageId != "onenotePageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenotePageId'", id.OnenotePageId, "onenotePageIdValue")
	}
}

func TestFormatGroupOnenoteSectionGroupSectionPageID(t *testing.T) {
	actual := NewGroupOnenoteSectionGroupSectionPageID("groupIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue").ID()
	expected := "/groups/groupIdValue/onenote/sectionGroups/sectionGroupIdValue/sections/onenoteSectionIdValue/pages/onenotePageIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupOnenoteSectionGroupSectionPageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupOnenoteSectionGroupSectionPageId
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
			Input: "/groups/groupIdValue/onenote",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/onenote/sectionGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/onenote/sectionGroups/sectionGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/onenote/sectionGroups/sectionGroupIdValue/sections",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/onenote/sectionGroups/sectionGroupIdValue/sections/onenoteSectionIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/onenote/sectionGroups/sectionGroupIdValue/sections/onenoteSectionIdValue/pages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/onenote/sectionGroups/sectionGroupIdValue/sections/onenoteSectionIdValue/pages/onenotePageIdValue",
			Expected: &GroupOnenoteSectionGroupSectionPageId{
				GroupId:          "groupIdValue",
				SectionGroupId:   "sectionGroupIdValue",
				OnenoteSectionId: "onenoteSectionIdValue",
				OnenotePageId:    "onenotePageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/onenote/sectionGroups/sectionGroupIdValue/sections/onenoteSectionIdValue/pages/onenotePageIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupOnenoteSectionGroupSectionPageID(v.Input)
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

		if actual.SectionGroupId != v.Expected.SectionGroupId {
			t.Fatalf("Expected %q but got %q for SectionGroupId", v.Expected.SectionGroupId, actual.SectionGroupId)
		}

		if actual.OnenoteSectionId != v.Expected.OnenoteSectionId {
			t.Fatalf("Expected %q but got %q for OnenoteSectionId", v.Expected.OnenoteSectionId, actual.OnenoteSectionId)
		}

		if actual.OnenotePageId != v.Expected.OnenotePageId {
			t.Fatalf("Expected %q but got %q for OnenotePageId", v.Expected.OnenotePageId, actual.OnenotePageId)
		}

	}
}

func TestParseGroupOnenoteSectionGroupSectionPageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupOnenoteSectionGroupSectionPageId
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
			Input: "/groups/groupIdValue/onenote",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/oNeNoTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/onenote/sectionGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/oNeNoTe/sEcTiOnGrOuPs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/onenote/sectionGroups/sectionGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/onenote/sectionGroups/sectionGroupIdValue/sections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiDvAlUe/sEcTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/onenote/sectionGroups/sectionGroupIdValue/sections/onenoteSectionIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiDvAlUe/sEcTiOnS/oNeNoTeSeCtIoNiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/onenote/sectionGroups/sectionGroupIdValue/sections/onenoteSectionIdValue/pages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiDvAlUe/sEcTiOnS/oNeNoTeSeCtIoNiDvAlUe/pAgEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/onenote/sectionGroups/sectionGroupIdValue/sections/onenoteSectionIdValue/pages/onenotePageIdValue",
			Expected: &GroupOnenoteSectionGroupSectionPageId{
				GroupId:          "groupIdValue",
				SectionGroupId:   "sectionGroupIdValue",
				OnenoteSectionId: "onenoteSectionIdValue",
				OnenotePageId:    "onenotePageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/onenote/sectionGroups/sectionGroupIdValue/sections/onenoteSectionIdValue/pages/onenotePageIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiDvAlUe/sEcTiOnS/oNeNoTeSeCtIoNiDvAlUe/pAgEs/oNeNoTePaGeIdVaLuE",
			Expected: &GroupOnenoteSectionGroupSectionPageId{
				GroupId:          "gRoUpIdVaLuE",
				SectionGroupId:   "sEcTiOnGrOuPiDvAlUe",
				OnenoteSectionId: "oNeNoTeSeCtIoNiDvAlUe",
				OnenotePageId:    "oNeNoTePaGeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiDvAlUe/sEcTiOnS/oNeNoTeSeCtIoNiDvAlUe/pAgEs/oNeNoTePaGeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupOnenoteSectionGroupSectionPageIDInsensitively(v.Input)
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

		if actual.SectionGroupId != v.Expected.SectionGroupId {
			t.Fatalf("Expected %q but got %q for SectionGroupId", v.Expected.SectionGroupId, actual.SectionGroupId)
		}

		if actual.OnenoteSectionId != v.Expected.OnenoteSectionId {
			t.Fatalf("Expected %q but got %q for OnenoteSectionId", v.Expected.OnenoteSectionId, actual.OnenoteSectionId)
		}

		if actual.OnenotePageId != v.Expected.OnenotePageId {
			t.Fatalf("Expected %q but got %q for OnenotePageId", v.Expected.OnenotePageId, actual.OnenotePageId)
		}

	}
}

func TestSegmentsForGroupOnenoteSectionGroupSectionPageId(t *testing.T) {
	segments := GroupOnenoteSectionGroupSectionPageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupOnenoteSectionGroupSectionPageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
