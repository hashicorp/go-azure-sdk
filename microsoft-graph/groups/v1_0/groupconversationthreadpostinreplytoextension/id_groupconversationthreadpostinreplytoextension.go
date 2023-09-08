package groupconversationthreadpostinreplytoextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupConversationThreadPostInReplyToExtensionId{}

// GroupConversationThreadPostInReplyToExtensionId is a struct representing the Resource ID for a Group Conversation Thread Post In Reply To Extension
type GroupConversationThreadPostInReplyToExtensionId struct {
	GroupId              string
	ConversationId       string
	ConversationThreadId string
	PostId               string
	ExtensionId          string
}

// NewGroupConversationThreadPostInReplyToExtensionID returns a new GroupConversationThreadPostInReplyToExtensionId struct
func NewGroupConversationThreadPostInReplyToExtensionID(groupId string, conversationId string, conversationThreadId string, postId string, extensionId string) GroupConversationThreadPostInReplyToExtensionId {
	return GroupConversationThreadPostInReplyToExtensionId{
		GroupId:              groupId,
		ConversationId:       conversationId,
		ConversationThreadId: conversationThreadId,
		PostId:               postId,
		ExtensionId:          extensionId,
	}
}

// ParseGroupConversationThreadPostInReplyToExtensionID parses 'input' into a GroupConversationThreadPostInReplyToExtensionId
func ParseGroupConversationThreadPostInReplyToExtensionID(input string) (*GroupConversationThreadPostInReplyToExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupConversationThreadPostInReplyToExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupConversationThreadPostInReplyToExtensionId{}

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

// ParseGroupConversationThreadPostInReplyToExtensionIDInsensitively parses 'input' case-insensitively into a GroupConversationThreadPostInReplyToExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupConversationThreadPostInReplyToExtensionIDInsensitively(input string) (*GroupConversationThreadPostInReplyToExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupConversationThreadPostInReplyToExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupConversationThreadPostInReplyToExtensionId{}

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

// ValidateGroupConversationThreadPostInReplyToExtensionID checks that 'input' can be parsed as a Group Conversation Thread Post In Reply To Extension ID
func ValidateGroupConversationThreadPostInReplyToExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupConversationThreadPostInReplyToExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Conversation Thread Post In Reply To Extension ID
func (id GroupConversationThreadPostInReplyToExtensionId) ID() string {
	fmtString := "/groups/%s/conversations/%s/threads/%s/posts/%s/inReplyTo/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationId, id.ConversationThreadId, id.PostId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Conversation Thread Post In Reply To Extension ID
func (id GroupConversationThreadPostInReplyToExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticConversations", "conversations", "conversations"),
		resourceids.UserSpecifiedSegment("conversationId", "conversationIdValue"),
		resourceids.StaticSegment("staticThreads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadIdValue"),
		resourceids.StaticSegment("staticPosts", "posts", "posts"),
		resourceids.UserSpecifiedSegment("postId", "postIdValue"),
		resourceids.StaticSegment("staticInReplyTo", "inReplyTo", "inReplyTo"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Group Conversation Thread Post In Reply To Extension ID
func (id GroupConversationThreadPostInReplyToExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation: %q", id.ConversationId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
		fmt.Sprintf("Post: %q", id.PostId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Conversation Thread Post In Reply To Extension (%s)", strings.Join(components, "\n"))
}
