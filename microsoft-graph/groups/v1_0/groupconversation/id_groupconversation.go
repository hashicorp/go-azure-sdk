package groupconversation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupConversationId{}

// GroupConversationId is a struct representing the Resource ID for a Group Conversation
type GroupConversationId struct {
	GroupId        string
	ConversationId string
}

// NewGroupConversationID returns a new GroupConversationId struct
func NewGroupConversationID(groupId string, conversationId string) GroupConversationId {
	return GroupConversationId{
		GroupId:        groupId,
		ConversationId: conversationId,
	}
}

// ParseGroupConversationID parses 'input' into a GroupConversationId
func ParseGroupConversationID(input string) (*GroupConversationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupConversationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupConversationId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ConversationId, ok = parsed.Parsed["conversationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationId", *parsed)
	}

	return &id, nil
}

// ParseGroupConversationIDInsensitively parses 'input' case-insensitively into a GroupConversationId
// note: this method should only be used for API response data and not user input
func ParseGroupConversationIDInsensitively(input string) (*GroupConversationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupConversationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupConversationId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ConversationId, ok = parsed.Parsed["conversationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationId", *parsed)
	}

	return &id, nil
}

// ValidateGroupConversationID checks that 'input' can be parsed as a Group Conversation ID
func ValidateGroupConversationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupConversationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Conversation ID
func (id GroupConversationId) ID() string {
	fmtString := "/groups/%s/conversations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Conversation ID
func (id GroupConversationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticConversations", "conversations", "conversations"),
		resourceids.UserSpecifiedSegment("conversationId", "conversationIdValue"),
	}
}

// String returns a human-readable description of this Group Conversation ID
func (id GroupConversationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation: %q", id.ConversationId),
	}
	return fmt.Sprintf("Group Conversation (%s)", strings.Join(components, "\n"))
}
