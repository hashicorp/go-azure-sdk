package mejoinedteamtagmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamTagMemberId{}

// MeJoinedTeamTagMemberId is a struct representing the Resource ID for a Me Joined Team Tag Member
type MeJoinedTeamTagMemberId struct {
	TeamId              string
	TeamworkTagId       string
	TeamworkTagMemberId string
}

// NewMeJoinedTeamTagMemberID returns a new MeJoinedTeamTagMemberId struct
func NewMeJoinedTeamTagMemberID(teamId string, teamworkTagId string, teamworkTagMemberId string) MeJoinedTeamTagMemberId {
	return MeJoinedTeamTagMemberId{
		TeamId:              teamId,
		TeamworkTagId:       teamworkTagId,
		TeamworkTagMemberId: teamworkTagMemberId,
	}
}

// ParseMeJoinedTeamTagMemberID parses 'input' into a MeJoinedTeamTagMemberId
func ParseMeJoinedTeamTagMemberID(input string) (*MeJoinedTeamTagMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamTagMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamTagMemberId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TeamworkTagId, ok = parsed.Parsed["teamworkTagId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagId", *parsed)
	}

	if id.TeamworkTagMemberId, ok = parsed.Parsed["teamworkTagMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagMemberId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamTagMemberIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamTagMemberId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamTagMemberIDInsensitively(input string) (*MeJoinedTeamTagMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamTagMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamTagMemberId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TeamworkTagId, ok = parsed.Parsed["teamworkTagId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagId", *parsed)
	}

	if id.TeamworkTagMemberId, ok = parsed.Parsed["teamworkTagMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagMemberId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamTagMemberID checks that 'input' can be parsed as a Me Joined Team Tag Member ID
func ValidateMeJoinedTeamTagMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamTagMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Tag Member ID
func (id MeJoinedTeamTagMemberId) ID() string {
	fmtString := "/me/joinedTeams/%s/tags/%s/members/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.TeamworkTagId, id.TeamworkTagMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Tag Member ID
func (id MeJoinedTeamTagMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticTags", "tags", "tags"),
		resourceids.UserSpecifiedSegment("teamworkTagId", "teamworkTagIdValue"),
		resourceids.StaticSegment("staticMembers", "members", "members"),
		resourceids.UserSpecifiedSegment("teamworkTagMemberId", "teamworkTagMemberIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Tag Member ID
func (id MeJoinedTeamTagMemberId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Teamwork Tag: %q", id.TeamworkTagId),
		fmt.Sprintf("Teamwork Tag Member: %q", id.TeamworkTagMemberId),
	}
	return fmt.Sprintf("Me Joined Team Tag Member (%s)", strings.Join(components, "\n"))
}
