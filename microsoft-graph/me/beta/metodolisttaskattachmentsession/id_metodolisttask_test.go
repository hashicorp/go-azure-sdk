package metodolisttaskattachmentsession

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeTodoListTaskId{}

func TestNewMeTodoListTaskID(t *testing.T) {
	id := NewMeTodoListTaskID("todoTaskListIdValue", "todoTaskIdValue")

	if id.TodoTaskListId != "todoTaskListIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TodoTaskListId'", id.TodoTaskListId, "todoTaskListIdValue")
	}

	if id.TodoTaskId != "todoTaskIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TodoTaskId'", id.TodoTaskId, "todoTaskIdValue")
	}
}

func TestFormatMeTodoListTaskID(t *testing.T) {
	actual := NewMeTodoListTaskID("todoTaskListIdValue", "todoTaskIdValue").ID()
	expected := "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeTodoListTaskID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeTodoListTaskId
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
			// Valid URI
			Input: "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue",
			Expected: &MeTodoListTaskId{
				TodoTaskListId: "todoTaskListIdValue",
				TodoTaskId:     "todoTaskIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeTodoListTaskID(v.Input)
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

	}
}

func TestParseMeTodoListTaskIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeTodoListTaskId
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
			// Valid URI
			Input: "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue",
			Expected: &MeTodoListTaskId{
				TodoTaskListId: "todoTaskListIdValue",
				TodoTaskId:     "todoTaskIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe",
			Expected: &MeTodoListTaskId{
				TodoTaskListId: "tOdOtAsKlIsTiDvAlUe",
				TodoTaskId:     "tOdOtAsKiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeTodoListTaskIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForMeTodoListTaskId(t *testing.T) {
	segments := MeTodoListTaskId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeTodoListTaskId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
