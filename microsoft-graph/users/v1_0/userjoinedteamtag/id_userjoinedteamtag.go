package userjoinedteamtag

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamTagId{}

// UserJoinedTeamTagId is a struct representing the Resource ID for a User Joined Team Tag
type UserJoinedTeamTagId struct {
	UserId        string
	TeamId        string
	TeamworkTagId string
}

// NewUserJoinedTeamTagID returns a new UserJoinedTeamTagId struct
func NewUserJoinedTeamTagID(userId string, teamId string, teamworkTagId string) UserJoinedTeamTagId {
	return UserJoinedTeamTagId{
		UserId:        userId,
		TeamId:        teamId,
		TeamworkTagId: teamworkTagId,
	}
}

// ParseUserJoinedTeamTagID parses 'input' into a UserJoinedTeamTagId
func ParseUserJoinedTeamTagID(input string) (*UserJoinedTeamTagId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamTagId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamTagId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TeamworkTagId, ok = parsed.Parsed["teamworkTagId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamTagIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamTagId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamTagIDInsensitively(input string) (*UserJoinedTeamTagId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamTagId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamTagId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TeamworkTagId, ok = parsed.Parsed["teamworkTagId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamTagID checks that 'input' can be parsed as a User Joined Team Tag ID
func ValidateUserJoinedTeamTagID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamTagID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Tag ID
func (id UserJoinedTeamTagId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/tags/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.TeamworkTagId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Tag ID
func (id UserJoinedTeamTagId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticTags", "tags", "tags"),
		resourceids.UserSpecifiedSegment("teamworkTagId", "teamworkTagIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Tag ID
func (id UserJoinedTeamTagId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Teamwork Tag: %q", id.TeamworkTagId),
	}
	return fmt.Sprintf("User Joined Team Tag (%s)", strings.Join(components, "\n"))
}
