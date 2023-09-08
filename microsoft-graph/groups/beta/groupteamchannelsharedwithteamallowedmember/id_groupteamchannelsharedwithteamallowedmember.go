package groupteamchannelsharedwithteamallowedmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamChannelSharedWithTeamAllowedMemberId{}

// GroupTeamChannelSharedWithTeamAllowedMemberId is a struct representing the Resource ID for a Group Team Channel Shared With Team Allowed Member
type GroupTeamChannelSharedWithTeamAllowedMemberId struct {
	GroupId                     string
	ChannelId                   string
	SharedWithChannelTeamInfoId string
	ConversationMemberId        string
}

// NewGroupTeamChannelSharedWithTeamAllowedMemberID returns a new GroupTeamChannelSharedWithTeamAllowedMemberId struct
func NewGroupTeamChannelSharedWithTeamAllowedMemberID(groupId string, channelId string, sharedWithChannelTeamInfoId string, conversationMemberId string) GroupTeamChannelSharedWithTeamAllowedMemberId {
	return GroupTeamChannelSharedWithTeamAllowedMemberId{
		GroupId:                     groupId,
		ChannelId:                   channelId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
		ConversationMemberId:        conversationMemberId,
	}
}

// ParseGroupTeamChannelSharedWithTeamAllowedMemberID parses 'input' into a GroupTeamChannelSharedWithTeamAllowedMemberId
func ParseGroupTeamChannelSharedWithTeamAllowedMemberID(input string) (*GroupTeamChannelSharedWithTeamAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamChannelSharedWithTeamAllowedMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamChannelSharedWithTeamAllowedMemberId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamChannelSharedWithTeamAllowedMemberIDInsensitively parses 'input' case-insensitively into a GroupTeamChannelSharedWithTeamAllowedMemberId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamChannelSharedWithTeamAllowedMemberIDInsensitively(input string) (*GroupTeamChannelSharedWithTeamAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamChannelSharedWithTeamAllowedMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamChannelSharedWithTeamAllowedMemberId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamChannelSharedWithTeamAllowedMemberID checks that 'input' can be parsed as a Group Team Channel Shared With Team Allowed Member ID
func ValidateGroupTeamChannelSharedWithTeamAllowedMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamChannelSharedWithTeamAllowedMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Channel Shared With Team Allowed Member ID
func (id GroupTeamChannelSharedWithTeamAllowedMemberId) ID() string {
	fmtString := "/groups/%s/team/channels/%s/sharedWithTeams/%s/allowedMembers/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId, id.SharedWithChannelTeamInfoId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Channel Shared With Team Allowed Member ID
func (id GroupTeamChannelSharedWithTeamAllowedMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
		resourceids.StaticSegment("staticSharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoIdValue"),
		resourceids.StaticSegment("staticAllowedMembers", "allowedMembers", "allowedMembers"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberIdValue"),
	}
}

// String returns a human-readable description of this Group Team Channel Shared With Team Allowed Member ID
func (id GroupTeamChannelSharedWithTeamAllowedMemberId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Group Team Channel Shared With Team Allowed Member (%s)", strings.Join(components, "\n"))
}
