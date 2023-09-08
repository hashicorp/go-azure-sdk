package mejoinedteamchannelsharedwithteamallowedmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamChannelSharedWithTeamId{}

// MeJoinedTeamChannelSharedWithTeamId is a struct representing the Resource ID for a Me Joined Team Channel Shared With Team
type MeJoinedTeamChannelSharedWithTeamId struct {
	TeamId                      string
	ChannelId                   string
	SharedWithChannelTeamInfoId string
}

// NewMeJoinedTeamChannelSharedWithTeamID returns a new MeJoinedTeamChannelSharedWithTeamId struct
func NewMeJoinedTeamChannelSharedWithTeamID(teamId string, channelId string, sharedWithChannelTeamInfoId string) MeJoinedTeamChannelSharedWithTeamId {
	return MeJoinedTeamChannelSharedWithTeamId{
		TeamId:                      teamId,
		ChannelId:                   channelId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
	}
}

// ParseMeJoinedTeamChannelSharedWithTeamID parses 'input' into a MeJoinedTeamChannelSharedWithTeamId
func ParseMeJoinedTeamChannelSharedWithTeamID(input string) (*MeJoinedTeamChannelSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamChannelSharedWithTeamId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamChannelSharedWithTeamId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamChannelSharedWithTeamIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamChannelSharedWithTeamId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamChannelSharedWithTeamIDInsensitively(input string) (*MeJoinedTeamChannelSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamChannelSharedWithTeamId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamChannelSharedWithTeamId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamChannelSharedWithTeamID checks that 'input' can be parsed as a Me Joined Team Channel Shared With Team ID
func ValidateMeJoinedTeamChannelSharedWithTeamID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamChannelSharedWithTeamID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Channel Shared With Team ID
func (id MeJoinedTeamChannelSharedWithTeamId) ID() string {
	fmtString := "/me/joinedTeams/%s/channels/%s/sharedWithTeams/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId, id.SharedWithChannelTeamInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Channel Shared With Team ID
func (id MeJoinedTeamChannelSharedWithTeamId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
		resourceids.StaticSegment("staticSharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Channel Shared With Team ID
func (id MeJoinedTeamChannelSharedWithTeamId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
	}
	return fmt.Sprintf("Me Joined Team Channel Shared With Team (%s)", strings.Join(components, "\n"))
}
