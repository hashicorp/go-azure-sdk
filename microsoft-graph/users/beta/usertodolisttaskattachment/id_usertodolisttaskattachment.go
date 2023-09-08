package usertodolisttaskattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserTodoListTaskAttachmentId{}

// UserTodoListTaskAttachmentId is a struct representing the Resource ID for a User Todo List Task Attachment
type UserTodoListTaskAttachmentId struct {
	UserId           string
	TodoTaskListId   string
	TodoTaskId       string
	AttachmentBaseId string
}

// NewUserTodoListTaskAttachmentID returns a new UserTodoListTaskAttachmentId struct
func NewUserTodoListTaskAttachmentID(userId string, todoTaskListId string, todoTaskId string, attachmentBaseId string) UserTodoListTaskAttachmentId {
	return UserTodoListTaskAttachmentId{
		UserId:           userId,
		TodoTaskListId:   todoTaskListId,
		TodoTaskId:       todoTaskId,
		AttachmentBaseId: attachmentBaseId,
	}
}

// ParseUserTodoListTaskAttachmentID parses 'input' into a UserTodoListTaskAttachmentId
func ParseUserTodoListTaskAttachmentID(input string) (*UserTodoListTaskAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTodoListTaskAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTodoListTaskAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	if id.TodoTaskId, ok = parsed.Parsed["todoTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskId", *parsed)
	}

	if id.AttachmentBaseId, ok = parsed.Parsed["attachmentBaseId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentBaseId", *parsed)
	}

	return &id, nil
}

// ParseUserTodoListTaskAttachmentIDInsensitively parses 'input' case-insensitively into a UserTodoListTaskAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserTodoListTaskAttachmentIDInsensitively(input string) (*UserTodoListTaskAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTodoListTaskAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTodoListTaskAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	if id.TodoTaskId, ok = parsed.Parsed["todoTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskId", *parsed)
	}

	if id.AttachmentBaseId, ok = parsed.Parsed["attachmentBaseId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentBaseId", *parsed)
	}

	return &id, nil
}

// ValidateUserTodoListTaskAttachmentID checks that 'input' can be parsed as a User Todo List Task Attachment ID
func ValidateUserTodoListTaskAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserTodoListTaskAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Todo List Task Attachment ID
func (id UserTodoListTaskAttachmentId) ID() string {
	fmtString := "/users/%s/todo/lists/%s/tasks/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TodoTaskListId, id.TodoTaskId, id.AttachmentBaseId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Todo List Task Attachment ID
func (id UserTodoListTaskAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticTodo", "todo", "todo"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentBaseId", "attachmentBaseIdValue"),
	}
}

// String returns a human-readable description of this User Todo List Task Attachment ID
func (id UserTodoListTaskAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
		fmt.Sprintf("Attachment Base: %q", id.AttachmentBaseId),
	}
	return fmt.Sprintf("User Todo List Task Attachment (%s)", strings.Join(components, "\n"))
}
