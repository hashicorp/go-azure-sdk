package groupthreadpostinreplytomention

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupThreadPostInReplyToMentionId{}

// GroupThreadPostInReplyToMentionId is a struct representing the Resource ID for a Group Thread Post In Reply To Mention
type GroupThreadPostInReplyToMentionId struct {
	GroupId              string
	ConversationThreadId string
	PostId               string
	MentionId            string
}

// NewGroupThreadPostInReplyToMentionID returns a new GroupThreadPostInReplyToMentionId struct
func NewGroupThreadPostInReplyToMentionID(groupId string, conversationThreadId string, postId string, mentionId string) GroupThreadPostInReplyToMentionId {
	return GroupThreadPostInReplyToMentionId{
		GroupId:              groupId,
		ConversationThreadId: conversationThreadId,
		PostId:               postId,
		MentionId:            mentionId,
	}
}

// ParseGroupThreadPostInReplyToMentionID parses 'input' into a GroupThreadPostInReplyToMentionId
func ParseGroupThreadPostInReplyToMentionID(input string) (*GroupThreadPostInReplyToMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupThreadPostInReplyToMentionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupThreadPostInReplyToMentionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
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

// ParseGroupThreadPostInReplyToMentionIDInsensitively parses 'input' case-insensitively into a GroupThreadPostInReplyToMentionId
// note: this method should only be used for API response data and not user input
func ParseGroupThreadPostInReplyToMentionIDInsensitively(input string) (*GroupThreadPostInReplyToMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupThreadPostInReplyToMentionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupThreadPostInReplyToMentionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
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

// ValidateGroupThreadPostInReplyToMentionID checks that 'input' can be parsed as a Group Thread Post In Reply To Mention ID
func ValidateGroupThreadPostInReplyToMentionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupThreadPostInReplyToMentionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Thread Post In Reply To Mention ID
func (id GroupThreadPostInReplyToMentionId) ID() string {
	fmtString := "/groups/%s/threads/%s/posts/%s/inReplyTo/mentions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationThreadId, id.PostId, id.MentionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Thread Post In Reply To Mention ID
func (id GroupThreadPostInReplyToMentionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticThreads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadIdValue"),
		resourceids.StaticSegment("staticPosts", "posts", "posts"),
		resourceids.UserSpecifiedSegment("postId", "postIdValue"),
		resourceids.StaticSegment("staticInReplyTo", "inReplyTo", "inReplyTo"),
		resourceids.StaticSegment("staticMentions", "mentions", "mentions"),
		resourceids.UserSpecifiedSegment("mentionId", "mentionIdValue"),
	}
}

// String returns a human-readable description of this Group Thread Post In Reply To Mention ID
func (id GroupThreadPostInReplyToMentionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
		fmt.Sprintf("Post: %q", id.PostId),
		fmt.Sprintf("Mention: %q", id.MentionId),
	}
	return fmt.Sprintf("Group Thread Post In Reply To Mention (%s)", strings.Join(components, "\n"))
}
