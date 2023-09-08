package userjoinedteamscheduleshift

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamId{}

// UserJoinedTeamId is a struct representing the Resource ID for a User Joined Team
type UserJoinedTeamId struct {
	UserId string
	TeamId string
}

// NewUserJoinedTeamID returns a new UserJoinedTeamId struct
func NewUserJoinedTeamID(userId string, teamId string) UserJoinedTeamId {
	return UserJoinedTeamId{
		UserId: userId,
		TeamId: teamId,
	}
}

// ParseUserJoinedTeamID parses 'input' into a UserJoinedTeamId
func ParseUserJoinedTeamID(input string) (*UserJoinedTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamIDInsensitively(input string) (*UserJoinedTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamID checks that 'input' can be parsed as a User Joined Team ID
func ValidateUserJoinedTeamID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team ID
func (id UserJoinedTeamId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team ID
func (id UserJoinedTeamId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team ID
func (id UserJoinedTeamId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
	}
	return fmt.Sprintf("User Joined Team (%s)", strings.Join(components, "\n"))
}
