package userjoinedteammember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamMemberId{}

// UserJoinedTeamMemberId is a struct representing the Resource ID for a User Joined Team Member
type UserJoinedTeamMemberId struct {
	UserId               string
	TeamId               string
	ConversationMemberId string
}

// NewUserJoinedTeamMemberID returns a new UserJoinedTeamMemberId struct
func NewUserJoinedTeamMemberID(userId string, teamId string, conversationMemberId string) UserJoinedTeamMemberId {
	return UserJoinedTeamMemberId{
		UserId:               userId,
		TeamId:               teamId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseUserJoinedTeamMemberID parses 'input' into a UserJoinedTeamMemberId
func ParseUserJoinedTeamMemberID(input string) (*UserJoinedTeamMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamMemberId{}

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

// ParseUserJoinedTeamMemberIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamMemberId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamMemberIDInsensitively(input string) (*UserJoinedTeamMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamMemberId{}

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

// ValidateUserJoinedTeamMemberID checks that 'input' can be parsed as a User Joined Team Member ID
func ValidateUserJoinedTeamMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Member ID
func (id UserJoinedTeamMemberId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/members/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Member ID
func (id UserJoinedTeamMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticMembers", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Member ID
func (id UserJoinedTeamMemberId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("User Joined Team Member (%s)", strings.Join(components, "\n"))
}
