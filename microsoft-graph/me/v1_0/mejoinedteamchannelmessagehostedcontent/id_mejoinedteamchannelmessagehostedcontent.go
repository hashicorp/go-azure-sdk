package mejoinedteamchannelmessagehostedcontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamChannelMessageHostedContentId{}

// MeJoinedTeamChannelMessageHostedContentId is a struct representing the Resource ID for a Me Joined Team Channel Message Hosted Content
type MeJoinedTeamChannelMessageHostedContentId struct {
	TeamId                     string
	ChannelId                  string
	ChatMessageId              string
	ChatMessageHostedContentId string
}

// NewMeJoinedTeamChannelMessageHostedContentID returns a new MeJoinedTeamChannelMessageHostedContentId struct
func NewMeJoinedTeamChannelMessageHostedContentID(teamId string, channelId string, chatMessageId string, chatMessageHostedContentId string) MeJoinedTeamChannelMessageHostedContentId {
	return MeJoinedTeamChannelMessageHostedContentId{
		TeamId:                     teamId,
		ChannelId:                  channelId,
		ChatMessageId:              chatMessageId,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseMeJoinedTeamChannelMessageHostedContentID parses 'input' into a MeJoinedTeamChannelMessageHostedContentId
func ParseMeJoinedTeamChannelMessageHostedContentID(input string) (*MeJoinedTeamChannelMessageHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamChannelMessageHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamChannelMessageHostedContentId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
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

// ParseMeJoinedTeamChannelMessageHostedContentIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamChannelMessageHostedContentId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamChannelMessageHostedContentIDInsensitively(input string) (*MeJoinedTeamChannelMessageHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamChannelMessageHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamChannelMessageHostedContentId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
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

// ValidateMeJoinedTeamChannelMessageHostedContentID checks that 'input' can be parsed as a Me Joined Team Channel Message Hosted Content ID
func ValidateMeJoinedTeamChannelMessageHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamChannelMessageHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Channel Message Hosted Content ID
func (id MeJoinedTeamChannelMessageHostedContentId) ID() string {
	fmtString := "/me/joinedTeams/%s/channels/%s/messages/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId, id.ChatMessageId, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Channel Message Hosted Content ID
func (id MeJoinedTeamChannelMessageHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
		resourceids.StaticSegment("staticHostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Channel Message Hosted Content ID
func (id MeJoinedTeamChannelMessageHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("Me Joined Team Channel Message Hosted Content (%s)", strings.Join(components, "\n"))
}
