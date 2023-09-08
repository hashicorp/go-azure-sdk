package groupteamchannelsharedwithteamallowedmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamChannelSharedWithTeamId{}

// GroupTeamChannelSharedWithTeamId is a struct representing the Resource ID for a Group Team Channel Shared With Team
type GroupTeamChannelSharedWithTeamId struct {
	GroupId                     string
	ChannelId                   string
	SharedWithChannelTeamInfoId string
}

// NewGroupTeamChannelSharedWithTeamID returns a new GroupTeamChannelSharedWithTeamId struct
func NewGroupTeamChannelSharedWithTeamID(groupId string, channelId string, sharedWithChannelTeamInfoId string) GroupTeamChannelSharedWithTeamId {
	return GroupTeamChannelSharedWithTeamId{
		GroupId:                     groupId,
		ChannelId:                   channelId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
	}
}

// ParseGroupTeamChannelSharedWithTeamID parses 'input' into a GroupTeamChannelSharedWithTeamId
func ParseGroupTeamChannelSharedWithTeamID(input string) (*GroupTeamChannelSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamChannelSharedWithTeamId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamChannelSharedWithTeamId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamChannelSharedWithTeamIDInsensitively parses 'input' case-insensitively into a GroupTeamChannelSharedWithTeamId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamChannelSharedWithTeamIDInsensitively(input string) (*GroupTeamChannelSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamChannelSharedWithTeamId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamChannelSharedWithTeamId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamChannelSharedWithTeamID checks that 'input' can be parsed as a Group Team Channel Shared With Team ID
func ValidateGroupTeamChannelSharedWithTeamID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamChannelSharedWithTeamID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Channel Shared With Team ID
func (id GroupTeamChannelSharedWithTeamId) ID() string {
	fmtString := "/groups/%s/team/channels/%s/sharedWithTeams/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId, id.SharedWithChannelTeamInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Channel Shared With Team ID
func (id GroupTeamChannelSharedWithTeamId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
		resourceids.StaticSegment("staticSharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoIdValue"),
	}
}

// String returns a human-readable description of this Group Team Channel Shared With Team ID
func (id GroupTeamChannelSharedWithTeamId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
	}
	return fmt.Sprintf("Group Team Channel Shared With Team (%s)", strings.Join(components, "\n"))
}
