package userjoinedteamchannelsharedwithteamallowedmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamChannelSharedWithTeamAllowedMemberId{}

// UserJoinedTeamChannelSharedWithTeamAllowedMemberId is a struct representing the Resource ID for a User Joined Team Channel Shared With Team Allowed Member
type UserJoinedTeamChannelSharedWithTeamAllowedMemberId struct {
	UserId                      string
	TeamId                      string
	ChannelId                   string
	SharedWithChannelTeamInfoId string
	ConversationMemberId        string
}

// NewUserJoinedTeamChannelSharedWithTeamAllowedMemberID returns a new UserJoinedTeamChannelSharedWithTeamAllowedMemberId struct
func NewUserJoinedTeamChannelSharedWithTeamAllowedMemberID(userId string, teamId string, channelId string, sharedWithChannelTeamInfoId string, conversationMemberId string) UserJoinedTeamChannelSharedWithTeamAllowedMemberId {
	return UserJoinedTeamChannelSharedWithTeamAllowedMemberId{
		UserId:                      userId,
		TeamId:                      teamId,
		ChannelId:                   channelId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
		ConversationMemberId:        conversationMemberId,
	}
}

// ParseUserJoinedTeamChannelSharedWithTeamAllowedMemberID parses 'input' into a UserJoinedTeamChannelSharedWithTeamAllowedMemberId
func ParseUserJoinedTeamChannelSharedWithTeamAllowedMemberID(input string) (*UserJoinedTeamChannelSharedWithTeamAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamChannelSharedWithTeamAllowedMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamChannelSharedWithTeamAllowedMemberId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamChannelSharedWithTeamAllowedMemberIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamChannelSharedWithTeamAllowedMemberId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamChannelSharedWithTeamAllowedMemberIDInsensitively(input string) (*UserJoinedTeamChannelSharedWithTeamAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamChannelSharedWithTeamAllowedMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamChannelSharedWithTeamAllowedMemberId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamChannelSharedWithTeamAllowedMemberID checks that 'input' can be parsed as a User Joined Team Channel Shared With Team Allowed Member ID
func ValidateUserJoinedTeamChannelSharedWithTeamAllowedMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamChannelSharedWithTeamAllowedMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Channel Shared With Team Allowed Member ID
func (id UserJoinedTeamChannelSharedWithTeamAllowedMemberId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/channels/%s/sharedWithTeams/%s/allowedMembers/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId, id.SharedWithChannelTeamInfoId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Channel Shared With Team Allowed Member ID
func (id UserJoinedTeamChannelSharedWithTeamAllowedMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
		resourceids.StaticSegment("staticSharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoIdValue"),
		resourceids.StaticSegment("staticAllowedMembers", "allowedMembers", "allowedMembers"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Channel Shared With Team Allowed Member ID
func (id UserJoinedTeamChannelSharedWithTeamAllowedMemberId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("User Joined Team Channel Shared With Team Allowed Member (%s)", strings.Join(components, "\n"))
}
