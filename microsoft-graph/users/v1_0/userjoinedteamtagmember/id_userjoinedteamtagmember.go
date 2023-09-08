package userjoinedteamtagmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamTagMemberId{}

// UserJoinedTeamTagMemberId is a struct representing the Resource ID for a User Joined Team Tag Member
type UserJoinedTeamTagMemberId struct {
	UserId              string
	TeamId              string
	TeamworkTagId       string
	TeamworkTagMemberId string
}

// NewUserJoinedTeamTagMemberID returns a new UserJoinedTeamTagMemberId struct
func NewUserJoinedTeamTagMemberID(userId string, teamId string, teamworkTagId string, teamworkTagMemberId string) UserJoinedTeamTagMemberId {
	return UserJoinedTeamTagMemberId{
		UserId:              userId,
		TeamId:              teamId,
		TeamworkTagId:       teamworkTagId,
		TeamworkTagMemberId: teamworkTagMemberId,
	}
}

// ParseUserJoinedTeamTagMemberID parses 'input' into a UserJoinedTeamTagMemberId
func ParseUserJoinedTeamTagMemberID(input string) (*UserJoinedTeamTagMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamTagMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamTagMemberId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TeamworkTagId, ok = parsed.Parsed["teamworkTagId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagId", *parsed)
	}

	if id.TeamworkTagMemberId, ok = parsed.Parsed["teamworkTagMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagMemberId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamTagMemberIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamTagMemberId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamTagMemberIDInsensitively(input string) (*UserJoinedTeamTagMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamTagMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamTagMemberId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TeamworkTagId, ok = parsed.Parsed["teamworkTagId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagId", *parsed)
	}

	if id.TeamworkTagMemberId, ok = parsed.Parsed["teamworkTagMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagMemberId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamTagMemberID checks that 'input' can be parsed as a User Joined Team Tag Member ID
func ValidateUserJoinedTeamTagMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamTagMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Tag Member ID
func (id UserJoinedTeamTagMemberId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/tags/%s/members/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.TeamworkTagId, id.TeamworkTagMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Tag Member ID
func (id UserJoinedTeamTagMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticTags", "tags", "tags"),
		resourceids.UserSpecifiedSegment("teamworkTagId", "teamworkTagIdValue"),
		resourceids.StaticSegment("staticMembers", "members", "members"),
		resourceids.UserSpecifiedSegment("teamworkTagMemberId", "teamworkTagMemberIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Tag Member ID
func (id UserJoinedTeamTagMemberId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Teamwork Tag: %q", id.TeamworkTagId),
		fmt.Sprintf("Teamwork Tag Member: %q", id.TeamworkTagMemberId),
	}
	return fmt.Sprintf("User Joined Team Tag Member (%s)", strings.Join(components, "\n"))
}
