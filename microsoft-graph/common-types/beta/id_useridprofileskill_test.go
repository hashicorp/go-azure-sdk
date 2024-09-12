package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileSkillId{}

func TestNewUserIdProfileSkillID(t *testing.T) {
	id := NewUserIdProfileSkillID("userIdValue", "skillProficiencyIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.SkillProficiencyId != "skillProficiencyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SkillProficiencyId'", id.SkillProficiencyId, "skillProficiencyIdValue")
	}
}

func TestFormatUserIdProfileSkillID(t *testing.T) {
	actual := NewUserIdProfileSkillID("userIdValue", "skillProficiencyIdValue").ID()
	expected := "/users/userIdValue/profile/skills/skillProficiencyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdProfileSkillID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileSkillId
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
			Input: "/users/userIdValue/profile",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/profile/skills",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/skills/skillProficiencyIdValue",
			Expected: &UserIdProfileSkillId{
				UserId:             "userIdValue",
				SkillProficiencyId: "skillProficiencyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/skills/skillProficiencyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileSkillID(v.Input)
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

		if actual.SkillProficiencyId != v.Expected.SkillProficiencyId {
			t.Fatalf("Expected %q but got %q for SkillProficiencyId", v.Expected.SkillProficiencyId, actual.SkillProficiencyId)
		}

	}
}

func TestParseUserIdProfileSkillIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileSkillId
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
			Input: "/users/userIdValue/profile",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/profile/skills",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/sKiLlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/skills/skillProficiencyIdValue",
			Expected: &UserIdProfileSkillId{
				UserId:             "userIdValue",
				SkillProficiencyId: "skillProficiencyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/skills/skillProficiencyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/sKiLlS/sKiLlPrOfIcIeNcYiDvAlUe",
			Expected: &UserIdProfileSkillId{
				UserId:             "uSeRiDvAlUe",
				SkillProficiencyId: "sKiLlPrOfIcIeNcYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/sKiLlS/sKiLlPrOfIcIeNcYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileSkillIDInsensitively(v.Input)
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

		if actual.SkillProficiencyId != v.Expected.SkillProficiencyId {
			t.Fatalf("Expected %q but got %q for SkillProficiencyId", v.Expected.SkillProficiencyId, actual.SkillProficiencyId)
		}

	}
}

func TestSegmentsForUserIdProfileSkillId(t *testing.T) {
	segments := UserIdProfileSkillId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdProfileSkillId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
