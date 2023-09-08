package metodolisttaskattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeTodoListTaskId{}

// MeTodoListTaskId is a struct representing the Resource ID for a Me Todo List Task
type MeTodoListTaskId struct {
	TodoTaskListId string
	TodoTaskId     string
}

// NewMeTodoListTaskID returns a new MeTodoListTaskId struct
func NewMeTodoListTaskID(todoTaskListId string, todoTaskId string) MeTodoListTaskId {
	return MeTodoListTaskId{
		TodoTaskListId: todoTaskListId,
		TodoTaskId:     todoTaskId,
	}
}

// ParseMeTodoListTaskID parses 'input' into a MeTodoListTaskId
func ParseMeTodoListTaskID(input string) (*MeTodoListTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeTodoListTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeTodoListTaskId{}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	if id.TodoTaskId, ok = parsed.Parsed["todoTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskId", *parsed)
	}

	return &id, nil
}

// ParseMeTodoListTaskIDInsensitively parses 'input' case-insensitively into a MeTodoListTaskId
// note: this method should only be used for API response data and not user input
func ParseMeTodoListTaskIDInsensitively(input string) (*MeTodoListTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeTodoListTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeTodoListTaskId{}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	if id.TodoTaskId, ok = parsed.Parsed["todoTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskId", *parsed)
	}

	return &id, nil
}

// ValidateMeTodoListTaskID checks that 'input' can be parsed as a Me Todo List Task ID
func ValidateMeTodoListTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeTodoListTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Todo List Task ID
func (id MeTodoListTaskId) ID() string {
	fmtString := "/me/todo/lists/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.TodoTaskListId, id.TodoTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Todo List Task ID
func (id MeTodoListTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticTodo", "todo", "todo"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskIdValue"),
	}
}

// String returns a human-readable description of this Me Todo List Task ID
func (id MeTodoListTaskId) String() string {
	components := []string{
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
	}
	return fmt.Sprintf("Me Todo List Task (%s)", strings.Join(components, "\n"))
}
