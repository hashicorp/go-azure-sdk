package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdConversationIdThreadIdPostId{}

// GroupIdConversationIdThreadIdPostId is a struct representing the Resource ID for a Group Id Conversation Id Thread Id Post
type GroupIdConversationIdThreadIdPostId struct {
	GroupId              string
	ConversationId       string
	ConversationThreadId string
	PostId               string
}

// NewGroupIdConversationIdThreadIdPostID returns a new GroupIdConversationIdThreadIdPostId struct
func NewGroupIdConversationIdThreadIdPostID(groupId string, conversationId string, conversationThreadId string, postId string) GroupIdConversationIdThreadIdPostId {
	return GroupIdConversationIdThreadIdPostId{
		GroupId:              groupId,
		ConversationId:       conversationId,
		ConversationThreadId: conversationThreadId,
		PostId:               postId,
	}
}

// ParseGroupIdConversationIdThreadIdPostID parses 'input' into a GroupIdConversationIdThreadIdPostId
func ParseGroupIdConversationIdThreadIdPostID(input string) (*GroupIdConversationIdThreadIdPostId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdConversationIdThreadIdPostId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdConversationIdThreadIdPostId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdConversationIdThreadIdPostIDInsensitively parses 'input' case-insensitively into a GroupIdConversationIdThreadIdPostId
// note: this method should only be used for API response data and not user input
func ParseGroupIdConversationIdThreadIdPostIDInsensitively(input string) (*GroupIdConversationIdThreadIdPostId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdConversationIdThreadIdPostId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdConversationIdThreadIdPostId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdConversationIdThreadIdPostId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ConversationId, ok = input.Parsed["conversationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationId", input)
	}

	if id.ConversationThreadId, ok = input.Parsed["conversationThreadId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationThreadId", input)
	}

	if id.PostId, ok = input.Parsed["postId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "postId", input)
	}

	return nil
}

// ValidateGroupIdConversationIdThreadIdPostID checks that 'input' can be parsed as a Group Id Conversation Id Thread Id Post ID
func ValidateGroupIdConversationIdThreadIdPostID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdConversationIdThreadIdPostID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Conversation Id Thread Id Post ID
func (id GroupIdConversationIdThreadIdPostId) ID() string {
	fmtString := "/groups/%s/conversations/%s/threads/%s/posts/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationId, id.ConversationThreadId, id.PostId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Conversation Id Thread Id Post ID
func (id GroupIdConversationIdThreadIdPostId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("conversations", "conversations", "conversations"),
		resourceids.UserSpecifiedSegment("conversationId", "conversationId"),
		resourceids.StaticSegment("threads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadId"),
		resourceids.StaticSegment("posts", "posts", "posts"),
		resourceids.UserSpecifiedSegment("postId", "postId"),
	}
}

// String returns a human-readable description of this Group Id Conversation Id Thread Id Post ID
func (id GroupIdConversationIdThreadIdPostId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation: %q", id.ConversationId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
		fmt.Sprintf("Post: %q", id.PostId),
	}
	return fmt.Sprintf("Group Id Conversation Id Thread Id Post (%s)", strings.Join(components, "\n"))
}
