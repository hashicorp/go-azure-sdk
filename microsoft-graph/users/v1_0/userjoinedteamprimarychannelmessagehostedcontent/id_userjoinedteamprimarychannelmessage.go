package userjoinedteamprimarychannelmessagehostedcontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamPrimaryChannelMessageId{}

// UserJoinedTeamPrimaryChannelMessageId is a struct representing the Resource ID for a User Joined Team Primary Channel Message
type UserJoinedTeamPrimaryChannelMessageId struct {
	UserId        string
	TeamId        string
	ChatMessageId string
}

// NewUserJoinedTeamPrimaryChannelMessageID returns a new UserJoinedTeamPrimaryChannelMessageId struct
func NewUserJoinedTeamPrimaryChannelMessageID(userId string, teamId string, chatMessageId string) UserJoinedTeamPrimaryChannelMessageId {
	return UserJoinedTeamPrimaryChannelMessageId{
		UserId:        userId,
		TeamId:        teamId,
		ChatMessageId: chatMessageId,
	}
}

// ParseUserJoinedTeamPrimaryChannelMessageID parses 'input' into a UserJoinedTeamPrimaryChannelMessageId
func ParseUserJoinedTeamPrimaryChannelMessageID(input string) (*UserJoinedTeamPrimaryChannelMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamPrimaryChannelMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamPrimaryChannelMessageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamPrimaryChannelMessageIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamPrimaryChannelMessageId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamPrimaryChannelMessageIDInsensitively(input string) (*UserJoinedTeamPrimaryChannelMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamPrimaryChannelMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamPrimaryChannelMessageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamPrimaryChannelMessageID checks that 'input' can be parsed as a User Joined Team Primary Channel Message ID
func ValidateUserJoinedTeamPrimaryChannelMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamPrimaryChannelMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Primary Channel Message ID
func (id UserJoinedTeamPrimaryChannelMessageId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/primaryChannel/messages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChatMessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Primary Channel Message ID
func (id UserJoinedTeamPrimaryChannelMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Primary Channel Message ID
func (id UserJoinedTeamPrimaryChannelMessageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
	}
	return fmt.Sprintf("User Joined Team Primary Channel Message (%s)", strings.Join(components, "\n"))
}
