package metodolisttask

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeTodoListId{}

// MeTodoListId is a struct representing the Resource ID for a Me Todo List
type MeTodoListId struct {
	TodoTaskListId string
}

// NewMeTodoListID returns a new MeTodoListId struct
func NewMeTodoListID(todoTaskListId string) MeTodoListId {
	return MeTodoListId{
		TodoTaskListId: todoTaskListId,
	}
}

// ParseMeTodoListID parses 'input' into a MeTodoListId
func ParseMeTodoListID(input string) (*MeTodoListId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeTodoListId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeTodoListId{}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	return &id, nil
}

// ParseMeTodoListIDInsensitively parses 'input' case-insensitively into a MeTodoListId
// note: this method should only be used for API response data and not user input
func ParseMeTodoListIDInsensitively(input string) (*MeTodoListId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeTodoListId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeTodoListId{}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	return &id, nil
}

// ValidateMeTodoListID checks that 'input' can be parsed as a Me Todo List ID
func ValidateMeTodoListID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeTodoListID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Todo List ID
func (id MeTodoListId) ID() string {
	fmtString := "/me/todo/lists/%s"
	return fmt.Sprintf(fmtString, id.TodoTaskListId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Todo List ID
func (id MeTodoListId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticTodo", "todo", "todo"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListIdValue"),
	}
}

// String returns a human-readable description of this Me Todo List ID
func (id MeTodoListId) String() string {
	components := []string{
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
	}
	return fmt.Sprintf("Me Todo List (%s)", strings.Join(components, "\n"))
}
