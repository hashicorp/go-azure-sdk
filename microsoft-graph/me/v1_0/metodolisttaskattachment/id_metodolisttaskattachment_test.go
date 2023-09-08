package metodolisttaskattachment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeTodoListTaskAttachmentId{}

func TestNewMeTodoListTaskAttachmentID(t *testing.T) {
	id := NewMeTodoListTaskAttachmentID("todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue")

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

func TestFormatMeTodoListTaskAttachmentID(t *testing.T) {
	actual := NewMeTodoListTaskAttachmentID("todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue").ID()
	expected := "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachments/attachmentBaseIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeTodoListTaskAttachmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeTodoListTaskAttachmentId
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
			Input: "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachments/attachmentBaseIdValue",
			Expected: &MeTodoListTaskAttachmentId{
				TodoTaskListId:   "todoTaskListIdValue",
				TodoTaskId:       "todoTaskIdValue",
				AttachmentBaseId: "attachmentBaseIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachments/attachmentBaseIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeTodoListTaskAttachmentID(v.Input)
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

		if actual.AttachmentBaseId != v.Expected.AttachmentBaseId {
			t.Fatalf("Expected %q but got %q for AttachmentBaseId", v.Expected.AttachmentBaseId, actual.AttachmentBaseId)
		}

	}
}

func TestParseMeTodoListTaskAttachmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeTodoListTaskAttachmentId
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
			Input: "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe/aTtAcHmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachments/attachmentBaseIdValue",
			Expected: &MeTodoListTaskAttachmentId{
				TodoTaskListId:   "todoTaskListIdValue",
				TodoTaskId:       "todoTaskIdValue",
				AttachmentBaseId: "attachmentBaseIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/todo/lists/todoTaskListIdValue/tasks/todoTaskIdValue/attachments/attachmentBaseIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe/aTtAcHmEnTs/aTtAcHmEnTbAsEiDvAlUe",
			Expected: &MeTodoListTaskAttachmentId{
				TodoTaskListId:   "tOdOtAsKlIsTiDvAlUe",
				TodoTaskId:       "tOdOtAsKiDvAlUe",
				AttachmentBaseId: "aTtAcHmEnTbAsEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/tAsKs/tOdOtAsKiDvAlUe/aTtAcHmEnTs/aTtAcHmEnTbAsEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeTodoListTaskAttachmentIDInsensitively(v.Input)
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

		if actual.AttachmentBaseId != v.Expected.AttachmentBaseId {
			t.Fatalf("Expected %q but got %q for AttachmentBaseId", v.Expected.AttachmentBaseId, actual.AttachmentBaseId)
		}

	}
}

func TestSegmentsForMeTodoListTaskAttachmentId(t *testing.T) {
	segments := MeTodoListTaskAttachmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeTodoListTaskAttachmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
