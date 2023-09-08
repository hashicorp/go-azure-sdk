package mejoinedteamchanneltabteamsapp

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamChannelTabId{}

// MeJoinedTeamChannelTabId is a struct representing the Resource ID for a Me Joined Team Channel Tab
type MeJoinedTeamChannelTabId struct {
	TeamId     string
	ChannelId  string
	TeamsTabId string
}

// NewMeJoinedTeamChannelTabID returns a new MeJoinedTeamChannelTabId struct
func NewMeJoinedTeamChannelTabID(teamId string, channelId string, teamsTabId string) MeJoinedTeamChannelTabId {
	return MeJoinedTeamChannelTabId{
		TeamId:     teamId,
		ChannelId:  channelId,
		TeamsTabId: teamsTabId,
	}
}

// ParseMeJoinedTeamChannelTabID parses 'input' into a MeJoinedTeamChannelTabId
func ParseMeJoinedTeamChannelTabID(input string) (*MeJoinedTeamChannelTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamChannelTabId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamChannelTabId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.TeamsTabId, ok = parsed.Parsed["teamsTabId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamChannelTabIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamChannelTabId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamChannelTabIDInsensitively(input string) (*MeJoinedTeamChannelTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamChannelTabId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamChannelTabId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.TeamsTabId, ok = parsed.Parsed["teamsTabId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamChannelTabID checks that 'input' can be parsed as a Me Joined Team Channel Tab ID
func ValidateMeJoinedTeamChannelTabID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamChannelTabID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Channel Tab ID
func (id MeJoinedTeamChannelTabId) ID() string {
	fmtString := "/me/joinedTeams/%s/channels/%s/tabs/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId, id.TeamsTabId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Channel Tab ID
func (id MeJoinedTeamChannelTabId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
		resourceids.StaticSegment("staticTabs", "tabs", "tabs"),
		resourceids.UserSpecifiedSegment("teamsTabId", "teamsTabIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Channel Tab ID
func (id MeJoinedTeamChannelTabId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Teams Tab: %q", id.TeamsTabId),
	}
	return fmt.Sprintf("Me Joined Team Channel Tab (%s)", strings.Join(components, "\n"))
}
