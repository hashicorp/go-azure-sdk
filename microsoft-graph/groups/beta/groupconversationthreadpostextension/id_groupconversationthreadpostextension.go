package groupconversationthreadpostextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupConversationThreadPostExtensionId{}

// GroupConversationThreadPostExtensionId is a struct representing the Resource ID for a Group Conversation Thread Post Extension
type GroupConversationThreadPostExtensionId struct {
	GroupId              string
	ConversationId       string
	ConversationThreadId string
	PostId               string
	ExtensionId          string
}

// NewGroupConversationThreadPostExtensionID returns a new GroupConversationThreadPostExtensionId struct
func NewGroupConversationThreadPostExtensionID(groupId string, conversationId string, conversationThreadId string, postId string, extensionId string) GroupConversationThreadPostExtensionId {
	return GroupConversationThreadPostExtensionId{
		GroupId:              groupId,
		ConversationId:       conversationId,
		ConversationThreadId: conversationThreadId,
		PostId:               postId,
		ExtensionId:          extensionId,
	}
}

// ParseGroupConversationThreadPostExtensionID parses 'input' into a GroupConversationThreadPostExtensionId
func ParseGroupConversationThreadPostExtensionID(input string) (*GroupConversationThreadPostExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupConversationThreadPostExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupConversationThreadPostExtensionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ConversationId, ok = parsed.Parsed["conversationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationId", *parsed)
	}

	if id.ConversationThreadId, ok = parsed.Parsed["conversationThreadId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationThreadId", *parsed)
	}

	if id.PostId, ok = parsed.Parsed["postId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "postId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseGroupConversationThreadPostExtensionIDInsensitively parses 'input' case-insensitively into a GroupConversationThreadPostExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupConversationThreadPostExtensionIDInsensitively(input string) (*GroupConversationThreadPostExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupConversationThreadPostExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupConversationThreadPostExtensionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ConversationId, ok = parsed.Parsed["conversationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationId", *parsed)
	}

	if id.ConversationThreadId, ok = parsed.Parsed["conversationThreadId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationThreadId", *parsed)
	}

	if id.PostId, ok = parsed.Parsed["postId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "postId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupConversationThreadPostExtensionID checks that 'input' can be parsed as a Group Conversation Thread Post Extension ID
func ValidateGroupConversationThreadPostExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupConversationThreadPostExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Conversation Thread Post Extension ID
func (id GroupConversationThreadPostExtensionId) ID() string {
	fmtString := "/groups/%s/conversations/%s/threads/%s/posts/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationId, id.ConversationThreadId, id.PostId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Conversation Thread Post Extension ID
func (id GroupConversationThreadPostExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticConversations", "conversations", "conversations"),
		resourceids.UserSpecifiedSegment("conversationId", "conversationIdValue"),
		resourceids.StaticSegment("staticThreads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadIdValue"),
		resourceids.StaticSegment("staticPosts", "posts", "posts"),
		resourceids.UserSpecifiedSegment("postId", "postIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Group Conversation Thread Post Extension ID
func (id GroupConversationThreadPostExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation: %q", id.ConversationId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
		fmt.Sprintf("Post: %q", id.PostId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Conversation Thread Post Extension (%s)", strings.Join(components, "\n"))
}
