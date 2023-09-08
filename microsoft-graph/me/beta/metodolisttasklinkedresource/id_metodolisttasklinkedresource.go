package metodolisttasklinkedresource

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeTodoListTaskLinkedResourceId{}

// MeTodoListTaskLinkedResourceId is a struct representing the Resource ID for a Me Todo List Task Linked Resource
type MeTodoListTaskLinkedResourceId struct {
	TodoTaskListId   string
	TodoTaskId       string
	LinkedResourceId string
}

// NewMeTodoListTaskLinkedResourceID returns a new MeTodoListTaskLinkedResourceId struct
func NewMeTodoListTaskLinkedResourceID(todoTaskListId string, todoTaskId string, linkedResourceId string) MeTodoListTaskLinkedResourceId {
	return MeTodoListTaskLinkedResourceId{
		TodoTaskListId:   todoTaskListId,
		TodoTaskId:       todoTaskId,
		LinkedResourceId: linkedResourceId,
	}
}

// ParseMeTodoListTaskLinkedResourceID parses 'input' into a MeTodoListTaskLinkedResourceId
func ParseMeTodoListTaskLinkedResourceID(input string) (*MeTodoListTaskLinkedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeTodoListTaskLinkedResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeTodoListTaskLinkedResourceId{}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	if id.TodoTaskId, ok = parsed.Parsed["todoTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskId", *parsed)
	}

	if id.LinkedResourceId, ok = parsed.Parsed["linkedResourceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "linkedResourceId", *parsed)
	}

	return &id, nil
}

// ParseMeTodoListTaskLinkedResourceIDInsensitively parses 'input' case-insensitively into a MeTodoListTaskLinkedResourceId
// note: this method should only be used for API response data and not user input
func ParseMeTodoListTaskLinkedResourceIDInsensitively(input string) (*MeTodoListTaskLinkedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeTodoListTaskLinkedResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeTodoListTaskLinkedResourceId{}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	if id.TodoTaskId, ok = parsed.Parsed["todoTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskId", *parsed)
	}

	if id.LinkedResourceId, ok = parsed.Parsed["linkedResourceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "linkedResourceId", *parsed)
	}

	return &id, nil
}

// ValidateMeTodoListTaskLinkedResourceID checks that 'input' can be parsed as a Me Todo List Task Linked Resource ID
func ValidateMeTodoListTaskLinkedResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeTodoListTaskLinkedResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Todo List Task Linked Resource ID
func (id MeTodoListTaskLinkedResourceId) ID() string {
	fmtString := "/me/todo/lists/%s/tasks/%s/linkedResources/%s"
	return fmt.Sprintf(fmtString, id.TodoTaskListId, id.TodoTaskId, id.LinkedResourceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Todo List Task Linked Resource ID
func (id MeTodoListTaskLinkedResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticTodo", "todo", "todo"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskIdValue"),
		resourceids.StaticSegment("staticLinkedResources", "linkedResources", "linkedResources"),
		resourceids.UserSpecifiedSegment("linkedResourceId", "linkedResourceIdValue"),
	}
}

// String returns a human-readable description of this Me Todo List Task Linked Resource ID
func (id MeTodoListTaskLinkedResourceId) String() string {
	components := []string{
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
		fmt.Sprintf("Linked Resource: %q", id.LinkedResourceId),
	}
	return fmt.Sprintf("Me Todo List Task Linked Resource (%s)", strings.Join(components, "\n"))
}
