package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenoteSectionGroupId{}

func TestNewMeOnenoteSectionGroupID(t *testing.T) {
	id := NewMeOnenoteSectionGroupID("sectionGroupId")

	if id.SectionGroupId != "sectionGroupId" {
		t.Fatalf("Expected %q but got %q for Segment 'SectionGroupId'", id.SectionGroupId, "sectionGroupId")
	}
}

func TestFormatMeOnenoteSectionGroupID(t *testing.T) {
	actual := NewMeOnenoteSectionGroupID("sectionGroupId").ID()
	expected := "/me/onenote/sectionGroups/sectionGroupId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnenoteSectionGroupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnenoteSectionGroupId
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
			// Valid URI
			Input: "/me/onenote/sectionGroups/sectionGroupId",
			Expected: &MeOnenoteSectionGroupId{
				SectionGroupId: "sectionGroupId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onenote/sectionGroups/sectionGroupId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnenoteSectionGroupID(v.Input)
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

	}
}

func TestParseMeOnenoteSectionGroupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnenoteSectionGroupId
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
			// Valid URI
			Input: "/me/onenote/sectionGroups/sectionGroupId",
			Expected: &MeOnenoteSectionGroupId{
				SectionGroupId: "sectionGroupId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onenote/sectionGroups/sectionGroupId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiD",
			Expected: &MeOnenoteSectionGroupId{
				SectionGroupId: "sEcTiOnGrOuPiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnenoteSectionGroupIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForMeOnenoteSectionGroupId(t *testing.T) {
	segments := MeOnenoteSectionGroupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnenoteSectionGroupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
