package usertodolistextension

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserTodoListExtensionId{}

func TestNewUserTodoListExtensionID(t *testing.T) {
	id := NewUserTodoListExtensionID("userIdValue", "todoTaskListIdValue", "extensionIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.TodoTaskListId != "todoTaskListIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TodoTaskListId'", id.TodoTaskListId, "todoTaskListIdValue")
	}

	if id.ExtensionId != "extensionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionId'", id.ExtensionId, "extensionIdValue")
	}
}

func TestFormatUserTodoListExtensionID(t *testing.T) {
	actual := NewUserTodoListExtensionID("userIdValue", "todoTaskListIdValue", "extensionIdValue").ID()
	expected := "/users/userIdValue/todo/lists/todoTaskListIdValue/extensions/extensionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserTodoListExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserTodoListExtensionId
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
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/extensions/extensionIdValue",
			Expected: &UserTodoListExtensionId{
				UserId:         "userIdValue",
				TodoTaskListId: "todoTaskListIdValue",
				ExtensionId:    "extensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/extensions/extensionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserTodoListExtensionID(v.Input)
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

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestParseUserTodoListExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserTodoListExtensionId
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
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/extensions/extensionIdValue",
			Expected: &UserTodoListExtensionId{
				UserId:         "userIdValue",
				TodoTaskListId: "todoTaskListIdValue",
				ExtensionId:    "extensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/todo/lists/todoTaskListIdValue/extensions/extensionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/eXtEnSiOnS/eXtEnSiOnIdVaLuE",
			Expected: &UserTodoListExtensionId{
				UserId:         "uSeRiDvAlUe",
				TodoTaskListId: "tOdOtAsKlIsTiDvAlUe",
				ExtensionId:    "eXtEnSiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tOdO/lIsTs/tOdOtAsKlIsTiDvAlUe/eXtEnSiOnS/eXtEnSiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserTodoListExtensionIDInsensitively(v.Input)
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

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestSegmentsForUserTodoListExtensionId(t *testing.T) {
	segments := UserTodoListExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserTodoListExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
