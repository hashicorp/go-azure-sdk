package usertodolisttasklinkedresource

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserTodoListTaskLinkedResourceId{}

func TestNewUserTodoListTaskLinkedResourceID(t *testing.T) {
	id := NewUserTodoListTaskLinkedResourceID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue", "linkedResourceIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.TodoTaskListId != "todoTaskListIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TodoTaskListId'", id.TodoTaskListId, "todoTaskListIdValue")
	}

	if id.TodoTaskId != "todoTaskIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TodoTaskId'", id.TodoTaskId, "todoTaskIdValue")
	}

	if id.LinkedResourceId != "linkedResourceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LinkedResourceId'", id.LinkedResourceId, "linkedResourceIdValue")
	}
}

func TestFormatUserTodoListTaskLinkedResourceID(t *testing.T) {
	actual := NewUserTodoListTaskLinkedResourceID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue", "linkedResourceIdValue").ID()
	expected := "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/linkedResources/linkedResourceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserTodoListTaskLinkedResourceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserTodoListTaskLinkedResourceId
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
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/linkedResources",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/linkedResources/linkedResourceIdValue",
			Expected: &UserTodoListTaskLinkedResourceId{
				UserId:           "userIdValue",
				TodoTaskListId:   "todoTaskListIdValue",
				TodoTaskId:       "todoTaskIdValue",
				LinkedResourceId: "linkedResourceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/linkedResources/linkedResourceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserTodoListTaskLinkedResourceID(v.Input)
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

		if actual.LinkedResourceId != v.Expected.LinkedResourceId {
			t.Fatalf("Expected %q but got %q for LinkedResourceId", v.Expected.LinkedResourceId, actual.LinkedResourceId)
		}

	}
}

func TestParseUserTodoListTaskLinkedResourceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserTodoListTaskLinkedResourceId
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
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/linkedResources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe/lInKeDrEsOuRcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/linkedResources/linkedResourceIdValue",
			Expected: &UserTodoListTaskLinkedResourceId{
				UserId:           "userIdValue",
				TodoTaskListId:   "todoTaskListIdValue",
				TodoTaskId:       "todoTaskIdValue",
				LinkedResourceId: "linkedResourceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/linkedResources/linkedResourceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe/lInKeDrEsOuRcEs/lInKeDrEsOuRcEiDvAlUe",
			Expected: &UserTodoListTaskLinkedResourceId{
				UserId:           "uSeRiDvAlUe",
				TodoTaskListId:   "tOdOtAsKlIsTiDvAlUe",
				TodoTaskId:       "tOdOtAsKiDvAlUe",
				LinkedResourceId: "lInKeDrEsOuRcEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe/lInKeDrEsOuRcEs/lInKeDrEsOuRcEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserTodoListTaskLinkedResourceIDInsensitively(v.Input)
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

		if actual.LinkedResourceId != v.Expected.LinkedResourceId {
			t.Fatalf("Expected %q but got %q for LinkedResourceId", v.Expected.LinkedResourceId, actual.LinkedResourceId)
		}

	}
}

func TestSegmentsForUserTodoListTaskLinkedResourceId(t *testing.T) {
	segments := UserTodoListTaskLinkedResourceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserTodoListTaskLinkedResourceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
