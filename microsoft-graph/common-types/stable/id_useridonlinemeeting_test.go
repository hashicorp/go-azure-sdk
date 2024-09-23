package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnlineMeetingId{}

func TestNewUserIdOnlineMeetingID(t *testing.T) {
	id := NewUserIdOnlineMeetingID("userId", "onlineMeetingId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.OnlineMeetingId != "onlineMeetingId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingId")
	}
}

func TestFormatUserIdOnlineMeetingID(t *testing.T) {
	actual := NewUserIdOnlineMeetingID("userId", "onlineMeetingId").ID()
	expected := "/users/userId/onlineMeetings/onlineMeetingId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdOnlineMeetingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOnlineMeetingId
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
			Input: "/users/userId/onlineMeetings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId",
			Expected: &UserIdOnlineMeetingId{
				UserId:          "userId",
				OnlineMeetingId: "onlineMeetingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/onlineMeetings/onlineMeetingId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOnlineMeetingID(v.Input)
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

		if actual.OnlineMeetingId != v.Expected.OnlineMeetingId {
			t.Fatalf("Expected %q but got %q for OnlineMeetingId", v.Expected.OnlineMeetingId, actual.OnlineMeetingId)
		}

	}
}

func TestParseUserIdOnlineMeetingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOnlineMeetingId
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
			Input: "/users/userId/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId",
			Expected: &UserIdOnlineMeetingId{
				UserId:          "userId",
				OnlineMeetingId: "onlineMeetingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/onlineMeetings/onlineMeetingId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId",
			Expected: &UserIdOnlineMeetingId{
				UserId:          "uSeRiD",
				OnlineMeetingId: "oNlInEmEeTiNgId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOnlineMeetingIDInsensitively(v.Input)
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

		if actual.OnlineMeetingId != v.Expected.OnlineMeetingId {
			t.Fatalf("Expected %q but got %q for OnlineMeetingId", v.Expected.OnlineMeetingId, actual.OnlineMeetingId)
		}

	}
}

func TestSegmentsForUserIdOnlineMeetingId(t *testing.T) {
	segments := UserIdOnlineMeetingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdOnlineMeetingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
