package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeActivityIdHistoryItemId{}

func TestNewMeActivityIdHistoryItemID(t *testing.T) {
	id := NewMeActivityIdHistoryItemID("userActivityId", "activityHistoryItemId")

	if id.UserActivityId != "userActivityId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserActivityId'", id.UserActivityId, "userActivityId")
	}

	if id.ActivityHistoryItemId != "activityHistoryItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'ActivityHistoryItemId'", id.ActivityHistoryItemId, "activityHistoryItemId")
	}
}

func TestFormatMeActivityIdHistoryItemID(t *testing.T) {
	actual := NewMeActivityIdHistoryItemID("userActivityId", "activityHistoryItemId").ID()
	expected := "/me/activities/userActivityId/historyItems/activityHistoryItemId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeActivityIdHistoryItemID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeActivityIdHistoryItemId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/activities",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/activities/userActivityId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/activities/userActivityId/historyItems",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/activities/userActivityId/historyItems/activityHistoryItemId",
			Expected: &MeActivityIdHistoryItemId{
				UserActivityId:        "userActivityId",
				ActivityHistoryItemId: "activityHistoryItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/activities/userActivityId/historyItems/activityHistoryItemId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeActivityIdHistoryItemID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserActivityId != v.Expected.UserActivityId {
			t.Fatalf("Expected %q but got %q for UserActivityId", v.Expected.UserActivityId, actual.UserActivityId)
		}

		if actual.ActivityHistoryItemId != v.Expected.ActivityHistoryItemId {
			t.Fatalf("Expected %q but got %q for ActivityHistoryItemId", v.Expected.ActivityHistoryItemId, actual.ActivityHistoryItemId)
		}

	}
}

func TestParseMeActivityIdHistoryItemIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeActivityIdHistoryItemId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/activities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aCtIvItIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/activities/userActivityId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aCtIvItIeS/uSeRaCtIvItYiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/activities/userActivityId/historyItems",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aCtIvItIeS/uSeRaCtIvItYiD/hIsToRyItEmS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/activities/userActivityId/historyItems/activityHistoryItemId",
			Expected: &MeActivityIdHistoryItemId{
				UserActivityId:        "userActivityId",
				ActivityHistoryItemId: "activityHistoryItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/activities/userActivityId/historyItems/activityHistoryItemId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aCtIvItIeS/uSeRaCtIvItYiD/hIsToRyItEmS/aCtIvItYhIsToRyItEmId",
			Expected: &MeActivityIdHistoryItemId{
				UserActivityId:        "uSeRaCtIvItYiD",
				ActivityHistoryItemId: "aCtIvItYhIsToRyItEmId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aCtIvItIeS/uSeRaCtIvItYiD/hIsToRyItEmS/aCtIvItYhIsToRyItEmId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeActivityIdHistoryItemIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserActivityId != v.Expected.UserActivityId {
			t.Fatalf("Expected %q but got %q for UserActivityId", v.Expected.UserActivityId, actual.UserActivityId)
		}

		if actual.ActivityHistoryItemId != v.Expected.ActivityHistoryItemId {
			t.Fatalf("Expected %q but got %q for ActivityHistoryItemId", v.Expected.ActivityHistoryItemId, actual.ActivityHistoryItemId)
		}

	}
}

func TestSegmentsForMeActivityIdHistoryItemId(t *testing.T) {
	segments := MeActivityIdHistoryItemId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeActivityIdHistoryItemId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
