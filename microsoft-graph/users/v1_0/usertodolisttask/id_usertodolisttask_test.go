package usertodolisttask

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserTodoListTaskId{}

func TestNewUserTodoListTaskID(t *testing.T) {
	id := NewUserTodoListTaskID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.TodoTaskListId != "todoTaskListIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TodoTaskListId'", id.TodoTaskListId, "todoTaskListIdValue")
	}

	if id.TodoTaskId != "todoTaskIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TodoTaskId'", id.TodoTaskId, "todoTaskIdValue")
	}
}

func TestFormatUserTodoListTaskID(t *testing.T) {
	actual := NewUserTodoListTaskID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue").ID()
	expected := "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserTodoListTaskID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserTodoListTaskId
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
			// Valid URI
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue",
			Expected: &UserTodoListTaskId{
				UserId:         "userIdValue",
				TodoTaskListId: "todoTaskListIdValue",
				TodoTaskId:     "todoTaskIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserTodoListTaskID(v.Input)
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

	}
}

func TestParseUserTodoListTaskIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserTodoListTaskId
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
			// Valid URI
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue",
			Expected: &UserTodoListTaskId{
				UserId:         "userIdValue",
				TodoTaskListId: "todoTaskListIdValue",
				TodoTaskId:     "todoTaskIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe",
			Expected: &UserTodoListTaskId{
				UserId:         "uSeRiDvAlUe",
				TodoTaskListId: "tOdOtAsKlIsTiDvAlUe",
				TodoTaskId:     "tOdOtAsKiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserTodoListTaskIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForUserTodoListTaskId(t *testing.T) {
	segments := UserTodoListTaskId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserTodoListTaskId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
