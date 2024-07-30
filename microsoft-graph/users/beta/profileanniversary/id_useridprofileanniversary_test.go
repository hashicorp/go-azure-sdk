package profileanniversary

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileAnniversaryId{}

func TestNewUserIdProfileAnniversaryID(t *testing.T) {
	id := NewUserIdProfileAnniversaryID("userIdValue", "personAnnualEventIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.PersonAnnualEventId != "personAnnualEventIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonAnnualEventId'", id.PersonAnnualEventId, "personAnnualEventIdValue")
	}
}

func TestFormatUserIdProfileAnniversaryID(t *testing.T) {
	actual := NewUserIdProfileAnniversaryID("userIdValue", "personAnnualEventIdValue").ID()
	expected := "/users/userIdValue/profile/anniversaries/personAnnualEventIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdProfileAnniversaryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileAnniversaryId
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
			Input: "/users/userIdValue/profile/anniversaries",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/anniversaries/personAnnualEventIdValue",
			Expected: &UserIdProfileAnniversaryId{
				UserId:              "userIdValue",
				PersonAnnualEventId: "personAnnualEventIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/anniversaries/personAnnualEventIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileAnniversaryID(v.Input)
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

		if actual.PersonAnnualEventId != v.Expected.PersonAnnualEventId {
			t.Fatalf("Expected %q but got %q for PersonAnnualEventId", v.Expected.PersonAnnualEventId, actual.PersonAnnualEventId)
		}

	}
}

func TestParseUserIdProfileAnniversaryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileAnniversaryId
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
			Input: "/users/userIdValue/profile/anniversaries",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/aNnIvErSaRiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/anniversaries/personAnnualEventIdValue",
			Expected: &UserIdProfileAnniversaryId{
				UserId:              "userIdValue",
				PersonAnnualEventId: "personAnnualEventIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/anniversaries/personAnnualEventIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/aNnIvErSaRiEs/pErSoNaNnUaLeVeNtIdVaLuE",
			Expected: &UserIdProfileAnniversaryId{
				UserId:              "uSeRiDvAlUe",
				PersonAnnualEventId: "pErSoNaNnUaLeVeNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/aNnIvErSaRiEs/pErSoNaNnUaLeVeNtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileAnniversaryIDInsensitively(v.Input)
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

		if actual.PersonAnnualEventId != v.Expected.PersonAnnualEventId {
			t.Fatalf("Expected %q but got %q for PersonAnnualEventId", v.Expected.PersonAnnualEventId, actual.PersonAnnualEventId)
		}

	}
}

func TestSegmentsForUserIdProfileAnniversaryId(t *testing.T) {
	segments := UserIdProfileAnniversaryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdProfileAnniversaryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
