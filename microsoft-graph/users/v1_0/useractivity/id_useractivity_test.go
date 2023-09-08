package useractivity

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserActivityId{}

func TestNewUserActivityID(t *testing.T) {
	id := NewUserActivityID("userIdValue", "userActivityIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.UserActivityId != "userActivityIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserActivityId'", id.UserActivityId, "userActivityIdValue")
	}
}

func TestFormatUserActivityID(t *testing.T) {
	actual := NewUserActivityID("userIdValue", "userActivityIdValue").ID()
	expected := "/users/userIdValue/activities/userActivityIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserActivityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserActivityId
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
			Input: "/users/userIdValue/activities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/activities/userActivityIdValue",
			Expected: &UserActivityId{
				UserId:         "userIdValue",
				UserActivityId: "userActivityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/activities/userActivityIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserActivityID(v.Input)
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

	}
}

func TestParseUserActivityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserActivityId
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
			Input: "/users/userIdValue/activities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aCtIvItIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/activities/userActivityIdValue",
			Expected: &UserActivityId{
				UserId:         "userIdValue",
				UserActivityId: "userActivityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/activities/userActivityIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aCtIvItIeS/uSeRaCtIvItYiDvAlUe",
			Expected: &UserActivityId{
				UserId:         "uSeRiDvAlUe",
				UserActivityId: "uSeRaCtIvItYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aCtIvItIeS/uSeRaCtIvItYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserActivityIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForUserActivityId(t *testing.T) {
	segments := UserActivityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserActivityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
