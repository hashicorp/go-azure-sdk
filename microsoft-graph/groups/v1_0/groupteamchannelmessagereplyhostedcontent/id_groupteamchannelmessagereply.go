package groupteamchannelmessagereplyhostedcontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamChannelMessageReplyId{}

// GroupTeamChannelMessageReplyId is a struct representing the Resource ID for a Group Team Channel Message Reply
type GroupTeamChannelMessageReplyId struct {
	GroupId        string
	ChannelId      string
	ChatMessageId  string
	ChatMessageId1 string
}

// NewGroupTeamChannelMessageReplyID returns a new GroupTeamChannelMessageReplyId struct
func NewGroupTeamChannelMessageReplyID(groupId string, channelId string, chatMessageId string, chatMessageId1 string) GroupTeamChannelMessageReplyId {
	return GroupTeamChannelMessageReplyId{
		GroupId:        groupId,
		ChannelId:      channelId,
		ChatMessageId:  chatMessageId,
		ChatMessageId1: chatMessageId1,
	}
}

// ParseGroupTeamChannelMessageReplyID parses 'input' into a GroupTeamChannelMessageReplyId
func ParseGroupTeamChannelMessageReplyID(input string) (*GroupTeamChannelMessageReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamChannelMessageReplyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamChannelMessageReplyId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageId1, ok = parsed.Parsed["chatMessageId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamChannelMessageReplyIDInsensitively parses 'input' case-insensitively into a GroupTeamChannelMessageReplyId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamChannelMessageReplyIDInsensitively(input string) (*GroupTeamChannelMessageReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamChannelMessageReplyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamChannelMessageReplyId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageId1, ok = parsed.Parsed["chatMessageId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamChannelMessageReplyID checks that 'input' can be parsed as a Group Team Channel Message Reply ID
func ValidateGroupTeamChannelMessageReplyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamChannelMessageReplyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Channel Message Reply ID
func (id GroupTeamChannelMessageReplyId) ID() string {
	fmtString := "/groups/%s/team/channels/%s/messages/%s/replies/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId, id.ChatMessageId, id.ChatMessageId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Channel Message Reply ID
func (id GroupTeamChannelMessageReplyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
		resourceids.StaticSegment("staticReplies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1Value"),
	}
}

// String returns a human-readable description of this Group Team Channel Message Reply ID
func (id GroupTeamChannelMessageReplyId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
	}
	return fmt.Sprintf("Group Team Channel Message Reply (%s)", strings.Join(components, "\n"))
}
