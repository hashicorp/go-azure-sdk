package mejoinedteamchannelmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamChannelMemberId{}

// MeJoinedTeamChannelMemberId is a struct representing the Resource ID for a Me Joined Team Channel Member
type MeJoinedTeamChannelMemberId struct {
	TeamId               string
	ChannelId            string
	ConversationMemberId string
}

// NewMeJoinedTeamChannelMemberID returns a new MeJoinedTeamChannelMemberId struct
func NewMeJoinedTeamChannelMemberID(teamId string, channelId string, conversationMemberId string) MeJoinedTeamChannelMemberId {
	return MeJoinedTeamChannelMemberId{
		TeamId:               teamId,
		ChannelId:            channelId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseMeJoinedTeamChannelMemberID parses 'input' into a MeJoinedTeamChannelMemberId
func ParseMeJoinedTeamChannelMemberID(input string) (*MeJoinedTeamChannelMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamChannelMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamChannelMemberId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamChannelMemberIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamChannelMemberId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamChannelMemberIDInsensitively(input string) (*MeJoinedTeamChannelMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamChannelMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamChannelMemberId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamChannelMemberID checks that 'input' can be parsed as a Me Joined Team Channel Member ID
func ValidateMeJoinedTeamChannelMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamChannelMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Channel Member ID
func (id MeJoinedTeamChannelMemberId) ID() string {
	fmtString := "/me/joinedTeams/%s/channels/%s/members/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Channel Member ID
func (id MeJoinedTeamChannelMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
		resourceids.StaticSegment("staticMembers", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Channel Member ID
func (id MeJoinedTeamChannelMemberId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Me Joined Team Channel Member (%s)", strings.Join(components, "\n"))
}
