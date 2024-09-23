package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileEmailId{}

func TestNewUserIdProfileEmailID(t *testing.T) {
	id := NewUserIdProfileEmailID("userId", "itemEmailId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ItemEmailId != "itemEmailId" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemEmailId'", id.ItemEmailId, "itemEmailId")
	}
}

func TestFormatUserIdProfileEmailID(t *testing.T) {
	actual := NewUserIdProfileEmailID("userId", "itemEmailId").ID()
	expected := "/users/userId/profile/emails/itemEmailId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdProfileEmailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileEmailId
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
			Input: "/users/userId/profile/emails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/emails/itemEmailId",
			Expected: &UserIdProfileEmailId{
				UserId:      "userId",
				ItemEmailId: "itemEmailId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/emails/itemEmailId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileEmailID(v.Input)
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

		if actual.ItemEmailId != v.Expected.ItemEmailId {
			t.Fatalf("Expected %q but got %q for ItemEmailId", v.Expected.ItemEmailId, actual.ItemEmailId)
		}

	}
}

func TestParseUserIdProfileEmailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileEmailId
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
			Input: "/users/userId/profile/emails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/eMaIlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/emails/itemEmailId",
			Expected: &UserIdProfileEmailId{
				UserId:      "userId",
				ItemEmailId: "itemEmailId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/emails/itemEmailId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/eMaIlS/iTeMeMaIlId",
			Expected: &UserIdProfileEmailId{
				UserId:      "uSeRiD",
				ItemEmailId: "iTeMeMaIlId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/eMaIlS/iTeMeMaIlId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileEmailIDInsensitively(v.Input)
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

		if actual.ItemEmailId != v.Expected.ItemEmailId {
			t.Fatalf("Expected %q but got %q for ItemEmailId", v.Expected.ItemEmailId, actual.ItemEmailId)
		}

	}
}

func TestSegmentsForUserIdProfileEmailId(t *testing.T) {
	segments := UserIdProfileEmailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdProfileEmailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
