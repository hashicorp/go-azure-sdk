package usertodolisttask

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserTodoListTaskId{}

// UserTodoListTaskId is a struct representing the Resource ID for a User Todo List Task
type UserTodoListTaskId struct {
	UserId         string
	TodoTaskListId string
	TodoTaskId     string
}

// NewUserTodoListTaskID returns a new UserTodoListTaskId struct
func NewUserTodoListTaskID(userId string, todoTaskListId string, todoTaskId string) UserTodoListTaskId {
	return UserTodoListTaskId{
		UserId:         userId,
		TodoTaskListId: todoTaskListId,
		TodoTaskId:     todoTaskId,
	}
}

// ParseUserTodoListTaskID parses 'input' into a UserTodoListTaskId
func ParseUserTodoListTaskID(input string) (*UserTodoListTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTodoListTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTodoListTaskId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	if id.TodoTaskId, ok = parsed.Parsed["todoTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskId", *parsed)
	}

	return &id, nil
}

// ParseUserTodoListTaskIDInsensitively parses 'input' case-insensitively into a UserTodoListTaskId
// note: this method should only be used for API response data and not user input
func ParseUserTodoListTaskIDInsensitively(input string) (*UserTodoListTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTodoListTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTodoListTaskId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	if id.TodoTaskId, ok = parsed.Parsed["todoTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskId", *parsed)
	}

	return &id, nil
}

// ValidateUserTodoListTaskID checks that 'input' can be parsed as a User Todo List Task ID
func ValidateUserTodoListTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserTodoListTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Todo List Task ID
func (id UserTodoListTaskId) ID() string {
	fmtString := "/users/%s/todo/lists/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TodoTaskListId, id.TodoTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Todo List Task ID
func (id UserTodoListTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticTodo", "todo", "todo"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskIdValue"),
	}
}

// String returns a human-readable description of this User Todo List Task ID
func (id UserTodoListTaskId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
	}
	return fmt.Sprintf("User Todo List Task (%s)", strings.Join(components, "\n"))
}
