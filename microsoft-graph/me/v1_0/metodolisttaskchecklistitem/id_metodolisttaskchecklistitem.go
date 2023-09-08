package metodolisttaskchecklistitem

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeTodoListTaskChecklistItemId{}

// MeTodoListTaskChecklistItemId is a struct representing the Resource ID for a Me Todo List Task Checklist Item
type MeTodoListTaskChecklistItemId struct {
	TodoTaskListId  string
	TodoTaskId      string
	ChecklistItemId string
}

// NewMeTodoListTaskChecklistItemID returns a new MeTodoListTaskChecklistItemId struct
func NewMeTodoListTaskChecklistItemID(todoTaskListId string, todoTaskId string, checklistItemId string) MeTodoListTaskChecklistItemId {
	return MeTodoListTaskChecklistItemId{
		TodoTaskListId:  todoTaskListId,
		TodoTaskId:      todoTaskId,
		ChecklistItemId: checklistItemId,
	}
}

// ParseMeTodoListTaskChecklistItemID parses 'input' into a MeTodoListTaskChecklistItemId
func ParseMeTodoListTaskChecklistItemID(input string) (*MeTodoListTaskChecklistItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeTodoListTaskChecklistItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeTodoListTaskChecklistItemId{}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	if id.TodoTaskId, ok = parsed.Parsed["todoTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskId", *parsed)
	}

	if id.ChecklistItemId, ok = parsed.Parsed["checklistItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "checklistItemId", *parsed)
	}

	return &id, nil
}

// ParseMeTodoListTaskChecklistItemIDInsensitively parses 'input' case-insensitively into a MeTodoListTaskChecklistItemId
// note: this method should only be used for API response data and not user input
func ParseMeTodoListTaskChecklistItemIDInsensitively(input string) (*MeTodoListTaskChecklistItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeTodoListTaskChecklistItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeTodoListTaskChecklistItemId{}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	if id.TodoTaskId, ok = parsed.Parsed["todoTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskId", *parsed)
	}

	if id.ChecklistItemId, ok = parsed.Parsed["checklistItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "checklistItemId", *parsed)
	}

	return &id, nil
}

// ValidateMeTodoListTaskChecklistItemID checks that 'input' can be parsed as a Me Todo List Task Checklist Item ID
func ValidateMeTodoListTaskChecklistItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeTodoListTaskChecklistItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Todo List Task Checklist Item ID
func (id MeTodoListTaskChecklistItemId) ID() string {
	fmtString := "/me/todo/lists/%s/tasks/%s/checklistItems/%s"
	return fmt.Sprintf(fmtString, id.TodoTaskListId, id.TodoTaskId, id.ChecklistItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Todo List Task Checklist Item ID
func (id MeTodoListTaskChecklistItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticTodo", "todo", "todo"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskIdValue"),
		resourceids.StaticSegment("staticChecklistItems", "checklistItems", "checklistItems"),
		resourceids.UserSpecifiedSegment("checklistItemId", "checklistItemIdValue"),
	}
}

// String returns a human-readable description of this Me Todo List Task Checklist Item ID
func (id MeTodoListTaskChecklistItemId) String() string {
	components := []string{
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
		fmt.Sprintf("Checklist Item: %q", id.ChecklistItemId),
	}
	return fmt.Sprintf("Me Todo List Task Checklist Item (%s)", strings.Join(components, "\n"))
}
