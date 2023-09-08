package mejoinedteamprimarychanneltabteamsapp

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamPrimaryChannelTabId{}

// MeJoinedTeamPrimaryChannelTabId is a struct representing the Resource ID for a Me Joined Team Primary Channel Tab
type MeJoinedTeamPrimaryChannelTabId struct {
	TeamId     string
	TeamsTabId string
}

// NewMeJoinedTeamPrimaryChannelTabID returns a new MeJoinedTeamPrimaryChannelTabId struct
func NewMeJoinedTeamPrimaryChannelTabID(teamId string, teamsTabId string) MeJoinedTeamPrimaryChannelTabId {
	return MeJoinedTeamPrimaryChannelTabId{
		TeamId:     teamId,
		TeamsTabId: teamsTabId,
	}
}

// ParseMeJoinedTeamPrimaryChannelTabID parses 'input' into a MeJoinedTeamPrimaryChannelTabId
func ParseMeJoinedTeamPrimaryChannelTabID(input string) (*MeJoinedTeamPrimaryChannelTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamPrimaryChannelTabId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamPrimaryChannelTabId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TeamsTabId, ok = parsed.Parsed["teamsTabId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamPrimaryChannelTabIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamPrimaryChannelTabId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamPrimaryChannelTabIDInsensitively(input string) (*MeJoinedTeamPrimaryChannelTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamPrimaryChannelTabId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamPrimaryChannelTabId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TeamsTabId, ok = parsed.Parsed["teamsTabId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamPrimaryChannelTabID checks that 'input' can be parsed as a Me Joined Team Primary Channel Tab ID
func ValidateMeJoinedTeamPrimaryChannelTabID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamPrimaryChannelTabID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Primary Channel Tab ID
func (id MeJoinedTeamPrimaryChannelTabId) ID() string {
	fmtString := "/me/joinedTeams/%s/primaryChannel/tabs/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.TeamsTabId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Primary Channel Tab ID
func (id MeJoinedTeamPrimaryChannelTabId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticTabs", "tabs", "tabs"),
		resourceids.UserSpecifiedSegment("teamsTabId", "teamsTabIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Primary Channel Tab ID
func (id MeJoinedTeamPrimaryChannelTabId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Teams Tab: %q", id.TeamsTabId),
	}
	return fmt.Sprintf("Me Joined Team Primary Channel Tab (%s)", strings.Join(components, "\n"))
}
