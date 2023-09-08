package groupconversationthreadpostinreplytoattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupConversationThreadPostInReplyToAttachmentId{}

// GroupConversationThreadPostInReplyToAttachmentId is a struct representing the Resource ID for a Group Conversation Thread Post In Reply To Attachment
type GroupConversationThreadPostInReplyToAttachmentId struct {
	GroupId              string
	ConversationId       string
	ConversationThreadId string
	PostId               string
	AttachmentId         string
}

// NewGroupConversationThreadPostInReplyToAttachmentID returns a new GroupConversationThreadPostInReplyToAttachmentId struct
func NewGroupConversationThreadPostInReplyToAttachmentID(groupId string, conversationId string, conversationThreadId string, postId string, attachmentId string) GroupConversationThreadPostInReplyToAttachmentId {
	return GroupConversationThreadPostInReplyToAttachmentId{
		GroupId:              groupId,
		ConversationId:       conversationId,
		ConversationThreadId: conversationThreadId,
		PostId:               postId,
		AttachmentId:         attachmentId,
	}
}

// ParseGroupConversationThreadPostInReplyToAttachmentID parses 'input' into a GroupConversationThreadPostInReplyToAttachmentId
func ParseGroupConversationThreadPostInReplyToAttachmentID(input string) (*GroupConversationThreadPostInReplyToAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupConversationThreadPostInReplyToAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupConversationThreadPostInReplyToAttachmentId{}

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

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseGroupConversationThreadPostInReplyToAttachmentIDInsensitively parses 'input' case-insensitively into a GroupConversationThreadPostInReplyToAttachmentId
// note: this method should only be used for API response data and not user input
func ParseGroupConversationThreadPostInReplyToAttachmentIDInsensitively(input string) (*GroupConversationThreadPostInReplyToAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupConversationThreadPostInReplyToAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupConversationThreadPostInReplyToAttachmentId{}

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

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateGroupConversationThreadPostInReplyToAttachmentID checks that 'input' can be parsed as a Group Conversation Thread Post In Reply To Attachment ID
func ValidateGroupConversationThreadPostInReplyToAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupConversationThreadPostInReplyToAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Conversation Thread Post In Reply To Attachment ID
func (id GroupConversationThreadPostInReplyToAttachmentId) ID() string {
	fmtString := "/groups/%s/conversations/%s/threads/%s/posts/%s/inReplyTo/attachments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationId, id.ConversationThreadId, id.PostId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Conversation Thread Post In Reply To Attachment ID
func (id GroupConversationThreadPostInReplyToAttachmentId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Group Conversation Thread Post In Reply To Attachment ID
func (id GroupConversationThreadPostInReplyToAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation: %q", id.ConversationId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
		fmt.Sprintf("Post: %q", id.PostId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Group Conversation Thread Post In Reply To Attachment (%s)", strings.Join(components, "\n"))
}
