package mejoinedteaminstalledappteamsappdefinition

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamInstalledAppId{}

// MeJoinedTeamInstalledAppId is a struct representing the Resource ID for a Me Joined Team Installed App
type MeJoinedTeamInstalledAppId struct {
	TeamId                 string
	TeamsAppInstallationId string
}

// NewMeJoinedTeamInstalledAppID returns a new MeJoinedTeamInstalledAppId struct
func NewMeJoinedTeamInstalledAppID(teamId string, teamsAppInstallationId string) MeJoinedTeamInstalledAppId {
	return MeJoinedTeamInstalledAppId{
		TeamId:                 teamId,
		TeamsAppInstallationId: teamsAppInstallationId,
	}
}

// ParseMeJoinedTeamInstalledAppID parses 'input' into a MeJoinedTeamInstalledAppId
func ParseMeJoinedTeamInstalledAppID(input string) (*MeJoinedTeamInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamInstalledAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamInstalledAppId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TeamsAppInstallationId, ok = parsed.Parsed["teamsAppInstallationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsAppInstallationId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamInstalledAppIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamInstalledAppId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamInstalledAppIDInsensitively(input string) (*MeJoinedTeamInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamInstalledAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamInstalledAppId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TeamsAppInstallationId, ok = parsed.Parsed["teamsAppInstallationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsAppInstallationId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamInstalledAppID checks that 'input' can be parsed as a Me Joined Team Installed App ID
func ValidateMeJoinedTeamInstalledAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamInstalledAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Installed App ID
func (id MeJoinedTeamInstalledAppId) ID() string {
	fmtString := "/me/joinedTeams/%s/installedApps/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.TeamsAppInstallationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Installed App ID
func (id MeJoinedTeamInstalledAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticInstalledApps", "installedApps", "installedApps"),
		resourceids.UserSpecifiedSegment("teamsAppInstallationId", "teamsAppInstallationIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Installed App ID
func (id MeJoinedTeamInstalledAppId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Teams App Installation: %q", id.TeamsAppInstallationId),
	}
	return fmt.Sprintf("Me Joined Team Installed App (%s)", strings.Join(components, "\n"))
}
