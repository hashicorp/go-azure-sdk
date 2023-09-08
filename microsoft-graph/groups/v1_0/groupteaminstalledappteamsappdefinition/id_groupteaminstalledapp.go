package groupteaminstalledappteamsappdefinition

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamInstalledAppId{}

// GroupTeamInstalledAppId is a struct representing the Resource ID for a Group Team Installed App
type GroupTeamInstalledAppId struct {
	GroupId                string
	TeamsAppInstallationId string
}

// NewGroupTeamInstalledAppID returns a new GroupTeamInstalledAppId struct
func NewGroupTeamInstalledAppID(groupId string, teamsAppInstallationId string) GroupTeamInstalledAppId {
	return GroupTeamInstalledAppId{
		GroupId:                groupId,
		TeamsAppInstallationId: teamsAppInstallationId,
	}
}

// ParseGroupTeamInstalledAppID parses 'input' into a GroupTeamInstalledAppId
func ParseGroupTeamInstalledAppID(input string) (*GroupTeamInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamInstalledAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamInstalledAppId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.TeamsAppInstallationId, ok = parsed.Parsed["teamsAppInstallationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsAppInstallationId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamInstalledAppIDInsensitively parses 'input' case-insensitively into a GroupTeamInstalledAppId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamInstalledAppIDInsensitively(input string) (*GroupTeamInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamInstalledAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamInstalledAppId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.TeamsAppInstallationId, ok = parsed.Parsed["teamsAppInstallationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsAppInstallationId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamInstalledAppID checks that 'input' can be parsed as a Group Team Installed App ID
func ValidateGroupTeamInstalledAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamInstalledAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Installed App ID
func (id GroupTeamInstalledAppId) ID() string {
	fmtString := "/groups/%s/team/installedApps/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.TeamsAppInstallationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Installed App ID
func (id GroupTeamInstalledAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticInstalledApps", "installedApps", "installedApps"),
		resourceids.UserSpecifiedSegment("teamsAppInstallationId", "teamsAppInstallationIdValue"),
	}
}

// String returns a human-readable description of this Group Team Installed App ID
func (id GroupTeamInstalledAppId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Teams App Installation: %q", id.TeamsAppInstallationId),
	}
	return fmt.Sprintf("Group Team Installed App (%s)", strings.Join(components, "\n"))
}
