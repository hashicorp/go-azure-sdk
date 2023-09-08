package userjoinedteamprimarychannelmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamPrimaryChannelMemberId{}

// UserJoinedTeamPrimaryChannelMemberId is a struct representing the Resource ID for a User Joined Team Primary Channel Member
type UserJoinedTeamPrimaryChannelMemberId struct {
	UserId               string
	TeamId               string
	ConversationMemberId string
}

// NewUserJoinedTeamPrimaryChannelMemberID returns a new UserJoinedTeamPrimaryChannelMemberId struct
func NewUserJoinedTeamPrimaryChannelMemberID(userId string, teamId string, conversationMemberId string) UserJoinedTeamPrimaryChannelMemberId {
	return UserJoinedTeamPrimaryChannelMemberId{
		UserId:               userId,
		TeamId:               teamId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseUserJoinedTeamPrimaryChannelMemberID parses 'input' into a UserJoinedTeamPrimaryChannelMemberId
func ParseUserJoinedTeamPrimaryChannelMemberID(input string) (*UserJoinedTeamPrimaryChannelMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamPrimaryChannelMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamPrimaryChannelMemberId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamPrimaryChannelMemberIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamPrimaryChannelMemberId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamPrimaryChannelMemberIDInsensitively(input string) (*UserJoinedTeamPrimaryChannelMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamPrimaryChannelMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamPrimaryChannelMemberId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamPrimaryChannelMemberID checks that 'input' can be parsed as a User Joined Team Primary Channel Member ID
func ValidateUserJoinedTeamPrimaryChannelMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamPrimaryChannelMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Primary Channel Member ID
func (id UserJoinedTeamPrimaryChannelMemberId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/primaryChannel/members/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Primary Channel Member ID
func (id UserJoinedTeamPrimaryChannelMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticMembers", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Primary Channel Member ID
func (id UserJoinedTeamPrimaryChannelMemberId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("User Joined Team Primary Channel Member (%s)", strings.Join(components, "\n"))
}
