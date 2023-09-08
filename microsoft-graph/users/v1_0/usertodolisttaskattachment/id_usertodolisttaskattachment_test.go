package usertodolisttaskattachment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserTodoListTaskAttachmentId{}

func TestNewUserTodoListTaskAttachmentID(t *testing.T) {
	id := NewUserTodoListTaskAttachmentID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.TodoTaskListId != "todoTaskListIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TodoTaskListId'", id.TodoTaskListId, "todoTaskListIdValue")
	}

	if id.TodoTaskId != "todoTaskIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TodoTaskId'", id.TodoTaskId, "todoTaskIdValue")
	}

	if id.AttachmentBaseId != "attachmentBaseIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AttachmentBaseId'", id.AttachmentBaseId, "attachmentBaseIdValue")
	}
}

func TestFormatUserTodoListTaskAttachmentID(t *testing.T) {
	actual := NewUserTodoListTaskAttachmentID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue").ID()
	expected := "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachments/attachmentBaseIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserTodoListTaskAttachmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserTodoListTaskAttachmentId
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
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachments/attachmentBaseIdValue",
			Expected: &UserTodoListTaskAttachmentId{
				UserId:           "userIdValue",
				TodoTaskListId:   "todoTaskListIdValue",
				TodoTaskId:       "todoTaskIdValue",
				AttachmentBaseId: "attachmentBaseIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachments/attachmentBaseIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserTodoListTaskAttachmentID(v.Input)
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

		if actual.AttachmentBaseId != v.Expected.AttachmentBaseId {
			t.Fatalf("Expected %q but got %q for AttachmentBaseId", v.Expected.AttachmentBaseId, actual.AttachmentBaseId)
		}

	}
}

func TestParseUserTodoListTaskAttachmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserTodoListTaskAttachmentId
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
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe/aTtAcHmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachments/attachmentBaseIdValue",
			Expected: &UserTodoListTaskAttachmentId{
				UserId:           "userIdValue",
				TodoTaskListId:   "todoTaskListIdValue",
				TodoTaskId:       "todoTaskIdValue",
				AttachmentBaseId: "attachmentBaseIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachments/attachmentBaseIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe/aTtAcHmEnTs/aTtAcHmEnTbAsEiDvAlUe",
			Expected: &UserTodoListTaskAttachmentId{
				UserId:           "uSeRiDvAlUe",
				TodoTaskListId:   "tOdOtAsKlIsTiDvAlUe",
				TodoTaskId:       "tOdOtAsKiDvAlUe",
				AttachmentBaseId: "aTtAcHmEnTbAsEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe/aTtAcHmEnTs/aTtAcHmEnTbAsEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserTodoListTaskAttachmentIDInsensitively(v.Input)
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

		if actual.AttachmentBaseId != v.Expected.AttachmentBaseId {
			t.Fatalf("Expected %q but got %q for AttachmentBaseId", v.Expected.AttachmentBaseId, actual.AttachmentBaseId)
		}

	}
}

func TestSegmentsForUserTodoListTaskAttachmentId(t *testing.T) {
	segments := UserTodoListTaskAttachmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserTodoListTaskAttachmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
