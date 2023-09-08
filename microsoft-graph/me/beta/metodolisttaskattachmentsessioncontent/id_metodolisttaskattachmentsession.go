package metodolisttaskattachmentsessioncontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeTodoListTaskAttachmentSessionId{}

// MeTodoListTaskAttachmentSessionId is a struct representing the Resource ID for a Me Todo List Task Attachment Session
type MeTodoListTaskAttachmentSessionId struct {
	TodoTaskListId      string
	TodoTaskId          string
	AttachmentSessionId string
}

// NewMeTodoListTaskAttachmentSessionID returns a new MeTodoListTaskAttachmentSessionId struct
func NewMeTodoListTaskAttachmentSessionID(todoTaskListId string, todoTaskId string, attachmentSessionId string) MeTodoListTaskAttachmentSessionId {
	return MeTodoListTaskAttachmentSessionId{
		TodoTaskListId:      todoTaskListId,
		TodoTaskId:          todoTaskId,
		AttachmentSessionId: attachmentSessionId,
	}
}

// ParseMeTodoListTaskAttachmentSessionID parses 'input' into a MeTodoListTaskAttachmentSessionId
func ParseMeTodoListTaskAttachmentSessionID(input string) (*MeTodoListTaskAttachmentSessionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeTodoListTaskAttachmentSessionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeTodoListTaskAttachmentSessionId{}

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

// ParseMeTodoListTaskAttachmentSessionIDInsensitively parses 'input' case-insensitively into a MeTodoListTaskAttachmentSessionId
// note: this method should only be used for API response data and not user input
func ParseMeTodoListTaskAttachmentSessionIDInsensitively(input string) (*MeTodoListTaskAttachmentSessionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeTodoListTaskAttachmentSessionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeTodoListTaskAttachmentSessionId{}

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

// ValidateMeTodoListTaskAttachmentSessionID checks that 'input' can be parsed as a Me Todo List Task Attachment Session ID
func ValidateMeTodoListTaskAttachmentSessionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeTodoListTaskAttachmentSessionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Todo List Task Attachment Session ID
func (id MeTodoListTaskAttachmentSessionId) ID() string {
	fmtString := "/me/todo/lists/%s/tasks/%s/attachmentSessions/%s"
	return fmt.Sprintf(fmtString, id.TodoTaskListId, id.TodoTaskId, id.AttachmentSessionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Todo List Task Attachment Session ID
func (id MeTodoListTaskAttachmentSessionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticTodo", "todo", "todo"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskIdValue"),
		resourceids.StaticSegment("staticAttachmentSessions", "attachmentSessions", "attachmentSessions"),
		resourceids.UserSpecifiedSegment("attachmentSessionId", "attachmentSessionIdValue"),
	}
}

// String returns a human-readable description of this Me Todo List Task Attachment Session ID
func (id MeTodoListTaskAttachmentSessionId) String() string {
	components := []string{
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
		fmt.Sprintf("Attachment Session: %q", id.AttachmentSessionId),
	}
	return fmt.Sprintf("Me Todo List Task Attachment Session (%s)", strings.Join(components, "\n"))
}
