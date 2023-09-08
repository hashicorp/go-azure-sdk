package groupteamtagmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamTagMemberId{}

// GroupTeamTagMemberId is a struct representing the Resource ID for a Group Team Tag Member
type GroupTeamTagMemberId struct {
	GroupId             string
	TeamworkTagId       string
	TeamworkTagMemberId string
}

// NewGroupTeamTagMemberID returns a new GroupTeamTagMemberId struct
func NewGroupTeamTagMemberID(groupId string, teamworkTagId string, teamworkTagMemberId string) GroupTeamTagMemberId {
	return GroupTeamTagMemberId{
		GroupId:             groupId,
		TeamworkTagId:       teamworkTagId,
		TeamworkTagMemberId: teamworkTagMemberId,
	}
}

// ParseGroupTeamTagMemberID parses 'input' into a GroupTeamTagMemberId
func ParseGroupTeamTagMemberID(input string) (*GroupTeamTagMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamTagMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamTagMemberId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.TeamworkTagId, ok = parsed.Parsed["teamworkTagId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagId", *parsed)
	}

	if id.TeamworkTagMemberId, ok = parsed.Parsed["teamworkTagMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagMemberId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamTagMemberIDInsensitively parses 'input' case-insensitively into a GroupTeamTagMemberId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamTagMemberIDInsensitively(input string) (*GroupTeamTagMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamTagMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamTagMemberId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.TeamworkTagId, ok = parsed.Parsed["teamworkTagId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagId", *parsed)
	}

	if id.TeamworkTagMemberId, ok = parsed.Parsed["teamworkTagMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagMemberId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamTagMemberID checks that 'input' can be parsed as a Group Team Tag Member ID
func ValidateGroupTeamTagMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamTagMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Tag Member ID
func (id GroupTeamTagMemberId) ID() string {
	fmtString := "/groups/%s/team/tags/%s/members/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.TeamworkTagId, id.TeamworkTagMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Tag Member ID
func (id GroupTeamTagMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticTags", "tags", "tags"),
		resourceids.UserSpecifiedSegment("teamworkTagId", "teamworkTagIdValue"),
		resourceids.StaticSegment("staticMembers", "members", "members"),
		resourceids.UserSpecifiedSegment("teamworkTagMemberId", "teamworkTagMemberIdValue"),
	}
}

// String returns a human-readable description of this Group Team Tag Member ID
func (id GroupTeamTagMemberId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Teamwork Tag: %q", id.TeamworkTagId),
		fmt.Sprintf("Teamwork Tag Member: %q", id.TeamworkTagMemberId),
	}
	return fmt.Sprintf("Group Team Tag Member (%s)", strings.Join(components, "\n"))
}
