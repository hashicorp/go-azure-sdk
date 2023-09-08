package useronenotesectiongroupsectionpage

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteSectionGroupSectionId{}

func TestNewUserOnenoteSectionGroupSectionID(t *testing.T) {
	id := NewUserOnenoteSectionGroupSectionID("userIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.SectionGroupId != "sectionGroupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SectionGroupId'", id.SectionGroupId, "sectionGroupIdValue")
	}

	if id.OnenoteSectionId != "onenoteSectionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenoteSectionId'", id.OnenoteSectionId, "onenoteSectionIdValue")
	}
}

func TestFormatUserOnenoteSectionGroupSectionID(t *testing.T) {
	actual := NewUserOnenoteSectionGroupSectionID("userIdValue", "sectionGroupIdValue", "onenoteSectionIdValue").ID()
	expected := "/users/userIdValue/onenote/sectionGroups/sectionGroupIdValue/sections/onenoteSectionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserOnenoteSectionGroupSectionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOnenoteSectionGroupSectionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote/sectionGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote/sectionGroups/sectionGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote/sectionGroups/sectionGroupIdValue/sections",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onenote/sectionGroups/sectionGroupIdValue/sections/onenoteSectionIdValue",
			Expected: &UserOnenoteSectionGroupSectionId{
				UserId:           "userIdValue",
				SectionGroupId:   "sectionGroupIdValue",
				OnenoteSectionId: "onenoteSectionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onenote/sectionGroups/sectionGroupIdValue/sections/onenoteSectionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOnenoteSectionGroupSectionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.SectionGroupId != v.Expected.SectionGroupId {
			t.Fatalf("Expected %q but got %q for SectionGroupId", v.Expected.SectionGroupId, actual.SectionGroupId)
		}

		if actual.OnenoteSectionId != v.Expected.OnenoteSectionId {
			t.Fatalf("Expected %q but got %q for OnenoteSectionId", v.Expected.OnenoteSectionId, actual.OnenoteSectionId)
		}

	}
}

func TestParseUserOnenoteSectionGroupSectionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOnenoteSectionGroupSectionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote/sectionGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/sEcTiOnGrOuPs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote/sectionGroups/sectionGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onenote/sectionGroups/sectionGroupIdValue/sections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiDvAlUe/sEcTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onenote/sectionGroups/sectionGroupIdValue/sections/onenoteSectionIdValue",
			Expected: &UserOnenoteSectionGroupSectionId{
				UserId:           "userIdValue",
				SectionGroupId:   "sectionGroupIdValue",
				OnenoteSectionId: "onenoteSectionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onenote/sectionGroups/sectionGroupIdValue/sections/onenoteSectionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiDvAlUe/sEcTiOnS/oNeNoTeSeCtIoNiDvAlUe",
			Expected: &UserOnenoteSectionGroupSectionId{
				UserId:           "uSeRiDvAlUe",
				SectionGroupId:   "sEcTiOnGrOuPiDvAlUe",
				OnenoteSectionId: "oNeNoTeSeCtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiDvAlUe/sEcTiOnS/oNeNoTeSeCtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOnenoteSectionGroupSectionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.SectionGroupId != v.Expected.SectionGroupId {
			t.Fatalf("Expected %q but got %q for SectionGroupId", v.Expected.SectionGroupId, actual.SectionGroupId)
		}

		if actual.OnenoteSectionId != v.Expected.OnenoteSectionId {
			t.Fatalf("Expected %q but got %q for OnenoteSectionId", v.Expected.OnenoteSectionId, actual.OnenoteSectionId)
		}

	}
}

func TestSegmentsForUserOnenoteSectionGroupSectionId(t *testing.T) {
	segments := UserOnenoteSectionGroupSectionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserOnenoteSectionGroupSectionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
