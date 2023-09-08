package groupconversationthreadpostmention

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupConversationThreadPostMentionId{}

// GroupConversationThreadPostMentionId is a struct representing the Resource ID for a Group Conversation Thread Post Mention
type GroupConversationThreadPostMentionId struct {
	GroupId              string
	ConversationId       string
	ConversationThreadId string
	PostId               string
	MentionId            string
}

// NewGroupConversationThreadPostMentionID returns a new GroupConversationThreadPostMentionId struct
func NewGroupConversationThreadPostMentionID(groupId string, conversationId string, conversationThreadId string, postId string, mentionId string) GroupConversationThreadPostMentionId {
	return GroupConversationThreadPostMentionId{
		GroupId:              groupId,
		ConversationId:       conversationId,
		ConversationThreadId: conversationThreadId,
		PostId:               postId,
		MentionId:            mentionId,
	}
}

// ParseGroupConversationThreadPostMentionID parses 'input' into a GroupConversationThreadPostMentionId
func ParseGroupConversationThreadPostMentionID(input string) (*GroupConversationThreadPostMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupConversationThreadPostMentionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupConversationThreadPostMentionId{}

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

	if id.MentionId, ok = parsed.Parsed["mentionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mentionId", *parsed)
	}

	return &id, nil
}

// ParseGroupConversationThreadPostMentionIDInsensitively parses 'input' case-insensitively into a GroupConversationThreadPostMentionId
// note: this method should only be used for API response data and not user input
func ParseGroupConversationThreadPostMentionIDInsensitively(input string) (*GroupConversationThreadPostMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupConversationThreadPostMentionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupConversationThreadPostMentionId{}

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

	if id.MentionId, ok = parsed.Parsed["mentionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mentionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupConversationThreadPostMentionID checks that 'input' can be parsed as a Group Conversation Thread Post Mention ID
func ValidateGroupConversationThreadPostMentionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupConversationThreadPostMentionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Conversation Thread Post Mention ID
func (id GroupConversationThreadPostMentionId) ID() string {
	fmtString := "/groups/%s/conversations/%s/threads/%s/posts/%s/mentions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationId, id.ConversationThreadId, id.PostId, id.MentionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Conversation Thread Post Mention ID
func (id GroupConversationThreadPostMentionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticConversations", "conversations", "conversations"),
		resourceids.UserSpecifiedSegment("conversationId", "conversationIdValue"),
		resourceids.StaticSegment("staticThreads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadIdValue"),
		resourceids.StaticSegment("staticPosts", "posts", "posts"),
		resourceids.UserSpecifiedSegment("postId", "postIdValue"),
		resourceids.StaticSegment("staticMentions", "mentions", "mentions"),
		resourceids.UserSpecifiedSegment("mentionId", "mentionIdValue"),
	}
}

// String returns a human-readable description of this Group Conversation Thread Post Mention ID
func (id GroupConversationThreadPostMentionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation: %q", id.ConversationId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
		fmt.Sprintf("Post: %q", id.PostId),
		fmt.Sprintf("Mention: %q", id.MentionId),
	}
	return fmt.Sprintf("Group Conversation Thread Post Mention (%s)", strings.Join(components, "\n"))
}
