package usertodolisttasklinkedresource

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserTodoListTaskLinkedResourceId{}

// UserTodoListTaskLinkedResourceId is a struct representing the Resource ID for a User Todo List Task Linked Resource
type UserTodoListTaskLinkedResourceId struct {
	UserId           string
	TodoTaskListId   string
	TodoTaskId       string
	LinkedResourceId string
}

// NewUserTodoListTaskLinkedResourceID returns a new UserTodoListTaskLinkedResourceId struct
func NewUserTodoListTaskLinkedResourceID(userId string, todoTaskListId string, todoTaskId string, linkedResourceId string) UserTodoListTaskLinkedResourceId {
	return UserTodoListTaskLinkedResourceId{
		UserId:           userId,
		TodoTaskListId:   todoTaskListId,
		TodoTaskId:       todoTaskId,
		LinkedResourceId: linkedResourceId,
	}
}

// ParseUserTodoListTaskLinkedResourceID parses 'input' into a UserTodoListTaskLinkedResourceId
func ParseUserTodoListTaskLinkedResourceID(input string) (*UserTodoListTaskLinkedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTodoListTaskLinkedResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTodoListTaskLinkedResourceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ParseUserTodoListTaskLinkedResourceIDInsensitively parses 'input' case-insensitively into a UserTodoListTaskLinkedResourceId
// note: this method should only be used for API response data and not user input
func ParseUserTodoListTaskLinkedResourceIDInsensitively(input string) (*UserTodoListTaskLinkedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTodoListTaskLinkedResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTodoListTaskLinkedResourceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ValidateUserTodoListTaskLinkedResourceID checks that 'input' can be parsed as a User Todo List Task Linked Resource ID
func ValidateUserTodoListTaskLinkedResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserTodoListTaskLinkedResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Todo List Task Linked Resource ID
func (id UserTodoListTaskLinkedResourceId) ID() string {
	fmtString := "/users/%s/todo/lists/%s/tasks/%s/linkedResources/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TodoTaskListId, id.TodoTaskId, id.LinkedResourceId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Todo List Task Linked Resource ID
func (id UserTodoListTaskLinkedResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticTodo", "todo", "todo"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskIdValue"),
		resourceids.StaticSegment("staticLinkedResources", "linkedResources", "linkedResources"),
		resourceids.UserSpecifiedSegment("linkedResourceId", "linkedResourceIdValue"),
	}
}

// String returns a human-readable description of this User Todo List Task Linked Resource ID
func (id UserTodoListTaskLinkedResourceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
		fmt.Sprintf("Linked Resource: %q", id.LinkedResourceId),
	}
	return fmt.Sprintf("User Todo List Task Linked Resource (%s)", strings.Join(components, "\n"))
}
