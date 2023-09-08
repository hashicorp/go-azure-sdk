package groupteamchannelmessagehostedcontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamChannelMessageHostedContentId{}

// GroupTeamChannelMessageHostedContentId is a struct representing the Resource ID for a Group Team Channel Message Hosted Content
type GroupTeamChannelMessageHostedContentId struct {
	GroupId                    string
	ChannelId                  string
	ChatMessageId              string
	ChatMessageHostedContentId string
}

// NewGroupTeamChannelMessageHostedContentID returns a new GroupTeamChannelMessageHostedContentId struct
func NewGroupTeamChannelMessageHostedContentID(groupId string, channelId string, chatMessageId string, chatMessageHostedContentId string) GroupTeamChannelMessageHostedContentId {
	return GroupTeamChannelMessageHostedContentId{
		GroupId:                    groupId,
		ChannelId:                  channelId,
		ChatMessageId:              chatMessageId,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseGroupTeamChannelMessageHostedContentID parses 'input' into a GroupTeamChannelMessageHostedContentId
func ParseGroupTeamChannelMessageHostedContentID(input string) (*GroupTeamChannelMessageHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamChannelMessageHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamChannelMessageHostedContentId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageHostedContentId, ok = parsed.Parsed["chatMessageHostedContentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamChannelMessageHostedContentIDInsensitively parses 'input' case-insensitively into a GroupTeamChannelMessageHostedContentId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamChannelMessageHostedContentIDInsensitively(input string) (*GroupTeamChannelMessageHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamChannelMessageHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamChannelMessageHostedContentId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageHostedContentId, ok = parsed.Parsed["chatMessageHostedContentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamChannelMessageHostedContentID checks that 'input' can be parsed as a Group Team Channel Message Hosted Content ID
func ValidateGroupTeamChannelMessageHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamChannelMessageHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Channel Message Hosted Content ID
func (id GroupTeamChannelMessageHostedContentId) ID() string {
	fmtString := "/groups/%s/team/channels/%s/messages/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId, id.ChatMessageId, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Channel Message Hosted Content ID
func (id GroupTeamChannelMessageHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
		resourceids.StaticSegment("staticHostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentIdValue"),
	}
}

// String returns a human-readable description of this Group Team Channel Message Hosted Content ID
func (id GroupTeamChannelMessageHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("Group Team Channel Message Hosted Content (%s)", strings.Join(components, "\n"))
}
