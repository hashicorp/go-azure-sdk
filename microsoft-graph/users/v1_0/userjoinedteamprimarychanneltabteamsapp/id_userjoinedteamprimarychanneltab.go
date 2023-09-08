package userjoinedteamprimarychanneltabteamsapp

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamPrimaryChannelTabId{}

// UserJoinedTeamPrimaryChannelTabId is a struct representing the Resource ID for a User Joined Team Primary Channel Tab
type UserJoinedTeamPrimaryChannelTabId struct {
	UserId     string
	TeamId     string
	TeamsTabId string
}

// NewUserJoinedTeamPrimaryChannelTabID returns a new UserJoinedTeamPrimaryChannelTabId struct
func NewUserJoinedTeamPrimaryChannelTabID(userId string, teamId string, teamsTabId string) UserJoinedTeamPrimaryChannelTabId {
	return UserJoinedTeamPrimaryChannelTabId{
		UserId:     userId,
		TeamId:     teamId,
		TeamsTabId: teamsTabId,
	}
}

// ParseUserJoinedTeamPrimaryChannelTabID parses 'input' into a UserJoinedTeamPrimaryChannelTabId
func ParseUserJoinedTeamPrimaryChannelTabID(input string) (*UserJoinedTeamPrimaryChannelTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamPrimaryChannelTabId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamPrimaryChannelTabId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TeamsTabId, ok = parsed.Parsed["teamsTabId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamPrimaryChannelTabIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamPrimaryChannelTabId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamPrimaryChannelTabIDInsensitively(input string) (*UserJoinedTeamPrimaryChannelTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamPrimaryChannelTabId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamPrimaryChannelTabId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TeamsTabId, ok = parsed.Parsed["teamsTabId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamPrimaryChannelTabID checks that 'input' can be parsed as a User Joined Team Primary Channel Tab ID
func ValidateUserJoinedTeamPrimaryChannelTabID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamPrimaryChannelTabID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Primary Channel Tab ID
func (id UserJoinedTeamPrimaryChannelTabId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/primaryChannel/tabs/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.TeamsTabId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Primary Channel Tab ID
func (id UserJoinedTeamPrimaryChannelTabId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticTabs", "tabs", "tabs"),
		resourceids.UserSpecifiedSegment("teamsTabId", "teamsTabIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Primary Channel Tab ID
func (id UserJoinedTeamPrimaryChannelTabId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Teams Tab: %q", id.TeamsTabId),
	}
	return fmt.Sprintf("User Joined Team Primary Channel Tab (%s)", strings.Join(components, "\n"))
}
