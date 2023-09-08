package userjoinedteamprimarychannelmessagereplyhostedcontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamPrimaryChannelMessageReplyId{}

// UserJoinedTeamPrimaryChannelMessageReplyId is a struct representing the Resource ID for a User Joined Team Primary Channel Message Reply
type UserJoinedTeamPrimaryChannelMessageReplyId struct {
	UserId         string
	TeamId         string
	ChatMessageId  string
	ChatMessageId1 string
}

// NewUserJoinedTeamPrimaryChannelMessageReplyID returns a new UserJoinedTeamPrimaryChannelMessageReplyId struct
func NewUserJoinedTeamPrimaryChannelMessageReplyID(userId string, teamId string, chatMessageId string, chatMessageId1 string) UserJoinedTeamPrimaryChannelMessageReplyId {
	return UserJoinedTeamPrimaryChannelMessageReplyId{
		UserId:         userId,
		TeamId:         teamId,
		ChatMessageId:  chatMessageId,
		ChatMessageId1: chatMessageId1,
	}
}

// ParseUserJoinedTeamPrimaryChannelMessageReplyID parses 'input' into a UserJoinedTeamPrimaryChannelMessageReplyId
func ParseUserJoinedTeamPrimaryChannelMessageReplyID(input string) (*UserJoinedTeamPrimaryChannelMessageReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamPrimaryChannelMessageReplyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamPrimaryChannelMessageReplyId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageId1, ok = parsed.Parsed["chatMessageId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamPrimaryChannelMessageReplyIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamPrimaryChannelMessageReplyId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamPrimaryChannelMessageReplyIDInsensitively(input string) (*UserJoinedTeamPrimaryChannelMessageReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamPrimaryChannelMessageReplyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamPrimaryChannelMessageReplyId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageId1, ok = parsed.Parsed["chatMessageId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamPrimaryChannelMessageReplyID checks that 'input' can be parsed as a User Joined Team Primary Channel Message Reply ID
func ValidateUserJoinedTeamPrimaryChannelMessageReplyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamPrimaryChannelMessageReplyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Primary Channel Message Reply ID
func (id UserJoinedTeamPrimaryChannelMessageReplyId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/primaryChannel/messages/%s/replies/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChatMessageId, id.ChatMessageId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Primary Channel Message Reply ID
func (id UserJoinedTeamPrimaryChannelMessageReplyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
		resourceids.StaticSegment("staticReplies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1Value"),
	}
}

// String returns a human-readable description of this User Joined Team Primary Channel Message Reply ID
func (id UserJoinedTeamPrimaryChannelMessageReplyId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
	}
	return fmt.Sprintf("User Joined Team Primary Channel Message Reply (%s)", strings.Join(components, "\n"))
}
