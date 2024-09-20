package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdTodoListIdTaskIdAttachmentSessionId{}

func TestNewUserIdTodoListIdTaskIdAttachmentSessionID(t *testing.T) {
	id := NewUserIdTodoListIdTaskIdAttachmentSessionID("userId", "todoTaskListId", "todoTaskId", "attachmentSessionId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.TodoTaskListId != "todoTaskListId" {
		t.Fatalf("Expected %q but got %q for Segment 'TodoTaskListId'", id.TodoTaskListId, "todoTaskListId")
	}

	if id.TodoTaskId != "todoTaskId" {
		t.Fatalf("Expected %q but got %q for Segment 'TodoTaskId'", id.TodoTaskId, "todoTaskId")
	}

	if id.AttachmentSessionId != "attachmentSessionId" {
		t.Fatalf("Expected %q but got %q for Segment 'AttachmentSessionId'", id.AttachmentSessionId, "attachmentSessionId")
	}
}

func TestFormatUserIdTodoListIdTaskIdAttachmentSessionID(t *testing.T) {
	actual := NewUserIdTodoListIdTaskIdAttachmentSessionID("userId", "todoTaskListId", "todoTaskId", "attachmentSessionId").ID()
	expected := "/users/userId/todo/lists/todoTaskListId/tasks/todoTaskId/attachmentSessions/attachmentSessionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdTodoListIdTaskIdAttachmentSessionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdTodoListIdTaskIdAttachmentSessionId
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
			Input: "/users/userId/todo",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/todo/lists",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/todo/lists/todoTaskListId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/todo/lists/todoTaskListId/tasks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/todo/lists/todoTaskListId/tasks/todoTaskId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/todo/lists/todoTaskListId/tasks/todoTaskId/attachmentSessions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/todo/lists/todoTaskListId/tasks/todoTaskId/attachmentSessions/attachmentSessionId",
			Expected: &UserIdTodoListIdTaskIdAttachmentSessionId{
				UserId:              "userId",
				TodoTaskListId:      "todoTaskListId",
				TodoTaskId:          "todoTaskId",
				AttachmentSessionId: "attachmentSessionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/todo/lists/todoTaskListId/tasks/todoTaskId/attachmentSessions/attachmentSessionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdTodoListIdTaskIdAttachmentSessionID(v.Input)
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

		if actual.TodoTaskListId != v.Expected.TodoTaskListId {
			t.Fatalf("Expected %q but got %q for TodoTaskListId", v.Expected.TodoTaskListId, actual.TodoTaskListId)
		}

		if actual.TodoTaskId != v.Expected.TodoTaskId {
			t.Fatalf("Expected %q but got %q for TodoTaskId", v.Expected.TodoTaskId, actual.TodoTaskId)
		}

		if actual.AttachmentSessionId != v.Expected.AttachmentSessionId {
			t.Fatalf("Expected %q but got %q for AttachmentSessionId", v.Expected.AttachmentSessionId, actual.AttachmentSessionId)
		}

	}
}

func TestParseUserIdTodoListIdTaskIdAttachmentSessionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdTodoListIdTaskIdAttachmentSessionId
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
			Input: "/users/userId/todo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/tOdO",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/todo/lists",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/tOdO/lIsTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/todo/lists/todoTaskListId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/tOdO/lIsTs/tOdOtAsKlIsTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/todo/lists/todoTaskListId/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/tOdO/lIsTs/tOdOtAsKlIsTiD/tAsKs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/todo/lists/todoTaskListId/tasks/todoTaskId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/tOdO/lIsTs/tOdOtAsKlIsTiD/tAsKs/tOdOtAsKiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/todo/lists/todoTaskListId/tasks/todoTaskId/attachmentSessions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/tOdO/lIsTs/tOdOtAsKlIsTiD/tAsKs/tOdOtAsKiD/aTtAcHmEnTsEsSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/todo/lists/todoTaskListId/tasks/todoTaskId/attachmentSessions/attachmentSessionId",
			Expected: &UserIdTodoListIdTaskIdAttachmentSessionId{
				UserId:              "userId",
				TodoTaskListId:      "todoTaskListId",
				TodoTaskId:          "todoTaskId",
				AttachmentSessionId: "attachmentSessionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/todo/lists/todoTaskListId/tasks/todoTaskId/attachmentSessions/attachmentSessionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/tOdO/lIsTs/tOdOtAsKlIsTiD/tAsKs/tOdOtAsKiD/aTtAcHmEnTsEsSiOnS/aTtAcHmEnTsEsSiOnId",
			Expected: &UserIdTodoListIdTaskIdAttachmentSessionId{
				UserId:              "uSeRiD",
				TodoTaskListId:      "tOdOtAsKlIsTiD",
				TodoTaskId:          "tOdOtAsKiD",
				AttachmentSessionId: "aTtAcHmEnTsEsSiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/tOdO/lIsTs/tOdOtAsKlIsTiD/tAsKs/tOdOtAsKiD/aTtAcHmEnTsEsSiOnS/aTtAcHmEnTsEsSiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdTodoListIdTaskIdAttachmentSessionIDInsensitively(v.Input)
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

		if actual.TodoTaskListId != v.Expected.TodoTaskListId {
			t.Fatalf("Expected %q but got %q for TodoTaskListId", v.Expected.TodoTaskListId, actual.TodoTaskListId)
		}

		if actual.TodoTaskId != v.Expected.TodoTaskId {
			t.Fatalf("Expected %q but got %q for TodoTaskId", v.Expected.TodoTaskId, actual.TodoTaskId)
		}

		if actual.AttachmentSessionId != v.Expected.AttachmentSessionId {
			t.Fatalf("Expected %q but got %q for AttachmentSessionId", v.Expected.AttachmentSessionId, actual.AttachmentSessionId)
		}

	}
}

func TestSegmentsForUserIdTodoListIdTaskIdAttachmentSessionId(t *testing.T) {
	segments := UserIdTodoListIdTaskIdAttachmentSessionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdTodoListIdTaskIdAttachmentSessionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
