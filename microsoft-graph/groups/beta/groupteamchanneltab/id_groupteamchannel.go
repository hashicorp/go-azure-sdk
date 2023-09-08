package groupteamchanneltab

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamChannelId{}

// GroupTeamChannelId is a struct representing the Resource ID for a Group Team Channel
type GroupTeamChannelId struct {
	GroupId   string
	ChannelId string
}

// NewGroupTeamChannelID returns a new GroupTeamChannelId struct
func NewGroupTeamChannelID(groupId string, channelId string) GroupTeamChannelId {
	return GroupTeamChannelId{
		GroupId:   groupId,
		ChannelId: channelId,
	}
}

// ParseGroupTeamChannelID parses 'input' into a GroupTeamChannelId
func ParseGroupTeamChannelID(input string) (*GroupTeamChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamChannelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamChannelId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamChannelIDInsensitively parses 'input' case-insensitively into a GroupTeamChannelId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamChannelIDInsensitively(input string) (*GroupTeamChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamChannelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamChannelId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamChannelID checks that 'input' can be parsed as a Group Team Channel ID
func ValidateGroupTeamChannelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamChannelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Channel ID
func (id GroupTeamChannelId) ID() string {
	fmtString := "/groups/%s/team/channels/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Channel ID
func (id GroupTeamChannelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
	}
}

// String returns a human-readable description of this Group Team Channel ID
func (id GroupTeamChannelId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
	}
	return fmt.Sprintf("Group Team Channel (%s)", strings.Join(components, "\n"))
}
