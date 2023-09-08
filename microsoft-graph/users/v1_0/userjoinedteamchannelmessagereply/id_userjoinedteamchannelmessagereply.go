package userjoinedteamchannelmessagereply

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamChannelMessageReplyId{}

// UserJoinedTeamChannelMessageReplyId is a struct representing the Resource ID for a User Joined Team Channel Message Reply
type UserJoinedTeamChannelMessageReplyId struct {
	UserId         string
	TeamId         string
	ChannelId      string
	ChatMessageId  string
	ChatMessageId1 string
}

// NewUserJoinedTeamChannelMessageReplyID returns a new UserJoinedTeamChannelMessageReplyId struct
func NewUserJoinedTeamChannelMessageReplyID(userId string, teamId string, channelId string, chatMessageId string, chatMessageId1 string) UserJoinedTeamChannelMessageReplyId {
	return UserJoinedTeamChannelMessageReplyId{
		UserId:         userId,
		TeamId:         teamId,
		ChannelId:      channelId,
		ChatMessageId:  chatMessageId,
		ChatMessageId1: chatMessageId1,
	}
}

// ParseUserJoinedTeamChannelMessageReplyID parses 'input' into a UserJoinedTeamChannelMessageReplyId
func ParseUserJoinedTeamChannelMessageReplyID(input string) (*UserJoinedTeamChannelMessageReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamChannelMessageReplyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamChannelMessageReplyId{}

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

	if id.ChatMessageId1, ok = parsed.Parsed["chatMessageId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamChannelMessageReplyIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamChannelMessageReplyId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamChannelMessageReplyIDInsensitively(input string) (*UserJoinedTeamChannelMessageReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamChannelMessageReplyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamChannelMessageReplyId{}

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

	if id.ChatMessageId1, ok = parsed.Parsed["chatMessageId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamChannelMessageReplyID checks that 'input' can be parsed as a User Joined Team Channel Message Reply ID
func ValidateUserJoinedTeamChannelMessageReplyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamChannelMessageReplyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Channel Message Reply ID
func (id UserJoinedTeamChannelMessageReplyId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/channels/%s/messages/%s/replies/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId, id.ChatMessageId, id.ChatMessageId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Channel Message Reply ID
func (id UserJoinedTeamChannelMessageReplyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
		resourceids.StaticSegment("staticReplies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1Value"),
	}
}

// String returns a human-readable description of this User Joined Team Channel Message Reply ID
func (id UserJoinedTeamChannelMessageReplyId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
	}
	return fmt.Sprintf("User Joined Team Channel Message Reply (%s)", strings.Join(components, "\n"))
}
