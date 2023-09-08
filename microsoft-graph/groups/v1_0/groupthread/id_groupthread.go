package groupthread

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupThreadId{}

// GroupThreadId is a struct representing the Resource ID for a Group Thread
type GroupThreadId struct {
	GroupId              string
	ConversationThreadId string
}

// NewGroupThreadID returns a new GroupThreadId struct
func NewGroupThreadID(groupId string, conversationThreadId string) GroupThreadId {
	return GroupThreadId{
		GroupId:              groupId,
		ConversationThreadId: conversationThreadId,
	}
}

// ParseGroupThreadID parses 'input' into a GroupThreadId
func ParseGroupThreadID(input string) (*GroupThreadId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupThreadId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupThreadId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ConversationThreadId, ok = parsed.Parsed["conversationThreadId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationThreadId", *parsed)
	}

	return &id, nil
}

// ParseGroupThreadIDInsensitively parses 'input' case-insensitively into a GroupThreadId
// note: this method should only be used for API response data and not user input
func ParseGroupThreadIDInsensitively(input string) (*GroupThreadId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupThreadId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupThreadId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ConversationThreadId, ok = parsed.Parsed["conversationThreadId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationThreadId", *parsed)
	}

	return &id, nil
}

// ValidateGroupThreadID checks that 'input' can be parsed as a Group Thread ID
func ValidateGroupThreadID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupThreadID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Thread ID
func (id GroupThreadId) ID() string {
	fmtString := "/groups/%s/threads/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationThreadId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Thread ID
func (id GroupThreadId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticThreads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadIdValue"),
	}
}

// String returns a human-readable description of this Group Thread ID
func (id GroupThreadId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
	}
	return fmt.Sprintf("Group Thread (%s)", strings.Join(components, "\n"))
}
