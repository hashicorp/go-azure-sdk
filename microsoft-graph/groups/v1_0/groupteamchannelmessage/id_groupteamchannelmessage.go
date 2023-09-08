package groupteamchannelmessage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamChannelMessageId{}

// GroupTeamChannelMessageId is a struct representing the Resource ID for a Group Team Channel Message
type GroupTeamChannelMessageId struct {
	GroupId       string
	ChannelId     string
	ChatMessageId string
}

// NewGroupTeamChannelMessageID returns a new GroupTeamChannelMessageId struct
func NewGroupTeamChannelMessageID(groupId string, channelId string, chatMessageId string) GroupTeamChannelMessageId {
	return GroupTeamChannelMessageId{
		GroupId:       groupId,
		ChannelId:     channelId,
		ChatMessageId: chatMessageId,
	}
}

// ParseGroupTeamChannelMessageID parses 'input' into a GroupTeamChannelMessageId
func ParseGroupTeamChannelMessageID(input string) (*GroupTeamChannelMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamChannelMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamChannelMessageId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamChannelMessageIDInsensitively parses 'input' case-insensitively into a GroupTeamChannelMessageId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamChannelMessageIDInsensitively(input string) (*GroupTeamChannelMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamChannelMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamChannelMessageId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamChannelMessageID checks that 'input' can be parsed as a Group Team Channel Message ID
func ValidateGroupTeamChannelMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamChannelMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Channel Message ID
func (id GroupTeamChannelMessageId) ID() string {
	fmtString := "/groups/%s/team/channels/%s/messages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId, id.ChatMessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Channel Message ID
func (id GroupTeamChannelMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
	}
}

// String returns a human-readable description of this Group Team Channel Message ID
func (id GroupTeamChannelMessageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
	}
	return fmt.Sprintf("Group Team Channel Message (%s)", strings.Join(components, "\n"))
}
