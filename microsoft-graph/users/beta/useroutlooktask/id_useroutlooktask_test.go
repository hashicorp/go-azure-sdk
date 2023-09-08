package useroutlooktask

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOutlookTaskId{}

func TestNewUserOutlookTaskID(t *testing.T) {
	id := NewUserOutlookTaskID("userIdValue", "outlookTaskIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.OutlookTaskId != "outlookTaskIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskId'", id.OutlookTaskId, "outlookTaskIdValue")
	}
}

func TestFormatUserOutlookTaskID(t *testing.T) {
	actual := NewUserOutlookTaskID("userIdValue", "outlookTaskIdValue").ID()
	expected := "/users/userIdValue/outlook/tasks/outlookTaskIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserOutlookTaskID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOutlookTaskId
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
			Input: "/users/userIdValue/outlook",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/outlook/tasks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/outlook/tasks/outlookTaskIdValue",
			Expected: &UserOutlookTaskId{
				UserId:        "userIdValue",
				OutlookTaskId: "outlookTaskIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/outlook/tasks/outlookTaskIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOutlookTaskID(v.Input)
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

func TestParseUserOutlookTaskIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOutlookTaskId
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
			Input: "/users/userIdValue/outlook",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oUtLoOk",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/outlook/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oUtLoOk/tAsKs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/outlook/tasks/outlookTaskIdValue",
			Expected: &UserOutlookTaskId{
				UserId:        "userIdValue",
				OutlookTaskId: "outlookTaskIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/outlook/tasks/outlookTaskIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oUtLoOk/tAsKs/oUtLoOkTaSkIdVaLuE",
			Expected: &UserOutlookTaskId{
				UserId:        "uSeRiDvAlUe",
				OutlookTaskId: "oUtLoOkTaSkIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oUtLoOk/tAsKs/oUtLoOkTaSkIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOutlookTaskIDInsensitively(v.Input)
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

func TestSegmentsForUserOutlookTaskId(t *testing.T) {
	segments := UserOutlookTaskId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserOutlookTaskId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
