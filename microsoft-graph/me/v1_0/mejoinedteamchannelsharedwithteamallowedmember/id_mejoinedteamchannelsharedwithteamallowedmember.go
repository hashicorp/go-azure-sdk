package mejoinedteamchannelsharedwithteamallowedmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamChannelSharedWithTeamAllowedMemberId{}

// MeJoinedTeamChannelSharedWithTeamAllowedMemberId is a struct representing the Resource ID for a Me Joined Team Channel Shared With Team Allowed Member
type MeJoinedTeamChannelSharedWithTeamAllowedMemberId struct {
	TeamId                      string
	ChannelId                   string
	SharedWithChannelTeamInfoId string
	ConversationMemberId        string
}

// NewMeJoinedTeamChannelSharedWithTeamAllowedMemberID returns a new MeJoinedTeamChannelSharedWithTeamAllowedMemberId struct
func NewMeJoinedTeamChannelSharedWithTeamAllowedMemberID(teamId string, channelId string, sharedWithChannelTeamInfoId string, conversationMemberId string) MeJoinedTeamChannelSharedWithTeamAllowedMemberId {
	return MeJoinedTeamChannelSharedWithTeamAllowedMemberId{
		TeamId:                      teamId,
		ChannelId:                   channelId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
		ConversationMemberId:        conversationMemberId,
	}
}

// ParseMeJoinedTeamChannelSharedWithTeamAllowedMemberID parses 'input' into a MeJoinedTeamChannelSharedWithTeamAllowedMemberId
func ParseMeJoinedTeamChannelSharedWithTeamAllowedMemberID(input string) (*MeJoinedTeamChannelSharedWithTeamAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamChannelSharedWithTeamAllowedMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamChannelSharedWithTeamAllowedMemberId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
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

// ParseMeJoinedTeamChannelSharedWithTeamAllowedMemberIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamChannelSharedWithTeamAllowedMemberId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamChannelSharedWithTeamAllowedMemberIDInsensitively(input string) (*MeJoinedTeamChannelSharedWithTeamAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamChannelSharedWithTeamAllowedMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamChannelSharedWithTeamAllowedMemberId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
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

// ValidateMeJoinedTeamChannelSharedWithTeamAllowedMemberID checks that 'input' can be parsed as a Me Joined Team Channel Shared With Team Allowed Member ID
func ValidateMeJoinedTeamChannelSharedWithTeamAllowedMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamChannelSharedWithTeamAllowedMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Channel Shared With Team Allowed Member ID
func (id MeJoinedTeamChannelSharedWithTeamAllowedMemberId) ID() string {
	fmtString := "/me/joinedTeams/%s/channels/%s/sharedWithTeams/%s/allowedMembers/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId, id.SharedWithChannelTeamInfoId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Channel Shared With Team Allowed Member ID
func (id MeJoinedTeamChannelSharedWithTeamAllowedMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
		resourceids.StaticSegment("staticSharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoIdValue"),
		resourceids.StaticSegment("staticAllowedMembers", "allowedMembers", "allowedMembers"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Channel Shared With Team Allowed Member ID
func (id MeJoinedTeamChannelSharedWithTeamAllowedMemberId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Me Joined Team Channel Shared With Team Allowed Member (%s)", strings.Join(components, "\n"))
}
