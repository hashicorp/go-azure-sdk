package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileLanguageId{}

func TestNewUserIdProfileLanguageID(t *testing.T) {
	id := NewUserIdProfileLanguageID("userIdValue", "languageProficiencyIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.LanguageProficiencyId != "languageProficiencyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LanguageProficiencyId'", id.LanguageProficiencyId, "languageProficiencyIdValue")
	}
}

func TestFormatUserIdProfileLanguageID(t *testing.T) {
	actual := NewUserIdProfileLanguageID("userIdValue", "languageProficiencyIdValue").ID()
	expected := "/users/userIdValue/profile/languages/languageProficiencyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdProfileLanguageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileLanguageId
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
			Input: "/users/userIdValue/profile/languages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/languages/languageProficiencyIdValue",
			Expected: &UserIdProfileLanguageId{
				UserId:                "userIdValue",
				LanguageProficiencyId: "languageProficiencyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/languages/languageProficiencyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileLanguageID(v.Input)
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

		if actual.LanguageProficiencyId != v.Expected.LanguageProficiencyId {
			t.Fatalf("Expected %q but got %q for LanguageProficiencyId", v.Expected.LanguageProficiencyId, actual.LanguageProficiencyId)
		}

	}
}

func TestParseUserIdProfileLanguageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileLanguageId
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
			Input: "/users/userIdValue/profile/languages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/lAnGuAgEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/languages/languageProficiencyIdValue",
			Expected: &UserIdProfileLanguageId{
				UserId:                "userIdValue",
				LanguageProficiencyId: "languageProficiencyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/languages/languageProficiencyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/lAnGuAgEs/lAnGuAgEpRoFiCiEnCyIdVaLuE",
			Expected: &UserIdProfileLanguageId{
				UserId:                "uSeRiDvAlUe",
				LanguageProficiencyId: "lAnGuAgEpRoFiCiEnCyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/lAnGuAgEs/lAnGuAgEpRoFiCiEnCyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileLanguageIDInsensitively(v.Input)
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

		if actual.LanguageProficiencyId != v.Expected.LanguageProficiencyId {
			t.Fatalf("Expected %q but got %q for LanguageProficiencyId", v.Expected.LanguageProficiencyId, actual.LanguageProficiencyId)
		}

	}
}

func TestSegmentsForUserIdProfileLanguageId(t *testing.T) {
	segments := UserIdProfileLanguageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdProfileLanguageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
