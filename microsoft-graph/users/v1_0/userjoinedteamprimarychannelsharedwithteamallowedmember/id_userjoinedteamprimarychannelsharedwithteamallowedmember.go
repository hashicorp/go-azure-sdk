package userjoinedteamprimarychannelsharedwithteamallowedmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId{}

// UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId is a struct representing the Resource ID for a User Joined Team Primary Channel Shared With Team Allowed Member
type UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId struct {
	UserId                      string
	TeamId                      string
	SharedWithChannelTeamInfoId string
	ConversationMemberId        string
}

// NewUserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberID returns a new UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId struct
func NewUserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberID(userId string, teamId string, sharedWithChannelTeamInfoId string, conversationMemberId string) UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId {
	return UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId{
		UserId:                      userId,
		TeamId:                      teamId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
		ConversationMemberId:        conversationMemberId,
	}
}

// ParseUserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberID parses 'input' into a UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId
func ParseUserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberID(input string) (*UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberIDInsensitively(input string) (*UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberID checks that 'input' can be parsed as a User Joined Team Primary Channel Shared With Team Allowed Member ID
func ValidateUserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Primary Channel Shared With Team Allowed Member ID
func (id UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/primaryChannel/sharedWithTeams/%s/allowedMembers/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.SharedWithChannelTeamInfoId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Primary Channel Shared With Team Allowed Member ID
func (id UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticSharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoIdValue"),
		resourceids.StaticSegment("staticAllowedMembers", "allowedMembers", "allowedMembers"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Primary Channel Shared With Team Allowed Member ID
func (id UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("User Joined Team Primary Channel Shared With Team Allowed Member (%s)", strings.Join(components, "\n"))
}
