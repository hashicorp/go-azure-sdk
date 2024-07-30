package todolisttaskattachmentsessioncontent

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdTodoListIdTaskIdAttachmentSessionId{}

func TestNewUserIdTodoListIdTaskIdAttachmentSessionID(t *testing.T) {
	id := NewUserIdTodoListIdTaskIdAttachmentSessionID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue", "attachmentSessionIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.TodoTaskListId != "todoTaskListIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TodoTaskListId'", id.TodoTaskListId, "todoTaskListIdValue")
	}

	if id.TodoTaskId != "todoTaskIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TodoTaskId'", id.TodoTaskId, "todoTaskIdValue")
	}

	if id.AttachmentSessionId != "attachmentSessionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AttachmentSessionId'", id.AttachmentSessionId, "attachmentSessionIdValue")
	}
}

func TestFormatUserIdTodoListIdTaskIdAttachmentSessionID(t *testing.T) {
	actual := NewUserIdTodoListIdTaskIdAttachmentSessionID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue", "attachmentSessionIdValue").ID()
	expected := "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachmentSessions/attachmentSessionIdValue"
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/todo",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/todo/lists",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachmentSessions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachmentSessions/attachmentSessionIdValue",
			Expected: &UserIdTodoListIdTaskIdAttachmentSessionId{
				UserId:              "userIdValue",
				TodoTaskListId:      "todoTaskListIdValue",
				TodoTaskId:          "todoTaskIdValue",
				AttachmentSessionId: "attachmentSessionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachmentSessions/attachmentSessionIdValue/extra",
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
			Input: "/users/userIdValue/todo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tOdO",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/todo/lists",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tOdO/lIsTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachmentSessions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe/aTtAcHmEnTsEsSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachmentSessions/attachmentSessionIdValue",
			Expected: &UserIdTodoListIdTaskIdAttachmentSessionId{
				UserId:              "userIdValue",
				TodoTaskListId:      "todoTaskListIdValue",
				TodoTaskId:          "todoTaskIdValue",
				AttachmentSessionId: "attachmentSessionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachmentSessions/attachmentSessionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe/aTtAcHmEnTsEsSiOnS/aTtAcHmEnTsEsSiOnIdVaLuE",
			Expected: &UserIdTodoListIdTaskIdAttachmentSessionId{
				UserId:              "uSeRiDvAlUe",
				TodoTaskListId:      "tOdOtAsKlIsTiDvAlUe",
				TodoTaskId:          "tOdOtAsKiDvAlUe",
				AttachmentSessionId: "aTtAcHmEnTsEsSiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe/aTtAcHmEnTsEsSiOnS/aTtAcHmEnTsEsSiOnIdVaLuE/extra",
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
