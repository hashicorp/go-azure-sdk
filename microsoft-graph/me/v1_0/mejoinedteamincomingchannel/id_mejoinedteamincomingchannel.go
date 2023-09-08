package mejoinedteamincomingchannel

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamIncomingChannelId{}

// MeJoinedTeamIncomingChannelId is a struct representing the Resource ID for a Me Joined Team Incoming Channel
type MeJoinedTeamIncomingChannelId struct {
	TeamId    string
	ChannelId string
}

// NewMeJoinedTeamIncomingChannelID returns a new MeJoinedTeamIncomingChannelId struct
func NewMeJoinedTeamIncomingChannelID(teamId string, channelId string) MeJoinedTeamIncomingChannelId {
	return MeJoinedTeamIncomingChannelId{
		TeamId:    teamId,
		ChannelId: channelId,
	}
}

// ParseMeJoinedTeamIncomingChannelID parses 'input' into a MeJoinedTeamIncomingChannelId
func ParseMeJoinedTeamIncomingChannelID(input string) (*MeJoinedTeamIncomingChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamIncomingChannelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamIncomingChannelId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamIncomingChannelIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIncomingChannelId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIncomingChannelIDInsensitively(input string) (*MeJoinedTeamIncomingChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamIncomingChannelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamIncomingChannelId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamIncomingChannelID checks that 'input' can be parsed as a Me Joined Team Incoming Channel ID
func ValidateMeJoinedTeamIncomingChannelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIncomingChannelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Incoming Channel ID
func (id MeJoinedTeamIncomingChannelId) ID() string {
	fmtString := "/me/joinedTeams/%s/incomingChannels/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Incoming Channel ID
func (id MeJoinedTeamIncomingChannelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticIncomingChannels", "incomingChannels", "incomingChannels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Incoming Channel ID
func (id MeJoinedTeamIncomingChannelId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
	}
	return fmt.Sprintf("Me Joined Team Incoming Channel (%s)", strings.Join(components, "\n"))
}
