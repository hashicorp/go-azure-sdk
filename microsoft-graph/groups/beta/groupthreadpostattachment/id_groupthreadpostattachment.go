package groupthreadpostattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupThreadPostAttachmentId{}

// GroupThreadPostAttachmentId is a struct representing the Resource ID for a Group Thread Post Attachment
type GroupThreadPostAttachmentId struct {
	GroupId              string
	ConversationThreadId string
	PostId               string
	AttachmentId         string
}

// NewGroupThreadPostAttachmentID returns a new GroupThreadPostAttachmentId struct
func NewGroupThreadPostAttachmentID(groupId string, conversationThreadId string, postId string, attachmentId string) GroupThreadPostAttachmentId {
	return GroupThreadPostAttachmentId{
		GroupId:              groupId,
		ConversationThreadId: conversationThreadId,
		PostId:               postId,
		AttachmentId:         attachmentId,
	}
}

// ParseGroupThreadPostAttachmentID parses 'input' into a GroupThreadPostAttachmentId
func ParseGroupThreadPostAttachmentID(input string) (*GroupThreadPostAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupThreadPostAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupThreadPostAttachmentId{}

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

// ParseGroupThreadPostAttachmentIDInsensitively parses 'input' case-insensitively into a GroupThreadPostAttachmentId
// note: this method should only be used for API response data and not user input
func ParseGroupThreadPostAttachmentIDInsensitively(input string) (*GroupThreadPostAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupThreadPostAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupThreadPostAttachmentId{}

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

// ValidateGroupThreadPostAttachmentID checks that 'input' can be parsed as a Group Thread Post Attachment ID
func ValidateGroupThreadPostAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupThreadPostAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Thread Post Attachment ID
func (id GroupThreadPostAttachmentId) ID() string {
	fmtString := "/groups/%s/threads/%s/posts/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationThreadId, id.PostId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Thread Post Attachment ID
func (id GroupThreadPostAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticThreads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadIdValue"),
		resourceids.StaticSegment("staticPosts", "posts", "posts"),
		resourceids.UserSpecifiedSegment("postId", "postIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Group Thread Post Attachment ID
func (id GroupThreadPostAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
		fmt.Sprintf("Post: %q", id.PostId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Group Thread Post Attachment (%s)", strings.Join(components, "\n"))
}
