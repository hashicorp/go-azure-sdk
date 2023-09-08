package mejoinedteamprimarychannelsharedwithteamteam

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamPrimaryChannelSharedWithTeamId{}

// MeJoinedTeamPrimaryChannelSharedWithTeamId is a struct representing the Resource ID for a Me Joined Team Primary Channel Shared With Team
type MeJoinedTeamPrimaryChannelSharedWithTeamId struct {
	TeamId                      string
	SharedWithChannelTeamInfoId string
}

// NewMeJoinedTeamPrimaryChannelSharedWithTeamID returns a new MeJoinedTeamPrimaryChannelSharedWithTeamId struct
func NewMeJoinedTeamPrimaryChannelSharedWithTeamID(teamId string, sharedWithChannelTeamInfoId string) MeJoinedTeamPrimaryChannelSharedWithTeamId {
	return MeJoinedTeamPrimaryChannelSharedWithTeamId{
		TeamId:                      teamId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
	}
}

// ParseMeJoinedTeamPrimaryChannelSharedWithTeamID parses 'input' into a MeJoinedTeamPrimaryChannelSharedWithTeamId
func ParseMeJoinedTeamPrimaryChannelSharedWithTeamID(input string) (*MeJoinedTeamPrimaryChannelSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamPrimaryChannelSharedWithTeamId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamPrimaryChannelSharedWithTeamId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamPrimaryChannelSharedWithTeamIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamPrimaryChannelSharedWithTeamId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamPrimaryChannelSharedWithTeamIDInsensitively(input string) (*MeJoinedTeamPrimaryChannelSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamPrimaryChannelSharedWithTeamId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamPrimaryChannelSharedWithTeamId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamPrimaryChannelSharedWithTeamID checks that 'input' can be parsed as a Me Joined Team Primary Channel Shared With Team ID
func ValidateMeJoinedTeamPrimaryChannelSharedWithTeamID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamPrimaryChannelSharedWithTeamID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Primary Channel Shared With Team ID
func (id MeJoinedTeamPrimaryChannelSharedWithTeamId) ID() string {
	fmtString := "/me/joinedTeams/%s/primaryChannel/sharedWithTeams/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.SharedWithChannelTeamInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Primary Channel Shared With Team ID
func (id MeJoinedTeamPrimaryChannelSharedWithTeamId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticSharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Primary Channel Shared With Team ID
func (id MeJoinedTeamPrimaryChannelSharedWithTeamId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
	}
	return fmt.Sprintf("Me Joined Team Primary Channel Shared With Team (%s)", strings.Join(components, "\n"))
}
