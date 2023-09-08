package groupteamtag

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamTagId{}

// GroupTeamTagId is a struct representing the Resource ID for a Group Team Tag
type GroupTeamTagId struct {
	GroupId       string
	TeamworkTagId string
}

// NewGroupTeamTagID returns a new GroupTeamTagId struct
func NewGroupTeamTagID(groupId string, teamworkTagId string) GroupTeamTagId {
	return GroupTeamTagId{
		GroupId:       groupId,
		TeamworkTagId: teamworkTagId,
	}
}

// ParseGroupTeamTagID parses 'input' into a GroupTeamTagId
func ParseGroupTeamTagID(input string) (*GroupTeamTagId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamTagId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamTagId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.TeamworkTagId, ok = parsed.Parsed["teamworkTagId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamTagIDInsensitively parses 'input' case-insensitively into a GroupTeamTagId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamTagIDInsensitively(input string) (*GroupTeamTagId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamTagId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamTagId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.TeamworkTagId, ok = parsed.Parsed["teamworkTagId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamTagID checks that 'input' can be parsed as a Group Team Tag ID
func ValidateGroupTeamTagID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamTagID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Tag ID
func (id GroupTeamTagId) ID() string {
	fmtString := "/groups/%s/team/tags/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.TeamworkTagId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Tag ID
func (id GroupTeamTagId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticTags", "tags", "tags"),
		resourceids.UserSpecifiedSegment("teamworkTagId", "teamworkTagIdValue"),
	}
}

// String returns a human-readable description of this Group Team Tag ID
func (id GroupTeamTagId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Teamwork Tag: %q", id.TeamworkTagId),
	}
	return fmt.Sprintf("Group Team Tag (%s)", strings.Join(components, "\n"))
}
