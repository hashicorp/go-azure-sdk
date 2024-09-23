package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnenoteSectionGroupIdSectionIdPageId{}

func TestNewUserIdOnenoteSectionGroupIdSectionIdPageID(t *testing.T) {
	id := NewUserIdOnenoteSectionGroupIdSectionIdPageID("userId", "sectionGroupId", "onenoteSectionId", "onenotePageId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.SectionGroupId != "sectionGroupId" {
		t.Fatalf("Expected %q but got %q for Segment 'SectionGroupId'", id.SectionGroupId, "sectionGroupId")
	}

	if id.OnenoteSectionId != "onenoteSectionId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenoteSectionId'", id.OnenoteSectionId, "onenoteSectionId")
	}

	if id.OnenotePageId != "onenotePageId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenotePageId'", id.OnenotePageId, "onenotePageId")
	}
}

func TestFormatUserIdOnenoteSectionGroupIdSectionIdPageID(t *testing.T) {
	actual := NewUserIdOnenoteSectionGroupIdSectionIdPageID("userId", "sectionGroupId", "onenoteSectionId", "onenotePageId").ID()
	expected := "/users/userId/onenote/sectionGroups/sectionGroupId/sections/onenoteSectionId/pages/onenotePageId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdOnenoteSectionGroupIdSectionIdPageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOnenoteSectionGroupIdSectionIdPageId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onenote",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onenote/sectionGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onenote/sectionGroups/sectionGroupId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onenote/sectionGroups/sectionGroupId/sections",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onenote/sectionGroups/sectionGroupId/sections/onenoteSectionId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onenote/sectionGroups/sectionGroupId/sections/onenoteSectionId/pages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/onenote/sectionGroups/sectionGroupId/sections/onenoteSectionId/pages/onenotePageId",
			Expected: &UserIdOnenoteSectionGroupIdSectionIdPageId{
				UserId:           "userId",
				SectionGroupId:   "sectionGroupId",
				OnenoteSectionId: "onenoteSectionId",
				OnenotePageId:    "onenotePageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/onenote/sectionGroups/sectionGroupId/sections/onenoteSectionId/pages/onenotePageId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOnenoteSectionGroupIdSectionIdPageID(v.Input)
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

		if actual.OnenotePageId != v.Expected.OnenotePageId {
			t.Fatalf("Expected %q but got %q for OnenotePageId", v.Expected.OnenotePageId, actual.OnenotePageId)
		}

	}
}

func TestParseUserIdOnenoteSectionGroupIdSectionIdPageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOnenoteSectionGroupIdSectionIdPageId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onenote",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNeNoTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onenote/sectionGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNeNoTe/sEcTiOnGrOuPs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onenote/sectionGroups/sectionGroupId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onenote/sectionGroups/sectionGroupId/sections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiD/sEcTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onenote/sectionGroups/sectionGroupId/sections/onenoteSectionId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiD/sEcTiOnS/oNeNoTeSeCtIoNiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onenote/sectionGroups/sectionGroupId/sections/onenoteSectionId/pages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiD/sEcTiOnS/oNeNoTeSeCtIoNiD/pAgEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/onenote/sectionGroups/sectionGroupId/sections/onenoteSectionId/pages/onenotePageId",
			Expected: &UserIdOnenoteSectionGroupIdSectionIdPageId{
				UserId:           "userId",
				SectionGroupId:   "sectionGroupId",
				OnenoteSectionId: "onenoteSectionId",
				OnenotePageId:    "onenotePageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/onenote/sectionGroups/sectionGroupId/sections/onenoteSectionId/pages/onenotePageId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiD/sEcTiOnS/oNeNoTeSeCtIoNiD/pAgEs/oNeNoTePaGeId",
			Expected: &UserIdOnenoteSectionGroupIdSectionIdPageId{
				UserId:           "uSeRiD",
				SectionGroupId:   "sEcTiOnGrOuPiD",
				OnenoteSectionId: "oNeNoTeSeCtIoNiD",
				OnenotePageId:    "oNeNoTePaGeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNeNoTe/sEcTiOnGrOuPs/sEcTiOnGrOuPiD/sEcTiOnS/oNeNoTeSeCtIoNiD/pAgEs/oNeNoTePaGeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOnenoteSectionGroupIdSectionIdPageIDInsensitively(v.Input)
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

		if actual.OnenotePageId != v.Expected.OnenotePageId {
			t.Fatalf("Expected %q but got %q for OnenotePageId", v.Expected.OnenotePageId, actual.OnenotePageId)
		}

	}
}

func TestSegmentsForUserIdOnenoteSectionGroupIdSectionIdPageId(t *testing.T) {
	segments := UserIdOnenoteSectionGroupIdSectionIdPageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdOnenoteSectionGroupIdSectionIdPageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
