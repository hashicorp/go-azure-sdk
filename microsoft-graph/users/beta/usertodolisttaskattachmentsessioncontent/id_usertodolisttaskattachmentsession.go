package usertodolisttaskattachmentsessioncontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserTodoListTaskAttachmentSessionId{}

// UserTodoListTaskAttachmentSessionId is a struct representing the Resource ID for a User Todo List Task Attachment Session
type UserTodoListTaskAttachmentSessionId struct {
	UserId              string
	TodoTaskListId      string
	TodoTaskId          string
	AttachmentSessionId string
}

// NewUserTodoListTaskAttachmentSessionID returns a new UserTodoListTaskAttachmentSessionId struct
func NewUserTodoListTaskAttachmentSessionID(userId string, todoTaskListId string, todoTaskId string, attachmentSessionId string) UserTodoListTaskAttachmentSessionId {
	return UserTodoListTaskAttachmentSessionId{
		UserId:              userId,
		TodoTaskListId:      todoTaskListId,
		TodoTaskId:          todoTaskId,
		AttachmentSessionId: attachmentSessionId,
	}
}

// ParseUserTodoListTaskAttachmentSessionID parses 'input' into a UserTodoListTaskAttachmentSessionId
func ParseUserTodoListTaskAttachmentSessionID(input string) (*UserTodoListTaskAttachmentSessionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTodoListTaskAttachmentSessionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTodoListTaskAttachmentSessionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	if id.TodoTaskId, ok = parsed.Parsed["todoTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskId", *parsed)
	}

	if id.AttachmentSessionId, ok = parsed.Parsed["attachmentSessionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentSessionId", *parsed)
	}

	return &id, nil
}

// ParseUserTodoListTaskAttachmentSessionIDInsensitively parses 'input' case-insensitively into a UserTodoListTaskAttachmentSessionId
// note: this method should only be used for API response data and not user input
func ParseUserTodoListTaskAttachmentSessionIDInsensitively(input string) (*UserTodoListTaskAttachmentSessionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTodoListTaskAttachmentSessionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTodoListTaskAttachmentSessionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	if id.TodoTaskId, ok = parsed.Parsed["todoTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskId", *parsed)
	}

	if id.AttachmentSessionId, ok = parsed.Parsed["attachmentSessionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentSessionId", *parsed)
	}

	return &id, nil
}

// ValidateUserTodoListTaskAttachmentSessionID checks that 'input' can be parsed as a User Todo List Task Attachment Session ID
func ValidateUserTodoListTaskAttachmentSessionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserTodoListTaskAttachmentSessionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Todo List Task Attachment Session ID
func (id UserTodoListTaskAttachmentSessionId) ID() string {
	fmtString := "/users/%s/todo/lists/%s/tasks/%s/attachmentSessions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TodoTaskListId, id.TodoTaskId, id.AttachmentSessionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Todo List Task Attachment Session ID
func (id UserTodoListTaskAttachmentSessionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticTodo", "todo", "todo"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskIdValue"),
		resourceids.StaticSegment("staticAttachmentSessions", "attachmentSessions", "attachmentSessions"),
		resourceids.UserSpecifiedSegment("attachmentSessionId", "attachmentSessionIdValue"),
	}
}

// String returns a human-readable description of this User Todo List Task Attachment Session ID
func (id UserTodoListTaskAttachmentSessionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
		fmt.Sprintf("Attachment Session: %q", id.AttachmentSessionId),
	}
	return fmt.Sprintf("User Todo List Task Attachment Session (%s)", strings.Join(components, "\n"))
}
