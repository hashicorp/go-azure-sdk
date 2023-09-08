package groupthreadpostinreplytoextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupThreadPostInReplyToExtensionId{}

// GroupThreadPostInReplyToExtensionId is a struct representing the Resource ID for a Group Thread Post In Reply To Extension
type GroupThreadPostInReplyToExtensionId struct {
	GroupId              string
	ConversationThreadId string
	PostId               string
	ExtensionId          string
}

// NewGroupThreadPostInReplyToExtensionID returns a new GroupThreadPostInReplyToExtensionId struct
func NewGroupThreadPostInReplyToExtensionID(groupId string, conversationThreadId string, postId string, extensionId string) GroupThreadPostInReplyToExtensionId {
	return GroupThreadPostInReplyToExtensionId{
		GroupId:              groupId,
		ConversationThreadId: conversationThreadId,
		PostId:               postId,
		ExtensionId:          extensionId,
	}
}

// ParseGroupThreadPostInReplyToExtensionID parses 'input' into a GroupThreadPostInReplyToExtensionId
func ParseGroupThreadPostInReplyToExtensionID(input string) (*GroupThreadPostInReplyToExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupThreadPostInReplyToExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupThreadPostInReplyToExtensionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
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

// ParseGroupThreadPostInReplyToExtensionIDInsensitively parses 'input' case-insensitively into a GroupThreadPostInReplyToExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupThreadPostInReplyToExtensionIDInsensitively(input string) (*GroupThreadPostInReplyToExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupThreadPostInReplyToExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupThreadPostInReplyToExtensionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
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

// ValidateGroupThreadPostInReplyToExtensionID checks that 'input' can be parsed as a Group Thread Post In Reply To Extension ID
func ValidateGroupThreadPostInReplyToExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupThreadPostInReplyToExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Thread Post In Reply To Extension ID
func (id GroupThreadPostInReplyToExtensionId) ID() string {
	fmtString := "/groups/%s/threads/%s/posts/%s/inReplyTo/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationThreadId, id.PostId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Thread Post In Reply To Extension ID
func (id GroupThreadPostInReplyToExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticThreads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadIdValue"),
		resourceids.StaticSegment("staticPosts", "posts", "posts"),
		resourceids.UserSpecifiedSegment("postId", "postIdValue"),
		resourceids.StaticSegment("staticInReplyTo", "inReplyTo", "inReplyTo"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Group Thread Post In Reply To Extension ID
func (id GroupThreadPostInReplyToExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
		fmt.Sprintf("Post: %q", id.PostId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Thread Post In Reply To Extension (%s)", strings.Join(components, "\n"))
}
