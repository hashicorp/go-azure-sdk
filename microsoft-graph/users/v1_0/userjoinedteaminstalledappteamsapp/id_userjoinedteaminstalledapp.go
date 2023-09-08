package userjoinedteaminstalledappteamsapp

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamInstalledAppId{}

// UserJoinedTeamInstalledAppId is a struct representing the Resource ID for a User Joined Team Installed App
type UserJoinedTeamInstalledAppId struct {
	UserId                 string
	TeamId                 string
	TeamsAppInstallationId string
}

// NewUserJoinedTeamInstalledAppID returns a new UserJoinedTeamInstalledAppId struct
func NewUserJoinedTeamInstalledAppID(userId string, teamId string, teamsAppInstallationId string) UserJoinedTeamInstalledAppId {
	return UserJoinedTeamInstalledAppId{
		UserId:                 userId,
		TeamId:                 teamId,
		TeamsAppInstallationId: teamsAppInstallationId,
	}
}

// ParseUserJoinedTeamInstalledAppID parses 'input' into a UserJoinedTeamInstalledAppId
func ParseUserJoinedTeamInstalledAppID(input string) (*UserJoinedTeamInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamInstalledAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamInstalledAppId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TeamsAppInstallationId, ok = parsed.Parsed["teamsAppInstallationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsAppInstallationId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamInstalledAppIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamInstalledAppId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamInstalledAppIDInsensitively(input string) (*UserJoinedTeamInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamInstalledAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamInstalledAppId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TeamsAppInstallationId, ok = parsed.Parsed["teamsAppInstallationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsAppInstallationId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamInstalledAppID checks that 'input' can be parsed as a User Joined Team Installed App ID
func ValidateUserJoinedTeamInstalledAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamInstalledAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Installed App ID
func (id UserJoinedTeamInstalledAppId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/installedApps/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.TeamsAppInstallationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Installed App ID
func (id UserJoinedTeamInstalledAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticInstalledApps", "installedApps", "installedApps"),
		resourceids.UserSpecifiedSegment("teamsAppInstallationId", "teamsAppInstallationIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Installed App ID
func (id UserJoinedTeamInstalledAppId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Teams App Installation: %q", id.TeamsAppInstallationId),
	}
	return fmt.Sprintf("User Joined Team Installed App (%s)", strings.Join(components, "\n"))
}
