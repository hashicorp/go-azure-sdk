package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdActivityIdHistoryItemId{}

func TestNewUserIdActivityIdHistoryItemID(t *testing.T) {
	id := NewUserIdActivityIdHistoryItemID("userId", "userActivityId", "activityHistoryItemId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.UserActivityId != "userActivityId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserActivityId'", id.UserActivityId, "userActivityId")
	}

	if id.ActivityHistoryItemId != "activityHistoryItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'ActivityHistoryItemId'", id.ActivityHistoryItemId, "activityHistoryItemId")
	}
}

func TestFormatUserIdActivityIdHistoryItemID(t *testing.T) {
	actual := NewUserIdActivityIdHistoryItemID("userId", "userActivityId", "activityHistoryItemId").ID()
	expected := "/users/userId/activities/userActivityId/historyItems/activityHistoryItemId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdActivityIdHistoryItemID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdActivityIdHistoryItemId
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
			Input: "/users/userId/activities",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/activities/userActivityId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/activities/userActivityId/historyItems",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/activities/userActivityId/historyItems/activityHistoryItemId",
			Expected: &UserIdActivityIdHistoryItemId{
				UserId:                "userId",
				UserActivityId:        "userActivityId",
				ActivityHistoryItemId: "activityHistoryItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/activities/userActivityId/historyItems/activityHistoryItemId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdActivityIdHistoryItemID(v.Input)
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

		if actual.UserActivityId != v.Expected.UserActivityId {
			t.Fatalf("Expected %q but got %q for UserActivityId", v.Expected.UserActivityId, actual.UserActivityId)
		}

		if actual.ActivityHistoryItemId != v.Expected.ActivityHistoryItemId {
			t.Fatalf("Expected %q but got %q for ActivityHistoryItemId", v.Expected.ActivityHistoryItemId, actual.ActivityHistoryItemId)
		}

	}
}

func TestParseUserIdActivityIdHistoryItemIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdActivityIdHistoryItemId
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
			Input: "/users/userId/activities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aCtIvItIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/activities/userActivityId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aCtIvItIeS/uSeRaCtIvItYiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/activities/userActivityId/historyItems",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aCtIvItIeS/uSeRaCtIvItYiD/hIsToRyItEmS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/activities/userActivityId/historyItems/activityHistoryItemId",
			Expected: &UserIdActivityIdHistoryItemId{
				UserId:                "userId",
				UserActivityId:        "userActivityId",
				ActivityHistoryItemId: "activityHistoryItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/activities/userActivityId/historyItems/activityHistoryItemId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aCtIvItIeS/uSeRaCtIvItYiD/hIsToRyItEmS/aCtIvItYhIsToRyItEmId",
			Expected: &UserIdActivityIdHistoryItemId{
				UserId:                "uSeRiD",
				UserActivityId:        "uSeRaCtIvItYiD",
				ActivityHistoryItemId: "aCtIvItYhIsToRyItEmId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aCtIvItIeS/uSeRaCtIvItYiD/hIsToRyItEmS/aCtIvItYhIsToRyItEmId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdActivityIdHistoryItemIDInsensitively(v.Input)
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

		if actual.UserActivityId != v.Expected.UserActivityId {
			t.Fatalf("Expected %q but got %q for UserActivityId", v.Expected.UserActivityId, actual.UserActivityId)
		}

		if actual.ActivityHistoryItemId != v.Expected.ActivityHistoryItemId {
			t.Fatalf("Expected %q but got %q for ActivityHistoryItemId", v.Expected.ActivityHistoryItemId, actual.ActivityHistoryItemId)
		}

	}
}

func TestSegmentsForUserIdActivityIdHistoryItemId(t *testing.T) {
	segments := UserIdActivityIdHistoryItemId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdActivityIdHistoryItemId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
