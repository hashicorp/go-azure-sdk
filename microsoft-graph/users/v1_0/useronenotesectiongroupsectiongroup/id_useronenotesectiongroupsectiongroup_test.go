package useronenotesectiongroupsectiongroup

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteSectionGroupSectionGroupId{}

func TestNewUserOnenoteSectionGroupSectionGroupID(t *testing.T) {
	id := NewUserOnenoteSectionGroupSectionGroupID("userIdValue", "sectionGroupIdValue", "sectionGroupId1Value")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.SectionGroupId != "sectionGroupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SectionGroupId'", id.SectionGroupId, "sectionGroupIdValue")
	}

	if id.SectionGroupId1 != "sectionGroupId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'SectionGroupId1'", id.SectionGroupId1, "sectionGroupId1Value")
	}
}

func TestFormatUserOnenoteSectionGroupSectionGroupID(t *testing.T) {
	actual := NewUserOnenoteSectionGroupSectionGroupID("userIdValue", "sectionGroupIdValue", "sectionGroupId1Value").ID()
	expected := "/users/userIdValue/onenote/sectionGroups/sectionGroupIdValue/sectionGroups/sectionGroupId1Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserOnenoteSectionGroupSectionGroupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOnenoteSectionGroupSectionGroupId
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
			Input: "/users/userIdValue/onenote/sectionGroups/sectionGroupIdValue/sectionGroups",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onenote/sectionGroups/sectionGroupIdValue/sectionGroups/sectionGroupId1Value",
			Expected: &UserOnenoteSectionGroupSectionGroupId{
				UserId:          "userIdValue",
				SectionGroupId:  "sectionGroupIdValue",
				SectionGroupId1: "sectionGroupId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onenote/sectionGroups/sectionGroupIdValue/sectionGroups/sectionGroupId1Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOnenoteSectionGroupSectionGroupID(v.Input)
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

		if actual.SectionGroupId1 != v.Expected.SectionGroupId1 {
			t.Fatalf("Expected %q but got %q for SectionGroupId1", v.Expected.SectionGroupId1, actual.SectionGroupId1)
		}

	}
}

func TestParseUserOnenoteSectionGroupSectionGroupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOnenoteSectionGroupSectionGroupId
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
			Input: "/users/userIdValue/onenote/sectionGroups/sectionGroupIdValue/sectionGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiDvAlUe/sEcTiOnGrOuPs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onenote/sectionGroups/sectionGroupIdValue/sectionGroups/sectionGroupId1Value",
			Expected: &UserOnenoteSectionGroupSectionGroupId{
				UserId:          "userIdValue",
				SectionGroupId:  "sectionGroupIdValue",
				SectionGroupId1: "sectionGroupId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onenote/sectionGroups/sectionGroupIdValue/sectionGroups/sectionGroupId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiDvAlUe/sEcTiOnGrOuPs/sEcTiOnGrOuPiD1VaLuE",
			Expected: &UserOnenoteSectionGroupSectionGroupId{
				UserId:          "uSeRiDvAlUe",
				SectionGroupId:  "sEcTiOnGrOuPiDvAlUe",
				SectionGroupId1: "sEcTiOnGrOuPiD1VaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiDvAlUe/sEcTiOnGrOuPs/sEcTiOnGrOuPiD1VaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOnenoteSectionGroupSectionGroupIDInsensitively(v.Input)
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

		if actual.SectionGroupId1 != v.Expected.SectionGroupId1 {
			t.Fatalf("Expected %q but got %q for SectionGroupId1", v.Expected.SectionGroupId1, actual.SectionGroupId1)
		}

	}
}

func TestSegmentsForUserOnenoteSectionGroupSectionGroupId(t *testing.T) {
	segments := UserOnenoteSectionGroupSectionGroupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserOnenoteSectionGroupSectionGroupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
