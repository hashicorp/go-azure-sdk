package groupthreadpostinreplytoattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupThreadPostInReplyToAttachmentId{}

// GroupThreadPostInReplyToAttachmentId is a struct representing the Resource ID for a Group Thread Post In Reply To Attachment
type GroupThreadPostInReplyToAttachmentId struct {
	GroupId              string
	ConversationThreadId string
	PostId               string
	AttachmentId         string
}

// NewGroupThreadPostInReplyToAttachmentID returns a new GroupThreadPostInReplyToAttachmentId struct
func NewGroupThreadPostInReplyToAttachmentID(groupId string, conversationThreadId string, postId string, attachmentId string) GroupThreadPostInReplyToAttachmentId {
	return GroupThreadPostInReplyToAttachmentId{
		GroupId:              groupId,
		ConversationThreadId: conversationThreadId,
		PostId:               postId,
		AttachmentId:         attachmentId,
	}
}

// ParseGroupThreadPostInReplyToAttachmentID parses 'input' into a GroupThreadPostInReplyToAttachmentId
func ParseGroupThreadPostInReplyToAttachmentID(input string) (*GroupThreadPostInReplyToAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupThreadPostInReplyToAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupThreadPostInReplyToAttachmentId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
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

// ParseGroupThreadPostInReplyToAttachmentIDInsensitively parses 'input' case-insensitively into a GroupThreadPostInReplyToAttachmentId
// note: this method should only be used for API response data and not user input
func ParseGroupThreadPostInReplyToAttachmentIDInsensitively(input string) (*GroupThreadPostInReplyToAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupThreadPostInReplyToAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupThreadPostInReplyToAttachmentId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
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

// ValidateGroupThreadPostInReplyToAttachmentID checks that 'input' can be parsed as a Group Thread Post In Reply To Attachment ID
func ValidateGroupThreadPostInReplyToAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupThreadPostInReplyToAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Thread Post In Reply To Attachment ID
func (id GroupThreadPostInReplyToAttachmentId) ID() string {
	fmtString := "/groups/%s/threads/%s/posts/%s/inReplyTo/attachments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationThreadId, id.PostId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Thread Post In Reply To Attachment ID
func (id GroupThreadPostInReplyToAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticThreads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadIdValue"),
		resourceids.StaticSegment("staticPosts", "posts", "posts"),
		resourceids.UserSpecifiedSegment("postId", "postIdValue"),
		resourceids.StaticSegment("staticInReplyTo", "inReplyTo", "inReplyTo"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Group Thread Post In Reply To Attachment ID
func (id GroupThreadPostInReplyToAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
		fmt.Sprintf("Post: %q", id.PostId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Group Thread Post In Reply To Attachment (%s)", strings.Join(components, "\n"))
}
