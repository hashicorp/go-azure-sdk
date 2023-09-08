package groupteamchanneltab

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamChannelTabId{}

// GroupTeamChannelTabId is a struct representing the Resource ID for a Group Team Channel Tab
type GroupTeamChannelTabId struct {
	GroupId    string
	ChannelId  string
	TeamsTabId string
}

// NewGroupTeamChannelTabID returns a new GroupTeamChannelTabId struct
func NewGroupTeamChannelTabID(groupId string, channelId string, teamsTabId string) GroupTeamChannelTabId {
	return GroupTeamChannelTabId{
		GroupId:    groupId,
		ChannelId:  channelId,
		TeamsTabId: teamsTabId,
	}
}

// ParseGroupTeamChannelTabID parses 'input' into a GroupTeamChannelTabId
func ParseGroupTeamChannelTabID(input string) (*GroupTeamChannelTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamChannelTabId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamChannelTabId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.TeamsTabId, ok = parsed.Parsed["teamsTabId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamChannelTabIDInsensitively parses 'input' case-insensitively into a GroupTeamChannelTabId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamChannelTabIDInsensitively(input string) (*GroupTeamChannelTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamChannelTabId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamChannelTabId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.TeamsTabId, ok = parsed.Parsed["teamsTabId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamChannelTabID checks that 'input' can be parsed as a Group Team Channel Tab ID
func ValidateGroupTeamChannelTabID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamChannelTabID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Channel Tab ID
func (id GroupTeamChannelTabId) ID() string {
	fmtString := "/groups/%s/team/channels/%s/tabs/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId, id.TeamsTabId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Channel Tab ID
func (id GroupTeamChannelTabId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
		resourceids.StaticSegment("staticTabs", "tabs", "tabs"),
		resourceids.UserSpecifiedSegment("teamsTabId", "teamsTabIdValue"),
	}
}

// String returns a human-readable description of this Group Team Channel Tab ID
func (id GroupTeamChannelTabId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Teams Tab: %q", id.TeamsTabId),
	}
	return fmt.Sprintf("Group Team Channel Tab (%s)", strings.Join(components, "\n"))
}
