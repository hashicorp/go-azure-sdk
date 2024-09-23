package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfilePatentId{}

func TestNewUserIdProfilePatentID(t *testing.T) {
	id := NewUserIdProfilePatentID("userId", "itemPatentId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ItemPatentId != "itemPatentId" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemPatentId'", id.ItemPatentId, "itemPatentId")
	}
}

func TestFormatUserIdProfilePatentID(t *testing.T) {
	actual := NewUserIdProfilePatentID("userId", "itemPatentId").ID()
	expected := "/users/userId/profile/patents/itemPatentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdProfilePatentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfilePatentId
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
			Input: "/users/userId/profile/patents",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/patents/itemPatentId",
			Expected: &UserIdProfilePatentId{
				UserId:       "userId",
				ItemPatentId: "itemPatentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/patents/itemPatentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfilePatentID(v.Input)
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

		if actual.ItemPatentId != v.Expected.ItemPatentId {
			t.Fatalf("Expected %q but got %q for ItemPatentId", v.Expected.ItemPatentId, actual.ItemPatentId)
		}

	}
}

func TestParseUserIdProfilePatentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfilePatentId
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
			Input: "/users/userId/profile/patents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/pAtEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/patents/itemPatentId",
			Expected: &UserIdProfilePatentId{
				UserId:       "userId",
				ItemPatentId: "itemPatentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/patents/itemPatentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/pAtEnTs/iTeMpAtEnTiD",
			Expected: &UserIdProfilePatentId{
				UserId:       "uSeRiD",
				ItemPatentId: "iTeMpAtEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/pAtEnTs/iTeMpAtEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfilePatentIDInsensitively(v.Input)
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

		if actual.ItemPatentId != v.Expected.ItemPatentId {
			t.Fatalf("Expected %q but got %q for ItemPatentId", v.Expected.ItemPatentId, actual.ItemPatentId)
		}

	}
}

func TestSegmentsForUserIdProfilePatentId(t *testing.T) {
	segments := UserIdProfilePatentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdProfilePatentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
