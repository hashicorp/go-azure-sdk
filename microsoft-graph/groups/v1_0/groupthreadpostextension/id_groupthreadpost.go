package groupthreadpostextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupThreadPostId{}

// GroupThreadPostId is a struct representing the Resource ID for a Group Thread Post
type GroupThreadPostId struct {
	GroupId              string
	ConversationThreadId string
	PostId               string
}

// NewGroupThreadPostID returns a new GroupThreadPostId struct
func NewGroupThreadPostID(groupId string, conversationThreadId string, postId string) GroupThreadPostId {
	return GroupThreadPostId{
		GroupId:              groupId,
		ConversationThreadId: conversationThreadId,
		PostId:               postId,
	}
}

// ParseGroupThreadPostID parses 'input' into a GroupThreadPostId
func ParseGroupThreadPostID(input string) (*GroupThreadPostId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupThreadPostId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupThreadPostId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ConversationThreadId, ok = parsed.Parsed["conversationThreadId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationThreadId", *parsed)
	}

	if id.PostId, ok = parsed.Parsed["postId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "postId", *parsed)
	}

	return &id, nil
}

// ParseGroupThreadPostIDInsensitively parses 'input' case-insensitively into a GroupThreadPostId
// note: this method should only be used for API response data and not user input
func ParseGroupThreadPostIDInsensitively(input string) (*GroupThreadPostId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupThreadPostId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupThreadPostId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ConversationThreadId, ok = parsed.Parsed["conversationThreadId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationThreadId", *parsed)
	}

	if id.PostId, ok = parsed.Parsed["postId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "postId", *parsed)
	}

	return &id, nil
}

// ValidateGroupThreadPostID checks that 'input' can be parsed as a Group Thread Post ID
func ValidateGroupThreadPostID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupThreadPostID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Thread Post ID
func (id GroupThreadPostId) ID() string {
	fmtString := "/groups/%s/threads/%s/posts/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationThreadId, id.PostId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Thread Post ID
func (id GroupThreadPostId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticThreads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadIdValue"),
		resourceids.StaticSegment("staticPosts", "posts", "posts"),
		resourceids.UserSpecifiedSegment("postId", "postIdValue"),
	}
}

// String returns a human-readable description of this Group Thread Post ID
func (id GroupThreadPostId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
		fmt.Sprintf("Post: %q", id.PostId),
	}
	return fmt.Sprintf("Group Thread Post (%s)", strings.Join(components, "\n"))
}
