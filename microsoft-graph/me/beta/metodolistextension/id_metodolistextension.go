package metodolistextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeTodoListExtensionId{}

// MeTodoListExtensionId is a struct representing the Resource ID for a Me Todo List Extension
type MeTodoListExtensionId struct {
	TodoTaskListId string
	ExtensionId    string
}

// NewMeTodoListExtensionID returns a new MeTodoListExtensionId struct
func NewMeTodoListExtensionID(todoTaskListId string, extensionId string) MeTodoListExtensionId {
	return MeTodoListExtensionId{
		TodoTaskListId: todoTaskListId,
		ExtensionId:    extensionId,
	}
}

// ParseMeTodoListExtensionID parses 'input' into a MeTodoListExtensionId
func ParseMeTodoListExtensionID(input string) (*MeTodoListExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeTodoListExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeTodoListExtensionId{}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseMeTodoListExtensionIDInsensitively parses 'input' case-insensitively into a MeTodoListExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeTodoListExtensionIDInsensitively(input string) (*MeTodoListExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeTodoListExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeTodoListExtensionId{}

	if id.TodoTaskListId, ok = parsed.Parsed["todoTaskListId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateMeTodoListExtensionID checks that 'input' can be parsed as a Me Todo List Extension ID
func ValidateMeTodoListExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeTodoListExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Todo List Extension ID
func (id MeTodoListExtensionId) ID() string {
	fmtString := "/me/todo/lists/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.TodoTaskListId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Todo List Extension ID
func (id MeTodoListExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticTodo", "todo", "todo"),
		resourceids.StaticSegment("staticLists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Todo List Extension ID
func (id MeTodoListExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Todo List Extension (%s)", strings.Join(components, "\n"))
}
