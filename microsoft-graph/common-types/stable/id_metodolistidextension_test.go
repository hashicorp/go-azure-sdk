package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeTodoListIdExtensionId{}

func TestNewMeTodoListIdExtensionID(t *testing.T) {
	id := NewMeTodoListIdExtensionID("todoTaskListId", "extensionId")

	if id.TodoTaskListId != "todoTaskListId" {
		t.Fatalf("Expected %q but got %q for Segment 'TodoTaskListId'", id.TodoTaskListId, "todoTaskListId")
	}

	if id.ExtensionId != "extensionId" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionId'", id.ExtensionId, "extensionId")
	}
}

func TestFormatMeTodoListIdExtensionID(t *testing.T) {
	actual := NewMeTodoListIdExtensionID("todoTaskListId", "extensionId").ID()
	expected := "/me/todo/lists/todoTaskListId/extensions/extensionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeTodoListIdExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeTodoListIdExtensionId
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
			Input: "/me/todo/lists/todoTaskListId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/todo/lists/todoTaskListId/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/todo/lists/todoTaskListId/extensions/extensionId",
			Expected: &MeTodoListIdExtensionId{
				TodoTaskListId: "todoTaskListId",
				ExtensionId:    "extensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/todo/lists/todoTaskListId/extensions/extensionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeTodoListIdExtensionID(v.Input)
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

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestParseMeTodoListIdExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeTodoListIdExtensionId
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
			Input: "/me/todo/lists/todoTaskListId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/tOdO/lIsTs/tOdOtAsKlIsTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/todo/lists/todoTaskListId/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/tOdO/lIsTs/tOdOtAsKlIsTiD/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/todo/lists/todoTaskListId/extensions/extensionId",
			Expected: &MeTodoListIdExtensionId{
				TodoTaskListId: "todoTaskListId",
				ExtensionId:    "extensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/todo/lists/todoTaskListId/extensions/extensionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/tOdO/lIsTs/tOdOtAsKlIsTiD/eXtEnSiOnS/eXtEnSiOnId",
			Expected: &MeTodoListIdExtensionId{
				TodoTaskListId: "tOdOtAsKlIsTiD",
				ExtensionId:    "eXtEnSiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/tOdO/lIsTs/tOdOtAsKlIsTiD/eXtEnSiOnS/eXtEnSiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeTodoListIdExtensionIDInsensitively(v.Input)
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

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestSegmentsForMeTodoListIdExtensionId(t *testing.T) {
	segments := MeTodoListIdExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeTodoListIdExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
