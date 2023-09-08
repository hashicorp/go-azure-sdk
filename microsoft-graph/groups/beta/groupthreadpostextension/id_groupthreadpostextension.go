package groupthreadpostextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupThreadPostExtensionId{}

// GroupThreadPostExtensionId is a struct representing the Resource ID for a Group Thread Post Extension
type GroupThreadPostExtensionId struct {
	GroupId              string
	ConversationThreadId string
	PostId               string
	ExtensionId          string
}

// NewGroupThreadPostExtensionID returns a new GroupThreadPostExtensionId struct
func NewGroupThreadPostExtensionID(groupId string, conversationThreadId string, postId string, extensionId string) GroupThreadPostExtensionId {
	return GroupThreadPostExtensionId{
		GroupId:              groupId,
		ConversationThreadId: conversationThreadId,
		PostId:               postId,
		ExtensionId:          extensionId,
	}
}

// ParseGroupThreadPostExtensionID parses 'input' into a GroupThreadPostExtensionId
func ParseGroupThreadPostExtensionID(input string) (*GroupThreadPostExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupThreadPostExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupThreadPostExtensionId{}

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

// ParseGroupThreadPostExtensionIDInsensitively parses 'input' case-insensitively into a GroupThreadPostExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupThreadPostExtensionIDInsensitively(input string) (*GroupThreadPostExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupThreadPostExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupThreadPostExtensionId{}

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

// ValidateGroupThreadPostExtensionID checks that 'input' can be parsed as a Group Thread Post Extension ID
func ValidateGroupThreadPostExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupThreadPostExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Thread Post Extension ID
func (id GroupThreadPostExtensionId) ID() string {
	fmtString := "/groups/%s/threads/%s/posts/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationThreadId, id.PostId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Thread Post Extension ID
func (id GroupThreadPostExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticThreads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadIdValue"),
		resourceids.StaticSegment("staticPosts", "posts", "posts"),
		resourceids.UserSpecifiedSegment("postId", "postIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Group Thread Post Extension ID
func (id GroupThreadPostExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
		fmt.Sprintf("Post: %q", id.PostId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Thread Post Extension (%s)", strings.Join(components, "\n"))
}
