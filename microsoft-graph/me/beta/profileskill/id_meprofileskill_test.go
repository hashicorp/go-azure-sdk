package profileskill

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileSkillId{}

func TestNewMeProfileSkillID(t *testing.T) {
	id := NewMeProfileSkillID("skillProficiencyIdValue")

	if id.SkillProficiencyId != "skillProficiencyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SkillProficiencyId'", id.SkillProficiencyId, "skillProficiencyIdValue")
	}
}

func TestFormatMeProfileSkillID(t *testing.T) {
	actual := NewMeProfileSkillID("skillProficiencyIdValue").ID()
	expected := "/me/profile/skills/skillProficiencyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeProfileSkillID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileSkillId
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
			Input: "/me/profile",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/profile/skills",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/skills/skillProficiencyIdValue",
			Expected: &MeProfileSkillId{
				SkillProficiencyId: "skillProficiencyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/skills/skillProficiencyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileSkillID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SkillProficiencyId != v.Expected.SkillProficiencyId {
			t.Fatalf("Expected %q but got %q for SkillProficiencyId", v.Expected.SkillProficiencyId, actual.SkillProficiencyId)
		}

	}
}

func TestParseMeProfileSkillIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileSkillId
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
			Input: "/me/profile",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/profile/skills",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/sKiLlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/skills/skillProficiencyIdValue",
			Expected: &MeProfileSkillId{
				SkillProficiencyId: "skillProficiencyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/skills/skillProficiencyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/sKiLlS/sKiLlPrOfIcIeNcYiDvAlUe",
			Expected: &MeProfileSkillId{
				SkillProficiencyId: "sKiLlPrOfIcIeNcYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/sKiLlS/sKiLlPrOfIcIeNcYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileSkillIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SkillProficiencyId != v.Expected.SkillProficiencyId {
			t.Fatalf("Expected %q but got %q for SkillProficiencyId", v.Expected.SkillProficiencyId, actual.SkillProficiencyId)
		}

	}
}

func TestSegmentsForMeProfileSkillId(t *testing.T) {
	segments := MeProfileSkillId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeProfileSkillId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
