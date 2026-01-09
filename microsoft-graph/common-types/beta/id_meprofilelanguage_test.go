package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileLanguageId{}

func TestNewMeProfileLanguageID(t *testing.T) {
	id := NewMeProfileLanguageID("languageProficiencyId")

	if id.LanguageProficiencyId != "languageProficiencyId" {
		t.Fatalf("Expected %q but got %q for Segment 'LanguageProficiencyId'", id.LanguageProficiencyId, "languageProficiencyId")
	}
}

func TestFormatMeProfileLanguageID(t *testing.T) {
	actual := NewMeProfileLanguageID("languageProficiencyId").ID()
	expected := "/me/profile/languages/languageProficiencyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeProfileLanguageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileLanguageId
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
			Input: "/me/profile/languages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/languages/languageProficiencyId",
			Expected: &MeProfileLanguageId{
				LanguageProficiencyId: "languageProficiencyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/languages/languageProficiencyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileLanguageID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.LanguageProficiencyId != v.Expected.LanguageProficiencyId {
			t.Fatalf("Expected %q but got %q for LanguageProficiencyId", v.Expected.LanguageProficiencyId, actual.LanguageProficiencyId)
		}

	}
}

func TestParseMeProfileLanguageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileLanguageId
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
			Input: "/me/profile/languages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/lAnGuAgEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/languages/languageProficiencyId",
			Expected: &MeProfileLanguageId{
				LanguageProficiencyId: "languageProficiencyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/languages/languageProficiencyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/lAnGuAgEs/lAnGuAgEpRoFiCiEnCyId",
			Expected: &MeProfileLanguageId{
				LanguageProficiencyId: "lAnGuAgEpRoFiCiEnCyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/lAnGuAgEs/lAnGuAgEpRoFiCiEnCyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileLanguageIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.LanguageProficiencyId != v.Expected.LanguageProficiencyId {
			t.Fatalf("Expected %q but got %q for LanguageProficiencyId", v.Expected.LanguageProficiencyId, actual.LanguageProficiencyId)
		}

	}
}

func TestSegmentsForMeProfileLanguageId(t *testing.T) {
	segments := MeProfileLanguageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeProfileLanguageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
