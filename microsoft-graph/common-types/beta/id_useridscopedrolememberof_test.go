package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdScopedRoleMemberOfId{}

func TestNewUserIdScopedRoleMemberOfID(t *testing.T) {
	id := NewUserIdScopedRoleMemberOfID("userId", "scopedRoleMembershipId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ScopedRoleMembershipId != "scopedRoleMembershipId" {
		t.Fatalf("Expected %q but got %q for Segment 'ScopedRoleMembershipId'", id.ScopedRoleMembershipId, "scopedRoleMembershipId")
	}
}

func TestFormatUserIdScopedRoleMemberOfID(t *testing.T) {
	actual := NewUserIdScopedRoleMemberOfID("userId", "scopedRoleMembershipId").ID()
	expected := "/users/userId/scopedRoleMemberOf/scopedRoleMembershipId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdScopedRoleMemberOfID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdScopedRoleMemberOfId
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
			Input: "/users/userId/scopedRoleMemberOf",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/scopedRoleMemberOf/scopedRoleMembershipId",
			Expected: &UserIdScopedRoleMemberOfId{
				UserId:                 "userId",
				ScopedRoleMembershipId: "scopedRoleMembershipId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/scopedRoleMemberOf/scopedRoleMembershipId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdScopedRoleMemberOfID(v.Input)
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

		if actual.ScopedRoleMembershipId != v.Expected.ScopedRoleMembershipId {
			t.Fatalf("Expected %q but got %q for ScopedRoleMembershipId", v.Expected.ScopedRoleMembershipId, actual.ScopedRoleMembershipId)
		}

	}
}

func TestParseUserIdScopedRoleMemberOfIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdScopedRoleMemberOfId
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
			Input: "/users/userId/scopedRoleMemberOf",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/sCoPeDrOlEmEmBeRoF",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/scopedRoleMemberOf/scopedRoleMembershipId",
			Expected: &UserIdScopedRoleMemberOfId{
				UserId:                 "userId",
				ScopedRoleMembershipId: "scopedRoleMembershipId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/scopedRoleMemberOf/scopedRoleMembershipId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/sCoPeDrOlEmEmBeRoF/sCoPeDrOlEmEmBeRsHiPiD",
			Expected: &UserIdScopedRoleMemberOfId{
				UserId:                 "uSeRiD",
				ScopedRoleMembershipId: "sCoPeDrOlEmEmBeRsHiPiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/sCoPeDrOlEmEmBeRoF/sCoPeDrOlEmEmBeRsHiPiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdScopedRoleMemberOfIDInsensitively(v.Input)
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

		if actual.ScopedRoleMembershipId != v.Expected.ScopedRoleMembershipId {
			t.Fatalf("Expected %q but got %q for ScopedRoleMembershipId", v.Expected.ScopedRoleMembershipId, actual.ScopedRoleMembershipId)
		}

	}
}

func TestSegmentsForUserIdScopedRoleMemberOfId(t *testing.T) {
	segments := UserIdScopedRoleMemberOfId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdScopedRoleMemberOfId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
