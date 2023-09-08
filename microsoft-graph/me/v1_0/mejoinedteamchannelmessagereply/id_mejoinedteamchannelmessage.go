package mejoinedteamchannelmessagereply

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamChannelMessageId{}

// MeJoinedTeamChannelMessageId is a struct representing the Resource ID for a Me Joined Team Channel Message
type MeJoinedTeamChannelMessageId struct {
	TeamId        string
	ChannelId     string
	ChatMessageId string
}

// NewMeJoinedTeamChannelMessageID returns a new MeJoinedTeamChannelMessageId struct
func NewMeJoinedTeamChannelMessageID(teamId string, channelId string, chatMessageId string) MeJoinedTeamChannelMessageId {
	return MeJoinedTeamChannelMessageId{
		TeamId:        teamId,
		ChannelId:     channelId,
		ChatMessageId: chatMessageId,
	}
}

// ParseMeJoinedTeamChannelMessageID parses 'input' into a MeJoinedTeamChannelMessageId
func ParseMeJoinedTeamChannelMessageID(input string) (*MeJoinedTeamChannelMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamChannelMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamChannelMessageId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamChannelMessageIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamChannelMessageId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamChannelMessageIDInsensitively(input string) (*MeJoinedTeamChannelMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamChannelMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamChannelMessageId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamChannelMessageID checks that 'input' can be parsed as a Me Joined Team Channel Message ID
func ValidateMeJoinedTeamChannelMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamChannelMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Channel Message ID
func (id MeJoinedTeamChannelMessageId) ID() string {
	fmtString := "/me/joinedTeams/%s/channels/%s/messages/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId, id.ChatMessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Channel Message ID
func (id MeJoinedTeamChannelMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Channel Message ID
func (id MeJoinedTeamChannelMessageId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
	}
	return fmt.Sprintf("Me Joined Team Channel Message (%s)", strings.Join(components, "\n"))
}
