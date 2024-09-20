package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileAnniversaryId{}

func TestNewUserIdProfileAnniversaryID(t *testing.T) {
	id := NewUserIdProfileAnniversaryID("userId", "personAnnualEventId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.PersonAnnualEventId != "personAnnualEventId" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonAnnualEventId'", id.PersonAnnualEventId, "personAnnualEventId")
	}
}

func TestFormatUserIdProfileAnniversaryID(t *testing.T) {
	actual := NewUserIdProfileAnniversaryID("userId", "personAnnualEventId").ID()
	expected := "/users/userId/profile/anniversaries/personAnnualEventId"
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
			Input: "/users/userId/profile/anniversaries",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/anniversaries/personAnnualEventId",
			Expected: &UserIdProfileAnniversaryId{
				UserId:              "userId",
				PersonAnnualEventId: "personAnnualEventId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/anniversaries/personAnnualEventId/extra",
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
			Input: "/users/userId/profile/anniversaries",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/aNnIvErSaRiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/anniversaries/personAnnualEventId",
			Expected: &UserIdProfileAnniversaryId{
				UserId:              "userId",
				PersonAnnualEventId: "personAnnualEventId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/anniversaries/personAnnualEventId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/aNnIvErSaRiEs/pErSoNaNnUaLeVeNtId",
			Expected: &UserIdProfileAnniversaryId{
				UserId:              "uSeRiD",
				PersonAnnualEventId: "pErSoNaNnUaLeVeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/aNnIvErSaRiEs/pErSoNaNnUaLeVeNtId/extra",
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
