package metodolisttaskchecklistitem

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeTodoListTaskChecklistItemId{}

func TestNewMeTodoListTaskChecklistItemID(t *testing.T) {
	id := NewMeTodoListTaskChecklistItemID("todoTaskListIdValue", "todoTaskIdValue", "checklistItemIdValue")

	if id.TodoTaskListId != "todoTaskListIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TodoTaskListId'", id.TodoTaskListId, "todoTaskListIdValue")
	}

	if id.TodoTaskId != "todoTaskIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TodoTaskId'", id.TodoTaskId, "todoTaskIdValue")
	}

	if id.ChecklistItemId != "checklistItemIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChecklistItemId'", id.ChecklistItemId, "checklistItemIdValue")
	}
}

func TestFormatMeTodoListTaskChecklistItemID(t *testing.T) {
	actual := NewMeTodoListTaskChecklistItemID("todoTaskListIdValue", "todoTaskIdValue", "checklistItemIdValue").ID()
	expected := "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/checklistItems/checklistItemIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeTodoListTaskChecklistItemID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeTodoListTaskChecklistItemId
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
			Input: "/me/todo",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/todo/lists",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/todo/lists/todoTaskListIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/todo/lists/todoTaskListIdValue/tasks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/checklistItems",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/checklistItems/checklistItemIdValue",
			Expected: &MeTodoListTaskChecklistItemId{
				TodoTaskListId:  "todoTaskListIdValue",
				TodoTaskId:      "todoTaskIdValue",
				ChecklistItemId: "checklistItemIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/checklistItems/checklistItemIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeTodoListTaskChecklistItemID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
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

func TestParseMeTodoListTaskChecklistItemIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeTodoListTaskChecklistItemId
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
			Input: "/me/todo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/tOdO",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/todo/lists",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/tOdO/lIsTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/todo/lists/todoTaskListIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/todo/lists/todoTaskListIdValue/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/checklistItems",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe/cHeCkLiStItEmS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/checklistItems/checklistItemIdValue",
			Expected: &MeTodoListTaskChecklistItemId{
				TodoTaskListId:  "todoTaskListIdValue",
				TodoTaskId:      "todoTaskIdValue",
				ChecklistItemId: "checklistItemIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/checklistItems/checklistItemIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe/cHeCkLiStItEmS/cHeCkLiStItEmIdVaLuE",
			Expected: &MeTodoListTaskChecklistItemId{
				TodoTaskListId:  "tOdOtAsKlIsTiDvAlUe",
				TodoTaskId:      "tOdOtAsKiDvAlUe",
				ChecklistItemId: "cHeCkLiStItEmIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe/cHeCkLiStItEmS/cHeCkLiStItEmIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeTodoListTaskChecklistItemIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
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

func TestSegmentsForMeTodoListTaskChecklistItemId(t *testing.T) {
	segments := MeTodoListTaskChecklistItemId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeTodoListTaskChecklistItemId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
