package userjoinedteamchannelmessagereplyhostedcontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamChannelMessageReplyHostedContentId{}

// UserJoinedTeamChannelMessageReplyHostedContentId is a struct representing the Resource ID for a User Joined Team Channel Message Reply Hosted Content
type UserJoinedTeamChannelMessageReplyHostedContentId struct {
	UserId                     string
	TeamId                     string
	ChannelId                  string
	ChatMessageId              string
	ChatMessageId1             string
	ChatMessageHostedContentId string
}

// NewUserJoinedTeamChannelMessageReplyHostedContentID returns a new UserJoinedTeamChannelMessageReplyHostedContentId struct
func NewUserJoinedTeamChannelMessageReplyHostedContentID(userId string, teamId string, channelId string, chatMessageId string, chatMessageId1 string, chatMessageHostedContentId string) UserJoinedTeamChannelMessageReplyHostedContentId {
	return UserJoinedTeamChannelMessageReplyHostedContentId{
		UserId:                     userId,
		TeamId:                     teamId,
		ChannelId:                  channelId,
		ChatMessageId:              chatMessageId,
		ChatMessageId1:             chatMessageId1,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseUserJoinedTeamChannelMessageReplyHostedContentID parses 'input' into a UserJoinedTeamChannelMessageReplyHostedContentId
func ParseUserJoinedTeamChannelMessageReplyHostedContentID(input string) (*UserJoinedTeamChannelMessageReplyHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamChannelMessageReplyHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamChannelMessageReplyHostedContentId{}

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

	if id.ChatMessageHostedContentId, ok = parsed.Parsed["chatMessageHostedContentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamChannelMessageReplyHostedContentIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamChannelMessageReplyHostedContentId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamChannelMessageReplyHostedContentIDInsensitively(input string) (*UserJoinedTeamChannelMessageReplyHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamChannelMessageReplyHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamChannelMessageReplyHostedContentId{}

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

	if id.ChatMessageHostedContentId, ok = parsed.Parsed["chatMessageHostedContentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamChannelMessageReplyHostedContentID checks that 'input' can be parsed as a User Joined Team Channel Message Reply Hosted Content ID
func ValidateUserJoinedTeamChannelMessageReplyHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamChannelMessageReplyHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Channel Message Reply Hosted Content ID
func (id UserJoinedTeamChannelMessageReplyHostedContentId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/channels/%s/messages/%s/replies/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId, id.ChatMessageId, id.ChatMessageId1, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Channel Message Reply Hosted Content ID
func (id UserJoinedTeamChannelMessageReplyHostedContentId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticHostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Channel Message Reply Hosted Content ID
func (id UserJoinedTeamChannelMessageReplyHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("User Joined Team Channel Message Reply Hosted Content (%s)", strings.Join(components, "\n"))
}
