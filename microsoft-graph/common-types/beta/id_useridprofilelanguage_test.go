package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileLanguageId{}

func TestNewUserIdProfileLanguageID(t *testing.T) {
	id := NewUserIdProfileLanguageID("userId", "languageProficiencyId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.LanguageProficiencyId != "languageProficiencyId" {
		t.Fatalf("Expected %q but got %q for Segment 'LanguageProficiencyId'", id.LanguageProficiencyId, "languageProficiencyId")
	}
}

func TestFormatUserIdProfileLanguageID(t *testing.T) {
	actual := NewUserIdProfileLanguageID("userId", "languageProficiencyId").ID()
	expected := "/users/userId/profile/languages/languageProficiencyId"
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/profile",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/profile/languages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/languages/languageProficiencyId",
			Expected: &UserIdProfileLanguageId{
				UserId:                "userId",
				LanguageProficiencyId: "languageProficiencyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/languages/languageProficiencyId/extra",
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
			Input: "/users/userId/profile",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/profile/languages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/lAnGuAgEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/languages/languageProficiencyId",
			Expected: &UserIdProfileLanguageId{
				UserId:                "userId",
				LanguageProficiencyId: "languageProficiencyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/languages/languageProficiencyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/lAnGuAgEs/lAnGuAgEpRoFiCiEnCyId",
			Expected: &UserIdProfileLanguageId{
				UserId:                "uSeRiD",
				LanguageProficiencyId: "lAnGuAgEpRoFiCiEnCyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/lAnGuAgEs/lAnGuAgEpRoFiCiEnCyId/extra",
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
