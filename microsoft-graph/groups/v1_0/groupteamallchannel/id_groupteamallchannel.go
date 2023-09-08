package groupteamallchannel

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamAllChannelId{}

// GroupTeamAllChannelId is a struct representing the Resource ID for a Group Team All Channel
type GroupTeamAllChannelId struct {
	GroupId   string
	ChannelId string
}

// NewGroupTeamAllChannelID returns a new GroupTeamAllChannelId struct
func NewGroupTeamAllChannelID(groupId string, channelId string) GroupTeamAllChannelId {
	return GroupTeamAllChannelId{
		GroupId:   groupId,
		ChannelId: channelId,
	}
}

// ParseGroupTeamAllChannelID parses 'input' into a GroupTeamAllChannelId
func ParseGroupTeamAllChannelID(input string) (*GroupTeamAllChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamAllChannelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamAllChannelId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamAllChannelIDInsensitively parses 'input' case-insensitively into a GroupTeamAllChannelId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamAllChannelIDInsensitively(input string) (*GroupTeamAllChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamAllChannelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamAllChannelId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamAllChannelID checks that 'input' can be parsed as a Group Team All Channel ID
func ValidateGroupTeamAllChannelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamAllChannelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team All Channel ID
func (id GroupTeamAllChannelId) ID() string {
	fmtString := "/groups/%s/team/allChannels/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team All Channel ID
func (id GroupTeamAllChannelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticAllChannels", "allChannels", "allChannels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
	}
}

// String returns a human-readable description of this Group Team All Channel ID
func (id GroupTeamAllChannelId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
	}
	return fmt.Sprintf("Group Team All Channel (%s)", strings.Join(components, "\n"))
}
