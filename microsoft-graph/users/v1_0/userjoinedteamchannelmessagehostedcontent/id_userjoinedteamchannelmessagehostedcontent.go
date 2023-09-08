package userjoinedteamchannelmessagehostedcontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamChannelMessageHostedContentId{}

// UserJoinedTeamChannelMessageHostedContentId is a struct representing the Resource ID for a User Joined Team Channel Message Hosted Content
type UserJoinedTeamChannelMessageHostedContentId struct {
	UserId                     string
	TeamId                     string
	ChannelId                  string
	ChatMessageId              string
	ChatMessageHostedContentId string
}

// NewUserJoinedTeamChannelMessageHostedContentID returns a new UserJoinedTeamChannelMessageHostedContentId struct
func NewUserJoinedTeamChannelMessageHostedContentID(userId string, teamId string, channelId string, chatMessageId string, chatMessageHostedContentId string) UserJoinedTeamChannelMessageHostedContentId {
	return UserJoinedTeamChannelMessageHostedContentId{
		UserId:                     userId,
		TeamId:                     teamId,
		ChannelId:                  channelId,
		ChatMessageId:              chatMessageId,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseUserJoinedTeamChannelMessageHostedContentID parses 'input' into a UserJoinedTeamChannelMessageHostedContentId
func ParseUserJoinedTeamChannelMessageHostedContentID(input string) (*UserJoinedTeamChannelMessageHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamChannelMessageHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamChannelMessageHostedContentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ParseUserJoinedTeamChannelMessageHostedContentIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamChannelMessageHostedContentId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamChannelMessageHostedContentIDInsensitively(input string) (*UserJoinedTeamChannelMessageHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamChannelMessageHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamChannelMessageHostedContentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ValidateUserJoinedTeamChannelMessageHostedContentID checks that 'input' can be parsed as a User Joined Team Channel Message Hosted Content ID
func ValidateUserJoinedTeamChannelMessageHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamChannelMessageHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Channel Message Hosted Content ID
func (id UserJoinedTeamChannelMessageHostedContentId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/channels/%s/messages/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId, id.ChatMessageId, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Channel Message Hosted Content ID
func (id UserJoinedTeamChannelMessageHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
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

// String returns a human-readable description of this User Joined Team Channel Message Hosted Content ID
func (id UserJoinedTeamChannelMessageHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("User Joined Team Channel Message Hosted Content (%s)", strings.Join(components, "\n"))
}
