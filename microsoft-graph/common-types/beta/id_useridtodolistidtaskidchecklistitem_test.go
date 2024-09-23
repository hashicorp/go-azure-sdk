package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdTodoListIdTaskIdChecklistItemId{}

func TestNewUserIdTodoListIdTaskIdChecklistItemID(t *testing.T) {
	id := NewUserIdTodoListIdTaskIdChecklistItemID("userId", "todoTaskListId", "todoTaskId", "checklistItemId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.TodoTaskListId != "todoTaskListId" {
		t.Fatalf("Expected %q but got %q for Segment 'TodoTaskListId'", id.TodoTaskListId, "todoTaskListId")
	}

	if id.TodoTaskId != "todoTaskId" {
		t.Fatalf("Expected %q but got %q for Segment 'TodoTaskId'", id.TodoTaskId, "todoTaskId")
	}

	if id.ChecklistItemId != "checklistItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChecklistItemId'", id.ChecklistItemId, "checklistItemId")
	}
}

func TestFormatUserIdTodoListIdTaskIdChecklistItemID(t *testing.T) {
	actual := NewUserIdTodoListIdTaskIdChecklistItemID("userId", "todoTaskListId", "todoTaskId", "checklistItemId").ID()
	expected := "/users/userId/todo/lists/todoTaskListId/tasks/todoTaskId/checklistItems/checklistItemId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdTodoListIdTaskIdChecklistItemID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdTodoListIdTaskIdChecklistItemId
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
			Input: "/users/userId/todo/lists/todoTaskListId/tasks/todoTaskId/checklistItems",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/todo/lists/todoTaskListId/tasks/todoTaskId/checklistItems/checklistItemId",
			Expected: &UserIdTodoListIdTaskIdChecklistItemId{
				UserId:          "userId",
				TodoTaskListId:  "todoTaskListId",
				TodoTaskId:      "todoTaskId",
				ChecklistItemId: "checklistItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/todo/lists/todoTaskListId/tasks/todoTaskId/checklistItems/checklistItemId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdTodoListIdTaskIdChecklistItemID(v.Input)
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

		if actual.ChecklistItemId != v.Expected.ChecklistItemId {
			t.Fatalf("Expected %q but got %q for ChecklistItemId", v.Expected.ChecklistItemId, actual.ChecklistItemId)
		}

	}
}

func TestParseUserIdTodoListIdTaskIdChecklistItemIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdTodoListIdTaskIdChecklistItemId
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
			Input: "/users/userId/todo/lists/todoTaskListId/tasks/todoTaskId/checklistItems",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/tOdO/lIsTs/tOdOtAsKlIsTiD/tAsKs/tOdOtAsKiD/cHeCkLiStItEmS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/todo/lists/todoTaskListId/tasks/todoTaskId/checklistItems/checklistItemId",
			Expected: &UserIdTodoListIdTaskIdChecklistItemId{
				UserId:          "userId",
				TodoTaskListId:  "todoTaskListId",
				TodoTaskId:      "todoTaskId",
				ChecklistItemId: "checklistItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/todo/lists/todoTaskListId/tasks/todoTaskId/checklistItems/checklistItemId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/tOdO/lIsTs/tOdOtAsKlIsTiD/tAsKs/tOdOtAsKiD/cHeCkLiStItEmS/cHeCkLiStItEmId",
			Expected: &UserIdTodoListIdTaskIdChecklistItemId{
				UserId:          "uSeRiD",
				TodoTaskListId:  "tOdOtAsKlIsTiD",
				TodoTaskId:      "tOdOtAsKiD",
				ChecklistItemId: "cHeCkLiStItEmId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/tOdO/lIsTs/tOdOtAsKlIsTiD/tAsKs/tOdOtAsKiD/cHeCkLiStItEmS/cHeCkLiStItEmId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdTodoListIdTaskIdChecklistItemIDInsensitively(v.Input)
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

		if actual.ChecklistItemId != v.Expected.ChecklistItemId {
			t.Fatalf("Expected %q but got %q for ChecklistItemId", v.Expected.ChecklistItemId, actual.ChecklistItemId)
		}

	}
}

func TestSegmentsForUserIdTodoListIdTaskIdChecklistItemId(t *testing.T) {
	segments := UserIdTodoListIdTaskIdChecklistItemId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdTodoListIdTaskIdChecklistItemId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
