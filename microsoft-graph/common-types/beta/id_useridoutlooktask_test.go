package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOutlookTaskId{}

func TestNewUserIdOutlookTaskID(t *testing.T) {
	id := NewUserIdOutlookTaskID("userId", "outlookTaskId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.OutlookTaskId != "outlookTaskId" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskId'", id.OutlookTaskId, "outlookTaskId")
	}
}

func TestFormatUserIdOutlookTaskID(t *testing.T) {
	actual := NewUserIdOutlookTaskID("userId", "outlookTaskId").ID()
	expected := "/users/userId/outlook/tasks/outlookTaskId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdOutlookTaskID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOutlookTaskId
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
			Input: "/users/userId/outlook",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook/tasks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/outlook/tasks/outlookTaskId",
			Expected: &UserIdOutlookTaskId{
				UserId:        "userId",
				OutlookTaskId: "outlookTaskId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/outlook/tasks/outlookTaskId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOutlookTaskID(v.Input)
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

		if actual.OutlookTaskId != v.Expected.OutlookTaskId {
			t.Fatalf("Expected %q but got %q for OutlookTaskId", v.Expected.OutlookTaskId, actual.OutlookTaskId)
		}

	}
}

func TestParseUserIdOutlookTaskIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOutlookTaskId
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
			Input: "/users/userId/outlook",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk/tAsKs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/outlook/tasks/outlookTaskId",
			Expected: &UserIdOutlookTaskId{
				UserId:        "userId",
				OutlookTaskId: "outlookTaskId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/outlook/tasks/outlookTaskId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk/tAsKs/oUtLoOkTaSkId",
			Expected: &UserIdOutlookTaskId{
				UserId:        "uSeRiD",
				OutlookTaskId: "oUtLoOkTaSkId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk/tAsKs/oUtLoOkTaSkId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOutlookTaskIDInsensitively(v.Input)
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

		if actual.OutlookTaskId != v.Expected.OutlookTaskId {
			t.Fatalf("Expected %q but got %q for OutlookTaskId", v.Expected.OutlookTaskId, actual.OutlookTaskId)
		}

	}
}

func TestSegmentsForUserIdOutlookTaskId(t *testing.T) {
	segments := UserIdOutlookTaskId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdOutlookTaskId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
