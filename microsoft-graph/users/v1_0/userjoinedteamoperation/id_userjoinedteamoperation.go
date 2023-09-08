package userjoinedteamoperation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamOperationId{}

// UserJoinedTeamOperationId is a struct representing the Resource ID for a User Joined Team Operation
type UserJoinedTeamOperationId struct {
	UserId                string
	TeamId                string
	TeamsAsyncOperationId string
}

// NewUserJoinedTeamOperationID returns a new UserJoinedTeamOperationId struct
func NewUserJoinedTeamOperationID(userId string, teamId string, teamsAsyncOperationId string) UserJoinedTeamOperationId {
	return UserJoinedTeamOperationId{
		UserId:                userId,
		TeamId:                teamId,
		TeamsAsyncOperationId: teamsAsyncOperationId,
	}
}

// ParseUserJoinedTeamOperationID parses 'input' into a UserJoinedTeamOperationId
func ParseUserJoinedTeamOperationID(input string) (*UserJoinedTeamOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamOperationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TeamsAsyncOperationId, ok = parsed.Parsed["teamsAsyncOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsAsyncOperationId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamOperationIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamOperationId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamOperationIDInsensitively(input string) (*UserJoinedTeamOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamOperationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TeamsAsyncOperationId, ok = parsed.Parsed["teamsAsyncOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsAsyncOperationId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamOperationID checks that 'input' can be parsed as a User Joined Team Operation ID
func ValidateUserJoinedTeamOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Operation ID
func (id UserJoinedTeamOperationId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.TeamsAsyncOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Operation ID
func (id UserJoinedTeamOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("teamsAsyncOperationId", "teamsAsyncOperationIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Operation ID
func (id UserJoinedTeamOperationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Teams Async Operation: %q", id.TeamsAsyncOperationId),
	}
	return fmt.Sprintf("User Joined Team Operation (%s)", strings.Join(components, "\n"))
}
