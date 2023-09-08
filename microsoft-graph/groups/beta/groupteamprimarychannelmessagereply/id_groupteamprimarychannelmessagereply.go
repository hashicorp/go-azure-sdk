package groupteamprimarychannelmessagereply

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamPrimaryChannelMessageReplyId{}

// GroupTeamPrimaryChannelMessageReplyId is a struct representing the Resource ID for a Group Team Primary Channel Message Reply
type GroupTeamPrimaryChannelMessageReplyId struct {
	GroupId        string
	ChatMessageId  string
	ChatMessageId1 string
}

// NewGroupTeamPrimaryChannelMessageReplyID returns a new GroupTeamPrimaryChannelMessageReplyId struct
func NewGroupTeamPrimaryChannelMessageReplyID(groupId string, chatMessageId string, chatMessageId1 string) GroupTeamPrimaryChannelMessageReplyId {
	return GroupTeamPrimaryChannelMessageReplyId{
		GroupId:        groupId,
		ChatMessageId:  chatMessageId,
		ChatMessageId1: chatMessageId1,
	}
}

// ParseGroupTeamPrimaryChannelMessageReplyID parses 'input' into a GroupTeamPrimaryChannelMessageReplyId
func ParseGroupTeamPrimaryChannelMessageReplyID(input string) (*GroupTeamPrimaryChannelMessageReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamPrimaryChannelMessageReplyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamPrimaryChannelMessageReplyId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageId1, ok = parsed.Parsed["chatMessageId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamPrimaryChannelMessageReplyIDInsensitively parses 'input' case-insensitively into a GroupTeamPrimaryChannelMessageReplyId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamPrimaryChannelMessageReplyIDInsensitively(input string) (*GroupTeamPrimaryChannelMessageReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamPrimaryChannelMessageReplyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamPrimaryChannelMessageReplyId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageId1, ok = parsed.Parsed["chatMessageId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamPrimaryChannelMessageReplyID checks that 'input' can be parsed as a Group Team Primary Channel Message Reply ID
func ValidateGroupTeamPrimaryChannelMessageReplyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamPrimaryChannelMessageReplyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Primary Channel Message Reply ID
func (id GroupTeamPrimaryChannelMessageReplyId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/messages/%s/replies/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChatMessageId, id.ChatMessageId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Primary Channel Message Reply ID
func (id GroupTeamPrimaryChannelMessageReplyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
		resourceids.StaticSegment("staticReplies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1Value"),
	}
}

// String returns a human-readable description of this Group Team Primary Channel Message Reply ID
func (id GroupTeamPrimaryChannelMessageReplyId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
	}
	return fmt.Sprintf("Group Team Primary Channel Message Reply (%s)", strings.Join(components, "\n"))
}
