package usertodolisttaskextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserTodoListTaskExtensionId{}

// UserTodoListTaskExtensionId is a struct representing the Resource ID for a User Todo List Task Extension
type UserTodoListTaskExtensionId struct {
	UserId         string
	TodoTaskListId string
	TodoTaskId     string
	ExtensionId    string
}

// NewUserTodoListTaskExtensionID returns a new UserTodoListTaskExtensionId struct
func NewUserTodoListTaskExtensionID(userId string, todoTaskListId string, todoTaskId string, extensionId string) UserTodoListTaskExtensionId {
	return UserTodoListTaskExtensionId{
		UserId:         userId,
		TodoTaskListId: todoTaskListId,
		TodoTaskId:     todoTaskId,
		ExtensionId:    extensionId,
	}
}

// ParseUserTodoListTaskExtensionID parses 'input' into a UserTodoListTaskExtensionId
func ParseUserTodoListTaskExtensionID(input string) (*UserTodoListTaskExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTodoListTaskExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTodoListTaskExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ParseUserTodoListTaskExtensionIDInsensitively parses 'input' case-insensitively into a UserTodoListTaskExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserTodoListTaskExtensionIDInsensitively(input string) (*UserTodoListTaskExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTodoListTaskExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTodoListTaskExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ValidateUserTodoListTaskExtensionID checks that 'input' can be parsed as a User Todo List Task Extension ID
func ValidateUserTodoListTaskExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserTodoListTaskExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Todo List Task Extension ID
func (id UserTodoListTaskExtensionId) ID() string {
	fmtString := "/users/%s/todo/lists/%s/tasks/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TodoTaskListId, id.TodoTaskId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Todo List Task Extension ID
func (id UserTodoListTaskExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticTodo", "todo", "todo"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Todo List Task Extension ID
func (id UserTodoListTaskExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Todo List Task Extension (%s)", strings.Join(components, "\n"))
}
