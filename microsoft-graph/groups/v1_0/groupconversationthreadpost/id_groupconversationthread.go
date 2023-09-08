package groupconversationthreadpost

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupConversationThreadId{}

// GroupConversationThreadId is a struct representing the Resource ID for a Group Conversation Thread
type GroupConversationThreadId struct {
	GroupId              string
	ConversationId       string
	ConversationThreadId string
}

// NewGroupConversationThreadID returns a new GroupConversationThreadId struct
func NewGroupConversationThreadID(groupId string, conversationId string, conversationThreadId string) GroupConversationThreadId {
	return GroupConversationThreadId{
		GroupId:              groupId,
		ConversationId:       conversationId,
		ConversationThreadId: conversationThreadId,
	}
}

// ParseGroupConversationThreadID parses 'input' into a GroupConversationThreadId
func ParseGroupConversationThreadID(input string) (*GroupConversationThreadId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupConversationThreadId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupConversationThreadId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ConversationId, ok = parsed.Parsed["conversationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationId", *parsed)
	}

	if id.ConversationThreadId, ok = parsed.Parsed["conversationThreadId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationThreadId", *parsed)
	}

	return &id, nil
}

// ParseGroupConversationThreadIDInsensitively parses 'input' case-insensitively into a GroupConversationThreadId
// note: this method should only be used for API response data and not user input
func ParseGroupConversationThreadIDInsensitively(input string) (*GroupConversationThreadId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupConversationThreadId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupConversationThreadId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ConversationId, ok = parsed.Parsed["conversationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationId", *parsed)
	}

	if id.ConversationThreadId, ok = parsed.Parsed["conversationThreadId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationThreadId", *parsed)
	}

	return &id, nil
}

// ValidateGroupConversationThreadID checks that 'input' can be parsed as a Group Conversation Thread ID
func ValidateGroupConversationThreadID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupConversationThreadID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Conversation Thread ID
func (id GroupConversationThreadId) ID() string {
	fmtString := "/groups/%s/conversations/%s/threads/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationId, id.ConversationThreadId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Conversation Thread ID
func (id GroupConversationThreadId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticConversations", "conversations", "conversations"),
		resourceids.UserSpecifiedSegment("conversationId", "conversationIdValue"),
		resourceids.StaticSegment("staticThreads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadIdValue"),
	}
}

// String returns a human-readable description of this Group Conversation Thread ID
func (id GroupConversationThreadId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation: %q", id.ConversationId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
	}
	return fmt.Sprintf("Group Conversation Thread (%s)", strings.Join(components, "\n"))
}
