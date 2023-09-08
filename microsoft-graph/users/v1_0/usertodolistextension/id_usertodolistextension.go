package usertodolistextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserTodoListExtensionId{}

// UserTodoListExtensionId is a struct representing the Resource ID for a User Todo List Extension
type UserTodoListExtensionId struct {
	UserId         string
	TodoTaskListId string
	ExtensionId    string
}

// NewUserTodoListExtensionID returns a new UserTodoListExtensionId struct
func NewUserTodoListExtensionID(userId string, todoTaskListId string, extensionId string) UserTodoListExtensionId {
	return UserTodoListExtensionId{
		UserId:         userId,
		TodoTaskListId: todoTaskListId,
		ExtensionId:    extensionId,
	}
}

// ParseUserTodoListExtensionID parses 'input' into a UserTodoListExtensionId
func ParseUserTodoListExtensionID(input string) (*UserTodoListExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTodoListExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTodoListExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseUserTodoListExtensionIDInsensitively parses 'input' case-insensitively into a UserTodoListExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserTodoListExtensionIDInsensitively(input string) (*UserTodoListExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTodoListExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTodoListExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateUserTodoListExtensionID checks that 'input' can be parsed as a User Todo List Extension ID
func ValidateUserTodoListExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserTodoListExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Todo List Extension ID
func (id UserTodoListExtensionId) ID() string {
	fmtString := "/users/%s/todo/lists/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TodoTaskListId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Todo List Extension ID
func (id UserTodoListExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticTodo", "todo", "todo"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Todo List Extension ID
func (id UserTodoListExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Todo List Extension (%s)", strings.Join(components, "\n"))
}
