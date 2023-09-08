package userjoinedteamchannelmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamChannelMemberId{}

// UserJoinedTeamChannelMemberId is a struct representing the Resource ID for a User Joined Team Channel Member
type UserJoinedTeamChannelMemberId struct {
	UserId               string
	TeamId               string
	ChannelId            string
	ConversationMemberId string
}

// NewUserJoinedTeamChannelMemberID returns a new UserJoinedTeamChannelMemberId struct
func NewUserJoinedTeamChannelMemberID(userId string, teamId string, channelId string, conversationMemberId string) UserJoinedTeamChannelMemberId {
	return UserJoinedTeamChannelMemberId{
		UserId:               userId,
		TeamId:               teamId,
		ChannelId:            channelId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseUserJoinedTeamChannelMemberID parses 'input' into a UserJoinedTeamChannelMemberId
func ParseUserJoinedTeamChannelMemberID(input string) (*UserJoinedTeamChannelMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamChannelMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamChannelMemberId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamChannelMemberIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamChannelMemberId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamChannelMemberIDInsensitively(input string) (*UserJoinedTeamChannelMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamChannelMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamChannelMemberId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamChannelMemberID checks that 'input' can be parsed as a User Joined Team Channel Member ID
func ValidateUserJoinedTeamChannelMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamChannelMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Channel Member ID
func (id UserJoinedTeamChannelMemberId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/channels/%s/members/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Channel Member ID
func (id UserJoinedTeamChannelMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
		resourceids.StaticSegment("staticMembers", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Channel Member ID
func (id UserJoinedTeamChannelMemberId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("User Joined Team Channel Member (%s)", strings.Join(components, "\n"))
}
