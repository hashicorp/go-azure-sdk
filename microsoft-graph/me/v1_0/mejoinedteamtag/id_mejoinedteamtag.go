package mejoinedteamtag

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamTagId{}

// MeJoinedTeamTagId is a struct representing the Resource ID for a Me Joined Team Tag
type MeJoinedTeamTagId struct {
	TeamId        string
	TeamworkTagId string
}

// NewMeJoinedTeamTagID returns a new MeJoinedTeamTagId struct
func NewMeJoinedTeamTagID(teamId string, teamworkTagId string) MeJoinedTeamTagId {
	return MeJoinedTeamTagId{
		TeamId:        teamId,
		TeamworkTagId: teamworkTagId,
	}
}

// ParseMeJoinedTeamTagID parses 'input' into a MeJoinedTeamTagId
func ParseMeJoinedTeamTagID(input string) (*MeJoinedTeamTagId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamTagId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamTagId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TeamworkTagId, ok = parsed.Parsed["teamworkTagId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamTagIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamTagId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamTagIDInsensitively(input string) (*MeJoinedTeamTagId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamTagId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamTagId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TeamworkTagId, ok = parsed.Parsed["teamworkTagId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamTagID checks that 'input' can be parsed as a Me Joined Team Tag ID
func ValidateMeJoinedTeamTagID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamTagID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Tag ID
func (id MeJoinedTeamTagId) ID() string {
	fmtString := "/me/joinedTeams/%s/tags/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.TeamworkTagId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Tag ID
func (id MeJoinedTeamTagId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticTags", "tags", "tags"),
		resourceids.UserSpecifiedSegment("teamworkTagId", "teamworkTagIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Tag ID
func (id MeJoinedTeamTagId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Teamwork Tag: %q", id.TeamworkTagId),
	}
	return fmt.Sprintf("Me Joined Team Tag (%s)", strings.Join(components, "\n"))
}
