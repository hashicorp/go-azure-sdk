package groupteamincomingchannel

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamIncomingChannelId{}

// GroupTeamIncomingChannelId is a struct representing the Resource ID for a Group Team Incoming Channel
type GroupTeamIncomingChannelId struct {
	GroupId   string
	ChannelId string
}

// NewGroupTeamIncomingChannelID returns a new GroupTeamIncomingChannelId struct
func NewGroupTeamIncomingChannelID(groupId string, channelId string) GroupTeamIncomingChannelId {
	return GroupTeamIncomingChannelId{
		GroupId:   groupId,
		ChannelId: channelId,
	}
}

// ParseGroupTeamIncomingChannelID parses 'input' into a GroupTeamIncomingChannelId
func ParseGroupTeamIncomingChannelID(input string) (*GroupTeamIncomingChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamIncomingChannelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamIncomingChannelId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamIncomingChannelIDInsensitively parses 'input' case-insensitively into a GroupTeamIncomingChannelId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamIncomingChannelIDInsensitively(input string) (*GroupTeamIncomingChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamIncomingChannelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamIncomingChannelId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamIncomingChannelID checks that 'input' can be parsed as a Group Team Incoming Channel ID
func ValidateGroupTeamIncomingChannelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamIncomingChannelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Incoming Channel ID
func (id GroupTeamIncomingChannelId) ID() string {
	fmtString := "/groups/%s/team/incomingChannels/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Incoming Channel ID
func (id GroupTeamIncomingChannelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticIncomingChannels", "incomingChannels", "incomingChannels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
	}
}

// String returns a human-readable description of this Group Team Incoming Channel ID
func (id GroupTeamIncomingChannelId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
	}
	return fmt.Sprintf("Group Team Incoming Channel (%s)", strings.Join(components, "\n"))
}
