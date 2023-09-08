package groupteamchannelmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamChannelMemberId{}

// GroupTeamChannelMemberId is a struct representing the Resource ID for a Group Team Channel Member
type GroupTeamChannelMemberId struct {
	GroupId              string
	ChannelId            string
	ConversationMemberId string
}

// NewGroupTeamChannelMemberID returns a new GroupTeamChannelMemberId struct
func NewGroupTeamChannelMemberID(groupId string, channelId string, conversationMemberId string) GroupTeamChannelMemberId {
	return GroupTeamChannelMemberId{
		GroupId:              groupId,
		ChannelId:            channelId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseGroupTeamChannelMemberID parses 'input' into a GroupTeamChannelMemberId
func ParseGroupTeamChannelMemberID(input string) (*GroupTeamChannelMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamChannelMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamChannelMemberId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamChannelMemberIDInsensitively parses 'input' case-insensitively into a GroupTeamChannelMemberId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamChannelMemberIDInsensitively(input string) (*GroupTeamChannelMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamChannelMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamChannelMemberId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamChannelMemberID checks that 'input' can be parsed as a Group Team Channel Member ID
func ValidateGroupTeamChannelMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamChannelMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Channel Member ID
func (id GroupTeamChannelMemberId) ID() string {
	fmtString := "/groups/%s/team/channels/%s/members/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Channel Member ID
func (id GroupTeamChannelMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
		resourceids.StaticSegment("staticMembers", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberIdValue"),
	}
}

// String returns a human-readable description of this Group Team Channel Member ID
func (id GroupTeamChannelMemberId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Group Team Channel Member (%s)", strings.Join(components, "\n"))
}
