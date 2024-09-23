package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfilePublicationId{}

func TestNewUserIdProfilePublicationID(t *testing.T) {
	id := NewUserIdProfilePublicationID("userId", "itemPublicationId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ItemPublicationId != "itemPublicationId" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemPublicationId'", id.ItemPublicationId, "itemPublicationId")
	}
}

func TestFormatUserIdProfilePublicationID(t *testing.T) {
	actual := NewUserIdProfilePublicationID("userId", "itemPublicationId").ID()
	expected := "/users/userId/profile/publications/itemPublicationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdProfilePublicationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfilePublicationId
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
			Input: "/users/userId/profile/publications",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/publications/itemPublicationId",
			Expected: &UserIdProfilePublicationId{
				UserId:            "userId",
				ItemPublicationId: "itemPublicationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/publications/itemPublicationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfilePublicationID(v.Input)
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

		if actual.ItemPublicationId != v.Expected.ItemPublicationId {
			t.Fatalf("Expected %q but got %q for ItemPublicationId", v.Expected.ItemPublicationId, actual.ItemPublicationId)
		}

	}
}

func TestParseUserIdProfilePublicationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfilePublicationId
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
			Input: "/users/userId/profile/publications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/pUbLiCaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/publications/itemPublicationId",
			Expected: &UserIdProfilePublicationId{
				UserId:            "userId",
				ItemPublicationId: "itemPublicationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/publications/itemPublicationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/pUbLiCaTiOnS/iTeMpUbLiCaTiOnId",
			Expected: &UserIdProfilePublicationId{
				UserId:            "uSeRiD",
				ItemPublicationId: "iTeMpUbLiCaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/pUbLiCaTiOnS/iTeMpUbLiCaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfilePublicationIDInsensitively(v.Input)
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

		if actual.ItemPublicationId != v.Expected.ItemPublicationId {
			t.Fatalf("Expected %q but got %q for ItemPublicationId", v.Expected.ItemPublicationId, actual.ItemPublicationId)
		}

	}
}

func TestSegmentsForUserIdProfilePublicationId(t *testing.T) {
	segments := UserIdProfilePublicationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdProfilePublicationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
