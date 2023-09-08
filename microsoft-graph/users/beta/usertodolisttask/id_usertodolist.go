package usertodolisttask

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserTodoListId{}

// UserTodoListId is a struct representing the Resource ID for a User Todo List
type UserTodoListId struct {
	UserId         string
	TodoTaskListId string
}

// NewUserTodoListID returns a new UserTodoListId struct
func NewUserTodoListID(userId string, todoTaskListId string) UserTodoListId {
	return UserTodoListId{
		UserId:         userId,
		TodoTaskListId: todoTaskListId,
	}
}

// ParseUserTodoListID parses 'input' into a UserTodoListId
func ParseUserTodoListID(input string) (*UserTodoListId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTodoListId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTodoListId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	return &id, nil
}

// ParseUserTodoListIDInsensitively parses 'input' case-insensitively into a UserTodoListId
// note: this method should only be used for API response data and not user input
func ParseUserTodoListIDInsensitively(input string) (*UserTodoListId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTodoListId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTodoListId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	return &id, nil
}

// ValidateUserTodoListID checks that 'input' can be parsed as a User Todo List ID
func ValidateUserTodoListID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserTodoListID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Todo List ID
func (id UserTodoListId) ID() string {
	fmtString := "/users/%s/todo/lists/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TodoTaskListId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Todo List ID
func (id UserTodoListId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticTodo", "todo", "todo"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListIdValue"),
	}
}

// String returns a human-readable description of this User Todo List ID
func (id UserTodoListId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
	}
	return fmt.Sprintf("User Todo List (%s)", strings.Join(components, "\n"))
}
