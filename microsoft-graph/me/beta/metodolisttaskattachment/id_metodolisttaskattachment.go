package metodolisttaskattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeTodoListTaskAttachmentId{}

// MeTodoListTaskAttachmentId is a struct representing the Resource ID for a Me Todo List Task Attachment
type MeTodoListTaskAttachmentId struct {
	TodoTaskListId   string
	TodoTaskId       string
	AttachmentBaseId string
}

// NewMeTodoListTaskAttachmentID returns a new MeTodoListTaskAttachmentId struct
func NewMeTodoListTaskAttachmentID(todoTaskListId string, todoTaskId string, attachmentBaseId string) MeTodoListTaskAttachmentId {
	return MeTodoListTaskAttachmentId{
		TodoTaskListId:   todoTaskListId,
		TodoTaskId:       todoTaskId,
		AttachmentBaseId: attachmentBaseId,
	}
}

// ParseMeTodoListTaskAttachmentID parses 'input' into a MeTodoListTaskAttachmentId
func ParseMeTodoListTaskAttachmentID(input string) (*MeTodoListTaskAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeTodoListTaskAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeTodoListTaskAttachmentId{}

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

// ParseMeTodoListTaskAttachmentIDInsensitively parses 'input' case-insensitively into a MeTodoListTaskAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeTodoListTaskAttachmentIDInsensitively(input string) (*MeTodoListTaskAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeTodoListTaskAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeTodoListTaskAttachmentId{}

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

// ValidateMeTodoListTaskAttachmentID checks that 'input' can be parsed as a Me Todo List Task Attachment ID
func ValidateMeTodoListTaskAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeTodoListTaskAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Todo List Task Attachment ID
func (id MeTodoListTaskAttachmentId) ID() string {
	fmtString := "/me/todo/lists/%s/tasks/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.TodoTaskListId, id.TodoTaskId, id.AttachmentBaseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Todo List Task Attachment ID
func (id MeTodoListTaskAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticTodo", "todo", "todo"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentBaseId", "attachmentBaseIdValue"),
	}
}

// String returns a human-readable description of this Me Todo List Task Attachment ID
func (id MeTodoListTaskAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
		fmt.Sprintf("Attachment Base: %q", id.AttachmentBaseId),
	}
	return fmt.Sprintf("Me Todo List Task Attachment (%s)", strings.Join(components, "\n"))
}
