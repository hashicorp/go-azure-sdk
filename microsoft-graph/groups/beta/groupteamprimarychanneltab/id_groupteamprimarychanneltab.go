package groupteamprimarychanneltab

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamPrimaryChannelTabId{}

// GroupTeamPrimaryChannelTabId is a struct representing the Resource ID for a Group Team Primary Channel Tab
type GroupTeamPrimaryChannelTabId struct {
	GroupId    string
	TeamsTabId string
}

// NewGroupTeamPrimaryChannelTabID returns a new GroupTeamPrimaryChannelTabId struct
func NewGroupTeamPrimaryChannelTabID(groupId string, teamsTabId string) GroupTeamPrimaryChannelTabId {
	return GroupTeamPrimaryChannelTabId{
		GroupId:    groupId,
		TeamsTabId: teamsTabId,
	}
}

// ParseGroupTeamPrimaryChannelTabID parses 'input' into a GroupTeamPrimaryChannelTabId
func ParseGroupTeamPrimaryChannelTabID(input string) (*GroupTeamPrimaryChannelTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamPrimaryChannelTabId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamPrimaryChannelTabId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.TeamsTabId, ok = parsed.Parsed["teamsTabId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamPrimaryChannelTabIDInsensitively parses 'input' case-insensitively into a GroupTeamPrimaryChannelTabId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamPrimaryChannelTabIDInsensitively(input string) (*GroupTeamPrimaryChannelTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamPrimaryChannelTabId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamPrimaryChannelTabId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.TeamsTabId, ok = parsed.Parsed["teamsTabId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamPrimaryChannelTabID checks that 'input' can be parsed as a Group Team Primary Channel Tab ID
func ValidateGroupTeamPrimaryChannelTabID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamPrimaryChannelTabID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Primary Channel Tab ID
func (id GroupTeamPrimaryChannelTabId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/tabs/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.TeamsTabId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Primary Channel Tab ID
func (id GroupTeamPrimaryChannelTabId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticTabs", "tabs", "tabs"),
		resourceids.UserSpecifiedSegment("teamsTabId", "teamsTabIdValue"),
	}
}

// String returns a human-readable description of this Group Team Primary Channel Tab ID
func (id GroupTeamPrimaryChannelTabId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Teams Tab: %q", id.TeamsTabId),
	}
	return fmt.Sprintf("Group Team Primary Channel Tab (%s)", strings.Join(components, "\n"))
}
