package usertodolisttaskchecklistitem

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserTodoListTaskChecklistItemId{}

// UserTodoListTaskChecklistItemId is a struct representing the Resource ID for a User Todo List Task Checklist Item
type UserTodoListTaskChecklistItemId struct {
	UserId          string
	TodoTaskListId  string
	TodoTaskId      string
	ChecklistItemId string
}

// NewUserTodoListTaskChecklistItemID returns a new UserTodoListTaskChecklistItemId struct
func NewUserTodoListTaskChecklistItemID(userId string, todoTaskListId string, todoTaskId string, checklistItemId string) UserTodoListTaskChecklistItemId {
	return UserTodoListTaskChecklistItemId{
		UserId:          userId,
		TodoTaskListId:  todoTaskListId,
		TodoTaskId:      todoTaskId,
		ChecklistItemId: checklistItemId,
	}
}

// ParseUserTodoListTaskChecklistItemID parses 'input' into a UserTodoListTaskChecklistItemId
func ParseUserTodoListTaskChecklistItemID(input string) (*UserTodoListTaskChecklistItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTodoListTaskChecklistItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTodoListTaskChecklistItemId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ParseUserTodoListTaskChecklistItemIDInsensitively parses 'input' case-insensitively into a UserTodoListTaskChecklistItemId
// note: this method should only be used for API response data and not user input
func ParseUserTodoListTaskChecklistItemIDInsensitively(input string) (*UserTodoListTaskChecklistItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTodoListTaskChecklistItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTodoListTaskChecklistItemId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ValidateUserTodoListTaskChecklistItemID checks that 'input' can be parsed as a User Todo List Task Checklist Item ID
func ValidateUserTodoListTaskChecklistItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserTodoListTaskChecklistItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Todo List Task Checklist Item ID
func (id UserTodoListTaskChecklistItemId) ID() string {
	fmtString := "/users/%s/todo/lists/%s/tasks/%s/checklistItems/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TodoTaskListId, id.TodoTaskId, id.ChecklistItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Todo List Task Checklist Item ID
func (id UserTodoListTaskChecklistItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticTodo", "todo", "todo"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskIdValue"),
		resourceids.StaticSegment("staticChecklistItems", "checklistItems", "checklistItems"),
		resourceids.UserSpecifiedSegment("checklistItemId", "checklistItemIdValue"),
	}
}

// String returns a human-readable description of this User Todo List Task Checklist Item ID
func (id UserTodoListTaskChecklistItemId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
		fmt.Sprintf("Checklist Item: %q", id.ChecklistItemId),
	}
	return fmt.Sprintf("User Todo List Task Checklist Item (%s)", strings.Join(components, "\n"))
}
