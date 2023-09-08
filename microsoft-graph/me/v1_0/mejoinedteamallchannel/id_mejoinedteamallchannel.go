package mejoinedteamallchannel

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamAllChannelId{}

// MeJoinedTeamAllChannelId is a struct representing the Resource ID for a Me Joined Team All Channel
type MeJoinedTeamAllChannelId struct {
	TeamId    string
	ChannelId string
}

// NewMeJoinedTeamAllChannelID returns a new MeJoinedTeamAllChannelId struct
func NewMeJoinedTeamAllChannelID(teamId string, channelId string) MeJoinedTeamAllChannelId {
	return MeJoinedTeamAllChannelId{
		TeamId:    teamId,
		ChannelId: channelId,
	}
}

// ParseMeJoinedTeamAllChannelID parses 'input' into a MeJoinedTeamAllChannelId
func ParseMeJoinedTeamAllChannelID(input string) (*MeJoinedTeamAllChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamAllChannelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamAllChannelId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamAllChannelIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamAllChannelId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamAllChannelIDInsensitively(input string) (*MeJoinedTeamAllChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamAllChannelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamAllChannelId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamAllChannelID checks that 'input' can be parsed as a Me Joined Team All Channel ID
func ValidateMeJoinedTeamAllChannelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamAllChannelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team All Channel ID
func (id MeJoinedTeamAllChannelId) ID() string {
	fmtString := "/me/joinedTeams/%s/allChannels/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team All Channel ID
func (id MeJoinedTeamAllChannelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticAllChannels", "allChannels", "allChannels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team All Channel ID
func (id MeJoinedTeamAllChannelId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
	}
	return fmt.Sprintf("Me Joined Team All Channel (%s)", strings.Join(components, "\n"))
}
