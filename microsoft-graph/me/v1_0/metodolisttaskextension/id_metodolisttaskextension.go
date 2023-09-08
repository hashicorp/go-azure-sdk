package metodolisttaskextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeTodoListTaskExtensionId{}

// MeTodoListTaskExtensionId is a struct representing the Resource ID for a Me Todo List Task Extension
type MeTodoListTaskExtensionId struct {
	TodoTaskListId string
	TodoTaskId     string
	ExtensionId    string
}

// NewMeTodoListTaskExtensionID returns a new MeTodoListTaskExtensionId struct
func NewMeTodoListTaskExtensionID(todoTaskListId string, todoTaskId string, extensionId string) MeTodoListTaskExtensionId {
	return MeTodoListTaskExtensionId{
		TodoTaskListId: todoTaskListId,
		TodoTaskId:     todoTaskId,
		ExtensionId:    extensionId,
	}
}

// ParseMeTodoListTaskExtensionID parses 'input' into a MeTodoListTaskExtensionId
func ParseMeTodoListTaskExtensionID(input string) (*MeTodoListTaskExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeTodoListTaskExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeTodoListTaskExtensionId{}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	if id.TodoTaskId, ok = parsed.Parsed["todoTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseMeTodoListTaskExtensionIDInsensitively parses 'input' case-insensitively into a MeTodoListTaskExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeTodoListTaskExtensionIDInsensitively(input string) (*MeTodoListTaskExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeTodoListTaskExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeTodoListTaskExtensionId{}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	if id.TodoTaskId, ok = parsed.Parsed["todoTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateMeTodoListTaskExtensionID checks that 'input' can be parsed as a Me Todo List Task Extension ID
func ValidateMeTodoListTaskExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeTodoListTaskExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Todo List Task Extension ID
func (id MeTodoListTaskExtensionId) ID() string {
	fmtString := "/me/todo/lists/%s/tasks/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.TodoTaskListId, id.TodoTaskId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Todo List Task Extension ID
func (id MeTodoListTaskExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticTodo", "todo", "todo"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Todo List Task Extension ID
func (id MeTodoListTaskExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Todo List Task Extension (%s)", strings.Join(components, "\n"))
}
