package mejoinedteamchanneltab

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamChannelId{}

// MeJoinedTeamChannelId is a struct representing the Resource ID for a Me Joined Team Channel
type MeJoinedTeamChannelId struct {
	TeamId    string
	ChannelId string
}

// NewMeJoinedTeamChannelID returns a new MeJoinedTeamChannelId struct
func NewMeJoinedTeamChannelID(teamId string, channelId string) MeJoinedTeamChannelId {
	return MeJoinedTeamChannelId{
		TeamId:    teamId,
		ChannelId: channelId,
	}
}

// ParseMeJoinedTeamChannelID parses 'input' into a MeJoinedTeamChannelId
func ParseMeJoinedTeamChannelID(input string) (*MeJoinedTeamChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamChannelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamChannelId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamChannelIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamChannelId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamChannelIDInsensitively(input string) (*MeJoinedTeamChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamChannelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamChannelId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamChannelID checks that 'input' can be parsed as a Me Joined Team Channel ID
func ValidateMeJoinedTeamChannelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamChannelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Channel ID
func (id MeJoinedTeamChannelId) ID() string {
	fmtString := "/me/joinedTeams/%s/channels/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Channel ID
func (id MeJoinedTeamChannelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Channel ID
func (id MeJoinedTeamChannelId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
	}
	return fmt.Sprintf("Me Joined Team Channel (%s)", strings.Join(components, "\n"))
}
