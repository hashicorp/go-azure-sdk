package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenoteSectionGroupIdSectionId{}

func TestNewMeOnenoteSectionGroupIdSectionID(t *testing.T) {
	id := NewMeOnenoteSectionGroupIdSectionID("sectionGroupId", "onenoteSectionId")

	if id.SectionGroupId != "sectionGroupId" {
		t.Fatalf("Expected %q but got %q for Segment 'SectionGroupId'", id.SectionGroupId, "sectionGroupId")
	}

	if id.OnenoteSectionId != "onenoteSectionId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenoteSectionId'", id.OnenoteSectionId, "onenoteSectionId")
	}
}

func TestFormatMeOnenoteSectionGroupIdSectionID(t *testing.T) {
	actual := NewMeOnenoteSectionGroupIdSectionID("sectionGroupId", "onenoteSectionId").ID()
	expected := "/me/onenote/sectionGroups/sectionGroupId/sections/onenoteSectionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnenoteSectionGroupIdSectionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnenoteSectionGroupIdSectionId
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
			Input: "/me/onenote",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/sectionGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/sectionGroups/sectionGroupId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/sectionGroups/sectionGroupId/sections",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onenote/sectionGroups/sectionGroupId/sections/onenoteSectionId",
			Expected: &MeOnenoteSectionGroupIdSectionId{
				SectionGroupId:   "sectionGroupId",
				OnenoteSectionId: "onenoteSectionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onenote/sectionGroups/sectionGroupId/sections/onenoteSectionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnenoteSectionGroupIdSectionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SectionGroupId != v.Expected.SectionGroupId {
			t.Fatalf("Expected %q but got %q for SectionGroupId", v.Expected.SectionGroupId, actual.SectionGroupId)
		}

		if actual.OnenoteSectionId != v.Expected.OnenoteSectionId {
			t.Fatalf("Expected %q but got %q for OnenoteSectionId", v.Expected.OnenoteSectionId, actual.OnenoteSectionId)
		}

	}
}

func TestParseMeOnenoteSectionGroupIdSectionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnenoteSectionGroupIdSectionId
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
			Input: "/me/onenote",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/sectionGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/sEcTiOnGrOuPs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/sectionGroups/sectionGroupId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/sectionGroups/sectionGroupId/sections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiD/sEcTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onenote/sectionGroups/sectionGroupId/sections/onenoteSectionId",
			Expected: &MeOnenoteSectionGroupIdSectionId{
				SectionGroupId:   "sectionGroupId",
				OnenoteSectionId: "onenoteSectionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onenote/sectionGroups/sectionGroupId/sections/onenoteSectionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiD/sEcTiOnS/oNeNoTeSeCtIoNiD",
			Expected: &MeOnenoteSectionGroupIdSectionId{
				SectionGroupId:   "sEcTiOnGrOuPiD",
				OnenoteSectionId: "oNeNoTeSeCtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiD/sEcTiOnS/oNeNoTeSeCtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnenoteSectionGroupIdSectionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SectionGroupId != v.Expected.SectionGroupId {
			t.Fatalf("Expected %q but got %q for SectionGroupId", v.Expected.SectionGroupId, actual.SectionGroupId)
		}

		if actual.OnenoteSectionId != v.Expected.OnenoteSectionId {
			t.Fatalf("Expected %q but got %q for OnenoteSectionId", v.Expected.OnenoteSectionId, actual.OnenoteSectionId)
		}

	}
}

func TestSegmentsForMeOnenoteSectionGroupIdSectionId(t *testing.T) {
	segments := MeOnenoteSectionGroupIdSectionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnenoteSectionGroupIdSectionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
