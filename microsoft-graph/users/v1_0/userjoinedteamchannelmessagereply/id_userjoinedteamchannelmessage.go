package userjoinedteamchannelmessagereply

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamChannelMessageId{}

// UserJoinedTeamChannelMessageId is a struct representing the Resource ID for a User Joined Team Channel Message
type UserJoinedTeamChannelMessageId struct {
	UserId        string
	TeamId        string
	ChannelId     string
	ChatMessageId string
}

// NewUserJoinedTeamChannelMessageID returns a new UserJoinedTeamChannelMessageId struct
func NewUserJoinedTeamChannelMessageID(userId string, teamId string, channelId string, chatMessageId string) UserJoinedTeamChannelMessageId {
	return UserJoinedTeamChannelMessageId{
		UserId:        userId,
		TeamId:        teamId,
		ChannelId:     channelId,
		ChatMessageId: chatMessageId,
	}
}

// ParseUserJoinedTeamChannelMessageID parses 'input' into a UserJoinedTeamChannelMessageId
func ParseUserJoinedTeamChannelMessageID(input string) (*UserJoinedTeamChannelMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamChannelMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamChannelMessageId{}

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

	return &id, nil
}

// ParseUserJoinedTeamChannelMessageIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamChannelMessageId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamChannelMessageIDInsensitively(input string) (*UserJoinedTeamChannelMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamChannelMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamChannelMessageId{}

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

	return &id, nil
}

// ValidateUserJoinedTeamChannelMessageID checks that 'input' can be parsed as a User Joined Team Channel Message ID
func ValidateUserJoinedTeamChannelMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamChannelMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Channel Message ID
func (id UserJoinedTeamChannelMessageId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/channels/%s/messages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId, id.ChatMessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Channel Message ID
func (id UserJoinedTeamChannelMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Channel Message ID
func (id UserJoinedTeamChannelMessageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
	}
	return fmt.Sprintf("User Joined Team Channel Message (%s)", strings.Join(components, "\n"))
}
